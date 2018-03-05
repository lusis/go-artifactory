package artifactory

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/stretchr/testify/require"
)

func TestNewClientCustomTransport(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, "pong")
	}))
	transport := &http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			return url.Parse(server.URL)
		},
	}
	httpClient := &http.Client{Transport: transport}
	defer server.Close()
	conf := &ClientConfig{
		BaseURL:    "http://127.0.0.1:8080/",
		Username:   "username",
		Password:   "password",
		VerifySSL:  false,
		HTTPClient: httpClient,
	}

	client, err := NewClient(conf)
	require.NoError(t, err)
	res, err := client.httpGet("/ping")
	require.NoError(t, err)
	require.Equal(t, "pong", string(res))
}

func TestClientHTTPVerifySSLTrue(t *testing.T) {
	conf := &ClientConfig{VerifySSL: true}
	client, err := NewClient(conf)
	require.NoError(t, err)
	require.False(t, client.HTTPClient.Transport.(*http.Transport).TLSClientConfig.InsecureSkipVerify)
}

func TestClientHTTPVerifySSLFalse(t *testing.T) {
	conf := &ClientConfig{VerifySSL: false}
	client, err := NewClient(conf)
	require.NoError(t, err)
	require.True(t, client.HTTPClient.Transport.(*http.Transport).TLSClientConfig.InsecureSkipVerify)
}

func TestClientFromEnvWithBasicAuth(t *testing.T) {
	testStashEnvVars(t)
	os.Setenv("ARTIFACTORY_URL", "http://artifactory.server.com") //nolint
	os.Setenv("ARTIFACTORY_USERNAME", "admin")                    //nolint
	os.Setenv("ARTIFACTORY_PASSWORD", "password")                 //nolint
	os.Setenv("ARTIFACTORY_TOKEN", "")                            //nolint
	defer func() {
		testUnStashEnvVars(t)
	}()
	client, err := NewClientFromEnv()
	require.Nil(t, err)
	require.NotNil(t, client)
	require.Equal(t, "http://artifactory.server.com", client.Config.BaseURL)
	require.Equal(t, "basic", client.Config.AuthMethod)
	require.Equal(t, "admin", client.Config.Username)
	require.Equal(t, "password", client.Config.Password)
}

func TestClientFromEnvWithTokenAuth(t *testing.T) {
	testStashEnvVars(t)
	os.Setenv("ARTIFACTORY_URL", "http://artifactory.server.com") //nolint
	os.Setenv("ARTIFACTORY_TOKEN", "someToken")                   //nolint
	defer func() {
		testUnStashEnvVars(t)
	}()
	client, err := NewClientFromEnv()
	require.NotNil(t, client)
	require.Nil(t, err)
	require.Equal(t, "http://artifactory.server.com", client.Config.BaseURL)
	require.Equal(t, "token", client.Config.AuthMethod)
	require.Equal(t, "someToken", client.Config.Token)
}

func TestClientWithDebug(t *testing.T) {
	testStashEnvVars(t)
	os.Setenv("ARTIFACTORY_URL", "http://artifactory.server.com") //nolint
	os.Setenv("ARTIFACTORY_TOKEN", "someToken")                   //nolint
	os.Setenv("ARTIFACTORY_DEBUG", "1")                           // nolint
	defer func() {
		testUnStashEnvVars(t)
	}()
	client, err := NewClientFromEnv()
	assert.NoError(t, err)
	assert.True(t, client.Config.Debug)
}

func TestClientWithVersion(t *testing.T) {
	testStashEnvVars(t)
	os.Setenv("ARTIFACTORY_URL", "http://artifactory.server.com") //nolint
	os.Setenv("ARTIFACTORY_TOKEN", "someToken")                   //nolint
	os.Setenv("ARTIFACTORY_VERSION", "1.0")                       // nolint
	defer func() {
		testUnStashEnvVars(t)
	}()
	client, err := NewClientFromEnv()
	assert.NoError(t, err)
	assert.True(t, client.Config.APIVersion.Equal(versionMustParse("1.0").Version))
}

func TestClientFromEnvWithNoURL(t *testing.T) {
	testStashEnvVars(t)
	defer func() {
		testUnStashEnvVars(t)
	}()
	client, err := NewClientFromEnv()
	require.Error(t, err)
	require.Nil(t, client)
}

func TestClientFromEnvWithNoCredentials(t *testing.T) {
	testStashEnvVars(t)
	os.Setenv("ARTIFACTORY_URL", "http://artifactory.server.com") //nolint
	defer func() {
		testUnStashEnvVars(t)
	}()
	client, err := NewClientFromEnv()
	require.Error(t, err)
	require.Nil(t, client)

}
