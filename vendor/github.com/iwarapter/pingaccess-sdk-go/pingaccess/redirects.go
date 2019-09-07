package pingaccess

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type RedirectsService service

//GetRedirectsCommand - Get all Redirects
//RequestType: GET
//Input: input *GetRedirectsCommandInput
func (s *RedirectsService) GetRedirectsCommand(input *GetRedirectsCommandInput) (result *RedirectsView, resp *http.Response, err error) {
	path := "/redirects"
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
	if input.Source != "" {
		q.Set("source", input.Source)
	}
	if input.Target != "" {
		q.Set("target", input.Target)
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

type GetRedirectsCommandInput struct {
	Page          string
	NumberPerPage string
	Filter        string
	Source        string
	Target        string
	SortKey       string
	Order         string
}

//AddRedirectCommand - Add a Redirect
//RequestType: POST
//Input: input *AddRedirectCommandInput
func (s *RedirectsService) AddRedirectCommand(input *AddRedirectCommandInput) (result *RedirectView, resp *http.Response, err error) {
	path := "/redirects"
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

type AddRedirectCommandInput struct {
	Body RedirectView
}

//DeleteRedirectCommand - Delete a Redirect
//RequestType: DELETE
//Input: input *DeleteRedirectCommandInput
func (s *RedirectsService) DeleteRedirectCommand(input *DeleteRedirectCommandInput) (resp *http.Response, err error) {
	path := "/redirects/{id}"
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

type DeleteRedirectCommandInput struct {
	Id string
}

//GetRedirectCommand - Get a Redirect
//RequestType: GET
//Input: input *GetRedirectCommandInput
func (s *RedirectsService) GetRedirectCommand(input *GetRedirectCommandInput) (result *RedirectView, resp *http.Response, err error) {
	path := "/redirects/{id}"
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

type GetRedirectCommandInput struct {
	Id string
}

//UpdateRedirectCommand - Update a Redirect
//RequestType: PUT
//Input: input *UpdateRedirectCommandInput
func (s *RedirectsService) UpdateRedirectCommand(input *UpdateRedirectCommandInput) (result *RedirectView, resp *http.Response, err error) {
	path := "/redirects/{id}"
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

type UpdateRedirectCommandInput struct {
	Body RedirectView
	Id   string
}
