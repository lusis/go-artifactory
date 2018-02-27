package requests

// AQLSearchRequest represents a AQLSearch request
type AQLSearchRequest struct{}

func (r AQLSearchRequest) minVersion() Version { return versionMustParse("3.5.0") }
func (r AQLSearchRequest) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (r AQLSearchRequest) deprecated() bool    { return false }
