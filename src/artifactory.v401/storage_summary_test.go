package artifactory

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"
)

func TestGetStorageSummary(t *testing.T) {
	responseFile, err := os.Open("assets/test/storage_summary.json")
	if err != nil {
		t.Fatalf("Unable to read test data: %s", err.Error())
	}
	defer responseFile.Close()
	responseBody, _ := ioutil.ReadAll(responseFile)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		w.Write(responseBody)
	}))
	defer server.Close()

	transport := &http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			return url.Parse(server.URL)
		},
	}

	conf := &ClientConfig{
		BaseURL:   "http://127.0.0.1:8080/",
		Username:  "username",
		Password:  "password",
		VerifySSL: false,
		Transport: transport,
	}

	client := NewClient(conf)
	storageSummary, err := client.GetStorageSummary()
	assert.NoError(t, err, "should not return an error")
	assert.Equal(t, storageSummary.BinariesSummary.BinariesCount, "125,726", "BinariesSummary.BinariesCount should be 125,726")
	assert.Equal(t, storageSummary.BinariesSummary.BinariesSize, "3.48 GB", "BinariesSummary.BinariesSize should be 3.48 GB")
	assert.Equal(t, storageSummary.BinariesSummary.ArtifactsSize, "59.77 GB", "BinariesSummary.ArtifactsSize should be 59.77 GB")
	assert.Equal(t, storageSummary.BinariesSummary.Optimization, "5.82%", "BinariesSummary.Optimization should be 5.82%")
	assert.Equal(t, storageSummary.BinariesSummary.ItemsCount, "2,176,580", "BinariesSummary.Optimization should be 2,176,580")
	assert.Equal(t, storageSummary.BinariesSummary.ArtifactsCount, "2,084,408", "BinariesSummary.ArtifactsCount should be 2,084,408")

	assert.Equal(t, storageSummary.FileStoreSummary.StorageType, "filesystem", "FileStoreSummary.StorageType should be filesystem")
	assert.Equal(t, storageSummary.FileStoreSummary.StorageDirectory, "/home/.../artifactory/devenv/.artifactory/data/filestore", "FileStoreSummary.StorageDirectory should be /home/.../artifactory/devenv/.artifactory/data/filestore")
	assert.Equal(t, storageSummary.FileStoreSummary.TotalSpace, "204.28 GB", "FileStoreSummary.TotalSpace should be 204.28 GB")
	assert.Equal(t, storageSummary.FileStoreSummary.UsedSpace, "32.22 GB (15.77%)", "FileStoreSummary.UsedSpace should be 32.22 GB (15.77%)")
	assert.Equal(t, storageSummary.FileStoreSummary.FreeSpace, "172.06 GB (84.23%)", "FileStoreSummary.FreeSpace should be 172.06 GB (84.23%)")

	assert.Equal(t, len(storageSummary.RepositoriesSummaryList), 3, "storageSummary.RepositoriesSummaryList count should be 3")
}

