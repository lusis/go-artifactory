package responses

import (
	"encoding/json"
	"testing"

	"github.com/lusis/go-artifactory/pkg/artifactory/responses/testdata"
	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/require"
)

func TestGetUsersResponse(t *testing.T) {
	obj := &GetUsersResponse{}
	data, dataErr := testdata.GetBytes(GetUsersResponseTestData)
	if dataErr != nil {
		t.Fatalf(dataErr.Error())
	}
	placeholder := make(map[string]interface{})
	_ = json.Unmarshal(data, &placeholder)
	config := newMSDecoderConfig()
	config.Result = obj
	decoder, newErr := mapstructure.NewDecoder(config)
	require.NoError(t, newErr)
	dErr := decoder.Decode(placeholder)
	require.NoError(t, dErr)
	require.Implements(t, (*VersionedResponse)(nil), obj)
}

func TestGetUserDetailsResponse(t *testing.T) {
	obj := &GetUserDetailsResponse{}
	data, dataErr := testdata.GetBytes(GetUserDetailsResponseTestData)
	if dataErr != nil {
		t.Fatalf(dataErr.Error())
	}
	placeholder := make(map[string]interface{})
	_ = json.Unmarshal(data, &placeholder)
	config := newMSDecoderConfig()
	config.Result = obj
	decoder, newErr := mapstructure.NewDecoder(config)
	require.NoError(t, newErr)
	dErr := decoder.Decode(placeholder)
	require.NoError(t, dErr)
	require.Implements(t, (*VersionedResponse)(nil), obj)
}
