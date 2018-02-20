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

// EncryptedPasswordResponse is a placeholder type for getting the current user's encrypted password
type EncryptedPasswordResponse struct{}

func (u EncryptedPasswordResponse) minVersion() Version { return versionMustParse("3.3.0") }
func (u EncryptedPasswordResponse) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (u EncryptedPasswordResponse) deprecated() bool    { return false }
