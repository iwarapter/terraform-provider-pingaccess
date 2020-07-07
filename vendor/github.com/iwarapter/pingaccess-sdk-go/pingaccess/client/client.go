package client

import (
	"io"

	"github.com/iwarapter/pingaccess-sdk-go/pingaccess/client/metadata"
	"github.com/iwarapter/pingaccess-sdk-go/pingaccess/config"
	"github.com/iwarapter/pingaccess-sdk-go/pingaccess/request"
)

type Options struct {
	Config         config.Config
	CustomCABundle io.Reader
}

// ConfigProvider provides a generic way for a service client to receive
// the ClientConfig without circular dependencies.
type ConfigProvider interface {
	ClientConfig(serviceName string, cfgs ...*config.Config) config.Config
}

// A Client implements the base client request and response handling
// used by all service clients.
type Client struct {
	//request.Retryer
	metadata.ClientInfo

	Config config.Config
	//Handlers request.Handlers
}

func New(cfg config.Config, info metadata.ClientInfo) *Client {
	svc := &Client{
		Config:     cfg,
		ClientInfo: info,
		//Handlers:   handlers.Copy(),
	}
	return svc
}

// NewRequest returns a new Request pointer for the service API
// operation and parameters.
func (c *Client) NewRequest(operation *request.Operation, params interface{}, data interface{}) *request.Request {
	return request.New(c.Config, c.ClientInfo, operation, params, data)
}
