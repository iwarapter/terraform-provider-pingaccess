package pingaccess

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type EnginesService service

//GetEnginesCommand - Get all Engines
//RequestType: GET
//Input: input *GetEnginesCommandInput
func (s *EnginesService) GetEnginesCommand(input *GetEnginesCommandInput) (result *EnginesView, resp *http.Response, err error) {
	path := "/engines"
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

type GetEnginesCommandInput struct {
	Page          string
	NumberPerPage string
	Filter        string
	Name          string
	SortKey       string
	Order         string
}

//AddEngineCommand - Add an Engine
//RequestType: POST
//Input: input *AddEngineCommandInput
func (s *EnginesService) AddEngineCommand(input *AddEngineCommandInput) (result *EngineView, resp *http.Response, err error) {
	path := "/engines"
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

type AddEngineCommandInput struct {
	Body EngineView
}

//GetEngineCertificatesCommand - Get all Engine Certificates
//RequestType: GET
//Input: input *GetEngineCertificatesCommandInput
func (s *EnginesService) GetEngineCertificatesCommand(input *GetEngineCertificatesCommandInput) (result *EngineCertificateView, resp *http.Response, err error) {
	path := "/engines/certificates"
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
	if input.Alias != "" {
		q.Set("alias", input.Alias)
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

type GetEngineCertificatesCommandInput struct {
	Page          string
	NumberPerPage string
	Filter        string
	Alias         string
	SortKey       string
	Order         string
}

//GetEngineCertificateCommand - Get an Engine Certificate
//RequestType: GET
//Input: input *GetEngineCertificateCommandInput
func (s *EnginesService) GetEngineCertificateCommand(input *GetEngineCertificateCommandInput) (result *EngineCertificateView, resp *http.Response, err error) {
	path := "/engines/certificates/{id}"
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

type GetEngineCertificateCommandInput struct {
	Id string
}

//GetEngineStatusCommand - Get health status of all Engines
//RequestType: GET
//Input:
func (s *EnginesService) GetEngineStatusCommand() (result *EngineHealthStatusView, resp *http.Response, err error) {
	path := "/engines/status"
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

//DeleteEngineCommand - Delete an Engine
//RequestType: DELETE
//Input: input *DeleteEngineCommandInput
func (s *EnginesService) DeleteEngineCommand(input *DeleteEngineCommandInput) (resp *http.Response, err error) {
	path := "/engines/{id}"
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

type DeleteEngineCommandInput struct {
	Id string
}

//GetEngineCommand - Get an Engine
//RequestType: GET
//Input: input *GetEngineCommandInput
func (s *EnginesService) GetEngineCommand(input *GetEngineCommandInput) (result *EngineView, resp *http.Response, err error) {
	path := "/engines/{id}"
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

type GetEngineCommandInput struct {
	Id string
}

//UpdateEngineCommand - Update an Engine
//RequestType: PUT
//Input: input *UpdateEngineCommandInput
func (s *EnginesService) UpdateEngineCommand(input *UpdateEngineCommandInput) (result *EngineView, resp *http.Response, err error) {
	path := "/engines/{id}"
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

type UpdateEngineCommandInput struct {
	Body EngineView
	Id   string
}

//GetEngineConfigFileCommand - Get configuration file for an Engine
//RequestType: POST
//Input: input *GetEngineConfigFileCommandInput
func (s *EnginesService) GetEngineConfigFileCommand(input *GetEngineConfigFileCommandInput) (resp *http.Response, err error) {
	path := "/engines/{id}/config"
	path = strings.Replace(path, "{id}", input.Id, -1)

	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("POST", rel, nil)
	if err != nil {
		return nil, err
	}

	resp, err = s.client.do(req, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil

}

type GetEngineConfigFileCommandInput struct {
	Id string
}
