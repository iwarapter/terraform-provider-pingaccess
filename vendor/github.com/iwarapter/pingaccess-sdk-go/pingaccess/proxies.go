package pingaccess

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type ProxiesService service

type ProxiesAPI interface {
	GetProxiesCommand(input *GetProxiesCommandInput) (result *HttpClientProxyView, resp *http.Response, err error)
	AddProxyCommand(input *AddProxyCommandInput) (result *HttpClientProxyView, resp *http.Response, err error)
	DeleteProxyCommand(input *DeleteProxyCommandInput) (resp *http.Response, err error)
	GetProxyCommand(input *GetProxyCommandInput) (result *HttpClientProxyView, resp *http.Response, err error)
	UpdateProxyCommand(input *UpdateProxyCommandInput) (result *HttpClientProxyView, resp *http.Response, err error)
}

//GetProxiesCommand - Get all Proxies
//RequestType: GET
//Input: input *GetProxiesCommandInput
func (s *ProxiesService) GetProxiesCommand(input *GetProxiesCommandInput) (result *HttpClientProxyView, resp *http.Response, err error) {
	path := "/proxies"
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

type GetProxiesCommandInput struct {
	Page          string
	NumberPerPage string
	Filter        string
	Name          string
	SortKey       string
	Order         string
}

//AddProxyCommand - Create a Proxy
//RequestType: POST
//Input: input *AddProxyCommandInput
func (s *ProxiesService) AddProxyCommand(input *AddProxyCommandInput) (result *HttpClientProxyView, resp *http.Response, err error) {
	path := "/proxies"
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

type AddProxyCommandInput struct {
	Body HttpClientProxyView
}

//DeleteProxyCommand - Delete a Proxy
//RequestType: DELETE
//Input: input *DeleteProxyCommandInput
func (s *ProxiesService) DeleteProxyCommand(input *DeleteProxyCommandInput) (resp *http.Response, err error) {
	path := "/proxies/{id}"
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

type DeleteProxyCommandInput struct {
	Id string
}

//GetProxyCommand - Get a Proxy
//RequestType: GET
//Input: input *GetProxyCommandInput
func (s *ProxiesService) GetProxyCommand(input *GetProxyCommandInput) (result *HttpClientProxyView, resp *http.Response, err error) {
	path := "/proxies/{id}"
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

type GetProxyCommandInput struct {
	Id string
}

//UpdateProxyCommand - Update a Proxy
//RequestType: PUT
//Input: input *UpdateProxyCommandInput
func (s *ProxiesService) UpdateProxyCommand(input *UpdateProxyCommandInput) (result *HttpClientProxyView, resp *http.Response, err error) {
	path := "/proxies/{id}"
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

type UpdateProxyCommandInput struct {
	Body HttpClientProxyView
	Id   string
}
