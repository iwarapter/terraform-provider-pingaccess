package pingaccess

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type HsmProvidersService service

type HsmProvidersAPI interface {
	GetHsmProvidersCommand(input *GetHsmProvidersCommandInput) (result *HsmProviderView, resp *http.Response, err error)
	AddHsmProviderCommand(input *AddHsmProviderCommandInput) (result *HsmProviderView, resp *http.Response, err error)
	GetHsmProviderDescriptorsCommand() (result *DescriptorsView, resp *http.Response, err error)
	DeleteHsmProviderCommand(input *DeleteHsmProviderCommandInput) (resp *http.Response, err error)
	GetHsmProviderCommand(input *GetHsmProviderCommandInput) (result *HsmProviderView, resp *http.Response, err error)
	UpdateHsmProviderCommand(input *UpdateHsmProviderCommandInput) (result *HsmProviderView, resp *http.Response, err error)
}

//GetHsmProvidersCommand - Get all HSM Providers
//RequestType: GET
//Input: input *GetHsmProvidersCommandInput
func (s *HsmProvidersService) GetHsmProvidersCommand(input *GetHsmProvidersCommandInput) (result *HsmProviderView, resp *http.Response, err error) {
	path := "/hsmProviders"
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

type GetHsmProvidersCommandInput struct {
	Page          string
	NumberPerPage string
	Filter        string
	Name          string
	SortKey       string
	Order         string
}

//AddHsmProviderCommand - Create an HSM Provider
//RequestType: POST
//Input: input *AddHsmProviderCommandInput
func (s *HsmProvidersService) AddHsmProviderCommand(input *AddHsmProviderCommandInput) (result *HsmProviderView, resp *http.Response, err error) {
	path := "/hsmProviders"
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

type AddHsmProviderCommandInput struct {
	Body HsmProviderView
}

//GetHsmProviderDescriptorsCommand - Get descriptors for all HSM Providers
//RequestType: GET
//Input:
func (s *HsmProvidersService) GetHsmProviderDescriptorsCommand() (result *DescriptorsView, resp *http.Response, err error) {
	path := "/hsmProviders/descriptors"
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

//DeleteHsmProviderCommand - Delete an HSM Provider
//RequestType: DELETE
//Input: input *DeleteHsmProviderCommandInput
func (s *HsmProvidersService) DeleteHsmProviderCommand(input *DeleteHsmProviderCommandInput) (resp *http.Response, err error) {
	path := "/hsmProviders/{id}"
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

type DeleteHsmProviderCommandInput struct {
	Id string
}

//GetHsmProviderCommand - Get an HSM Provider
//RequestType: GET
//Input: input *GetHsmProviderCommandInput
func (s *HsmProvidersService) GetHsmProviderCommand(input *GetHsmProviderCommandInput) (result *HsmProviderView, resp *http.Response, err error) {
	path := "/hsmProviders/{id}"
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

type GetHsmProviderCommandInput struct {
	Id string
}

//UpdateHsmProviderCommand - Update an HSM Provider
//RequestType: PUT
//Input: input *UpdateHsmProviderCommandInput
func (s *HsmProvidersService) UpdateHsmProviderCommand(input *UpdateHsmProviderCommandInput) (result *HsmProviderView, resp *http.Response, err error) {
	path := "/hsmProviders/{id}"
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

type UpdateHsmProviderCommandInput struct {
	Body HsmProviderView
	Id   string
}
