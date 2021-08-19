package globalUnprotectedResources

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
	ServiceName = "GlobalUnprotectedResources"
)

//GlobalUnprotectedResourcesService provides the API operations for making requests to
// GlobalUnprotectedResources endpoint.
type GlobalUnprotectedResourcesService struct {
	*client.Client
}

//New createa a new instance of the GlobalUnprotectedResourcesService client.
//
// Example:
//   cfg := config.NewConfig().WithUsername("Administrator").WithPassword("2Access").WithEndpoint(paURL)
//
//   //Create a GlobalUnprotectedResourcesService from the configuration
//   svc := globalUnprotectedResources.New(cfg)
//
func New(cfg *config.Config) *GlobalUnprotectedResourcesService {

	return &GlobalUnprotectedResourcesService{Client: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingaccess.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a GlobalUnprotectedResources operation
func (s *GlobalUnprotectedResourcesService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := s.NewRequest(op, params, data)

	return req
}

//GetGlobalUnprotectedResourcesCommand - Get all Global Unprotected Resources
//RequestType: GET
//Input: input *GetGlobalUnprotectedResourcesCommandInput
func (s *GlobalUnprotectedResourcesService) GetGlobalUnprotectedResourcesCommand(input *GetGlobalUnprotectedResourcesCommandInput) (output *models.GlobalUnprotectedResourcesView, resp *http.Response, err error) {
	path := "/globalUnprotectedResources"
	op := &request.Operation{
		Name:       "GetGlobalUnprotectedResourcesCommand",
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
	output = &models.GlobalUnprotectedResourcesView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// GetGlobalUnprotectedResourcesCommandInput - Inputs for GetGlobalUnprotectedResourcesCommand
type GetGlobalUnprotectedResourcesCommandInput struct {
	Page          string
	NumberPerPage string
	Filter        string
	Name          string
	SortKey       string
	Order         string
}

//AddGlobalUnprotectedResourceCommand - Add a Global Unprotected Resource
//RequestType: POST
//Input: input *AddGlobalUnprotectedResourceCommandInput
func (s *GlobalUnprotectedResourcesService) AddGlobalUnprotectedResourceCommand(input *AddGlobalUnprotectedResourceCommandInput) (output *models.GlobalUnprotectedResourceView, resp *http.Response, err error) {
	path := "/globalUnprotectedResources"
	op := &request.Operation{
		Name:        "AddGlobalUnprotectedResourceCommand",
		HTTPMethod:  "POST",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.GlobalUnprotectedResourceView{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// AddGlobalUnprotectedResourceCommandInput - Inputs for AddGlobalUnprotectedResourceCommand
type AddGlobalUnprotectedResourceCommandInput struct {
	Body models.GlobalUnprotectedResourceView
}

//DeleteGlobalUnprotectedResourceCommand - Delete a Global Unprotected Resource
//RequestType: DELETE
//Input: input *DeleteGlobalUnprotectedResourceCommandInput
func (s *GlobalUnprotectedResourcesService) DeleteGlobalUnprotectedResourceCommand(input *DeleteGlobalUnprotectedResourceCommandInput) (resp *http.Response, err error) {
	path := "/globalUnprotectedResources/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "DeleteGlobalUnprotectedResourceCommand",
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

// DeleteGlobalUnprotectedResourceCommandInput - Inputs for DeleteGlobalUnprotectedResourceCommand
type DeleteGlobalUnprotectedResourceCommandInput struct {
	Id string
}

//GetGlobalUnprotectedResourceCommand - Get a Global Unprotected Resource
//RequestType: GET
//Input: input *GetGlobalUnprotectedResourceCommandInput
func (s *GlobalUnprotectedResourcesService) GetGlobalUnprotectedResourceCommand(input *GetGlobalUnprotectedResourceCommandInput) (output *models.GlobalUnprotectedResourceView, resp *http.Response, err error) {
	path := "/globalUnprotectedResources/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "GetGlobalUnprotectedResourceCommand",
		HTTPMethod:  "GET",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.GlobalUnprotectedResourceView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// GetGlobalUnprotectedResourceCommandInput - Inputs for GetGlobalUnprotectedResourceCommand
type GetGlobalUnprotectedResourceCommandInput struct {
	Id string
}

//UpdateGlobalUnprotectedResourceCommand - Update a Global Unprotected Resource
//RequestType: PUT
//Input: input *UpdateGlobalUnprotectedResourceCommandInput
func (s *GlobalUnprotectedResourcesService) UpdateGlobalUnprotectedResourceCommand(input *UpdateGlobalUnprotectedResourceCommandInput) (output *models.GlobalUnprotectedResourceView, resp *http.Response, err error) {
	path := "/globalUnprotectedResources/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "UpdateGlobalUnprotectedResourceCommand",
		HTTPMethod:  "PUT",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.GlobalUnprotectedResourceView{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// UpdateGlobalUnprotectedResourceCommandInput - Inputs for UpdateGlobalUnprotectedResourceCommand
type UpdateGlobalUnprotectedResourceCommandInput struct {
	Body models.GlobalUnprotectedResourceView
	Id   string
}
