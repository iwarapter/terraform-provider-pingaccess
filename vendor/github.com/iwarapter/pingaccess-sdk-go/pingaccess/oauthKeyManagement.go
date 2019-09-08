package pingaccess

import (
	"fmt"
	"net/http"
	"net/url"
)

type OauthKeyManagementService service

//DeleteOAuthKeyManagementCommand - Resets the OAuth Key Management configuration to default values
//RequestType: DELETE
//Input:
func (s *OauthKeyManagementService) DeleteOAuthKeyManagementCommand() (resp *http.Response, err error) {
	path := "/oauthKeyManagement"
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

//GetOAuthKeyManagementCommand - Get the OAuth Key Management configuration
//RequestType: GET
//Input:
func (s *OauthKeyManagementService) GetOAuthKeyManagementCommand() (result *OAuthKeyManagementView, resp *http.Response, err error) {
	path := "/oauthKeyManagement"
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

//UpdateOAuthKeyManagementCommand - Update the OAuth Key Management configuration
//RequestType: PUT
//Input: input *UpdateOAuthKeyManagementCommandInput
func (s *OauthKeyManagementService) UpdateOAuthKeyManagementCommand(input *UpdateOAuthKeyManagementCommandInput) (result *OAuthKeyManagementView, resp *http.Response, err error) {
	path := "/oauthKeyManagement"
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

type UpdateOAuthKeyManagementCommandInput struct {
	Body OAuthKeyManagementView
}

//GetOAuthKeySetCommand - Get the OAuth key set configuration
//RequestType: GET
//Input:
func (s *OauthKeyManagementService) GetOAuthKeySetCommand() (result *KeySetView, resp *http.Response, err error) {
	path := "/oauthKeyManagement/keySet"
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

//UpdateOAuthKeySetCommand - Update the OAuth key set configuration
//RequestType: PUT
//Input: input *UpdateOAuthKeySetCommandInput
func (s *OauthKeyManagementService) UpdateOAuthKeySetCommand(input *UpdateOAuthKeySetCommandInput) (result *KeySetView, resp *http.Response, err error) {
	path := "/oauthKeyManagement/keySet"
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

type UpdateOAuthKeySetCommandInput struct {
	Body KeySetView
}
