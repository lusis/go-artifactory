package responses

// AllBuildsResponse represents a AllBuilds response
type AllBuildsResponse struct {
	URI    string `json:"uri"`
	Builds []struct {
		URI         string       `json:"uri"`
		LastStarted *ISO8601Time `json:"lastStarted"`
	} `json:"builds"`
}

// AllBuildsResponseTestData is test data for AllBuildsResponse
const AllBuildsResponseTestData = "builds.json"

func (r AllBuildsResponse) minVersion() Version { return versionMustParse("2.2.0") }
func (r AllBuildsResponse) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (r AllBuildsResponse) deprecated() bool    { return false }

// BuildRunsResponse represents a Build response
type BuildRunsResponse struct {
	URI           string `json:"uri"`
	BuildsNumbers []struct {
		URI     string       `json:"uri"`
		Started *ISO8601Time `json:"started"`
	} `json:"buildsNumbers"`
}

// BuildRunsResponseTestData is test data for BuildResponse
const BuildRunsResponseTestData = "build.json"

func (r BuildRunsResponse) minVersion() Version { return versionMustParse("2.2.0") }
func (r BuildRunsResponse) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (r BuildRunsResponse) deprecated() bool    { return false }

// BuildInfoResponse represents a BuildInfo response
type BuildInfoResponse struct {
	URI       string                 `json:"uri"`
	BuildInfo map[string]interface{} `json:"buildInfo"`
}

// BuildInfoResponseTestData is test data for BuildInfoResponse
const BuildInfoResponseTestData = "build_info.json"

func (r BuildInfoResponse) minVersion() Version { return versionMustParse("2.2.0") }
func (r BuildInfoResponse) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (r BuildInfoResponse) deprecated() bool    { return false }

// BuildDiffCommon represents common elements of build diffs
type BuildDiffCommon struct {
	Updated   []map[string]interface{} `json:"updated"`
	Unchanged []map[string]interface{} `json:"unchanged"`
	Removed   []map[string]interface{} `json:"removed"`
	New       []map[string]interface{} `json:"new"`
}

// BuildDiffResponse represents a BuildDiff response
type BuildDiffResponse struct {
	Artifacts    BuildDiffCommon `json:"artifacts"`
	Dependencies BuildDiffCommon `json:"dependencies"`
	Properties   BuildDiffCommon `json:"properties"`
}

// BuildDiffResponseTestData is test data for BuildDiffResponse
const BuildDiffResponseTestData = "build_diff.json"

func (r BuildDiffResponse) minVersion() Version { return versionMustParse("2.6.6") }
func (r BuildDiffResponse) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (r BuildDiffResponse) deprecated() bool    { return false }
