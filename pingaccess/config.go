package pingaccess

import (
	"crypto/tls"
	"net/http"
	"net/url"
	"os"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"

	"github.com/iwarapter/pingaccess-sdk-go/pingaccess"
)

type config struct {
	Username string
	Password string
	Context  string
	BaseURL  string
}

// Client configures and returns a fully initialized PAClient
func (c *config) Client() (interface{}, diag.Diagnostics) {
	/* #nosec */
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	url, _ := url.Parse(c.BaseURL)
	client := pingaccess.NewClient(c.Username, c.Password, url, c.Context, nil)

	if os.Getenv("TF_LOG") == "DEBUG" || os.Getenv("TF_LOG") == "TRACE" {
		client.LogDebug = true
	}
	return client, nil
}
