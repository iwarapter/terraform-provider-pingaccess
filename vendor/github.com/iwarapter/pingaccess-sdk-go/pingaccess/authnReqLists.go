package pingaccess

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type AuthnReqListsService service

type AuthnReqListsAPI interface {
	GetAuthnReqListsCommand(input *GetAuthnReqListsCommandInput) (result *AuthnReqListsView, resp *http.Response, err error)
	AddAuthnReqListCommand(input *AddAuthnReqListCommandInput) (result *AuthnReqListView, resp *http.Response, err error)
	DeleteAuthnReqListCommand(input *DeleteAuthnReqListCommandInput) (resp *http.Response, err error)
	GetAuthnReqListCommand(input *GetAuthnReqListCommandInput) (result *AuthnReqListView, resp *http.Response, err error)
	UpdateAuthnReqListCommand(input *UpdateAuthnReqListCommandInput) (result *AuthnReqListView, resp *http.Response, err error)
}

//GetAuthnReqListsCommand - Get all Authentication Requirement Lists
//RequestType: GET
//Input: input *GetAuthnReqListsCommandInput
func (s *AuthnReqListsService) GetAuthnReqListsCommand(input *GetAuthnReqListsCommandInput) (result *AuthnReqListsView, resp *http.Response, err error) {
	path := "/authnReqLists"
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

type GetAuthnReqListsCommandInput struct {
	Page          string
	NumberPerPage string
	Filter        string
	Name          string
	SortKey       string
	Order         string
}

//AddAuthnReqListCommand - Add an Authentication Requirement List
//RequestType: POST
//Input: input *AddAuthnReqListCommandInput
func (s *AuthnReqListsService) AddAuthnReqListCommand(input *AddAuthnReqListCommandInput) (result *AuthnReqListView, resp *http.Response, err error) {
	path := "/authnReqLists"
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

type AddAuthnReqListCommandInput struct {
	Body AuthnReqListView
}

//DeleteAuthnReqListCommand - Delete an Authentication Requirement List
//RequestType: DELETE
//Input: input *DeleteAuthnReqListCommandInput
func (s *AuthnReqListsService) DeleteAuthnReqListCommand(input *DeleteAuthnReqListCommandInput) (resp *http.Response, err error) {
	path := "/authnReqLists/{id}"
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

type DeleteAuthnReqListCommandInput struct {
	Id string
}

//GetAuthnReqListCommand - Get an Authentication Requirement List
//RequestType: GET
//Input: input *GetAuthnReqListCommandInput
func (s *AuthnReqListsService) GetAuthnReqListCommand(input *GetAuthnReqListCommandInput) (result *AuthnReqListView, resp *http.Response, err error) {
	path := "/authnReqLists/{id}"
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

type GetAuthnReqListCommandInput struct {
	Id string
}

//UpdateAuthnReqListCommand - Update an Authentication Requirement List
//RequestType: PUT
//Input: input *UpdateAuthnReqListCommandInput
func (s *AuthnReqListsService) UpdateAuthnReqListCommand(input *UpdateAuthnReqListCommandInput) (result *AuthnReqListView, resp *http.Response, err error) {
	path := "/authnReqLists/{id}"
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

type UpdateAuthnReqListCommandInput struct {
	Body AuthnReqListView
	Id   string
}
