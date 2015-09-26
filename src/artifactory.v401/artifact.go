package artifactory

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func (c *ArtifactoryClient) DeployArtifact(repoKey string, filename string, path string, properties map[string]string) (CreatedStorageItem, error) {
	var res CreatedStorageItem
	var fileProps []string
	var finalUrl string
	finalUrl = "/" + repoKey + "/"
	if &path != nil {
		finalUrl = finalUrl + path
	}
	baseFile := filepath.Base(filename)
	finalUrl = finalUrl + "/" + baseFile
	if len(properties) > 0 {
		finalUrl = finalUrl + ";"
		for k, v := range properties {
			fileProps = append(fileProps, k+"="+v)
		}
		finalUrl = finalUrl + strings.Join(fileProps, ";")
	}
	data, err := os.Open(filename)
	if err != nil {
		return res, err
	}
	defer data.Close()
	b, _ := ioutil.ReadAll(data)
	d, err := c.Put(finalUrl, string(b), make(map[string]string))
	if err != nil {
		return res, err
	} else {
		e := json.Unmarshal(d, &res)
		if e != nil {
			return res, e
		} else {
			return res, nil
		}
	}
}
