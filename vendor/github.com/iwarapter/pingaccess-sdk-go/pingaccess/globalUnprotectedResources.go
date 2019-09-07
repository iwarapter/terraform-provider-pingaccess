package pingaccess

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type GlobalUnprotectedResourcesService service

//GetGlobalUnprotectedResourcesCommand - Get all Global Unprotected Resources
//RequestType: GET
//Input: input *GetGlobalUnprotectedResourcesCommandInput
func (s *GlobalUnprotectedResourcesService) GetGlobalUnprotectedResourcesCommand(input *GetGlobalUnprotectedResourcesCommandInput) (result *GlobalUnprotectedResourcesView, resp *http.Response, err error) {
	path := "/globalUnprotectedResources"
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

type GetGlobalUnprotectedResourcesCommandInput struct {
	Page          string
	NumberPerPage string
	Filter        string
	Name          string
	SortKey       string
	Order         string
}

//AddGlobalUnprotectedResourceCommand - Add a Global Unprotected Resource
//RequestType: POST
//Input: input *AddGlobalUnprotectedResourceCommandInput
func (s *GlobalUnprotectedResourcesService) AddGlobalUnprotectedResourceCommand(input *AddGlobalUnprotectedResourceCommandInput) (result *GlobalUnprotectedResourceView, resp *http.Response, err error) {
	path := "/globalUnprotectedResources"
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

type AddGlobalUnprotectedResourceCommandInput struct {
	Body GlobalUnprotectedResourceView
}

//DeleteGlobalUnprotectedResourceCommand - Delete a Global Unprotected Resource
//RequestType: DELETE
//Input: input *DeleteGlobalUnprotectedResourceCommandInput
func (s *GlobalUnprotectedResourcesService) DeleteGlobalUnprotectedResourceCommand(input *DeleteGlobalUnprotectedResourceCommandInput) (resp *http.Response, err error) {
	path := "/globalUnprotectedResources/{id}"
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

type DeleteGlobalUnprotectedResourceCommandInput struct {
	Id string
}

//GetGlobalUnprotectedResourceCommand - Get a Global Unprotected Resource
//RequestType: GET
//Input: input *GetGlobalUnprotectedResourceCommandInput
func (s *GlobalUnprotectedResourcesService) GetGlobalUnprotectedResourceCommand(input *GetGlobalUnprotectedResourceCommandInput) (result *GlobalUnprotectedResourceView, resp *http.Response, err error) {
	path := "/globalUnprotectedResources/{id}"
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

type GetGlobalUnprotectedResourceCommandInput struct {
	Id string
}

//UpdateGlobalUnprotectedResourceCommand - Update a Global Unprotected Resource
//RequestType: PUT
//Input: input *UpdateGlobalUnprotectedResourceCommandInput
func (s *GlobalUnprotectedResourcesService) UpdateGlobalUnprotectedResourceCommand(input *UpdateGlobalUnprotectedResourceCommandInput) (result *GlobalUnprotectedResourceView, resp *http.Response, err error) {
	path := "/globalUnprotectedResources/{id}"
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

type UpdateGlobalUnprotectedResourceCommandInput struct {
	Body GlobalUnprotectedResourceView
	Id   string
}
