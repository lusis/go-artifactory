package artifactory

import (
	"testing"

	"github.com/lusis/go-artifactory/pkg/artifactory/responses"
	"github.com/lusis/go-artifactory/pkg/artifactory/responses/testdata"
	"github.com/stretchr/testify/require"
)

func TestGetBackgroundTasks(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.GetBackgroundTasksResponseTestData)
	require.NoError(t, err)
	client, server, err := newTestClient(jsonfile, "application/json", 200)
	defer server.Close()
	require.NoError(t, err)
	obj, err := client.GetBackgroundTasks()
	require.NoError(t, err)
	require.NotNil(t, obj)
}

func TestEmptyTrashCan(t *testing.T) {
	client, server, err := newTestClient([]byte(""), "application/json", 200)
	defer server.Close()
	require.NoError(t, err)
	err = client.EmptyTrashCan()
	require.NoError(t, err)
}

func TestDeleteItemFromTrashCan(t *testing.T) {
	client, server, err := newTestClient([]byte(""), "application/json", 200)
	defer server.Close()
	require.NoError(t, err)
	err = client.DeleteItemFromTrashCan("foo")
	require.NoError(t, err)
}

func TestRestoreItemFromTrashCan(t *testing.T) {
	client, server, err := newTestClient([]byte(""), "application/json", 200)
	defer server.Close()
	require.NoError(t, err)
	err = client.RestoreItemFromTrashCan("foo", "bar")
	require.NoError(t, err)
}
