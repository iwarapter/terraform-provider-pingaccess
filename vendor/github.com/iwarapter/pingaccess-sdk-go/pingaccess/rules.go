package pingaccess

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type RulesService service

type RulesAPI interface {
	GetRulesCommand(input *GetRulesCommandInput) (result *RulesView, resp *http.Response, err error)
	AddRuleCommand(input *AddRuleCommandInput) (result *RuleView, resp *http.Response, err error)
	GetRuleDescriptorsCommand() (result *RuleDescriptorsView, resp *http.Response, err error)
	GetRuleDescriptorCommand(input *GetRuleDescriptorCommandInput) (result *RuleDescriptorView, resp *http.Response, err error)
	DeleteRuleCommand(input *DeleteRuleCommandInput) (resp *http.Response, err error)
	GetRuleCommand(input *GetRuleCommandInput) (result *RuleView, resp *http.Response, err error)
	UpdateRuleCommand(input *UpdateRuleCommandInput) (result *RuleView, resp *http.Response, err error)
}

//GetRulesCommand - Get all Rules
//RequestType: GET
//Input: input *GetRulesCommandInput
func (s *RulesService) GetRulesCommand(input *GetRulesCommandInput) (result *RulesView, resp *http.Response, err error) {
	path := "/rules"
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
func (s *RulesService) AddRuleCommand(input *AddRuleCommandInput) (result *RuleView, resp *http.Response, err error) {
	path := "/rules"
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

type AddRuleCommandInput struct {
	Body RuleView
}

//GetRuleDescriptorsCommand - Get descriptors for all supported Rule types
//RequestType: GET
//Input:
func (s *RulesService) GetRuleDescriptorsCommand() (result *RuleDescriptorsView, resp *http.Response, err error) {
	path := "/rules/descriptors"
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

//GetRuleDescriptorCommand - Get descriptor for a Rule type
//RequestType: GET
//Input: input *GetRuleDescriptorCommandInput
func (s *RulesService) GetRuleDescriptorCommand(input *GetRuleDescriptorCommandInput) (result *RuleDescriptorView, resp *http.Response, err error) {
	path := "/rules/descriptors/{ruleType}"
	path = strings.Replace(path, "{ruleType}", input.RuleType, -1)

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

type GetRuleDescriptorCommandInput struct {
	RuleType string
}

//DeleteRuleCommand - Delete a Rule
//RequestType: DELETE
//Input: input *DeleteRuleCommandInput
func (s *RulesService) DeleteRuleCommand(input *DeleteRuleCommandInput) (resp *http.Response, err error) {
	path := "/rules/{id}"
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

type DeleteRuleCommandInput struct {
	Id string
}

//GetRuleCommand - Get a Rule
//RequestType: GET
//Input: input *GetRuleCommandInput
func (s *RulesService) GetRuleCommand(input *GetRuleCommandInput) (result *RuleView, resp *http.Response, err error) {
	path := "/rules/{id}"
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

type GetRuleCommandInput struct {
	Id string
}

//UpdateRuleCommand - Update a Rule
//RequestType: PUT
//Input: input *UpdateRuleCommandInput
func (s *RulesService) UpdateRuleCommand(input *UpdateRuleCommandInput) (result *RuleView, resp *http.Response, err error) {
	path := "/rules/{id}"
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

type UpdateRuleCommandInput struct {
	Body RuleView
	Id   string
}
