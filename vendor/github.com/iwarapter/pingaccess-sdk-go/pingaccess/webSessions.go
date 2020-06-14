package pingaccess

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type WebSessionsService service

type WebSessionsAPI interface {
	GetWebSessionsCommand(input *GetWebSessionsCommandInput) (result *WebSessionsView, resp *http.Response, err error)
	AddWebSessionCommand(input *AddWebSessionCommandInput) (result *WebSessionView, resp *http.Response, err error)
	DeleteWebSessionCommand(input *DeleteWebSessionCommandInput) (resp *http.Response, err error)
	GetWebSessionCommand(input *GetWebSessionCommandInput) (result *WebSessionView, resp *http.Response, err error)
	UpdateWebSessionCommand(input *UpdateWebSessionCommandInput) (result *WebSessionView, resp *http.Response, err error)
}

//GetWebSessionsCommand - Get all WebSessions
//RequestType: GET
//Input: input *GetWebSessionsCommandInput
func (s *WebSessionsService) GetWebSessionsCommand(input *GetWebSessionsCommandInput) (result *WebSessionsView, resp *http.Response, err error) {
	path := "/webSessions"
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

type GetWebSessionsCommandInput struct {
	Page          string
	NumberPerPage string
	Filter        string
	Name          string
	SortKey       string
	Order         string
}

//AddWebSessionCommand - Create a WebSession
//RequestType: POST
//Input: input *AddWebSessionCommandInput
func (s *WebSessionsService) AddWebSessionCommand(input *AddWebSessionCommandInput) (result *WebSessionView, resp *http.Response, err error) {
	path := "/webSessions"
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

type AddWebSessionCommandInput struct {
	Body WebSessionView
}

//DeleteWebSessionCommand - Delete a WebSession
//RequestType: DELETE
//Input: input *DeleteWebSessionCommandInput
func (s *WebSessionsService) DeleteWebSessionCommand(input *DeleteWebSessionCommandInput) (resp *http.Response, err error) {
	path := "/webSessions/{id}"
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

type DeleteWebSessionCommandInput struct {
	Id string
}

//GetWebSessionCommand - Get a WebSession
//RequestType: GET
//Input: input *GetWebSessionCommandInput
func (s *WebSessionsService) GetWebSessionCommand(input *GetWebSessionCommandInput) (result *WebSessionView, resp *http.Response, err error) {
	path := "/webSessions/{id}"
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

type GetWebSessionCommandInput struct {
	Id string
}

//UpdateWebSessionCommand - Update a WebSession
//RequestType: PUT
//Input: input *UpdateWebSessionCommandInput
func (s *WebSessionsService) UpdateWebSessionCommand(input *UpdateWebSessionCommandInput) (result *WebSessionView, resp *http.Response, err error) {
	path := "/webSessions/{id}"
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

type UpdateWebSessionCommandInput struct {
	Body WebSessionView
	Id   string
}
