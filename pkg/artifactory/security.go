package artifactory

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"strconv"
	"strings"

	multierror "github.com/hashicorp/go-multierror"
	"github.com/lusis/go-artifactory/pkg/artifactory/requests"
	"github.com/lusis/go-artifactory/pkg/artifactory/responses"
)

// Users is a collection of User
type Users []responses.UserResponse

// UserDetails is a user's details
type UserDetails struct {
	responses.GetUserDetailsResponse
}

// PasswordExpirationPolicy represents a password expiration Policy
type PasswordExpirationPolicy struct {
	responses.GetPasswordExpirationPolicyResponse
}

// UserLockPolicy represents a user lock policy
type UserLockPolicy struct {
	responses.RetrieveUserLockPolicyResponse
}

// Groups is a list of groups
type Groups []responses.GroupResponse

// GroupDetails is a group's details
type GroupDetails struct {
	responses.GetGroupDetailsResponse
}

// PermissionTargets is a list of permission targets
type PermissionTargets []responses.PermissionTargetEntry

// PermissionTargetDetails is a permission target's details
type PermissionTargetDetails struct {
	responses.GetPermissionTargetDetailsResponse
}

// EffectiveItemPermissions returns the effective item permissions for the specified item or folder
type EffectiveItemPermissions struct {
	responses.EffectiveItemPermissionsResponse
}

// Token represents an Access Token
type Token struct {
	responses.CreateTokenResponse
}

// TokenOption is a functional option for creating tokens
type TokenOption func(*url.Values) error

// Certificates is a collection of certificates
type Certificates responses.GetCertificatesResponse

// GetUsers returns a list of users
func (c *Client) GetUsers() (Users, error) {
	if err := c.checkRequiredResponseVersion(responses.GetUsersResponse{}); err != nil {
		return nil, err
	}
	users := Users{}
	res, err := c.httpGet("security/users", requestJSON(), requestExpects(200))
	if err != nil {
		return nil, err
	}
	if jsonErr := json.Unmarshal(res, &users); jsonErr != nil {
		return nil, &UnmarshalError{msg: multierror.Append(errDecoding, jsonErr).Error()}
	}
	return users, nil
}

// GetUserDetails gets a user's details
func (c *Client) GetUserDetails(u string) (*UserDetails, error) {
	if err := c.checkRequiredResponseVersion(responses.GetUserDetailsResponse{}); err != nil {
		return nil, err
	}
	users := &UserDetails{}
	res, err := c.httpGet("security/users/"+u, requestJSON(), requestExpects(200))
	if err != nil {
		return nil, err
	}
	if jsonErr := json.Unmarshal(res, users); jsonErr != nil {
		return nil, &UnmarshalError{msg: multierror.Append(errDecoding, jsonErr).Error()}
	}
	return users, nil
}

// CreateOrReplaceUser creates or updates the provided user with the provided details
func (c *Client) CreateOrReplaceUser(u string, details requests.CreateOrReplaceUserRequest) error {
	if err := c.checkRequiredRequestVersion(details); err != nil {
		return err
	}
	j, err := json.Marshal(details)
	if err != nil {
		return &MarshalError{msg: multierror.Append(errDecoding, err).Error()}
	}
	url := fmt.Sprintf("security/users/%s", u)
	_, err = c.httpPut(url, requestExpects(200), requestJSON(), withBody(bytes.NewReader(j)))
	return err

}

// UpdateUser creates or updates the provided user with the provided details
func (c *Client) UpdateUser(u string, details requests.UpdateUserRequest) error {
	if err := c.checkRequiredRequestVersion(details); err != nil {
		return err
	}
	j, err := json.Marshal(details)
	if err != nil {
		return &MarshalError{msg: multierror.Append(errDecoding, err).Error()}
	}
	url := fmt.Sprintf("security/users/%s", u)
	_, err = c.httpPost(url, requestExpects(200), requestJSON(), withBody(bytes.NewReader(j)))
	return err
}

// DeleteUser creates or updates the provided user with the provided details
func (c *Client) DeleteUser(u string) error {
	if err := c.checkRequiredRequestVersion(requests.GenericVersionedRequest{}); err != nil {
		return err
	}
	url := fmt.Sprintf("security/users/%s", u)
	return c.httpDelete(url, requestExpects(200), requestJSON())
}

// GetUserEncryptedPassword returns the encrypted password for the current user
func (c *Client) GetUserEncryptedPassword() (string, error) {
	if err := c.checkRequiredResponseVersion(responses.GetEncryptedPasswordResponse{}); err != nil {
		return "", err
	}
	res, err := c.httpGet("security/encryptedPassword", requestExpects(200))
	if err != nil {
		return "", err
	}
	return string(res), nil
}

// ExpirePasswordForUser expires the password for a user
func (c *Client) ExpirePasswordForUser(user string) error {
	if err := c.checkRequiredRequestVersion(requests.ExpirePasswordForUserRequest{}); err != nil {
		return err
	}
	_, err := c.httpPost("security/users/authorization/expirePassword/"+user, requestExpects(200), requestJSON())
	return err
}

// ExpirePasswordForMultipleUsers expires the password for multiple users
func (c *Client) ExpirePasswordForMultipleUsers(users []string) error {
	if err := c.checkRequiredRequestVersion(requests.ExpirePasswordForMultipleUsersRequest{}); err != nil {
		return err
	}
	j, err := json.Marshal(users)
	if err != nil {
		return &MarshalError{msg: multierror.Append(errDecoding, err).Error()}
	}
	_, err = c.httpPost("security/users/authorization/expirePassword", requestExpects(200), requestJSON(), withBody(bytes.NewReader(j)))
	return err
}

// ExpirePasswordForAllUsers expires the password for multiple users
func (c *Client) ExpirePasswordForAllUsers() error {
	if err := c.checkRequiredRequestVersion(requests.ExpirePasswordForAllUsersRequest{}); err != nil {
		return err
	}

	_, err := c.httpPost("security/users/authorization/expirePasswordForAllUsers", requestExpects(200), requestJSON())
	return err
}

// UnExpirePasswordForUser expires the password for a user
func (c *Client) UnExpirePasswordForUser(user string) error {
	if err := c.checkRequiredRequestVersion(requests.UnexpirePasswordForUserRequest{}); err != nil {
		return err
	}
	_, err := c.httpPost("security/users/authorization/unexpirePassword/"+user, requestExpects(200), requestJSON())
	return err
}

// ChangePassword changes a user's password
func (c *Client) ChangePassword(user, old, new string) error {
	if err := c.checkRequiredRequestVersion(requests.ChangePasswordRequest{}); err != nil {
		return err
	}
	cp := &requests.ChangePasswordRequest{
		UserName:     user,
		OldPassword:  old,
		NewPassword1: new,
		NewPassword2: new,
	}
	j, err := json.Marshal(cp)
	if err != nil {
		return &MarshalError{msg: multierror.Append(errDecoding, err).Error()}
	}
	_, err = c.httpPost("security/users/authorization/changePassword", requestExpects(200), requestJSON(), withBody(bytes.NewReader(j)))
	return err
}

// GetPasswordExpirationPolicy returns the password expiration policy
func (c *Client) GetPasswordExpirationPolicy() (*PasswordExpirationPolicy, error) {
	if err := c.checkRequiredResponseVersion(responses.GetPasswordExpirationPolicyResponse{}); err != nil {
		return nil, err
	}
	obj := &PasswordExpirationPolicy{}
	res, err := c.httpGet("security/configuration/passwordExpirationPolicy", requestJSON(), requestExpects(200))
	if err != nil {
		return nil, err
	}
	if jsonErr := json.Unmarshal(res, obj); jsonErr != nil {
		return nil, &UnmarshalError{msg: multierror.Append(errDecoding, jsonErr).Error()}
	}
	return obj, nil
}

// SetPasswordExpirationPolicy sets the password expiration policy
func (c *Client) SetPasswordExpirationPolicy(passwordMaxAge int, enabled, notifyByEmail bool) error {
	if err := c.checkRequiredRequestVersion(requests.SetPasswordExpirationPolicyRequest{}); err != nil {
		return err
	}
	obj := &requests.SetPasswordExpirationPolicyRequest{
		Enabled:        enabled,
		PasswordMaxAge: passwordMaxAge,
		NotifyByEmail:  notifyByEmail,
	}
	j, err := json.Marshal(obj)
	if err != nil {
		return &MarshalError{msg: multierror.Append(errDecoding, err).Error()}
	}
	_, err = c.httpPut("security/configuration/passwordExpirationPolicy", requestExpects(200), requestJSON(), withBody(bytes.NewReader(j)))
	return err
}

// ConfigureUserLockPolicy sets the user lock policy
func (c *Client) ConfigureUserLockPolicy(loginAttempts int, enabled bool) error {
	if err := c.checkRequiredRequestVersion(requests.ConfigureUserLockPolicyRequest{}); err != nil {
		return err
	}
	obj := &requests.ConfigureUserLockPolicyRequest{
		Enabled:       enabled,
		LoginAttempts: loginAttempts,
	}
	j, err := json.Marshal(obj)
	if err != nil {
		return &MarshalError{msg: multierror.Append(errDecoding, err).Error()}
	}
	_, err = c.httpPut("security/userLockPolicy", requestExpects(200), requestJSON(), withBody(bytes.NewReader(j)))
	return err
}

// RetrieveUserLockPolicy returns the password expiration policy
func (c *Client) RetrieveUserLockPolicy() (*UserLockPolicy, error) {
	if err := c.checkRequiredResponseVersion(responses.RetrieveUserLockPolicyResponse{}); err != nil {
		return nil, err
	}
	obj := &UserLockPolicy{}
	res, err := c.httpGet("security/userLockPolicy", requestJSON(), requestExpects(200))
	if err != nil {
		return nil, err
	}
	if jsonErr := json.Unmarshal(res, obj); jsonErr != nil {
		return nil, &UnmarshalError{msg: multierror.Append(errDecoding, jsonErr).Error()}
	}
	return obj, nil
}

// GetLockedOutUsers gets a list of locked out users
func (c *Client) GetLockedOutUsers() ([]string, error) {
	if err := c.checkRequiredResponseVersion(responses.GetLockedOutUsersResponse{}); err != nil {
		return nil, err
	}
	res, err := c.httpGet("security/lockedUsers", requestJSON(), requestExpects(200))
	if err != nil {
		return nil, err
	}
	users := []string{}
	if jsonErr := json.Unmarshal(res, &users); jsonErr != nil {
		return nil, &UnmarshalError{msg: multierror.Append(errDecoding, jsonErr).Error()}
	}
	return users, nil
}

// UnlockLockedOutUser unlocks a locked user
func (c *Client) UnlockLockedOutUser(u string) error {
	if err := c.checkRequiredRequestVersion(requests.UnlockLockedOutUserRequest{}); err != nil {
		return err
	}

	_, err := c.httpPost("security/unlockUsers/"+u, requestExpects(200), requestJSON())
	return err
}

// UnlockLockedUsers unlocks a list of users
func (c *Client) UnlockLockedUsers(users []string) error {
	if err := c.checkRequiredRequestVersion(requests.UnlockLockedOutUsersRequest{}); err != nil {
		return err
	}
	j, err := json.Marshal(users)
	if err != nil {
		return &MarshalError{msg: multierror.Append(errDecoding, err).Error()}
	}
	_, err = c.httpPost("security/unlockUsers", requestExpects(200), requestJSON(), withBody(bytes.NewReader(j)))
	return err
}

// UnlockAllLockedUsers unlocks all locked out users
func (c *Client) UnlockAllLockedUsers() error {
	if err := c.checkRequiredRequestVersion(requests.UnlockAllLockedOutUsersRequest{}); err != nil {
		return err
	}
	_, err := c.httpPost("security/unlockAllUsers", requestExpects(200), requestJSON())
	return err
}

// CreateAPIKey creates an apikey for the current user
func (c *Client) CreateAPIKey() (string, error) {
	if err := c.checkRequiredRequestVersion(requests.CreateAPIKeyRequest{}); err != nil {
		return "", err
	}
	res, err := c.httpPost("security/apiKey", requestJSON(), requestExpects(200))
	if err != nil {
		return "", err
	}
	data := &responses.CreateAPIKeyResponse{}
	if jsonErr := json.Unmarshal(res, data); jsonErr != nil {
		return "", &UnmarshalError{msg: multierror.Append(errDecoding, jsonErr).Error()}
	}
	return data.APIKey, nil
}

// RegenerateAPIKey regenerate an apikey for the current user
func (c *Client) RegenerateAPIKey() (string, error) {
	if err := c.checkRequiredRequestVersion(requests.RegenerateAPIKeyRequest{}); err != nil {
		return "", err
	}
	res, err := c.httpPut("security/apiKey", requestJSON(), requestExpects(200))
	if err != nil {
		return "", err
	}
	data := &responses.RegenerateAPIKeyResponse{}
	if jsonErr := json.Unmarshal(res, data); jsonErr != nil {
		return "", &UnmarshalError{msg: multierror.Append(errDecoding, jsonErr).Error()}
	}
	return data.APIKey, nil
}

// GetAPIKey creates an apikey for the current user
func (c *Client) GetAPIKey() (string, error) {
	if err := c.checkRequiredResponseVersion(responses.GetAPIKeyResponse{}); err != nil {
		return "", err
	}
	res, err := c.httpPost("security/apiKey", requestJSON(), requestExpects(200))
	if err != nil {
		return "", err
	}
	data := &responses.GetAPIKeyResponse{}
	if jsonErr := json.Unmarshal(res, data); jsonErr != nil {
		return "", &UnmarshalError{msg: multierror.Append(errDecoding, jsonErr).Error()}
	}
	return data.APIKey, nil
}

// RevokeUserAPIKey revokes an apikey for the specified user
func (c *Client) RevokeUserAPIKey(u string) error {
	if err := c.checkRequiredRequestVersion(requests.RevokeUserAPIKeyRequest{}); err != nil {
		return err
	}
	return c.httpDelete("security/apiKey"+u, requestJSON(), requestExpects(200))
}

// RevokeAPIKey revokes an apikey for the current user
func (c *Client) RevokeAPIKey() error {
	if err := c.checkRequiredRequestVersion(requests.RevokeAPIKeyRequest{}); err != nil {
		return err
	}
	return c.httpDelete("security/apiKey", requestJSON(), requestExpects(200))
}

// RevokeAllAPIKeys revokes all api keys
func (c *Client) RevokeAllAPIKeys() error {
	if err := c.checkRequiredRequestVersion(requests.RevokeAllAPIKeysRequest{}); err != nil {
		return err
	}
	return c.httpDelete("security/apiKey", requestJSON(), requestExpects(200), queryParams(map[string]string{"deleteAll": "1"}))
}

// GetGroups gets a list of groups
func (c *Client) GetGroups() (Groups, error) {
	data := Groups{}
	if err := c.checkRequiredResponseVersion(responses.GetGroupsResponse{}); err != nil {
		return data, err
	}
	res, err := c.httpGet("security/groups", requestJSON(), requestExpects(200))
	if err != nil {
		return data, err
	}
	if jsonErr := json.Unmarshal(res, &data); jsonErr != nil {
		return data, &UnmarshalError{msg: multierror.Append(errDecoding, jsonErr).Error()}
	}
	return data, nil
}

// GetGroupDetails gets a group's details
func (c *Client) GetGroupDetails(g string) (*GroupDetails, error) {
	if err := c.checkRequiredResponseVersion(responses.GetGroupDetailsResponse{}); err != nil {
		return nil, err
	}
	res, err := c.httpGet("security/groups/"+g, requestJSON(), requestExpects(200))
	if err != nil {
		return nil, err
	}
	data := &GroupDetails{}
	if jsonErr := json.Unmarshal(res, data); jsonErr != nil {
		return nil, &UnmarshalError{msg: multierror.Append(errDecoding, jsonErr).Error()}
	}
	return data, nil
}

// CreateOrReplaceGroup creates or updates the provided group with the provided details
func (c *Client) CreateOrReplaceGroup(g string, details requests.CreateOrReplaceGroupRequest) error {
	if err := c.checkRequiredRequestVersion(details); err != nil {
		return err
	}
	j, err := json.Marshal(details)
	if err != nil {
		return &MarshalError{msg: multierror.Append(errEncoding, err).Error()}
	}
	url := fmt.Sprintf("security/groups/%s", g)
	_, err = c.httpPut(url, requestExpects(200), requestJSON(), withBody(bytes.NewReader(j)))
	return err
}

// UpdateGroup creates or updates the provided group with the provided details
func (c *Client) UpdateGroup(g string, details requests.UpdateGroupRequest) error {
	if err := c.checkRequiredRequestVersion(details); err != nil {
		return err
	}
	j, err := json.Marshal(details)
	if err != nil {
		return &MarshalError{msg: multierror.Append(errEncoding, err).Error()}
	}
	url := fmt.Sprintf("security/groups/%s", g)
	_, err = c.httpPost(url, requestExpects(200), requestJSON(), withBody(bytes.NewReader(j)))
	return err
}

// DeleteGroup deletes the provided group
func (c *Client) DeleteGroup(g string) error {
	if err := c.checkRequiredRequestVersion(requests.DeleteGroupRequest{}); err != nil {
		return err
	}
	url := fmt.Sprintf("security/groups/%s", g)
	return c.httpDelete(url, requestExpects(200), requestJSON())
}

// GetPermissionTargets gets a list of permission targets
func (c *Client) GetPermissionTargets() (PermissionTargets, error) {
	if err := c.checkRequiredResponseVersion(responses.GetPermissionTargetsResponse{}); err != nil {
		return nil, err
	}
	res, err := c.httpGet("security/permissions", requestJSON(), requestExpects(200))
	if err != nil {
		return nil, err
	}
	data := PermissionTargets{}
	if jsonErr := json.Unmarshal(res, &data); jsonErr != nil {
		return nil, &UnmarshalError{msg: multierror.Append(errDecoding, jsonErr).Error()}
	}
	return data, nil
}

// GetPermissionTargetDetails gets details about a permission target
func (c *Client) GetPermissionTargetDetails(g string) (*PermissionTargetDetails, error) {
	if err := c.checkRequiredResponseVersion(responses.GetPermissionTargetDetailsResponse{}); err != nil {
		return nil, err
	}
	res, err := c.httpGet("security/permissions/"+g, requestJSON(), requestExpects(200))
	if err != nil {
		return nil, err
	}
	data := &PermissionTargetDetails{}
	if jsonErr := json.Unmarshal(res, data); jsonErr != nil {
		return nil, &UnmarshalError{msg: multierror.Append(errDecoding, jsonErr).Error()}
	}
	return data, nil
}

// CreateOrReplacePermissionTarget creates or replaces the permission target with provided details
func (c *Client) CreateOrReplacePermissionTarget(g string, details requests.CreateOrReplacePermissionTargetRequest) error {
	if err := c.checkRequiredRequestVersion(details); err != nil {
		return err
	}
	j, err := json.Marshal(details)
	if err != nil {
		return &MarshalError{msg: multierror.Append(errEncoding, err).Error()}
	}
	url := fmt.Sprintf("security/permissions/%s", g)
	_, err = c.httpPut(url, requestExpects(200), requestJSON(), withBody(bytes.NewReader(j)))
	return err
}

// DeletePermissionTarget deletes the specified permission target
func (c *Client) DeletePermissionTarget(g string) error {
	if err := c.checkRequiredRequestVersion(requests.DeletePermissionTargetRequest{}); err != nil {
		return err
	}
	url := fmt.Sprintf("security/permissions/%s", g)
	return c.httpDelete(url, requestExpects(200), requestJSON())
}

// GetEffectiveItemPermissions gets an item or folder's effective permissions
func (c *Client) GetEffectiveItemPermissions(p string) (*EffectiveItemPermissions, error) {
	if err := c.checkRequiredResponseVersion(responses.EffectiveItemPermissionsResponse{}); err != nil {
		return nil, err
	}
	res, err := c.httpGet("storage/"+p, requestJSON(), requestExpects(200))
	if err != nil {
		return nil, err
	}
	data := &EffectiveItemPermissions{}
	if jsonErr := json.Unmarshal(res, data); jsonErr != nil {
		return nil, &UnmarshalError{msg: multierror.Append(errDecoding, jsonErr).Error()}
	}
	return data, nil
}

// GetSecurityConfiguration gets the system security configuration
func (c *Client) GetSecurityConfiguration() (string, error) {
	if err := c.checkRequiredResponseVersion(responses.SecurityConfigurationResponse{}); err != nil {
		return "", err
	}
	res, err := c.httpGet("system/security", accept("application/xml"), requestExpects(200))
	if err != nil {
		return "", err
	}
	return string(res), nil
}

// ActivateArtifactoryKeyEncryption creates a new artifactory encryption key and activates key encryption
func (c *Client) ActivateArtifactoryKeyEncryption() error {
	if err := c.checkRequiredRequestVersion(requests.ActivateArtifactoryKeyEncryptionRequest{}); err != nil {
		return err
	}
	_, err := c.httpPost("system/encrypt", requestJSON(), requestExpects(200))
	return err
}

// DeactivateArtifactoryKeyEncryption removes a new artifactory encryption key and deactivates key encryption
func (c *Client) DeactivateArtifactoryKeyEncryption() error {
	if err := c.checkRequiredRequestVersion(requests.DeactivateArtifactoryKeyEncryptionRequest{}); err != nil {
		return err
	}
	_, err := c.httpPost("system/decrypt", requestJSON(), requestExpects(200))
	return err
}

// GetGPGPublicKey gets the public key that artifactory provides to debian and opkg clients to verify packages
func (c *Client) GetGPGPublicKey() (string, error) {
	if err := c.checkRequiredResponseVersion(responses.GetGPGPublicKeyResponse{}); err != nil {
		return "", err
	}
	res, err := c.httpGet("gpg/key/public", accept("text/plain"), requestExpects(200))
	return string(res), err
}

// SetGPGPublicKey sets the public key that artifactory provides to debian and opkg clients to verify packages
func (c *Client) SetGPGPublicKey(data io.Reader) error {
	if err := c.checkRequiredRequestVersion(requests.SetGPGPublicKeyRequest{}); err != nil {
		return err
	}
	_, err := c.httpPut("gpg/key/public", contentType("text/plain"), requestExpects(200), withBody(data))
	return err
}

// SetGPGPrivateKey sets the public key that artifactory provides to debian and opkg clients to verify packages
func (c *Client) SetGPGPrivateKey(data io.Reader) error {
	if err := c.checkRequiredRequestVersion(requests.SetGPGPrivateKeyRequest{}); err != nil {
		return err
	}
	_, err := c.httpPut("gpg/key/private", contentType("text/plain"), requestExpects(200), withBody(data))
	return err
}

// SetGPGPassPhrase sets the public key that artifactory provides to debian and opkg clients to verify packages
func (c *Client) SetGPGPassPhrase(passphrase string) error {
	if err := c.checkRequiredRequestVersion(requests.SetGPGPassphraseRequest{}); err != nil {
		return err
	}
	_, err := c.httpPut("gpg/key/public", contentType("text/plain"), requestExpects(200), addHeader("X-GPG-PASSPHRASE", passphrase))
	return err
}

// TokenGrantType sets the grant_type of the token
func TokenGrantType(t string) TokenOption {
	return func(u *url.Values) error {
		u.Add("grant_type", t)
		return nil
	}
}

// TokenScope is the option for setting a token's scope
func TokenScope(scope string) TokenOption {
	return func(u *url.Values) error {
		u.Add("scope", scope)
		return nil
	}
}

// TokenExpires is the option for setting when a token expires
func TokenExpires(expires int) TokenOption {
	return func(u *url.Values) error {
		e := strconv.Itoa(expires)
		u.Add("expires_in", e)
		return nil
	}
}

// TokenRefreshable is the option for setting a token refreshable
func TokenRefreshable(r bool) TokenOption {
	return func(u *url.Values) error {
		u.Add("refreshable", fmt.Sprintf("%t", r))
		return nil
	}
}

// TokenAudience is a space separate list of other artifactory instances that should accept this token
func TokenAudience(a string) TokenOption {
	return func(u *url.Values) error {
		u.Add("audience", a)
		return nil
	}
}

// TokenUsername is the username for a token.
// This is required for creating new tokens
func TokenUsername(username string) TokenOption {
	return func(u *url.Values) error {
		u.Add("username", username)
		return nil
	}
}

// CreateToken creates an access token
func (c *Client) CreateToken(opts ...TokenOption) (*Token, error) {
	if err := c.checkRequiredRequestVersion(requests.CreateTokenRequest{}); err != nil {
		return nil, err
	}
	u := &url.Values{}
	for _, o := range opts {
		if err := o(u); err != nil {
			return nil, err
		}
	}
	data := &Token{}
	res, err := c.httpPost("security/token",
		withBody(strings.NewReader(u.Encode())),
		contentType("application/x-www-form-urlencoded"),
		requestExpects(200),
		accept("application/json"))
	if err != nil {
		return nil, err
	}
	if jsonErr := json.Unmarshal(res, data); jsonErr != nil {
		return nil, &UnmarshalError{msg: multierror.Append(errDecoding, jsonErr).Error()}
	}
	return data, nil
}

// RefreshToken refreshes an access token
func (c *Client) RefreshToken(accessToken, refreshToken string, opts ...TokenOption) (*Token, error) {
	if err := c.checkRequiredRequestVersion(requests.RefreshTokenRequest{}); err != nil {
		return nil, err
	}
	u := &url.Values{}
	for _, o := range opts {
		if err := o(u); err != nil {
			return nil, err
		}
	}
	u.Add("access_token", accessToken)
	u.Add("refresh_token", refreshToken)
	u.Add("grant_type", "refresh_token")
	data := &Token{}
	res, err := c.httpPost("security/token",
		withBody(strings.NewReader(u.Encode())),
		contentType("application/x-www-form-urlencoded"),
		requestExpects(200),
		accept("application/json"))
	if err != nil {
		return nil, err
	}
	if jsonErr := json.Unmarshal(res, data); jsonErr != nil {
		return nil, &UnmarshalError{msg: multierror.Append(errDecoding, jsonErr).Error()}
	}
	return data, nil
}

// RevokeToken revokes a token
func (c *Client) RevokeToken(token string) error {
	if err := c.checkRequiredRequestVersion(requests.RevokeTokenRequest{}); err != nil {
		return err
	}
	u := &url.Values{}
	u.Add("token", token)
	_, err := c.httpPost("security/token/revoke",
		withBody(strings.NewReader(u.Encode())),
		contentType("application/x-www-form-urlencoded"),
		requestExpects(200),
		accept("application/json"))

	return err
}

// GetServiceID gets the service id of an artifactory instance or cluster
func (c *Client) GetServiceID() (string, error) {
	if err := c.checkRequiredResponseVersion(responses.GetServiceIDResponse{}); err != nil {
		return "", err
	}
	res, err := c.httpGet("system/serviceid", requestExpects(200))
	return string(res), err
}

// GetCertificates gets a list of installed SSL Certificates
func (c *Client) GetCertificates() (Certificates, error) {
	if err := c.checkRequiredResponseVersion(responses.GetCertificatesResponse{}); err != nil {
		return nil, err
	}
	data := Certificates{}
	res, err := c.httpGet("system/security/certificates", requestJSON(), requestExpects(200))
	if err != nil {
		return nil, err
	}
	if jsonErr := json.Unmarshal(res, &data); jsonErr != nil {
		return nil, &UnmarshalError{msg: multierror.Append(errDecoding, jsonErr).Error()}
	}
	return data, nil
}

// AddCertificate adds an SSL Certificate
func (c *Client) AddCertificate(alias string, i io.Reader) error {
	if err := c.checkRequiredRequestVersion(requests.AddCertificateRequest{}); err != nil {
		return err
	}
	_, err := c.httpPost("security/certificates/"+alias, requestExpects(200), withBody(i))
	return err
}

// DeleteCertificate deletes an SSL certificate
func (c *Client) DeleteCertificate(alias string) error {
	if err := c.checkRequiredRequestVersion(requests.DeleteCertificateRequest{}); err != nil {
		return err
	}
	return c.httpDelete("security/certificates/"+alias, requestExpects(200))
}
