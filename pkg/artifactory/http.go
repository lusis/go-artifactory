package artifactory

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"

	multierror "github.com/hashicorp/go-multierror"
	"github.com/lusis/go-artifactory/pkg/artifactory/responses"
	httpclient "github.com/lusis/go-artifactory/pkg/httpclient"
)

// withBody sets the post/put body for a request
// nolint: deadcode
func withBody(body io.Reader) httpclient.RequestOption {
	return httpclient.WithBody(body)
}

// queryParams sets the query params for a request
// nolint: deadcode
func queryParams(m map[string]string) httpclient.RequestOption {
	return httpclient.QueryParams(m)
}

// requestJSON sets a request to accept and respond with json
// nolint: deadcode
func requestJSON() httpclient.RequestOption {
	return httpclient.JSON()
}

// contentType allows setting the content-type for the request
// nolint: deadcode
func contentType(ct string) httpclient.RequestOption {
	return httpclient.ContentType(ct)
}

// accept allows setting the Accept header individually
// nolint: deadcode
func accept(ct string) httpclient.RequestOption {
	return httpclient.Accept(ct)
}

// requestExpects sets the expected status codes for a request
func requestExpects(code int) httpclient.RequestOption {
	return httpclient.ExpectStatus(code)
}

func (rc *Client) makeAPIPath(path string) string {
	return rc.Config.BaseURL + "/api/" + path
}

// Get performs an http get
func (rc *Client) Get(path string, opts ...httpclient.RequestOption) ([]byte, error) {
	return rc.httpGet(path, opts...)
}

func parseErrorBodytoError(b []byte) error {
	e := &responses.ErrorResponse{}
	compoundErr := errors.New("Artifactory error")
	err := json.Unmarshal(b, e)
	if err != nil {
		return multierror.Append(err, fmt.Errorf("%s", b))
	}
	for _, res := range e.Errors {
		if err := multierror.Append(compoundErr, fmt.Errorf("%d: %s", res.Status, res.Message)); err != nil {
			return err
		}
	}
	return compoundErr
}
func (rc *Client) httpGet(path string, opts ...httpclient.RequestOption) ([]byte, error) {
	authOpt, authErr := rc.authWrap()
	if authErr != nil {
		return nil, authErr
	}
	authOpt = append(authOpt, opts...)
	resp, err := httpclient.Get(rc.makeAPIPath(path), authOpt...)
	if err != nil {
		if resp.Body != nil {
			return nil, parseErrorBodytoError(resp.Body)
		}
		return nil, err
	}
	return resp.Body, nil
}

func (rc *Client) httpPost(path string, opts ...httpclient.RequestOption) ([]byte, error) {
	authOpt, authErr := rc.authWrap()
	if authErr != nil {
		return nil, authErr
	}
	opts = append(opts, authOpt...)
	resp, err := httpclient.Post(rc.makeAPIPath(path), opts...)
	if err != nil {
		if resp.Body != nil {
			return nil, parseErrorBodytoError(resp.Body)
		}
		return nil, err
	}
	return resp.Body, nil
}

func (rc *Client) httpPut(path string, opts ...httpclient.RequestOption) ([]byte, error) {
	authOpt, authErr := rc.authWrap()
	if authErr != nil {
		return nil, authErr
	}
	opts = append(opts, authOpt...)
	resp, err := httpclient.Put(rc.makeAPIPath(path), opts...)
	if err != nil {
		if resp.Body != nil {
			return nil, parseErrorBodytoError(resp.Body)
		}
		return nil, err
	}

	return resp.Body, err
}

func (rc *Client) httpDelete(path string, opts ...httpclient.RequestOption) error {
	authOpt, authErr := rc.authWrap()
	if authErr != nil {
		return authErr
	}
	opts = append(opts, authOpt...)
	resp, err := httpclient.Delete(rc.makeAPIPath(path), opts...)
	if err != nil {
		if resp.Body != nil {
			return parseErrorBodytoError(resp.Body)
		}
		return err
	}
	return nil
}

func (rc *Client) authWrap() ([]httpclient.RequestOption, error) {
	if rc.Config.AuthMethod == basicAuthType {
		return []httpclient.RequestOption{
			httpclient.AddHeaders(map[string]string{
				"User-Agent": "artifactory-go.v" + rc.Config.APIVersion.String(),
			}),
			httpclient.BasicAuth(rc.Config.Username, rc.Config.Password),
			httpclient.SetClient(rc.HTTPClient),
		}, nil

	}
	headers := make(map[string]string, 2)
	headers["X-JFrog-Art-Api"] = rc.Config.Token
	headers["User-Agent"] = "artifactory-go.v" + rc.Config.APIVersion.String()

	return []httpclient.RequestOption{
		httpclient.AddHeaders(headers),
		httpclient.SetClient(rc.HTTPClient),
	}, nil
}
