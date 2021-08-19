package engines

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
	ServiceName = "Engines"
)

//EnginesService provides the API operations for making requests to
// Engines endpoint.
type EnginesService struct {
	*client.Client
}

//New createa a new instance of the EnginesService client.
//
// Example:
//   cfg := config.NewConfig().WithUsername("Administrator").WithPassword("2Access").WithEndpoint(paURL)
//
//   //Create a EnginesService from the configuration
//   svc := engines.New(cfg)
//
func New(cfg *config.Config) *EnginesService {

	return &EnginesService{Client: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingaccess.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a Engines operation
func (s *EnginesService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := s.NewRequest(op, params, data)

	return req
}

//GetEnginesCommand - Get all Engines
//RequestType: GET
//Input: input *GetEnginesCommandInput
func (s *EnginesService) GetEnginesCommand(input *GetEnginesCommandInput) (output *models.EnginesView, resp *http.Response, err error) {
	path := "/engines"
	op := &request.Operation{
		Name:       "GetEnginesCommand",
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
	output = &models.EnginesView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// GetEnginesCommandInput - Inputs for GetEnginesCommand
type GetEnginesCommandInput struct {
	Page          string
	NumberPerPage string
	Filter        string
	Name          string
	SortKey       string
	Order         string
}

//AddEngineCommand - Add an Engine
//RequestType: POST
//Input: input *AddEngineCommandInput
func (s *EnginesService) AddEngineCommand(input *AddEngineCommandInput) (output *models.EngineView, resp *http.Response, err error) {
	path := "/engines"
	op := &request.Operation{
		Name:        "AddEngineCommand",
		HTTPMethod:  "POST",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.EngineView{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// AddEngineCommandInput - Inputs for AddEngineCommand
type AddEngineCommandInput struct {
	Body models.EngineView
}

//GetEngineCertificatesCommand - Get all Engine Certificates
//RequestType: GET
//Input: input *GetEngineCertificatesCommandInput
func (s *EnginesService) GetEngineCertificatesCommand(input *GetEngineCertificatesCommandInput) (output *models.EngineCertificateView, resp *http.Response, err error) {
	path := "/engines/certificates"
	op := &request.Operation{
		Name:       "GetEngineCertificatesCommand",
		HTTPMethod: "GET",
		HTTPPath:   path,
		QueryParams: map[string]string{
			"page":          input.Page,
			"numberPerPage": input.NumberPerPage,
			"filter":        input.Filter,
			"alias":         input.Alias,
			"sortKey":       input.SortKey,
			"order":         input.Order,
		},
	}
	output = &models.EngineCertificateView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// GetEngineCertificatesCommandInput - Inputs for GetEngineCertificatesCommand
type GetEngineCertificatesCommandInput struct {
	Page          string
	NumberPerPage string
	Filter        string
	Alias         string
	SortKey       string
	Order         string
}

//GetEngineCertificateCommand - Get an Engine Certificate
//RequestType: GET
//Input: input *GetEngineCertificateCommandInput
func (s *EnginesService) GetEngineCertificateCommand(input *GetEngineCertificateCommandInput) (output *models.EngineCertificateView, resp *http.Response, err error) {
	path := "/engines/certificates/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "GetEngineCertificateCommand",
		HTTPMethod:  "GET",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.EngineCertificateView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// GetEngineCertificateCommandInput - Inputs for GetEngineCertificateCommand
type GetEngineCertificateCommandInput struct {
	Id string
}

//GetEngineStatusCommand - Get health status of all Engines
//RequestType: GET
//Input:
func (s *EnginesService) GetEngineStatusCommand() (output *models.EngineHealthStatusView, resp *http.Response, err error) {
	path := "/engines/status"
	op := &request.Operation{
		Name:       "GetEngineStatusCommand",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.EngineHealthStatusView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//DeleteEngineCommand - Delete an Engine
//RequestType: DELETE
//Input: input *DeleteEngineCommandInput
func (s *EnginesService) DeleteEngineCommand(input *DeleteEngineCommandInput) (resp *http.Response, err error) {
	path := "/engines/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "DeleteEngineCommand",
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

// DeleteEngineCommandInput - Inputs for DeleteEngineCommand
type DeleteEngineCommandInput struct {
	Id string
}

//GetEngineCommand - Get an Engine
//RequestType: GET
//Input: input *GetEngineCommandInput
func (s *EnginesService) GetEngineCommand(input *GetEngineCommandInput) (output *models.EngineView, resp *http.Response, err error) {
	path := "/engines/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "GetEngineCommand",
		HTTPMethod:  "GET",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.EngineView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// GetEngineCommandInput - Inputs for GetEngineCommand
type GetEngineCommandInput struct {
	Id string
}

//UpdateEngineCommand - Update an Engine
//RequestType: PUT
//Input: input *UpdateEngineCommandInput
func (s *EnginesService) UpdateEngineCommand(input *UpdateEngineCommandInput) (output *models.EngineView, resp *http.Response, err error) {
	path := "/engines/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "UpdateEngineCommand",
		HTTPMethod:  "PUT",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.EngineView{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// UpdateEngineCommandInput - Inputs for UpdateEngineCommand
type UpdateEngineCommandInput struct {
	Body models.EngineView
	Id   string
}

//GetEngineConfigFileCommand - Get configuration file for an Engine
//RequestType: POST
//Input: input *GetEngineConfigFileCommandInput
func (s *EnginesService) GetEngineConfigFileCommand(input *GetEngineConfigFileCommandInput) (resp *http.Response, err error) {
	path := "/engines/{id}/config"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "GetEngineConfigFileCommand",
		HTTPMethod:  "POST",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}

	req := s.newRequest(op, nil, nil)

	if req.Send() == nil {
		return req.HTTPResponse, nil
	}
	return req.HTTPResponse, req.Error
}

// GetEngineConfigFileCommandInput - Inputs for GetEngineConfigFileCommand
type GetEngineConfigFileCommandInput struct {
	Id string
}
