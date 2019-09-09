package pingaccess

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type AccessTokenValidatorsService service

//GetAccessTokenValidatorsCommand - Get all Access Token Validators
//RequestType: GET
//Input: input *GetAccessTokenValidatorsCommandInput
func (s *AccessTokenValidatorsService) GetAccessTokenValidatorsCommand(input *GetAccessTokenValidatorsCommandInput) (result *AccessTokenValidatorsView, resp *http.Response, err error) {
	path := "/accessTokenValidators"
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

type GetAccessTokenValidatorsCommandInput struct {
	Page          string
	NumberPerPage string
	Filter        string
	Name          string
	SortKey       string
	Order         string
}

//AddAccessTokenValidatorCommand - Create an Access Token Validator
//RequestType: POST
//Input: input *AddAccessTokenValidatorCommandInput
func (s *AccessTokenValidatorsService) AddAccessTokenValidatorCommand(input *AddAccessTokenValidatorCommandInput) (result *AccessTokenValidatorView, resp *http.Response, err error) {
	path := "/accessTokenValidators"
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

type AddAccessTokenValidatorCommandInput struct {
	Body AccessTokenValidatorView
}

//GetAccessTokenValidatorDescriptorsCommand - Get descriptors for all Access Token Validators
//RequestType: GET
//Input:
func (s *AccessTokenValidatorsService) GetAccessTokenValidatorDescriptorsCommand() (result *DescriptorsView, resp *http.Response, err error) {
	path := "/accessTokenValidators/descriptors"
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

//DeleteAccessTokenValidatorCommand - Delete a Access Token Validator
//RequestType: DELETE
//Input: input *DeleteAccessTokenValidatorCommandInput
func (s *AccessTokenValidatorsService) DeleteAccessTokenValidatorCommand(input *DeleteAccessTokenValidatorCommandInput) (resp *http.Response, err error) {
	path := "/accessTokenValidators/{id}"
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

type DeleteAccessTokenValidatorCommandInput struct {
	Id string
}

//GetAccessTokenValidatorCommand - Get an Access Token Validator
//RequestType: GET
//Input: input *GetAccessTokenValidatorCommandInput
func (s *AccessTokenValidatorsService) GetAccessTokenValidatorCommand(input *GetAccessTokenValidatorCommandInput) (result *AccessTokenValidatorView, resp *http.Response, err error) {
	path := "/accessTokenValidators/{id}"
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

type GetAccessTokenValidatorCommandInput struct {
	Id string
}

//UpdateAccessTokenValidatorCommand - Update an Access Token Validator
//RequestType: PUT
//Input: input *UpdateAccessTokenValidatorCommandInput
func (s *AccessTokenValidatorsService) UpdateAccessTokenValidatorCommand(input *UpdateAccessTokenValidatorCommandInput) (result *AccessTokenValidatorView, resp *http.Response, err error) {
	path := "/accessTokenValidators/{id}"
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

type UpdateAccessTokenValidatorCommandInput struct {
	Body AccessTokenValidatorView
	Id   string
}
