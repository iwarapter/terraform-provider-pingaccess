package pingaccess

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type SitesService service

//GetSitesCommand - Get all Sites
//RequestType: GET
//Input: input *GetSitesCommandInput
func (s *SitesService) GetSitesCommand(input *GetSitesCommandInput) (result *SitesView, resp *http.Response, err error) {
	path := "/sites"
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

type GetSitesCommandInput struct {
	Page          string
	NumberPerPage string
	Filter        string
	Name          string
	SortKey       string
	Order         string
}

//AddSiteCommand - Create a Site
//RequestType: POST
//Input: input *AddSiteCommandInput
func (s *SitesService) AddSiteCommand(input *AddSiteCommandInput) (result *SiteView, resp *http.Response, err error) {
	path := "/sites"
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

type AddSiteCommandInput struct {
	Body SiteView
}

//DeleteSiteCommand - Delete a Site
//RequestType: DELETE
//Input: input *DeleteSiteCommandInput
func (s *SitesService) DeleteSiteCommand(input *DeleteSiteCommandInput) (resp *http.Response, err error) {
	path := "/sites/{id}"
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

type DeleteSiteCommandInput struct {
	Id string
}

//GetSiteCommand - Get a Site
//RequestType: GET
//Input: input *GetSiteCommandInput
func (s *SitesService) GetSiteCommand(input *GetSiteCommandInput) (result *SiteView, resp *http.Response, err error) {
	path := "/sites/{id}"
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

type GetSiteCommandInput struct {
	Id string
}

//UpdateSiteCommand - Update a Site
//RequestType: PUT
//Input: input *UpdateSiteCommandInput
func (s *SitesService) UpdateSiteCommand(input *UpdateSiteCommandInput) (result *SiteView, resp *http.Response, err error) {
	path := "/sites/{id}"
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

type UpdateSiteCommandInput struct {
	Body SiteView
	Id   string
}
