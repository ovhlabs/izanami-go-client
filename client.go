package client

import (
	"net/http"
	"time"

	"github.com/facebookgo/httpcontrol"
)

// Client represents the izanami client
type Client struct {
	hostname     string
	clientID     string
	clientSecret string
	HttpClient   *http.Client
}

// FeatureClient represents a client for feature management
type FeatureClient struct {
	client *Client
}

// New creates a new izanami client
func New(host, clientID, secret string) *Client {
	return &Client{
		hostname:     host,
		clientID:     clientID,
		clientSecret: secret,
		HttpClient: &http.Client{
			Transport: &httpcontrol.Transport{
				RequestTimeout: time.Second * 30,
				MaxTries:       5,
			},
		},
	}
}

// Feature creates a specific client for feature management
func (c *Client) Feature() *FeatureClient {
	return &FeatureClient{
		c,
	}
}
