package responses

// ExportSystemSettingsResponse represents a ExportSystemSettings response
type ExportSystemSettingsResponse struct {
	ExportPath      string `json:"exportPath"`
	IncludeMetaData bool   `json:"includeMetadata"`
	CreateArchive   bool   `json:"createArchive"`
	BypassFiltering bool   `json:"bypassFiltering"`
	Verbose         bool   `json:"verbose"`
	FailOnError     bool   `json:"failOnError"`
	FailIfEmpty     bool   `json:"failIfEmpty"`
	M2              bool   `json:"m2"`
	Incremental     bool   `json:"incremental"`
	ExcludeContent  bool   `json:"excludeContent"`
}

// ExportSystemSettingsResponseTestData is test data for ExportSystemSettingsResponse
const ExportSystemSettingsResponseTestData = "export_system_settings.json"

func (r ExportSystemSettingsResponse) minVersion() Version { return versionMustParse("2.4.0") }
func (r ExportSystemSettingsResponse) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (r ExportSystemSettingsResponse) deprecated() bool    { return false }
