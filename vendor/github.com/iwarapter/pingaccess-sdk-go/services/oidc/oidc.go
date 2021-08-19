package oidc

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
	ServiceName = "Oidc"
)

//OidcService provides the API operations for making requests to
// Oidc endpoint.
type OidcService struct {
	*client.Client
}

//New createa a new instance of the OidcService client.
//
// Example:
//   cfg := config.NewConfig().WithUsername("Administrator").WithPassword("2Access").WithEndpoint(paURL)
//
//   //Create a OidcService from the configuration
//   svc := oidc.New(cfg)
//
func New(cfg *config.Config) *OidcService {

	return &OidcService{Client: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingaccess.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a Oidc operation
func (s *OidcService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := s.NewRequest(op, params, data)

	return req
}

//DeleteOIDCProviderCommand - Resets the OpenID Connect Provider configuration to default values
//RequestType: DELETE
//Input:
func (s *OidcService) DeleteOIDCProviderCommand() (resp *http.Response, err error) {
	path := "/oidc/provider"
	op := &request.Operation{
		Name:       "DeleteOIDCProviderCommand",
		HTTPMethod: "DELETE",
		HTTPPath:   path,
	}

	req := s.newRequest(op, nil, nil)

	if req.Send() == nil {
		return req.HTTPResponse, nil
	}
	return req.HTTPResponse, req.Error
}

//GetOIDCProviderCommand - Get the OpenID Connect Provider configuration
//RequestType: GET
//Input:
func (s *OidcService) GetOIDCProviderCommand() (output *models.OIDCProviderView, resp *http.Response, err error) {
	path := "/oidc/provider"
	op := &request.Operation{
		Name:       "GetOIDCProviderCommand",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.OIDCProviderView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateOIDCProviderCommand - Update the OpenID Connect Provider configuration
//RequestType: PUT
//Input: input *UpdateOIDCProviderCommandInput
func (s *OidcService) UpdateOIDCProviderCommand(input *UpdateOIDCProviderCommandInput) (output *models.OIDCProviderView, resp *http.Response, err error) {
	path := "/oidc/provider"
	op := &request.Operation{
		Name:        "UpdateOIDCProviderCommand",
		HTTPMethod:  "PUT",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.OIDCProviderView{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// UpdateOIDCProviderCommandInput - Inputs for UpdateOIDCProviderCommand
type UpdateOIDCProviderCommandInput struct {
	Body models.OIDCProviderView
}

//GetOIDCProviderPluginDescriptorsCommand - Get descriptors for all OIDC Provider plugins
//RequestType: GET
//Input:
func (s *OidcService) GetOIDCProviderPluginDescriptorsCommand() (output *models.DescriptorsView, resp *http.Response, err error) {
	path := "/oidc/provider/descriptors"
	op := &request.Operation{
		Name:       "GetOIDCProviderPluginDescriptorsCommand",
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

//GetOIDCProviderPluginDescriptorCommand - Get a descriptor for a OIDC Provider plugin
//RequestType: GET
//Input: input *GetOIDCProviderPluginDescriptorCommandInput
func (s *OidcService) GetOIDCProviderPluginDescriptorCommand(input *GetOIDCProviderPluginDescriptorCommandInput) (output *models.DescriptorView, resp *http.Response, err error) {
	path := "/oidc/provider/descriptors/{pluginType}"
	path = strings.Replace(path, "{pluginType}", input.PluginType, -1)

	op := &request.Operation{
		Name:        "GetOIDCProviderPluginDescriptorCommand",
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

// GetOIDCProviderPluginDescriptorCommandInput - Inputs for GetOIDCProviderPluginDescriptorCommand
type GetOIDCProviderPluginDescriptorCommandInput struct {
	PluginType string
}

//GetOIDCProviderMetadataCommand - Get the OpenID Connect Provider's metadata
//RequestType: GET
//Input:
func (s *OidcService) GetOIDCProviderMetadataCommand() (output *models.OIDCProviderMetadata, resp *http.Response, err error) {
	path := "/oidc/provider/metadata"
	op := &request.Operation{
		Name:       "GetOIDCProviderMetadataCommand",
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
