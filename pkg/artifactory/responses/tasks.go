package responses

// GetBackgroundTasksResponse represents a GetBackgroundTasks response
type GetBackgroundTasksResponse struct {
	Tasks []struct {
		ID          string `json:"id"`
		Type        string `json:"type"`
		State       string `json:"state"`
		Description string `json:"description"`
		NodeID      string `json:"nodeId"`
	}
}

// GetBackgroundTasksResponseTestData is test data for GetBackgroundTasksResponse
const GetBackgroundTasksResponseTestData = "get_background_tasks.json"

func (r GetBackgroundTasksResponse) minVersion() Version { return versionMustParse("4.4.0") }
func (r GetBackgroundTasksResponse) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (r GetBackgroundTasksResponse) deprecated() bool    { return false }
