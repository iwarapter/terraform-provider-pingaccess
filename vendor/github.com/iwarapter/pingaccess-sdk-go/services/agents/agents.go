package agents

import (
	"net/http"
	"strings"

	"github.com/iwarapter/pingaccess-sdk-go/pingaccess"
	"github.com/iwarapter/pingaccess-sdk-go/pingaccess/client"
	"github.com/iwarapter/pingaccess-sdk-go/pingaccess/client/metadata"
	"github.com/iwarapter/pingaccess-sdk-go/pingaccess/config"
	"github.com/iwarapter/pingaccess-sdk-go/pingaccess/models"
	"github.com/iwarapter/pingaccess-sdk-go/pingaccess/request"
)

const (
	// ServiceName - The name of service.
	ServiceName = "Agents"
)

//AgentsService provides the API operations for making requests to
// Agents endpoint.
type AgentsService struct {
	*client.Client
}

//New createa a new instance of the AgentsService client.
//
// Example:
//   cfg := config.NewConfig().WithUsername("Administrator").WithPassword("2Access").WithEndpoint(paURL)
//
//   //Create a AgentsService from the configuration
//   svc := agents.New(cfg)
//
func New(cfg *config.Config) *AgentsService {

	return &AgentsService{Client: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingaccess.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a Agents operation
func (s *AgentsService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := s.NewRequest(op, params, data)

	return req
}

//GetAgentsCommand - Get all Agents
//RequestType: GET
//Input: input *GetAgentsCommandInput
func (s *AgentsService) GetAgentsCommand(input *GetAgentsCommandInput) (output *models.AgentsView, resp *http.Response, err error) {
	path := "/agents"
	op := &request.Operation{
		Name:       "GetAgentsCommand",
		HTTPMethod: "GET",
		HTTPPath:   path,
		QueryParams: map[string]string{
			"page":          input.Page,
			"numberPerPage": input.NumberPerPage,
			"filter":        input.Filter,
			"name":          input.Name,
			"sortKey":       input.SortKey,
			"order":         input.Order,
		},
	}
	output = &models.AgentsView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// GetAgentsCommandInput - Inputs for GetAgentsCommand
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
func (s *AgentsService) AddAgentCommand(input *AddAgentCommandInput) (output *models.AgentView, resp *http.Response, err error) {
	path := "/agents"
	op := &request.Operation{
		Name:        "AddAgentCommand",
		HTTPMethod:  "POST",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.AgentView{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// AddAgentCommandInput - Inputs for AddAgentCommand
type AddAgentCommandInput struct {
	Body models.AgentView
}

//GetAgentCertificatesCommand - Get all Agent Certificates
//RequestType: GET
//Input: input *GetAgentCertificatesCommandInput
func (s *AgentsService) GetAgentCertificatesCommand(input *GetAgentCertificatesCommandInput) (output *models.AgentCertificatesView, resp *http.Response, err error) {
	path := "/agents/certificates"
	op := &request.Operation{
		Name:       "GetAgentCertificatesCommand",
		HTTPMethod: "GET",
		HTTPPath:   path,
		QueryParams: map[string]string{
			"page":          input.Page,
			"numberPerPage": input.NumberPerPage,
			"filter":        input.Filter,
			"alias":         input.Alias,
			"sortKey":       input.SortKey,
			"order":         input.Order,
		},
	}
	output = &models.AgentCertificatesView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// GetAgentCertificatesCommandInput - Inputs for GetAgentCertificatesCommand
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
func (s *AgentsService) GetAgentCertificateCommand(input *GetAgentCertificateCommandInput) (output *models.AgentCertificateView, resp *http.Response, err error) {
	path := "/agents/certificates/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "GetAgentCertificateCommand",
		HTTPMethod:  "GET",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.AgentCertificateView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// GetAgentCertificateCommandInput - Inputs for GetAgentCertificateCommand
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

	op := &request.Operation{
		Name:        "GetAgentFileCommand",
		HTTPMethod:  "GET",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}

	req := s.newRequest(op, nil, nil)

	if req.Send() == nil {
		return req.HTTPResponse, nil
	}
	return req.HTTPResponse, req.Error
}

// GetAgentFileCommandInput - Inputs for GetAgentFileCommand
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

	op := &request.Operation{
		Name:        "DeleteAgentCommand",
		HTTPMethod:  "DELETE",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}

	req := s.newRequest(op, nil, nil)

	if req.Send() == nil {
		return req.HTTPResponse, nil
	}
	return req.HTTPResponse, req.Error
}

// DeleteAgentCommandInput - Inputs for DeleteAgentCommand
type DeleteAgentCommandInput struct {
	Id string
}

//GetAgentCommand - Get an Agent
//RequestType: GET
//Input: input *GetAgentCommandInput
func (s *AgentsService) GetAgentCommand(input *GetAgentCommandInput) (output *models.AgentView, resp *http.Response, err error) {
	path := "/agents/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "GetAgentCommand",
		HTTPMethod:  "GET",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.AgentView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// GetAgentCommandInput - Inputs for GetAgentCommand
type GetAgentCommandInput struct {
	Id string
}

//UpdateAgentCommand - Update an Agent
//RequestType: PUT
//Input: input *UpdateAgentCommandInput
func (s *AgentsService) UpdateAgentCommand(input *UpdateAgentCommandInput) (output *models.AgentView, resp *http.Response, err error) {
	path := "/agents/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "UpdateAgentCommand",
		HTTPMethod:  "PUT",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.AgentView{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// UpdateAgentCommandInput - Inputs for UpdateAgentCommand
type UpdateAgentCommandInput struct {
	Body models.AgentView
	Id   string
}
