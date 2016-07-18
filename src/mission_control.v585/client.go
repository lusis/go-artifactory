package mission_control

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"time"
	"net"
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

type MissionControlClient struct {
	Client    *http.Client
	Config    *ClientConfig
	Transport *http.Transport
}

func NewClient(config *ClientConfig) (c MissionControlClient) {
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
	return MissionControlClient{Client: config.Client, Config: config}
}

func clientConfigFrom(from string) (c *ClientConfig) {
	switch from {
	case "environment":
		if os.Getenv("MISSION_CONTROL_URL") == "" || os.Getenv("MISSION_CONTROL_USERNAME") == "" || os.Getenv("MISSION_CONTROL_PASSWORD") == "" {
			fmt.Printf("You must set the environment variables MISSION_CONTROL_URL/MISSION_CONTROL_USERNAME/MISSION_CONTROL_PASSWORD\n")
			os.Exit(1)
		}
	}
	conf := ClientConfig{
		BaseURL:  os.Getenv("MISSION_CONTROL_URL"),
		Username: os.Getenv("MISSION_CONTROL_USERNAME"),
		Password: os.Getenv("MISSION_CONTROL_PASSWORD"),
	}
	return &conf
}

func NewClientFromEnv() (c MissionControlClient) {
	config := clientConfigFrom("environment")
	client := NewClient(config)
	return client
}
