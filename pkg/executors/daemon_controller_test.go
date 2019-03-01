package executors

import (
	"testing"

	"bitbucket.org/openbankingteam/conformance-suite/pkg/executors/results"
	"github.com/stretchr/testify/assert"
)

func TestNewDaemonController(t *testing.T) {
	testChan := make(chan results.TestCase, 100)

	controller := NewDaemonController(testChan)

	assert.Equal(t, testChan, controller.Results())
	assert.NotNil(t, controller.stopLock)
	assert.False(t, controller.shouldStop)
	assert.False(t, controller.ShouldStop())
}

func TestNewBufferedDaemonController(t *testing.T) {
	controller := NewBufferedDaemonController()

	assert.NotNil(t, controller.Results())
	assert.NotNil(t, controller.stopLock)
	assert.False(t, controller.shouldStop)
	assert.False(t, controller.ShouldStop())
}

func TestDaemonControllerStops(t *testing.T) {
	testChan := make(chan results.TestCase, 100)
	controller := NewDaemonController(testChan)

	controller.Stop()

	assert.True(t, controller.shouldStop)
	assert.True(t, controller.ShouldStop())
}
