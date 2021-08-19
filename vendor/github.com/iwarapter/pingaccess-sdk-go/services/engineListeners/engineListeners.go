package engineListeners

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
	ServiceName = "EngineListeners"
)

//EngineListenersService provides the API operations for making requests to
// EngineListeners endpoint.
type EngineListenersService struct {
	*client.Client
}

//New createa a new instance of the EngineListenersService client.
//
// Example:
//   cfg := config.NewConfig().WithUsername("Administrator").WithPassword("2Access").WithEndpoint(paURL)
//
//   //Create a EngineListenersService from the configuration
//   svc := engineListeners.New(cfg)
//
func New(cfg *config.Config) *EngineListenersService {

	return &EngineListenersService{Client: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingaccess.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a EngineListeners operation
func (s *EngineListenersService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := s.NewRequest(op, params, data)

	return req
}

//GetEngineListenersCommand - Get all Engine Listeners
//RequestType: GET
//Input: input *GetEngineListenersCommandInput
func (s *EngineListenersService) GetEngineListenersCommand(input *GetEngineListenersCommandInput) (output *models.EngineListenersView, resp *http.Response, err error) {
	path := "/engineListeners"
	op := &request.Operation{
		Name:       "GetEngineListenersCommand",
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
	output = &models.EngineListenersView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// GetEngineListenersCommandInput - Inputs for GetEngineListenersCommand
type GetEngineListenersCommandInput struct {
	Page          string
	NumberPerPage string
	Filter        string
	Name          string
	SortKey       string
	Order         string
}

//AddEngineListenerCommand - Create an Engine Listener
//RequestType: POST
//Input: input *AddEngineListenerCommandInput
func (s *EngineListenersService) AddEngineListenerCommand(input *AddEngineListenerCommandInput) (output *models.EngineListenerView, resp *http.Response, err error) {
	path := "/engineListeners"
	op := &request.Operation{
		Name:        "AddEngineListenerCommand",
		HTTPMethod:  "POST",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.EngineListenerView{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// AddEngineListenerCommandInput - Inputs for AddEngineListenerCommand
type AddEngineListenerCommandInput struct {
	Body models.EngineListenerView
}

//DeleteEngineListenerCommand - Delete an Engine Listener
//RequestType: DELETE
//Input: input *DeleteEngineListenerCommandInput
func (s *EngineListenersService) DeleteEngineListenerCommand(input *DeleteEngineListenerCommandInput) (resp *http.Response, err error) {
	path := "/engineListeners/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "DeleteEngineListenerCommand",
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

// DeleteEngineListenerCommandInput - Inputs for DeleteEngineListenerCommand
type DeleteEngineListenerCommandInput struct {
	Id string
}

//GetEngineListenerCommand - Get an Engine Listener
//RequestType: GET
//Input: input *GetEngineListenerCommandInput
func (s *EngineListenersService) GetEngineListenerCommand(input *GetEngineListenerCommandInput) (output *models.EngineListenerView, resp *http.Response, err error) {
	path := "/engineListeners/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "GetEngineListenerCommand",
		HTTPMethod:  "GET",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.EngineListenerView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// GetEngineListenerCommandInput - Inputs for GetEngineListenerCommand
type GetEngineListenerCommandInput struct {
	Id string
}

//UpdateEngineListenerCommand - Update an Engine Listener
//RequestType: PUT
//Input: input *UpdateEngineListenerCommandInput
func (s *EngineListenersService) UpdateEngineListenerCommand(input *UpdateEngineListenerCommandInput) (output *models.EngineListenerView, resp *http.Response, err error) {
	path := "/engineListeners/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "UpdateEngineListenerCommand",
		HTTPMethod:  "PUT",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.EngineListenerView{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// UpdateEngineListenerCommandInput - Inputs for UpdateEngineListenerCommand
type UpdateEngineListenerCommandInput struct {
	Body models.EngineListenerView
	Id   string
}
