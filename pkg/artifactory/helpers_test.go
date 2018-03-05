package artifactory

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"
)

func newTestClient(content []byte, contentType string, statusCode int) (*Client, *httptest.Server, error) {

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", contentType)
		w.WriteHeader(statusCode)
		fmt.Fprintf(w, string(content))
	}))

	transport := http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			return url.Parse(server.URL)
		},
	}

	httpClient := http.Client{}
	httpClient.Transport = &transport
	conf := &ClientConfig{
		BaseURL:    "http://localhost:8080/artifactory",
		Token:      "XXXXXXXXXXXXX",
		VerifySSL:  false,
		AuthMethod: "token",
		APIVersion: MaxArtifactoryVersion,
		HTTPClient: &httpClient,
	}
	client, err := NewClient(conf)
	if err != nil {
		return nil, nil, err
	}
	return client, server, nil
}

func testStashEnvVars(t *testing.T) {
	for _, v := range []string{
		"ARTIFACTORY_TOKEN",
		"ARTIFACTORY_URL",
		"ARTIFACTORY_USERNAME",
		"ARTIFACTORY_PASSWORD",
		"ARTIFACTORY_VERSION",
		"ARTIFACTORY_DEBUG",
	} {
		if os.Getenv(v) != "" {
			//stash it
			toStash := os.Getenv(v)
			_ = os.Setenv("TST_"+v, toStash)
			// clear it for the tests
			_ = os.Unsetenv(v)
		}
	}
}

func testUnStashEnvVars(t *testing.T) {
	for _, v := range []string{
		"ARTIFACTORY_TOKEN",
		"ARTIFACTORY_URL",
		"ARTIFACTORY_USERNAME",
		"ARTIFACTORY_PASSWORD",
		"ARTIFACTORY_VERSION",
		"ARTIFACTORY_DEBUG",
	} {
		if os.Getenv("TST_"+v) != "" {
			//unstash it
			fromStash := os.Getenv("TST_" + v)
			// clear existing
			_ = os.Unsetenv(v)
			// set from stash
			_ = os.Setenv(v, fromStash)
			// clear stash
			_ = os.Unsetenv("TST_" + v)
		}
	}
}
