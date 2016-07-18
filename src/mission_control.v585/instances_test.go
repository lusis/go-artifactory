package mission_control

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"
)

func TestGetInstances(t *testing.T) {
	responseFile, err := os.Open("assets/test/instances.json")
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
	instances, err := client.GetInstances()
	assert.NoError(t, err, "should not return an error")
	assert.Equal(t, len(instances), 2, "instances size should be 2")
	assert.Equal(t, instances[0].Name, "QA", "instances[0].Name should be QA")
	assert.Equal(t, instances[0].URL, "http://10.0.0.110:8080/artifactory", "instances[0].URL should be http://10.0.0.110:8080/artifactory")
	assert.Equal(t, instances[1].Name, "DEV", "instances[1].Name should be DEV")
	assert.Equal(t, instances[1].URL, "http://10.0.0.120:8080/artifactory", "instances[1].URL should be http://10.0.0.120:8080/artifactory")
}
