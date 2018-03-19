package izanami_client

type client struct {
	hostname     string
	clientID     string
	clientSecret string
}

var c *client

func New(conf Configuration) {
	c = &client{
		hostname:     conf.Host,
		clientID:     conf.ClientID,
		clientSecret: conf.ClientSecret,
	}
}
