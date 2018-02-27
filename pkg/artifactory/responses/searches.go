package responses

// AQLSearchResponse represents a AQL response
type AQLSearchResponse struct {
	Results []struct {
		Repo       string    `json:"repo"`
		Path       string    `json:"path"`
		Name       string    `json:"name"`
		Type       string    `json:"type"`
		Size       int       `json:"size"`
		Created    *JSONTime `json:"created"`
		CreatedBy  string    `json:"created_by"`
		Modified   *JSONTime `json:"modified"`
		ModifiedBy string    `json:"modified_by"`
		Updated    *JSONTime `json:"updated"`
		Archives   []struct {
			Entries []struct {
				EntryName string `json:"entry.name"`
				EntryPath string `json:"entry.path"`
			}
		} `json:"archives,omitempty"`
	} `json:"results"`
	Range struct {
		StartPos int `json:"start_pos"`
		EndPos   int `json:"end_pos"`
		Total    int `json:"total"`
	} `json:"range"`
}

// AQLSearchResponseTestData is test data for AQLResponse
const AQLSearchResponseTestData = "aql_search.json"

// AQLSearchResponseWithArchiveTestData is test data for AQLResponse
const AQLSearchResponseWithArchiveTestData = "aql_search_with_archive.json"

func (r AQLSearchResponse) minVersion() Version { return versionMustParse("3.5.0") }
func (r AQLSearchResponse) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (r AQLSearchResponse) deprecated() bool    { return false }

// ArtifactSearchResponse represents a ArtifactSearch response
type ArtifactSearchResponse struct {
	Results []struct {
		URI string `json:"uri"`
	} `json:"results"`
}

// ArtifactSearchResponseTestData is test data for ArtifactSearchResponse
const ArtifactSearchResponseTestData = "artifact_search.json"

func (r ArtifactSearchResponse) minVersion() Version { return versionMustParse("2.2.0") }
func (r ArtifactSearchResponse) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (r ArtifactSearchResponse) deprecated() bool    { return false }

// ArchiveEntriesSearchResponse represents a ArchiveEntriesSearch response
type ArchiveEntriesSearchResponse struct {
	Results []struct {
		Entry       string   `json:"entry"`
		ArchiveURIS []string `json:"archiveUris"`
	} `json:"results"`
}

// ArchiveEntriesSearchResponseTestData is test data for ArchiveEntriesSearchResponse
const ArchiveEntriesSearchResponseTestData = "archive_entries_search.json"

func (r ArchiveEntriesSearchResponse) minVersion() Version { return versionMustParse("2.2.0") }
func (r ArchiveEntriesSearchResponse) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (r ArchiveEntriesSearchResponse) deprecated() bool    { return false }

// GAVCSearchResponse represents a GAVCSearch response
type GAVCSearchResponse struct {
	Results []struct {
		Repo         string              `json:"repo"`
		Path         string              `json:"path"`
		Created      *JSONTime           `json:"created"`
		CreatedBy    string              `json:"createdBy"`
		LastModified *JSONTime           `json:"lastModified"`
		LastUpdated  *JSONTime           `json:"lastUpdated"`
		ModifiedBy   string              `json:"modifiedBy"`
		Properties   map[string][]string `json:"properties,omitempty"`
		DownloadURI  string              `json:"downloadUri"`
		RemoteURL    string              `json:"remoteUrl,omitempty"`
		MimeType     string              `json:"mimeType"`
		Size         string              `json:"size"`
		URI          string              `json:"uri"`
		Checksums    struct {
			SHA1   string `json:"sha1,omitempty"`
			MD5    string `json:"md5,omitempty"`
			SHA256 string `json:"sha256,omitempty"`
		} `json:"checksums"`
		OriginalChecksums struct {
			SHA1   string `json:"sha1,omitempty"`
			MD5    string `json:"md5,omitempty"`
			SHA256 string `json:"sha256,omitempty"`
		} `json:"originalChecksums"`
	} `json:"results"`
}

// GAVCSearchResponseTestData is test data for GAVCSearchResponse
const GAVCSearchResponseTestData = "gavc_search.json"

func (r GAVCSearchResponse) minVersion() Version { return versionMustParse("2.2.0") }
func (r GAVCSearchResponse) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (r GAVCSearchResponse) deprecated() bool    { return false }

// PropertySearchResponse represents a PropertySearch response
type PropertySearchResponse struct {
	Results []struct {
		Repo         string              `json:"repo"`
		Path         string              `json:"path"`
		Created      *JSONTime           `json:"created"`
		CreatedBy    string              `json:"createdBy"`
		LastModified *JSONTime           `json:"lastModified"`
		LastUpdated  *JSONTime           `json:"lastUpdated"`
		ModifiedBy   string              `json:"modifiedBy"`
		Properties   map[string][]string `json:"properties,omitempty"`
		DownloadURI  string              `json:"downloadUri"`
		RemoteURL    string              `json:"remoteUrl,omitempty"`
		MimeType     string              `json:"mimeType"`
		Size         string              `json:"size"`
		URI          string              `json:"uri"`
		Checksums    struct {
			SHA1   string `json:"sha1,omitempty"`
			MD5    string `json:"md5,omitempty"`
			SHA256 string `json:"sha256,omitempty"`
		} `json:"checksums"`
		OriginalChecksums struct {
			SHA1   string `json:"sha1,omitempty"`
			MD5    string `json:"md5,omitempty"`
			SHA256 string `json:"sha256,omitempty"`
		} `json:"originalChecksums"`
	} `json:"results"`
}

// PropertySearchResponseTestData is test data for PropertySearchResponse
const PropertySearchResponseTestData = "property_search.json"

func (r PropertySearchResponse) minVersion() Version { return versionMustParse("2.2.0") }
func (r PropertySearchResponse) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (r PropertySearchResponse) deprecated() bool    { return false }

// ChecksumSearchResponse represents a ChecksumSearch response
type ChecksumSearchResponse struct {
	Results []struct {
		Repo         string              `json:"repo"`
		Path         string              `json:"path"`
		Created      *JSONTime           `json:"created"`
		CreatedBy    string              `json:"createdBy"`
		LastModified *JSONTime           `json:"lastModified"`
		LastUpdated  *JSONTime           `json:"lastUpdated"`
		ModifiedBy   string              `json:"modifiedBy"`
		Properties   map[string][]string `json:"properties,omitempty"`
		DownloadURI  string              `json:"downloadUri"`
		RemoteURL    string              `json:"remoteUrl,omitempty"`
		MimeType     string              `json:"mimeType"`
		Size         string              `json:"size"`
		URI          string              `json:"uri"`
		Checksums    struct {
			SHA1   string `json:"sha1,omitempty"`
			MD5    string `json:"md5,omitempty"`
			SHA256 string `json:"sha256,omitempty"`
		} `json:"checksums"`
		OriginalChecksums struct {
			SHA1   string `json:"sha1,omitempty"`
			MD5    string `json:"md5,omitempty"`
			SHA256 string `json:"sha256,omitempty"`
		} `json:"originalChecksums"`
	} `json:"results"`
}

// ChecksumSearchResponseTestData is test data for ChecksumSearchResponse
const ChecksumSearchResponseTestData = "checksum_search.json"

func (r ChecksumSearchResponse) minVersion() Version { return versionMustParse("2.3.0") }
func (r ChecksumSearchResponse) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (r ChecksumSearchResponse) deprecated() bool    { return false }

// BadChecksumSearchResponse represents a BadChecksumSearch response
type BadChecksumSearchResponse struct {
	Results []struct {
		URI       string `json:"uri"`
		ServerMD5 string `json:"serverMd5"`
		ClientMD5 string `json:"clientMd5"`
	} `json:"results"`
}

// BadChecksumSearchResponseTestData is test data for BadChecksumSearchResponse
const BadChecksumSearchResponseTestData = "bad_checksum_search.json"

func (r BadChecksumSearchResponse) minVersion() Version { return versionMustParse("2.3.4") }
func (r BadChecksumSearchResponse) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (r BadChecksumSearchResponse) deprecated() bool    { return false }

// ArtifactsNotDownloadedSinceSearchResponse represents a ArtifactsNotDownloadedSinceSearch response
type ArtifactsNotDownloadedSinceSearchResponse struct {
	Results []struct {
		URI                  string    `json:"uri"`
		LastDownload         *JSONTime `json:"lastDownloaded"`
		DownloadCount        int       `json:"downloadCount"`
		RemoteLastDownloaded *JSONTime `json:"remoteLastDownloaded"`
		RemoteDownloadCount  int       `json:"remoteDownloadCount"`
	} `json:"results"`
}

// ArtifactsNotDownloadedSinceSearchResponseTestData is test data for ArtifactsNotDownloadedSinceSearchResponse
const ArtifactsNotDownloadedSinceSearchResponseTestData = "artifacts_not_downloaded_since_search.json"

func (r ArtifactsNotDownloadedSinceSearchResponse) minVersion() Version {
	return versionMustParse("2.2.4")
}
func (r ArtifactsNotDownloadedSinceSearchResponse) maxVersion() Version {
	return versionMustParse(CurrentVersion)
}
func (r ArtifactsNotDownloadedSinceSearchResponse) deprecated() bool { return false }

// ArtifactsWithDataInRangeSearchResponse represents a ArtifactsWithDataInRangeSearch response
type ArtifactsWithDataInRangeSearchResponse struct {
	Results []struct {
		URI            string    `json:"uri"`
		Created        *JSONTime `json:"created"`
		LastModified   *JSONTime `json:"lastModified"`
		LastDownloaded *JSONTime `json:"lastDownloaded"`
	} `json:"results"`
}

// ArtifactsWithDataInRangeSearchResponseTestData is test data for ArtifactsWithDataInRangeSearchResponse
const ArtifactsWithDataInRangeSearchResponseTestData = "artifacts_with_date_in_range_search.json"

func (r ArtifactsWithDataInRangeSearchResponse) minVersion() Version { return versionMustParse("3.2.1") }
func (r ArtifactsWithDataInRangeSearchResponse) maxVersion() Version {
	return versionMustParse(CurrentVersion)
}
func (r ArtifactsWithDataInRangeSearchResponse) deprecated() bool { return false }

// ArtifactsCreatedInDateRangeSearchResponse represents a ArtifactsCreatedInDateRangeSearch response
type ArtifactsCreatedInDateRangeSearchResponse struct {
	Results []struct {
		URI     string    `json:"uri"`
		Created *JSONTime `json:"created"`
	} `json:"results"`
}

// ArtifactsCreatedInDateRangeSearchResponseTestData is test data for ArtifactsCreatedInDateRangeSearchResponse
const ArtifactsCreatedInDateRangeSearchResponseTestData = "artifacts_created_in_date_range_search.json"

func (r ArtifactsCreatedInDateRangeSearchResponse) minVersion() Version {
	return versionMustParse("2.2.0")
}
func (r ArtifactsCreatedInDateRangeSearchResponse) maxVersion() Version {
	return versionMustParse(CurrentVersion)
}
func (r ArtifactsCreatedInDateRangeSearchResponse) deprecated() bool { return false }

// PatternSearchResponse represents a PatternSearch response
type PatternSearchResponse struct {
	RepoURI       string   `json:"repoUri"`
	SourcePattern string   `json:"sourcePattern"`
	Files         []string `json:"files"`
}

// PatternSearchResponseTestData is test data for PatternSearchResponse
const PatternSearchResponseTestData = "pattern_search.json"

func (r PatternSearchResponse) minVersion() Version { return versionMustParse("2.3.4") }
func (r PatternSearchResponse) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (r PatternSearchResponse) deprecated() bool    { return false }

// BuildsForDependencyResponse represents a BuildsForDepdency response
type BuildsForDependencyResponse struct {
	Results []struct {
		URI string `json:"uri"`
	} `json:"results"`
}

// BuildsForDependencyResponseTestData is test data for BuildsForDepdencyResponse
const BuildsForDependencyResponseTestData = "build_for_dependency.json"

func (r BuildsForDependencyResponse) minVersion() Version { return versionMustParse("2.3.4") }
func (r BuildsForDependencyResponse) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (r BuildsForDependencyResponse) deprecated() bool    { return false }

// LicenseSearchResponse represents a LicenseSearch response
type LicenseSearchResponse struct {
	Results []struct {
		URI     string `json:"uri"`
		License string `json:"license"`
		Found   string `json:"found"`
		Status  string `json:"status"`
	} `json:"results"`
}

// LicenseSearchResponseTestData is test data for LicenseSearchResponse
const LicenseSearchResponseTestData = "license_search.json"

func (r LicenseSearchResponse) minVersion() Version { return versionMustParse("2.3.0") }
func (r LicenseSearchResponse) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (r LicenseSearchResponse) deprecated() bool    { return false }

// ArtifactVersionSearchResponse represents a ArtifactVersionSearch response
type ArtifactVersionSearchResponse struct {
	Results []struct {
		Version     string `json:"version"`
		Integration bool   `json:"integration"`
	} `json:"results"`
}

// ArtifactVersionSearchResponseTestData is test data for ArtifactVersionSearchResponse
const ArtifactVersionSearchResponseTestData = "artifact_version_search.json"

func (r ArtifactVersionSearchResponse) minVersion() Version { return versionMustParse("2.6.0") }
func (r ArtifactVersionSearchResponse) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (r ArtifactVersionSearchResponse) deprecated() bool    { return false }

// ArtifactLatestVersionBasedOnLayoutSearchResponse represents a ArtifactLatestVersionBasedOnLayout response
type ArtifactLatestVersionBasedOnLayoutSearchResponse struct{}

func (r ArtifactLatestVersionBasedOnLayoutSearchResponse) minVersion() Version {
	return versionMustParse("2.6.0")
}
func (r ArtifactLatestVersionBasedOnLayoutSearchResponse) maxVersion() Version {
	return versionMustParse(CurrentVersion)
}
func (r ArtifactLatestVersionBasedOnLayoutSearchResponse) deprecated() bool { return false }

// ArtifactLatestVersionBasedOnPropertiesSearchResponse represents a ArtifactLatestVersionBasedOnPropertiesSearch response
type ArtifactLatestVersionBasedOnPropertiesSearchResponse struct {
	Version   string `json:"version"`
	Artifacts []struct {
		URI string `json:"uri"`
	} `json:"artifacts"`
}

// ArtifactLatestVersionBasedOnPropertiesSearchResponseTestData is test data for ArtifactLatestVersionBasedOnPropertiesSearchResponse
const ArtifactLatestVersionBasedOnPropertiesSearchResponseTestData = "artifact_latest_version_by_props_search.json"

func (r ArtifactLatestVersionBasedOnPropertiesSearchResponse) minVersion() Version {
	return versionMustParse("3.1.1")
}
func (r ArtifactLatestVersionBasedOnPropertiesSearchResponse) maxVersion() Version {
	return versionMustParse(CurrentVersion)
}
func (r ArtifactLatestVersionBasedOnPropertiesSearchResponse) deprecated() bool { return false }

// BuildArtifactsSearchResponse represents a BuildArtifactsSearch response
type BuildArtifactsSearchResponse struct {
	Results []struct {
		DownloadURI string `json:"downloadUri"`
	} `json:"results"`
}

// BuildArtifactsSearchResponseTestData is test data for BuildArtifactsSearchResponse
const BuildArtifactsSearchResponseTestData = "build_artifacts_search.json"

func (r BuildArtifactsSearchResponse) minVersion() Version { return versionMustParse("2.6.5") }
func (r BuildArtifactsSearchResponse) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (r BuildArtifactsSearchResponse) deprecated() bool    { return false }

// ListDockerRepositoriesResponse represents a ListDockerRepositories response
type ListDockerRepositoriesResponse struct {
	Repositories []string `json:"repositories"`
}

// ListDockerRepositoriesResponseTestData is test data for ListDockerRepositoriesResponse
const ListDockerRepositoriesResponseTestData = "list_docker_repos.json"

func (r ListDockerRepositoriesResponse) minVersion() Version { return versionMustParse("4.4.3") }
func (r ListDockerRepositoriesResponse) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (r ListDockerRepositoriesResponse) deprecated() bool    { return false }

// ListDockerTagsResponse represents a ListDockerTags response
type ListDockerTagsResponse struct {
	Name string   `json:"name"`
	Tags []string `json:"tags"`
}

// ListDockerTagsResponseTestData is test data for ListDockerTagsResponse
const ListDockerTagsResponseTestData = "list_docker_tags.json"

func (r ListDockerTagsResponse) minVersion() Version { return versionMustParse("4.4.3") }
func (r ListDockerTagsResponse) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (r ListDockerTagsResponse) deprecated() bool    { return false }
