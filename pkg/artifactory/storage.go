package artifactory

import "github.com/lusis/go-artifactory/pkg/artifactory/requests"

// OptimizeSystemStorage optimizes system storage
func (c *Client) OptimizeSystemStorage() error {
	if err := c.checkRequiredRequestVersion(requests.OptimizeSystemStorageRequest{}); err != nil {
		return err
	}
	_, err := c.httpPost("/api/system/storage/optimize", requestExpects(200))
	return err
}
