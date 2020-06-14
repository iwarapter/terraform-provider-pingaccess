package pingaccess

import (
	"fmt"
	"net/http"
	"net/url"
)

type TokenProviderService service

type TokenProviderAPI interface {
	DeleteTokenProviderSettingCommand() (resp *http.Response, err error)
	GetTokenProviderSettingCommand() (result *TokenProviderSettingView, resp *http.Response, err error)
	UpdateTokenProviderSettingCommand(input *UpdateTokenProviderSettingCommandInput) (result *TokenProviderSettingView, resp *http.Response, err error)
}

//DeleteTokenProviderSettingCommand - Resets the Token Provider settings to default values
//RequestType: DELETE
//Input:
func (s *TokenProviderService) DeleteTokenProviderSettingCommand() (resp *http.Response, err error) {
	path := "/tokenProvider/settings"
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

//GetTokenProviderSettingCommand - Get the Token Provider settings
//RequestType: GET
//Input:
func (s *TokenProviderService) GetTokenProviderSettingCommand() (result *TokenProviderSettingView, resp *http.Response, err error) {
	path := "/tokenProvider/settings"
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

//UpdateTokenProviderSettingCommand - Update the Token Provider setting
//RequestType: PUT
//Input: input *UpdateTokenProviderSettingCommandInput
func (s *TokenProviderService) UpdateTokenProviderSettingCommand(input *UpdateTokenProviderSettingCommandInput) (result *TokenProviderSettingView, resp *http.Response, err error) {
	path := "/tokenProvider/settings"
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

type UpdateTokenProviderSettingCommandInput struct {
	Body TokenProviderSettingView
}
