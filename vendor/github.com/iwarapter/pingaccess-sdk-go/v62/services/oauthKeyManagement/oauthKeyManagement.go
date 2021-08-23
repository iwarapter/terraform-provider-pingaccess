package oauthKeyManagement

import (
	"net/http"

	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess"
	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess/client"
	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess/client/metadata"
	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess/config"
	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess/models"
	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess/request"
)

const (
	// ServiceName - The name of service.
	ServiceName = "OauthKeyManagement"
)

//OauthKeyManagementService provides the API operations for making requests to
// OauthKeyManagement endpoint.
type OauthKeyManagementService struct {
	*client.Client
}

//New createa a new instance of the OauthKeyManagementService client.
//
// Example:
//   cfg := config.NewConfig().WithUsername("Administrator").WithPassword("2Access").WithEndpoint(paURL)
//
//   //Create a OauthKeyManagementService from the configuration
//   svc := oauthKeyManagement.New(cfg)
//
func New(cfg *config.Config) *OauthKeyManagementService {

	return &OauthKeyManagementService{Client: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingaccess.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a OauthKeyManagement operation
func (s *OauthKeyManagementService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := s.NewRequest(op, params, data)

	return req
}

//DeleteOAuthKeyManagementCommand - Resets the OAuth Key Management configuration to default values
//RequestType: DELETE
//Input:
func (s *OauthKeyManagementService) DeleteOAuthKeyManagementCommand() (resp *http.Response, err error) {
	path := "/oauthKeyManagement"
	op := &request.Operation{
		Name:       "DeleteOAuthKeyManagementCommand",
		HTTPMethod: "DELETE",
		HTTPPath:   path,
	}

	req := s.newRequest(op, nil, nil)

	if req.Send() == nil {
		return req.HTTPResponse, nil
	}
	return req.HTTPResponse, req.Error
}

//GetOAuthKeyManagementCommand - Get the OAuth Key Management configuration
//RequestType: GET
//Input:
func (s *OauthKeyManagementService) GetOAuthKeyManagementCommand() (output *models.OAuthKeyManagementView, resp *http.Response, err error) {
	path := "/oauthKeyManagement"
	op := &request.Operation{
		Name:       "GetOAuthKeyManagementCommand",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.OAuthKeyManagementView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateOAuthKeyManagementCommand - Update the OAuth Key Management configuration
//RequestType: PUT
//Input: input *UpdateOAuthKeyManagementCommandInput
func (s *OauthKeyManagementService) UpdateOAuthKeyManagementCommand(input *UpdateOAuthKeyManagementCommandInput) (output *models.OAuthKeyManagementView, resp *http.Response, err error) {
	path := "/oauthKeyManagement"
	op := &request.Operation{
		Name:        "UpdateOAuthKeyManagementCommand",
		HTTPMethod:  "PUT",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.OAuthKeyManagementView{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// UpdateOAuthKeyManagementCommandInput - Inputs for UpdateOAuthKeyManagementCommand
type UpdateOAuthKeyManagementCommandInput struct {
	Body models.OAuthKeyManagementView
}

//GetOAuthKeySetCommand - Get the OAuth key set configuration
//RequestType: GET
//Input:
func (s *OauthKeyManagementService) GetOAuthKeySetCommand() (output *models.KeySetView, resp *http.Response, err error) {
	path := "/oauthKeyManagement/keySet"
	op := &request.Operation{
		Name:       "GetOAuthKeySetCommand",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.KeySetView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateOAuthKeySetCommand - Update the OAuth key set configuration
//RequestType: PUT
//Input: input *UpdateOAuthKeySetCommandInput
func (s *OauthKeyManagementService) UpdateOAuthKeySetCommand(input *UpdateOAuthKeySetCommandInput) (output *models.KeySetView, resp *http.Response, err error) {
	path := "/oauthKeyManagement/keySet"
	op := &request.Operation{
		Name:        "UpdateOAuthKeySetCommand",
		HTTPMethod:  "PUT",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.KeySetView{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// UpdateOAuthKeySetCommandInput - Inputs for UpdateOAuthKeySetCommand
type UpdateOAuthKeySetCommandInput struct {
	Body models.KeySetView
}
