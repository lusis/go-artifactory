package responses

// FolderInfoResponse represents a FolderInfo response
type FolderInfoResponse struct {
	Repo         string    `json:"repo"`
	Path         string    `json:"path"`
	Created      *JSONTime `json:"created"`
	LastModified *JSONTime `json:"lastModified"`
	LastUpdated  *JSONTime `json:"lastUpdated"`
	Children     []struct {
		URI    string `json:"uri"`
		Folder bool   `json:"folder"`
	} `json:"children"`
	URI string `json:"uri"`
}

// FolderInfoResponseTestData is test data for FolderInfoResponse
const FolderInfoResponseTestData = "folder_info.json"

func (r FolderInfoResponse) minVersion() Version { return versionMustParse("2.2.0") }
func (r FolderInfoResponse) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (r FolderInfoResponse) deprecated() bool    { return false }

// FileInfoResponse represents a FileInfo response
type FileInfoResponse struct {
	Repo         string    `json:"repo"`
	Path         string    `json:"path"`
	Created      *JSONTime `json:"created"`
	LastModified *JSONTime `json:"lastModified"`
	LastUpdated  *JSONTime `json:"lastUpdated"`
	CreatedBy    string    `json:"createdBy"`
	ModifiedBy   string    `json:"modifiedBy"`
	DownloadURI  string    `json:"downloadUri"`
	MimeType     string    `json:"mimeType"`
	Size         string    `json:"size"`
	Checksums    struct {
		SHA1   string `json:"sha1"`
		MD5    string `json:"md5"`
		SHA256 string `json:"sha256,omitempty"`
	} `json:"checksums"`
	OriginalChecksums struct {
		SHA1   string `json:"sha1"`
		MD5    string `json:"md5"`
		SHA256 string `json:"sha256,omitempty"`
	} `json:"originalChecksums"`
	URI string `json:"uri"`
}

// FileInfoResponseTestData is test data for FileInfoResponse
const FileInfoResponseTestData = "file_info.json"

func (r FileInfoResponse) minVersion() Version { return versionMustParse("2.2.0") }
func (r FileInfoResponse) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (r FileInfoResponse) deprecated() bool    { return false }

// RepoSummaryListEntry is a repo entry for GetStorageSummaryInfoResponse
type RepoSummaryListEntry struct {
	RepoKey      string `json:"repoKey"`
	RepoType     string `json:"repoType"`
	FoldersCount int    `json:"foldersCount"`
	FilesCount   int    `json:"filesCount"`
	UsedSpace    string `json:"usedSpace"`
	ItemsCount   int    `json:"itemsCount"`
	PackageType  string `json:"packageType"`
	Percentage   string `json:"percentage"`
}

// BinariesSummaryEntry is a
type BinariesSummaryEntry struct {
	BinariesCount  string `json:"binariesCount"`
	BinariesSize   string `json:"binariesSize"`
	ArtifactsSize  string `json:"artifactsSize"`
	Optimization   string `json:"optimization"`
	ItemsCount     string `json:"itemsCount"`
	ArtifactsCount string `json:"artifactsCount"`
}

// FileStoreSummaryEntry is a
type FileStoreSummaryEntry struct {
	StorageType      string `json:"storageType"`
	StorageDirectory string `json:"storageDirectory"`
	TotalSpace       string `json:"totalSpace"`
	UsedSpace        string `json:"usedSpace"`
	FreeSpace        string `json:"freeSpace"`
}

// GetStorageSummaryInfoResponse represents a GetStorageSummaryInfo response
type GetStorageSummaryInfoResponse struct {
	StorageSummary struct {
		BinariesSummary         BinariesSummaryEntry   `json:"binariesSummary"`
		RepositoriesSummaryList []RepoSummaryListEntry `json:"repositoriesSummaryList"`
		FileStoreSummary        FileStoreSummaryEntry  `json:"fileStoreSummary"`
	}
	RepositoriesSummaryList []RepoSummaryListEntry `json:"repositoriesSummaryList"`
	BinariesSummary         BinariesSummaryEntry   `json:"binariesSummary"`
	FileStoreSummary        FileStoreSummaryEntry  `json:"fileStoreSummary"`
}

// GetStorageSummaryInfoResponseTestData is test data for GetStorageSummaryInfoResponse
const GetStorageSummaryInfoResponseTestData = "storageinfo.json"

func (r GetStorageSummaryInfoResponse) minVersion() Version { return versionMustParse("4.2.0") }
func (r GetStorageSummaryInfoResponse) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (r GetStorageSummaryInfoResponse) deprecated() bool    { return false }

// ItemLastModifiedResponse represents a ItemLastModified response
type ItemLastModifiedResponse struct {
	URI          string       `json:"uri"`
	LastModified *ISO8601Time `json:"lastModified"`
}

// ItemLastModifiedResponseTestData is test data for ItemLastModifiedResponse
const ItemLastModifiedResponseTestData = "last_modified.json"

func (r ItemLastModifiedResponse) minVersion() Version { return versionMustParse("2.2.5") }
func (r ItemLastModifiedResponse) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (r ItemLastModifiedResponse) deprecated() bool    { return false }

// FileStatisticsResponse represents a FileStatistics response
type FileStatisticsResponse struct {
	URI                  string `json:"uri"`
	DownloadCount        int    `json:"downloadCount"`
	LastDownloaded       int    `json:"lastDownloaded"`
	RemoteDownloadCount  int    `json:"remoteDownloadCount"`
	RemoteLastDownloaded int    `json:"remoteLastDownloaded"`
}

// FileStatisticsResponseTestData is test data for FileStatisticsResponse
const FileStatisticsResponseTestData = "file_stats.json"

func (r FileStatisticsResponse) minVersion() Version { return versionMustParse("3.1.0") }
func (r FileStatisticsResponse) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (r FileStatisticsResponse) deprecated() bool    { return false }

// ItemPropertiesResponse represents a ItemProperties response
type ItemPropertiesResponse struct {
	URI        string              `json:"uri"`
	Properties map[string][]string `json:"properties"`
}

// ItemPropertiesResponseTestData is test data for ItemPropertiesResponse
const ItemPropertiesResponseTestData = "item_properties.json"

func (r ItemPropertiesResponse) minVersion() Version { return versionMustParse("2.2.1") }
func (r ItemPropertiesResponse) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (r ItemPropertiesResponse) deprecated() bool    { return false }

// FileComplianceInfoResponse represents a FileComplianceInfo response
type FileComplianceInfoResponse struct {
	Licenses []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"licenses"`
	Vulnerabilities []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"vulnerabilities"`
}

// FileComplianceInfoResponseTestData is test data for FileComplianceInfoResponse
const FileComplianceInfoResponseTestData = "file_compliance.json"

func (r FileComplianceInfoResponse) minVersion() Version { return versionMustParse("3.0.0") }
func (r FileComplianceInfoResponse) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (r FileComplianceInfoResponse) deprecated() bool    { return false }

// CreateDirectoryResponse represents a CreateDirectory response
type CreateDirectoryResponse struct {
	Repo    string    `json:"repo"`
	Path    string    `json:"path"`
	Created *JSONTime `json:"created"`
	URI     string    `json:"uri"`
}

// CreateDirectoryResponseTestData is test data for CreateDirectoryResponse
const CreateDirectoryResponseTestData = "create_directory.json"

func (r CreateDirectoryResponse) minVersion() Version { return versionMustParse("2.2.0") }
func (r CreateDirectoryResponse) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (r CreateDirectoryResponse) deprecated() bool    { return false }

// DeployArtifactResponse represents a DeployArtifact response
type DeployArtifactResponse struct {
	URI         string    `json:"uri"`
	DownloadURI string    `json:"downloadUri"`
	Repo        string    `json:"repo"`
	Path        string    `json:"path"`
	Created     *JSONTime `json:"created"`
	CreatedBy   string    `json:"createdBy"`
	Size        string    `json:"size"`
	MimeType    string    `json:"mimeType"`
	Checksums   struct {
		SHA1   string `json:"sha1"`
		MD5    string `json:"md5"`
		SHA256 string `json:"sha256,omitempty"`
	} `json:"checksums"`
	OriginalChecksums struct {
		SHA1   string `json:"sha1"`
		MD5    string `json:"md5"`
		SHA256 string `json:"sha256,omitempty"`
	} `json:"originalChecksums"`
}

// DeployArtifactResponseTestData is test data for DeployArtifactResponse
const DeployArtifactResponseTestData = "deploy_artifact.json"

func (r DeployArtifactResponse) minVersion() Version { return versionMustParse("2.2.0") }
func (r DeployArtifactResponse) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (r DeployArtifactResponse) deprecated() bool    { return false }

// PushToBintrayResponse represents a PushToBintray response
type PushToBintrayResponse struct {
	Message string `json:"Message"`
}

// PushToBintrayResponseTestData is test data for PushToBintrayResponse
const PushToBintrayResponseTestData = "bintray_push.json"

func (r PushToBintrayResponse) minVersion() Version { return versionMustParse("3.6.0") }
func (r PushToBintrayResponse) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (r PushToBintrayResponse) deprecated() bool    { return false }

// CopyItemResponse represents a CopyItem response
type CopyItemResponse struct {
	Messages []struct {
		Level   string `json:"level"`
		Message string `json:"message"`
	} `json:"messages"`
}

// CopyItemResponseTestData is test data for CopyItemResponse
const CopyItemResponseTestData = "copy_item.json"

func (r CopyItemResponse) minVersion() Version { return versionMustParse("2.2.2") }
func (r CopyItemResponse) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (r CopyItemResponse) deprecated() bool    { return false }

// MoveItemResponse represents a MoveItem response
type MoveItemResponse struct {
	Messages []struct {
		Level   string `json:"level"`
		Message string `json:"message"`
	} `json:"messages"`
}

// MoveItemResponseTestData is test data for MoveItemResponse
const MoveItemResponseTestData = "move_item.json"

func (r MoveItemResponse) minVersion() Version { return versionMustParse("2.2.2") }
func (r MoveItemResponse) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (r MoveItemResponse) deprecated() bool    { return false }

// FileListItemEntry represents a FileListResponse item
type FileListItemEntry struct {
	URI          string    `json:"uri"`
	Size         int       `json:"size"`
	LastModified *JSONTime `json:"lastModified"`
	Folder       bool      `json:"folder"`
	SHA1         string    `json:"sha1,omitempty"`
	SHA2         string    `json:"sha2,omitempty"`
	MDTimestamps struct {
		Properties *JSONTime `json:"properties"`
	} `json:"mdTimestamps,omitempty"`
}

// FileListResponse represents a FileList response
type FileListResponse struct {
	URI     string              `json:"uri"`
	Created *JSONTime           `json:"created"`
	Files   []FileListItemEntry `json:"files"`
}

// FileListResponseTestData is test data for FileListResponse
const FileListResponseTestData = "file_list.json"

// FileListDeepResponseTestData is test data for FileListResponse with deep option
const FileListDeepResponseTestData = "file_list_deep.json"

// FileListMdTimestampsResponseTestData is test data for FileListResponse with mdTimestamps option
const FileListMdTimestampsResponseTestData = "file_list_md_timestamps.json"

// FileListWithFoldersResponseTestData is test data for FileListResponse with listFolders option
const FileListWithFoldersResponseTestData = "file_list_with_folders.json"

// FileListAllOptsResponseTestData is test data for FileListResponse with all options provided
const FileListAllOptsResponseTestData = "file_list_all_opts.json"

func (r FileListResponse) minVersion() Version { return versionMustParse("2.2.4") }
func (r FileListResponse) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (r FileListResponse) deprecated() bool    { return false }

// EffectiveItemPermissionsResponse represents a EffectiveItemPermissions response
type EffectiveItemPermissionsResponse struct {
	URI        string                         `json:"uri"`
	Principals map[string]map[string][]string `json:"principals"`
}

// EffectiveItemPermissionsResponseTestData is test data for EffectiveItemPermissionsResponse
const EffectiveItemPermissionsResponseTestData = "effective_item_permissions.json"

func (r EffectiveItemPermissionsResponse) minVersion() Version { return versionMustParse("2.3.4") }
func (r EffectiveItemPermissionsResponse) maxVersion() Version {
	return versionMustParse(CurrentVersion)
}
func (r EffectiveItemPermissionsResponse) deprecated() bool { return false }
