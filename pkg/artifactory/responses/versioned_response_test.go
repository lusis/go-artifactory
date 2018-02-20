package responses

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type testVersionResponse struct{}

func (u testVersionResponse) minVersion() Version { return versionMustParse(AbsoluteMinimumVersion) }
func (u testVersionResponse) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (u testVersionResponse) deprecated() bool    { return false }

func TestVersionedResponseMinVersion(t *testing.T) {
	x := testVersionResponse{}
	smallVer := versionMustParse("1.1.1")
	reqVer := GetMinVersionFor(x)

	require.False(t, reqVer.LessThan(smallVer.Version))
}

func TestVersionedResponseMaxVersion(t *testing.T) {
	x := testVersionResponse{}
	bigVer := versionMustParse("6.0.0")
	reqVer := GetMinVersionFor(x)

	require.False(t, reqVer.GreaterThan(bigVer.Version))
}

func TestVersionedDeprecated(t *testing.T) {
	x := testVersionResponse{}

	require.False(t, IsDeprecated(x))
}
