package artifactory

import (
	"encoding/json"
)

// User represents a user in artifactory
type UserShort struct {
	Name string `json:"name"`
	URI  string `json:"uri"`
}

// User represents the details of a user in artifactory
type User struct {
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

func (c *Client) GetUsers() ([]UserShort, error) {
	var res []UserShort
	d, err := c.Get("/api/security/users", make(map[string]string))
	if err != nil {
		return res, err
	}
	err = json.Unmarshal(d, &res)
	return res, err
}

func (c *Client) CreateUser(key string, u User, q map[string]string) error {
	j, err := json.Marshal(u)
	if err != nil {
		return err
	}
	_, err = c.Put("/api/security/users/"+key, j, q)
	return err
}

func (c *Client) GetUser(key string, q map[string]string) (User, error) {
	var res User
	d, err := c.Get("/api/security/users/"+key, q)
	if err != nil {
		return res, err
	}
	err = json.Unmarshal(d, &res)
	return res, err
}


func (c *Client) UpdateUser(key string, u User, q map[string]string) error {
	j, err := json.Marshal(u)
	if err != nil {
		return err
	}
	_, err = c.Post("/api/security/users/"+key, j, q)
	return err
}

func (c *Client) DeleteUser(key string) error {
	err := c.Delete("/api/security/users/" + key)
	return err
}
