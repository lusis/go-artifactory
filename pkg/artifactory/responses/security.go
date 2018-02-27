package responses

// UserResponse represents a user response
type UserResponse struct {
	Name string `json:"name"`
	URI  string `json:"uri"`
}

func (u UserResponse) minVersion() Version { return versionMustParse("2.4.0") }
func (u UserResponse) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (u UserResponse) deprecated() bool    { return false }

// UserResponseTestData is test data for a user response
const UserResponseTestData = "user.json"

// GetUserDetailsResponse represents a user detail response
type GetUserDetailsResponse struct {
	Name                     string    `json:"name"`
	Email                    string    `json:"email"`
	Password                 string    `json:"password"`
	Admin                    bool      `json:"admin"`
	ProfileUpdatable         bool      `json:"profileUpdatable"`
	DisableUIAccess          bool      `json:"disableUIAccess"`
	InternalPasswordDisabled bool      `json:"internalPasswordDisabled"`
	LastLoggedIn             *JSONTime `json:"lastLoggedIn"`
	Realm                    string    `json:"realm"`
	Groups                   []string  `json:"groups"`
}

func (u GetUserDetailsResponse) minVersion() Version { return versionMustParse("2.4.0") }
func (u GetUserDetailsResponse) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (u GetUserDetailsResponse) deprecated() bool    { return false }

// GetUserDetailsResponseTestData is test data for a user detail response
const GetUserDetailsResponseTestData = "user_detail.json"

// GetUsersResponse represents a GetUsers response
type GetUsersResponse []UserResponse

// GetUsersResponseTestData is test data for a get users response
const GetUsersResponseTestData = "users.json"

func (u GetUsersResponse) minVersion() Version { return versionMustParse("2.4.0") }
func (u GetUsersResponse) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (u GetUsersResponse) deprecated() bool    { return false }

// GetEncryptedPasswordResponse is a placeholder type for getting the current user's encrypted password
type GetEncryptedPasswordResponse struct{}

func (u GetEncryptedPasswordResponse) minVersion() Version { return versionMustParse("3.3.0") }
func (u GetEncryptedPasswordResponse) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (u GetEncryptedPasswordResponse) deprecated() bool    { return false }

// GetPasswordExpirationPolicyResponse represents a GetPasswordExpirationPolicy response
type GetPasswordExpirationPolicyResponse struct {
	Enabled        bool `json:"enabled"`
	PasswordMaxAge int  `json:"passwordMaxAge"`
	NotifyByEmail  bool `json:"notifyByEmail"`
}

// GetPasswordExpirationPolicyResponseTestData is test data for GetPasswordExpirationPolicyResponse
const GetPasswordExpirationPolicyResponseTestData = "password_expiration_policy.json"

func (r GetPasswordExpirationPolicyResponse) minVersion() Version { return versionMustParse("4.4.2") }
func (r GetPasswordExpirationPolicyResponse) maxVersion() Version {
	return versionMustParse(CurrentVersion)
}
func (r GetPasswordExpirationPolicyResponse) deprecated() bool { return false }

// RetrieveUserLockPolicyResponse represents a RetrieveUserLockPolicy response
type RetrieveUserLockPolicyResponse struct {
	Enabled       bool `json:"enabled"`
	LoginAttempts int  `json:"loginAttempts"`
}

// RetrieveUserLockPolicyResponseTestData is test data for RetrieveUserLockPolicyResponse
const RetrieveUserLockPolicyResponseTestData = "user_lock_policy.json"

func (r RetrieveUserLockPolicyResponse) minVersion() Version { return versionMustParse("4.4") }
func (r RetrieveUserLockPolicyResponse) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (r RetrieveUserLockPolicyResponse) deprecated() bool    { return false }

// CreateAPIKeyResponse represents a CreateAPIKey response
type CreateAPIKeyResponse struct {
	APIKey string `json:"apiKey"`
}

// CreateAPIKeyResponseTestData is test data for CreateAPIKeyResponse
const CreateAPIKeyResponseTestData = "api_key.json"

func (r CreateAPIKeyResponse) minVersion() Version { return versionMustParse("4.3.0") }
func (r CreateAPIKeyResponse) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (r CreateAPIKeyResponse) deprecated() bool    { return false }

// RegenerateAPIKeyResponse represents a RegenerateAPIKey response
type RegenerateAPIKeyResponse struct {
	APIKey string `json:"apiKey"`
}

// RegenerateAPIKeyResponseTestData is test data for RegenerateAPIKeyResponse
const RegenerateAPIKeyResponseTestData = "api_key.json"

func (r RegenerateAPIKeyResponse) minVersion() Version { return versionMustParse("4.3.0") }
func (r RegenerateAPIKeyResponse) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (r RegenerateAPIKeyResponse) deprecated() bool    { return false }

// GetAPIKeyResponse represents a GetAPIKey response
type GetAPIKeyResponse struct {
	APIKey string `json:"apiKey"`
}

// GetAPIKeyResponseTestData is test data for GetAPIKeyResponse
const GetAPIKeyResponseTestData = "api_key.json"

func (r GetAPIKeyResponse) minVersion() Version { return versionMustParse("4.3.0") }
func (r GetAPIKeyResponse) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (r GetAPIKeyResponse) deprecated() bool    { return false }

// GetGroupsResponse represents a GetGroups response
type GetGroupsResponse []struct {
	Name string `json:"name"`
	URI  string `json:"uri"`
}

// GetGroupsResponseTestData is test data for GetGroupsResponse
const GetGroupsResponseTestData = "groups.json"

func (r GetGroupsResponse) minVersion() Version { return versionMustParse("2.4.0") }
func (r GetGroupsResponse) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (r GetGroupsResponse) deprecated() bool    { return false }

// GetGroupDetailsResponse represents a GetGroupDetails response
type GetGroupDetailsResponse struct {
	Name            string `json:"name"`
	Description     string `json:"description"`
	AutoJoin        bool   `json:"autojoin"`
	Realm           string `json:"realm"`
	AdminPrivileges bool   `json:"adminPrivileges"`
}

// GetGroupDetailsResponseTestData is test data for GetGroupDetailsResponse
const GetGroupDetailsResponseTestData = "group.json"

func (r GetGroupDetailsResponse) minVersion() Version { return versionMustParse("2.4.0") }
func (r GetGroupDetailsResponse) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (r GetGroupDetailsResponse) deprecated() bool    { return false }

// GetPermissionTargetsResponse represents a GetPermissionTargets response
type GetPermissionTargetsResponse struct {
	Name string `json:"name"`
	URI  string `json:"uri"`
}

// GetPermissionTargetsResponseTestData is test data for GetPermissionTargetsResponse
const GetPermissionTargetsResponseTestData = "permission_targets.json"

func (r GetPermissionTargetsResponse) minVersion() Version { return versionMustParse("2.4.0") }
func (r GetPermissionTargetsResponse) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (r GetPermissionTargetsResponse) deprecated() bool    { return false }

// GetPermissionTargetDetailsResponse represents a GetPermissionTarget response
type GetPermissionTargetDetailsResponse struct {
	Name            string                         `json:"name"`
	IncludesPattern string                         `json:"includesPattern"`
	ExcludesPattern string                         `json:"excludesPattern"`
	Repositories    []string                       `json:"repositories"`
	Principals      map[string]map[string][]string `json:"principals"`
}

// GetPermissionTargetDetailsResponseTestData is test data for GetPermissionTargetResponse
const GetPermissionTargetDetailsResponseTestData = "permission_target.json"

func (r GetPermissionTargetDetailsResponse) minVersion() Version { return versionMustParse("2.4.0") }
func (r GetPermissionTargetDetailsResponse) maxVersion() Version {
	return versionMustParse(CurrentVersion)
}
func (r GetPermissionTargetDetailsResponse) deprecated() bool { return false }

// CreateTokenResponse represents a CreateToken response
type CreateTokenResponse struct {
	Scope        string `json:"scope"`
	ExpiresIn    int    `json:"expires_in"`
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

// CreateTokenResponseTestData is test data for CreateTokenResponse
const CreateTokenResponseTestData = "create_token.json"

func (r CreateTokenResponse) minVersion() Version { return versionMustParse("5.0.0") }
func (r CreateTokenResponse) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (r CreateTokenResponse) deprecated() bool    { return false }

// RefreshTokenResponse represents a RefreshToken response
type RefreshTokenResponse struct{}

func (r RefreshTokenResponse) minVersion() Version { return versionMustParse("5.0.0") }
func (r RefreshTokenResponse) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (r RefreshTokenResponse) deprecated() bool    { return false }

// RevokeTokenResponse represents a RevokeToken response
type RevokeTokenResponse struct{}

func (r RevokeTokenResponse) minVersion() Version { return versionMustParse("5.0.0") }
func (r RevokeTokenResponse) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (r RevokeTokenResponse) deprecated() bool    { return false }
