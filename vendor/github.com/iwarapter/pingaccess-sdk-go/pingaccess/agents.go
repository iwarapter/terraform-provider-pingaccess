package pingaccess

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type AgentsService service

type AgentsAPI interface {
	GetAgentsCommand(input *GetAgentsCommandInput) (result *AgentsView, resp *http.Response, err error)
	AddAgentCommand(input *AddAgentCommandInput) (result *AgentView, resp *http.Response, err error)
	GetAgentCertificatesCommand(input *GetAgentCertificatesCommandInput) (result *AgentCertificatesView, resp *http.Response, err error)
	GetAgentCertificateCommand(input *GetAgentCertificateCommandInput) (result *AgentCertificateView, resp *http.Response, err error)
	GetAgentFileCommand(input *GetAgentFileCommandInput) (resp *http.Response, err error)
	DeleteAgentCommand(input *DeleteAgentCommandInput) (resp *http.Response, err error)
	GetAgentCommand(input *GetAgentCommandInput) (result *AgentView, resp *http.Response, err error)
	UpdateAgentCommand(input *UpdateAgentCommandInput) (result *AgentView, resp *http.Response, err error)
}

//GetAgentsCommand - Get all Agents
//RequestType: GET
//Input: input *GetAgentsCommandInput
func (s *AgentsService) GetAgentsCommand(input *GetAgentsCommandInput) (result *AgentsView, resp *http.Response, err error) {
	path := "/agents"
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

type GetAgentsCommandInput struct {
	Page          string
	NumberPerPage string
	Filter        string
	Name          string
	SortKey       string
	Order         string
}

//AddAgentCommand - Add an Agent
//RequestType: POST
//Input: input *AddAgentCommandInput
func (s *AgentsService) AddAgentCommand(input *AddAgentCommandInput) (result *AgentView, resp *http.Response, err error) {
	path := "/agents"
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

type AddAgentCommandInput struct {
	Body AgentView
}

//GetAgentCertificatesCommand - Get all Agent Certificates
//RequestType: GET
//Input: input *GetAgentCertificatesCommandInput
func (s *AgentsService) GetAgentCertificatesCommand(input *GetAgentCertificatesCommandInput) (result *AgentCertificatesView, resp *http.Response, err error) {
	path := "/agents/certificates"
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

type GetAgentCertificatesCommandInput struct {
	Page          string
	NumberPerPage string
	Filter        string
	Alias         string
	SortKey       string
	Order         string
}

//GetAgentCertificateCommand - Get an Agent Certificate
//RequestType: GET
//Input: input *GetAgentCertificateCommandInput
func (s *AgentsService) GetAgentCertificateCommand(input *GetAgentCertificateCommandInput) (result *AgentCertificateView, resp *http.Response, err error) {
	path := "/agents/certificates/{id}"
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

type GetAgentCertificateCommandInput struct {
	Id string
}

//GetAgentFileCommand - Get a configuration file for an Agent
//RequestType: GET
//Input: input *GetAgentFileCommandInput
func (s *AgentsService) GetAgentFileCommand(input *GetAgentFileCommandInput) (resp *http.Response, err error) {
	path := "/agents/{agentId}/config/{sharedSecretId}"
	path = strings.Replace(path, "{agentId}", input.AgentId, -1)

	path = strings.Replace(path, "{sharedSecretId}", input.SharedSecretId, -1)

	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("GET", rel, nil)
	if err != nil {
		return nil, err
	}

	resp, err = s.client.do(req, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil

}

type GetAgentFileCommandInput struct {
	AgentId        string
	SharedSecretId string
}

//DeleteAgentCommand - Delete an Agent
//RequestType: DELETE
//Input: input *DeleteAgentCommandInput
func (s *AgentsService) DeleteAgentCommand(input *DeleteAgentCommandInput) (resp *http.Response, err error) {
	path := "/agents/{id}"
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

type DeleteAgentCommandInput struct {
	Id string
}

//GetAgentCommand - Get an Agent
//RequestType: GET
//Input: input *GetAgentCommandInput
func (s *AgentsService) GetAgentCommand(input *GetAgentCommandInput) (result *AgentView, resp *http.Response, err error) {
	path := "/agents/{id}"
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

type GetAgentCommandInput struct {
	Id string
}

//UpdateAgentCommand - Update an Agent
//RequestType: PUT
//Input: input *UpdateAgentCommandInput
func (s *AgentsService) UpdateAgentCommand(input *UpdateAgentCommandInput) (result *AgentView, resp *http.Response, err error) {
	path := "/agents/{id}"
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

type UpdateAgentCommandInput struct {
	Body AgentView
	Id   string
}
