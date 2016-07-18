package mission_control

import "encoding/json"

type Instance struct {
	URL string `json:"url"`
	Name string `json:"name"`
}

type Instances struct {
	Instances []Instance `json:"data"`
}

func (c *MissionControlClient) GetInstances() ([]Instance, error) {
	var res Instances
	d, e := c.Get("/api/v1/instances", make(map[string]string))
	if e != nil {
		return nil, e
	} else {
		err := json.Unmarshal(d, &res)
		if err != nil {
			return nil, err
		} else {
			return res.Instances, e
		}
	}
}