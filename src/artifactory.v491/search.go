package artifactory

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Gavc struct {
	GroupID    string
	ArtifactID string
	Version    string
	Classifier string
	Repos      []string
}

func (c *ArtifactoryClient) GAVCSearch(coords *Gavc) (files []FileInfo, e error) {
	url := "/api/search/gavc"
	params := make(map[string]string)
	if &coords.GroupID != nil {
		params["g"] = coords.GroupID
	}
	if &coords.ArtifactID != nil {
		params["a"] = coords.ArtifactID
	}
	if &coords.Version != nil {
		params["v"] = coords.Version
	}
	if &coords.Classifier != nil {
		params["c"] = coords.Classifier
	}
	if &coords.Repos != nil {
		params["repos"] = strings.Join(coords.Repos, ",")
	}
	d, err := c.Get(url, params)
	if err != nil {
		return files, err
	} else {
		var dat GavcSearchResults
		err := json.Unmarshal(d, &dat)
		if err != nil {
			return files, err
		} else {
			files = dat.Results
			return files, nil
		}
	}
}

func (c *ArtifactoryClient) DockerSearch(name string) (files []FileInfo, e error) {
	var request ArtifactoryRequest
	params := make(map[string]string)
	params["docker.repoName"] = fmt.Sprintf("*%s*", name)
	request.Verb = "GET"
	request.Path = "/api/search/prop"
	request.QueryParams = params
	request.ContentType = "application/json"
	data, err := c.HttpRequest(request)
	if err != nil {
		return files, err
	}
	var dat GavcSearchResults
	uerr := json.Unmarshal(data, &dat)
	if uerr != nil {
		return files, uerr
	}
	files = dat.Results
	return files, nil
}
