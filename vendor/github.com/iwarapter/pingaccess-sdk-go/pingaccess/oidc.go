package pingaccess

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type OidcService service

type OidcAPI interface {
	DeleteOIDCProviderCommand() (resp *http.Response, err error)
	GetOIDCProviderCommand() (result *OIDCProviderView, resp *http.Response, err error)
	UpdateOIDCProviderCommand(input *UpdateOIDCProviderCommandInput) (result *OIDCProviderView, resp *http.Response, err error)
	GetOIDCProviderPluginDescriptorsCommand() (result *DescriptorsView, resp *http.Response, err error)
	GetOIDCProviderPluginDescriptorCommand(input *GetOIDCProviderPluginDescriptorCommandInput) (result *DescriptorView, resp *http.Response, err error)
	GetOIDCProviderMetadataCommand() (result *OIDCProviderMetadata, resp *http.Response, err error)
}

//DeleteOIDCProviderCommand - Resets the OpenID Connect Provider configuration to default values
//RequestType: DELETE
//Input:
func (s *OidcService) DeleteOIDCProviderCommand() (resp *http.Response, err error) {
	path := "/oidc/provider"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("DELETE", rel, nil)
	if err != nil {
		return nil, err
	}

	resp, err = s.client.do(req, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil

}

//GetOIDCProviderCommand - Get the OpenID Connect Provider configuration
//RequestType: GET
//Input:
func (s *OidcService) GetOIDCProviderCommand() (result *OIDCProviderView, resp *http.Response, err error) {
	path := "/oidc/provider"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("GET", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//UpdateOIDCProviderCommand - Update the OpenID Connect Provider configuration
//RequestType: PUT
//Input: input *UpdateOIDCProviderCommandInput
func (s *OidcService) UpdateOIDCProviderCommand(input *UpdateOIDCProviderCommandInput) (result *OIDCProviderView, resp *http.Response, err error) {
	path := "/oidc/provider"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("PUT", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

type UpdateOIDCProviderCommandInput struct {
	Body OIDCProviderView
}

//GetOIDCProviderPluginDescriptorsCommand - Get descriptors for all OIDC Provider plugins
//RequestType: GET
//Input:
func (s *OidcService) GetOIDCProviderPluginDescriptorsCommand() (result *DescriptorsView, resp *http.Response, err error) {
	path := "/oidc/provider/descriptors"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("GET", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//GetOIDCProviderPluginDescriptorCommand - Get a descriptor for a OIDC Provider plugin
//RequestType: GET
//Input: input *GetOIDCProviderPluginDescriptorCommandInput
func (s *OidcService) GetOIDCProviderPluginDescriptorCommand(input *GetOIDCProviderPluginDescriptorCommandInput) (result *DescriptorView, resp *http.Response, err error) {
	path := "/oidc/provider/descriptors/{pluginType}"
	path = strings.Replace(path, "{pluginType}", input.PluginType, -1)

	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("GET", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

type GetOIDCProviderPluginDescriptorCommandInput struct {
	PluginType string
}

//GetOIDCProviderMetadataCommand - Get the OpenID Connect Provider's metadata
//RequestType: GET
//Input:
func (s *OidcService) GetOIDCProviderMetadataCommand() (result *OIDCProviderMetadata, resp *http.Response, err error) {
	path := "/oidc/provider/metadata"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("GET", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}
