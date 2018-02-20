package responses

// ErrorResponse represents an artifactory api error response
type ErrorResponse struct {
	Errors []struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	} `json:"errors"`
}

func (u ErrorResponse) minVersion() Version { return versionMustParse(AbsoluteMinimumVersion) }
func (u ErrorResponse) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (u ErrorResponse) deprecated() bool    { return false }

// ErrorResponseTestData is test data for a user response
const ErrorResponseTestData = "error.json"
