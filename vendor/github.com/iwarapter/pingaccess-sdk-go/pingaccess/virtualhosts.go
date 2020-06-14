package pingaccess

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type VirtualhostsService service

type VirtualhostsAPI interface {
	GetVirtualHostsCommand(input *GetVirtualHostsCommandInput) (result *VirtualHostsView, resp *http.Response, err error)
	AddVirtualHostCommand(input *AddVirtualHostCommandInput) (result *VirtualHostView, resp *http.Response, err error)
	DeleteVirtualHostCommand(input *DeleteVirtualHostCommandInput) (resp *http.Response, err error)
	GetVirtualHostCommand(input *GetVirtualHostCommandInput) (result *VirtualHostView, resp *http.Response, err error)
	UpdateVirtualHostCommand(input *UpdateVirtualHostCommandInput) (result *VirtualHostView, resp *http.Response, err error)
}

//GetVirtualHostsCommand - Get all Virtual Hosts
//RequestType: GET
//Input: input *GetVirtualHostsCommandInput
func (s *VirtualhostsService) GetVirtualHostsCommand(input *GetVirtualHostsCommandInput) (result *VirtualHostsView, resp *http.Response, err error) {
	path := "/virtualhosts"
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
	if input.VirtualHost != "" {
		q.Set("virtualHost", input.VirtualHost)
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

type GetVirtualHostsCommandInput struct {
	Page          string
	NumberPerPage string
	Filter        string
	VirtualHost   string
	SortKey       string
	Order         string
}

//AddVirtualHostCommand - Create a Virtual Host
//RequestType: POST
//Input: input *AddVirtualHostCommandInput
func (s *VirtualhostsService) AddVirtualHostCommand(input *AddVirtualHostCommandInput) (result *VirtualHostView, resp *http.Response, err error) {
	path := "/virtualhosts"
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

type AddVirtualHostCommandInput struct {
	Body VirtualHostView
}

//DeleteVirtualHostCommand - Delete a Virtual Host
//RequestType: DELETE
//Input: input *DeleteVirtualHostCommandInput
func (s *VirtualhostsService) DeleteVirtualHostCommand(input *DeleteVirtualHostCommandInput) (resp *http.Response, err error) {
	path := "/virtualhosts/{id}"
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

type DeleteVirtualHostCommandInput struct {
	Id string
}

//GetVirtualHostCommand - Get a Virtual Host
//RequestType: GET
//Input: input *GetVirtualHostCommandInput
func (s *VirtualhostsService) GetVirtualHostCommand(input *GetVirtualHostCommandInput) (result *VirtualHostView, resp *http.Response, err error) {
	path := "/virtualhosts/{id}"
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

type GetVirtualHostCommandInput struct {
	Id string
}

//UpdateVirtualHostCommand - Update a Virtual Host
//RequestType: PUT
//Input: input *UpdateVirtualHostCommandInput
func (s *VirtualhostsService) UpdateVirtualHostCommand(input *UpdateVirtualHostCommandInput) (result *VirtualHostView, resp *http.Response, err error) {
	path := "/virtualhosts/{id}"
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

type UpdateVirtualHostCommandInput struct {
	Body VirtualHostView
	Id   string
}
