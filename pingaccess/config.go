package pingaccess

import (
	"crypto/tls"
	"net/http"
	"net/url"

	"github.com/iwarapter/pingaccess-sdk-go/pingaccess"
)

type Config struct {
	Username string
	Password string
	BaseURL  string
}

// Client configures and returns a fully initialized PAClient
func (c *Config) Client() (interface{}, error) {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	url, _ := url.Parse("https://localhost:9000/")
	client := pingaccess.NewClient("Administrator", "2Access2", url, nil)

	return client, nil
}
