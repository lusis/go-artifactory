package artifactory

import (
	"fmt"
	"net/url"
	"strings"
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

func TestExpirePasswordForUser(t *testing.T) {
	client, server, err := newTestClient([]byte("foo"), "text/plain", 200)
	defer server.Close()
	require.NoError(t, err)
	err = client.ExpirePasswordForUser("foo")
	require.NoError(t, err)
}

func TestExpirePasswordForMultipleUsers(t *testing.T) {
	client, server, err := newTestClient([]byte("foo"), "text/plain", 200)
	defer server.Close()
	require.NoError(t, err)
	err = client.ExpirePasswordForMultipleUsers([]string{"foo", "bar"})
	require.NoError(t, err)
}

func TestExpirePasswordForAllUsers(t *testing.T) {
	client, server, err := newTestClient([]byte("foo"), "text/plain", 200)
	defer server.Close()
	require.NoError(t, err)
	err = client.ExpirePasswordForAllUsers()
	require.NoError(t, err)
}

func TestUnExpirePasswordForUser(t *testing.T) {
	client, server, err := newTestClient([]byte("foo"), "text/plain", 200)
	defer server.Close()
	require.NoError(t, err)
	err = client.UnExpirePasswordForUser("foo")
	require.NoError(t, err)
}

func TestChangePassword(t *testing.T) {
	client, server, err := newTestClient([]byte("foo"), "text/plain", 200)
	defer server.Close()
	require.NoError(t, err)
	err = client.ChangePassword("foo", "password1", "password2")
	require.NoError(t, err)
}

func TestGetPasswordExpirationPolicy(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.GetPasswordExpirationPolicyResponseTestData)
	require.NoError(t, err)
	client, server, err := newTestClient(jsonfile, "application/json", 200)
	defer server.Close()
	require.NoError(t, err)
	obj, err := client.GetPasswordExpirationPolicy()
	require.NoError(t, err)
	require.NotNil(t, obj)
}

func TestSetPasswordExpirationPolicy(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.GetPasswordExpirationPolicyResponseTestData)
	require.NoError(t, err)
	client, server, err := newTestClient(jsonfile, "application/json", 200)
	defer server.Close()
	require.NoError(t, err)
	err = client.SetPasswordExpirationPolicy(3600, false, false)
	require.NoError(t, err)
}

func TestConfigureUserLockPolicy(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.GetPasswordExpirationPolicyResponseTestData)
	require.NoError(t, err)
	client, server, err := newTestClient(jsonfile, "application/json", 200)
	defer server.Close()
	require.NoError(t, err)
	err = client.ConfigureUserLockPolicy(1, false)
	require.NoError(t, err)
}

func TestRetrieveUserLockPolicy(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.RetrieveUserLockPolicyResponseTestData)
	require.NoError(t, err)
	client, server, err := newTestClient(jsonfile, "application/json", 200)
	defer server.Close()
	require.NoError(t, err)
	obj, err := client.RetrieveUserLockPolicy()
	require.NoError(t, err)
	require.NotNil(t, obj)
}

func TestGetLockedOutUsers(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.GetLockedOutUsersResponseTestData)
	require.NoError(t, err)
	client, server, err := newTestClient(jsonfile, "application/json", 200)
	defer server.Close()
	require.NoError(t, err)
	obj, err := client.GetLockedOutUsers()
	require.NoError(t, err)
	require.NotNil(t, obj)
}

func TestUnlockLockedOutUser(t *testing.T) {
	client, server, err := newTestClient([]byte(""), "application/json", 200)
	defer server.Close()
	require.NoError(t, err)
	err = client.UnlockLockedOutUser("bob")
	require.NoError(t, err)
}

func TestUnlockLockedUsers(t *testing.T) {
	client, server, err := newTestClient([]byte(""), "application/json", 200)
	defer server.Close()
	require.NoError(t, err)
	err = client.UnlockLockedUsers([]string{"bob"})
	require.NoError(t, err)
}

func TestUnlockAllLockedUsers(t *testing.T) {
	client, server, err := newTestClient([]byte(""), "application/json", 200)
	defer server.Close()
	require.NoError(t, err)
	err = client.UnlockAllLockedUsers()
	require.NoError(t, err)
}

func TestCreateAPIKey(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.CreateAPIKeyResponseTestData)
	require.NoError(t, err)
	client, server, err := newTestClient(jsonfile, "application/json", 200)
	defer server.Close()
	require.NoError(t, err)
	obj, err := client.CreateAPIKey()
	require.NoError(t, err)
	require.NotNil(t, obj)
}

func TestRegenerateAPIKey(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.RegenerateAPIKeyResponseTestData)
	require.NoError(t, err)
	client, server, err := newTestClient(jsonfile, "application/json", 200)
	defer server.Close()
	require.NoError(t, err)
	obj, err := client.RegenerateAPIKey()
	require.NoError(t, err)
	require.NotNil(t, obj)
}

func TestGetAPIKey(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.GetAPIKeyResponseTestData)
	require.NoError(t, err)
	client, server, err := newTestClient(jsonfile, "application/json", 200)
	defer server.Close()
	require.NoError(t, err)
	obj, err := client.GetAPIKey()
	require.NoError(t, err)
	require.NotNil(t, obj)
}

func TestRevokeAPIKey(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.GetAPIKeyResponseTestData)
	require.NoError(t, err)
	client, server, err := newTestClient(jsonfile, "application/json", 200)
	defer server.Close()
	require.NoError(t, err)
	err = client.RevokeAPIKey()
	require.NoError(t, err)
}

func TestRevokeUserAPIKey(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.GetAPIKeyResponseTestData)
	require.NoError(t, err)
	client, server, err := newTestClient(jsonfile, "application/json", 200)
	defer server.Close()
	require.NoError(t, err)
	err = client.RevokeUserAPIKey("bob")
	require.NoError(t, err)
}

func TestRevokeAllAPIKeys(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.GetAPIKeyResponseTestData)
	require.NoError(t, err)
	client, server, err := newTestClient(jsonfile, "application/json", 200)
	defer server.Close()
	require.NoError(t, err)
	err = client.RevokeAllAPIKeys()
	require.NoError(t, err)
}

func TestGetGroups(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.GetGroupsResponseTestData)
	require.NoError(t, err)
	client, server, err := newTestClient(jsonfile, "application/json", 200)
	defer server.Close()
	require.NoError(t, err)

	obj, err := client.GetGroups()
	require.NoError(t, err)
	require.NotNil(t, obj)
}

func TestGetGroupDetails(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.GetGroupDetailsResponseTestData)
	require.NoError(t, err)
	client, server, err := newTestClient(jsonfile, "application/json", 200)
	defer server.Close()
	require.NoError(t, err)

	obj, err := client.GetGroupDetails("foo")
	require.NoError(t, err)
	require.NotNil(t, obj)
}

func TestCreateOrReplaceGroup(t *testing.T) {
	client, server, err := newTestClient([]byte(""), "application/json", 200)
	defer server.Close()
	require.NoError(t, err)

	details := requests.CreateOrReplaceGroupRequest{
		Description: "my new group",
	}
	err = client.CreateOrReplaceGroup("foo", details)
	require.NoError(t, err)
}

func TestUpdateGroup(t *testing.T) {
	client, server, err := newTestClient([]byte(""), "application/json", 200)
	defer server.Close()
	require.NoError(t, err)

	details := requests.UpdateGroupRequest{
		Description: "my new group",
	}
	err = client.UpdateGroup("foo", details)
	require.NoError(t, err)
}

func TestDeleteGroup(t *testing.T) {
	client, server, err := newTestClient([]byte(""), "application/json", 200)
	defer server.Close()
	require.NoError(t, err)

	err = client.DeleteGroup("foo")
	require.NoError(t, err)
}

func TestGetPermissionTargets(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.GetPermissionTargetsResponseTestData)
	require.NoError(t, err)
	client, server, err := newTestClient(jsonfile, "application/json", 200)
	defer server.Close()
	require.NoError(t, err)

	obj, err := client.GetPermissionTargets()
	require.NoError(t, err)
	require.NotNil(t, obj)
}

func TestGetPermissionTargetDetails(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.GetPermissionTargetDetailsResponseTestData)
	require.NoError(t, err)
	client, server, err := newTestClient(jsonfile, "application/json", 200)
	defer server.Close()
	require.NoError(t, err)

	obj, err := client.GetPermissionTargetDetails("foo")
	require.NoError(t, err)
	require.NotNil(t, obj)
}

func TestCreateOrReplacePermissionTarget(t *testing.T) {
	client, server, err := newTestClient([]byte(""), "application/json", 200)
	defer server.Close()
	require.NoError(t, err)

	details := requests.CreateOrReplacePermissionTargetRequest{}
	err = client.CreateOrReplacePermissionTarget("foo", details)
	require.NoError(t, err)
}

func TestDeletePermissionTarget(t *testing.T) {
	client, server, err := newTestClient([]byte(""), "application/json", 200)
	defer server.Close()
	require.NoError(t, err)

	err = client.DeletePermissionTarget("foo")
	require.NoError(t, err)
}

func TestGetEffectiveItemPermissions(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.EffectiveItemPermissionsResponseTestData)
	require.NoError(t, err)
	client, server, err := newTestClient(jsonfile, "application/json", 200)
	defer server.Close()
	require.NoError(t, err)

	obj, err := client.GetEffectiveItemPermissions("foo")
	require.NoError(t, err)
	require.NotNil(t, obj)
}

func TestGetSecurityConfiguration(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.EffectiveItemPermissionsResponseTestData)
	require.NoError(t, err)
	client, server, err := newTestClient(jsonfile, "application/json", 200)
	defer server.Close()
	require.NoError(t, err)

	obj, err := client.GetSecurityConfiguration()
	require.NoError(t, err)
	require.NotNil(t, obj)
}

func TestActivateArtifactoryKeyEncryption(t *testing.T) {
	client, server, err := newTestClient([]byte(""), "application/json", 200)
	defer server.Close()
	require.NoError(t, err)

	err = client.ActivateArtifactoryKeyEncryption()
	require.NoError(t, err)
}

func TestDeactivateArtifactoryKeyEncryption(t *testing.T) {
	client, server, err := newTestClient([]byte(""), "application/json", 200)
	defer server.Close()
	require.NoError(t, err)

	err = client.DeactivateArtifactoryKeyEncryption()
	require.NoError(t, err)
}

func TestGetGPGPublicKey(t *testing.T) {
	client, server, err := newTestClient([]byte("abcdefg"), "application/json", 200)
	defer server.Close()
	require.NoError(t, err)

	key, err := client.GetGPGPublicKey()
	require.NoError(t, err)
	require.NotEmpty(t, key)
}

func TestSetGPGPublicKey(t *testing.T) {
	client, server, err := newTestClient([]byte("abcdefg"), "application/json", 200)
	defer server.Close()
	require.NoError(t, err)

	err = client.SetGPGPublicKey(strings.NewReader("abcdefg"))
	require.NoError(t, err)
}

func TestSetGPGPrivateKey(t *testing.T) {
	client, server, err := newTestClient([]byte("abcdefg"), "application/json", 200)
	defer server.Close()
	require.NoError(t, err)

	err = client.SetGPGPrivateKey(strings.NewReader("abcdefg"))
	require.NoError(t, err)
}

func TestSetGPGPassPhrase(t *testing.T) {
	client, server, err := newTestClient([]byte("abcdefg"), "application/json", 200)
	defer server.Close()
	require.NoError(t, err)

	err = client.SetGPGPassPhrase("abcdefg")
	require.NoError(t, err)
}

func TestCreateToken(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.CreateTokenResponseTestData)
	require.NoError(t, err)
	client, server, err := newTestClient(jsonfile, "application/json", 200)
	defer server.Close()
	require.NoError(t, err)
	res, err := client.CreateToken()
	require.NoError(t, err)
	require.NotNil(t, res)
}

func TestCreateTokenWithOptions(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.CreateTokenResponseTestData)
	require.NoError(t, err)
	client, server, err := newTestClient(jsonfile, "application/json", 200)
	defer server.Close()
	require.NoError(t, err)
	opts := []TokenOption{
		TokenGrantType("foo"),
		TokenScope("foo"),
		TokenExpires(3600),
		TokenRefreshable(true),
		TokenAudience("foo"),
		TokenUsername("foo"),
	}
	res, err := client.CreateToken(opts...)
	require.NoError(t, err)
	require.NotNil(t, res)
}

func TestCreateTokenWithBadOption(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.CreateTokenResponseTestData)
	require.NoError(t, err)
	client, server, err := newTestClient(jsonfile, "application/json", 200)
	defer server.Close()
	require.NoError(t, err)
	opt := func(u *url.Values) error {
		return fmt.Errorf("invalid option")
	}
	res, err := client.CreateToken(opt)
	require.Error(t, err)
	require.Nil(t, res)
}

func TestRefreshToken(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.RefreshTokenResponseTestData)
	require.NoError(t, err)
	client, server, err := newTestClient(jsonfile, "application/json", 200)
	defer server.Close()
	require.NoError(t, err)
	res, err := client.RefreshToken("foo", "bar")
	require.NoError(t, err)
	require.NotNil(t, res)
}

func TestRefreshTokenWithOptions(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.RefreshTokenResponseTestData)
	require.NoError(t, err)
	client, server, err := newTestClient(jsonfile, "application/json", 200)
	defer server.Close()
	require.NoError(t, err)
	opts := []TokenOption{
		TokenGrantType("foo"),
		TokenScope("foo"),
		TokenExpires(3600),
		TokenRefreshable(true),
		TokenAudience("foo"),
		TokenUsername("foo"),
	}
	res, err := client.RefreshToken("foo", "bar", opts...)
	require.NoError(t, err)
	require.NotNil(t, res)
}

func TestRefreshTokenWithBadOption(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.RefreshTokenResponseTestData)
	require.NoError(t, err)
	client, server, err := newTestClient(jsonfile, "application/json", 200)
	defer server.Close()
	require.NoError(t, err)
	opt := func(u *url.Values) error {
		return fmt.Errorf("invalid option")
	}
	res, err := client.RefreshToken("foo", "bar", opt)
	require.Error(t, err)
	require.Nil(t, res)
}

func TestRevokeToken(t *testing.T) {
	client, server, err := newTestClient([]byte(""), "application/json", 200)
	defer server.Close()
	require.NoError(t, err)
	err = client.RevokeToken("abcdef")
	require.NoError(t, err)
}

func TestGetServiceID(t *testing.T) {
	client, server, err := newTestClient([]byte("abcdefg"), "application/json", 200)
	defer server.Close()
	require.NoError(t, err)
	res, err := client.GetServiceID()
	require.NoError(t, err)
	require.Equal(t, "abcdefg", res)
}

func TestGetCertificates(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.GetCertificatesResponseTestData)
	require.NoError(t, err)
	client, server, err := newTestClient(jsonfile, "application/json", 200)
	defer server.Close()
	require.NoError(t, err)

	res, err := client.GetCertificates()
	require.NoError(t, err)
	require.NotNil(t, res)
}

func TestAddCertificate(t *testing.T) {
	client, server, err := newTestClient([]byte("foo"), "application/json", 200)
	defer server.Close()
	require.NoError(t, err)

	err = client.AddCertificate("foo", strings.NewReader("bar"))
	require.NoError(t, err)
}

func TestDeleteCertificate(t *testing.T) {
	client, server, err := newTestClient([]byte("foo"), "application/json", 200)
	defer server.Close()
	require.NoError(t, err)

	err = client.DeleteCertificate("foo")
	require.NoError(t, err)
}
