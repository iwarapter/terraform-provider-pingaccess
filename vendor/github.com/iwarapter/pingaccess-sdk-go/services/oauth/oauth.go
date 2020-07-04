package oauth

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
	ServiceName = "Oauth"
)

//OauthService provides the API operations for making requests to
// Oauth endpoint.
type OauthService struct {
	*client.Client
}

//New createa a new instance of the OauthService client.
//
// Example:
//   cfg := config.NewConfig().WithUsername("Administrator").WithPassword("2FederateM0re").WithEndpoint(paURL.String())
//
//   //Create a OauthService from the configuration
//   svc := oauth.New(cfg)
//
func New(cfg *config.Config) *OauthService {

	return &OauthService{Client: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingaccess.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a Oauth operation
func (s *OauthService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := s.NewRequest(op, params, data)

	return req
}

//DeleteAuthorizationServerCommand - Resets the OpenID Connect Provider configuration to default values
//RequestType: DELETE
//Input:
func (s *OauthService) DeleteAuthorizationServerCommand() (resp *http.Response, err error) {
	path := "/oauth/authServer"
	op := &request.Operation{
		Name:       "DeleteAuthorizationServerCommand",
		HTTPMethod: "DELETE",
		HTTPPath:   path,
	}

	req := s.newRequest(op, nil, nil)

	if req.Send() == nil {
		return req.HTTPResponse, nil
	}
	return req.HTTPResponse, req.Error
}

//GetAuthorizationServerCommand - Get Authorization Server configuration
//RequestType: GET
//Input:
func (s *OauthService) GetAuthorizationServerCommand() (output *models.AuthorizationServerView, resp *http.Response, err error) {
	path := "/oauth/authServer"
	op := &request.Operation{
		Name:       "GetAuthorizationServerCommand",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.AuthorizationServerView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateAuthorizationServerCommand - Update OAuth 2.0 Authorization Server configuration
//RequestType: PUT
//Input: input *UpdateAuthorizationServerCommandInput
func (s *OauthService) UpdateAuthorizationServerCommand(input *UpdateAuthorizationServerCommandInput) (output *models.AuthorizationServerView, resp *http.Response, err error) {
	path := "/oauth/authServer"
	op := &request.Operation{
		Name:        "UpdateAuthorizationServerCommand",
		HTTPMethod:  "PUT",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.AuthorizationServerView{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// UpdateAuthorizationServerCommandInput - Inputs for UpdateAuthorizationServerCommand
type UpdateAuthorizationServerCommandInput struct {
	Body models.AuthorizationServerView
}
