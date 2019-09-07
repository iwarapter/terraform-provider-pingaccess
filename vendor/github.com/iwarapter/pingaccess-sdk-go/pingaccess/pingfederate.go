package pingaccess

import (
	"fmt"
	"net/http"
	"net/url"
)

type PingfederateService service

//DeletePingFederateCommand - Resets the PingFederate configuration to default values
//RequestType: DELETE
//Input:
func (s *PingfederateService) DeletePingFederateCommand() (resp *http.Response, err error) {
	path := "/pingfederate"
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

//GetPingFederateCommand - Get the PingFederate configuration
//RequestType: GET
//Input:
func (s *PingfederateService) GetPingFederateCommand() (result *PingFederateRuntimeView, resp *http.Response, err error) {
	path := "/pingfederate"
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

//UpdatePingFederateCommand - Update the PingFederate configuration
//RequestType: PUT
//Input: input *UpdatePingFederateCommandInput
func (s *PingfederateService) UpdatePingFederateCommand(input *UpdatePingFederateCommandInput) (result *PingFederateRuntimeView, resp *http.Response, err error) {
	path := "/pingfederate"
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

type UpdatePingFederateCommandInput struct {
	Body PingFederateRuntimeView
}

//DeletePingFederateAccessTokensCommand - Resets the PingAccess OAuth Client configuration to default values
//RequestType: DELETE
//Input:
func (s *PingfederateService) DeletePingFederateAccessTokensCommand() (resp *http.Response, err error) {
	path := "/pingfederate/accessTokens"
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

//GetPingFederateAccessTokensCommand - Get the PingAccess OAuth Client configuration
//RequestType: GET
//Input:
func (s *PingfederateService) GetPingFederateAccessTokensCommand() (result *PingFederateAccessTokenView, resp *http.Response, err error) {
	path := "/pingfederate/accessTokens"
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

//UpdatePingFederateAccessTokensCommand - Update the PingFederate OAuth Client configuration
//RequestType: PUT
//Input: input *UpdatePingFederateAccessTokensCommandInput
func (s *PingfederateService) UpdatePingFederateAccessTokensCommand(input *UpdatePingFederateAccessTokensCommandInput) (result *PingFederateAccessTokenView, resp *http.Response, err error) {
	path := "/pingfederate/accessTokens"
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

type UpdatePingFederateAccessTokensCommandInput struct {
	Body PingFederateAccessTokenView
}

//DeletePingFederateAdminCommand - Resets the PingFederate Admin configuration to default values
//RequestType: DELETE
//Input:
func (s *PingfederateService) DeletePingFederateAdminCommand() (resp *http.Response, err error) {
	path := "/pingfederate/admin"
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

//GetPingFederateAdminCommand - Get the PingFederate Admin configuration
//RequestType: GET
//Input:
func (s *PingfederateService) GetPingFederateAdminCommand() (result *PingFederateAdminView, resp *http.Response, err error) {
	path := "/pingfederate/admin"
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

//UpdatePingFederateAdminCommand - Update the PingFederate Admin configuration
//RequestType: PUT
//Input: input *UpdatePingFederateAdminCommandInput
func (s *PingfederateService) UpdatePingFederateAdminCommand(input *UpdatePingFederateAdminCommandInput) (result *PingFederateAdminView, resp *http.Response, err error) {
	path := "/pingfederate/admin"
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

type UpdatePingFederateAdminCommandInput struct {
	Body PingFederateAdminView
}
