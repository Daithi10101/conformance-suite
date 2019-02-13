package server

import (
	"sync"

	"bitbucket.org/openbankingteam/conformance-suite/pkg/authentication"
	"bitbucket.org/openbankingteam/conformance-suite/pkg/discovery"
	"bitbucket.org/openbankingteam/conformance-suite/pkg/executors"
	"bitbucket.org/openbankingteam/conformance-suite/pkg/generation"
	"bitbucket.org/openbankingteam/conformance-suite/pkg/model"
	"github.com/pkg/errors"
)

var (
	errDiscoveryModelNotSet        = errors.New("error discovery model not set")
	errTestCasesNotGenerated       = errors.New("error test cases not generated")
	errNotFinishedCollectingTokens = errors.New("error not finished collecting tokens")
)

// Journey represents all possible steps for a user test conformance journey
//
// Happy path journey would look like:
// 1. SetCertificates - sets configuration to run test cases
// 2. SetDiscoveryModel - this validates and if successful set this as your discovery model
// 3. TestCases - Generates test cases, generates permission set requirements to run tests and starts a token collector
// 3.1 CollectToken - collects all tokens required to RunTest
// 4. RunTest - Runs triggers a background run on all generated test from previous steps, needs all token to be already collected
// 5. Results - returns a background process control, so we can monitor on finished tests
//
type Journey interface {
	SetDiscoveryModel(discoveryModel *discovery.Model) (discovery.ValidationFailures, error)
	TestCases() (generation.TestCasesRun, error)
	CollectToken(setName, token string) error
	AllTokenCollected() bool
	RunTests() error
	StopTestRun()
	Results() executors.DaemonController
	SetCertificates(signing, transport authentication.Certificate)
}

type journey struct {
	generator        generation.Generator
	validator        discovery.Validator
	daemonController executors.DaemonController

	journeyLock           *sync.Mutex
	testCasesRun          generation.TestCasesRun
	testCasesRunGenerated bool
	collector             executors.Collector
	allCollected          bool
	validDiscoveryModel   *discovery.Model
	certificateSigning    authentication.Certificate
	certificateTransport  authentication.Certificate
	context               model.Context
}

// NewJourney creates an instance for a user journey
func NewJourney(generator generation.Generator, validator discovery.Validator) *journey {
	return &journey{
		generator:             generator,
		validator:             validator,
		daemonController:      executors.NewBufferedDaemonController(),
		journeyLock:           &sync.Mutex{},
		allCollected:          false,
		testCasesRunGenerated: false,
		context:               model.Context{},
	}
}

func (wj *journey) SetDiscoveryModel(discoveryModel *discovery.Model) (discovery.ValidationFailures, error) {
	failures, err := wj.validator.Validate(discoveryModel)
	if err != nil {
		return nil, errors.Wrap(err, "error setting discovery model")
	}

	if !failures.Empty() {
		return failures, nil
	}

	wj.journeyLock.Lock()
	wj.validDiscoveryModel = discoveryModel
	wj.testCasesRunGenerated = false
	wj.allCollected = false
	wj.journeyLock.Unlock()

	return discovery.NoValidationFailures, nil
}

func (wj *journey) TestCases() (generation.TestCasesRun, error) {
	wj.journeyLock.Lock()
	defer wj.journeyLock.Unlock()

	if wj.validDiscoveryModel == nil {
		return generation.TestCasesRun{}, errDiscoveryModelNotSet
	}

	if !wj.testCasesRunGenerated {
		wj.testCasesRun = wj.generator.GenerateSpecificationTestCases(wj.validDiscoveryModel.DiscoveryModel)
		// replace this with a NewCollector to a real implementation
		runDefinition := executors.RunDefinition{
			DiscoModel:    wj.validDiscoveryModel,
			TestCaseRun:   wj.testCasesRun,
			SigningCert:   wj.certificateSigning,
			TransportCert: wj.certificateTransport,
			Context:       &wj.context,
		}
		wj.customTestParametersToJourneyContext()
		if wj.validDiscoveryModel.DiscoveryModel.TokenAcquisition == "psu" {
			executors.InitiationConsentAcquisition(wj.testCasesRun.SpecConsentRequirements, runDefinition)
			wj.collector = executors.NewNullCollector(wj.doneCollectionCallback)
		}
		wj.testCasesRunGenerated = true
		wj.allCollected = false
	}

	return wj.testCasesRun, nil
}

func (wj *journey) CollectToken(setName, token string) error {
	wj.journeyLock.Lock()
	defer wj.journeyLock.Unlock()

	if !wj.testCasesRunGenerated {
		return errTestCasesNotGenerated
	}

	return wj.collector.Collect(setName, token)
}

func (wj *journey) AllTokenCollected() bool {
	wj.journeyLock.Lock()
	defer wj.journeyLock.Unlock()

	return wj.allCollected
}

func (wj *journey) doneCollectionCallback() {
	wj.journeyLock.Lock()
	wj.allCollected = true
	wj.journeyLock.Unlock()
}

func (wj *journey) RunTests() error {
	wj.journeyLock.Lock()
	defer wj.journeyLock.Unlock()

	if !wj.testCasesRunGenerated {
		return errTestCasesNotGenerated
	}

	if !wj.allCollected {
		return errNotFinishedCollectingTokens
	}

	runDefinition := executors.RunDefinition{
		DiscoModel:    wj.validDiscoveryModel,
		TestCaseRun:   wj.testCasesRun,
		SigningCert:   wj.certificateSigning,
		TransportCert: wj.certificateTransport,
		Context:       &wj.context,
	}

	runner := executors.NewTestCaseRunner(runDefinition, wj.daemonController)
	return runner.RunTestCases()
}

func (wj *journey) Results() executors.DaemonController {
	return wj.daemonController
}

func (wj *journey) StopTestRun() {
	wj.daemonController.Stop()
}

func (wj *journey) SetCertificates(signing, transport authentication.Certificate) {
	wj.journeyLock.Lock()
	wj.certificateSigning = signing
	wj.certificateTransport = transport
	wj.journeyLock.Unlock()
}

func (wj *journey) customTestParametersToJourneyContext() {
	for _, customTest := range wj.validDiscoveryModel.DiscoveryModel.CustomTests { // assume ordering is prerun i.e. customtest run before other tests
		for k, v := range customTest.Replacements {
			wj.context.PutString(k, v)
		}
	}
}
