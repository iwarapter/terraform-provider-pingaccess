package authnReqLists

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
	ServiceName = "AuthnReqLists"
)

//AuthnReqListsService provides the API operations for making requests to
// AuthnReqLists endpoint.
type AuthnReqListsService struct {
	*client.Client
}

//New createa a new instance of the AuthnReqListsService client.
//
// Example:
//   cfg := config.NewConfig().WithUsername("Administrator").WithPassword("2FederateM0re").WithEndpoint(paURL.String())
//
//   //Create a AuthnReqListsService from the configuration
//   svc := authnReqLists.New(cfg)
//
func New(cfg *config.Config) *AuthnReqListsService {

	return &AuthnReqListsService{Client: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingaccess.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a AuthnReqLists operation
func (s *AuthnReqListsService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := s.NewRequest(op, params, data)

	return req
}

//GetAuthnReqListsCommand - Get all Authentication Requirement Lists
//RequestType: GET
//Input: input *GetAuthnReqListsCommandInput
func (s *AuthnReqListsService) GetAuthnReqListsCommand(input *GetAuthnReqListsCommandInput) (output *models.AuthnReqListsView, resp *http.Response, err error) {
	path := "/authnReqLists"
	op := &request.Operation{
		Name:       "GetAuthnReqListsCommand",
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
	output = &models.AuthnReqListsView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// GetAuthnReqListsCommandInput - Inputs for GetAuthnReqListsCommand
type GetAuthnReqListsCommandInput struct {
	Page          string
	NumberPerPage string
	Filter        string
	Name          string
	SortKey       string
	Order         string
}

//AddAuthnReqListCommand - Add an Authentication Requirement List
//RequestType: POST
//Input: input *AddAuthnReqListCommandInput
func (s *AuthnReqListsService) AddAuthnReqListCommand(input *AddAuthnReqListCommandInput) (output *models.AuthnReqListView, resp *http.Response, err error) {
	path := "/authnReqLists"
	op := &request.Operation{
		Name:        "AddAuthnReqListCommand",
		HTTPMethod:  "POST",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.AuthnReqListView{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// AddAuthnReqListCommandInput - Inputs for AddAuthnReqListCommand
type AddAuthnReqListCommandInput struct {
	Body models.AuthnReqListView
}

//DeleteAuthnReqListCommand - Delete an Authentication Requirement List
//RequestType: DELETE
//Input: input *DeleteAuthnReqListCommandInput
func (s *AuthnReqListsService) DeleteAuthnReqListCommand(input *DeleteAuthnReqListCommandInput) (resp *http.Response, err error) {
	path := "/authnReqLists/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "DeleteAuthnReqListCommand",
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

// DeleteAuthnReqListCommandInput - Inputs for DeleteAuthnReqListCommand
type DeleteAuthnReqListCommandInput struct {
	Id string
}

//GetAuthnReqListCommand - Get an Authentication Requirement List
//RequestType: GET
//Input: input *GetAuthnReqListCommandInput
func (s *AuthnReqListsService) GetAuthnReqListCommand(input *GetAuthnReqListCommandInput) (output *models.AuthnReqListView, resp *http.Response, err error) {
	path := "/authnReqLists/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "GetAuthnReqListCommand",
		HTTPMethod:  "GET",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.AuthnReqListView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// GetAuthnReqListCommandInput - Inputs for GetAuthnReqListCommand
type GetAuthnReqListCommandInput struct {
	Id string
}

//UpdateAuthnReqListCommand - Update an Authentication Requirement List
//RequestType: PUT
//Input: input *UpdateAuthnReqListCommandInput
func (s *AuthnReqListsService) UpdateAuthnReqListCommand(input *UpdateAuthnReqListCommandInput) (output *models.AuthnReqListView, resp *http.Response, err error) {
	path := "/authnReqLists/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "UpdateAuthnReqListCommand",
		HTTPMethod:  "PUT",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.AuthnReqListView{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// UpdateAuthnReqListCommandInput - Inputs for UpdateAuthnReqListCommand
type UpdateAuthnReqListCommandInput struct {
	Body models.AuthnReqListView
	Id   string
}
