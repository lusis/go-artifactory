package artifactory

import (
	"bytes"
	"encoding/json"
	"fmt"

	multierror "github.com/hashicorp/go-multierror"
	"github.com/lusis/go-artifactory/pkg/artifactory/requests"
	"github.com/lusis/go-artifactory/pkg/artifactory/responses"
)

// Users is a collection of User
type Users []responses.UserResponse

// UserDetails is a user's details
type UserDetails struct {
	responses.GetUserDetailsResponse
}

// GetUsers returns a list of users
func (c *Client) GetUsers() (Users, error) {
	if err := c.checkRequiredResponseVersion(responses.GetUsersResponse{}); err != nil {
		return nil, err
	}
	users := Users{}
	res, err := c.httpGet("security/users", requestJSON(), requestExpects(200))
	if err != nil {
		return nil, err
	}
	if jsonErr := json.Unmarshal(res, &users); jsonErr != nil {
		return nil, &UnmarshalError{msg: multierror.Append(errDecoding, jsonErr).Error()}
	}
	return users, nil
}

// GetUserDetails gets a user's details
func (c *Client) GetUserDetails(u string) (*UserDetails, error) {
	if err := c.checkRequiredResponseVersion(responses.GetUserDetailsResponse{}); err != nil {
		return nil, err
	}
	users := &UserDetails{}
	res, err := c.httpGet("security/users/"+u, requestJSON(), requestExpects(200))
	if err != nil {
		return nil, err
	}
	if jsonErr := json.Unmarshal(res, users); jsonErr != nil {
		return nil, &UnmarshalError{msg: multierror.Append(errDecoding, jsonErr).Error()}
	}
	return users, nil
}

// CreateOrReplaceUser creates or updates the provided user with the provided details
func (c *Client) CreateOrReplaceUser(u string, details requests.CreateOrReplaceUserRequest) error {
	if err := c.checkRequiredRequestVersion(details); err != nil {
		return err
	}
	j, err := json.Marshal(details)
	if err != nil {
		return &MarshalError{msg: multierror.Append(errDecoding, err).Error()}
	}
	url := fmt.Sprintf("security/users/%s", u)
	_, err = c.httpPut(url, requestExpects(200), requestJSON(), withBody(bytes.NewReader(j)))
	return err

}

// UpdateUser creates or updates the provided user with the provided details
func (c *Client) UpdateUser(u string, details requests.UpdateUserRequest) error {
	if err := c.checkRequiredRequestVersion(details); err != nil {
		return err
	}
	j, err := json.Marshal(details)
	if err != nil {
		return &MarshalError{msg: multierror.Append(errDecoding, err).Error()}
	}
	url := fmt.Sprintf("security/users/%s", u)
	_, err = c.httpPost(url, requestExpects(200), requestJSON(), withBody(bytes.NewReader(j)))
	return err
}

// DeleteUser creates or updates the provided user with the provided details
func (c *Client) DeleteUser(u string) error {
	if err := c.checkRequiredRequestVersion(requests.GenericVersionedRequest{}); err != nil {
		return err
	}
	url := fmt.Sprintf("security/users/%s", u)
	return c.httpDelete(url, requestExpects(200), requestJSON())
}

// GetUserEncryptedPassword returns the encrypted password for the current user
func (c *Client) GetUserEncryptedPassword() (string, error) {
	if err := c.checkRequiredResponseVersion(responses.GetEncryptedPasswordResponse{}); err != nil {
		return "", err
	}
	res, err := c.httpGet("security/encryptedPassword", requestExpects(200))
	if err != nil {
		return "", err
	}
	return string(res), nil
}
