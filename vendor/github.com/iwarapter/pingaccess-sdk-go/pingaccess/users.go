package pingaccess

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type UsersService service

type UsersAPI interface {
	GetUsersCommand(input *GetUsersCommandInput) (result *UsersView, resp *http.Response, err error)
	GetUserCommand(input *GetUserCommandInput) (result *UserView, resp *http.Response, err error)
	UpdateUserCommand(input *UpdateUserCommandInput) (result *UserView, resp *http.Response, err error)
	UpdateUserPasswordCommand(input *UpdateUserPasswordCommandInput) (result *UserPasswordView, resp *http.Response, err error)
}

//GetUsersCommand - Get all Users
//RequestType: GET
//Input: input *GetUsersCommandInput
func (s *UsersService) GetUsersCommand(input *GetUsersCommandInput) (result *UsersView, resp *http.Response, err error) {
	path := "/users"
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
	if input.Username != "" {
		q.Set("username", input.Username)
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

type GetUsersCommandInput struct {
	Page          string
	NumberPerPage string
	Filter        string
	Username      string
	SortKey       string
	Order         string
}

//GetUserCommand - Get a User
//RequestType: GET
//Input: input *GetUserCommandInput
func (s *UsersService) GetUserCommand(input *GetUserCommandInput) (result *UserView, resp *http.Response, err error) {
	path := "/users/{id}"
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

type GetUserCommandInput struct {
	Id string
}

//UpdateUserCommand - Update a User
//RequestType: PUT
//Input: input *UpdateUserCommandInput
func (s *UsersService) UpdateUserCommand(input *UpdateUserCommandInput) (result *UserView, resp *http.Response, err error) {
	path := "/users/{id}"
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

type UpdateUserCommandInput struct {
	Body UserView
	Id   string
}

//UpdateUserPasswordCommand - Update a User's Password
//RequestType: PUT
//Input: input *UpdateUserPasswordCommandInput
func (s *UsersService) UpdateUserPasswordCommand(input *UpdateUserPasswordCommandInput) (result *UserPasswordView, resp *http.Response, err error) {
	path := "/users/{id}/password"
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

type UpdateUserPasswordCommandInput struct {
	Body UserPasswordView
	Id   string
}
