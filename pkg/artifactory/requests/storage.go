package requests

// OptimizeSystemStorageRequest represents a OptimizeSystemStorage request
type OptimizeSystemStorageRequest struct{}

func (r OptimizeSystemStorageRequest) minVersion() Version { return versionMustParse("4.6.0") }
func (r OptimizeSystemStorageRequest) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (r OptimizeSystemStorageRequest) deprecated() bool    { return false }
