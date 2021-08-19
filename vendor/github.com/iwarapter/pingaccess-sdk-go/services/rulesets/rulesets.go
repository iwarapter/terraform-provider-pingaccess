package rulesets

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
	ServiceName = "Rulesets"
)

//RulesetsService provides the API operations for making requests to
// Rulesets endpoint.
type RulesetsService struct {
	*client.Client
}

//New createa a new instance of the RulesetsService client.
//
// Example:
//   cfg := config.NewConfig().WithUsername("Administrator").WithPassword("2Access").WithEndpoint(paURL)
//
//   //Create a RulesetsService from the configuration
//   svc := rulesets.New(cfg)
//
func New(cfg *config.Config) *RulesetsService {

	return &RulesetsService{Client: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingaccess.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a Rulesets operation
func (s *RulesetsService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := s.NewRequest(op, params, data)

	return req
}

//GetRuleSetsCommand - Get all Rule Sets
//RequestType: GET
//Input: input *GetRuleSetsCommandInput
func (s *RulesetsService) GetRuleSetsCommand(input *GetRuleSetsCommandInput) (output *models.RuleSetsView, resp *http.Response, err error) {
	path := "/rulesets"
	op := &request.Operation{
		Name:       "GetRuleSetsCommand",
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
	output = &models.RuleSetsView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// GetRuleSetsCommandInput - Inputs for GetRuleSetsCommand
type GetRuleSetsCommandInput struct {
	Page          string
	NumberPerPage string
	Filter        string
	Name          string
	SortKey       string
	Order         string
}

//AddRuleSetCommand - Add a Rule Set
//RequestType: POST
//Input: input *AddRuleSetCommandInput
func (s *RulesetsService) AddRuleSetCommand(input *AddRuleSetCommandInput) (output *models.RuleSetView, resp *http.Response, err error) {
	path := "/rulesets"
	op := &request.Operation{
		Name:        "AddRuleSetCommand",
		HTTPMethod:  "POST",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.RuleSetView{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// AddRuleSetCommandInput - Inputs for AddRuleSetCommand
type AddRuleSetCommandInput struct {
	Body models.RuleSetView
}

//GetRuleSetElementTypesCommand - Get all Rule Set Element Types
//RequestType: GET
//Input:
func (s *RulesetsService) GetRuleSetElementTypesCommand() (output *models.RuleSetElementTypesView, resp *http.Response, err error) {
	path := "/rulesets/elementTypes"
	op := &request.Operation{
		Name:       "GetRuleSetElementTypesCommand",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.RuleSetElementTypesView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetRuleSetSuccessCriteriaCommand - Get all Success Criteria
//RequestType: GET
//Input:
func (s *RulesetsService) GetRuleSetSuccessCriteriaCommand() (output *models.RuleSetSuccessCriteriaView, resp *http.Response, err error) {
	path := "/rulesets/successCriteria"
	op := &request.Operation{
		Name:       "GetRuleSetSuccessCriteriaCommand",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.RuleSetSuccessCriteriaView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//DeleteRuleSetCommand - Delete a Rule Set
//RequestType: DELETE
//Input: input *DeleteRuleSetCommandInput
func (s *RulesetsService) DeleteRuleSetCommand(input *DeleteRuleSetCommandInput) (resp *http.Response, err error) {
	path := "/rulesets/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "DeleteRuleSetCommand",
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

// DeleteRuleSetCommandInput - Inputs for DeleteRuleSetCommand
type DeleteRuleSetCommandInput struct {
	Id string
}

//GetRuleSetCommand - Get a Rule Set
//RequestType: GET
//Input: input *GetRuleSetCommandInput
func (s *RulesetsService) GetRuleSetCommand(input *GetRuleSetCommandInput) (output *models.RuleSetView, resp *http.Response, err error) {
	path := "/rulesets/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "GetRuleSetCommand",
		HTTPMethod:  "GET",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.RuleSetView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// GetRuleSetCommandInput - Inputs for GetRuleSetCommand
type GetRuleSetCommandInput struct {
	Id string
}

//UpdateRuleSetCommand - Update a Rule Set
//RequestType: PUT
//Input: input *UpdateRuleSetCommandInput
func (s *RulesetsService) UpdateRuleSetCommand(input *UpdateRuleSetCommandInput) (output *models.RuleSetView, resp *http.Response, err error) {
	path := "/rulesets/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "UpdateRuleSetCommand",
		HTTPMethod:  "PUT",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.RuleSetView{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// UpdateRuleSetCommandInput - Inputs for UpdateRuleSetCommand
type UpdateRuleSetCommandInput struct {
	Body models.RuleSetView
	Id   string
}
