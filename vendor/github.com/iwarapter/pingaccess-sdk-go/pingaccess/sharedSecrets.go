package pingaccess

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type SharedSecretsService service

//GetSharedSecretsCommand - Get all Shared Secrets
//RequestType: GET
//Input: input *GetSharedSecretsCommandInput
func (s *SharedSecretsService) GetSharedSecretsCommand(input *GetSharedSecretsCommandInput) (result *SharedSecretsView, resp *http.Response, err error) {
	path := "/sharedSecrets"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	q := rel.Query()
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

type GetSharedSecretsCommandInput struct {
	SortKey string
	Order   string
}

//AddSharedSecretCommand - Create a Shared Secret
//RequestType: POST
//Input: input *AddSharedSecretCommandInput
func (s *SharedSecretsService) AddSharedSecretCommand(input *AddSharedSecretCommandInput) (result *SharedSecretView, resp *http.Response, err error) {
	path := "/sharedSecrets"
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

type AddSharedSecretCommandInput struct {
	Body SharedSecretView
}

//DeleteSharedSecretCommand - Delete a Shared Secret
//RequestType: DELETE
//Input: input *DeleteSharedSecretCommandInput
func (s *SharedSecretsService) DeleteSharedSecretCommand(input *DeleteSharedSecretCommandInput) (resp *http.Response, err error) {
	path := "/sharedSecrets/{id}"
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

type DeleteSharedSecretCommandInput struct {
	Id string
}

//GetSharedSecretCommand - Get a Shared Secret
//RequestType: GET
//Input: input *GetSharedSecretCommandInput
func (s *SharedSecretsService) GetSharedSecretCommand(input *GetSharedSecretCommandInput) (result *SharedSecretView, resp *http.Response, err error) {
	path := "/sharedSecrets/{id}"
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

type GetSharedSecretCommandInput struct {
	Id string
}
