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
