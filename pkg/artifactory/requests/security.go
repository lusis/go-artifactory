package requests

// CreateOrReplaceUserRequest represents a user to be created or updated
// https://www.jfrog.com/confluence/display/RTF/Security+Configuration+JSON
type CreateOrReplaceUserRequest struct {
	Name                     string   `json:"name,omitempty"`
	Email                    string   `json:"email"`
	Password                 string   `json:"password"`
	Admin                    bool     `json:"admin,omitempty"`
	ProfileUpdatable         bool     `json:"profileUpdatable,omitempty"`
	DisableUIAccess          bool     `json:"disableUIAccess,omitempty"`
	InternalPasswordDisabled bool     `json:"internalPasswordDisabled,omitempty"`
	Groups                   []string `json:"groups,omitempty"`
}

func (u CreateOrReplaceUserRequest) minVersion() Version { return versionMustParse("2.4.0") }
func (u CreateOrReplaceUserRequest) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (u CreateOrReplaceUserRequest) deprecated() bool    { return false }

// UpdateUserRequest represents a user to be updated
type UpdateUserRequest struct {
	Name                     string   `json:"name,omitempty"`
	Email                    string   `json:"email,omitempty"`
	Password                 string   `json:"password,omitempty"`
	Admin                    bool     `json:"admin,omitempty"`
	ProfileUpdatable         bool     `json:"profileUpdatable,omitempty"`
	DisableUIAccess          bool     `json:"disableUIAccess,omitempty"`
	InternalPasswordDisabled bool     `json:"internalPasswordDisabled,omitempty"`
	Groups                   []string `json:"groups,omitempty"`
}

func (u UpdateUserRequest) minVersion() Version { return versionMustParse("2.4.0") }
func (u UpdateUserRequest) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (u UpdateUserRequest) deprecated() bool    { return false }

// DeleteUserRequest represents a DeleteUser Request
type DeleteUserRequest struct{}

func (r DeleteUserRequest) minVersion() Version { return versionMustParse("2.4.0") }
func (r DeleteUserRequest) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (r DeleteUserRequest) deprecated() bool    { return false }

// ExpirePasswordForUserRequest represents a ExpirePasswordForUser Request
type ExpirePasswordForUserRequest struct{}

func (r ExpirePasswordForUserRequest) minVersion() Version { return versionMustParse("4.4.2") }
func (r ExpirePasswordForUserRequest) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (r ExpirePasswordForUserRequest) deprecated() bool    { return false }

// ExpirePasswordForMultipleUsersRequest represents a ExpirePasswordForMultipleUsers Request
type ExpirePasswordForMultipleUsersRequest struct{}

func (r ExpirePasswordForMultipleUsersRequest) minVersion() Version { return versionMustParse("4.4.2") }
func (r ExpirePasswordForMultipleUsersRequest) maxVersion() Version {
	return versionMustParse(CurrentVersion)
}
func (r ExpirePasswordForMultipleUsersRequest) deprecated() bool { return false }

// ExpirePasswordForAllUsersRequest represents a ExpirePasswordForAllUsers Request
type ExpirePasswordForAllUsersRequest struct{}

func (r ExpirePasswordForAllUsersRequest) minVersion() Version { return versionMustParse("4.4.2") }
func (r ExpirePasswordForAllUsersRequest) maxVersion() Version {
	return versionMustParse(CurrentVersion)
}
func (r ExpirePasswordForAllUsersRequest) deprecated() bool { return false }

// UnexpirePasswordForUserRequest represents a UnexpirePasswordForUser Request
type UnexpirePasswordForUserRequest struct{}

func (r UnexpirePasswordForUserRequest) minVersion() Version { return versionMustParse("4.4.2") }
func (r UnexpirePasswordForUserRequest) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (r UnexpirePasswordForUserRequest) deprecated() bool    { return false }

// ChangePasswordRequest represents a ChangePassword request
type ChangePasswordRequest struct {
	UserName     string `json:"userName"`
	OldPassword  string `json:"oldPassword"`
	NewPassword1 string `json:"newPassword1"`
	NewPassword2 string `json:"newpassword2"`
}

func (r ChangePasswordRequest) minVersion() Version { return versionMustParse("4.4.2") }
func (r ChangePasswordRequest) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (r ChangePasswordRequest) deprecated() bool    { return false }

// SetPasswordExpirationPolicyRequest represents a SetPasswordExpirationPolicy request
type SetPasswordExpirationPolicyRequest struct {
	Enabled        bool `json:"enabled"`
	PasswordMaxAge int  `json:"passwordMaxAge"`
	NotifyByEmail  bool `json:"notifyByEmail"`
}

func (r SetPasswordExpirationPolicyRequest) minVersion() Version { return versionMustParse("4.4.2") }
func (r SetPasswordExpirationPolicyRequest) maxVersion() Version {
	return versionMustParse(CurrentVersion)
}
func (r SetPasswordExpirationPolicyRequest) deprecated() bool { return false }

// ConfigureUserLockPolicyRequest represents a ConfigureUserLockPolicy request
type ConfigureUserLockPolicyRequest struct {
	Enabled       bool `json:"enabled"`
	LoginAttempts int  `json:"loginAttempts"`
}

func (r ConfigureUserLockPolicyRequest) minVersion() Version { return versionMustParse("4.4") }
func (r ConfigureUserLockPolicyRequest) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (r ConfigureUserLockPolicyRequest) deprecated() bool    { return false }

// UnlockLockedOutUserRequest represents a UnlockLockedOutUser request
type UnlockLockedOutUserRequest struct{}

func (r UnlockLockedOutUserRequest) minVersion() Version { return versionMustParse("4.4") }
func (r UnlockLockedOutUserRequest) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (r UnlockLockedOutUserRequest) deprecated() bool    { return false }

// UnlockLockedOutUsersRequest represents a UnlockLockedOutUsers request
type UnlockLockedOutUsersRequest []string

func (r UnlockLockedOutUsersRequest) minVersion() Version { return versionMustParse("4.4") }
func (r UnlockLockedOutUsersRequest) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (r UnlockLockedOutUsersRequest) deprecated() bool    { return false }

// UnlockAllLockedOutUsersRequest represents a UnlockAllLockedOutUsers request
type UnlockAllLockedOutUsersRequest struct{}

func (r UnlockAllLockedOutUsersRequest) minVersion() Version { return versionMustParse("4.4") }
func (r UnlockAllLockedOutUsersRequest) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (r UnlockAllLockedOutUsersRequest) deprecated() bool    { return false }

// CreateAPIKeyRequest represents a CreateAPIKey request
type CreateAPIKeyRequest struct{}

func (r CreateAPIKeyRequest) minVersion() Version { return versionMustParse("4.3.0") }
func (r CreateAPIKeyRequest) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (r CreateAPIKeyRequest) deprecated() bool    { return false }

// RegenerateAPIKeyRequest represents a RegenerateAPIKey request
type RegenerateAPIKeyRequest struct{}

func (r RegenerateAPIKeyRequest) minVersion() Version { return versionMustParse("4.3.0") }
func (r RegenerateAPIKeyRequest) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (r RegenerateAPIKeyRequest) deprecated() bool    { return false }

// RevokeAPIKeyRequest represents a RemoveAPIKey request
type RevokeAPIKeyRequest struct{}

func (r RevokeAPIKeyRequest) minVersion() Version { return versionMustParse("4.3.0") }
func (r RevokeAPIKeyRequest) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (r RevokeAPIKeyRequest) deprecated() bool    { return false }

// RevokeUserAPIKeyRequest represents a RevokeUserAPIKey request
type RevokeUserAPIKeyRequest struct{}

func (r RevokeUserAPIKeyRequest) minVersion() Version { return versionMustParse("4.3.0") }
func (r RevokeUserAPIKeyRequest) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (r RevokeUserAPIKeyRequest) deprecated() bool    { return false }

// RevokeAllAPIKeysRequest represents a RevokeAllAPIKeys request
type RevokeAllAPIKeysRequest struct{}

func (r RevokeAllAPIKeysRequest) minVersion() Version { return versionMustParse("4.3.0") }
func (r RevokeAllAPIKeysRequest) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (r RevokeAllAPIKeysRequest) deprecated() bool    { return false }

// CreateOrReplaceGroupRequest represents a CreateOrReplaceGroup request
type CreateOrReplaceGroupRequest struct {
	Name            string `json:"name,omitempty"`
	Description     string `json:"description,omitempty"`
	AutoJoin        bool   `json:"autoJoin,omitempty"`
	AdminPrivileges bool   `json:"adminPrivileges,omitempty"`
	Realm           string `json:"realm,omitempty"`
	RealmAttributes string `json:"realmAttributes,omitempty"`
}

func (r CreateOrReplaceGroupRequest) minVersion() Version { return versionMustParse("2.4.0") }
func (r CreateOrReplaceGroupRequest) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (r CreateOrReplaceGroupRequest) deprecated() bool    { return false }

// UpdateGroupRequest represents a UpdateGroup request
type UpdateGroupRequest struct {
	Name            string `json:"name,omitempty"`
	Description     string `json:"description,omitempty"`
	AutoJoin        bool   `json:"autoJoin,omitempty"`
	AdminPrivileges bool   `json:"adminPrivileges,omitempty"`
	Realm           string `json:"realm,omitempty"`
	RealmAttributes string `json:"realmAttributes,omitempty"`
}

func (r UpdateGroupRequest) minVersion() Version { return versionMustParse("2.4.0") }
func (r UpdateGroupRequest) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (r UpdateGroupRequest) deprecated() bool    { return false }

// DeleteGroupRequest represents a DeleteGroup request
type DeleteGroupRequest struct{}

func (r DeleteGroupRequest) minVersion() Version { return versionMustParse("2.4.0") }
func (r DeleteGroupRequest) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (r DeleteGroupRequest) deprecated() bool    { return false }

// CreateOrReplacePermissionTargetRequest represents a CreateOrReplacePermissionTarget request
type CreateOrReplacePermissionTargetRequest struct {
	Repositories    []string                       `json:"repositories"`
	Name            string                         `json:"name,omitempty"`
	ExcludesPattern string                         `json:"excludesPattern, omitempty"`
	IncludesPattern string                         `json:"includesPattern,omitempty"`
	Principals      map[string]map[string][]string `json:"principals,omitempty"`
}

func (r CreateOrReplacePermissionTargetRequest) minVersion() Version { return versionMustParse("2.4.0") }
func (r CreateOrReplacePermissionTargetRequest) maxVersion() Version {
	return versionMustParse(CurrentVersion)
}
func (r CreateOrReplacePermissionTargetRequest) deprecated() bool { return false }

// DeletePermissionTargetRequest represents a DeletePermissionTarget request
type DeletePermissionTargetRequest struct{}

func (r DeletePermissionTargetRequest) minVersion() Version { return versionMustParse("2.4.0") }
func (r DeletePermissionTargetRequest) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (r DeletePermissionTargetRequest) deprecated() bool    { return false }

// ActivateArtifactoryKeyEncryptionRequest represents a ActivateArtifactoryKeyEncryption request
type ActivateArtifactoryKeyEncryptionRequest struct{}

func (r ActivateArtifactoryKeyEncryptionRequest) minVersion() Version {
	return versionMustParse("3.2.2")
}
func (r ActivateArtifactoryKeyEncryptionRequest) maxVersion() Version {
	return versionMustParse(CurrentVersion)
}
func (r ActivateArtifactoryKeyEncryptionRequest) deprecated() bool { return false }

// DeactivateArtifactoryKeyEncryptionRequest represents a DeactivateArtifactoryKeyEncryption request
type DeactivateArtifactoryKeyEncryptionRequest struct{}

func (r DeactivateArtifactoryKeyEncryptionRequest) minVersion() Version {
	return versionMustParse("3.2.2")
}
func (r DeactivateArtifactoryKeyEncryptionRequest) maxVersion() Version {
	return versionMustParse(CurrentVersion)
}
func (r DeactivateArtifactoryKeyEncryptionRequest) deprecated() bool { return false }

// SetGPGPublicKeyRequest represents a SetGPGPublicKey request
type SetGPGPublicKeyRequest struct{}

func (r SetGPGPublicKeyRequest) minVersion() Version { return versionMustParse("3.3") }
func (r SetGPGPublicKeyRequest) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (r SetGPGPublicKeyRequest) deprecated() bool    { return false }

// SetGPGPrivateKeyRequest represents a SetGPGPrivateKey request
type SetGPGPrivateKeyRequest struct{}

func (r SetGPGPrivateKeyRequest) minVersion() Version { return versionMustParse("3.3") }
func (r SetGPGPrivateKeyRequest) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (r SetGPGPrivateKeyRequest) deprecated() bool    { return false }

// SetGPGPassphraseRequest represents a SetGPGPassphrase request
type SetGPGPassphraseRequest struct{}

func (r SetGPGPassphraseRequest) minVersion() Version { return versionMustParse("3.3") }
func (r SetGPGPassphraseRequest) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (r SetGPGPassphraseRequest) deprecated() bool    { return false }

// CreateTokenRequest represents a CreateToken request
type CreateTokenRequest struct{}

func (r CreateTokenRequest) minVersion() Version { return versionMustParse("5.0.0") }
func (r CreateTokenRequest) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (r CreateTokenRequest) deprecated() bool    { return false }

// RefreshTokenRequest represents a RefreshToken request
type RefreshTokenRequest struct{}

func (r RefreshTokenRequest) minVersion() Version { return versionMustParse("5.0.0") }
func (r RefreshTokenRequest) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (r RefreshTokenRequest) deprecated() bool    { return false }

// RevokeTokenRequest represents a RevokeToken request
type RevokeTokenRequest struct{}

func (r RevokeTokenRequest) minVersion() Version { return versionMustParse("5.0.0") }
func (r RevokeTokenRequest) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (r RevokeTokenRequest) deprecated() bool    { return false }

// AddCertificateRequest represents a AddCertificate request
type AddCertificateRequest struct{}

func (r AddCertificateRequest) minVersion() Version { return versionMustParse("5.4.0") }
func (r AddCertificateRequest) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (r AddCertificateRequest) deprecated() bool    { return false }

// DeleteCertificateRequest represents a DeleteCertificate request
type DeleteCertificateRequest struct{}

func (r DeleteCertificateRequest) minVersion() Version { return versionMustParse("5.4.0") }
func (r DeleteCertificateRequest) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (r DeleteCertificateRequest) deprecated() bool    { return false }
