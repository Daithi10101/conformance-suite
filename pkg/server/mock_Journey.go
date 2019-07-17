// Code generated by mockery v1.0.0. DO NOT EDIT.

package server

import discovery "bitbucket.org/openbankingteam/conformance-suite/pkg/discovery"
import events "bitbucket.org/openbankingteam/conformance-suite/pkg/executors/events"
import executors "bitbucket.org/openbankingteam/conformance-suite/pkg/executors"
import generation "bitbucket.org/openbankingteam/conformance-suite/pkg/generation"
import manifest "bitbucket.org/openbankingteam/conformance-suite/pkg/manifest"
import mock "github.com/stretchr/testify/mock"

// MockJourney is an autogenerated mock type for the Journey type
type MockJourney struct {
	mock.Mock
}

// AllTokenCollected provides a mock function with given fields:
func (_m *MockJourney) AllTokenCollected() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// CollectToken provides a mock function with given fields: code, state, scope
func (_m *MockJourney) CollectToken(code string, state string, scope string) error {
	ret := _m.Called(code, state, scope)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string) error); ok {
		r0 = rf(code, state, scope)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DiscoveryModel provides a mock function with given fields:
func (_m *MockJourney) DiscoveryModel() (discovery.Model, error) {
	ret := _m.Called()

	var r0 discovery.Model
	if rf, ok := ret.Get(0).(func() discovery.Model); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(discovery.Model)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Events provides a mock function with given fields:
func (_m *MockJourney) Events() events.Events {
	ret := _m.Called()

	var r0 events.Events
	if rf, ok := ret.Get(0).(func() events.Events); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(events.Events)
		}
	}

	return r0
}

// FilteredManifests provides a mock function with given fields:
func (_m *MockJourney) FilteredManifests() (manifest.Scripts, error) {
	ret := _m.Called()

	var r0 manifest.Scripts
	if rf, ok := ret.Get(0).(func() manifest.Scripts); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(manifest.Scripts)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewDaemonController provides a mock function with given fields:
func (_m *MockJourney) NewDaemonController() {
	_m.Called()
}

// Results provides a mock function with given fields:
func (_m *MockJourney) Results() executors.DaemonController {
	ret := _m.Called()

	var r0 executors.DaemonController
	if rf, ok := ret.Get(0).(func() executors.DaemonController); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(executors.DaemonController)
		}
	}

	return r0
}

// RunTests provides a mock function with given fields:
func (_m *MockJourney) RunTests() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetConfig provides a mock function with given fields: config
func (_m *MockJourney) SetConfig(config JourneyConfig) error {
	ret := _m.Called(config)

	var r0 error
	if rf, ok := ret.Get(0).(func(JourneyConfig) error); ok {
		r0 = rf(config)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetDiscoveryModel provides a mock function with given fields: discoveryModel
func (_m *MockJourney) SetDiscoveryModel(discoveryModel *discovery.Model) (discovery.ValidationFailures, error) {
	ret := _m.Called(discoveryModel)

	var r0 discovery.ValidationFailures
	if rf, ok := ret.Get(0).(func(*discovery.Model) discovery.ValidationFailures); ok {
		r0 = rf(discoveryModel)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(discovery.ValidationFailures)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*discovery.Model) error); ok {
		r1 = rf(discoveryModel)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetFilteredManifests provides a mock function with given fields: _a0
func (_m *MockJourney) SetFilteredManifests(_a0 manifest.Scripts) {
	_m.Called(_a0)
}

// StopTestRun provides a mock function with given fields:
func (_m *MockJourney) StopTestRun() {
	_m.Called()
}

// TLSVersionResult provides a mock function with given fields:
func (_m *MockJourney) TLSVersionResult() map[string]*discovery.TLSValidationResult {
	ret := _m.Called()

	var r0 map[string]*discovery.TLSValidationResult
	if rf, ok := ret.Get(0).(func() map[string]*discovery.TLSValidationResult); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]*discovery.TLSValidationResult)
		}
	}

	return r0
}

// TestCases provides a mock function with given fields:
func (_m *MockJourney) TestCases() (generation.TestCasesRun, error) {
	ret := _m.Called()

	var r0 generation.TestCasesRun
	if rf, ok := ret.Get(0).(func() generation.TestCasesRun); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(generation.TestCasesRun)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
