package rejectionHandlers

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
	ServiceName = "RejectionHandlers"
)

//RejectionHandlersService provides the API operations for making requests to
// RejectionHandlers endpoint.
type RejectionHandlersService struct {
	*client.Client
}

//New createa a new instance of the RejectionHandlersService client.
//
// Example:
//   cfg := config.NewConfig().WithUsername("Administrator").WithPassword("2Access").WithEndpoint(paURL)
//
//   //Create a RejectionHandlersService from the configuration
//   svc := rejectionHandlers.New(cfg)
//
func New(cfg *config.Config) *RejectionHandlersService {

	return &RejectionHandlersService{Client: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingaccess.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a RejectionHandlers operation
func (s *RejectionHandlersService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := s.NewRequest(op, params, data)

	return req
}

//GetRejectionHandlersCommand - Get all Rejection Handlers
//RequestType: GET
//Input: input *GetRejectionHandlersCommandInput
func (s *RejectionHandlersService) GetRejectionHandlersCommand(input *GetRejectionHandlersCommandInput) (output *models.RejectionHandlersView, resp *http.Response, err error) {
	path := "/rejectionHandlers"
	op := &request.Operation{
		Name:       "GetRejectionHandlersCommand",
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
	output = &models.RejectionHandlersView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// GetRejectionHandlersCommandInput - Inputs for GetRejectionHandlersCommand
type GetRejectionHandlersCommandInput struct {
	Page          string
	NumberPerPage string
	Filter        string
	Name          string
	SortKey       string
	Order         string
}

//AddRejectionHandlerCommand - Create a Rejection Handler
//RequestType: POST
//Input: input *AddRejectionHandlerCommandInput
func (s *RejectionHandlersService) AddRejectionHandlerCommand(input *AddRejectionHandlerCommandInput) (output *models.RejectionHandlerView, resp *http.Response, err error) {
	path := "/rejectionHandlers"
	op := &request.Operation{
		Name:        "AddRejectionHandlerCommand",
		HTTPMethod:  "POST",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.RejectionHandlerView{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// AddRejectionHandlerCommandInput - Inputs for AddRejectionHandlerCommand
type AddRejectionHandlerCommandInput struct {
	Body models.RejectionHandlerView
}

//GetRejectionHandlerDescriptorsCommand - Get descriptors for all supported Rejection Handler types
//RequestType: GET
//Input:
func (s *RejectionHandlersService) GetRejectionHandlerDescriptorsCommand() (output *models.DescriptorsView, resp *http.Response, err error) {
	path := "/rejectionHandlers/descriptors"
	op := &request.Operation{
		Name:       "GetRejectionHandlerDescriptorsCommand",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.DescriptorsView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetRejecitonHandlerDescriptorCommand - Get descriptor for a Rejection Handler type
//RequestType: GET
//Input: input *GetRejecitonHandlerDescriptorCommandInput
func (s *RejectionHandlersService) GetRejecitonHandlerDescriptorCommand(input *GetRejecitonHandlerDescriptorCommandInput) (output *models.DescriptorView, resp *http.Response, err error) {
	path := "/rejectionHandlers/descriptors/{rejectionHandlerType}"
	path = strings.Replace(path, "{rejectionHandlerType}", input.RejectionHandlerType, -1)

	op := &request.Operation{
		Name:        "GetRejecitonHandlerDescriptorCommand",
		HTTPMethod:  "GET",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.DescriptorView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// GetRejecitonHandlerDescriptorCommandInput - Inputs for GetRejecitonHandlerDescriptorCommand
type GetRejecitonHandlerDescriptorCommandInput struct {
	RejectionHandlerType string
}

//DeleteRejectionHandlerCommand - Delete a Rejection Handler
//RequestType: DELETE
//Input: input *DeleteRejectionHandlerCommandInput
func (s *RejectionHandlersService) DeleteRejectionHandlerCommand(input *DeleteRejectionHandlerCommandInput) (resp *http.Response, err error) {
	path := "/rejectionHandlers/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "DeleteRejectionHandlerCommand",
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

// DeleteRejectionHandlerCommandInput - Inputs for DeleteRejectionHandlerCommand
type DeleteRejectionHandlerCommandInput struct {
	Id string
}

//GetRejectionHandlerCommand - Get a Rejection Handler
//RequestType: GET
//Input: input *GetRejectionHandlerCommandInput
func (s *RejectionHandlersService) GetRejectionHandlerCommand(input *GetRejectionHandlerCommandInput) (output *models.RejectionHandlerView, resp *http.Response, err error) {
	path := "/rejectionHandlers/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "GetRejectionHandlerCommand",
		HTTPMethod:  "GET",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.RejectionHandlerView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// GetRejectionHandlerCommandInput - Inputs for GetRejectionHandlerCommand
type GetRejectionHandlerCommandInput struct {
	Id string
}

//UpdateRejectionHandlerCommand - Update a Rejection Handler
//RequestType: PUT
//Input: input *UpdateRejectionHandlerCommandInput
func (s *RejectionHandlersService) UpdateRejectionHandlerCommand(input *UpdateRejectionHandlerCommandInput) (output *models.RejectionHandlerView, resp *http.Response, err error) {
	path := "/rejectionHandlers/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "UpdateRejectionHandlerCommand",
		HTTPMethod:  "PUT",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.RejectionHandlerView{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// UpdateRejectionHandlerCommandInput - Inputs for UpdateRejectionHandlerCommand
type UpdateRejectionHandlerCommandInput struct {
	Body models.RejectionHandlerView
	Id   string
}
