package pingaccess

import (
	"fmt"
	"net/http"
	"net/url"
)

type AuthService service

type AuthAPI interface {
	DeleteBasicAuthCommand() (resp *http.Response, err error)
	GetBasicAuthCommand() (result *BasicConfig, resp *http.Response, err error)
	UpdateBasicAuthCommand(input *UpdateBasicAuthCommandInput) (result *BasicAuthConfigView, resp *http.Response, err error)
	DeleteOAuthAuthCommand() (resp *http.Response, err error)
	GetOAuthAuthCommand() (result *OAuthConfigView, resp *http.Response, err error)
	UpdateOAuthAuthCommand(input *UpdateOAuthAuthCommandInput) (result *OAuthConfigView, resp *http.Response, err error)
	DeleteOidcAuthCommand() (resp *http.Response, err error)
	GetOidcAuthCommand() (result *OidcConfigView, resp *http.Response, err error)
	UpdateOidcAuthCommand(input *UpdateOidcAuthCommandInput) (result *OidcConfigView, resp *http.Response, err error)
	DeleteAdminBasicWebSessionCommand() (resp *http.Response, err error)
	GetAdminBasicWebSessionCommand() (result *AdminBasicWebSessionView, resp *http.Response, err error)
	UpdateAdminBasicWebSessionCommand(input *UpdateAdminBasicWebSessionCommandInput) (result *AdminBasicWebSessionView, resp *http.Response, err error)
}

//DeleteBasicAuthCommand - Resets the HTTP Basic Authentication configuration to default values
//RequestType: DELETE
//Input:
func (s *AuthService) DeleteBasicAuthCommand() (resp *http.Response, err error) {
	path := "/auth/basic"
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

//GetBasicAuthCommand - Get the HTTP Basic Authentication configuration
//RequestType: GET
//Input:
func (s *AuthService) GetBasicAuthCommand() (result *BasicConfig, resp *http.Response, err error) {
	path := "/auth/basic"
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

//UpdateBasicAuthCommand - Update the Basic Authentication configuration
//RequestType: PUT
//Input: input *UpdateBasicAuthCommandInput
func (s *AuthService) UpdateBasicAuthCommand(input *UpdateBasicAuthCommandInput) (result *BasicAuthConfigView, resp *http.Response, err error) {
	path := "/auth/basic"
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

type UpdateBasicAuthCommandInput struct {
	Body BasicAuthConfigView
}

//DeleteOAuthAuthCommand - Resets the OAuth Authentication configuration to default values
//RequestType: DELETE
//Input:
func (s *AuthService) DeleteOAuthAuthCommand() (resp *http.Response, err error) {
	path := "/auth/oauth"
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

//GetOAuthAuthCommand - Get the OAuth Authentication configuration
//RequestType: GET
//Input:
func (s *AuthService) GetOAuthAuthCommand() (result *OAuthConfigView, resp *http.Response, err error) {
	path := "/auth/oauth"
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

//UpdateOAuthAuthCommand - Update the OAuth Authentication configuration
//RequestType: PUT
//Input: input *UpdateOAuthAuthCommandInput
func (s *AuthService) UpdateOAuthAuthCommand(input *UpdateOAuthAuthCommandInput) (result *OAuthConfigView, resp *http.Response, err error) {
	path := "/auth/oauth"
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

type UpdateOAuthAuthCommandInput struct {
	Body OAuthConfigView
}

//DeleteOidcAuthCommand - Resets the OIDC Authentication configuration to default values
//RequestType: DELETE
//Input:
func (s *AuthService) DeleteOidcAuthCommand() (resp *http.Response, err error) {
	path := "/auth/oidc"
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

//GetOidcAuthCommand - Get the OIDC Authentication configuration
//RequestType: GET
//Input:
func (s *AuthService) GetOidcAuthCommand() (result *OidcConfigView, resp *http.Response, err error) {
	path := "/auth/oidc"
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

//UpdateOidcAuthCommand - Update the OIDC Authentication configuration
//RequestType: PUT
//Input: input *UpdateOidcAuthCommandInput
func (s *AuthService) UpdateOidcAuthCommand(input *UpdateOidcAuthCommandInput) (result *OidcConfigView, resp *http.Response, err error) {
	path := "/auth/oidc"
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

type UpdateOidcAuthCommandInput struct {
	Body OidcConfigView
}

//DeleteAdminBasicWebSessionCommand - Resets the Admin Web Session configuration to default values
//RequestType: DELETE
//Input:
func (s *AuthService) DeleteAdminBasicWebSessionCommand() (resp *http.Response, err error) {
	path := "/auth/webSession"
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

//GetAdminBasicWebSessionCommand - Get the admin web session configuration
//RequestType: GET
//Input:
func (s *AuthService) GetAdminBasicWebSessionCommand() (result *AdminBasicWebSessionView, resp *http.Response, err error) {
	path := "/auth/webSession"
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

//UpdateAdminBasicWebSessionCommand - Update the admin web session configuration
//RequestType: PUT
//Input: input *UpdateAdminBasicWebSessionCommandInput
func (s *AuthService) UpdateAdminBasicWebSessionCommand(input *UpdateAdminBasicWebSessionCommandInput) (result *AdminBasicWebSessionView, resp *http.Response, err error) {
	path := "/auth/webSession"
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

type UpdateAdminBasicWebSessionCommandInput struct {
	Body AdminBasicWebSessionView
}
