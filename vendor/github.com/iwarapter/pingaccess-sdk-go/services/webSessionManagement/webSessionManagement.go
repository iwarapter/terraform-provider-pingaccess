package webSessionManagement

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
	ServiceName = "WebSessionManagement"
)

//WebSessionManagementService provides the API operations for making requests to
// WebSessionManagement endpoint.
type WebSessionManagementService struct {
	*client.Client
}

//New createa a new instance of the WebSessionManagementService client.
//
// Example:
//   cfg := config.NewConfig().WithUsername("Administrator").WithPassword("2Access").WithEndpoint(paURL)
//
//   //Create a WebSessionManagementService from the configuration
//   svc := webSessionManagement.New(cfg)
//
func New(cfg *config.Config) *WebSessionManagementService {

	return &WebSessionManagementService{Client: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingaccess.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a WebSessionManagement operation
func (s *WebSessionManagementService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := s.NewRequest(op, params, data)

	return req
}

//DeleteWebSessionManagementCommand - Resets the Web Session Management configuration to default values
//RequestType: DELETE
//Input:
func (s *WebSessionManagementService) DeleteWebSessionManagementCommand() (resp *http.Response, err error) {
	path := "/webSessionManagement"
	op := &request.Operation{
		Name:       "DeleteWebSessionManagementCommand",
		HTTPMethod: "DELETE",
		HTTPPath:   path,
	}

	req := s.newRequest(op, nil, nil)

	if req.Send() == nil {
		return req.HTTPResponse, nil
	}
	return req.HTTPResponse, req.Error
}

//GetWebSessionManagementCommand - Get the Web Session Management configuration
//RequestType: GET
//Input:
func (s *WebSessionManagementService) GetWebSessionManagementCommand() (output *models.WebSessionManagementView, resp *http.Response, err error) {
	path := "/webSessionManagement"
	op := &request.Operation{
		Name:       "GetWebSessionManagementCommand",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.WebSessionManagementView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateWebSessionManagementCommand - Update the Web Session Management configuration
//RequestType: PUT
//Input: input *UpdateWebSessionManagementCommandInput
func (s *WebSessionManagementService) UpdateWebSessionManagementCommand(input *UpdateWebSessionManagementCommandInput) (output *models.WebSessionManagementView, resp *http.Response, err error) {
	path := "/webSessionManagement"
	op := &request.Operation{
		Name:        "UpdateWebSessionManagementCommand",
		HTTPMethod:  "PUT",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.WebSessionManagementView{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// UpdateWebSessionManagementCommandInput - Inputs for UpdateWebSessionManagementCommand
type UpdateWebSessionManagementCommandInput struct {
	Body models.WebSessionManagementView
}

//GetCookieTypes - Get the valid OIDC Cookie Types
//RequestType: GET
//Input:
func (s *WebSessionManagementService) GetCookieTypes() (output *models.CookieTypesView, resp *http.Response, err error) {
	path := "/webSessionManagement/cookieTypes"
	op := &request.Operation{
		Name:       "GetCookieTypes",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.CookieTypesView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetWebSessionSupportedEncryptionAlgorithmsCommand - Get the valid algorithms for Web Session Cookie Encryption
//RequestType: GET
//Input:
func (s *WebSessionManagementService) GetWebSessionSupportedEncryptionAlgorithmsCommand() (output *models.AlgorithmsView, resp *http.Response, err error) {
	path := "/webSessionManagement/encryptionAlgorithms"
	op := &request.Operation{
		Name:       "GetWebSessionSupportedEncryptionAlgorithmsCommand",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.AlgorithmsView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetWebSessionKeySetCommand - Get the Web Session key set configuration
//RequestType: GET
//Input:
func (s *WebSessionManagementService) GetWebSessionKeySetCommand() (output *models.KeySetView, resp *http.Response, err error) {
	path := "/webSessionManagement/keySet"
	op := &request.Operation{
		Name:       "GetWebSessionKeySetCommand",
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

//UpdateWebSessionKeySetCommand - Update the WebSession key set configuration
//RequestType: PUT
//Input: input *UpdateWebSessionKeySetCommandInput
func (s *WebSessionManagementService) UpdateWebSessionKeySetCommand(input *UpdateWebSessionKeySetCommandInput) (output *models.KeySetView, resp *http.Response, err error) {
	path := "/webSessionManagement/keySet"
	op := &request.Operation{
		Name:        "UpdateWebSessionKeySetCommand",
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

// UpdateWebSessionKeySetCommandInput - Inputs for UpdateWebSessionKeySetCommand
type UpdateWebSessionKeySetCommandInput struct {
	Body models.KeySetView
}

//GetOidcLoginTypes - Get the valid OIDC Login Types
//RequestType: GET
//Input:
func (s *WebSessionManagementService) GetOidcLoginTypes() (output *models.OidcLoginTypesView, resp *http.Response, err error) {
	path := "/webSessionManagement/oidcLoginTypes"
	op := &request.Operation{
		Name:       "GetOidcLoginTypes",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.OidcLoginTypesView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetOidcScopesCommand - Get the scopes supported by the current OIDC Provider
//RequestType: GET
//Input: input *GetOidcScopesCommandInput
func (s *WebSessionManagementService) GetOidcScopesCommand(input *GetOidcScopesCommandInput) (output *models.SupportedScopesView, resp *http.Response, err error) {
	path := "/webSessionManagement/oidcScopes"
	op := &request.Operation{
		Name:       "GetOidcScopesCommand",
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

// GetOidcScopesCommandInput - Inputs for GetOidcScopesCommand
type GetOidcScopesCommandInput struct {
	ClientId string
}

//GetRequestPreservationTypes - Get the valid Request Preservation Types
//RequestType: GET
//Input:
func (s *WebSessionManagementService) GetRequestPreservationTypes() (output *models.RequestPreservationTypesView, resp *http.Response, err error) {
	path := "/webSessionManagement/requestPreservationTypes"
	op := &request.Operation{
		Name:       "GetRequestPreservationTypes",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.RequestPreservationTypesView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetWebSessionSupportedSigningAlgorithms - Get the valid algorithms for Web Session Cookie Signing
//RequestType: GET
//Input:
func (s *WebSessionManagementService) GetWebSessionSupportedSigningAlgorithms() (output *models.SigningAlgorithmsView, resp *http.Response, err error) {
	path := "/webSessionManagement/signingAlgorithms"
	op := &request.Operation{
		Name:       "GetWebSessionSupportedSigningAlgorithms",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.SigningAlgorithmsView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetWebStorageTypes - Get the valid Web Storage Types
//RequestType: GET
//Input:
func (s *WebSessionManagementService) GetWebStorageTypes() (output *models.WebStorageTypesView, resp *http.Response, err error) {
	path := "/webSessionManagement/webStorageTypes"
	op := &request.Operation{
		Name:       "GetWebStorageTypes",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.WebStorageTypesView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}
