package pingaccess

import (
	"fmt"
	"net/http"
	"net/url"
)

type AdminSessionInfoService service

//AdminSessionDeleteCommand - Invalidate the Admin session information
//RequestType: DELETE
//Input:
func (s *AdminSessionInfoService) AdminSessionDeleteCommand() (resp *http.Response, err error) {
	path := "/adminSessionInfo"
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

//AdminSessionInfoCommand - Return the Admin session information
//RequestType: GET
//Input:
func (s *AdminSessionInfoService) AdminSessionInfoCommand() (result *SessionInfo, resp *http.Response, err error) {
	path := "/adminSessionInfo"
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

//AdminSessionInfoCheckCommand - Return the Admin session information without affecting session expiration
//RequestType: GET
//Input:
func (s *AdminSessionInfoService) AdminSessionInfoCheckCommand() (result *SessionInfo, resp *http.Response, err error) {
	path := "/adminSessionInfo/checkOnly"
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
