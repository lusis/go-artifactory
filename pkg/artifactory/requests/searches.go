package requests

// AQLSearchRequest represents a AQLSearch request
type AQLSearchRequest struct{}

func (r AQLSearchRequest) minVersion() Version { return versionMustParse("3.5.0") }
func (r AQLSearchRequest) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (r AQLSearchRequest) deprecated() bool    { return false }

// BuildArtifactsSearchRequest represents a BuildArtifactsSearch request
type BuildArtifactsSearchRequest struct {
	BuildName   string              `json:"buildName"`
	BuildNumber string              `json:"buildNumber"`
	BuildStatus string              `json:"buildStatus,omitempty"`
	Repos       []string            `json:"repos,omitempty"`
	Mappings    []map[string]string `json:"mappings,omitempty"`
}

func (r BuildArtifactsSearchRequest) minVersion() Version { return versionMustParse("2.6.5") }
func (r BuildArtifactsSearchRequest) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (r BuildArtifactsSearchRequest) deprecated() bool    { return false }
