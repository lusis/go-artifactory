package responses

// CreateBundleResponse represents a CreateBundle response
type CreateBundleResponse struct {
	Bundles []string `json:"bundles"`
}

// CreateBundleResponseTestData is test data for CreateBundleResponse
const CreateBundleResponseTestData = "create_bundle.json"

func (r CreateBundleResponse) minVersion() Version { return versionMustParse("4.3.0") }
func (r CreateBundleResponse) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (r CreateBundleResponse) deprecated() bool    { return false }

// ListBundlesResponse represents a ListBundles response
type ListBundlesResponse struct {
	Bundles []string `json:"bundles"`
}

// ListBundlesResponseTestData is test data for ListBundlesResponse
const ListBundlesResponseTestData = "list_bundles.json"

func (r ListBundlesResponse) minVersion() Version { return versionMustParse("4.3.0") }
func (r ListBundlesResponse) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (r ListBundlesResponse) deprecated() bool    { return false }

// GetBundleResponse represents a GetBundle response
type GetBundleResponse struct{}

func (r GetBundleResponse) minVersion() Version { return versionMustParse("4.3.0") }
func (r GetBundleResponse) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (r GetBundleResponse) deprecated() bool    { return false }
