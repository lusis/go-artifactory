package artifactory

import (
	"encoding/json"
)

type Group struct {
	Name string `json:"name"`
	Uri  string `json:"uri"`
}

type GroupDetails struct {
	Name            string `json:"name,omitempty"`
	Description     string `json:"description,omitempty"`
	AutoJoin        bool   `json:"autoJoin,omitempty"`
	Realm           string `json:"realm,omitempty"`
	RealmAttributes string `json:"realmAttributes,omitempty"`
}

func (c *ArtifactoryClient) GetGroups() ([]Group, error) {
	var res []Group
	d, e := c.Get("/api/security/groups", make(map[string]string))
	if e != nil {
		return res, e
	} else {
		err := json.Unmarshal(d, &res)
		if err != nil {
			return res, err
		} else {
			return res, e
		}
	}
}

func (c *ArtifactoryClient) GetGroupDetails(u string) (GroupDetails, error) {
	var res GroupDetails
	d, e := c.Get("/api/security/groups/"+u, make(map[string]string))
	if e != nil {
		return res, e
	} else {
		err := json.Unmarshal(d, &res)
		if err != nil {
			return res, err
		} else {
			return res, e
		}
	}
}
