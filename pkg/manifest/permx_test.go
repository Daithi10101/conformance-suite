package manifest

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"bitbucket.org/openbankingteam/conformance-suite/pkg/model"
)

const manifestPath = "file://manifests/ob_3.1_payment_fca.json"

func TestPermx(t *testing.T) {
	tests, err := GenerateTestCases(accountSwaggerLocation31, "http://mybaseurl", &model.Context{}, readDiscovery(), manifestPath)
	assert.Nil(t, err)
	testcasePermissions, err := getTestCasePermissions(tests)
	assert.Nil(t, err)
	requiredTokens, err := getRequiredTokens(testcasePermissions)
	assert.Nil(t, err)
	dumpJSON(requiredTokens)
}

func TestGetScriptConsentTokens(t *testing.T) {
	tests, err := GenerateTestCases(accountSwaggerLocation31, "http://mybaseurl", &model.Context{}, readDiscovery(), manifestPath)
	assert.Nil(t, err)
	testcasePermissions, err := getTestCasePermissions(tests)
	assert.Nil(t, err)
	requiredTokens, err := getRequiredTokens(testcasePermissions)
	assert.Nil(t, err)
	populateTokens(t, requiredTokens)
	dumpJSON(requiredTokens)
}

func populateTokens(t *testing.T, gatherer []RequiredTokens) error {
	t.Helper()

	t.Logf("%d entries found\n", len(gatherer))
	for k, tokenGatherer := range gatherer {
		if len(tokenGatherer.Perms) == 0 {
			continue
		}
		token, err := getToken(tokenGatherer.Perms)
		if err != nil {
			return err
		}
		tokenGatherer.Token = token
		gatherer[k] = tokenGatherer

	}
	return nil
}

func getToken(perms []string) (string, error) {
	// for headless - get the okens for the permissions here
	return "abigfattoken", nil
}
