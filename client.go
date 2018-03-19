package client

type Client struct {
	hostname     string
	clientID     string
	clientSecret string
}

type FeatureClient struct {
	client *Client
}

// New creates a new izanami client
func New(host, clientID, secret string) *Client {
	return &Client{
		hostname:     host,
		clientID:     clientID,
		clientSecret: secret,
	}
}

// Feature creates a specific client for feature management
func (c *Client) Feature() *FeatureClient {
	return &FeatureClient{
		c,
	}
}
