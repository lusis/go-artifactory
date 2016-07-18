package mission_control

import (
	"net/http/httptest"
	"net/url"
	"testing"
	"net/http"
	"go.riotgames.com/telemetry-engineering/germes-utils/Godeps/_workspace/src/github.com/stretchr/testify/assert"
)

func TestHTTP(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, r.URL.String(), "http://127.0.0.1:8080/test/url/here", "Request string should be http://127.0.0.1:8080/test/url/here, but was: " + r.URL.String())
		        w.WriteHeader(200)
			w.Write([]byte("{}"))
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

	//Address ends with slash
	_, err := client.Get("/test/url/here", make(map[string]string))
	assert.NoError(t, err, "should not return an error")


	conf.BaseURL = "http://127.0.0.1:8080"
	client = NewClient(conf)

	//Address ends with no Slash
	_, err = client.Get("/test/url/here", make(map[string]string))
	assert.NoError(t, err, "should not return an error")
}
