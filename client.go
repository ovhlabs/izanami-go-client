package izanami_client

type client struct {
	hostname     string
	clientID     string
	clientSecret string
}

type featureClient struct {
	client *client
}

func New(host, clientID, secret string) *client {
	return &client{
		hostname:     host,
		clientID:     clientID,
		clientSecret: secret,
	}
}

func (c *client) Feature() *featureClient {
	return &featureClient{
		c,
	}
}
