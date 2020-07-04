package identityMappings

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
	ServiceName = "IdentityMappings"
)

//IdentityMappingsService provides the API operations for making requests to
// IdentityMappings endpoint.
type IdentityMappingsService struct {
	*client.Client
}

//New createa a new instance of the IdentityMappingsService client.
//
// Example:
//   cfg := config.NewConfig().WithUsername("Administrator").WithPassword("2FederateM0re").WithEndpoint(paURL.String())
//
//   //Create a IdentityMappingsService from the configuration
//   svc := identityMappings.New(cfg)
//
func New(cfg *config.Config) *IdentityMappingsService {

	return &IdentityMappingsService{Client: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingaccess.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a IdentityMappings operation
func (s *IdentityMappingsService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := s.NewRequest(op, params, data)

	return req
}

//GetIdentityMappingsCommand - Get all Identity Mappings
//RequestType: GET
//Input: input *GetIdentityMappingsCommandInput
func (s *IdentityMappingsService) GetIdentityMappingsCommand(input *GetIdentityMappingsCommandInput) (output *models.IdentityMappingsView, resp *http.Response, err error) {
	path := "/identityMappings"
	op := &request.Operation{
		Name:       "GetIdentityMappingsCommand",
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
	output = &models.IdentityMappingsView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// GetIdentityMappingsCommandInput - Inputs for GetIdentityMappingsCommand
type GetIdentityMappingsCommandInput struct {
	Page          string
	NumberPerPage string
	Filter        string
	Name          string
	SortKey       string
	Order         string
}

//AddIdentityMappingCommand - Create an Identity Mapping
//RequestType: POST
//Input: input *AddIdentityMappingCommandInput
func (s *IdentityMappingsService) AddIdentityMappingCommand(input *AddIdentityMappingCommandInput) (output *models.IdentityMappingView, resp *http.Response, err error) {
	path := "/identityMappings"
	op := &request.Operation{
		Name:        "AddIdentityMappingCommand",
		HTTPMethod:  "POST",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.IdentityMappingView{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// AddIdentityMappingCommandInput - Inputs for AddIdentityMappingCommand
type AddIdentityMappingCommandInput struct {
	Body models.IdentityMappingView
}

//GetIdentityMappingDescriptorsCommand - Get descriptors for all supported Identity Mapping types
//RequestType: GET
//Input:
func (s *IdentityMappingsService) GetIdentityMappingDescriptorsCommand() (output *models.DescriptorsView, resp *http.Response, err error) {
	path := "/identityMappings/descriptors"
	op := &request.Operation{
		Name:       "GetIdentityMappingDescriptorsCommand",
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

//GetIdentityMappingDescriptorCommand - Get descriptor for an Identity Mapping type
//RequestType: GET
//Input: input *GetIdentityMappingDescriptorCommandInput
func (s *IdentityMappingsService) GetIdentityMappingDescriptorCommand(input *GetIdentityMappingDescriptorCommandInput) (output *models.DescriptorView, resp *http.Response, err error) {
	path := "/identityMappings/descriptors/{identityMappingType}"
	path = strings.Replace(path, "{identityMappingType}", input.IdentityMappingType, -1)

	op := &request.Operation{
		Name:        "GetIdentityMappingDescriptorCommand",
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

// GetIdentityMappingDescriptorCommandInput - Inputs for GetIdentityMappingDescriptorCommand
type GetIdentityMappingDescriptorCommandInput struct {
	IdentityMappingType string
}

//DeleteIdentityMappingCommand - Delete an Identity Mapping
//RequestType: DELETE
//Input: input *DeleteIdentityMappingCommandInput
func (s *IdentityMappingsService) DeleteIdentityMappingCommand(input *DeleteIdentityMappingCommandInput) (resp *http.Response, err error) {
	path := "/identityMappings/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "DeleteIdentityMappingCommand",
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

// DeleteIdentityMappingCommandInput - Inputs for DeleteIdentityMappingCommand
type DeleteIdentityMappingCommandInput struct {
	Id string
}

//GetIdentityMappingCommand - Get an Identity Mapping
//RequestType: GET
//Input: input *GetIdentityMappingCommandInput
func (s *IdentityMappingsService) GetIdentityMappingCommand(input *GetIdentityMappingCommandInput) (output *models.IdentityMappingView, resp *http.Response, err error) {
	path := "/identityMappings/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "GetIdentityMappingCommand",
		HTTPMethod:  "GET",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.IdentityMappingView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// GetIdentityMappingCommandInput - Inputs for GetIdentityMappingCommand
type GetIdentityMappingCommandInput struct {
	Id string
}

//UpdateIdentityMappingCommand - Update an Identity Mapping
//RequestType: PUT
//Input: input *UpdateIdentityMappingCommandInput
func (s *IdentityMappingsService) UpdateIdentityMappingCommand(input *UpdateIdentityMappingCommandInput) (output *models.IdentityMappingView, resp *http.Response, err error) {
	path := "/identityMappings/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "UpdateIdentityMappingCommand",
		HTTPMethod:  "PUT",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.IdentityMappingView{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// UpdateIdentityMappingCommandInput - Inputs for UpdateIdentityMappingCommand
type UpdateIdentityMappingCommandInput struct {
	Body models.IdentityMappingView
	Id   string
}
