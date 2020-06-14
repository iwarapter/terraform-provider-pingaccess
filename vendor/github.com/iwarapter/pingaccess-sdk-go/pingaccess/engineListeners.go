package pingaccess

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type EngineListenersService service

type EngineListenersAPI interface {
	GetEngineListenersCommand(input *GetEngineListenersCommandInput) (result *EngineListenersView, resp *http.Response, err error)
	AddEngineListenerCommand(input *AddEngineListenerCommandInput) (result *EngineListenerView, resp *http.Response, err error)
	DeleteEngineListenerCommand(input *DeleteEngineListenerCommandInput) (resp *http.Response, err error)
	GetEngineListenerCommand(input *GetEngineListenerCommandInput) (result *EngineListenerView, resp *http.Response, err error)
	UpdateEngineListenerCommand(input *UpdateEngineListenerCommandInput) (result *EngineListenerView, resp *http.Response, err error)
}

//GetEngineListenersCommand - Get all Engine Listeners
//RequestType: GET
//Input: input *GetEngineListenersCommandInput
func (s *EngineListenersService) GetEngineListenersCommand(input *GetEngineListenersCommandInput) (result *EngineListenersView, resp *http.Response, err error) {
	path := "/engineListeners"
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

type GetEngineListenersCommandInput struct {
	Page          string
	NumberPerPage string
	Filter        string
	Name          string
	SortKey       string
	Order         string
}

//AddEngineListenerCommand - Create an Engine Listener
//RequestType: POST
//Input: input *AddEngineListenerCommandInput
func (s *EngineListenersService) AddEngineListenerCommand(input *AddEngineListenerCommandInput) (result *EngineListenerView, resp *http.Response, err error) {
	path := "/engineListeners"
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

type AddEngineListenerCommandInput struct {
	Body EngineListenerView
}

//DeleteEngineListenerCommand - Delete an Engine Listener
//RequestType: DELETE
//Input: input *DeleteEngineListenerCommandInput
func (s *EngineListenersService) DeleteEngineListenerCommand(input *DeleteEngineListenerCommandInput) (resp *http.Response, err error) {
	path := "/engineListeners/{id}"
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

type DeleteEngineListenerCommandInput struct {
	Id string
}

//GetEngineListenerCommand - Get an Engine Listener
//RequestType: GET
//Input: input *GetEngineListenerCommandInput
func (s *EngineListenersService) GetEngineListenerCommand(input *GetEngineListenerCommandInput) (result *EngineListenerView, resp *http.Response, err error) {
	path := "/engineListeners/{id}"
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

type GetEngineListenerCommandInput struct {
	Id string
}

//UpdateEngineListenerCommand - Update an Engine Listener
//RequestType: PUT
//Input: input *UpdateEngineListenerCommandInput
func (s *EngineListenersService) UpdateEngineListenerCommand(input *UpdateEngineListenerCommandInput) (result *EngineListenerView, resp *http.Response, err error) {
	path := "/engineListeners/{id}"
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

type UpdateEngineListenerCommandInput struct {
	Body EngineListenerView
	Id   string
}
