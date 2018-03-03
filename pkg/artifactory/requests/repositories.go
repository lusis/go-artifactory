package requests

// CalculateYumMetadataRequest represents a CalculateYumMetadata request
type CalculateYumMetadataRequest struct{}

func (r CalculateYumMetadataRequest) minVersion() Version { return versionMustParse("2.3.5") }
func (r CalculateYumMetadataRequest) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (r CalculateYumMetadataRequest) deprecated() bool    { return false }

// CalculateNuGetMetadataRequest represents a CalculateNuGetMetadata request
type CalculateNuGetMetadataRequest struct{}

func (r CalculateNuGetMetadataRequest) minVersion() Version { return versionMustParse("3.0.3") }
func (r CalculateNuGetMetadataRequest) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (r CalculateNuGetMetadataRequest) deprecated() bool    { return false }

// CalculateNPMMetadataRequest represents a CalculateNPMMetadata request
type CalculateNPMMetadataRequest struct{}

func (r CalculateNPMMetadataRequest) minVersion() Version { return versionMustParse("3.2.0") }
func (r CalculateNPMMetadataRequest) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (r CalculateNPMMetadataRequest) deprecated() bool    { return false }

// CalculateMavenIndexRequest represents a CalculateMavenIndex request
type CalculateMavenIndexRequest struct{}

func (r CalculateMavenIndexRequest) minVersion() Version { return versionMustParse("2.5.0") }
func (r CalculateMavenIndexRequest) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (r CalculateMavenIndexRequest) deprecated() bool    { return false }

// CalculateMavenMetadataRequest represents a CalculateMavenMetadata request
type CalculateMavenMetadataRequest struct{}

func (r CalculateMavenMetadataRequest) minVersion() Version { return versionMustParse("3.0.2") }
func (r CalculateMavenMetadataRequest) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (r CalculateMavenMetadataRequest) deprecated() bool    { return false }

// CalculateDebianMetadataRequest represents a CalculateDebianMetadata request
type CalculateDebianMetadataRequest struct{}

func (r CalculateDebianMetadataRequest) minVersion() Version { return versionMustParse("3.3") }
func (r CalculateDebianMetadataRequest) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (r CalculateDebianMetadataRequest) deprecated() bool    { return false }

// CalculateOpkgMetadataRequest represents a CalculateOpkgMetadata request
type CalculateOpkgMetadataRequest struct{}

func (r CalculateOpkgMetadataRequest) minVersion() Version { return versionMustParse("4.4") }
func (r CalculateOpkgMetadataRequest) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (r CalculateOpkgMetadataRequest) deprecated() bool    { return false }

// CalculateBowerIndexRequest represents a CalculateBowerIndex request
type CalculateBowerIndexRequest struct{}

func (r CalculateBowerIndexRequest) minVersion() Version { return versionMustParse("3.6.0") }
func (r CalculateBowerIndexRequest) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (r CalculateBowerIndexRequest) deprecated() bool    { return false }

// CalculateHelmChartIndexRequest represents a CalculateHelmChartIndex request
type CalculateHelmChartIndexRequest struct{}

func (r CalculateHelmChartIndexRequest) minVersion() Version { return versionMustParse("5.8") }
func (r CalculateHelmChartIndexRequest) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (r CalculateHelmChartIndexRequest) deprecated() bool    { return false }

// CreateLocalRepositoryRequest represents a CreateRepository request
type CreateLocalRepositoryRequest struct {
	Key                             string   `json:"key,omitempty"`
	PackageType                     string   `json:"packageType,omitempty"`
	RClass                          string   `json:"rclass"`
	Description                     string   `json:"description,omitempty"`
	Notes                           string   `json:"notes,omitempty"`
	IncludesPattern                 string   `json:"includesPattern,omitempty"`
	ExcludesPattern                 string   `json:"excludesPattern,omitempty"`
	RepoLayoutRef                   string   `json:"repoLayoutRef,omitempty"`
	DebianTrivialLayout             bool     `json:"debianTrivialLayout,omitempty"`
	ChecksumPolicyType              string   `json:"checksumPolicyType,omitempty"`
	HandleReleases                  bool     `json:"handleReleases,omitempty"`
	HandleSnapshots                 bool     `json:"handleSnapshots,omitempty"`
	MaxUniqueSnapshots              int      `json:"maxUniqueSnapshots,omitempty"`
	MaxUniqueTags                   int      `json:"maxUniqueTags,omitempty"`
	SnapshotVersionBehaviour        string   `json:"snapshotVersionBehavior,omitempty"`
	SuppressPomConsistencyChecks    bool     `json:"suppressPomConsistencyChecks,omitempty"`
	BlackedOut                      bool     `json:"blackedOut,omitempty"`
	PropertySets                    []string `json:"propertySets,omitempty"`
	ArchiveBrowsingEnabled          bool     `json:"archiveBrowsingEnabled,omitempty"`
	CalculateYumMetaData            bool     `json:"calculateYumMetadata,omitempty"`
	YumRootDepth                    int      `json:"yumRootDepth,omitempty"`
	DockerAPIVersion                string   `json:"dockerApiVersion,omitempty"`
	EnableFileListsIndexing         bool     `json:"enableFileListsIndexing,omitempty"`
	OptionalIndexCompressionFormats []string `json:"optionalIndexCompressionFormats,omitempty"`

	ForceNuGetAuthentication    bool   `json:"forceNugetAuthentication,omitempty"`
	BlockXRayUnscannedArtifacts bool   `json:"blockXrayUnscannedArtifacts,omitempty"`
	XRayIndex                   bool   `json:"xrayIndex,omitempty"`
	XrayMinimumBlockedSeverity  string `json:"xrayMinimumBlockedSeverity,omitempty"`
}

func (r CreateLocalRepositoryRequest) minVersion() Version { return versionMustParse("2.3.0") }
func (r CreateLocalRepositoryRequest) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (r CreateLocalRepositoryRequest) deprecated() bool    { return false }

// CreateRemoteRepositoryRequest represents a CreateRemoteRepository request
type CreateRemoteRepositoryRequest struct {
	Key                               string   `json:"key,omitempty"`
	RClass                            string   `json:"rclass"`
	PackageType                       string   `json:"packageType,omitempty"`
	URL                               string   `json:"url"`
	Username                          string   `json:"username,omitempty"`
	Password                          string   `json:"password,omitempty"`
	Proxy                             string   `json:"proxy,omitempty"`
	Description                       string   `json:"description,omitempty"`
	Notes                             string   `json:"notes,omitempty"`
	IncludesPattern                   string   `json:"includesPattern,omitempty"`
	ExcludesPattern                   string   `json:"excludesPattern,omitempty"`
	RepoLayoutRef                     string   `json:"repoLayoutRef,omitempty"`
	HandleReleases                    bool     `json:"handleReleases,omitempty"`
	HandleSnapshots                   bool     `json:"handleSnapshots,omitempty"`
	MaxUniqueSnapshots                int      `json:"maxUniqueSnapshots,omitempty"`
	MaxUniqueTags                     int      `json:"maxUniqueTags,omitempty"`
	SuppressPomConsistencyChecks      bool     `json:"suppressPomConsistencyChecks,omitempty"`
	HardFail                          bool     `json:"hardFail,omitempty"`
	Offline                           bool     `json:"offline,omitempty"`
	BlackedOut                        bool     `json:"blackedOut,omitempty"`
	StoreArtifactsLocally             bool     `json:"storeArtifactsLocally,omitempty"`
	SocketTimeoutMillis               int      `json:"socketTimeoutMillis,omitempty"`
	LocalAddress                      string   `json:"localAddress,omitempty"`
	RetrievalCachePeriodSecs          int      `json:"retrievalCachePeriodSecs,omitempty"`
	FailedRetrievalCachePeriodSecs    int      `json:"failedRetrievalCachePeriodSecs,omitempty"`
	MissedRetrievalCachePeriodSecs    int      `json:"missedRetrievalCachePeriodSecs,omitempty"`
	UnusedArtifactsCleanupPeriodHours int      `json:"unusedArtifactsCleanupPeriodHours,omitempty"`
	UnusedArtifactsCleanupEnabled     bool     `json:"unusedArtifactsCleanupEnabled,omitempty"`
	AssumedOfflinePeriodSecs          int      `json:"assumedOfflinePeriodSecs,omitempty"`
	FetchJarsEagerly                  bool     `json:"fetchJarsEagerly,omitempty"`
	FetchSourcesEagerly               bool     `json:"fetchSourcesEagerly,omitempty"`
	ShareConfiguration                bool     `json:"shareConfiguration,omitempty"`
	SynchronizeProperties             bool     `json:"synchronizeProperties,omitempty"`
	BlockMismatchingMimeTypes         bool     `json:"blockMismatchingMimeTypes,omitempty"`
	PropertySets                      []string `json:"propertySets,omitempty"`
	AllowAnyHostAuth                  bool     `json:"allowAnyHostAuth,omitempty"`
	EnableCookieManagement            bool     `json:"enableCookieManagement,omitempty"`
	BowerRegistryURL                  string   `json:"bowerRegistryUrl,omitempty"`
	VcsType                           string   `json:"vcsType,omitempty"`
	VcsGitProvider                    string   `json:"vcsGitProvider,omitempty"`
	VcsGitDownloadURL                 string   `json:"vcsGitDownloadUrl,omitempty"`
	BypassHeadRequests                bool     `json:"bypassHeadRequests,omitempty"`
	ClientTLSCertificate              string   `json:"clientTlsCertificate,omitempty"`

	// These are only documented in responses but I have a feeling they're options for requests as well
	DebianTrivialLayout              bool                   `json:"debianTrivialLayout,omitEmpty"`
	DockerAPIVersion                 string                 `json:"dockerApiVersion,omitempty"`
	ForceNuGetAuthentication         bool                   `json:"forceNugetAuthentication,omitempty"`
	ArchiveBrowsingEnabled           bool                   `json:"archiveBrowsingEnabled,omitempty"`
	BlockXRayUnscannedArtifacts      bool                   `json:"blockXrayUnscannedArtifacts,omitempty"`
	XRayIndex                        bool                   `json:"xrayIndex,omitempty"`
	XrayMinimumBlockedSeverity       string                 `json:"xrayMinimumBlockedSeverity,omitempty"`
	RemoteRepoChecksumPolicyType     string                 `json:"remoteRepoChecksumPolicyType,omitempty"`
	ContentSynchronisation           map[string]interface{} `json:"contentSynchronisation,omitempty"`
	EnableTokenAuthentication        bool                   `json:"enableTokenAuthentication,omitempty"`
	ListRemoteFolderItmes            bool                   `json:"listRemoteFolderItems,omitempty"`
	MismatchingMimeTypesOverrideList string                 `json:"mismatchingMimeTypesOverrideList,omitempty"`
	PropagateQueryParams             bool                   `json:"propagateQueryParams,omitempty"`
	RejectInvalidJars                bool                   `json:"rejectInvalidJars,omitempty"`
}

func (r CreateRemoteRepositoryRequest) minVersion() Version { return versionMustParse("2.3.0") }
func (r CreateRemoteRepositoryRequest) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (r CreateRemoteRepositoryRequest) deprecated() bool    { return false }

// CreateVirtualRepositoryRequest represents a CreateVirtualRepository request
type CreateVirtualRepositoryRequest struct {
	Key                                           string   `json:"key,omitempty"`
	PackageType                                   string   `json:"packageType"`
	RClass                                        string   `json:"rclass"`
	Description                                   string   `json:"description,omitempty"`
	Notes                                         string   `json:"notes,omitempty"`
	IncludesPattern                               string   `json:"includesPattern,omitempty"`
	ExcludesPattern                               string   `json:"excludesPattern,omitempty"`
	DebianTrivialLayout                           bool     `json:"debianTrivialLayout,omitempty"`
	Repositories                                  []string `json:"repositories,omitempty"`
	KeyPair                                       string   `json:"keyPair,omitempty"`
	VirtualRetrievalCachePeriodSecs               int      `json:"virtualRetrievalCachePeriodSecs,omitempty"`
	ExternalDependenciesEnabled                   bool     `json:"externalDependenciesEnabled,omitempty"`
	ArtifactoryRequestsCanRetrieveRemoteArtifacts bool     `json:"artifactoryRequestsCanRetrieveRemoteArtifacts,omitempty"`
	PomRepositoryReferencesCleanupPolicy          string   `json:"pomRepositoryReferencesCleanupPolicy,omitempty"`
	DefaultDeploymentRepo                         string   `json:"defaultDeploymentRepo,omitempty"`
}

func (r CreateVirtualRepositoryRequest) minVersion() Version { return versionMustParse("2.3.0") }
func (r CreateVirtualRepositoryRequest) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (r CreateVirtualRepositoryRequest) deprecated() bool    { return false }

// UpdateRepositoryConfigurationRequest represents a UpdateRepositoryConfiguration request
type UpdateRepositoryConfigurationRequest struct{}

func (r UpdateRepositoryConfigurationRequest) minVersion() Version { return versionMustParse("2.3.0") }
func (r UpdateRepositoryConfigurationRequest) maxVersion() Version {
	return versionMustParse(CurrentVersion)
}
func (r UpdateRepositoryConfigurationRequest) deprecated() bool { return false }
