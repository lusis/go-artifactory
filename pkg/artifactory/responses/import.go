package responses

// ImportSystemSettingsResponse represents the import system settings response
// https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-ImportSystemSettingsExample
type ImportSystemSettingsResponse struct {
	ImportPath      string `json:"importPath"`
	IncludeMetaData bool   `json:"includeMetadata"`
	Verbose         bool   `json:"verbose"`
	FailOnError     bool   `json:"failOnError"`
	FailIfEmpty     bool   `json:"failIfEmpty"`
}

// ImportSystemSettingsResponseTestData is test data for ImportSystemSettingsResponse
const ImportSystemSettingsResponseTestData = "import_system.json"

func (v ImportSystemSettingsResponse) minVersion() Version { return versionMustParse("2.5.2") }
func (v ImportSystemSettingsResponse) maxVersion() Version {
	return versionMustParse(CurrentVersion)
}
func (v ImportSystemSettingsResponse) deprecated() bool { return false }
