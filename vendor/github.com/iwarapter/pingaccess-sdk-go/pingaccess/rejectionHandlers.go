package pingaccess

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type RejectionHandlersService service

//GetRejectionHandlersCommand - Get all Rejection Handlers
//RequestType: GET
//Input: input *GetRejectionHandlersCommandInput
func (s *RejectionHandlersService) GetRejectionHandlersCommand(input *GetRejectionHandlersCommandInput) (result *RejectionHandlersView, resp *http.Response, err error) {
	path := "/rejectionHandlers"
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

type GetRejectionHandlersCommandInput struct {
	Page          string
	NumberPerPage string
	Filter        string
	Name          string
	SortKey       string
	Order         string
}

//AddRejectionHandlerCommand - Create a Rejection Handler
//RequestType: POST
//Input: input *AddRejectionHandlerCommandInput
func (s *RejectionHandlersService) AddRejectionHandlerCommand(input *AddRejectionHandlerCommandInput) (result *RejectionHandlerView, resp *http.Response, err error) {
	path := "/rejectionHandlers"
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

type AddRejectionHandlerCommandInput struct {
	Body RejectionHandlerView
}

//GetRejectionHandlerDescriptorsCommand - Get descriptors for all supported Rejection Handler types
//RequestType: GET
//Input:
func (s *RejectionHandlersService) GetRejectionHandlerDescriptorsCommand() (result *DescriptorsView, resp *http.Response, err error) {
	path := "/rejectionHandlers/descriptors"
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

//GetRejecitonHandlerDescriptorCommand - Get descriptor for a Rejection Handler type
//RequestType: GET
//Input: input *GetRejecitonHandlerDescriptorCommandInput
func (s *RejectionHandlersService) GetRejecitonHandlerDescriptorCommand(input *GetRejecitonHandlerDescriptorCommandInput) (result *DescriptorView, resp *http.Response, err error) {
	path := "/rejectionHandlers/descriptors/{rejectionHandlerType}"
	path = strings.Replace(path, "{rejectionHandlerType}", input.RejectionHandlerType, -1)

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

type GetRejecitonHandlerDescriptorCommandInput struct {
	RejectionHandlerType string
}

//DeleteRejectionHandlerCommand - Delete a Rejection Handler
//RequestType: DELETE
//Input: input *DeleteRejectionHandlerCommandInput
func (s *RejectionHandlersService) DeleteRejectionHandlerCommand(input *DeleteRejectionHandlerCommandInput) (resp *http.Response, err error) {
	path := "/rejectionHandlers/{id}"
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

type DeleteRejectionHandlerCommandInput struct {
	Id string
}

//GetRejectionHandlerCommand - Get a Rejection Handler
//RequestType: GET
//Input: input *GetRejectionHandlerCommandInput
func (s *RejectionHandlersService) GetRejectionHandlerCommand(input *GetRejectionHandlerCommandInput) (result *RejectionHandlerView, resp *http.Response, err error) {
	path := "/rejectionHandlers/{id}"
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

type GetRejectionHandlerCommandInput struct {
	Id string
}

//UpdateRejectionHandlerCommand - Update a Rejection Handler
//RequestType: PUT
//Input: input *UpdateRejectionHandlerCommandInput
func (s *RejectionHandlersService) UpdateRejectionHandlerCommand(input *UpdateRejectionHandlerCommandInput) (result *RejectionHandlerView, resp *http.Response, err error) {
	path := "/rejectionHandlers/{id}"
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

type UpdateRejectionHandlerCommandInput struct {
	Body RejectionHandlerView
	Id   string
}
