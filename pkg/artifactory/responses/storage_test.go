package responses

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/lusis/go-artifactory/pkg/artifactory/responses/testdata"
	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/require"
)

func TestStorageResponses(t *testing.T) {
	// Add any new test cases here
	// in the form of
	// {&FooResponse{}:FooResponseTestData}
	storageResponsesTestCases := []map[VersionedResponse]string{
		{&FileComplianceInfoResponse{}: FileComplianceInfoResponseTestData},
		{&FolderInfoResponse{}: FolderInfoResponseTestData},
		{&FileInfoResponse{}: FileInfoResponseTestData},
		{&GetStorageSummaryInfoResponse{}: GetStorageSummaryInfoResponseTestData},
		{&ItemLastModifiedResponse{}: ItemLastModifiedResponseTestData},
		{&FileStatisticsResponse{}: FileStatisticsResponseTestData},
		{&ItemPropertiesResponse{}: ItemPropertiesResponseTestData},
		{&CreateDirectoryResponse{}: CreateDirectoryResponseTestData},
		{&DeployArtifactResponse{}: DeployArtifactResponseTestData},
		{&PushToBintrayResponse{}: PushToBintrayResponseTestData},
		{&CopyItemResponse{}: CopyItemResponseTestData},
		{&MoveItemResponse{}: MoveItemResponseTestData},
		{&FileListResponse{}: FileListResponseTestData},
		{&FileListResponse{}: FileListDeepResponseTestData},
		{&FileListResponse{}: FileListMdTimestampsResponseTestData},
		{&FileListResponse{}: FileListWithFoldersResponseTestData},
		{&FileListResponse{}: FileListAllOptsResponseTestData},
		{&EffectiveItemPermissionsResponse{}: EffectiveItemPermissionsResponseTestData},
	}

	for _, testCase := range storageResponsesTestCases {
		for k, v := range testCase {
			data, err := testdata.GetBytes(v)
			require.NoError(t, err)
			placeholder := make(map[string]interface{})
			_ = json.Unmarshal(data, &placeholder)
			config := newMSDecoderConfig()
			config.Result = k
			decoder, newErr := mapstructure.NewDecoder(config)
			require.NoError(t, newErr)
			dErr := decoder.Decode(placeholder)
			require.NoError(t, dErr, fmt.Sprintf("should parse %s", v))
			require.NotNil(t, k)
			require.Implements(t, (*VersionedResponse)(nil), k)
		}
	}
}
