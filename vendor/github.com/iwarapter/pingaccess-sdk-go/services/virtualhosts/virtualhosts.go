package virtualhosts

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
	ServiceName = "Virtualhosts"
)

//VirtualhostsService provides the API operations for making requests to
// Virtualhosts endpoint.
type VirtualhostsService struct {
	*client.Client
}

//New createa a new instance of the VirtualhostsService client.
//
// Example:
//   cfg := config.NewConfig().WithUsername("Administrator").WithPassword("2Access").WithEndpoint(paURL)
//
//   //Create a VirtualhostsService from the configuration
//   svc := virtualhosts.New(cfg)
//
func New(cfg *config.Config) *VirtualhostsService {

	return &VirtualhostsService{Client: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingaccess.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a Virtualhosts operation
func (s *VirtualhostsService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := s.NewRequest(op, params, data)

	return req
}

//GetVirtualHostsCommand - Get all Virtual Hosts
//RequestType: GET
//Input: input *GetVirtualHostsCommandInput
func (s *VirtualhostsService) GetVirtualHostsCommand(input *GetVirtualHostsCommandInput) (output *models.VirtualHostsView, resp *http.Response, err error) {
	path := "/virtualhosts"
	op := &request.Operation{
		Name:       "GetVirtualHostsCommand",
		HTTPMethod: "GET",
		HTTPPath:   path,
		QueryParams: map[string]string{
			"page":          input.Page,
			"numberPerPage": input.NumberPerPage,
			"filter":        input.Filter,
			"virtualHost":   input.VirtualHost,
			"sortKey":       input.SortKey,
			"order":         input.Order,
		},
	}
	output = &models.VirtualHostsView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// GetVirtualHostsCommandInput - Inputs for GetVirtualHostsCommand
type GetVirtualHostsCommandInput struct {
	Page          string
	NumberPerPage string
	Filter        string
	VirtualHost   string
	SortKey       string
	Order         string
}

//AddVirtualHostCommand - Create a Virtual Host
//RequestType: POST
//Input: input *AddVirtualHostCommandInput
func (s *VirtualhostsService) AddVirtualHostCommand(input *AddVirtualHostCommandInput) (output *models.VirtualHostView, resp *http.Response, err error) {
	path := "/virtualhosts"
	op := &request.Operation{
		Name:        "AddVirtualHostCommand",
		HTTPMethod:  "POST",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.VirtualHostView{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// AddVirtualHostCommandInput - Inputs for AddVirtualHostCommand
type AddVirtualHostCommandInput struct {
	Body models.VirtualHostView
}

//DeleteVirtualHostCommand - Delete a Virtual Host
//RequestType: DELETE
//Input: input *DeleteVirtualHostCommandInput
func (s *VirtualhostsService) DeleteVirtualHostCommand(input *DeleteVirtualHostCommandInput) (resp *http.Response, err error) {
	path := "/virtualhosts/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "DeleteVirtualHostCommand",
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

// DeleteVirtualHostCommandInput - Inputs for DeleteVirtualHostCommand
type DeleteVirtualHostCommandInput struct {
	Id string
}

//GetVirtualHostCommand - Get a Virtual Host
//RequestType: GET
//Input: input *GetVirtualHostCommandInput
func (s *VirtualhostsService) GetVirtualHostCommand(input *GetVirtualHostCommandInput) (output *models.VirtualHostView, resp *http.Response, err error) {
	path := "/virtualhosts/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "GetVirtualHostCommand",
		HTTPMethod:  "GET",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.VirtualHostView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// GetVirtualHostCommandInput - Inputs for GetVirtualHostCommand
type GetVirtualHostCommandInput struct {
	Id string
}

//UpdateVirtualHostCommand - Update a Virtual Host
//RequestType: PUT
//Input: input *UpdateVirtualHostCommandInput
func (s *VirtualhostsService) UpdateVirtualHostCommand(input *UpdateVirtualHostCommandInput) (output *models.VirtualHostView, resp *http.Response, err error) {
	path := "/virtualhosts/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "UpdateVirtualHostCommand",
		HTTPMethod:  "PUT",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.VirtualHostView{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// UpdateVirtualHostCommandInput - Inputs for UpdateVirtualHostCommand
type UpdateVirtualHostCommandInput struct {
	Body models.VirtualHostView
	Id   string
}
