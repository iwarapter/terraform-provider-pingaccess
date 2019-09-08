package pingaccess

import (
	"fmt"
	"net/http"
	"net/url"
)

type HttpConfigService service

//DeleteHostSourceCommand - Resets the HTTP request Host Source type to default values
//RequestType: DELETE
//Input:
func (s *HttpConfigService) DeleteHostSourceCommand() (resp *http.Response, err error) {
	path := "/httpConfig/request/hostSource"
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

//GetHostSourceCommand - Get the HTTP request Host Source type
//RequestType: GET
//Input:
func (s *HttpConfigService) GetHostSourceCommand() (result *HostMultiValueSourceView, resp *http.Response, err error) {
	path := "/httpConfig/request/hostSource"
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

//UpdateHostSourceCommand - Update the HTTP request Host Source type
//RequestType: PUT
//Input: input *UpdateHostSourceCommandInput
func (s *HttpConfigService) UpdateHostSourceCommand(input *UpdateHostSourceCommandInput) (result *HostMultiValueSourceView, resp *http.Response, err error) {
	path := "/httpConfig/request/hostSource"
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

type UpdateHostSourceCommandInput struct {
	Body HostMultiValueSourceView
}

//DeleteIpSourceCommand - Resets the HTTP request IP Source type to default values
//RequestType: DELETE
//Input:
func (s *HttpConfigService) DeleteIpSourceCommand() (resp *http.Response, err error) {
	path := "/httpConfig/request/ipSource"
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

//GetIpSourceCommand - Get the HTTP request IP Source type
//RequestType: GET
//Input:
func (s *HttpConfigService) GetIpSourceCommand() (result *IpMultiValueSourceView, resp *http.Response, err error) {
	path := "/httpConfig/request/ipSource"
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

//UpdateIpSourceCommand - Update the HTTP request IP Source type
//RequestType: PUT
//Input: input *UpdateIpSourceCommandInput
func (s *HttpConfigService) UpdateIpSourceCommand(input *UpdateIpSourceCommandInput) (result *IpMultiValueSourceView, resp *http.Response, err error) {
	path := "/httpConfig/request/ipSource"
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

type UpdateIpSourceCommandInput struct {
	Body IpMultiValueSourceView
}

//DeleteProtoSourceCommand - Resets the HTTP request Protocol Source type to default values
//RequestType: DELETE
//Input:
func (s *HttpConfigService) DeleteProtoSourceCommand() (resp *http.Response, err error) {
	path := "/httpConfig/request/protocolSource"
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

//GetProtoSourceCommand - Get the HTTP request Protocol Source type
//RequestType: GET
//Input:
func (s *HttpConfigService) GetProtoSourceCommand() (result *ProtocolSourceView, resp *http.Response, err error) {
	path := "/httpConfig/request/protocolSource"
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

//UpdateProtocolSourceCommand - Update the HTTP request Protocol Source type
//RequestType: PUT
//Input: input *UpdateProtocolSourceCommandInput
func (s *HttpConfigService) UpdateProtocolSourceCommand(input *UpdateProtocolSourceCommandInput) (result *ProtocolSourceView, resp *http.Response, err error) {
	path := "/httpConfig/request/protocolSource"
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

type UpdateProtocolSourceCommandInput struct {
	Body ProtocolSourceView
}
