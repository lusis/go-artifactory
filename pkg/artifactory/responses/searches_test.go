package responses

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/lusis/go-artifactory/pkg/artifactory/responses/testdata"
	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/require"
)

func TestSearchesResponses(t *testing.T) {
	// Add any new test cases here
	// in the form of
	// {&FooResponse{}:FooResponseTestData}
	responsesTestCases := []map[VersionedResponse]string{
		{&AQLSearchResponse{}: AQLSearchResponseTestData},
		{&AQLSearchResponse{}: AQLSearchResponseWithArchiveTestData},
		{&ArtifactSearchResponse{}: ArtifactSearchResponseTestData},
		{&ArchiveEntriesSearchResponse{}: ArchiveEntriesSearchResponseTestData},
		{&GAVCSearchResponse{}: GAVCSearchResponseTestData},
		{&PropertySearchResponse{}: PropertySearchResponseTestData},
		{&ChecksumSearchResponse{}: ChecksumSearchResponseTestData},
		{&BadChecksumSearchResponse{}: BadChecksumSearchResponseTestData},
		{&ArtifactsNotDownloadedSinceSearchResponse{}: ArtifactsNotDownloadedSinceSearchResponseTestData},
		{&ArtifactsWithDataInRangeSearchResponse{}: ArtifactsWithDataInRangeSearchResponseTestData},
		{&ArtifactsCreatedInDateRangeSearchResponse{}: ArtifactsCreatedInDateRangeSearchResponseTestData},
		{&PatternSearchResponse{}: PatternSearchResponseTestData},
		{&BuildsForDependencyResponse{}: BuildsForDependencyResponseTestData},
		{&LicenseSearchResponse{}: LicenseSearchResponseTestData},
		{&ArtifactVersionSearchResponse{}: ArtifactVersionSearchResponseTestData},
		{&ArtifactLatestVersionBasedOnPropertiesSearchResponse{}: ArtifactLatestVersionBasedOnPropertiesSearchResponseTestData},
		{&BuildArtifactsSearchResponse{}: BuildArtifactsSearchResponseTestData},
		{&ListDockerRepositoriesResponse{}: ListDockerRepositoriesResponseTestData},
		{&ListDockerTagsResponse{}: ListDockerTagsResponseTestData},
	}
	t.Parallel()
	for _, testCase := range responsesTestCases {
		for k, v := range testCase {
			t.Run(v,
				func(*testing.T) {
					data, err := testdata.GetBytes(v)
					require.NoError(t, err)
					placeholder := make(map[string]interface{})
					_ = json.Unmarshal(data, &placeholder)
					config := newMSDecoderConfig()
					config.Result = k
					decoder, newErr := mapstructure.NewDecoder(config)
					require.NoError(t, newErr)
					dErr := decoder.Decode(placeholder)
					require.NoError(t, dErr, fmt.Sprintf("should parse file: %s", v))
					require.Implements(t, (*VersionedResponse)(nil), k)
				})
		}
	}
}
