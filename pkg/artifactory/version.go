package artifactory

import (
	"fmt"

	gover "github.com/hashicorp/go-version"
	"github.com/lusis/go-artifactory/pkg/artifactory/requests"
	"github.com/lusis/go-artifactory/pkg/artifactory/responses"
)

// Version is a wrapper around go-version Version
type Version struct {
	*gover.Version
}

func versionMustParse(v string) Version {
	ver, err := gover.NewVersion(v)
	if err != nil {
		panic("cannot parse version")
	}
	return Version{ver}
}

func (c *Client) checkRequiredResponseVersion(r responses.VersionedResponse) error {
	min := responses.GetMinVersionFor(r)
	max := responses.GetMaxVersionFor(r)

	if c.Config.APIVersion.Equal(max.Version) || c.Config.APIVersion.Equal(min.Version) {
		return nil
	}
	if c.Config.APIVersion.GreaterThan(min.Version) && c.Config.APIVersion.LessThan(max.Version) {
		return nil
	}
	return fmt.Errorf("Requested API version (%s) does not meet the requirements for this api call (min: %s, max: %s)",
		c.Config.APIVersion.String(), min.String(), max.String())
}

func (c *Client) checkRequiredRequestVersion(r requests.VersionedRequest) error {
	min := requests.GetMinVersionFor(r)
	max := requests.GetMaxVersionFor(r)

	if c.Config.APIVersion.Equal(max.Version) || c.Config.APIVersion.Equal(min.Version) {
		return nil
	}
	if c.Config.APIVersion.GreaterThan(min.Version) && c.Config.APIVersion.LessThan(max.Version) {
		return nil
	}
	return fmt.Errorf("Requested API version (%s) does not meet the requirements for this api call (min: %s, max: %s)",
		c.Config.APIVersion.String(), min.String(), max.String())
}
