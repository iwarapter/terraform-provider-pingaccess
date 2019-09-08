package pingaccess

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type IdentityMappingsService service

//GetIdentityMappingsCommand - Get all Identity Mappings
//RequestType: GET
//Input: input *GetIdentityMappingsCommandInput
func (s *IdentityMappingsService) GetIdentityMappingsCommand(input *GetIdentityMappingsCommandInput) (result *IdentityMappingsView, resp *http.Response, err error) {
	path := "/identityMappings"
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

type GetIdentityMappingsCommandInput struct {
	Page          string
	NumberPerPage string
	Filter        string
	Name          string
	SortKey       string
	Order         string
}

//AddIdentityMappingCommand - Create an Identity Mapping
//RequestType: POST
//Input: input *AddIdentityMappingCommandInput
func (s *IdentityMappingsService) AddIdentityMappingCommand(input *AddIdentityMappingCommandInput) (result *IdentityMappingView, resp *http.Response, err error) {
	path := "/identityMappings"
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

type AddIdentityMappingCommandInput struct {
	Body IdentityMappingView
}

//GetIdentityMappingDescriptorsCommand - Get descriptors for all supported Identity Mapping types
//RequestType: GET
//Input:
func (s *IdentityMappingsService) GetIdentityMappingDescriptorsCommand() (result *DescriptorsView, resp *http.Response, err error) {
	path := "/identityMappings/descriptors"
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

//GetIdentityMappingDescriptorCommand - Get descriptor for an Identity Mapping type
//RequestType: GET
//Input: input *GetIdentityMappingDescriptorCommandInput
func (s *IdentityMappingsService) GetIdentityMappingDescriptorCommand(input *GetIdentityMappingDescriptorCommandInput) (result *DescriptorView, resp *http.Response, err error) {
	path := "/identityMappings/descriptors/{identityMappingType}"
	path = strings.Replace(path, "{identityMappingType}", input.IdentityMappingType, -1)

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

type GetIdentityMappingDescriptorCommandInput struct {
	IdentityMappingType string
}

//DeleteIdentityMappingCommand - Delete an Identity Mapping
//RequestType: DELETE
//Input: input *DeleteIdentityMappingCommandInput
func (s *IdentityMappingsService) DeleteIdentityMappingCommand(input *DeleteIdentityMappingCommandInput) (resp *http.Response, err error) {
	path := "/identityMappings/{id}"
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

type DeleteIdentityMappingCommandInput struct {
	Id string
}

//GetIdentityMappingCommand - Get an Identity Mapping
//RequestType: GET
//Input: input *GetIdentityMappingCommandInput
func (s *IdentityMappingsService) GetIdentityMappingCommand(input *GetIdentityMappingCommandInput) (result *IdentityMappingView, resp *http.Response, err error) {
	path := "/identityMappings/{id}"
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

type GetIdentityMappingCommandInput struct {
	Id string
}

//UpdateIdentityMappingCommand - Update an Identity Mapping
//RequestType: PUT
//Input: input *UpdateIdentityMappingCommandInput
func (s *IdentityMappingsService) UpdateIdentityMappingCommand(input *UpdateIdentityMappingCommandInput) (result *IdentityMappingView, resp *http.Response, err error) {
	path := "/identityMappings/{id}"
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

type UpdateIdentityMappingCommandInput struct {
	Body IdentityMappingView
	Id   string
}
