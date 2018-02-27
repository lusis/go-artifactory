package artifactory

import (
	"encoding/json"
	"strings"

	multierror "github.com/hashicorp/go-multierror"
	"github.com/lusis/go-artifactory/pkg/artifactory/requests"
	"github.com/lusis/go-artifactory/pkg/artifactory/responses"
)

// AQLResults represents AQL search results
type AQLResults struct {
	responses.AQLSearchResponse
}

// AQLSearch performs an AQL search
func (c *Client) AQLSearch(query string) (*AQLResults, error) {
	if err := c.checkRequiredRequestVersion(requests.AQLSearchRequest{}); err != nil {
		return nil, err
	}
	body := strings.NewReader(query)
	res, err := c.httpPost("/api/search/aql", requestExpects(200), accept("application/json"), withBody(body))
	if err != nil {
		return nil, err
	}
	searchResults := &AQLResults{}
	if jsonErr := json.Unmarshal(res, searchResults); jsonErr != nil {
		return nil, &UnmarshalError{msg: multierror.Append(errDecoding, jsonErr).Error()}
	}
	return searchResults, nil
}
