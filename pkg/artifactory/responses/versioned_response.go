package responses

import (
	gover "github.com/hashicorp/go-version"
)

// Version is a self-contained go-version Version
type Version struct {
	*gover.Version
}

// VersionedResponse is an interface for a Rundeck Response that supports versioning information
type VersionedResponse interface {
	minVersion() Version
	maxVersion() Version
	deprecated() bool
}

// AbsoluteMinimumVersion is the absolute minimum version this library will support
const AbsoluteMinimumVersion = "5.1"

// CurrentVersion is the current version of the API that this library is tested against
const CurrentVersion = "5.5"

// GetMinVersionFor gets the minimum api version required for a response
func GetMinVersionFor(a VersionedResponse) Version { return a.minVersion() }

// GetMaxVersionFor gets the maximum api version required for a response
func GetMaxVersionFor(a VersionedResponse) Version { return a.maxVersion() }

// IsDeprecated indicates if a response is deprecated or not
func IsDeprecated(a VersionedResponse) bool { return a.deprecated() }

// GenericVersionedResponse is for version checking
// Some operations don't have a response (think DELETE or PUT)
// but we still want to do a version check on ALL functions anyway
// This response simply responds to that
type GenericVersionedResponse struct{}

func (g GenericVersionedResponse) minVersion() Version {
	return versionMustParse(AbsoluteMinimumVersion)
}
func (g GenericVersionedResponse) maxVersion() Version {
	return versionMustParse(CurrentVersion)
}
func (g GenericVersionedResponse) deprecated() bool { return false }

// versionMustParse is a panicing version of NewVersion
func versionMustParse(v string) Version {
	ver, err := gover.NewVersion(v)
	if err != nil {
		panic("cannot parse version")
	}
	return Version{ver}
}
