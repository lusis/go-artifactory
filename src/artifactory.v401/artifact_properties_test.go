package artifactory

import (
	"os"
	"io/ioutil"
	"net/http/httptest"
	"net/url"
	"testing"
	"net/http"
	"github.com/stretchr/testify/assert"
)

func TestGetArtifactProperties(t *testing.T) {
	responseFile, err := os.Open("assets/test/artifact_properties.json")
	if err != nil {
		t.Fatalf("Unable to read test data: %s", err.Error())
	}
	defer responseFile.Close()
	responseBody, _ := ioutil.ReadAll(responseFile)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		w.Write(responseBody)
	}))
	defer server.Close()

	transport := &http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			return url.Parse(server.URL)
		},
	}

	conf := &ClientConfig{
		BaseURL:   "http://127.0.0.1:8080/",
		Username:  "username",
		Password:  "password",
		VerifySSL: false,
		Transport: transport,
	}

	client := NewClient(conf)
	properties, err := client.GetArtifactProperties("docker-prod")
	assert.NoError(t, err, "should not return an error")
	assert.Equal(t, len(properties), 6, "properties size should be 6")
	assert.Equal(t, properties[0].Name, "artifactory.replication.docker-prod.result", "properties[0].Name should be artifactory.replication.docker-prod.result")
	assert.Equal(t, properties[0].Value, "ok", "properties[0].Value should be ok")
}
