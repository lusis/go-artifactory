package artifactory

import (
	"fmt"
	"testing"

	"github.com/lusis/go-artifactory/pkg/artifactory/requests"
	"github.com/lusis/go-artifactory/pkg/artifactory/responses"
	"github.com/lusis/go-artifactory/pkg/artifactory/responses/testdata"
	"github.com/stretchr/testify/require"
)

func TestAQLSearch(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.AQLSearchResponseTestData)
	require.NoError(t, err)
	client, server, err := newTestClient(jsonfile, "application/json", 200)
	defer server.Close()
	require.NoError(t, err)
	obj, err := client.AQLSearch(`items.find({"repo":"$eq":"libs-release-local"}`)
	require.NoError(t, err)
	require.NotNil(t, obj)
}

func TestArtifactSearch(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.ArtifactSearchResponseTestData)
	require.NoError(t, err)
	client, server, err := newTestClient(jsonfile, "application/json", 200)
	defer server.Close()
	require.NoError(t, err)
	obj, err := client.ArtifactSearch(map[string]string{})
	require.NoError(t, err)
	require.NotNil(t, obj)
}

func TestArchiveEntriesSearch(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.ArchiveEntriesSearchResponseTestData)
	require.NoError(t, err)
	client, server, err := newTestClient(jsonfile, "application/json", 200)
	defer server.Close()
	require.NoError(t, err)
	obj, err := client.ArchiveEntriesSearch(map[string]string{})
	require.NoError(t, err)
	require.NotNil(t, obj)
}

func TestGAVCSearch(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.GAVCSearchResponseTestData)
	require.NoError(t, err)
	client, server, err := newTestClient(jsonfile, "application/json", 200)
	defer server.Close()
	require.NoError(t, err)
	obj, err := client.GAVCSearch(map[string]string{})
	require.NoError(t, err)
	require.NotNil(t, obj)
}

func TestPropertySearch(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.PropertySearchResponseTestData)
	require.NoError(t, err)
	client, server, err := newTestClient(jsonfile, "application/json", 200)
	defer server.Close()
	require.NoError(t, err)
	obj, err := client.PropertySearch(map[string]string{})
	require.NoError(t, err)
	require.NotNil(t, obj)
}

func TestChecksumSearch(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.ChecksumSearchResponseTestData)
	require.NoError(t, err)
	client, server, err := newTestClient(jsonfile, "application/json", 200)
	defer server.Close()
	require.NoError(t, err)
	obj, err := client.ChecksumSearch(map[string]string{})
	require.NoError(t, err)
	require.NotNil(t, obj)
}

func TestBadChecksumSearch(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.BadChecksumSearchResponseTestData)
	require.NoError(t, err)
	client, server, err := newTestClient(jsonfile, "application/json", 200)
	defer server.Close()
	require.NoError(t, err)
	obj, err := client.BadChecksumSearch(map[string]string{})
	require.NoError(t, err)
	require.NotNil(t, obj)
}

func TestArtifactsNotDownloadedSinceSearch(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.ArtifactsNotDownloadedSinceSearchResponseTestData)
	require.NoError(t, err)
	client, server, err := newTestClient(jsonfile, "application/json", 200)
	defer server.Close()
	require.NoError(t, err)
	obj, err := client.ArtifactsNotDownloadedSinceSearch(map[string]string{})
	require.NoError(t, err)
	require.NotNil(t, obj)
}

func TestArtifactsWithDateInRangeSearch(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.ArtifactsWithDataInRangeSearchResponseTestData)
	require.NoError(t, err)
	client, server, err := newTestClient(jsonfile, "application/json", 200)
	defer server.Close()
	require.NoError(t, err)
	obj, err := client.ArtifactsWithDateInRangeSearch(map[string]string{})
	require.NoError(t, err)
	require.NotNil(t, obj)
}

func TestArtifactsCreatedInDateRangeSearch(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.ArtifactsCreatedInDateRangeSearchResponseTestData)
	require.NoError(t, err)
	client, server, err := newTestClient(jsonfile, "application/json", 200)
	defer server.Close()
	require.NoError(t, err)
	obj, err := client.ArtifactsCreatedInDateRangeSearch(map[string]string{})
	require.NoError(t, err)
	require.NotNil(t, obj)
}

func TestPatternSearch(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.PatternSearchResponseTestData)
	require.NoError(t, err)
	client, server, err := newTestClient(jsonfile, "application/json", 200)
	defer server.Close()
	require.NoError(t, err)
	obj, err := client.PatternSearch("foo")
	require.NoError(t, err)
	require.NotNil(t, obj)
}

func TestBuildsForDependencySearch(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.BuildsForDependencyResponseTestData)
	require.NoError(t, err)
	client, server, err := newTestClient(jsonfile, "application/json", 200)
	defer server.Close()
	require.NoError(t, err)
	obj, err := client.BuildsForDependencySearch(map[string]string{})
	require.NoError(t, err)
	require.NotNil(t, obj)
}

func TestLicenseSearch(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.LicenseSearchResponseTestData)
	require.NoError(t, err)
	client, server, err := newTestClient(jsonfile, "application/json", 200)
	defer server.Close()
	require.NoError(t, err)
	obj, err := client.LicenseSearch(map[string]string{})
	require.NoError(t, err)
	require.NotNil(t, obj)
}

func TestArtifactVersionSearch(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.ArtifactVersionSearchResponseTestData)
	require.NoError(t, err)
	client, server, err := newTestClient(jsonfile, "application/json", 200)
	defer server.Close()
	require.NoError(t, err)
	obj, err := client.ArtifactVersionSearch(map[string]string{})
	require.NoError(t, err)
	require.NotNil(t, obj)
}

func TestArtifactLatestVersionBasedOnLayoutSearch(t *testing.T) {
	client, server, err := newTestClient([]byte("1.0.0"), "application/json", 200)
	defer server.Close()
	require.NoError(t, err)
	obj, err := client.ArtifactLatestVersionBasedOnLayoutSearch(map[string]string{})
	require.NoError(t, err)
	require.NotNil(t, obj)
}

func TestArtifactLatestBasedOnPropertySearch(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.ArtifactLatestVersionBasedOnPropertiesSearchResponseTestData)
	require.NoError(t, err)
	client, server, err := newTestClient(jsonfile, "application/json", 200)
	defer server.Close()
	require.NoError(t, err)
	obj, err := client.ArtifactLatestBasedOnPropertiesSearch("path", map[string]string{})
	require.NoError(t, err)
	require.NotNil(t, obj)
}

func TestBuildArtifactsSearch(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.BuildArtifactsSearchResponseTestData)
	require.NoError(t, err)
	client, server, err := newTestClient(jsonfile, "application/json", 200)
	defer server.Close()
	require.NoError(t, err)
	obj, err := client.BuildArtifactsSearch("foo", "bar")
	require.NoError(t, err)
	require.NotNil(t, obj)
}

func TestBuildArtifactsSearchWithOptions(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.BuildArtifactsSearchResponseTestData)
	require.NoError(t, err)
	client, server, err := newTestClient(jsonfile, "application/json", 200)
	defer server.Close()
	require.NoError(t, err)
	opts := []BuildArtifactsSearchOption{
		BuildArtifactsSearchMappings([]map[string]string{}),
		BuildArtifactsSearchRepos("foo"),
		BuildArtifactsSearchStatus("bar"),
	}
	obj, err := client.BuildArtifactsSearch("foo", "bar", opts...)
	require.NoError(t, err)
	require.NotNil(t, obj)
}
func TestBuildArtifactsSearchWithBadOption(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.BuildArtifactsSearchResponseTestData)
	require.NoError(t, err)
	client, server, err := newTestClient(jsonfile, "application/json", 200)
	defer server.Close()
	require.NoError(t, err)
	badopt := func(r *requests.BuildArtifactsSearchRequest) error {
		return fmt.Errorf("bad option")
	}

	obj, err := client.BuildArtifactsSearch("foo", "bar", badopt)
	require.Error(t, err)
	require.Nil(t, obj)
}

func TestListDockerRepos(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.ListDockerRepositoriesResponseTestData)
	require.NoError(t, err)
	client, server, err := newTestClient(jsonfile, "application/json", 200)
	defer server.Close()
	require.NoError(t, err)
	obj, err := client.ListDockerRepositories("foo")
	require.NoError(t, err)
	require.NotNil(t, obj)
}

func TestListDockerTags(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.ListDockerTagsResponseTestData)
	require.NoError(t, err)
	client, server, err := newTestClient(jsonfile, "application/json", 200)
	defer server.Close()
	require.NoError(t, err)
	obj, err := client.ListDockerTags("foo", "bar")
	require.NoError(t, err)
	require.NotNil(t, obj)
}
