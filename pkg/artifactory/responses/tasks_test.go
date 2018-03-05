package responses

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/lusis/go-artifactory/pkg/artifactory/responses/testdata"
	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/require"
)

func TestTaskResponses(t *testing.T) {
	// Add any new test cases here
	// in the form of
	// {&FooResponse{}:FooResponseTestData}
	taskTestCases := []map[VersionedResponse]string{
		{&GetBackgroundTasksResponse{}: GetBackgroundTasksResponseTestData},
	}
	t.Parallel()
	for _, testCase := range taskTestCases {
		for k, v := range testCase {
			t.Run(v,
				func(*testing.T) {
					data, err := testdata.GetBytes(v)
					require.NoError(t, err)
					placeholder := make(map[string]interface{})
					_ = json.Unmarshal(data, &placeholder)
					config := newMSDecoderConfig()
					config.Result = k
					decoder, newErr := mapstructure.NewDecoder(config)
					require.NoError(t, newErr)
					dErr := decoder.Decode(placeholder)
					require.NoError(t, dErr, fmt.Sprintf("should parse %s", v))
					require.NotNil(t, k)
					require.Implements(t, (*VersionedResponse)(nil), k)
				})
		}
	}
}
