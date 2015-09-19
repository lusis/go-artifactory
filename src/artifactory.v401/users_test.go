package artifactory

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"
)

func TestGetUsers(t *testing.T) {
	responseFile, err := os.Open("assets/test/users.json")
	if err != nil {
		t.Fatalf("Unable to read test data: %s", err.Error())
	}
	defer responseFile.Close()
	responseBody, _ := ioutil.ReadAll(responseFile)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, string(responseBody))
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
	users, err := client.GetUsers()
	assert.Nil(t, err, "should not return an error")
	assert.Len(t, users, 2, "should have two users")
}

func TestGetUserDetails(t *testing.T) {
	responseFile, err := os.Open("assets/test/single_user.json")
	if err != nil {
		t.Fatalf("Unable to read test data: %s", err.Error())
	}
	defer responseFile.Close()
	responseBody, _ := ioutil.ReadAll(responseFile)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, string(responseBody))
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
	user, err := client.GetUserDetails("admin")
	assert.Nil(t, err, "should not return an error")
	assert.Equal(t, user.Name, "admin", "name should be admin")
	assert.Equal(t, user.Email, "admin@admin.com", "should have email of admin@admin.com")
	assert.True(t, user.Admin, "user should be an admin")
	assert.True(t, user.ProfileUpdatable, "profile updatable should be true")
	assert.False(t, user.InternalPasswordDisabled, "Internal password should not be disabled")
	assert.Len(t, user.Groups, 1, "User should be in one group")
	assert.Equal(t, user.Groups[0], "administrators", "user should be in the administrators group")
	assert.Equal(t, user.Realm, "internal", "user realm should be internal")
	assert.NotNil(t, user.LastLoggedIn, "lastLoggedIn should not be empty")
}
