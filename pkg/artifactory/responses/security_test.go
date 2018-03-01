package responses

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/lusis/go-artifactory/pkg/artifactory/responses/testdata"
	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/require"
)

func TestSecurityResponses(t *testing.T) {
	// Add any new test cases here
	// in the form of
	// {&FooResponse{}:FooResponseTestData}
	responsesTestCases := []map[VersionedResponse]string{
		{&GetUsersResponse{}: GetUsersResponseTestData},
		{&GetUserDetailsResponse{}: GetUserDetailsResponseTestData},
		{&GetPasswordExpirationPolicyResponse{}: GetPasswordExpirationPolicyResponseTestData},
		{&CreateAPIKeyResponse{}: CreateAPIKeyResponseTestData},
		{&RegenerateAPIKeyResponse{}: RegenerateAPIKeyResponseTestData},
		{&GetAPIKeyResponse{}: GetAPIKeyResponseTestData},
		{&GetGroupsResponse{}: GetGroupsResponseTestData},
		{&GetGroupDetailsResponse{}: GetGroupDetailsResponseTestData},
		{&GetPermissionTargetsResponse{}: GetPermissionTargetsResponseTestData},
		{&GetPermissionTargetDetailsResponse{}: GetPermissionTargetDetailsResponseTestData},
		{&CreateTokenResponse{}: CreateTokenResponseTestData},
		{&RefreshTokenResponse{}: RefreshTokenResponseTestData},
		{&GetLockedOutUsersResponse{}: GetLockedOutUsersResponseTestData},
		{&GetCertificatesResponse{}: GetCertificatesResponseTestData},
	}

	for _, testCase := range responsesTestCases {
		for k, v := range testCase {
			require.Implements(t, (*VersionedResponse)(nil), k)
			data, err := testdata.GetBytes(v)
			require.NoError(t, err)
			placeholder := make(map[string]interface{})
			_ = json.Unmarshal(data, &placeholder)
			config := newMSDecoderConfig()
			config.Result = k
			decoder, newErr := mapstructure.NewDecoder(config)
			require.NoError(t, newErr)
			dErr := decoder.Decode(placeholder)
			require.NoError(t, dErr, fmt.Sprintf("should parse %s", v))
		}
	}
}
