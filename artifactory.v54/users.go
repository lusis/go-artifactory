package artifactory

import (
	"encoding/json"
	"errors"
)

// User represents a user in artifactory
type User struct {
	Name string `json:"name"`
	URI  string `json:"uri"`
}

// UserDetails represents the details of a user in artifactory
type UserDetails struct {
	Name                     string   `json:"name,omitempty"`
	Email                    string   `json:"email"`
	Password                 string   `json:"password"`
	Admin                    bool     `json:"admin,omitempty"`
	ProfileUpdatable         bool     `json:"profileUpdatable,omitempty"`
	DisableUIAccess          bool     `json:"disableUIAccess,omitempty"`
	InternalPasswordDisabled bool     `json:"internalPasswordDisabled,omitempty"`
	LastLoggedIn             string   `json:"lastLoggedIn,omitempty"`
	Realm                    string   `json:"realm,omitempty"`
	Groups                   []string `json:"groups,omitempty"`
}

// GetUsers returns all users
func (c *Client) GetUsers() ([]User, error) {
	var res []User
	d, err := c.Get("/api/security/users", make(map[string]string))
	if err != nil {
		return res, err
	}
	err = json.Unmarshal(d, &res)
	return res, err
}

// GetUserDetails returns details for the named user
func (c *Client) GetUserDetails(key string, q map[string]string) (UserDetails, error) {
	var res UserDetails
	d, err := c.Get("/api/security/users/"+key, q)
	if err != nil {
		return res, err
	}
	err = json.Unmarshal(d, &res)
	return res, err
}

// CreateUser creates a user with the specified details
func (c *Client) CreateUser(key string, u UserDetails, q map[string]string) error {
	if &u.Email == nil || &u.Password == nil {
		return errors.New("Email and password are required to create users")
	}
	j, err := json.Marshal(u)
	if err != nil {
		return err
	}
	_, err = c.Put("/api/security/users/"+key, j, q)
	return err
}

// DeleteUser deletes a user
func (c *Client) DeleteUser(key string) error {
	err := c.Delete("/api/security/users/" + key)
	return err
}

// GetUserEncryptedPassword returns the current user's encrypted password
func (c *Client) GetUserEncryptedPassword() (string, error) {
	d, err := c.Get("/api/security/encryptedPassword", make(map[string]string))
	return string(d), err
}

// GetUserAPIKey returns the current user's api key
func (c *Client) GetUserAPIKey() (string, error) {
	var res UserAPIKey
	d, err := c.Get("/api/security/apiKey", make(map[string]string))
	if err != nil {
		return "", err
	}
	err = json.Unmarshal(d, &res)
	if err != nil {
		return "", err
	}
	return res.APIKey, nil
}

// CreateUserAPIKey creates an apikey for the current user
func (c *Client) CreateUserAPIKey() (string, error) {
	var res UserAPIKey
	d, err := c.Post("/api/security/apiKey", nil, make(map[string]string))
	if err != nil {
		return "", err
	}
	err = json.Unmarshal(d, &res)
	if err != nil {
		return "", err
	}
	return res.APIKey, nil
}
