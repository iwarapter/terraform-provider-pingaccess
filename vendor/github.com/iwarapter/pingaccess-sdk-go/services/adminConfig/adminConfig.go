package adminConfig

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
	ServiceName = "AdminConfig"
)

//AdminConfigService provides the API operations for making requests to
// AdminConfig endpoint.
type AdminConfigService struct {
	*client.Client
}

//New createa a new instance of the AdminConfigService client.
//
// Example:
//   cfg := config.NewConfig().WithUsername("Administrator").WithPassword("2FederateM0re").WithEndpoint(paURL.String())
//
//   //Create a AdminConfigService from the configuration
//   svc := adminConfig.New(cfg)
//
func New(cfg *config.Config) *AdminConfigService {

	return &AdminConfigService{Client: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingaccess.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a AdminConfig operation
func (s *AdminConfigService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := s.NewRequest(op, params, data)

	return req
}

//DeleteAdminConfigurationCommand - Resets the Admin Config to default values
//RequestType: DELETE
//Input:
func (s *AdminConfigService) DeleteAdminConfigurationCommand() (resp *http.Response, err error) {
	path := "/adminConfig"
	op := &request.Operation{
		Name:       "DeleteAdminConfigurationCommand",
		HTTPMethod: "DELETE",
		HTTPPath:   path,
	}

	req := s.newRequest(op, nil, nil)

	if req.Send() == nil {
		return req.HTTPResponse, nil
	}
	return req.HTTPResponse, req.Error
}

//GetAdminConfigurationCommand - Get the Admin Config
//RequestType: GET
//Input:
func (s *AdminConfigService) GetAdminConfigurationCommand() (output *models.AdminConfigurationView, resp *http.Response, err error) {
	path := "/adminConfig"
	op := &request.Operation{
		Name:       "GetAdminConfigurationCommand",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.AdminConfigurationView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateAdminConfigurationCommand - Update the Admin Config
//RequestType: PUT
//Input: input *UpdateAdminConfigurationCommandInput
func (s *AdminConfigService) UpdateAdminConfigurationCommand(input *UpdateAdminConfigurationCommandInput) (output *models.AdminConfigurationView, resp *http.Response, err error) {
	path := "/adminConfig"
	op := &request.Operation{
		Name:        "UpdateAdminConfigurationCommand",
		HTTPMethod:  "PUT",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.AdminConfigurationView{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// UpdateAdminConfigurationCommandInput - Inputs for UpdateAdminConfigurationCommand
type UpdateAdminConfigurationCommandInput struct {
	Body models.AdminConfigurationView
}

//GetReplicaAdminsCommand - Get list of ReplicaAdmins
//RequestType: GET
//Input:
func (s *AdminConfigService) GetReplicaAdminsCommand() (output *models.ReplicaAdminsView, resp *http.Response, err error) {
	path := "/adminConfig/replicaAdmins"
	op := &request.Operation{
		Name:       "GetReplicaAdminsCommand",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.ReplicaAdminsView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//AddReplicaAdminCommand - Add a ReplicaAdmin
//RequestType: POST
//Input: input *AddReplicaAdminCommandInput
func (s *AdminConfigService) AddReplicaAdminCommand(input *AddReplicaAdminCommandInput) (output *models.ReplicaAdminView, resp *http.Response, err error) {
	path := "/adminConfig/replicaAdmins"
	op := &request.Operation{
		Name:        "AddReplicaAdminCommand",
		HTTPMethod:  "POST",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.ReplicaAdminView{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// AddReplicaAdminCommandInput - Inputs for AddReplicaAdminCommand
type AddReplicaAdminCommandInput struct {
	Body models.ReplicaAdminView
}

//DeleteReplicaAdminCommand - Delete a ReplicaAdmin
//RequestType: DELETE
//Input: input *DeleteReplicaAdminCommandInput
func (s *AdminConfigService) DeleteReplicaAdminCommand(input *DeleteReplicaAdminCommandInput) (resp *http.Response, err error) {
	path := "/adminConfig/replicaAdmins/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "DeleteReplicaAdminCommand",
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

// DeleteReplicaAdminCommandInput - Inputs for DeleteReplicaAdminCommand
type DeleteReplicaAdminCommandInput struct {
	Id string
}

//GetReplicaAdminCommand - Get a ReplicaAdmin
//RequestType: GET
//Input: input *GetReplicaAdminCommandInput
func (s *AdminConfigService) GetReplicaAdminCommand(input *GetReplicaAdminCommandInput) (output *models.ReplicaAdminView, resp *http.Response, err error) {
	path := "/adminConfig/replicaAdmins/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "GetReplicaAdminCommand",
		HTTPMethod:  "GET",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.ReplicaAdminView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// GetReplicaAdminCommandInput - Inputs for GetReplicaAdminCommand
type GetReplicaAdminCommandInput struct {
	Id string
}

//UpdateAdminReplicaCommand - Update a ReplicaAdmin
//RequestType: PUT
//Input: input *UpdateAdminReplicaCommandInput
func (s *AdminConfigService) UpdateAdminReplicaCommand(input *UpdateAdminReplicaCommandInput) (output *models.ReplicaAdminView, resp *http.Response, err error) {
	path := "/adminConfig/replicaAdmins/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "UpdateAdminReplicaCommand",
		HTTPMethod:  "PUT",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.ReplicaAdminView{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// UpdateAdminReplicaCommandInput - Inputs for UpdateAdminReplicaCommand
type UpdateAdminReplicaCommandInput struct {
	Body models.ReplicaAdminView
	Id   string
}

//GetAdminReplicaFileCommand - Get configuration file for a given ReplicaAdmin
//RequestType: POST
//Input: input *GetAdminReplicaFileCommandInput
func (s *AdminConfigService) GetAdminReplicaFileCommand(input *GetAdminReplicaFileCommandInput) (resp *http.Response, err error) {
	path := "/adminConfig/replicaAdmins/{id}/config"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "GetAdminReplicaFileCommand",
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

// GetAdminReplicaFileCommandInput - Inputs for GetAdminReplicaFileCommand
type GetAdminReplicaFileCommandInput struct {
	Id string
}
