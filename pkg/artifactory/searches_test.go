package artifactory

import (
	"testing"

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
