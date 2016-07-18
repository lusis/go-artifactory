package artifactory

func (c *ArtifactoryClient) GetSystemHealthPing() (s string, e error) {
	d, e := c.Get("/api/system/ping", make(map[string]string))
	return string(d), e
}
