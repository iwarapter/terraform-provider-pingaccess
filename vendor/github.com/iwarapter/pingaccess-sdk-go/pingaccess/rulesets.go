package pingaccess

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type RulesetsService service

//GetRuleSetsCommand - Get all Rule Sets
//RequestType: GET
//Input: input *GetRuleSetsCommandInput
func (s *RulesetsService) GetRuleSetsCommand(input *GetRuleSetsCommandInput) (result *RuleSetsView, resp *http.Response, err error) {
	path := "/rulesets"
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
func (s *RulesetsService) AddRuleSetCommand(input *AddRuleSetCommandInput) (result *RuleSetView, resp *http.Response, err error) {
	path := "/rulesets"
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

type AddRuleSetCommandInput struct {
	Body RuleSetView
}

//GetRuleSetElementTypesCommand - Get all Rule Set Element Types
//RequestType: GET
//Input:
func (s *RulesetsService) GetRuleSetElementTypesCommand() (result *RuleSetElementTypesView, resp *http.Response, err error) {
	path := "/rulesets/elementTypes"
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

//GetRuleSetSuccessCriteriaCommand - Get all Success Criteria
//RequestType: GET
//Input:
func (s *RulesetsService) GetRuleSetSuccessCriteriaCommand() (result *RuleSetSuccessCriteriaView, resp *http.Response, err error) {
	path := "/rulesets/successCriteria"
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

//DeleteRuleSetCommand - Delete a Rule Set
//RequestType: DELETE
//Input: input *DeleteRuleSetCommandInput
func (s *RulesetsService) DeleteRuleSetCommand(input *DeleteRuleSetCommandInput) (resp *http.Response, err error) {
	path := "/rulesets/{id}"
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

type DeleteRuleSetCommandInput struct {
	Id string
}

//GetRuleSetCommand - Get a Rule Set
//RequestType: GET
//Input: input *GetRuleSetCommandInput
func (s *RulesetsService) GetRuleSetCommand(input *GetRuleSetCommandInput) (result *RuleSetView, resp *http.Response, err error) {
	path := "/rulesets/{id}"
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

type GetRuleSetCommandInput struct {
	Id string
}

//UpdateRuleSetCommand - Update a Rule Set
//RequestType: PUT
//Input: input *UpdateRuleSetCommandInput
func (s *RulesetsService) UpdateRuleSetCommand(input *UpdateRuleSetCommandInput) (result *RuleSetView, resp *http.Response, err error) {
	path := "/rulesets/{id}"
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

type UpdateRuleSetCommandInput struct {
	Body RuleSetView
	Id   string
}
