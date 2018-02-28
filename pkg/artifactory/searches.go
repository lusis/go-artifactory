package artifactory

import (
	"bytes"
	"encoding/json"
	"strings"

	multierror "github.com/hashicorp/go-multierror"
	"github.com/lusis/go-artifactory/pkg/artifactory/requests"
	"github.com/lusis/go-artifactory/pkg/artifactory/responses"
)

// AQLSearchResults represents AQL search results
type AQLSearchResults struct {
	responses.AQLSearchResponse
}

// ArtifactSearchResults represents a quick artifact search result
type ArtifactSearchResults struct {
	responses.ArtifactSearchResponse
}

// GAVCSearchResults represents GAVC search results
type GAVCSearchResults struct {
	responses.GAVCSearchResponse
}

// ArchiveEntriesSearchResults represents Archive entry search results
type ArchiveEntriesSearchResults struct {
	responses.ArchiveEntriesSearchResponse
}

// PropertySearchResults represents property-based search results
type PropertySearchResults struct {
	responses.PropertySearchResponse
}

// ChecksumSearchResults represents results of a checksum search
type ChecksumSearchResults struct {
	responses.ChecksumSearchResponse
}

// BadChecksumSearchResults represents results of a checksum search
type BadChecksumSearchResults struct {
	responses.BadChecksumSearchResponse
}

// ArtifactsNotDownloadedSinceSearchResults represents results of a checksum search
type ArtifactsNotDownloadedSinceSearchResults struct {
	responses.ArtifactsNotDownloadedSinceSearchResponse
}

// ArtifactsWithDateInRangeSearchResults represents results of a date-range artifact search
type ArtifactsWithDateInRangeSearchResults struct {
	responses.ArtifactsWithDataInRangeSearchResponse
}

// ArtifactsCreatedInDateRangeSearchResults represents results of a date-range artifact search based on creation date
type ArtifactsCreatedInDateRangeSearchResults struct {
	responses.ArtifactsCreatedInDateRangeSearchResponse
}

// PatternSearchResults represents results from a pattern search
type PatternSearchResults struct {
	responses.PatternSearchResponse
}

// BuildsForDependencySearchResults represents results from a search for builds based on provided dependency
type BuildsForDependencySearchResults struct {
	responses.BuildsForDependencyResponse
}

// LicenseSearchResults represents results from a search base on license
type LicenseSearchResults struct {
	responses.LicenseSearchResponse
}

// ArtifactVersionSearchResults represents results from a search based on version
type ArtifactVersionSearchResults struct {
	responses.ArtifactVersionSearchResponse
}

// ArtifactLatestBasedOnLayoutSearchResults represents the result of a search for an artifact based on layout
type ArtifactLatestBasedOnLayoutSearchResults struct {
	responses.ArtifactLatestVersionBasedOnLayoutSearchResponse
}

// ArtifactLatestBasedOnPropertiesSearchResults represents the result of a search for an artifact based on properties
type ArtifactLatestBasedOnPropertiesSearchResults struct {
	responses.ArtifactLatestVersionBasedOnPropertiesSearchResponse
}

// BuildArtifactsSearchResults represents the result of a search for build artifacts
type BuildArtifactsSearchResults struct {
	responses.BuildArtifactsSearchResponse
}

// BuildArtifactsSearchOption is a functional option for build artifacts searches
type BuildArtifactsSearchOption func(*requests.BuildArtifactsSearchRequest) error

// ListDockerRepositoriesResults represents the result of listing docker repositories
type ListDockerRepositoriesResults struct {
	responses.ListDockerRepositoriesResponse
}

// ListDockerTagsSearchResults represents the results of listing docker tags
type ListDockerTagsSearchResults struct {
	responses.ListDockerTagsResponse
}

// AQLSearch performs an AQL search
func (c *Client) AQLSearch(query string) (*AQLSearchResults, error) {
	if err := c.checkRequiredRequestVersion(requests.AQLSearchRequest{}); err != nil {
		return nil, err
	}
	body := strings.NewReader(query)
	res, err := c.httpPost("search/aql", requestExpects(200), accept("application/json"), withBody(body), contentType("text/plain"))
	if err != nil {
		return nil, err
	}
	searchResults := &AQLSearchResults{}
	if jsonErr := json.Unmarshal(res, searchResults); jsonErr != nil {
		return nil, &UnmarshalError{msg: multierror.Append(errDecoding, jsonErr).Error()}
	}
	return searchResults, nil
}

// ArtifactSearch performs an AQL search
// p is a map[string]string matching the format of
// the options specified in the API docs
// https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-ArtifactSearch(QuickSearch)
func (c *Client) ArtifactSearch(p map[string]string) (*ArtifactSearchResults, error) {
	if err := c.checkRequiredResponseVersion(responses.ArtifactSearchResponse{}); err != nil {
		return nil, err
	}
	res, err := c.httpGet("search/artifact", requestExpects(200), requestJSON(), queryParams(p))
	if err != nil {
		return nil, err
	}
	searchResults := &ArtifactSearchResults{}
	if jsonErr := json.Unmarshal(res, searchResults); jsonErr != nil {
		return nil, &UnmarshalError{msg: multierror.Append(errDecoding, jsonErr).Error()}
	}
	return searchResults, nil
}

// ArchiveEntriesSearch performs a class search
// p is a map[string]string matching the format of
// the options specified in the api docs
// https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-ArchiveEntriesSearch(ClassSearch)
func (c *Client) ArchiveEntriesSearch(p map[string]string) (*ArchiveEntriesSearchResults, error) {
	if err := c.checkRequiredResponseVersion(responses.ArchiveEntriesSearchResponse{}); err != nil {
		return nil, err
	}
	res, err := c.httpGet("search/archive", requestExpects(200), requestJSON(), queryParams(p))
	if err != nil {
		return nil, err
	}
	searchResults := &ArchiveEntriesSearchResults{}
	if jsonErr := json.Unmarshal(res, searchResults); jsonErr != nil {
		return nil, &UnmarshalError{msg: multierror.Append(errDecoding, jsonErr).Error()}
	}
	return searchResults, nil
}

// GAVCSearch performs a GAVC search
// p is a map[string]string matching the format of
// options specificied in the API docs
// https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-GAVCSearch
// p will be passed as is
func (c *Client) GAVCSearch(p map[string]string) (*GAVCSearchResults, error) {
	if err := c.checkRequiredResponseVersion(responses.GAVCSearchResponse{}); err != nil {
		return nil, err
	}
	res, err := c.httpGet("search/gavc", requestExpects(200), requestJSON(), queryParams(p))
	if err != nil {
		return nil, err
	}
	searchResults := &GAVCSearchResults{}
	if jsonErr := json.Unmarshal(res, searchResults); jsonErr != nil {
		return nil, &UnmarshalError{msg: multierror.Append(errDecoding, jsonErr).Error()}
	}
	return searchResults, nil
}

// PropertySearch performs a property search
// p is a map[string]string matching the format of
// options specificied in the API docs
// https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-PropertySearch
// p will be passed as is
func (c *Client) PropertySearch(p map[string]string) (*PropertySearchResults, error) {
	if err := c.checkRequiredResponseVersion(responses.PropertySearchResponse{}); err != nil {
		return nil, err
	}
	res, err := c.httpGet("search/prop", requestExpects(200), requestJSON(), queryParams(p))
	if err != nil {
		return nil, err
	}
	searchResults := &PropertySearchResults{}
	if jsonErr := json.Unmarshal(res, searchResults); jsonErr != nil {
		return nil, &UnmarshalError{msg: multierror.Append(errDecoding, jsonErr).Error()}
	}
	return searchResults, nil
}

// ChecksumSearch performs a property search
// p is a map[string]string matching the format of
// options specificied in the API docs
// https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-ChecksumSearch
// p will be passed as is
func (c *Client) ChecksumSearch(p map[string]string) (*ChecksumSearchResults, error) {
	if err := c.checkRequiredResponseVersion(responses.ChecksumSearchResponse{}); err != nil {
		return nil, err
	}
	res, err := c.httpGet("search/checksum", requestExpects(200), requestJSON(), queryParams(p))
	if err != nil {
		return nil, err
	}
	searchResults := &ChecksumSearchResults{}
	if jsonErr := json.Unmarshal(res, searchResults); jsonErr != nil {
		return nil, &UnmarshalError{msg: multierror.Append(errDecoding, jsonErr).Error()}
	}
	return searchResults, nil
}

// BadChecksumSearch performs a property search
// p is a map[string]string matching the format of
// options specificied in the API docs
// https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-ChecksumSearch
// p will be passed as is
func (c *Client) BadChecksumSearch(p map[string]string) (*BadChecksumSearchResults, error) {
	if err := c.checkRequiredResponseVersion(responses.BadChecksumSearchResponse{}); err != nil {
		return nil, err
	}
	res, err := c.httpGet("search/badChecksum", requestExpects(200), requestJSON(), queryParams(p))
	if err != nil {
		return nil, err
	}
	searchResults := &BadChecksumSearchResults{}
	if jsonErr := json.Unmarshal(res, searchResults); jsonErr != nil {
		return nil, &UnmarshalError{msg: multierror.Append(errDecoding, jsonErr).Error()}
	}
	return searchResults, nil
}

// ArtifactsNotDownloadedSinceSearch performs a search for artifacts not downloaded since given parameters
// p is a map[string]string matching the format of
// options specificied in the API docs
// https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-ArtifactsNotDownloadedSince
// p will be passed as is
func (c *Client) ArtifactsNotDownloadedSinceSearch(p map[string]string) (*ArtifactsNotDownloadedSinceSearchResults, error) {
	if err := c.checkRequiredResponseVersion(responses.ArtifactsNotDownloadedSinceSearchResponse{}); err != nil {
		return nil, err
	}
	res, err := c.httpGet("search/usage", requestExpects(200), requestJSON(), queryParams(p))
	if err != nil {
		return nil, err
	}
	searchResults := &ArtifactsNotDownloadedSinceSearchResults{}
	if jsonErr := json.Unmarshal(res, searchResults); jsonErr != nil {
		return nil, &UnmarshalError{msg: multierror.Append(errDecoding, jsonErr).Error()}
	}
	return searchResults, nil
}

// ArtifactsWithDateInRangeSearch performs a search for artifacts in a given date range
// p is a map[string]string matching the format of
// options specificied in the API docs
// https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-ArtifactsWithDateInDateRange
// p will be passed as is
func (c *Client) ArtifactsWithDateInRangeSearch(p map[string]string) (*ArtifactsWithDateInRangeSearchResults, error) {
	if err := c.checkRequiredResponseVersion(responses.ArtifactsWithDataInRangeSearchResponse{}); err != nil {
		return nil, err
	}
	res, err := c.httpGet("search/dates", requestExpects(200), requestJSON(), queryParams(p))
	if err != nil {
		return nil, err
	}
	searchResults := &ArtifactsWithDateInRangeSearchResults{}
	if jsonErr := json.Unmarshal(res, searchResults); jsonErr != nil {
		return nil, &UnmarshalError{msg: multierror.Append(errDecoding, jsonErr).Error()}
	}
	return searchResults, nil
}

// ArtifactsCreatedInDateRangeSearch performs a search for artifacts created in a given date range
// p is a map[string]string matching the format of
// options specificied in the API docs
// https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-ArtifactsCreatedInDateRange
// p will be passed as is
func (c *Client) ArtifactsCreatedInDateRangeSearch(p map[string]string) (*ArtifactsCreatedInDateRangeSearchResults, error) {
	if err := c.checkRequiredResponseVersion(responses.ArtifactsCreatedInDateRangeSearchResponse{}); err != nil {
		return nil, err
	}
	res, err := c.httpGet("search/created", requestExpects(200), requestJSON(), queryParams(p))
	if err != nil {
		return nil, err
	}
	searchResults := &ArtifactsCreatedInDateRangeSearchResults{}
	if jsonErr := json.Unmarshal(res, searchResults); jsonErr != nil {
		return nil, &UnmarshalError{msg: multierror.Append(errDecoding, jsonErr).Error()}
	}
	return searchResults, nil
}

// PatternSearch performs a search for artifacts based on a pattern
// https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-PatternSearch
// p will be passed as is
func (c *Client) PatternSearch(p string) (*PatternSearchResults, error) {
	if err := c.checkRequiredResponseVersion(responses.PatternSearchResponse{}); err != nil {
		return nil, err
	}
	qp := map[string]string{
		"pattern": p,
	}
	res, err := c.httpGet("search/pattern", requestExpects(200), requestJSON(), queryParams(qp))
	if err != nil {
		return nil, err
	}
	searchResults := &PatternSearchResults{}
	if jsonErr := json.Unmarshal(res, searchResults); jsonErr != nil {
		return nil, &UnmarshalError{msg: multierror.Append(errDecoding, jsonErr).Error()}
	}
	return searchResults, nil
}

// BuildsForDependencySearch finds all the builds an artifact is a dependency of
// https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-BuildsforDependency
// p is a map[string]string matching the format of
// options specificied in the API docs
func (c *Client) BuildsForDependencySearch(p map[string]string) (*BuildsForDependencySearchResults, error) {
	if err := c.checkRequiredResponseVersion(responses.BuildsForDependencyResponse{}); err != nil {
		return nil, err
	}

	res, err := c.httpGet("search/dependency", requestExpects(200), requestJSON(), queryParams(p))
	if err != nil {
		return nil, err
	}
	searchResults := &BuildsForDependencySearchResults{}
	if jsonErr := json.Unmarshal(res, searchResults); jsonErr != nil {
		return nil, &UnmarshalError{msg: multierror.Append(errDecoding, jsonErr).Error()}
	}
	return searchResults, nil
}

// LicenseSearch search for artifacts that were already tagged with license information
// https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-LicenseSearch
// p is a map[string]string matching the format of
// options specificied in the API docs
func (c *Client) LicenseSearch(p map[string]string) (*LicenseSearchResults, error) {
	if err := c.checkRequiredResponseVersion(responses.LicenseSearchResponse{}); err != nil {
		return nil, err
	}

	res, err := c.httpGet("search/license", requestExpects(200), requestJSON(), queryParams(p))
	if err != nil {
		return nil, err
	}
	searchResults := &LicenseSearchResults{}
	if jsonErr := json.Unmarshal(res, searchResults); jsonErr != nil {
		return nil, &UnmarshalError{msg: multierror.Append(errDecoding, jsonErr).Error()}
	}
	return searchResults, nil
}

// ArtifactVersionSearch search for artifacts that were already tagged with license information
// https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-ArtifactVersionSearch
// p is a map[string]string matching the format of
// options specificied in the API docs
func (c *Client) ArtifactVersionSearch(p map[string]string) (*ArtifactVersionSearchResults, error) {
	if err := c.checkRequiredResponseVersion(responses.ArtifactVersionSearchResponse{}); err != nil {
		return nil, err
	}

	res, err := c.httpGet("search/versions", requestExpects(200), requestJSON(), queryParams(p))
	if err != nil {
		return nil, err
	}
	searchResults := &ArtifactVersionSearchResults{}
	if jsonErr := json.Unmarshal(res, searchResults); jsonErr != nil {
		return nil, &UnmarshalError{msg: multierror.Append(errDecoding, jsonErr).Error()}
	}
	return searchResults, nil
}

// ArtifactLatestVersionBasedOnLayoutSearch search for artifacts that were already tagged with license information
// https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-ArtifactLatestVersionSearchBasedonLayout
// p is a map[string]string matching the format of
// options specificied in the API docs
func (c *Client) ArtifactLatestVersionBasedOnLayoutSearch(p map[string]string) (string, error) {
	if err := c.checkRequiredResponseVersion(responses.ArtifactLatestVersionBasedOnLayoutSearchResponse{}); err != nil {
		return "", err
	}

	res, err := c.httpGet("search/versions", requestExpects(200), requestJSON(), queryParams(p))
	if err != nil {
		return "", err
	}
	return string(res), nil
}

// ArtifactLatestBasedOnPropertiesSearch search for artifacts that were already tagged with license information
// https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-ArtifactLatestVersionSearchBasedOnProperties
// path is the path pattern specified in the api docs
// p is a map[string]string matching the format of options specificied in the API docs
func (c *Client) ArtifactLatestBasedOnPropertiesSearch(path string, p map[string]string) (*ArtifactLatestBasedOnPropertiesSearchResults, error) {
	if err := c.checkRequiredResponseVersion(responses.ArtifactLatestVersionBasedOnPropertiesSearchResponse{}); err != nil {
		return nil, err
	}

	res, err := c.httpGet("search/versions/"+path, requestExpects(200), requestJSON(), queryParams(p))
	if err != nil {
		return nil, err
	}
	searchResults := &ArtifactLatestBasedOnPropertiesSearchResults{}
	if jsonErr := json.Unmarshal(res, searchResults); jsonErr != nil {
		return nil, &UnmarshalError{msg: multierror.Append(errDecoding, jsonErr).Error()}
	}
	return searchResults, nil
}

// BuildArtifactsSearchStatus is the option for build status in build artifact searches
func BuildArtifactsSearchStatus(status string) BuildArtifactsSearchOption {
	return func(b *requests.BuildArtifactsSearchRequest) error {
		b.BuildStatus = status
		return nil
	}
}

// BuildArtifactsSearchRepos is the option for repos in build artifact searches
func BuildArtifactsSearchRepos(repos ...string) BuildArtifactsSearchOption {
	return func(b *requests.BuildArtifactsSearchRequest) error {
		b.Repos = repos
		return nil
	}
}

// BuildArtifactsSearchMappings is the option for mappings in build artifact searches
func BuildArtifactsSearchMappings(p []map[string]string) BuildArtifactsSearchOption {
	return func(b *requests.BuildArtifactsSearchRequest) error {
		b.Mappings = p
		return nil
	}
}

// BuildArtifactsSearch search for artifacts base on a specific build
// https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-BuildArtifactsSearch
// path is the path pattern specified in the api docs
// p is a map[string]string matching the format of options specificied in the API docs
func (c *Client) BuildArtifactsSearch(buildname, buildnumber string, opts ...BuildArtifactsSearchOption) (*BuildArtifactsSearchResults, error) {
	if err := c.checkRequiredResponseVersion(responses.BuildArtifactsSearchResponse{}); err != nil {
		return nil, err
	}

	b := &requests.BuildArtifactsSearchRequest{}
	for _, o := range opts {
		if err := o(b); err != nil {
			return nil, err
		}
	}
	data, err := json.Marshal(b)
	if err != nil {
		return nil, &MarshalError{msg: multierror.Append(errEncoding, err).Error()}
	}
	res, err := c.httpPost("search/buildArtifacts", requestExpects(200), requestJSON(), withBody(bytes.NewReader(data)))
	if err != nil {
		return nil, err
	}
	searchResults := &BuildArtifactsSearchResults{}
	if jsonErr := json.Unmarshal(res, searchResults); jsonErr != nil {
		return nil, &UnmarshalError{msg: multierror.Append(errDecoding, jsonErr).Error()}
	}
	return searchResults, nil
}

// ListDockerRepositories lists a docker registry's catalog
// https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-ListDockerRepositories
// r is the repository to list
func (c *Client) ListDockerRepositories(r string) (*ListDockerRepositoriesResults, error) {
	if err := c.checkRequiredResponseVersion(responses.ListDockerRepositoriesResponse{}); err != nil {
		return nil, err
	}
	searchResults := &ListDockerRepositoriesResults{}
	res, err := c.httpGet("api/docker/"+r+"/v2/_catalog", requestExpects(200), requestJSON())
	if err != nil {
		return nil, err
	}
	if jsonErr := json.Unmarshal(res, searchResults); jsonErr != nil {
		return nil, &UnmarshalError{msg: multierror.Append(errDecoding, jsonErr).Error()}
	}
	return searchResults, nil
}

// ListDockerTags lists a docker registry's catalog
// https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-ListDockerTags
// r is the repository to search
// i is the image name
func (c *Client) ListDockerTags(r, i string) (*ListDockerTagsSearchResults, error) {
	if err := c.checkRequiredResponseVersion(responses.ListDockerTagsResponse{}); err != nil {
		return nil, err
	}
	searchResults := &ListDockerTagsSearchResults{}
	res, err := c.httpGet("api/docker/"+r+"/v2/"+i+"/tags/list", requestExpects(200), requestJSON())
	if err != nil {
		return nil, err
	}
	if jsonErr := json.Unmarshal(res, searchResults); jsonErr != nil {
		return nil, &UnmarshalError{msg: multierror.Append(errDecoding, jsonErr).Error()}
	}
	return searchResults, nil
}
