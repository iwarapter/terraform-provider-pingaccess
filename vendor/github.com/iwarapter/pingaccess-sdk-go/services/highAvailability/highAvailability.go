package highAvailability

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
	ServiceName = "HighAvailability"
)

//HighAvailabilityService provides the API operations for making requests to
// HighAvailability endpoint.
type HighAvailabilityService struct {
	*client.Client
}

//New createa a new instance of the HighAvailabilityService client.
//
// Example:
//   cfg := config.NewConfig().WithUsername("Administrator").WithPassword("2FederateM0re").WithEndpoint(paURL.String())
//
//   //Create a HighAvailabilityService from the configuration
//   svc := highAvailability.New(cfg)
//
func New(cfg *config.Config) *HighAvailabilityService {

	return &HighAvailabilityService{Client: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingaccess.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a HighAvailability operation
func (s *HighAvailabilityService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := s.NewRequest(op, params, data)

	return req
}

//GetAvailabilityProfilesCommand - Get all Availability Profiles
//RequestType: GET
//Input: input *GetAvailabilityProfilesCommandInput
func (s *HighAvailabilityService) GetAvailabilityProfilesCommand(input *GetAvailabilityProfilesCommandInput) (output *models.AvailabilityProfilesView, resp *http.Response, err error) {
	path := "/highAvailability/availabilityProfiles"
	op := &request.Operation{
		Name:       "GetAvailabilityProfilesCommand",
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
	output = &models.AvailabilityProfilesView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// GetAvailabilityProfilesCommandInput - Inputs for GetAvailabilityProfilesCommand
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
func (s *HighAvailabilityService) AddAvailabilityProfileCommand(input *AddAvailabilityProfileCommandInput) (output *models.AvailabilityProfileView, resp *http.Response, err error) {
	path := "/highAvailability/availabilityProfiles"
	op := &request.Operation{
		Name:        "AddAvailabilityProfileCommand",
		HTTPMethod:  "POST",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.AvailabilityProfileView{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// AddAvailabilityProfileCommandInput - Inputs for AddAvailabilityProfileCommand
type AddAvailabilityProfileCommandInput struct {
	Body models.AvailabilityProfileView
}

//GetAvailabilityProfileDescriptorsCommand - Get descriptors for all Availability Profiles
//RequestType: GET
//Input:
func (s *HighAvailabilityService) GetAvailabilityProfileDescriptorsCommand() (output *models.DescriptorsView, resp *http.Response, err error) {
	path := "/highAvailability/availabilityProfiles/descriptors"
	op := &request.Operation{
		Name:       "GetAvailabilityProfileDescriptorsCommand",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.DescriptorsView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetAvailabilityProfileDescriptorCommand - Get a descriptor for an Availability Profile
//RequestType: GET
//Input: input *GetAvailabilityProfileDescriptorCommandInput
func (s *HighAvailabilityService) GetAvailabilityProfileDescriptorCommand(input *GetAvailabilityProfileDescriptorCommandInput) (output *models.DescriptorView, resp *http.Response, err error) {
	path := "/highAvailability/availabilityProfiles/descriptors/{availabilityProfileType}"
	path = strings.Replace(path, "{availabilityProfileType}", input.AvailabilityProfileType, -1)

	op := &request.Operation{
		Name:        "GetAvailabilityProfileDescriptorCommand",
		HTTPMethod:  "GET",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.DescriptorView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// GetAvailabilityProfileDescriptorCommandInput - Inputs for GetAvailabilityProfileDescriptorCommand
type GetAvailabilityProfileDescriptorCommandInput struct {
	AvailabilityProfileType string
}

//DeleteAvailabilityProfileCommand - Delete an Availability Profile
//RequestType: DELETE
//Input: input *DeleteAvailabilityProfileCommandInput
func (s *HighAvailabilityService) DeleteAvailabilityProfileCommand(input *DeleteAvailabilityProfileCommandInput) (resp *http.Response, err error) {
	path := "/highAvailability/availabilityProfiles/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "DeleteAvailabilityProfileCommand",
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

// DeleteAvailabilityProfileCommandInput - Inputs for DeleteAvailabilityProfileCommand
type DeleteAvailabilityProfileCommandInput struct {
	Id string
}

//GetAvailabilityProfileCommand - Get an Availability Profile
//RequestType: GET
//Input: input *GetAvailabilityProfileCommandInput
func (s *HighAvailabilityService) GetAvailabilityProfileCommand(input *GetAvailabilityProfileCommandInput) (output *models.AvailabilityProfileView, resp *http.Response, err error) {
	path := "/highAvailability/availabilityProfiles/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "GetAvailabilityProfileCommand",
		HTTPMethod:  "GET",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.AvailabilityProfileView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// GetAvailabilityProfileCommandInput - Inputs for GetAvailabilityProfileCommand
type GetAvailabilityProfileCommandInput struct {
	Id string
}

//UpdateAvailabilityProfileCommand - Update an Availability Profile
//RequestType: PUT
//Input: input *UpdateAvailabilityProfileCommandInput
func (s *HighAvailabilityService) UpdateAvailabilityProfileCommand(input *UpdateAvailabilityProfileCommandInput) (output *models.AvailabilityProfileView, resp *http.Response, err error) {
	path := "/highAvailability/availabilityProfiles/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "UpdateAvailabilityProfileCommand",
		HTTPMethod:  "PUT",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.AvailabilityProfileView{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// UpdateAvailabilityProfileCommandInput - Inputs for UpdateAvailabilityProfileCommand
type UpdateAvailabilityProfileCommandInput struct {
	Body models.AvailabilityProfileView
	Id   string
}

//GetLoadBalancingStrategiesCommand - Get all Load Balancing Strategies
//RequestType: GET
//Input: input *GetLoadBalancingStrategiesCommandInput
func (s *HighAvailabilityService) GetLoadBalancingStrategiesCommand(input *GetLoadBalancingStrategiesCommandInput) (output *models.LoadBalancingStrategiesView, resp *http.Response, err error) {
	path := "/highAvailability/loadBalancingStrategies"
	op := &request.Operation{
		Name:       "GetLoadBalancingStrategiesCommand",
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
	output = &models.LoadBalancingStrategiesView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// GetLoadBalancingStrategiesCommandInput - Inputs for GetLoadBalancingStrategiesCommand
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
func (s *HighAvailabilityService) AddLoadBalancingStrategyCommand(input *AddLoadBalancingStrategyCommandInput) (output *models.LoadBalancingStrategyView, resp *http.Response, err error) {
	path := "/highAvailability/loadBalancingStrategies"
	op := &request.Operation{
		Name:        "AddLoadBalancingStrategyCommand",
		HTTPMethod:  "POST",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.LoadBalancingStrategyView{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// AddLoadBalancingStrategyCommandInput - Inputs for AddLoadBalancingStrategyCommand
type AddLoadBalancingStrategyCommandInput struct {
	Body models.LoadBalancingStrategyView
}

//GetLoadBalancingStrategyDescriptorsCommand - Get descriptors for all Load Balancing Strategies
//RequestType: GET
//Input:
func (s *HighAvailabilityService) GetLoadBalancingStrategyDescriptorsCommand() (output *models.DescriptorsView, resp *http.Response, err error) {
	path := "/highAvailability/loadBalancingStrategies/descriptors"
	op := &request.Operation{
		Name:       "GetLoadBalancingStrategyDescriptorsCommand",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.DescriptorsView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetLoadBalancingStrategyDescriptorCommand - Get a descriptor for a Load Balancing Strategy
//RequestType: GET
//Input: input *GetLoadBalancingStrategyDescriptorCommandInput
func (s *HighAvailabilityService) GetLoadBalancingStrategyDescriptorCommand(input *GetLoadBalancingStrategyDescriptorCommandInput) (output *models.DescriptorView, resp *http.Response, err error) {
	path := "/highAvailability/loadBalancingStrategies/descriptors/{loadBalancingStrategyType}"
	path = strings.Replace(path, "{loadBalancingStrategyType}", input.LoadBalancingStrategyType, -1)

	op := &request.Operation{
		Name:        "GetLoadBalancingStrategyDescriptorCommand",
		HTTPMethod:  "GET",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.DescriptorView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// GetLoadBalancingStrategyDescriptorCommandInput - Inputs for GetLoadBalancingStrategyDescriptorCommand
type GetLoadBalancingStrategyDescriptorCommandInput struct {
	LoadBalancingStrategyType string
}

//DeleteLoadBalancingStrategyCommand - Delete a Load Balancing Strategy
//RequestType: DELETE
//Input: input *DeleteLoadBalancingStrategyCommandInput
func (s *HighAvailabilityService) DeleteLoadBalancingStrategyCommand(input *DeleteLoadBalancingStrategyCommandInput) (resp *http.Response, err error) {
	path := "/highAvailability/loadBalancingStrategies/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "DeleteLoadBalancingStrategyCommand",
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

// DeleteLoadBalancingStrategyCommandInput - Inputs for DeleteLoadBalancingStrategyCommand
type DeleteLoadBalancingStrategyCommandInput struct {
	Id string
}

//GetLoadBalancingStrategyCommand - Get a Load Balancing Strategy
//RequestType: GET
//Input: input *GetLoadBalancingStrategyCommandInput
func (s *HighAvailabilityService) GetLoadBalancingStrategyCommand(input *GetLoadBalancingStrategyCommandInput) (output *models.LoadBalancingStrategyView, resp *http.Response, err error) {
	path := "/highAvailability/loadBalancingStrategies/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "GetLoadBalancingStrategyCommand",
		HTTPMethod:  "GET",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.LoadBalancingStrategyView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// GetLoadBalancingStrategyCommandInput - Inputs for GetLoadBalancingStrategyCommand
type GetLoadBalancingStrategyCommandInput struct {
	Id string
}

//UpdateLoadBalancingStrategyCommand - Update a Load Balancing Strategy
//RequestType: PUT
//Input: input *UpdateLoadBalancingStrategyCommandInput
func (s *HighAvailabilityService) UpdateLoadBalancingStrategyCommand(input *UpdateLoadBalancingStrategyCommandInput) (output *models.LoadBalancingStrategyView, resp *http.Response, err error) {
	path := "/highAvailability/loadBalancingStrategies/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "UpdateLoadBalancingStrategyCommand",
		HTTPMethod:  "PUT",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.LoadBalancingStrategyView{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// UpdateLoadBalancingStrategyCommandInput - Inputs for UpdateLoadBalancingStrategyCommand
type UpdateLoadBalancingStrategyCommandInput struct {
	Body models.LoadBalancingStrategyView
	Id   string
}
