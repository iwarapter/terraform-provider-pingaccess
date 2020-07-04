package redirects

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
	ServiceName = "Redirects"
)

//RedirectsService provides the API operations for making requests to
// Redirects endpoint.
type RedirectsService struct {
	*client.Client
}

//New createa a new instance of the RedirectsService client.
//
// Example:
//   cfg := config.NewConfig().WithUsername("Administrator").WithPassword("2FederateM0re").WithEndpoint(paURL.String())
//
//   //Create a RedirectsService from the configuration
//   svc := redirects.New(cfg)
//
func New(cfg *config.Config) *RedirectsService {

	return &RedirectsService{Client: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingaccess.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a Redirects operation
func (s *RedirectsService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := s.NewRequest(op, params, data)

	return req
}

//GetRedirectsCommand - Get all Redirects
//RequestType: GET
//Input: input *GetRedirectsCommandInput
func (s *RedirectsService) GetRedirectsCommand(input *GetRedirectsCommandInput) (output *models.RedirectsView, resp *http.Response, err error) {
	path := "/redirects"
	op := &request.Operation{
		Name:       "GetRedirectsCommand",
		HTTPMethod: "GET",
		HTTPPath:   path,
		QueryParams: map[string]string{
			"page":          input.Page,
			"numberPerPage": input.NumberPerPage,
			"filter":        input.Filter,
			"source":        input.Source,
			"target":        input.Target,
			"sortKey":       input.SortKey,
			"order":         input.Order,
		},
	}
	output = &models.RedirectsView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// GetRedirectsCommandInput - Inputs for GetRedirectsCommand
type GetRedirectsCommandInput struct {
	Page          string
	NumberPerPage string
	Filter        string
	Source        string
	Target        string
	SortKey       string
	Order         string
}

//AddRedirectCommand - Add a Redirect
//RequestType: POST
//Input: input *AddRedirectCommandInput
func (s *RedirectsService) AddRedirectCommand(input *AddRedirectCommandInput) (output *models.RedirectView, resp *http.Response, err error) {
	path := "/redirects"
	op := &request.Operation{
		Name:        "AddRedirectCommand",
		HTTPMethod:  "POST",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.RedirectView{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// AddRedirectCommandInput - Inputs for AddRedirectCommand
type AddRedirectCommandInput struct {
	Body models.RedirectView
}

//DeleteRedirectCommand - Delete a Redirect
//RequestType: DELETE
//Input: input *DeleteRedirectCommandInput
func (s *RedirectsService) DeleteRedirectCommand(input *DeleteRedirectCommandInput) (resp *http.Response, err error) {
	path := "/redirects/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "DeleteRedirectCommand",
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

// DeleteRedirectCommandInput - Inputs for DeleteRedirectCommand
type DeleteRedirectCommandInput struct {
	Id string
}

//GetRedirectCommand - Get a Redirect
//RequestType: GET
//Input: input *GetRedirectCommandInput
func (s *RedirectsService) GetRedirectCommand(input *GetRedirectCommandInput) (output *models.RedirectView, resp *http.Response, err error) {
	path := "/redirects/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "GetRedirectCommand",
		HTTPMethod:  "GET",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.RedirectView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// GetRedirectCommandInput - Inputs for GetRedirectCommand
type GetRedirectCommandInput struct {
	Id string
}

//UpdateRedirectCommand - Update a Redirect
//RequestType: PUT
//Input: input *UpdateRedirectCommandInput
func (s *RedirectsService) UpdateRedirectCommand(input *UpdateRedirectCommandInput) (output *models.RedirectView, resp *http.Response, err error) {
	path := "/redirects/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "UpdateRedirectCommand",
		HTTPMethod:  "PUT",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.RedirectView{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// UpdateRedirectCommandInput - Inputs for UpdateRedirectCommand
type UpdateRedirectCommandInput struct {
	Body models.RedirectView
	Id   string
}
