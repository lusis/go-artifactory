package artifactory

import "encoding/json"

type ArtifactProperty struct {
	Name string `json:"name"`
	Value string `json:"value"`
}

type ArtifactProperties struct {
	ArtifactProperties []ArtifactProperty `json:"artifactProperties"`
}

func (c *ArtifactoryClient) GetArtifactProperties(repoKey string) ([]ArtifactProperty, error) {
	var res ArtifactProperties
	d, e := c.Get("/api/artifactproperties?repoKey=" + repoKey, make(map[string]string))
	if e != nil {
		return nil, e
	} else {
		err := json.Unmarshal(d, &res)
		if err != nil {
			return nil, err
		} else {
			return res.ArtifactProperties, e
		}
	}
}
