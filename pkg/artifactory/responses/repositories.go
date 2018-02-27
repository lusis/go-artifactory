package responses

// GetRepositoryReplicationConfigurationResponse represents a GetRepositoryReplicationConfiguration response
type GetRepositoryReplicationConfigurationResponse struct {
	URL                    string `json:"url"`
	SocketTimeoutMillis    int    `json:"socketTimeoutMillis"`
	Username               string `json:"username"`
	Password               string `json:"password"`
	EnableEventReplication bool   `json:"enableEventReplication"`
	Enabled                bool   `json:"enabled"`
	CronExp                string `json:"cronExp"`
	SyncDeletes            bool   `json:"syncDeletes"`
	SyncProperties         bool   `json:"syncProperties"`
	SyncStatistics         bool   `json:"syncStatistics"`
	RepoKey                string `json:"repoKey"`
	PathPrefix             string `json:"pathPrefix"`
}

// GetRepositoryReplicationConfigurationResponseTestData is test data for GetRepositoryReplicationConfigurationResponse
const GetRepositoryReplicationConfigurationResponseTestData = "get_repository_replication_configuration.json"

func (r GetRepositoryReplicationConfigurationResponse) minVersion() Version {
	return versionMustParse("3.1.1")
}
func (r GetRepositoryReplicationConfigurationResponse) maxVersion() Version {
	return versionMustParse(CurrentVersion)
}
func (r GetRepositoryReplicationConfigurationResponse) deprecated() bool { return false }

// ScheduledReplicationStatusResponse represents a ScheduledReplicationStatus response
type ScheduledReplicationStatusResponse struct {
	Status        string       `json:"status"`
	LastCompleted *ISO8601Time `json:"lastCompleted"`
	Targets       []struct {
		URL           string       `json:"url"`
		RepoKey       string       `json:"repoKey"`
		Status        string       `json:"status"`
		LastCompleted *ISO8601Time `json:"lastCompleted"`
	} `json:"targets"`
	Repositories map[string]struct {
		Status        string       `json:"status"`
		LastCompleted *ISO8601Time `json:"lastCompleted"`
	} `json:"repositories"`
}

// ScheduledReplicationStatusResponseTestData is test data for ScheduledReplicationStatusResponse
const ScheduledReplicationStatusResponseTestData = "scheduled_replication_status.json"

func (r ScheduledReplicationStatusResponse) minVersion() Version { return versionMustParse("2.4.2") }
func (r ScheduledReplicationStatusResponse) maxVersion() Version {
	return versionMustParse(CurrentVersion)
}
func (r ScheduledReplicationStatusResponse) deprecated() bool { return false }

// GetGlobalSystemReplicationConfigurationResponse represents a GetGlobalSystemReplicationConfiguration response
type GetGlobalSystemReplicationConfigurationResponse struct {
	BlockPullReplications bool `json:"blockPullReplications"`
	BlockPushReplications bool `json:"blockPushReplications"`
}

// GetGlobalSystemReplicationConfigurationResponseTestData is test data for GetGlobalSystemReplicationConfigurationResponse
const GetGlobalSystemReplicationConfigurationResponseTestData = "global_system_replication_configuration.json"

func (r GetGlobalSystemReplicationConfigurationResponse) minVersion() Version {
	return versionMustParse("4.7.2")
}
func (r GetGlobalSystemReplicationConfigurationResponse) maxVersion() Version {
	return versionMustParse(CurrentVersion)
}
func (r GetGlobalSystemReplicationConfigurationResponse) deprecated() bool { return false }

// GetRepositoriesResponse represents a GetRepositories response
type GetRepositoriesResponse []struct {
	Key         string `json:"key"`
	Type        string `json:"type"`
	Description string `json:"description"`
	URL         string `json:"url"`
}

// GetRepositoriesResponseTestData is test data for GetRepositoriesResponse
const GetRepositoriesResponseTestData = "get_repositories.json"

func (r GetRepositoriesResponse) minVersion() Version { return versionMustParse("2.2.0") }
func (r GetRepositoriesResponse) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (r GetRepositoriesResponse) deprecated() bool    { return false }

// RepositoryResponse is an interface for all repo types
type RepositoryResponse interface {
	GetType() string
}

// CommonRepositoryResponse represents the common elements in a repository  configuration response
type CommonRepositoryResponse struct {
	Key                      string `json:"key"`
	PackageType              string `json:"packageType"`
	RClass                   string `json:"rclass"`
	Description              string `json:"description"`
	Notes                    string `json:"notes"`
	IncludesPattern          string `json:"includesPattern"`
	ExcludesPattern          string `json:"excludesPattern"`
	RepoLayoutRef            string `json:"repoLayoutRef"`
	DebianTrivialLayout      bool   `json:"debianTrivialLayout,omitEmpty"`
	EnableComposerSupport    bool   `json:"enableComposerSupport,omitempty"`
	EnableNuGetSupport       bool   `json:"enableNuGetSupport,omitempty"`
	EnableGemsSupport        bool   `json:"enableGemsSupport,omitempty"`
	EnableNPMSupport         bool   `json:"enableNpmSupport,omitempty"`
	EnableBowerSupport       bool   `json:"enableBowerSupport,omitempty"`
	EnableCocoaPodsSupport   bool   `json:"enableCocoaPodsSupport,omitempty"`
	EnableConanSupport       bool   `json:"enableConanSupport,omitempty"`
	EnableDebianSupport      bool   `json:"enableDebianSupport,omitempty"`
	EnablePyPiSupport        bool   `json:"enablePypiSupport,omitempty"`
	EnablePuppetSupport      bool   `json:"enablePuppetSupport,omitempty"`
	EnableDockerSupport      bool   `json:"enableDockerSupport,omitempty"`
	EnableVagrantSupport     bool   `json:"enableVagrantSupport,omitempty"`
	EnableGitLFSSupport      bool   `json:"enableGitLfsSupport,omitempty"`
	EnableDistRepoSupport    bool   `json:"enableDistRepoSupport,omitempty"`
	EnabledChefSupport       bool   `json:"enabledChefSupport,omitempty"`
	DockerAPIVersion         string `json:"dockerApiVersion,omitempty"`
	ForceNuGetAuthentication bool   `json:"forceNugetAuthentication,omitempty"`
}

// GetType returns the type of the repository
func (r CommonRepositoryResponse) GetType() string {
	return r.RClass
}

func (r CommonRepositoryResponse) minVersion() Version { return versionMustParse("2.3.0") }
func (r CommonRepositoryResponse) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (r CommonRepositoryResponse) deprecated() bool    { return false }

// LocalRepositoryConfigurationResponse represents a LocalRepositoryConfiguration response
type LocalRepositoryConfigurationResponse struct {
	Key                          string   `json:"key"`
	PackageType                  string   `json:"packageType"`
	RClass                       string   `json:"rclass"`
	Description                  string   `json:"description"`
	Notes                        string   `json:"notes"`
	IncludesPattern              string   `json:"includesPattern"`
	ExcludesPattern              string   `json:"excludesPattern"`
	RepoLayoutRef                string   `json:"repoLayoutRef"`
	DebianTrivialLayout          bool     `json:"debianTrivialLayout,omitEmpty"`
	EnableComposerSupport        bool     `json:"enableComposerSupport,omitempty"`
	EnableNuGetSupport           bool     `json:"enableNuGetSupport,omitempty"`
	EnableGemsSupport            bool     `json:"enableGemsSupport,omitempty"`
	EnableNPMSupport             bool     `json:"enableNpmSupport,omitempty"`
	EnableBowerSupport           bool     `json:"enableBowerSupport,omitempty"`
	EnableCocoaPodsSupport       bool     `json:"enableCocoaPodsSupport,omitempty"`
	EnableConanSupport           bool     `json:"enableConanSupport,omitempty"`
	EnableDebianSupport          bool     `json:"enableDebianSupport,omitempty"`
	EnablePyPiSupport            bool     `json:"enablePypiSupport,omitempty"`
	EnablePuppetSupport          bool     `json:"enablePuppetSupport,omitempty"`
	EnableDockerSupport          bool     `json:"enableDockerSupport,omitempty"`
	EnableVagrantSupport         bool     `json:"enableVagrantSupport,omitempty"`
	EnableGitLFSSupport          bool     `json:"enableGitLfsSupport,omitempty"`
	EnableDistRepoSupport        bool     `json:"enableDistRepoSupport,omitempty"`
	EnabledChefSupport           bool     `json:"enabledChefSupport,omitempty"`
	DockerAPIVersion             string   `json:"dockerApiVersion,omitempty"`
	ForceNuGetAuthentication     bool     `json:"forceNugetAuthentication,omitempty"`
	PropertySets                 []string `json:"propertySets"`
	HandleReleases               bool     `json:"handleReleases"`
	HandleSnapshots              bool     `json:"handleSnapshots"`
	ArchiveBrowsingEnabled       bool     `json:"archiveBrowsingEnabled"`
	BlackedOut                   bool     `json:"blackedOut"`
	BlockXRayUnscannedArtifacts  bool     `json:"blockXrayUnscannedArtifacts"`
	CalculateYumMetaData         bool     `json:"calculateYumMetadata"`
	ChecksumPolicyType           string   `json:"checksumPolicyType"`
	MaxUniqueSnapshots           int      `json:"maxUniqueSnapshots"`
	MaxUniqueTags                int      `json:"maxUniqueTags"`
	SnapshotVersionBehaviour     string   `json:"snapshotVersionBehavior"`
	SuppressPomConsistencyChecks bool     `json:"suppressPomConsistencyChecks"`
	XRayIndex                    bool     `json:"xrayIndex"`
	XrayMinimumBlockedSeverity   string   `json:"xrayMinimumBlockedSeverity"`
	YumRootDepth                 int      `json:"yumRootDepth"`
	EnableFileListsIndexing      bool     `json:"enableFileListsIndexing"`
}

func (r LocalRepositoryConfigurationResponse) minVersion() Version { return versionMustParse("2.3.0") }
func (r LocalRepositoryConfigurationResponse) maxVersion() Version {
	return versionMustParse(CurrentVersion)
}
func (r LocalRepositoryConfigurationResponse) deprecated() bool { return false }

// LocalRepositoryConfigurationResponseTestData is test data for LocalRepositoryConfigurationResponse
const LocalRepositoryConfigurationResponseTestData = "local_repo.json"

// RemoteRepositoryConfigurationResponse represents a RemoteRepositoryConfiguration response
type RemoteRepositoryConfigurationResponse struct {
	Key                               string                 `json:"key"`
	PackageType                       string                 `json:"packageType"`
	RClass                            string                 `json:"rclass"`
	Description                       string                 `json:"description"`
	Notes                             string                 `json:"notes"`
	IncludesPattern                   string                 `json:"includesPattern"`
	ExcludesPattern                   string                 `json:"excludesPattern"`
	RepoLayoutRef                     string                 `json:"repoLayoutRef"`
	DebianTrivialLayout               bool                   `json:"debianTrivialLayout,omitEmpty"`
	EnableComposerSupport             bool                   `json:"enableComposerSupport,omitempty"`
	EnableNuGetSupport                bool                   `json:"enableNuGetSupport,omitempty"`
	EnableGemsSupport                 bool                   `json:"enableGemsSupport,omitempty"`
	EnableNPMSupport                  bool                   `json:"enableNpmSupport,omitempty"`
	EnableBowerSupport                bool                   `json:"enableBowerSupport,omitempty"`
	EnableCocoaPodsSupport            bool                   `json:"enableCocoaPodsSupport,omitempty"`
	EnableConanSupport                bool                   `json:"enableConanSupport,omitempty"`
	EnableDebianSupport               bool                   `json:"enableDebianSupport,omitempty"`
	EnablePyPiSupport                 bool                   `json:"enablePypiSupport,omitempty"`
	EnablePuppetSupport               bool                   `json:"enablePuppetSupport,omitempty"`
	EnableDockerSupport               bool                   `json:"enableDockerSupport,omitempty"`
	EnableVagrantSupport              bool                   `json:"enableVagrantSupport,omitempty"`
	EnableGitLFSSupport               bool                   `json:"enableGitLfsSupport,omitempty"`
	EnableDistRepoSupport             bool                   `json:"enableDistRepoSupport,omitempty"`
	EnabledChefSupport                bool                   `json:"enabledChefSupport,omitempty"`
	DockerAPIVersion                  string                 `json:"dockerApiVersion,omitempty"`
	ForceNuGetAuthentication          bool                   `json:"forceNugetAuthentication,omitempty"`
	URL                               string                 `json:"url"`
	ArchiveBrowsingEnabled            bool                   `json:"archiveBrowsingEnabled"`
	BlackedOut                        bool                   `json:"blackedOut"`
	BlockXRayUnscannedArtifacts       bool                   `json:"blockXrayUnscannedArtifacts"`
	MaxUniqueSnapshots                int                    `json:"maxUniqueSnapshots"`
	MaxUniqueTags                     int                    `json:"maxUniqueTags"`
	SuppressPomConsistencyChecks      bool                   `json:"suppressPomConsistencyChecks"`
	XRayIndex                         bool                   `json:"xrayIndex"`
	XrayMinimumBlockedSeverity        string                 `json:"xrayMinimumBlockedSeverity"`
	RemoteRepoChecksumPolicyType      string                 `json:"remoteRepoChecksumPolicyType"`
	AllowAnyHostAuth                  bool                   `json:"allowAnyHostAuth"`
	AssumedOfflinePeriodSecs          int                    `json:"assumedOfflinePeriodSecs"`
	BlockMismatchingMimeTypes         bool                   `json:"blockMismatchingMimeTypes"`
	BypassHeadRequests                bool                   `json:"bypassHeadRequests"`
	ContentSynchronisation            map[string]interface{} `json:"contentSynchronisation"`
	EnableCookieManagement            bool                   `json:"enableCookieManagement"`
	EnableTokenAuthentication         bool                   `json:"enableTokenAuthentication"`
	FetchJarsEagerly                  bool                   `json:"fetchJarsEagerly"`
	FetchSourcesEagerly               bool                   `json:"fetchSourcesEagerly"`
	HandleReleases                    bool                   `json:"handleReleases"`
	HandleSnapshots                   bool                   `json:"handleSnapshots"`
	HardFail                          bool                   `json:"hardFail"`
	ListRemoteFolderItmes             bool                   `json:"listRemoteFolderItems"`
	LocalAddress                      string                 `json:"localAddress"`
	Username                          string                 `json:"username"`
	Password                          string                 `json:"password"`
	Offline                           bool                   `json:"offline"`
	MismatchingMimeTypesOverrideList  string                 `json:"mismatchingMimeTypesOverrideList"`
	MissedRetrievalCachePeriodSecs    int                    `json:"missedRetrievalCachePeriodSecs"`
	PropagateQueryParams              bool                   `json:"propagateQueryParams"`
	PropertySets                      []string               `json:"propertySets"`
	RejectInvalidJars                 bool                   `json:"rejectInvalidJars"`
	RetrievalCAchePeriodSecs          int                    `json:"retrievalCachePeriodSecs"`
	ShareConfiguration                bool                   `json:"shareConfiguration"`
	SocketTimeoutMillis               int                    `json:"socketTimeoutMillis"`
	StoreArtifactsLocally             bool                   `json:"storeArtifactsLocally"`
	SynchronizeProperties             bool                   `json:"synchronizeProperties"`
	UnusedArtifactsCleanupPeriodHours int                    `json:"unusedArtifactsCleanupPeriodHours"`
	//*CommonRepositoryResponse
}

func (r RemoteRepositoryConfigurationResponse) minVersion() Version { return versionMustParse("2.3.0") }
func (r RemoteRepositoryConfigurationResponse) maxVersion() Version {
	return versionMustParse(CurrentVersion)
}
func (r RemoteRepositoryConfigurationResponse) deprecated() bool { return false }

// RemoteRepositoryConfigurationResponseTestData is test data for RemoteRepositoryConfigurationResponse
const RemoteRepositoryConfigurationResponseTestData = "remote_repo.json"

// VirtualRepositoryConfigurationResponse represents a VirtualRepositoryConfiguration response
type VirtualRepositoryConfigurationResponse struct {
	Key                                           string   `json:"key"`
	PackageType                                   string   `json:"packageType"`
	RClass                                        string   `json:"rclass"`
	Description                                   string   `json:"description"`
	Notes                                         string   `json:"notes"`
	IncludesPattern                               string   `json:"includesPattern"`
	ExcludesPattern                               string   `json:"excludesPattern"`
	RepoLayoutRef                                 string   `json:"repoLayoutRef"`
	DebianTrivialLayout                           bool     `json:"debianTrivialLayout,omitEmpty"`
	EnableComposerSupport                         bool     `json:"enableComposerSupport,omitempty"`
	EnableNuGetSupport                            bool     `json:"enableNuGetSupport,omitempty"`
	EnableGemsSupport                             bool     `json:"enableGemsSupport,omitempty"`
	EnableNPMSupport                              bool     `json:"enableNpmSupport,omitempty"`
	EnableBowerSupport                            bool     `json:"enableBowerSupport,omitempty"`
	EnableCocoaPodsSupport                        bool     `json:"enableCocoaPodsSupport,omitempty"`
	EnableConanSupport                            bool     `json:"enableConanSupport,omitempty"`
	EnableDebianSupport                           bool     `json:"enableDebianSupport,omitempty"`
	EnablePyPiSupport                             bool     `json:"enablePypiSupport,omitempty"`
	EnablePuppetSupport                           bool     `json:"enablePuppetSupport,omitempty"`
	EnableDockerSupport                           bool     `json:"enableDockerSupport,omitempty"`
	EnableVagrantSupport                          bool     `json:"enableVagrantSupport,omitempty"`
	EnableGitLFSSupport                           bool     `json:"enableGitLfsSupport,omitempty"`
	EnableDistRepoSupport                         bool     `json:"enableDistRepoSupport,omitempty"`
	EnabledChefSupport                            bool     `json:"enabledChefSupport,omitempty"`
	DockerAPIVersion                              string   `json:"dockerApiVersion,omitempty"`
	ForceNuGetAuthentication                      bool     `json:"forceNugetAuthentication,omitempty"`
	Repositories                                  []string `json:"repositories"`
	KeyPair                                       string   `json:"keyPair"`
	VirtualRetrievalCachePeriodSecs               int      `json:"virtualRetrievalCachePeriodSecs"`
	ExternalDependenciesEnabled                   bool     `json:"externalDependenciesEnabled"`
	ArtifactoryRequestsCanRetrieveRemoteArtifacts bool     `json:"artifactoryRequestsCanRetrieveRemoteArtifacts"`
	PomRepositoryReferencesCleanupPolicy          string   `json:"pomRepositoryReferencesCleanupPolicy"`
}

func (r VirtualRepositoryConfigurationResponse) minVersion() Version { return versionMustParse("2.3.0") }
func (r VirtualRepositoryConfigurationResponse) maxVersion() Version {
	return versionMustParse(CurrentVersion)
}
func (r VirtualRepositoryConfigurationResponse) deprecated() bool { return false }

// VirtualRepositoryConfigurationResponseTestData is test data for VirtualRepositoryConfigurationResponse
const VirtualRepositoryConfigurationResponseTestData = "virtual_repo.json"
