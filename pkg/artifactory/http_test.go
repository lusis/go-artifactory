package artifactory

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/lusis/go-artifactory/pkg/artifactory/responses"
	"github.com/lusis/go-artifactory/pkg/artifactory/responses/testdata"
	"github.com/lusis/go-artifactory/pkg/httpclient"
	"github.com/stretchr/testify/require"
)

func TestParseBodyToError(t *testing.T) {
	data, err := testdata.GetBytes(responses.ErrorResponseTestData)
	require.NoError(t, err)
	err = parseErrorBodytoError(data)
	require.Error(t, err, "should parse error properly")
	require.Contains(t, err.Error(), "Artifactory error", "should contain artifactory error")
	err = parseErrorBodytoError([]byte("plain string"))
	require.Error(t, err, "should not parse error as json")
	require.Contains(t, err.Error(), "Error parsing response body: plain string", "should contain response body")
}

func TestParseHTTPresp(t *testing.T) {
	jsonHeaders := http.Header{}
	jsonHeaders.Add("Content-Type", contentTypeJSON)
	testCases := map[string]struct {
		response *httpclient.Response
		e        error
		fail     bool
		contains string
	}{
		"explicit error": {
			response: &httpclient.Response{},
			e:        fmt.Errorf("this is an error"),
			fail:     true,
			contains: "this is an error",
		},
		"status >= 400 no body": {
			response: &httpclient.Response{
				Status: 404,
			},
			e:        nil,
			fail:     true,
			contains: "Got an error but no body",
		},
		"status >= 400 with body not json": {
			response: &httpclient.Response{
				Status: 404,
				Body:   []byte("my body"),
			},
			e:        nil,
			fail:     true,
			contains: "Body: my body",
		},
		"status >= 400 with body not json with json header": {
			response: &httpclient.Response{
				Status:  404,
				Body:    []byte("my body"),
				Headers: jsonHeaders,
			},
			e:        nil,
			fail:     true,
			contains: "Error parsing response body",
		},
		"status good with body": {
			response: &httpclient.Response{
				Status: 200,
				Body:   []byte("good body"),
			},
			e:        nil,
			fail:     false,
			contains: "good body",
		},
	}
	t.Parallel()
	for k, v := range testCases {
		t.Run(k,
			func(*testing.T) {
				data, err := parseHTTPresp(v.response, v.e)
				if v.fail {
					require.Nil(t, data, "data should be nil: "+k)
					require.Error(t, err, "should be an error: "+k)
					require.Contains(t, err.Error(), v.contains, "error should match: "+k)
					return
				}
				require.NotNil(t, data, "data should not be nil: "+k)
				require.Equal(t, v.contains, string(data), "body should match: "+k)
				require.NoError(t, err, "should not be an error: "+k)
			})
	}
}
func TestHTTPErrors(t *testing.T) {
	data, err := testdata.GetBytes(responses.ErrorResponseTestData)
	require.NoError(t, err)
	client, server, err := newTestClient(data, "application/json", 404)
	defer server.Close()
	require.NoError(t, err)
	_, err = client.httpGet("/foo")
	require.Error(t, err, "should return an error")
	require.Contains(t, err.Error(), "Artifactory error", "should contain artifactory error")
}
