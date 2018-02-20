package requests

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type testVersionRequest struct{}

func (u testVersionRequest) minVersion() Version { return versionMustParse(AbsoluteMinimumVersion) }
func (u testVersionRequest) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (u testVersionRequest) deprecated() bool    { return false }

func TestVersionedRequestMinVersion(t *testing.T) {
	x := testVersionRequest{}
	smallVer := versionMustParse("1.1.1")
	reqVer := GetMinVersionFor(x)

	require.False(t, reqVer.LessThan(smallVer.Version))
}

func TestVersionedRequestMaxVersion(t *testing.T) {
	x := testVersionRequest{}
	bigVer := versionMustParse("6.0.0")
	reqVer := GetMaxVersionFor(x)

	require.False(t, reqVer.GreaterThan(bigVer.Version))
}

func TestVersionedDeprecated(t *testing.T) {
	x := testVersionRequest{}

	require.False(t, IsDeprecated(x))
}
