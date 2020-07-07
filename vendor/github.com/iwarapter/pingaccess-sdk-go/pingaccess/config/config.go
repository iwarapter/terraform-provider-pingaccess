package config

import (
	"crypto/tls"
	"net/http"

	"github.com/iwarapter/pingaccess-sdk-go/pingaccess"
)

type Config struct {
	Username          *string
	Password          *string
	LogDebug          *bool
	MaskAuthorization *bool
	Endpoint          *string

	// The HTTP client to use when sending requests. Defaults to
	// `http.DefaultClient`.
	HTTPClient *http.Client
}

func NewConfig() *Config {
	//TODO this is epically shit dont leave it here, parse a ca bundle etc
	/* #nosec */
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	return &Config{
		MaskAuthorization: pingaccess.Bool(true),
		LogDebug:          pingaccess.Bool(false),
		HTTPClient:        http.DefaultClient,
	}
}

func (c *Config) WithPassword(password string) *Config {
	c.Password = pingaccess.String(password)
	return c
}

func (c *Config) WithUsername(username string) *Config {
	c.Username = pingaccess.String(username)
	return c
}

func (c *Config) WithEndpoint(endpoint string) *Config {
	c.Endpoint = pingaccess.String(endpoint)
	return c
}

func (c *Config) WithDebug(debug bool) *Config {
	c.LogDebug = pingaccess.Bool(debug)
	return c
}

func (c *Config) WithMaskAuthorization(debug bool) *Config {
	c.MaskAuthorization = pingaccess.Bool(debug)
	return c
}
