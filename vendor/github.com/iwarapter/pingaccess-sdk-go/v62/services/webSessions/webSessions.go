package webSessions

import (
	"net/http"
	"strings"

	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess"
	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess/client"
	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess/client/metadata"
	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess/config"
	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess/models"
	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess/request"
)

const (
	// ServiceName - The name of service.
	ServiceName = "WebSessions"
)

//WebSessionsService provides the API operations for making requests to
// WebSessions endpoint.
type WebSessionsService struct {
	*client.Client
}

//New createa a new instance of the WebSessionsService client.
//
// Example:
//   cfg := config.NewConfig().WithUsername("Administrator").WithPassword("2Access").WithEndpoint(paURL)
//
//   //Create a WebSessionsService from the configuration
//   svc := webSessions.New(cfg)
//
func New(cfg *config.Config) *WebSessionsService {

	return &WebSessionsService{Client: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingaccess.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a WebSessions operation
func (s *WebSessionsService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := s.NewRequest(op, params, data)

	return req
}

//GetWebSessionsCommand - Get all WebSessions
//RequestType: GET
//Input: input *GetWebSessionsCommandInput
func (s *WebSessionsService) GetWebSessionsCommand(input *GetWebSessionsCommandInput) (output *models.WebSessionsView, resp *http.Response, err error) {
	path := "/webSessions"
	op := &request.Operation{
		Name:       "GetWebSessionsCommand",
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
	output = &models.WebSessionsView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// GetWebSessionsCommandInput - Inputs for GetWebSessionsCommand
type GetWebSessionsCommandInput struct {
	Page          string
	NumberPerPage string
	Filter        string
	Name          string
	SortKey       string
	Order         string
}

//AddWebSessionCommand - Create a WebSession
//RequestType: POST
//Input: input *AddWebSessionCommandInput
func (s *WebSessionsService) AddWebSessionCommand(input *AddWebSessionCommandInput) (output *models.WebSessionView, resp *http.Response, err error) {
	path := "/webSessions"
	op := &request.Operation{
		Name:        "AddWebSessionCommand",
		HTTPMethod:  "POST",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.WebSessionView{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// AddWebSessionCommandInput - Inputs for AddWebSessionCommand
type AddWebSessionCommandInput struct {
	Body models.WebSessionView
}

//DeleteWebSessionCommand - Delete a WebSession
//RequestType: DELETE
//Input: input *DeleteWebSessionCommandInput
func (s *WebSessionsService) DeleteWebSessionCommand(input *DeleteWebSessionCommandInput) (resp *http.Response, err error) {
	path := "/webSessions/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "DeleteWebSessionCommand",
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

// DeleteWebSessionCommandInput - Inputs for DeleteWebSessionCommand
type DeleteWebSessionCommandInput struct {
	Id string
}

//GetWebSessionCommand - Get a WebSession
//RequestType: GET
//Input: input *GetWebSessionCommandInput
func (s *WebSessionsService) GetWebSessionCommand(input *GetWebSessionCommandInput) (output *models.WebSessionView, resp *http.Response, err error) {
	path := "/webSessions/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "GetWebSessionCommand",
		HTTPMethod:  "GET",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.WebSessionView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// GetWebSessionCommandInput - Inputs for GetWebSessionCommand
type GetWebSessionCommandInput struct {
	Id string
}

//UpdateWebSessionCommand - Update a WebSession
//RequestType: PUT
//Input: input *UpdateWebSessionCommandInput
func (s *WebSessionsService) UpdateWebSessionCommand(input *UpdateWebSessionCommandInput) (output *models.WebSessionView, resp *http.Response, err error) {
	path := "/webSessions/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "UpdateWebSessionCommand",
		HTTPMethod:  "PUT",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.WebSessionView{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// UpdateWebSessionCommandInput - Inputs for UpdateWebSessionCommand
type UpdateWebSessionCommandInput struct {
	Body models.WebSessionView
	Id   string
}
