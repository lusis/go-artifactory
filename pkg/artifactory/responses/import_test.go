package responses

import (
	"encoding/json"
	"testing"

	"github.com/lusis/go-artifactory/pkg/artifactory/responses/testdata"
	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/require"
)

func TestImportResponses(t *testing.T) {
	// Add any new test cases here
	// in the form of
	// {&FooResponse{}:FooResponseTestData}
	systemResponsesTestCases := []map[VersionedResponse]string{
		{&ImportSystemSettingsResponse{}: ImportSystemSettingsResponseTestData},
	}

	for _, testCase := range systemResponsesTestCases {
		for k, v := range testCase {
			data, err := testdata.GetBytes(v)
			require.NoError(t, err)
			placeholder := make(map[string]interface{})
			_ = json.Unmarshal(data, &placeholder)
			config := newMSDecoderConfig()
			config.Result = k
			decoder, newErr := mapstructure.NewDecoder(config)
			require.NoError(t, newErr)
			dErr := decoder.Decode(placeholder)
			require.NoError(t, dErr)
			require.Implements(t, (*VersionedResponse)(nil), k)
		}
	}
}
