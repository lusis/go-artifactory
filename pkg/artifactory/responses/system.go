package responses

// GetVersionAndAddOnResponse is the response for
// https://www.jfrog.com/confluence/display/RTF/System+Settings+JSON#SystemSettingsJSON-application/vnd.org.jfrog.artifactory.system.Version+json
type GetVersionAndAddOnResponse struct {
	Version  string   `json:"version"`
	Revision string   `json:"revision"`
	License  string   `json:"license"`
	AddOns   []string `json:"addons"`
}

func (v GetVersionAndAddOnResponse) minVersion() Version { return versionMustParse("2.2.2") }
func (v GetVersionAndAddOnResponse) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (v GetVersionAndAddOnResponse) deprecated() bool    { return false }

// GetVersionAndAddOnResponseTestData is test data for VersionAndAddOnResponse
const GetVersionAndAddOnResponseTestData = "version.json"

// GetLicenseResponse represents a license response
// https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-LicenseInformation
type GetLicenseResponse struct {
	Type         string `json:"type"`
	ValidThrough string `json:"validThrough"`
	LicensedTo   string `json:"licensedTo"`
}

func (v GetLicenseResponse) minVersion() Version { return versionMustParse("3.3.0") }
func (v GetLicenseResponse) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (v GetLicenseResponse) deprecated() bool    { return false }

// GetLicenseResponseTestData is test data for VersionAndAddOnResponse
const GetLicenseResponseTestData = "license.json"

// GetHALicenseResponseEntry represents an ha license response item
// https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-LicenseInformation
type GetHALicenseResponseEntry struct {
	Type         string `json:"type"`
	ValidThrough string `json:"validThrough"`
	LicensedTo   string `json:"licensedTo"`
	LicenseHash  string `json:"licenseHash"`
	NodeID       string `json:"nodeId"`
	NodeURL      string `json:"nodeUrl"`
	Expired      bool   `json:"expired"`
}

// GetHALicenseResponse is represents an HA License response
type GetHALicenseResponse []GetHALicenseResponseEntry

func (v GetHALicenseResponse) minVersion() Version { return versionMustParse("5.5.0") }
func (v GetHALicenseResponse) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (v GetHALicenseResponse) deprecated() bool    { return false }

// GetHALicenseResponseTestData is test data for VersionAndAddOnResponse
const GetHALicenseResponseTestData = "ha_license.json"

// GetReverseProxyConfigurationResponse represents the web server configuration response
// https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-GetReverseProxyConfiguration
type GetReverseProxyConfigurationResponse struct {
	Key                      string `json:"key"`
	WebServerType            string `json:"webServerType"`
	ArtifactoryAppContext    string `json:"artifactoryAppContext"`
	PublicAppContext         string `json:"publicAppContext"`
	ServerName               string `json:"serverName"`
	ServerNameExpression     string `json:"serverNameExpression"`
	ArtifactoryServerName    string `json:"artifactoryServerName"`
	ArtifactoryPort          int    `json:"artifactoryPort"`
	SSLCertificate           string `json:"sslCertificate"`
	SSLKey                   string `json:"sslKey"`
	DockerReverseProxyMethod string `json:"dockerReverseProxyMethod"`
	UseHTTPS                 bool   `json:"useHttps"`
	UseHTTP                  bool   `json:"useHttp"`
	SSLPort                  int    `json:"sslPort"`
	HTTPPort                 int    `json:"httpPort"`
}

// GetReverseProxyConfigurationResponseTestData is test data for ReverseProxyConfigurationResponse
const GetReverseProxyConfigurationResponseTestData = "webserver.json"

func (v GetReverseProxyConfigurationResponse) minVersion() Version { return versionMustParse("4.3.1") }
func (v GetReverseProxyConfigurationResponse) maxVersion() Version {
	return versionMustParse(CurrentVersion)
}
func (v GetReverseProxyConfigurationResponse) deprecated() bool { return false }

// GetReverseProxySnippetResponse represents a ReverseProxySnippet response
type GetReverseProxySnippetResponse struct{}

func (v GetReverseProxySnippetResponse) minVersion() Version { return versionMustParse("4.3.1") }
func (v GetReverseProxySnippetResponse) maxVersion() Version {
	return versionMustParse(CurrentVersion)
}
func (v GetReverseProxySnippetResponse) deprecated() bool { return false }
