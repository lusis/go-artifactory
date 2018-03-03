package artifactory

import (
	"bytes"
	"encoding/json"

	multierror "github.com/hashicorp/go-multierror"
	"github.com/lusis/go-artifactory/pkg/artifactory/requests"
	"github.com/lusis/go-artifactory/pkg/artifactory/responses"
	httpclient "github.com/lusis/go-artifactory/pkg/httpclient"
)

// RepositoryEntry is an entry in a Repositories list
type RepositoryEntry struct {
	responses.RepositoryItemEntry
}

// Repositories is a collection of repositories
type Repositories []RepositoryEntry

// RepositoryConfiguration is an individual repository
type RepositoryConfiguration interface {
	getType() string
}

type unknownRepository struct {
	RClass string `json:"rclass"`
}

// LocalRepository is a local repository
type LocalRepository responses.LocalRepositoryConfigurationResponse

func (r *LocalRepository) getType() string { return r.RClass }

// RemoteRepository is a remote repository
type RemoteRepository responses.RemoteRepositoryConfigurationResponse

func (r *RemoteRepository) getType() string { return r.RClass }

// VirtualRepository is a virtual repository
type VirtualRepository responses.VirtualRepositoryConfigurationResponse

func (r *VirtualRepository) getType() string { return r.RClass }

// GetRepositories lists all repositories
func (c *Client) GetRepositories() (Repositories, error) {
	if err := c.checkRequiredResponseVersion(responses.GetRepositoriesResponse{}); err != nil {
		return nil, err
	}
	r := Repositories{}
	res, err := c.httpGet("/repositories", requestJSON())
	if err != nil {
		return nil, err
	}
	if jsonErr := json.Unmarshal(res, &r); jsonErr != nil {
		return nil, &UnmarshalError{msg: multierror.Append(errDecoding, jsonErr).Error()}
	}
	return r, nil
}

// GetRepositoryConfiguration gets a repo's configuration
// Note that you will need to cast this to the appropriate type of repo type
// to operate on it
func (c *Client) GetRepositoryConfiguration(repo string) (RepositoryConfiguration, error) {
	if err := c.checkRequiredResponseVersion(responses.RepositoryConfigurationResponse{}); err != nil {
		return nil, err
	}
	repoTypes := map[string]RepositoryConfiguration{
		"local":   &LocalRepository{},
		"remote":  &RemoteRepository{},
		"virtual": &VirtualRepository{},
	}
	r := unknownRepository{}
	res, err := c.httpGet("repositories/"+repo, requestJSON())
	if err != nil {
		return nil, err
	}
	if jsonErr := json.Unmarshal(res, &r); jsonErr != nil {
		return nil, &UnmarshalError{msg: multierror.Append(errDecoding, jsonErr).Error()}
	}
	data := repoTypes[r.RClass]
	if jsonErr := json.Unmarshal(res, data); jsonErr != nil {
		return nil, &UnmarshalError{msg: multierror.Append(errDecoding, jsonErr).Error()}
	}

	return data, nil
}

// GetLocalRepositoryConfiguration gets a local repository configuration
func (c *Client) GetLocalRepositoryConfiguration(repo string) (*LocalRepository, error) {
	if err := c.checkRequiredResponseVersion(responses.LocalRepositoryConfigurationResponse{}); err != nil {
		return nil, err
	}
	i, err := c.GetRepositoryConfiguration(repo)
	if err != nil {
		return nil, err
	}
	return i.(*LocalRepository), nil
}

// GetRemoteRepositoryConfiguration gets a local repository configuration
func (c *Client) GetRemoteRepositoryConfiguration(repo string) (*RemoteRepository, error) {
	if err := c.checkRequiredResponseVersion(responses.RemoteRepositoryConfigurationResponse{}); err != nil {
		return nil, err
	}
	i, err := c.GetRepositoryConfiguration(repo)
	if err != nil {
		return nil, err
	}
	return i.(*RemoteRepository), nil
}

// GetVirtualRepositoryConfiguration gets a local repository configuration
func (c *Client) GetVirtualRepositoryConfiguration(repo string) (*VirtualRepository, error) {
	if err := c.checkRequiredResponseVersion(responses.VirtualRepositoryConfigurationResponse{}); err != nil {
		return nil, err
	}
	i, err := c.GetRepositoryConfiguration(repo)
	if err != nil {
		return nil, err
	}
	return i.(*VirtualRepository), nil
}

// CreateOrUpdateRepositoryOption is a functional option type for setting optional fields when creating a repository
type CreateOrUpdateRepositoryOption func(map[string]interface{}) error

// RepositoryIntOption sets a key's value to an int
func RepositoryIntOption(k string, v int) CreateOrUpdateRepositoryOption {
	return func(m map[string]interface{}) error {
		m[k] = v
		return nil
	}
}

// RepositoryStringOption sets a key's value to a string
func RepositoryStringOption(k, v string) CreateOrUpdateRepositoryOption {
	return func(m map[string]interface{}) error {
		m[k] = v
		return nil
	}
}

// RepositoryStringSliceOption sets a key's value to a slice of strings
func RepositoryStringSliceOption(k string, v []string) CreateOrUpdateRepositoryOption {
	return func(m map[string]interface{}) error {
		m[k] = v
		return nil
	}
}

// RepositoryBoolOption sets a key's value to a bool
func RepositoryBoolOption(k string, v bool) CreateOrUpdateRepositoryOption {
	return func(m map[string]interface{}) error {
		m[k] = v
		return nil
	}
}

// CreateLocalRepository creates a repository
func (c *Client) CreateLocalRepository(repoName string, opts ...CreateOrUpdateRepositoryOption) error {
	if err := c.checkRequiredRequestVersion(requests.CreateLocalRepositoryRequest{}); err != nil {
		return err
	}
	repo := map[string]interface{}{}

	for _, o := range opts {
		if err := o(repo); err != nil {
			return err
		}
	}
	repo["rclass"] = "local"
	data, err := json.Marshal(repo)
	if err != nil {
		return &MarshalError{msg: multierror.Append(errEncoding, err).Error()}
	}
	_, err = c.httpPut("repositories/"+repoName, contentType("application/json"), withBody(bytes.NewReader(data)))
	return err
}

// CreateRemoteRepository creates a repository
func (c *Client) CreateRemoteRepository(repoName, repoURL string, opts ...CreateOrUpdateRepositoryOption) error {
	if err := c.checkRequiredRequestVersion(requests.CreateRemoteRepositoryRequest{}); err != nil {
		return err
	}
	repo := map[string]interface{}{}

	for _, o := range opts {
		if err := o(repo); err != nil {
			return err
		}
	}
	repo["rclass"] = "remote"
	repo["url"] = repoURL
	data, err := json.Marshal(repo)
	if err != nil {
		return &MarshalError{msg: multierror.Append(errEncoding, err).Error()}
	}
	_, err = c.httpPut("repositories/"+repoName, contentType("application/json"), withBody(bytes.NewReader(data)))
	return err
}

// CreateVirtualRepository creates a repository
func (c *Client) CreateVirtualRepository(repoName, packageType string, opts ...CreateOrUpdateRepositoryOption) error {
	if err := c.checkRequiredRequestVersion(requests.CreateVirtualRepositoryRequest{}); err != nil {
		return err
	}
	repo := map[string]interface{}{}

	for _, o := range opts {
		if err := o(repo); err != nil {
			return err
		}
	}
	repo["rclass"] = "virtual"
	repo["packageType"] = packageType
	data, err := json.Marshal(repo)
	if err != nil {
		return &MarshalError{msg: multierror.Append(errEncoding, err).Error()}
	}
	_, err = c.httpPut("repositories/"+repoName, contentType("application/json"), withBody(bytes.NewReader(data)))
	return err
}

// UpdateRepositoryConfiguration updates a repo configuration
func (c *Client) UpdateRepositoryConfiguration(repoName string, opts ...CreateOrUpdateRepositoryOption) error {
	if err := c.checkRequiredRequestVersion(requests.UpdateRepositoryConfigurationRequest{}); err != nil {
		return err
	}
	repo := map[string]interface{}{}

	for _, o := range opts {
		if err := o(repo); err != nil {
			return err
		}
	}
	data, err := json.Marshal(repo)
	if err != nil {
		return &MarshalError{msg: multierror.Append(errEncoding, err).Error()}
	}
	_, err = c.httpPost("repositories/"+repoName, contentType("application/json"), withBody(bytes.NewReader(data)))
	return err
}

// DeleteRepository deletes a repo
func (c *Client) DeleteRepository(repo string) error {
	if err := c.checkRequiredResponseVersion(responses.DeleteRepositoryResponse{}); err != nil {
		return err
	}
	return c.httpDelete("repositories/" + repo)
}

// CalculateOption is a functional option type for passing options to various Calculate requests
type CalculateOption func(*[]httpclient.RequestOption) error

// CalculateGPGPassphrase sets the X-GPG-PASSPHRASE header
func CalculateGPGPassphrase(p string) CalculateOption {
	return func(opts *[]httpclient.RequestOption) error {
		*opts = append(*opts, addHeader("X-GPG-PASSPHRASE", p))
		return nil
	}
}

// CalculateQueryParams sets any custom query params required for a Calculate request
func CalculateQueryParams(m map[string]string) CalculateOption {
	return func(opts *[]httpclient.RequestOption) error {
		*opts = append(*opts, queryParams(m))
		return nil
	}
}

// CalculateYumRepositoryMetadata calculates a yum repo's metadata
func (c *Client) CalculateYumRepositoryMetadata(repoKey string, additionalOpts ...CalculateOption) error {
	if err := c.checkRequiredRequestVersion(requests.CalculateYumMetadataRequest{}); err != nil {
		return err
	}
	opts := []httpclient.RequestOption{}
	for _, o := range additionalOpts {
		if err := o(&opts); err != nil {
			return err
		}
	}
	opts = append(opts)
	_, err := c.httpPost("yum/"+repoKey, opts...)
	return err
}

// CalculateNuGetRepositoryMetadata calculates a nuget repo's metadata
func (c *Client) CalculateNuGetRepositoryMetadata(repoKey string) error {
	if err := c.checkRequiredRequestVersion(requests.CalculateNuGetMetadataRequest{}); err != nil {
		return err
	}
	_, err := c.httpPost("nuget/" + repoKey + "/reindex")
	return err
}

// CalculateNPMRepositoryMetadata calculates an npm repo's metadata
func (c *Client) CalculateNPMRepositoryMetadata(repoKey string) error {
	if err := c.checkRequiredRequestVersion(requests.CalculateNPMMetadataRequest{}); err != nil {
		return err
	}
	_, err := c.httpPost("npm/" + repoKey + "/reindex")
	return err
}

// CalculateMavenIndex calculates the index for a maven repository
func (c *Client) CalculateMavenIndex(additionalOpts ...CalculateOption) error {
	if err := c.checkRequiredRequestVersion(requests.CalculateMavenIndexRequest{}); err != nil {
		return err
	}
	opts := []httpclient.RequestOption{}
	for _, o := range additionalOpts {
		if err := o(&opts); err != nil {
			return err
		}
	}
	opts = append(opts)
	_, err := c.httpPost("maven", opts...)
	return err
}

// CalculateMavenMetadata calculates the maven metadata for a repository
func (c *Client) CalculateMavenMetadata(path string, additionalOpts ...CalculateOption) error {
	if err := c.checkRequiredRequestVersion(requests.CalculateMavenMetadataRequest{}); err != nil {
		return err
	}
	opts := []httpclient.RequestOption{}
	for _, o := range additionalOpts {
		if err := o(&opts); err != nil {
			return err
		}
	}
	opts = append(opts)
	_, err := c.httpPost("maven/calculateMetadata/"+path, opts...)
	return err
}

// CalculateDebianRepositoryMetadata calculates the metadata for a debian repository
func (c *Client) CalculateDebianRepositoryMetadata(repoKey string, additionalOpts ...CalculateOption) error {
	if err := c.checkRequiredRequestVersion(requests.CalculateDebianMetadataRequest{}); err != nil {
		return err
	}
	opts := []httpclient.RequestOption{}
	for _, o := range additionalOpts {
		if err := o(&opts); err != nil {
			return err
		}
	}
	opts = append(opts)
	_, err := c.httpPost("deb/reindex/"+repoKey, opts...)
	return err
}

// CalculateOpkgRepostitoryMetadata calculates the metadata for an opkg repository
func (c *Client) CalculateOpkgRepostitoryMetadata(repoKey string, additionalOpts ...CalculateOption) error {
	if err := c.checkRequiredRequestVersion(requests.CalculateOpkgMetadataRequest{}); err != nil {
		return err
	}
	opts := []httpclient.RequestOption{}
	for _, o := range additionalOpts {
		if err := o(&opts); err != nil {
			return err
		}
	}
	opts = append(opts)
	_, err := c.httpPost("opkg/reindex/"+repoKey, opts...)
	return err
}

// CalculateBowerIndex calculates the index for a bower repository
func (c *Client) CalculateBowerIndex(repoKey string) error {
	if err := c.checkRequiredRequestVersion(requests.CalculateBowerIndexRequest{}); err != nil {
		return err
	}
	_, err := c.httpPost("bower/" + repoKey + "/reindex")
	return err
}

// CalculateHelmChartIndex calculates the index for a helm chart repository
func (c *Client) CalculateHelmChartIndex(repoKey string) error {
	if err := c.checkRequiredRequestVersion(requests.CalculateHelmChartIndexRequest{}); err != nil {
		return err
	}
	_, err := c.httpPost("helm/" + repoKey + "/reindex")
	return err
}
