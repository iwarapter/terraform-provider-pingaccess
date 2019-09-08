package pingaccess

import (
	"fmt"
	"net/http"
	"net/url"
)

type AuthTokenManagementService service

//DeleteAuthTokenManagementCommand - Resets the Auth Token Management configuration to default values
//RequestType: DELETE
//Input:
func (s *AuthTokenManagementService) DeleteAuthTokenManagementCommand() (resp *http.Response, err error) {
	path := "/authTokenManagement"
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

//GetAuthTokenManagementCommand - Get the Auth Token Management configuration
//RequestType: GET
//Input:
func (s *AuthTokenManagementService) GetAuthTokenManagementCommand() (result *AuthTokenManagementView, resp *http.Response, err error) {
	path := "/authTokenManagement"
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

//UpdateAuthTokenManagementCommand - Update the Auth Token Management configuration
//RequestType: PUT
//Input: input *UpdateAuthTokenManagementCommandInput
func (s *AuthTokenManagementService) UpdateAuthTokenManagementCommand(input *UpdateAuthTokenManagementCommandInput) (result *AuthTokenManagementView, resp *http.Response, err error) {
	path := "/authTokenManagement"
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

type UpdateAuthTokenManagementCommandInput struct {
	Body AuthTokenManagementView
}

//GetAuthTokenKeySetCommand - Get the Auth Token key set configuration
//RequestType: GET
//Input:
func (s *AuthTokenManagementService) GetAuthTokenKeySetCommand() (result *KeySetView, resp *http.Response, err error) {
	path := "/authTokenManagement/keySet"
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

//UpdateAuthTokenKeySetCommand - Update the AuthToken key set configuration
//RequestType: PUT
//Input: input *UpdateAuthTokenKeySetCommandInput
func (s *AuthTokenManagementService) UpdateAuthTokenKeySetCommand(input *UpdateAuthTokenKeySetCommandInput) (result *KeySetView, resp *http.Response, err error) {
	path := "/authTokenManagement/keySet"
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

type UpdateAuthTokenKeySetCommandInput struct {
	Body KeySetView
}
