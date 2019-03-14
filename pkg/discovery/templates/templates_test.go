package templates

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"testing"

	"bitbucket.org/openbankingteam/conformance-suite/pkg/discovery"
	"bitbucket.org/openbankingteam/conformance-suite/pkg/model"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Test that all the *.json discovery files parse correctly.
func TestDiscoverySamples_Examples_Parse_Correctly(t *testing.T) {
	discoveryFiles, err := filepath.Glob("*.json")
	require.NoError(t, err)

	for _, discoveryFile := range discoveryFiles {
		t.Run("Parses_Without_Error_"+discoveryFile, func(t *testing.T) {
			assert := assert.New(t)

			discoveryJSON, err := ioutil.ReadFile(discoveryFile)
			assert.NoError(err)
			assert.NotNil(discoveryJSON)

			discoveryModel := &discovery.Model{}
			assert.NoError(json.Unmarshal(discoveryJSON, &discoveryModel))

			checker := model.NewConditionalityChecker()
			result, failures, err := discovery.Validate(checker, discoveryModel)
			require.NoError(t, err)
			require.Empty(t, failures)
			assert.True(result)
		})
	}
}

// TestDiscoverySamplesIfManifestIsURLHTTPSOnly Asserts that if the manifest field is populated as a URL
// then it must use the HTTPS scheme.
func TestDiscoverySamplesIfManifestIsURLHTTPSOnly(t *testing.T) {
	discoveryFiles, err := filepath.Glob("*.json")
	require.NoError(t, err)

	for _, discoveryFile := range discoveryFiles {
		t.Run("Parses_Without_Error_"+discoveryFile, func(t *testing.T) {
			assert := assert.New(t)

			discoveryJSON, err := ioutil.ReadFile(discoveryFile)
			assert.NoError(err)
			assert.NotNil(discoveryJSON)

			discoveryModel := &discovery.Model{}
			assert.NoError(json.Unmarshal(discoveryJSON, &discoveryModel))

			checker := model.NewConditionalityChecker()
			result, failures, err := discovery.Validate(checker, discoveryModel)
			require.NoError(t, err)
			require.Empty(t, failures)
			assert.True(result)
		})
	}
}
