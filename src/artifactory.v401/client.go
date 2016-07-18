package artifactory

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"net"
	"time"
)

type ClientConfig struct {
	BaseURL   string
	Username  string
	Password  string
	VerifySSL bool
	Client    *http.Client
	Transport *http.Transport
	Timeout   time.Duration
}

type ArtifactoryClient struct {
	Client    *http.Client
	Config    *ClientConfig
	Transport *http.Transport
}

func NewClient(config *ClientConfig) (c ArtifactoryClient) {
	verifySSL := func() bool {
		if config.VerifySSL != true {
			return false
		} else {
			return true
		}
	}
	if config.Transport == nil {
		config.Transport = new(http.Transport)
	}
	config.Transport.TLSClientConfig = &tls.Config{InsecureSkipVerify: verifySSL()}
	if config.Timeout != 0 {
		config.Transport.Dial = func(network, addr string) (net.Conn, error) {
			return net.DialTimeout(network, addr, config.Timeout)
		}
	}

	if config.Client == nil {
		config.Client = new(http.Client)
	}
	config.Client.Transport = config.Transport
	return ArtifactoryClient{Client: config.Client, Config: config}
}

func clientConfigFrom(from string) (c *ClientConfig) {
	switch from {
	case "environment":
		if os.Getenv("ARTIFACTORY_URL") == "" || os.Getenv("ARTIFACTORY_USERNAME") == "" || os.Getenv("ARTIFACTORY_PASSWORD") == "" {
			fmt.Printf("You must set the environment variables ARTIFACTORY_URL/ARTIFACTORY_USERNAME/ARTIFACTORY_PASSWORD\n")
			os.Exit(1)
		}
	}
	conf := ClientConfig{
		BaseURL:  os.Getenv("ARTIFACTORY_URL"),
		Username: os.Getenv("ARTIFACTORY_USERNAME"),
		Password: os.Getenv("ARTIFACTORY_PASSWORD"),
	}
	return &conf
}

func NewClientFromEnv() (c ArtifactoryClient) {
	config := clientConfigFrom("environment")
	client := NewClient(config)
	return client
}
