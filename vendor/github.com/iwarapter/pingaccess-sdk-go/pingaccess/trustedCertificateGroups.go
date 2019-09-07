package pingaccess

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type TrustedCertificateGroupsService service

//GetTrustedCertificateGroupsCommand - Get all Trusted Certificate Groups
//RequestType: GET
//Input: input *GetTrustedCertificateGroupsCommandInput
func (s *TrustedCertificateGroupsService) GetTrustedCertificateGroupsCommand(input *GetTrustedCertificateGroupsCommandInput) (result *TrustedCertificateGroupsView, resp *http.Response, err error) {
	path := "/trustedCertificateGroups"
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

type GetTrustedCertificateGroupsCommandInput struct {
	Page          string
	NumberPerPage string
	Filter        string
	Name          string
	SortKey       string
	Order         string
}

//AddTrustedCertificateGroupCommand - Create a Trusted Certificate Group
//RequestType: POST
//Input: input *AddTrustedCertificateGroupCommandInput
func (s *TrustedCertificateGroupsService) AddTrustedCertificateGroupCommand(input *AddTrustedCertificateGroupCommandInput) (result *TrustedCertificateGroupView, resp *http.Response, err error) {
	path := "/trustedCertificateGroups"
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

type AddTrustedCertificateGroupCommandInput struct {
	Body TrustedCertificateGroupView
}

//DeleteTrustedCertificateGroupCommand - Delete a Trusted Certificate Group
//RequestType: DELETE
//Input: input *DeleteTrustedCertificateGroupCommandInput
func (s *TrustedCertificateGroupsService) DeleteTrustedCertificateGroupCommand(input *DeleteTrustedCertificateGroupCommandInput) (resp *http.Response, err error) {
	path := "/trustedCertificateGroups/{id}"
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

type DeleteTrustedCertificateGroupCommandInput struct {
	Id string
}

//GetTrustedCertificateGroupCommand - Get a Trusted Certificate Group
//RequestType: GET
//Input: input *GetTrustedCertificateGroupCommandInput
func (s *TrustedCertificateGroupsService) GetTrustedCertificateGroupCommand(input *GetTrustedCertificateGroupCommandInput) (result *TrustedCertificateGroupView, resp *http.Response, err error) {
	path := "/trustedCertificateGroups/{id}"
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

type GetTrustedCertificateGroupCommandInput struct {
	Id string
}

//UpdateTrustedCertificateGroupCommand - Update a TrustedCertificateGroup
//RequestType: PUT
//Input: input *UpdateTrustedCertificateGroupCommandInput
func (s *TrustedCertificateGroupsService) UpdateTrustedCertificateGroupCommand(input *UpdateTrustedCertificateGroupCommandInput) (result *TrustedCertificateGroupView, resp *http.Response, err error) {
	path := "/trustedCertificateGroups/{id}"
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

type UpdateTrustedCertificateGroupCommandInput struct {
	Body TrustedCertificateGroupView
	Id   string
}
