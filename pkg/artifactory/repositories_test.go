package artifactory

import (
	"fmt"
	"testing"

	"github.com/lusis/go-artifactory/pkg/artifactory/responses"
	"github.com/lusis/go-artifactory/pkg/artifactory/responses/testdata"
	"github.com/lusis/go-artifactory/pkg/httpclient"
	"github.com/stretchr/testify/require"
)

func TestGetRepositories(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.GetRepositoriesResponseTestData)
	require.NoError(t, err)
	cases := map[string]struct {
		Data []byte
		Code int
		Fail bool
	}{
		"parse good data": {
			Data: jsonfile,
			Code: 200,
			Fail: false,
		},
		"not parse bad data": {
			Data: []byte(""),
			Code: 200,
			Fail: true,
		},
		"bad return code": {
			Data: jsonfile,
			Code: 500,
			Fail: true,
		},
	}

	for desc, testCase := range cases {
		client, server, err := newTestClient(testCase.Data, "application/json", testCase.Code)
		defer server.Close()
		require.NoError(t, err)
		obj, err := client.GetRepositories()
		if testCase.Fail {
			require.Error(t, err, desc)
			require.Nil(t, obj, desc)
		} else {
			require.NoError(t, err, desc)
			require.NotNil(t, obj, desc)
		}
	}
}

func TestGetRepositoryConfiguration(t *testing.T) {
	localrepo, err := testdata.GetBytes(responses.LocalRepositoryConfigurationResponseTestData)
	require.NoError(t, err)
	remoterepo, err := testdata.GetBytes(responses.RemoteRepositoryConfigurationResponseTestData)
	require.NoError(t, err)
	virtualrepo, err := testdata.GetBytes(responses.VirtualRepositoryConfigurationResponseTestData)
	require.NoError(t, err)
	cases := map[string]struct {
		Data     []byte
		Code     int
		Fail     bool
		RepoType string
	}{
		"parse good data - local repo": {
			Data:     localrepo,
			Code:     200,
			Fail:     false,
			RepoType: "local",
		},
		"parse good data - remote repo": {
			Data:     remoterepo,
			Code:     200,
			Fail:     false,
			RepoType: "remote",
		},
		"parse good data - virtual repo": {
			Data:     virtualrepo,
			Code:     200,
			Fail:     false,
			RepoType: "virtual",
		},
		"not parse bad data": {
			Data:     []byte(""),
			Code:     200,
			Fail:     true,
			RepoType: "local",
		},
		"bad return code": {
			Data:     localrepo,
			Code:     500,
			Fail:     true,
			RepoType: "local",
		},
	}

	for desc, testCase := range cases {
		client, server, err := newTestClient(testCase.Data, "application/json", testCase.Code)
		defer server.Close()
		require.NoError(t, err)
		var obj RepositoryConfiguration
		switch testCase.RepoType {
		case "local":
			obj, err = client.GetLocalRepositoryConfiguration("foo")
		case "remote":
			obj, err = client.GetRemoteRepositoryConfiguration("foo")
		case "virtual":
			obj, err = client.GetVirtualRepositoryConfiguration("foo")
		}

		if testCase.Fail {
			require.Error(t, err, desc)
			require.Nil(t, obj, desc)
		} else {
			require.NoError(t, err, desc)
			require.NotNil(t, obj, desc)
		}
	}
}

func TestCreateRepository(t *testing.T) {
	badOpt := func() CreateOrUpdateRepositoryOption {
		return func(m map[string]interface{}) error {
			return fmt.Errorf("option setting failed")
		}
	}
	testCases := map[string]struct {
		code int
		fail bool
		opts []CreateOrUpdateRepositoryOption
	}{
		"pass no opts":      {200, false, nil},
		"fail no opts":      {500, true, nil},
		"fail with bad opt": {200, true, []CreateOrUpdateRepositoryOption{badOpt()}},
		"pass with opt": {200, false, []CreateOrUpdateRepositoryOption{
			RepositoryBoolOption("boolval", false),
			RepositoryIntOption("intval", 1),
			RepositoryStringOption("stringval", "string"),
			RepositoryStringSliceOption("stringsliceval", []string{"foo", "bar"}),
		}},
		"fail with opt": {500, true, []CreateOrUpdateRepositoryOption{
			RepositoryBoolOption("boolval", false),
			RepositoryIntOption("intval", 1),
			RepositoryStringOption("stringval", "string"),
			RepositoryStringSliceOption("stringsliceval", []string{"foo", "bar"}),
		}},
	}
	for desc, testCase := range testCases {
		client, server, err := newTestClient([]byte(""), "application/json", testCase.code)
		defer server.Close()
		require.NoError(t, err)
		err = client.CreateLocalRepository("testrepo", testCase.opts...)
		if testCase.fail {
			require.Error(t, err, desc)
		} else {
			require.NoError(t, err, desc)
		}
		err = client.CreateRemoteRepository("testrepo", "http://foo", testCase.opts...)
		if testCase.fail {
			require.Error(t, err, desc)
		} else {
			require.NoError(t, err, desc)
		}
		err = client.CreateVirtualRepository("testrepo", "generic", testCase.opts...)
		if testCase.fail {
			require.Error(t, err, desc)
		} else {
			require.NoError(t, err, desc)
		}
	}
}

func TestUpdateRepositoryConfiguration(t *testing.T) {
	badOpt := func() CreateOrUpdateRepositoryOption {
		return func(m map[string]interface{}) error {
			return fmt.Errorf("option setting failed")
		}
	}
	testCases := map[string]struct {
		code int
		fail bool
		opts []CreateOrUpdateRepositoryOption
	}{
		"pass no opts":      {200, false, nil},
		"fail no opts":      {500, true, nil},
		"fail with bad opt": {200, true, []CreateOrUpdateRepositoryOption{badOpt()}},
		"pass with opt": {200, false, []CreateOrUpdateRepositoryOption{
			RepositoryBoolOption("boolval", false),
			RepositoryIntOption("intval", 1),
			RepositoryStringOption("stringval", "string"),
			RepositoryStringSliceOption("stringsliceval", []string{"foo", "bar"}),
		}},
		"fail with opt": {500, true, []CreateOrUpdateRepositoryOption{
			RepositoryBoolOption("boolval", false),
			RepositoryIntOption("intval", 1),
			RepositoryStringOption("stringval", "string"),
			RepositoryStringSliceOption("stringsliceval", []string{"foo", "bar"}),
		}},
	}
	for desc, testCase := range testCases {
		client, server, err := newTestClient([]byte(""), "application/json", testCase.code)
		defer server.Close()
		require.NoError(t, err)
		err = client.UpdateRepositoryConfiguration("testrepo", testCase.opts...)
		if testCase.fail {
			require.Error(t, err, desc)
		} else {
			require.NoError(t, err, desc)
		}
	}
}

func TestDeleteRepository(t *testing.T) {
	client, server, err := newTestClient([]byte(""), "text/plain", 200)
	defer server.Close()
	require.NoError(t, err)
	failClient, failServer, failErr := newTestClient([]byte("error"), "text/plain", 500)
	defer failServer.Close()
	require.NoError(t, failErr)
	err = client.DeleteRepository("foo")
	require.NoError(t, err, "delete repo - no error")
	err = failClient.DeleteRepository("foo")
	require.Error(t, err, "delete repo - error")
}

func TestCalculateAndIndex(t *testing.T) {
	client, server, err := newTestClient([]byte(""), "text/plain", 200)
	defer server.Close()
	require.NoError(t, err)
	failClient, failServer, failErr := newTestClient([]byte(""), "text/plain", 500)
	defer failServer.Close()
	require.NoError(t, failErr)

	noOptCases := map[string]func(string) error{
		"nuget": client.CalculateNuGetRepositoryMetadata,
		"npm":   client.CalculateNPMRepositoryMetadata,
		"bower": client.CalculateBowerIndex,
		"helm":  client.CalculateHelmChartIndex,
	}
	optCases := map[string]func(string, ...CalculateOption) error{
		"opkg":           client.CalculateOpkgRepostitoryMetadata,
		"debian":         client.CalculateDebianRepositoryMetadata,
		"maven-metadata": client.CalculateMavenMetadata,
		"yum":            client.CalculateYumRepositoryMetadata,
	}

	noOptCasesFail := map[string]func(string) error{
		"nuget": failClient.CalculateNuGetRepositoryMetadata,
		"npm":   failClient.CalculateNPMRepositoryMetadata,
		"bower": failClient.CalculateBowerIndex,
		"helm":  failClient.CalculateHelmChartIndex,
	}
	optCasesFail := map[string]func(string, ...CalculateOption) error{
		"opkg":           failClient.CalculateOpkgRepostitoryMetadata,
		"debian":         failClient.CalculateDebianRepositoryMetadata,
		"maven-metadata": failClient.CalculateMavenMetadata,
		"yum":            failClient.CalculateYumRepositoryMetadata,
	}

	for name, f := range noOptCases {
		err := f("foo")
		require.NoError(t, err, name+" - no error")
		err = client.CalculateMavenIndex()
		require.NoError(t, err, "maven - no error")
	}

	for name, f := range optCases {
		err := f("foo", CalculateGPGPassphrase("password"), CalculateQueryParams(map[string]string{"foo": "bar"}))
		require.NoError(t, err, name+" - no error with options")
		err = f("foo")
		require.NoError(t, err, name+" - no error with no options")
		failOption := func() CalculateOption {
			return func(h *[]httpclient.RequestOption) error {
				return fmt.Errorf("option setting failed")
			}
		}
		err = f("foo", failOption())
		require.Error(t, err, name+" - bad option")
		err = client.CalculateMavenIndex(failOption())
		require.Error(t, err, "maven - bad option")
	}

	for name, f := range noOptCasesFail {
		err := f("foo")
		require.Error(t, err, name+" - http error")
		err = failClient.CalculateMavenIndex()
		require.Error(t, err, "maven - http error")
	}

	for name, f := range optCasesFail {
		err := f("foo", CalculateGPGPassphrase("password"), CalculateQueryParams(map[string]string{"foo": "bar"}))
		require.Error(t, err, name+" - http error")
	}

}
