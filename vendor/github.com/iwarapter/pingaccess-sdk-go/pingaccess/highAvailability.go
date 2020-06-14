package pingaccess

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type HighAvailabilityService service

type HighAvailabilityAPI interface {
	GetAvailabilityProfilesCommand(input *GetAvailabilityProfilesCommandInput) (result *AvailabilityProfilesView, resp *http.Response, err error)
	AddAvailabilityProfileCommand(input *AddAvailabilityProfileCommandInput) (result *AvailabilityProfileView, resp *http.Response, err error)
	GetAvailabilityProfileDescriptorsCommand() (result *DescriptorsView, resp *http.Response, err error)
	GetAvailabilityProfileDescriptorCommand(input *GetAvailabilityProfileDescriptorCommandInput) (result *DescriptorView, resp *http.Response, err error)
	DeleteAvailabilityProfileCommand(input *DeleteAvailabilityProfileCommandInput) (resp *http.Response, err error)
	GetAvailabilityProfileCommand(input *GetAvailabilityProfileCommandInput) (result *AvailabilityProfileView, resp *http.Response, err error)
	UpdateAvailabilityProfileCommand(input *UpdateAvailabilityProfileCommandInput) (result *AvailabilityProfileView, resp *http.Response, err error)
	GetLoadBalancingStrategiesCommand(input *GetLoadBalancingStrategiesCommandInput) (result *LoadBalancingStrategiesView, resp *http.Response, err error)
	AddLoadBalancingStrategyCommand(input *AddLoadBalancingStrategyCommandInput) (result *LoadBalancingStrategyView, resp *http.Response, err error)
	GetLoadBalancingStrategyDescriptorsCommand() (result *DescriptorsView, resp *http.Response, err error)
	GetLoadBalancingStrategyDescriptorCommand(input *GetLoadBalancingStrategyDescriptorCommandInput) (result *DescriptorView, resp *http.Response, err error)
	DeleteLoadBalancingStrategyCommand(input *DeleteLoadBalancingStrategyCommandInput) (resp *http.Response, err error)
	GetLoadBalancingStrategyCommand(input *GetLoadBalancingStrategyCommandInput) (result *LoadBalancingStrategyView, resp *http.Response, err error)
	UpdateLoadBalancingStrategyCommand(input *UpdateLoadBalancingStrategyCommandInput) (result *LoadBalancingStrategyView, resp *http.Response, err error)
}

//GetAvailabilityProfilesCommand - Get all Availability Profiles
//RequestType: GET
//Input: input *GetAvailabilityProfilesCommandInput
func (s *HighAvailabilityService) GetAvailabilityProfilesCommand(input *GetAvailabilityProfilesCommandInput) (result *AvailabilityProfilesView, resp *http.Response, err error) {
	path := "/highAvailability/availabilityProfiles"
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

type GetAvailabilityProfilesCommandInput struct {
	Page          string
	NumberPerPage string
	Filter        string
	Name          string
	SortKey       string
	Order         string
}

//AddAvailabilityProfileCommand - Create an Availability Profile
//RequestType: POST
//Input: input *AddAvailabilityProfileCommandInput
func (s *HighAvailabilityService) AddAvailabilityProfileCommand(input *AddAvailabilityProfileCommandInput) (result *AvailabilityProfileView, resp *http.Response, err error) {
	path := "/highAvailability/availabilityProfiles"
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

type AddAvailabilityProfileCommandInput struct {
	Body AvailabilityProfileView
}

//GetAvailabilityProfileDescriptorsCommand - Get descriptors for all Availability Profiles
//RequestType: GET
//Input:
func (s *HighAvailabilityService) GetAvailabilityProfileDescriptorsCommand() (result *DescriptorsView, resp *http.Response, err error) {
	path := "/highAvailability/availabilityProfiles/descriptors"
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

//GetAvailabilityProfileDescriptorCommand - Get a descriptor for an Availability Profile
//RequestType: GET
//Input: input *GetAvailabilityProfileDescriptorCommandInput
func (s *HighAvailabilityService) GetAvailabilityProfileDescriptorCommand(input *GetAvailabilityProfileDescriptorCommandInput) (result *DescriptorView, resp *http.Response, err error) {
	path := "/highAvailability/availabilityProfiles/descriptors/{availabilityProfileType}"
	path = strings.Replace(path, "{availabilityProfileType}", input.AvailabilityProfileType, -1)

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

type GetAvailabilityProfileDescriptorCommandInput struct {
	AvailabilityProfileType string
}

//DeleteAvailabilityProfileCommand - Delete an Availability Profile
//RequestType: DELETE
//Input: input *DeleteAvailabilityProfileCommandInput
func (s *HighAvailabilityService) DeleteAvailabilityProfileCommand(input *DeleteAvailabilityProfileCommandInput) (resp *http.Response, err error) {
	path := "/highAvailability/availabilityProfiles/{id}"
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

type DeleteAvailabilityProfileCommandInput struct {
	Id string
}

//GetAvailabilityProfileCommand - Get an Availability Profile
//RequestType: GET
//Input: input *GetAvailabilityProfileCommandInput
func (s *HighAvailabilityService) GetAvailabilityProfileCommand(input *GetAvailabilityProfileCommandInput) (result *AvailabilityProfileView, resp *http.Response, err error) {
	path := "/highAvailability/availabilityProfiles/{id}"
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

type GetAvailabilityProfileCommandInput struct {
	Id string
}

//UpdateAvailabilityProfileCommand - Update an Availability Profile
//RequestType: PUT
//Input: input *UpdateAvailabilityProfileCommandInput
func (s *HighAvailabilityService) UpdateAvailabilityProfileCommand(input *UpdateAvailabilityProfileCommandInput) (result *AvailabilityProfileView, resp *http.Response, err error) {
	path := "/highAvailability/availabilityProfiles/{id}"
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

type UpdateAvailabilityProfileCommandInput struct {
	Body AvailabilityProfileView
	Id   string
}

//GetLoadBalancingStrategiesCommand - Get all Load Balancing Strategies
//RequestType: GET
//Input: input *GetLoadBalancingStrategiesCommandInput
func (s *HighAvailabilityService) GetLoadBalancingStrategiesCommand(input *GetLoadBalancingStrategiesCommandInput) (result *LoadBalancingStrategiesView, resp *http.Response, err error) {
	path := "/highAvailability/loadBalancingStrategies"
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

type GetLoadBalancingStrategiesCommandInput struct {
	Page          string
	NumberPerPage string
	Filter        string
	Name          string
	SortKey       string
	Order         string
}

//AddLoadBalancingStrategyCommand - Create a Load Balancing Strategy
//RequestType: POST
//Input: input *AddLoadBalancingStrategyCommandInput
func (s *HighAvailabilityService) AddLoadBalancingStrategyCommand(input *AddLoadBalancingStrategyCommandInput) (result *LoadBalancingStrategyView, resp *http.Response, err error) {
	path := "/highAvailability/loadBalancingStrategies"
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

type AddLoadBalancingStrategyCommandInput struct {
	Body LoadBalancingStrategyView
}

//GetLoadBalancingStrategyDescriptorsCommand - Get descriptors for all Load Balancing Strategies
//RequestType: GET
//Input:
func (s *HighAvailabilityService) GetLoadBalancingStrategyDescriptorsCommand() (result *DescriptorsView, resp *http.Response, err error) {
	path := "/highAvailability/loadBalancingStrategies/descriptors"
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

//GetLoadBalancingStrategyDescriptorCommand - Get a descriptor for a Load Balancing Strategy
//RequestType: GET
//Input: input *GetLoadBalancingStrategyDescriptorCommandInput
func (s *HighAvailabilityService) GetLoadBalancingStrategyDescriptorCommand(input *GetLoadBalancingStrategyDescriptorCommandInput) (result *DescriptorView, resp *http.Response, err error) {
	path := "/highAvailability/loadBalancingStrategies/descriptors/{loadBalancingStrategyType}"
	path = strings.Replace(path, "{loadBalancingStrategyType}", input.LoadBalancingStrategyType, -1)

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

type GetLoadBalancingStrategyDescriptorCommandInput struct {
	LoadBalancingStrategyType string
}

//DeleteLoadBalancingStrategyCommand - Delete a Load Balancing Strategy
//RequestType: DELETE
//Input: input *DeleteLoadBalancingStrategyCommandInput
func (s *HighAvailabilityService) DeleteLoadBalancingStrategyCommand(input *DeleteLoadBalancingStrategyCommandInput) (resp *http.Response, err error) {
	path := "/highAvailability/loadBalancingStrategies/{id}"
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

type DeleteLoadBalancingStrategyCommandInput struct {
	Id string
}

//GetLoadBalancingStrategyCommand - Get a Load Balancing Strategy
//RequestType: GET
//Input: input *GetLoadBalancingStrategyCommandInput
func (s *HighAvailabilityService) GetLoadBalancingStrategyCommand(input *GetLoadBalancingStrategyCommandInput) (result *LoadBalancingStrategyView, resp *http.Response, err error) {
	path := "/highAvailability/loadBalancingStrategies/{id}"
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

type GetLoadBalancingStrategyCommandInput struct {
	Id string
}

//UpdateLoadBalancingStrategyCommand - Update a Load Balancing Strategy
//RequestType: PUT
//Input: input *UpdateLoadBalancingStrategyCommandInput
func (s *HighAvailabilityService) UpdateLoadBalancingStrategyCommand(input *UpdateLoadBalancingStrategyCommandInput) (result *LoadBalancingStrategyView, resp *http.Response, err error) {
	path := "/highAvailability/loadBalancingStrategies/{id}"
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

type UpdateLoadBalancingStrategyCommandInput struct {
	Body LoadBalancingStrategyView
	Id   string
}
