package pingfederate

import (
	"net/http"

	"github.com/iwarapter/pingaccess-sdk-go/pingaccess"
	"github.com/iwarapter/pingaccess-sdk-go/pingaccess/client"
	"github.com/iwarapter/pingaccess-sdk-go/pingaccess/client/metadata"
	"github.com/iwarapter/pingaccess-sdk-go/pingaccess/config"
	"github.com/iwarapter/pingaccess-sdk-go/pingaccess/models"
	"github.com/iwarapter/pingaccess-sdk-go/pingaccess/request"
)

const (
	// ServiceName - The name of service.
	ServiceName = "Pingfederate"
)

//PingfederateService provides the API operations for making requests to
// Pingfederate endpoint.
type PingfederateService struct {
	*client.Client
}

//New createa a new instance of the PingfederateService client.
//
// Example:
//   cfg := config.NewConfig().WithUsername("Administrator").WithPassword("2Access").WithEndpoint(paURL)
//
//   //Create a PingfederateService from the configuration
//   svc := pingfederate.New(cfg)
//
func New(cfg *config.Config) *PingfederateService {

	return &PingfederateService{Client: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingaccess.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a Pingfederate operation
func (s *PingfederateService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := s.NewRequest(op, params, data)

	return req
}

//DeletePingFederateCommand - Resets the PingFederate configuration to default values
//RequestType: DELETE
//Input:
func (s *PingfederateService) DeletePingFederateCommand() (resp *http.Response, err error) {
	path := "/pingfederate"
	op := &request.Operation{
		Name:       "DeletePingFederateCommand",
		HTTPMethod: "DELETE",
		HTTPPath:   path,
	}

	req := s.newRequest(op, nil, nil)

	if req.Send() == nil {
		return req.HTTPResponse, nil
	}
	return req.HTTPResponse, req.Error
}

//GetPingFederateCommand - Get the PingFederate configuration
//RequestType: GET
//Input:
func (s *PingfederateService) GetPingFederateCommand() (output *models.PingFederateRuntimeView, resp *http.Response, err error) {
	path := "/pingfederate"
	op := &request.Operation{
		Name:       "GetPingFederateCommand",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.PingFederateRuntimeView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdatePingFederateCommand - Update the PingFederate configuration
//RequestType: PUT
//Input: input *UpdatePingFederateCommandInput
func (s *PingfederateService) UpdatePingFederateCommand(input *UpdatePingFederateCommandInput) (output *models.PingFederateRuntimeView, resp *http.Response, err error) {
	path := "/pingfederate"
	op := &request.Operation{
		Name:        "UpdatePingFederateCommand",
		HTTPMethod:  "PUT",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.PingFederateRuntimeView{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// UpdatePingFederateCommandInput - Inputs for UpdatePingFederateCommand
type UpdatePingFederateCommandInput struct {
	Body models.PingFederateRuntimeView
}

//DeletePingFederateAccessTokensCommand - Resets the PingAccess OAuth Client configuration to default values
//RequestType: DELETE
//Input:
func (s *PingfederateService) DeletePingFederateAccessTokensCommand() (resp *http.Response, err error) {
	path := "/pingfederate/accessTokens"
	op := &request.Operation{
		Name:       "DeletePingFederateAccessTokensCommand",
		HTTPMethod: "DELETE",
		HTTPPath:   path,
	}

	req := s.newRequest(op, nil, nil)

	if req.Send() == nil {
		return req.HTTPResponse, nil
	}
	return req.HTTPResponse, req.Error
}

//GetPingFederateAccessTokensCommand - Get the PingAccess OAuth Client configuration
//RequestType: GET
//Input:
func (s *PingfederateService) GetPingFederateAccessTokensCommand() (output *models.PingFederateAccessTokenView, resp *http.Response, err error) {
	path := "/pingfederate/accessTokens"
	op := &request.Operation{
		Name:       "GetPingFederateAccessTokensCommand",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.PingFederateAccessTokenView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdatePingFederateAccessTokensCommand - Update the PingFederate OAuth Client configuration
//RequestType: PUT
//Input: input *UpdatePingFederateAccessTokensCommandInput
func (s *PingfederateService) UpdatePingFederateAccessTokensCommand(input *UpdatePingFederateAccessTokensCommandInput) (output *models.PingFederateAccessTokenView, resp *http.Response, err error) {
	path := "/pingfederate/accessTokens"
	op := &request.Operation{
		Name:        "UpdatePingFederateAccessTokensCommand",
		HTTPMethod:  "PUT",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.PingFederateAccessTokenView{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// UpdatePingFederateAccessTokensCommandInput - Inputs for UpdatePingFederateAccessTokensCommand
type UpdatePingFederateAccessTokensCommandInput struct {
	Body models.PingFederateAccessTokenView
}

//DeletePingFederateAdminCommand - Resets the PingFederate Admin configuration to default values
//RequestType: DELETE
//Input:
func (s *PingfederateService) DeletePingFederateAdminCommand() (resp *http.Response, err error) {
	path := "/pingfederate/admin"
	op := &request.Operation{
		Name:       "DeletePingFederateAdminCommand",
		HTTPMethod: "DELETE",
		HTTPPath:   path,
	}

	req := s.newRequest(op, nil, nil)

	if req.Send() == nil {
		return req.HTTPResponse, nil
	}
	return req.HTTPResponse, req.Error
}

//GetPingFederateAdminCommand - Get the PingFederate Admin configuration
//RequestType: GET
//Input:
func (s *PingfederateService) GetPingFederateAdminCommand() (output *models.PingFederateAdminView, resp *http.Response, err error) {
	path := "/pingfederate/admin"
	op := &request.Operation{
		Name:       "GetPingFederateAdminCommand",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.PingFederateAdminView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdatePingFederateAdminCommand - Update the PingFederate Admin configuration
//RequestType: PUT
//Input: input *UpdatePingFederateAdminCommandInput
func (s *PingfederateService) UpdatePingFederateAdminCommand(input *UpdatePingFederateAdminCommandInput) (output *models.PingFederateAdminView, resp *http.Response, err error) {
	path := "/pingfederate/admin"
	op := &request.Operation{
		Name:        "UpdatePingFederateAdminCommand",
		HTTPMethod:  "PUT",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.PingFederateAdminView{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// UpdatePingFederateAdminCommandInput - Inputs for UpdatePingFederateAdminCommand
type UpdatePingFederateAdminCommandInput struct {
	Body models.PingFederateAdminView
}

//GetPingFederateMetadataCommand - Get the PingFederate metadata
//RequestType: GET
//Input:
func (s *PingfederateService) GetPingFederateMetadataCommand() (output *models.OIDCProviderMetadata, resp *http.Response, err error) {
	path := "/pingfederate/metadata"
	op := &request.Operation{
		Name:       "GetPingFederateMetadataCommand",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.OIDCProviderMetadata{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//DeletePingFederateRuntimeCommand - Resets the PingFederate configuration
//RequestType: DELETE
//Input:
func (s *PingfederateService) DeletePingFederateRuntimeCommand() (resp *http.Response, err error) {
	path := "/pingfederate/runtime"
	op := &request.Operation{
		Name:       "DeletePingFederateRuntimeCommand",
		HTTPMethod: "DELETE",
		HTTPPath:   path,
	}

	req := s.newRequest(op, nil, nil)

	if req.Send() == nil {
		return req.HTTPResponse, nil
	}
	return req.HTTPResponse, req.Error
}

//GetPingFederateRuntimeCommand - Get the PingFederate configuration
//RequestType: GET
//Input:
func (s *PingfederateService) GetPingFederateRuntimeCommand() (output *models.PingFederateMetadataRuntimeView, resp *http.Response, err error) {
	path := "/pingfederate/runtime"
	op := &request.Operation{
		Name:       "GetPingFederateRuntimeCommand",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.PingFederateMetadataRuntimeView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdatePingFederateRuntimeCommand - Update the PingFederate configuration
//RequestType: PUT
//Input: input *UpdatePingFederateRuntimeCommandInput
func (s *PingfederateService) UpdatePingFederateRuntimeCommand(input *UpdatePingFederateRuntimeCommandInput) (output *models.PingFederateMetadataRuntimeView, resp *http.Response, err error) {
	path := "/pingfederate/runtime"
	op := &request.Operation{
		Name:        "UpdatePingFederateRuntimeCommand",
		HTTPMethod:  "PUT",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.PingFederateMetadataRuntimeView{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// UpdatePingFederateRuntimeCommandInput - Inputs for UpdatePingFederateRuntimeCommand
type UpdatePingFederateRuntimeCommandInput struct {
	Body models.PingFederateMetadataRuntimeView
}

//GetPingFederateRuntimeMetadataCommand - Get the PingFederate Runtime metadata
//RequestType: GET
//Input:
func (s *PingfederateService) GetPingFederateRuntimeMetadataCommand() (output *models.OIDCProviderMetadata, resp *http.Response, err error) {
	path := "/pingfederate/runtime/metadata"
	op := &request.Operation{
		Name:       "GetPingFederateRuntimeMetadataCommand",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.OIDCProviderMetadata{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}
