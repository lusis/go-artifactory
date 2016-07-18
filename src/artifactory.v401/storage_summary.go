package artifactory

import (
	"encoding/json"
)

type BinariesSummary struct {
	BinariesCount string `json:"binariesCount"`
	BinariesSize string `json:"binariesSize"`
	ArtifactsSize string `json:"artifactsSize"`
	Optimization string `json:"optimization"`
	ItemsCount string `json:"itemsCount"`
	ArtifactsCount string `json:"artifactsCount"`
}

type FileStoreSummary struct {
	StorageType string `json:"storageType"`
	StorageDirectory string `json:"storageDirectory"`
	TotalSpace string `json:"totalSpace"`
	UsedSpace string `json:"usedSpace"`
	FreeSpace string `json:"freeSpace"`
}

type RepositorySummary struct {
	RepoKey string `json:"repoKey"`
	RepoType string `json:"repoType"`
	FoldersCount int `json:"foldersCount"`
	FilesCount int `json:"filesCount"`
	UsedSpace string `json:"usedSpace"`
	ItemsCount int `json:"itemsCount"`
	PackageType string `json:"packageType,omitempty"`
	Percentage string `json:"percentage,omitempty"`
}

type StorageSummary struct {
	BinariesSummary BinariesSummary `json:"binariesSummary"`
	FileStoreSummary FileStoreSummary `json:"fileStoreSummary"`
	RepositoriesSummaryList []RepositorySummary `json:"repositoriesSummaryList"`
}

func (c *ArtifactoryClient) GetStorageSummary() (StorageSummary, error) {
	var res StorageSummary
	d, e := c.Get("/api/storageinfo", make(map[string]string))
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