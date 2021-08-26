package auth

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
	ServiceName = "Auth"
)

//AuthService provides the API operations for making requests to
// Auth endpoint.
type AuthService struct {
	*client.Client
}

//New createa a new instance of the AuthService client.
//
// Example:
//   cfg := config.NewConfig().WithUsername("Administrator").WithPassword("2Access").WithEndpoint(paURL)
//
//   //Create a AuthService from the configuration
//   svc := auth.New(cfg)
//
func New(cfg *config.Config) *AuthService {

	return &AuthService{Client: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingaccess.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a Auth operation
func (s *AuthService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := s.NewRequest(op, params, data)

	return req
}

//DeleteBasicAuthCommand - Resets the HTTP Basic Authentication configuration to default values
//RequestType: DELETE
//Input:
func (s *AuthService) DeleteBasicAuthCommand() (resp *http.Response, err error) {
	path := "/auth/basic"
	op := &request.Operation{
		Name:       "DeleteBasicAuthCommand",
		HTTPMethod: "DELETE",
		HTTPPath:   path,
	}

	req := s.newRequest(op, nil, nil)

	if req.Send() == nil {
		return req.HTTPResponse, nil
	}
	return req.HTTPResponse, req.Error
}

//GetBasicAuthCommand - Get the HTTP Basic Authentication configuration
//RequestType: GET
//Input:
func (s *AuthService) GetBasicAuthCommand() (output *models.BasicConfig, resp *http.Response, err error) {
	path := "/auth/basic"
	op := &request.Operation{
		Name:       "GetBasicAuthCommand",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.BasicConfig{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateBasicAuthCommand - Update the Basic Authentication configuration
//RequestType: PUT
//Input: input *UpdateBasicAuthCommandInput
func (s *AuthService) UpdateBasicAuthCommand(input *UpdateBasicAuthCommandInput) (output *models.BasicAuthConfigView, resp *http.Response, err error) {
	path := "/auth/basic"
	op := &request.Operation{
		Name:        "UpdateBasicAuthCommand",
		HTTPMethod:  "PUT",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.BasicAuthConfigView{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// UpdateBasicAuthCommandInput - Inputs for UpdateBasicAuthCommand
type UpdateBasicAuthCommandInput struct {
	Body models.BasicAuthConfigView
}

//DeleteOAuthAuthCommand - Resets the OAuth Authentication configuration to default values
//RequestType: DELETE
//Input:
func (s *AuthService) DeleteOAuthAuthCommand() (resp *http.Response, err error) {
	path := "/auth/oauth"
	op := &request.Operation{
		Name:       "DeleteOAuthAuthCommand",
		HTTPMethod: "DELETE",
		HTTPPath:   path,
	}

	req := s.newRequest(op, nil, nil)

	if req.Send() == nil {
		return req.HTTPResponse, nil
	}
	return req.HTTPResponse, req.Error
}

//GetOAuthAuthCommand - Get the OAuth Authentication configuration
//RequestType: GET
//Input:
func (s *AuthService) GetOAuthAuthCommand() (output *models.OAuthConfigView, resp *http.Response, err error) {
	path := "/auth/oauth"
	op := &request.Operation{
		Name:       "GetOAuthAuthCommand",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.OAuthConfigView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateOAuthAuthCommand - Update the OAuth Authentication configuration
//RequestType: PUT
//Input: input *UpdateOAuthAuthCommandInput
func (s *AuthService) UpdateOAuthAuthCommand(input *UpdateOAuthAuthCommandInput) (output *models.OAuthConfigView, resp *http.Response, err error) {
	path := "/auth/oauth"
	op := &request.Operation{
		Name:        "UpdateOAuthAuthCommand",
		HTTPMethod:  "PUT",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.OAuthConfigView{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// UpdateOAuthAuthCommandInput - Inputs for UpdateOAuthAuthCommand
type UpdateOAuthAuthCommandInput struct {
	Body models.OAuthConfigView
}

//DeleteOidcAuthCommand - Resets the OIDC Authentication configuration to default values
//RequestType: DELETE
//Input:
func (s *AuthService) DeleteOidcAuthCommand() (resp *http.Response, err error) {
	path := "/auth/oidc"
	op := &request.Operation{
		Name:       "DeleteOidcAuthCommand",
		HTTPMethod: "DELETE",
		HTTPPath:   path,
	}

	req := s.newRequest(op, nil, nil)

	if req.Send() == nil {
		return req.HTTPResponse, nil
	}
	return req.HTTPResponse, req.Error
}

//GetOidcAuthCommand - Get the OIDC Authentication configuration
//RequestType: GET
//Input:
func (s *AuthService) GetOidcAuthCommand() (output *models.OidcConfigView, resp *http.Response, err error) {
	path := "/auth/oidc"
	op := &request.Operation{
		Name:       "GetOidcAuthCommand",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.OidcConfigView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateOidcAuthCommand - Update the OIDC Authentication configuration
//RequestType: PUT
//Input: input *UpdateOidcAuthCommandInput
func (s *AuthService) UpdateOidcAuthCommand(input *UpdateOidcAuthCommandInput) (output *models.OidcConfigView, resp *http.Response, err error) {
	path := "/auth/oidc"
	op := &request.Operation{
		Name:        "UpdateOidcAuthCommand",
		HTTPMethod:  "PUT",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.OidcConfigView{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// UpdateOidcAuthCommandInput - Inputs for UpdateOidcAuthCommand
type UpdateOidcAuthCommandInput struct {
	Body models.OidcConfigView
}

//GetAuthOidcScopesCommand - Get the scopes supported by the current Admin OIDC Provider
//RequestType: GET
//Input: input *GetAuthOidcScopesCommandInput
func (s *AuthService) GetAuthOidcScopesCommand(input *GetAuthOidcScopesCommandInput) (output *models.SupportedScopesView, resp *http.Response, err error) {
	path := "/auth/oidc/scopes"
	op := &request.Operation{
		Name:       "GetAuthOidcScopesCommand",
		HTTPMethod: "GET",
		HTTPPath:   path,
		QueryParams: map[string]string{
			"clientId": input.ClientId,
		},
	}
	output = &models.SupportedScopesView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// GetAuthOidcScopesCommandInput - Inputs for GetAuthOidcScopesCommand
type GetAuthOidcScopesCommandInput struct {
	ClientId string
}

//DeleteAdminTokenProviderCommand - Resets the Admin Token Provider configuration to default values
//RequestType: DELETE
//Input:
func (s *AuthService) DeleteAdminTokenProviderCommand() (resp *http.Response, err error) {
	path := "/auth/tokenProvider"
	op := &request.Operation{
		Name:       "DeleteAdminTokenProviderCommand",
		HTTPMethod: "DELETE",
		HTTPPath:   path,
	}

	req := s.newRequest(op, nil, nil)

	if req.Send() == nil {
		return req.HTTPResponse, nil
	}
	return req.HTTPResponse, req.Error
}

//GetAdminTokenProviderCommand - Get the Admin Token Provider configuration
//RequestType: GET
//Input:
func (s *AuthService) GetAdminTokenProviderCommand() (output *models.AdminTokenProviderView, resp *http.Response, err error) {
	path := "/auth/tokenProvider"
	op := &request.Operation{
		Name:       "GetAdminTokenProviderCommand",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.AdminTokenProviderView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateAdminTokenProviderCommand - Update the Admin Token Provider configuration
//RequestType: PUT
//Input: input *UpdateAdminTokenProviderCommandInput
func (s *AuthService) UpdateAdminTokenProviderCommand(input *UpdateAdminTokenProviderCommandInput) (output *models.AdminTokenProviderView, resp *http.Response, err error) {
	path := "/auth/tokenProvider"
	op := &request.Operation{
		Name:        "UpdateAdminTokenProviderCommand",
		HTTPMethod:  "PUT",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.AdminTokenProviderView{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// UpdateAdminTokenProviderCommandInput - Inputs for UpdateAdminTokenProviderCommand
type UpdateAdminTokenProviderCommandInput struct {
	Body models.AdminTokenProviderView
}

//GetAdminTokenProviderMetadataCommand - Get the Admin Token Provider metadata
//RequestType: GET
//Input:
func (s *AuthService) GetAdminTokenProviderMetadataCommand() (output *models.OIDCProviderMetadata, resp *http.Response, err error) {
	path := "/auth/tokenProvider/metadata"
	op := &request.Operation{
		Name:       "GetAdminTokenProviderMetadataCommand",
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

//DeleteAdminBasicWebSessionCommand - Resets the Admin Web Session configuration to default values
//RequestType: DELETE
//Input:
func (s *AuthService) DeleteAdminBasicWebSessionCommand() (resp *http.Response, err error) {
	path := "/auth/webSession"
	op := &request.Operation{
		Name:       "DeleteAdminBasicWebSessionCommand",
		HTTPMethod: "DELETE",
		HTTPPath:   path,
	}

	req := s.newRequest(op, nil, nil)

	if req.Send() == nil {
		return req.HTTPResponse, nil
	}
	return req.HTTPResponse, req.Error
}

//GetAdminBasicWebSessionCommand - Get the admin web session configuration
//RequestType: GET
//Input:
func (s *AuthService) GetAdminBasicWebSessionCommand() (output *models.AdminBasicWebSessionView, resp *http.Response, err error) {
	path := "/auth/webSession"
	op := &request.Operation{
		Name:       "GetAdminBasicWebSessionCommand",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.AdminBasicWebSessionView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateAdminBasicWebSessionCommand - Update the admin web session configuration
//RequestType: PUT
//Input: input *UpdateAdminBasicWebSessionCommandInput
func (s *AuthService) UpdateAdminBasicWebSessionCommand(input *UpdateAdminBasicWebSessionCommandInput) (output *models.AdminBasicWebSessionView, resp *http.Response, err error) {
	path := "/auth/webSession"
	op := &request.Operation{
		Name:        "UpdateAdminBasicWebSessionCommand",
		HTTPMethod:  "PUT",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.AdminBasicWebSessionView{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// UpdateAdminBasicWebSessionCommandInput - Inputs for UpdateAdminBasicWebSessionCommand
type UpdateAdminBasicWebSessionCommandInput struct {
	Body models.AdminBasicWebSessionView
}
