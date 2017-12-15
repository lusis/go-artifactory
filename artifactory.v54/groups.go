package artifactory

import (
	"encoding/json"
)

// GroupShort represents the json response for a list of groups in Artifactory
type GroupShort struct {
	Name string `json:"name"`
	URI  string `json:"uri"`
}

// Group represents the json response for a group in artifactory
type Group struct {
	Name            string `json:"name,omitempty"`
	Description     string `json:"description,omitempty"`
	AutoJoin        bool   `json:"autoJoin,omitempty"`
	Admin           bool   `json:"admin,omitempty"`
	Realm           string `json:"realm,omitempty"`
	RealmAttributes string `json:"realmAttributes,omitempty"`
}

func (c *Client) GetGroups() ([]GroupShort, error) {
	var res []GroupShort
	d, err := c.Get("/api/security/groups", make(map[string]string))
	if err != nil {
		return res, err
	}
	err = json.Unmarshal(d, &res)
	return res, err
}

func (c *Client) GetGroup(key string, q map[string]string) (Group, error) {
	var res Group
	d, err := c.Get("/api/security/groups/"+key, q)
	if err != nil {
		return res, err
	}
	err = json.Unmarshal(d, &res)
	return res, err
}

func (c *Client) CreateGroup(key string, g Group, q map[string]string) error {
	j, err := json.Marshal(g)
	if err != nil {
		return err
	}
	_, err = c.Put("/api/security/groups/"+key, j, q)
	return err
}

func (c *Client) UpdateGroup(key string, g Group, q map[string]string) error {
	j, err := json.Marshal(g)
	if err != nil {
		return err
	}
	_, err = c.Post("/api/security/groups/"+key, j, q)
	return err
}
func (c *Client) DeleteGroup(key string) error {
	err := c.Delete("/api/security/groups/"+key)
	return err
}