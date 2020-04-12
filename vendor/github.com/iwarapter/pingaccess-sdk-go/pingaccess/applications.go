package pingaccess

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type ApplicationsService service

//GetApplicationsCommand - Get all Applications
//RequestType: GET
//Input: input *GetApplicationsCommandInput
func (s *ApplicationsService) GetApplicationsCommand(input *GetApplicationsCommandInput) (result *ApplicationsView, resp *http.Response, err error) {
	path := "/applications"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	q := rel.Query()
	if input.Page != "" {
		q.Set("page", input.Page)
	}
	if input.SiteId != "" {
		q.Set("siteId", input.SiteId)
	}
	if input.NumberPerPage != "" {
		q.Set("numberPerPage", input.NumberPerPage)
	}
	if input.AgentId != "" {
		q.Set("agentId", input.AgentId)
	}
	if input.VirtualHostId != "" {
		q.Set("virtualHostId", input.VirtualHostId)
	}
	if input.RuleId != "" {
		q.Set("ruleId", input.RuleId)
	}
	if input.RulesetId != "" {
		q.Set("rulesetId", input.RulesetId)
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

type GetApplicationsCommandInput struct {
	Page          string
	SiteId        string
	NumberPerPage string
	AgentId       string
	VirtualHostId string
	RuleId        string
	RulesetId     string
	Filter        string
	Name          string
	SortKey       string
	Order         string
}

//AddApplicationCommand - Add an Application
//RequestType: POST
//Input: input *AddApplicationCommandInput
func (s *ApplicationsService) AddApplicationCommand(input *AddApplicationCommandInput) (result *ApplicationView, resp *http.Response, err error) {
	path := "/applications"
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

type AddApplicationCommandInput struct {
	Body ApplicationView
}

//DeleteReservedApplicationCommand - Resets the Reserved Application configuration to default values
//RequestType: DELETE
//Input:
func (s *ApplicationsService) DeleteReservedApplicationCommand() (resp *http.Response, err error) {
	path := "/applications/reserved"
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

//GetReservedApplicationCommand - Get Reserved Application configuration
//RequestType: GET
//Input:
func (s *ApplicationsService) GetReservedApplicationCommand() (result *ReservedApplicationView, resp *http.Response, err error) {
	path := "/applications/reserved"
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

//UpdateReservedApplicationCommand - Update Reserved Application configuration
//RequestType: PUT
//Input: input *UpdateReservedApplicationCommandInput
func (s *ApplicationsService) UpdateReservedApplicationCommand(input *UpdateReservedApplicationCommandInput) (result *ReservedApplicationView, resp *http.Response, err error) {
	path := "/applications/reserved"
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

type UpdateReservedApplicationCommandInput struct {
	Body ReservedApplicationView
}

//GetResourcesCommand - Get all Resources
//RequestType: GET
//Input: input *GetResourcesCommandInput
func (s *ApplicationsService) GetResourcesCommand(input *GetResourcesCommandInput) (result *ResourcesView, resp *http.Response, err error) {
	path := "/applications/resources"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	q := rel.Query()
	if input.Page != "" {
		q.Set("page", input.Page)
	}
	if input.NumberPerPage != "" {
		q.Set("numberPerPage", input.NumberPerPage)
	}
	if input.RuleId != "" {
		q.Set("ruleId", input.RuleId)
	}
	if input.RulesetId != "" {
		q.Set("rulesetId", input.RulesetId)
	}
	if input.Name != "" {
		q.Set("name", input.Name)
	}
	if input.Filter != "" {
		q.Set("filter", input.Filter)
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

type GetResourcesCommandInput struct {
	Page          string
	NumberPerPage string
	RuleId        string
	RulesetId     string
	Name          string
	Filter        string
	SortKey       string
	Order         string
}

//GetApplicationsResourcesMethodsCommand - Get Application Resource Methods
//RequestType: GET
//Input:
func (s *ApplicationsService) GetApplicationsResourcesMethodsCommand() (result *MethodsView, resp *http.Response, err error) {
	path := "/applications/resources/methods"
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

//DeleteApplicationResourceCommand - Delete an Application Resource
//RequestType: DELETE
//Input: input *DeleteApplicationResourceCommandInput
func (s *ApplicationsService) DeleteApplicationResourceCommand(input *DeleteApplicationResourceCommandInput) (resp *http.Response, err error) {
	path := "/applications/{applicationId}/resources/{resourceId}"
	path = strings.Replace(path, "{applicationId}", input.ApplicationId, -1)

	path = strings.Replace(path, "{resourceId}", input.ResourceId, -1)

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

type DeleteApplicationResourceCommandInput struct {
	ApplicationId string
	ResourceId    string
}

//GetApplicationResourceCommand - Get an Application Resource
//RequestType: GET
//Input: input *GetApplicationResourceCommandInput
func (s *ApplicationsService) GetApplicationResourceCommand(input *GetApplicationResourceCommandInput) (result *ResourceView, resp *http.Response, err error) {
	path := "/applications/{applicationId}/resources/{resourceId}"
	path = strings.Replace(path, "{applicationId}", input.ApplicationId, -1)

	path = strings.Replace(path, "{resourceId}", input.ResourceId, -1)

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

type GetApplicationResourceCommandInput struct {
	ApplicationId string
	ResourceId    string
}

//UpdateApplicationResourceCommand - Update an Application Resource
//RequestType: PUT
//Input: input *UpdateApplicationResourceCommandInput
func (s *ApplicationsService) UpdateApplicationResourceCommand(input *UpdateApplicationResourceCommandInput) (result *ResourceView, resp *http.Response, err error) {
	path := "/applications/{applicationId}/resources/{resourceId}"
	path = strings.Replace(path, "{applicationId}", input.ApplicationId, -1)

	path = strings.Replace(path, "{resourceId}", input.ResourceId, -1)

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

type UpdateApplicationResourceCommandInput struct {
	Body          ResourceView
	ApplicationId string
	ResourceId    string
}

//DeleteApplicationCommand - Delete an Application
//RequestType: DELETE
//Input: input *DeleteApplicationCommandInput
func (s *ApplicationsService) DeleteApplicationCommand(input *DeleteApplicationCommandInput) (result *ApplicationView, resp *http.Response, err error) {
	path := "/applications/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("DELETE", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

type DeleteApplicationCommandInput struct {
	Id string
}

//GetApplicationCommand - Get an Application
//RequestType: GET
//Input: input *GetApplicationCommandInput
func (s *ApplicationsService) GetApplicationCommand(input *GetApplicationCommandInput) (result *ApplicationView, resp *http.Response, err error) {
	path := "/applications/{id}"
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

type GetApplicationCommandInput struct {
	Id string
}

//UpdateApplicationCommand - Update an Application
//RequestType: PUT
//Input: input *UpdateApplicationCommandInput
func (s *ApplicationsService) UpdateApplicationCommand(input *UpdateApplicationCommandInput) (result *ApplicationView, resp *http.Response, err error) {
	path := "/applications/{id}"
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

type UpdateApplicationCommandInput struct {
	Body ApplicationView
	Id   string
}

//GetResourceMatchingEvaluationOrderCommand - Get the resource path ordering for an Application
//RequestType: GET
//Input: input *GetResourceMatchingEvaluationOrderCommandInput
func (s *ApplicationsService) GetResourceMatchingEvaluationOrderCommand(input *GetResourceMatchingEvaluationOrderCommandInput) (result *ResourceMatchingEvaluationOrderView, resp *http.Response, err error) {
	path := "/applications/{id}/resourceMatchingEvaluationOrder"
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

type GetResourceMatchingEvaluationOrderCommandInput struct {
	Id string
}

//GetApplicationResourcesCommand - Get Resources for an Application
//RequestType: GET
//Input: input *GetApplicationResourcesCommandInput
func (s *ApplicationsService) GetApplicationResourcesCommand(input *GetApplicationResourcesCommandInput) (result *ResourcesView, resp *http.Response, err error) {
	path := "/applications/{id}/resources"
	path = strings.Replace(path, "{id}", input.Id, -1)

	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	q := rel.Query()
	if input.Page != "" {
		q.Set("page", input.Page)
	}
	if input.NumberPerPage != "" {
		q.Set("numberPerPage", input.NumberPerPage)
	}
	if input.Name != "" {
		q.Set("name", input.Name)
	}
	if input.Filter != "" {
		q.Set("filter", input.Filter)
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

type GetApplicationResourcesCommandInput struct {
	Page          string
	NumberPerPage string
	Name          string
	Filter        string
	SortKey       string
	Order         string

	Id string
}

//AddApplicationResourceCommand - Add Resource to an Application
//RequestType: POST
//Input: input *AddApplicationResourceCommandInput
func (s *ApplicationsService) AddApplicationResourceCommand(input *AddApplicationResourceCommandInput) (result *ResourceView, resp *http.Response, err error) {
	path := "/applications/{id}/resources"
	path = strings.Replace(path, "{id}", input.Id, -1)

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

type AddApplicationResourceCommandInput struct {
	Body ResourceView
	Id   string
}

//GetResourceAutoOrderCommand - Computes an automatic, intelligent resource ordering for an Application.
//RequestType: GET
//Input: input *GetResourceAutoOrderCommandInput
func (s *ApplicationsService) GetResourceAutoOrderCommand(input *GetResourceAutoOrderCommandInput) (result *ResourceOrderView, resp *http.Response, err error) {
	path := "/applications/{id}/resources/autoOrder"
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

type GetResourceAutoOrderCommandInput struct {
	Id string
}
