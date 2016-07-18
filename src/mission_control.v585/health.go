package mission_control

import "encoding/json"

type Health struct {
	Data bool `json:"data"`
}

func (c *MissionControlClient) GetSystemHealthCheck() (Health, error) {
	var res Health
	d, e := c.Get("/api/v1/ping", make(map[string]string))
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
