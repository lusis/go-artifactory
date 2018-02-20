package artifactory

import (
	"testing"

	"github.com/lusis/go-artifactory/pkg/artifactory/requests"
	"github.com/lusis/go-artifactory/pkg/artifactory/responses"
	"github.com/lusis/go-artifactory/pkg/artifactory/responses/testdata"
	"github.com/stretchr/testify/require"
)

func TestGetUsers(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.GetUsersResponseTestData)
	require.NoError(t, err)
	client, server, err := newTestClient(jsonfile, "application/json", 200)
	defer server.Close()
	require.NoError(t, err)
	obj, err := client.GetUsers()
	require.NoError(t, err)
	require.NotNil(t, obj)
}

func TestGetUsersHTTPError(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.ErrorResponseTestData)
	require.NoError(t, err)
	client, server, err := newTestClient(jsonfile, "application/json", 404)
	defer server.Close()
	require.NoError(t, err)
	obj, err := client.GetUsers()
	require.Error(t, err)
	require.Nil(t, obj)
}

func TestGetUsersJSONError(t *testing.T) {
	client, server, err := newTestClient([]byte(""), "application/json", 200)
	defer server.Close()
	require.NoError(t, err)
	obj, err := client.GetUsers()
	require.Error(t, err)
	require.Nil(t, obj)
}

func TestGetUserDetails(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.GetUserDetailsResponseTestData)
	require.NoError(t, err)
	client, server, err := newTestClient(jsonfile, "application/json", 200)
	defer server.Close()
	require.NoError(t, err)
	obj, err := client.GetUserDetails("bob")
	require.NoError(t, err)
	require.NotNil(t, obj)
}

func TestGetUserDetailsHTTPError(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.ErrorResponseTestData)
	require.NoError(t, err)
	client, server, err := newTestClient(jsonfile, "application/json", 404)
	defer server.Close()
	require.NoError(t, err)
	obj, err := client.GetUserDetails("bob")
	require.Error(t, err)
	require.Nil(t, obj)
}

func TestGetUserDetailsJSONError(t *testing.T) {
	client, server, err := newTestClient([]byte(""), "application/json", 200)
	defer server.Close()
	require.NoError(t, err)
	obj, err := client.GetUserDetails("bob")
	require.Error(t, err)
	require.Nil(t, obj)
}

func TestCreateOrReplaceUser(t *testing.T) {
	client, server, err := newTestClient([]byte(""), "application/json", 200)
	defer server.Close()
	require.NoError(t, err)
	err = client.CreateOrReplaceUser("bob", requests.CreateOrReplaceUserRequest{})
	require.NoError(t, err)
}

func TestUpdateUser(t *testing.T) {
	client, server, err := newTestClient([]byte(""), "application/json", 200)
	defer server.Close()
	require.NoError(t, err)
	err = client.UpdateUser("bob", requests.UpdateUserRequest{})
	require.NoError(t, err)
}

func TestDeleteUser(t *testing.T) {
	client, server, err := newTestClient([]byte(""), "application/json", 200)
	defer server.Close()
	require.NoError(t, err)
	err = client.DeleteUser("bob")
	require.NoError(t, err)
}

func TestGetUserEncryptedPassword(t *testing.T) {
	client, server, err := newTestClient([]byte("abcdefg"), "text/plain", 200)
	defer server.Close()
	require.NoError(t, err)
	p, err := client.GetUserEncryptedPassword()
	require.NoError(t, err)
	require.Equal(t, "abcdefg", string(p))
}

func TestGetUserEncryptedPasswordHTTPError(t *testing.T) {
	client, server, err := newTestClient([]byte("abcdefg"), "text/plain", 500)
	defer server.Close()
	require.NoError(t, err)
	p, err := client.GetUserEncryptedPassword()
	require.Error(t, err)
	require.NotEqual(t, "abcdefg", string(p))
}
