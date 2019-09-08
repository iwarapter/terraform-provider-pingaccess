package pingaccess

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type SiteAuthenticatorsService service

//GetSiteAuthenticatorsCommand - Get all Site Authenticators
//RequestType: GET
//Input: input *GetSiteAuthenticatorsCommandInput
func (s *SiteAuthenticatorsService) GetSiteAuthenticatorsCommand(input *GetSiteAuthenticatorsCommandInput) (result *SiteAuthenticatorsView, resp *http.Response, err error) {
	path := "/siteAuthenticators"
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

type GetSiteAuthenticatorsCommandInput struct {
	Page          string
	NumberPerPage string
	Filter        string
	Name          string
	SortKey       string
	Order         string
}

//AddSiteAuthenticatorCommand - Create a Site Authenticator
//RequestType: POST
//Input: input *AddSiteAuthenticatorCommandInput
func (s *SiteAuthenticatorsService) AddSiteAuthenticatorCommand(input *AddSiteAuthenticatorCommandInput) (result *SiteAuthenticatorView, resp *http.Response, err error) {
	path := "/siteAuthenticators"
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

type AddSiteAuthenticatorCommandInput struct {
	Body SiteAuthenticatorView
}

//GetSiteAuthenticatorDescriptorsCommand - Get descriptors for all supported Site Authenticator types
//RequestType: GET
//Input:
func (s *SiteAuthenticatorsService) GetSiteAuthenticatorDescriptorsCommand() (result *DescriptorsView, resp *http.Response, err error) {
	path := "/siteAuthenticators/descriptors"
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

//GetSiteAuthenticatorDescriptorCommand - Get descriptor for a Site Authenticator type
//RequestType: GET
//Input: input *GetSiteAuthenticatorDescriptorCommandInput
func (s *SiteAuthenticatorsService) GetSiteAuthenticatorDescriptorCommand(input *GetSiteAuthenticatorDescriptorCommandInput) (result *DescriptorView, resp *http.Response, err error) {
	path := "/siteAuthenticators/descriptors/{siteAuthenticatorType}"
	path = strings.Replace(path, "{siteAuthenticatorType}", input.SiteAuthenticatorType, -1)

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

type GetSiteAuthenticatorDescriptorCommandInput struct {
	SiteAuthenticatorType string
}

//DeleteSiteAuthenticatorCommand - Delete a Site Authenticator
//RequestType: DELETE
//Input: input *DeleteSiteAuthenticatorCommandInput
func (s *SiteAuthenticatorsService) DeleteSiteAuthenticatorCommand(input *DeleteSiteAuthenticatorCommandInput) (resp *http.Response, err error) {
	path := "/siteAuthenticators/{id}"
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

type DeleteSiteAuthenticatorCommandInput struct {
	Id string
}

//GetSiteAuthenticatorCommand - Get a Site Authenticator
//RequestType: GET
//Input: input *GetSiteAuthenticatorCommandInput
func (s *SiteAuthenticatorsService) GetSiteAuthenticatorCommand(input *GetSiteAuthenticatorCommandInput) (result *SiteAuthenticatorView, resp *http.Response, err error) {
	path := "/siteAuthenticators/{id}"
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

type GetSiteAuthenticatorCommandInput struct {
	Id string
}

//UpdateSiteAuthenticatorCommand - Update a Site Authenticator
//RequestType: PUT
//Input: input *UpdateSiteAuthenticatorCommandInput
func (s *SiteAuthenticatorsService) UpdateSiteAuthenticatorCommand(input *UpdateSiteAuthenticatorCommandInput) (result *SiteAuthenticatorView, resp *http.Response, err error) {
	path := "/siteAuthenticators/{id}"
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

type UpdateSiteAuthenticatorCommandInput struct {
	Body SiteAuthenticatorView
	Id   string
}
