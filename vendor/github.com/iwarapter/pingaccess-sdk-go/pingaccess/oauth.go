package pingaccess

import (
	"fmt"
	"net/http"
	"net/url"
)

type OauthService service

//DeleteAuthorizationServerCommand - Resets the OpenID Connect Provider configuration to default values
//RequestType: DELETE
//Input:
func (s *OauthService) DeleteAuthorizationServerCommand() (resp *http.Response, err error) {
	path := "/oauth/authServer"
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

//GetAuthorizationServerCommand - Get Authorization Server configuration
//RequestType: GET
//Input:
func (s *OauthService) GetAuthorizationServerCommand() (result *AuthorizationServerView, resp *http.Response, err error) {
	path := "/oauth/authServer"
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

//UpdateAuthorizationServerCommand - Update OAuth 2.0 Authorization Server configuration
//RequestType: PUT
//Input: input *UpdateAuthorizationServerCommandInput
func (s *OauthService) UpdateAuthorizationServerCommand(input *UpdateAuthorizationServerCommandInput) (result *AuthorizationServerView, resp *http.Response, err error) {
	path := "/oauth/authServer"
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

type UpdateAuthorizationServerCommandInput struct {
	Body AuthorizationServerView
}
