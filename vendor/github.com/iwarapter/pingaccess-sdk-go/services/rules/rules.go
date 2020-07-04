package rules

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
	ServiceName = "Rules"
)

//RulesService provides the API operations for making requests to
// Rules endpoint.
type RulesService struct {
	*client.Client
}

//New createa a new instance of the RulesService client.
//
// Example:
//   cfg := config.NewConfig().WithUsername("Administrator").WithPassword("2FederateM0re").WithEndpoint(paURL.String())
//
//   //Create a RulesService from the configuration
//   svc := rules.New(cfg)
//
func New(cfg *config.Config) *RulesService {

	return &RulesService{Client: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingaccess.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a Rules operation
func (s *RulesService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := s.NewRequest(op, params, data)

	return req
}

//GetRulesCommand - Get all Rules
//RequestType: GET
//Input: input *GetRulesCommandInput
func (s *RulesService) GetRulesCommand(input *GetRulesCommandInput) (output *models.RulesView, resp *http.Response, err error) {
	path := "/rules"
	op := &request.Operation{
		Name:       "GetRulesCommand",
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
	output = &models.RulesView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// GetRulesCommandInput - Inputs for GetRulesCommand
type GetRulesCommandInput struct {
	Page          string
	NumberPerPage string
	Filter        string
	Name          string
	SortKey       string
	Order         string
}

//AddRuleCommand - Add a Rule
//RequestType: POST
//Input: input *AddRuleCommandInput
func (s *RulesService) AddRuleCommand(input *AddRuleCommandInput) (output *models.RuleView, resp *http.Response, err error) {
	path := "/rules"
	op := &request.Operation{
		Name:        "AddRuleCommand",
		HTTPMethod:  "POST",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.RuleView{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// AddRuleCommandInput - Inputs for AddRuleCommand
type AddRuleCommandInput struct {
	Body models.RuleView
}

//GetRuleDescriptorsCommand - Get descriptors for all supported Rule types
//RequestType: GET
//Input:
func (s *RulesService) GetRuleDescriptorsCommand() (output *models.RuleDescriptorsView, resp *http.Response, err error) {
	path := "/rules/descriptors"
	op := &request.Operation{
		Name:       "GetRuleDescriptorsCommand",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.RuleDescriptorsView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetRuleDescriptorCommand - Get descriptor for a Rule type
//RequestType: GET
//Input: input *GetRuleDescriptorCommandInput
func (s *RulesService) GetRuleDescriptorCommand(input *GetRuleDescriptorCommandInput) (output *models.RuleDescriptorView, resp *http.Response, err error) {
	path := "/rules/descriptors/{ruleType}"
	path = strings.Replace(path, "{ruleType}", input.RuleType, -1)

	op := &request.Operation{
		Name:        "GetRuleDescriptorCommand",
		HTTPMethod:  "GET",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.RuleDescriptorView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// GetRuleDescriptorCommandInput - Inputs for GetRuleDescriptorCommand
type GetRuleDescriptorCommandInput struct {
	RuleType string
}

//DeleteRuleCommand - Delete a Rule
//RequestType: DELETE
//Input: input *DeleteRuleCommandInput
func (s *RulesService) DeleteRuleCommand(input *DeleteRuleCommandInput) (resp *http.Response, err error) {
	path := "/rules/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "DeleteRuleCommand",
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

// DeleteRuleCommandInput - Inputs for DeleteRuleCommand
type DeleteRuleCommandInput struct {
	Id string
}

//GetRuleCommand - Get a Rule
//RequestType: GET
//Input: input *GetRuleCommandInput
func (s *RulesService) GetRuleCommand(input *GetRuleCommandInput) (output *models.RuleView, resp *http.Response, err error) {
	path := "/rules/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "GetRuleCommand",
		HTTPMethod:  "GET",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.RuleView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// GetRuleCommandInput - Inputs for GetRuleCommand
type GetRuleCommandInput struct {
	Id string
}

//UpdateRuleCommand - Update a Rule
//RequestType: PUT
//Input: input *UpdateRuleCommandInput
func (s *RulesService) UpdateRuleCommand(input *UpdateRuleCommandInput) (output *models.RuleView, resp *http.Response, err error) {
	path := "/rules/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "UpdateRuleCommand",
		HTTPMethod:  "PUT",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.RuleView{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// UpdateRuleCommandInput - Inputs for UpdateRuleCommand
type UpdateRuleCommandInput struct {
	Body models.RuleView
	Id   string
}
