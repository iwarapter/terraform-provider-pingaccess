package proxies

import (
	"net/http"
	"strings"

	"github.com/iwarapter/pingaccess-sdk-go/pingaccess"
	"github.com/iwarapter/pingaccess-sdk-go/pingaccess/client"
	"github.com/iwarapter/pingaccess-sdk-go/pingaccess/client/metadata"
	"github.com/iwarapter/pingaccess-sdk-go/pingaccess/config"
	"github.com/iwarapter/pingaccess-sdk-go/pingaccess/models"
	"github.com/iwarapter/pingaccess-sdk-go/pingaccess/request"
)

const (
	// ServiceName - The name of service.
	ServiceName = "Proxies"
)

//ProxiesService provides the API operations for making requests to
// Proxies endpoint.
type ProxiesService struct {
	*client.Client
}

//New createa a new instance of the ProxiesService client.
//
// Example:
//   cfg := config.NewConfig().WithUsername("Administrator").WithPassword("2Access").WithEndpoint(paURL)
//
//   //Create a ProxiesService from the configuration
//   svc := proxies.New(cfg)
//
func New(cfg *config.Config) *ProxiesService {

	return &ProxiesService{Client: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingaccess.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a Proxies operation
func (s *ProxiesService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := s.NewRequest(op, params, data)

	return req
}

//GetProxiesCommand - Get all Proxies
//RequestType: GET
//Input: input *GetProxiesCommandInput
func (s *ProxiesService) GetProxiesCommand(input *GetProxiesCommandInput) (output *models.HttpClientProxyView, resp *http.Response, err error) {
	path := "/proxies"
	op := &request.Operation{
		Name:       "GetProxiesCommand",
		HTTPMethod: "GET",
		HTTPPath:   path,
		QueryParams: map[string]string{
			"page":          input.Page,
			"numberPerPage": input.NumberPerPage,
			"filter":        input.Filter,
			"name":          input.Name,
			"sortKey":       input.SortKey,
			"order":         input.Order,
		},
	}
	output = &models.HttpClientProxyView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// GetProxiesCommandInput - Inputs for GetProxiesCommand
type GetProxiesCommandInput struct {
	Page          string
	NumberPerPage string
	Filter        string
	Name          string
	SortKey       string
	Order         string
}

//AddProxyCommand - Create a Proxy
//RequestType: POST
//Input: input *AddProxyCommandInput
func (s *ProxiesService) AddProxyCommand(input *AddProxyCommandInput) (output *models.HttpClientProxyView, resp *http.Response, err error) {
	path := "/proxies"
	op := &request.Operation{
		Name:        "AddProxyCommand",
		HTTPMethod:  "POST",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.HttpClientProxyView{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// AddProxyCommandInput - Inputs for AddProxyCommand
type AddProxyCommandInput struct {
	Body models.HttpClientProxyView
}

//DeleteProxyCommand - Delete a Proxy
//RequestType: DELETE
//Input: input *DeleteProxyCommandInput
func (s *ProxiesService) DeleteProxyCommand(input *DeleteProxyCommandInput) (resp *http.Response, err error) {
	path := "/proxies/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "DeleteProxyCommand",
		HTTPMethod:  "DELETE",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}

	req := s.newRequest(op, nil, nil)

	if req.Send() == nil {
		return req.HTTPResponse, nil
	}
	return req.HTTPResponse, req.Error
}

// DeleteProxyCommandInput - Inputs for DeleteProxyCommand
type DeleteProxyCommandInput struct {
	Id string
}

//GetProxyCommand - Get a Proxy
//RequestType: GET
//Input: input *GetProxyCommandInput
func (s *ProxiesService) GetProxyCommand(input *GetProxyCommandInput) (output *models.HttpClientProxyView, resp *http.Response, err error) {
	path := "/proxies/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "GetProxyCommand",
		HTTPMethod:  "GET",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.HttpClientProxyView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// GetProxyCommandInput - Inputs for GetProxyCommand
type GetProxyCommandInput struct {
	Id string
}

//UpdateProxyCommand - Update a Proxy
//RequestType: PUT
//Input: input *UpdateProxyCommandInput
func (s *ProxiesService) UpdateProxyCommand(input *UpdateProxyCommandInput) (output *models.HttpClientProxyView, resp *http.Response, err error) {
	path := "/proxies/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "UpdateProxyCommand",
		HTTPMethod:  "PUT",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.HttpClientProxyView{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// UpdateProxyCommandInput - Inputs for UpdateProxyCommand
type UpdateProxyCommandInput struct {
	Body models.HttpClientProxyView
	Id   string
}
