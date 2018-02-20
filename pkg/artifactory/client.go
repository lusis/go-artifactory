package artifactory

import (
	"crypto/tls"
	"errors"
	"net/http"
	"os"
)

// ClientConfig is an artifactory client configuration
type ClientConfig struct {
	BaseURL    string
	Username   string
	Password   string
	Token      string
	AuthMethod string
	VerifySSL  bool
	Debug      bool
	APIVersion Version
	HTTPClient *http.Client
}

// Client is an artifactory client
type Client struct {
	Config     *ClientConfig
	HTTPClient *http.Client
}

// NewClientFromEnv returns a new client configured from environment variables
func NewClientFromEnv() (*Client, error) {
	config, err := clientConfigFrom("environment")
	if err != nil {
		return nil, err
	}

	return NewClient(config)
}

// NewClient creates a new client from the provided `ClientConfig`
func NewClient(config *ClientConfig) (*Client, error) {
	if config.HTTPClient == nil {
		tlsClientConfig := tls.Config{InsecureSkipVerify: !config.VerifySSL}
		transport := &http.Transport{TLSClientConfig: &tlsClientConfig}
		client := &http.Client{Transport: transport}
		config.HTTPClient = client
	}
	client := Client{
		HTTPClient: config.HTTPClient,
		Config:     config,
	}
	if config.APIVersion.Version == nil {
		config.APIVersion = MaxArtifactoryVersion
	}
	return &client, nil
}

func clientConfigFrom(from string) (*ClientConfig, error) {
	conf := ClientConfig{}
	switch from {
	case "environment":
		if os.Getenv("ARTIFACTORY_URL") == "" {
			return nil, errors.New("You must set the environment variable ARTIFACTORY_URL")
		}

		conf.BaseURL = os.Getenv("ARTIFACTORY_URL")
		if os.Getenv("ARTIFACTORY_TOKEN") == "" {
			if os.Getenv("ARTIFACTORY_USERNAME") == "" || os.Getenv("ARTIFACTORY_PASSWORD") == "" {
				return nil, errors.New("You must set the environment variables ARTIFACTORY_USERNAME/ARTIFACTORY_PASSWORD if not using ARTIFACTORY_TOKEN")
			}

			conf.AuthMethod = basicAuthType
		} else {
			conf.AuthMethod = tokenAuthType
		}
	}
	if conf.AuthMethod == "token" {
		conf.Token = os.Getenv("ARTIFACTORY_TOKEN")
	} else {
		conf.Username = os.Getenv("ARTIFACTORY_USERNAME")
		conf.Password = os.Getenv("ARTIFACTORY_PASSWORD")
	}
	if os.Getenv("ARTIFACTORY_DEBUG") != "" {
		conf.Debug = true
	}
	if os.Getenv("ARTIFACTORY_VERSION") != "" {
		conf.APIVersion = versionMustParse(os.Getenv("ARTIFACTORY_VERSION"))
	}
	return &conf, nil
}
