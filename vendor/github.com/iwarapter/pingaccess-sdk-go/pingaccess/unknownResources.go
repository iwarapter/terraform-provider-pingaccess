package pingaccess

import (
	"fmt"
	"net/http"
	"net/url"
)

type UnknownResourcesService service

type UnknownResourcesAPI interface {
	Delete() (resp *http.Response, err error)
	Get() (result *UnknownResourceSettingsView, resp *http.Response, err error)
	Update(input *UpdateInput) (result *UnknownResourceSettingsView, resp *http.Response, err error)
}

//Delete - Resets the global settings for unknown resources to default values
//RequestType: DELETE
//Input:
func (s *UnknownResourcesService) Delete() (resp *http.Response, err error) {
	path := "/unknownResources/settings"
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

//Get - Retrieves the global settings for unknown resources
//RequestType: GET
//Input:
func (s *UnknownResourcesService) Get() (result *UnknownResourceSettingsView, resp *http.Response, err error) {
	path := "/unknownResources/settings"
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

//Update - Updates the global settings for unknown resources
//RequestType: PUT
//Input: input *UpdateInput
func (s *UnknownResourcesService) Update(input *UpdateInput) (result *UnknownResourceSettingsView, resp *http.Response, err error) {
	path := "/unknownResources/settings"
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

type UpdateInput struct {
	Body UnknownResourceSettingsView
}
