package authTokenManagement

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
	ServiceName = "AuthTokenManagement"
)

//AuthTokenManagementService provides the API operations for making requests to
// AuthTokenManagement endpoint.
type AuthTokenManagementService struct {
	*client.Client
}

//New createa a new instance of the AuthTokenManagementService client.
//
// Example:
//   cfg := config.NewConfig().WithUsername("Administrator").WithPassword("2FederateM0re").WithEndpoint(paURL.String())
//
//   //Create a AuthTokenManagementService from the configuration
//   svc := authTokenManagement.New(cfg)
//
func New(cfg *config.Config) *AuthTokenManagementService {

	return &AuthTokenManagementService{Client: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingaccess.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a AuthTokenManagement operation
func (s *AuthTokenManagementService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := s.NewRequest(op, params, data)

	return req
}

//DeleteAuthTokenManagementCommand - Resets the Auth Token Management configuration to default values
//RequestType: DELETE
//Input:
func (s *AuthTokenManagementService) DeleteAuthTokenManagementCommand() (resp *http.Response, err error) {
	path := "/authTokenManagement"
	op := &request.Operation{
		Name:       "DeleteAuthTokenManagementCommand",
		HTTPMethod: "DELETE",
		HTTPPath:   path,
	}

	req := s.newRequest(op, nil, nil)

	if req.Send() == nil {
		return req.HTTPResponse, nil
	}
	return req.HTTPResponse, req.Error
}

//GetAuthTokenManagementCommand - Get the Auth Token Management configuration
//RequestType: GET
//Input:
func (s *AuthTokenManagementService) GetAuthTokenManagementCommand() (output *models.AuthTokenManagementView, resp *http.Response, err error) {
	path := "/authTokenManagement"
	op := &request.Operation{
		Name:       "GetAuthTokenManagementCommand",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.AuthTokenManagementView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateAuthTokenManagementCommand - Update the Auth Token Management configuration
//RequestType: PUT
//Input: input *UpdateAuthTokenManagementCommandInput
func (s *AuthTokenManagementService) UpdateAuthTokenManagementCommand(input *UpdateAuthTokenManagementCommandInput) (output *models.AuthTokenManagementView, resp *http.Response, err error) {
	path := "/authTokenManagement"
	op := &request.Operation{
		Name:        "UpdateAuthTokenManagementCommand",
		HTTPMethod:  "PUT",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.AuthTokenManagementView{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// UpdateAuthTokenManagementCommandInput - Inputs for UpdateAuthTokenManagementCommand
type UpdateAuthTokenManagementCommandInput struct {
	Body models.AuthTokenManagementView
}

//GetAuthTokenKeySetCommand - Get the Auth Token key set configuration
//RequestType: GET
//Input:
func (s *AuthTokenManagementService) GetAuthTokenKeySetCommand() (output *models.KeySetView, resp *http.Response, err error) {
	path := "/authTokenManagement/keySet"
	op := &request.Operation{
		Name:       "GetAuthTokenKeySetCommand",
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

//UpdateAuthTokenKeySetCommand - Update the AuthToken key set configuration
//RequestType: PUT
//Input: input *UpdateAuthTokenKeySetCommandInput
func (s *AuthTokenManagementService) UpdateAuthTokenKeySetCommand(input *UpdateAuthTokenKeySetCommandInput) (output *models.KeySetView, resp *http.Response, err error) {
	path := "/authTokenManagement/keySet"
	op := &request.Operation{
		Name:        "UpdateAuthTokenKeySetCommand",
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

// UpdateAuthTokenKeySetCommandInput - Inputs for UpdateAuthTokenKeySetCommand
type UpdateAuthTokenKeySetCommandInput struct {
	Body models.KeySetView
}
