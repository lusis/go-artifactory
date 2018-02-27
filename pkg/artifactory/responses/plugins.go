package responses

// RetrievePluginItemEntry represents an entry in RetrievePluginInfoResponse
type RetrievePluginItemEntry struct {
	Name        string            `json:"name"`
	Version     string            `json:"version"`
	Description string            `json:"description"`
	Users       []string          `json:"users, omitempty"`
	Groups      []string          `json:"groups, omitempty"`
	Params      map[string]string `json:"params"`
	HTTPMethod  string            `json:"httpMethod, omitempty"`
}

// RetrievePluginInfoResponse represents a a retrieve plugin info response
// https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-RetrievePluginInfo
type RetrievePluginInfoResponse struct {
	Executions []RetrievePluginItemEntry `json:"executions, omitempty"`
	Staging    []RetrievePluginItemEntry `json:"staging, omitempty"`
}

func (v RetrievePluginInfoResponse) minVersion() Version { return versionMustParse("2.5.2") }
func (v RetrievePluginInfoResponse) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (v RetrievePluginInfoResponse) deprecated() bool    { return false }

// RetrievePluginInfoResponseTestData is test data for RetrievePluginInfoResponse
const RetrievePluginInfoResponseTestData = "plugins.json"

// RetrieveNamedPluginInfoResponseTestData is test data for RetrievePluginInfoResponse for a specific plugin type
const RetrieveNamedPluginInfoResponseTestData = "executions_plugins.json"

// RetrieveBuildStagingStrategyResponse represents a build staging strategy response
// https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-RetrieveBuildStagingStrategy
type RetrieveBuildStagingStrategyResponse struct {
	DefaultModuleVersion struct {
		ModuleID        string `json:"moduleId"`
		NextRelease     string `json:"nextRelease"`
		NextDevelopment string `json:"nextDevelopment"`
	} `json:"defaultModuleVersion"`
	VCSConfig struct {
		UseReleaseBranch              bool   `json:"useReleaseBranch"`
		ReleaseBranchName             string `json:"releaseBranchName"`
		CreateTag                     bool   `json:"createTag"`
		TagOrURLName                  string `json:"tagUrlOrName"`
		TagComment                    string `json:"tagComment"`
		NextDevelopmentVersionComment string `json:"nextDevelopmentVersionComment"`
	} `json:"vcsConfig"`
	PromotionConfig struct {
		TargetRepository string `json:"targetRepository"`
		Comment          string `json:"comment"`
		Status           string `json:"status"`
	} `json:"promotionConfig"`
}

// RetrieveBuildStagingStrategyResponseTestData is test data for RetrieveBuildStagingStrategyResponse
const RetrieveBuildStagingStrategyResponseTestData = "build_staging_strategy.json"

func (v RetrieveBuildStagingStrategyResponse) minVersion() Version { return versionMustParse("2.5.2") }
func (v RetrieveBuildStagingStrategyResponse) maxVersion() Version {
	return versionMustParse(CurrentVersion)
}
func (v RetrieveBuildStagingStrategyResponse) deprecated() bool { return false }
