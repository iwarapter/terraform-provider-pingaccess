package pingaccess

import (
	"fmt"
	"net/http"
	"net/url"
)

type PingoneService service

//DeletePingOne4CCommand - Resets the PingOne For Customers configuration to default values
//RequestType: DELETE
//Input:
func (s *PingoneService) DeletePingOne4CCommand() (resp *http.Response, err error) {
	path := "/pingone/customers"
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

//GetPingOne4CCommand - Get the PingOne For Customers configuration
//RequestType: GET
//Input:
func (s *PingoneService) GetPingOne4CCommand() (result *PingOne4CView, resp *http.Response, err error) {
	path := "/pingone/customers"
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

//UpdatePingOne4CCommand - Update the PingOne For Customers configuration
//RequestType: PUT
//Input: input *UpdatePingOne4CCommandInput
func (s *PingoneService) UpdatePingOne4CCommand(input *UpdatePingOne4CCommandInput) (result *PingOne4CView, resp *http.Response, err error) {
	path := "/pingone/customers"
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

type UpdatePingOne4CCommandInput struct {
	Body PingOne4CView
}

//GetPingOne4CMetadataCommand - Get the Ping One for Customers metadata
//RequestType: GET
//Input:
func (s *PingoneService) GetPingOne4CMetadataCommand() (result *OIDCProviderMetadata, resp *http.Response, err error) {
	path := "/pingone/customers/metadata"
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
