package pingaccess

import (
	"github.com/iwarapter/pingaccess-sdk-go/pingaccess"
	"github.com/iwarapter/pingaccess-sdk-go/service/applications"
	"github.com/iwarapter/pingaccess-sdk-go/service/sites"
	"github.com/iwarapter/pingaccess-sdk-go/service/virtualhosts"
)

type Config struct {
	Username string
	Password string
	BaseURL  string
}

type PAClient struct {
	vhconn   *virtualhosts.Virtualhosts
	siteconn *sites.Sites
	appconn  *applications.Applications
}

// Client configures and returns a fully initialized PAClient
func (c *Config) Client() (interface{}, error) {
	var client PAClient

	var conf = pingaccess.Config{
		Username: c.Username,
		Password: c.Password,
		BaseURL:  c.BaseURL,
	}

	client.vhconn = virtualhosts.New(&conf)
	client.siteconn = sites.New(&conf)
	client.appconn = applications.New(&conf)

	return &client, nil
}
