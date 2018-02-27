package artifactory

import (
	"encoding/json"

	multierror "github.com/hashicorp/go-multierror"
	"github.com/lusis/go-artifactory/pkg/artifactory/requests"
	"github.com/lusis/go-artifactory/pkg/artifactory/responses"
)

// BackgroundTasks represents Background Tasks
type BackgroundTasks struct {
	responses.GetBackgroundTasksResponse
}

// GetBackgroundTasks returns the list of background tasks
func (c *Client) GetBackgroundTasks() (*BackgroundTasks, error) {
	if err := c.checkRequiredResponseVersion(responses.GetBackgroundTasksResponse{}); err != nil {
		return nil, err
	}
	tasks := &BackgroundTasks{}
	res, err := c.httpGet("api/tasks", requestJSON(), requestExpects(200))
	if err != nil {
		return nil, err
	}
	if jsonErr := json.Unmarshal(res, tasks); jsonErr != nil {
		return nil, &UnmarshalError{msg: multierror.Append(errDecoding, jsonErr).Error()}
	}
	return tasks, nil
}

// EmptyTrashCan empties the trash can
func (c *Client) EmptyTrashCan() error {
	if err := c.checkRequiredRequestVersion(requests.EmptyTrashCanRequest{}); err != nil {
		return err
	}
	_, err := c.httpPost("/api/trash/empty", requestExpects(200))
	return err
}

// DeleteItemFromTrashCan empties the trash can
func (c *Client) DeleteItemFromTrashCan(path string) error {
	if err := c.checkRequiredRequestVersion(requests.DeleteItemFromTrashCanRequest{}); err != nil {
		return err
	}
	return c.httpDelete("/api/trash/empty/"+path, requestExpects(200))
}

// RestoreItemFromTrashCan restores the specified item from the trashcan into the specified path
func (c *Client) RestoreItemFromTrashCan(from, to string) error {
	if err := c.checkRequiredRequestVersion(requests.RestoreItemFromTrashCanRequest{}); err != nil {
		return err
	}
	_, err := c.httpPost("/api/trash/restore/"+from, requestExpects(200), queryParams(map[string]string{"to": to}))
	return err
}
