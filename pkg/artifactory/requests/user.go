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
