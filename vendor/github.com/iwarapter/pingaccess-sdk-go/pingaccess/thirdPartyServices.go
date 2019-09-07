package pingaccess

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type ThirdPartyServicesService service

//GetThirdPartyServicesCommand - Get all Third-Party Services
//RequestType: GET
//Input: input *GetThirdPartyServicesCommandInput
func (s *ThirdPartyServicesService) GetThirdPartyServicesCommand(input *GetThirdPartyServicesCommandInput) (result *ThirdPartyServicesView, resp *http.Response, err error) {
	path := "/thirdPartyServices"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	q := rel.Query()
	if input.Page != "" {
		q.Set("page", input.Page)
	}
	if input.NumberPerPage != "" {
		q.Set("numberPerPage", input.NumberPerPage)
	}
	if input.Filter != "" {
		q.Set("filter", input.Filter)
	}
	if input.Name != "" {
		q.Set("name", input.Name)
	}
	if input.SortKey != "" {
		q.Set("sortKey", input.SortKey)
	}
	if input.Order != "" {
		q.Set("order", input.Order)
	}
	rel.RawQuery = q.Encode()
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

type GetThirdPartyServicesCommandInput struct {
	Page          string
	NumberPerPage string
	Filter        string
	Name          string
	SortKey       string
	Order         string
}

//AddThirdPartyServiceCommand - Create a Third-Party Service
//RequestType: POST
//Input: input *AddThirdPartyServiceCommandInput
func (s *ThirdPartyServicesService) AddThirdPartyServiceCommand(input *AddThirdPartyServiceCommandInput) (result *ThirdPartyServiceView, resp *http.Response, err error) {
	path := "/thirdPartyServices"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("POST", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

type AddThirdPartyServiceCommandInput struct {
	Body ThirdPartyServiceView
}

//DeleteThirdPartyServiceCommand - Delete a Third-Party Service
//RequestType: DELETE
//Input: input *DeleteThirdPartyServiceCommandInput
func (s *ThirdPartyServicesService) DeleteThirdPartyServiceCommand(input *DeleteThirdPartyServiceCommandInput) (resp *http.Response, err error) {
	path := "/thirdPartyServices/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

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

type DeleteThirdPartyServiceCommandInput struct {
	Id string
}

//GetThirdPartyServiceCommand - Get a Third-Party Service
//RequestType: GET
//Input: input *GetThirdPartyServiceCommandInput
func (s *ThirdPartyServicesService) GetThirdPartyServiceCommand(input *GetThirdPartyServiceCommandInput) (result *ThirdPartyServiceView, resp *http.Response, err error) {
	path := "/thirdPartyServices/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

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

type GetThirdPartyServiceCommandInput struct {
	Id string
}

//UpdateThirdPartyServiceCommand - Update a Third-Party Service
//RequestType: PUT
//Input: input *UpdateThirdPartyServiceCommandInput
func (s *ThirdPartyServicesService) UpdateThirdPartyServiceCommand(input *UpdateThirdPartyServiceCommandInput) (result *ThirdPartyServiceView, resp *http.Response, err error) {
	path := "/thirdPartyServices/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

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

type UpdateThirdPartyServiceCommandInput struct {
	Body ThirdPartyServiceView
	Id   string
}
