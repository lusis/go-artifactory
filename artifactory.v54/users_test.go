package artifactory

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUsers(t *testing.T) {
	responseFile, err := os.Open("assets/test/users.json")
	if err != nil {
		t.Fatalf("Unable to read test data: %s", err.Error())
	}
	defer func() { _ = responseFile.Close() }()
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
	assert.NoError(t, err, "should not return an error")
	assert.Len(t, users, 2, "should have two users")
}

func TestGetUserDetails(t *testing.T) {
	responseFile, err := os.Open("assets/test/single_user.json")
	if err != nil {
		t.Fatalf("Unable to read test data: %s", err.Error())
	}
	defer func() { _ = responseFile.Close() }()
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
	user, err := client.GetUser("admin", make(map[string]string))
	assert.NoError(t, err, "should not return an error")
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

func TestCreateUser(t *testing.T) {
	var buf bytes.Buffer
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		req, _ := ioutil.ReadAll(r.Body)
		buf.Write(req)
		fmt.Fprintf(w, "")
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

	user := User{
		Name:                     "admin",
		Email:                    "test@test.com",
		Password:                 "somepass",
		Admin:                    true,
		ProfileUpdatable:         true,
		DisableUIAccess:          false,
		InternalPasswordDisabled: false,
		Groups:       []string{"administrators"},
		LastLoggedIn: "2015-08-11T14:04:11.472Z",
		Realm:        "internal",
	}

	expectedJSON, _ := json.Marshal(user)
	err := client.CreateUser("admin", user, make(map[string]string))
	assert.NoError(t, err, "should not return an error")
	assert.Equal(t, string(expectedJSON), buf.String(), "should send user json")
}

func TestCreateUserFailure(t *testing.T) {
	conf := &ClientConfig{
		BaseURL:   "http://127.0.0.1:8080/",
		Username:  "username",
		Password:  "password",
		VerifySSL: false,
	}

	client := NewClient(conf)
	var details = User{}
	err := client.CreateUser("testuser", details, make(map[string]string))
	assert.Error(t, err, "should return an error")
}

func TestDeleteUser(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "")
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
	err := client.DeleteUser("testuser")
	assert.NoError(t, err, "should not return an error")
}