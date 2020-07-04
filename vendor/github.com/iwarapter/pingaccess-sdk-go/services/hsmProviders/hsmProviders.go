package hsmProviders

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
	ServiceName = "HsmProviders"
)

//HsmProvidersService provides the API operations for making requests to
// HsmProviders endpoint.
type HsmProvidersService struct {
	*client.Client
}

//New createa a new instance of the HsmProvidersService client.
//
// Example:
//   cfg := config.NewConfig().WithUsername("Administrator").WithPassword("2FederateM0re").WithEndpoint(paURL.String())
//
//   //Create a HsmProvidersService from the configuration
//   svc := hsmProviders.New(cfg)
//
func New(cfg *config.Config) *HsmProvidersService {

	return &HsmProvidersService{Client: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingaccess.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a HsmProviders operation
func (s *HsmProvidersService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := s.NewRequest(op, params, data)

	return req
}

//GetHsmProvidersCommand - Get all HSM Providers
//RequestType: GET
//Input: input *GetHsmProvidersCommandInput
func (s *HsmProvidersService) GetHsmProvidersCommand(input *GetHsmProvidersCommandInput) (output *models.HsmProviderView, resp *http.Response, err error) {
	path := "/hsmProviders"
	op := &request.Operation{
		Name:       "GetHsmProvidersCommand",
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
	output = &models.HsmProviderView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// GetHsmProvidersCommandInput - Inputs for GetHsmProvidersCommand
type GetHsmProvidersCommandInput struct {
	Page          string
	NumberPerPage string
	Filter        string
	Name          string
	SortKey       string
	Order         string
}

//AddHsmProviderCommand - Create an HSM Provider
//RequestType: POST
//Input: input *AddHsmProviderCommandInput
func (s *HsmProvidersService) AddHsmProviderCommand(input *AddHsmProviderCommandInput) (output *models.HsmProviderView, resp *http.Response, err error) {
	path := "/hsmProviders"
	op := &request.Operation{
		Name:        "AddHsmProviderCommand",
		HTTPMethod:  "POST",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.HsmProviderView{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// AddHsmProviderCommandInput - Inputs for AddHsmProviderCommand
type AddHsmProviderCommandInput struct {
	Body models.HsmProviderView
}

//GetHsmProviderDescriptorsCommand - Get descriptors for all HSM Providers
//RequestType: GET
//Input:
func (s *HsmProvidersService) GetHsmProviderDescriptorsCommand() (output *models.DescriptorsView, resp *http.Response, err error) {
	path := "/hsmProviders/descriptors"
	op := &request.Operation{
		Name:       "GetHsmProviderDescriptorsCommand",
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

//DeleteHsmProviderCommand - Delete an HSM Provider
//RequestType: DELETE
//Input: input *DeleteHsmProviderCommandInput
func (s *HsmProvidersService) DeleteHsmProviderCommand(input *DeleteHsmProviderCommandInput) (resp *http.Response, err error) {
	path := "/hsmProviders/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "DeleteHsmProviderCommand",
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

// DeleteHsmProviderCommandInput - Inputs for DeleteHsmProviderCommand
type DeleteHsmProviderCommandInput struct {
	Id string
}

//GetHsmProviderCommand - Get an HSM Provider
//RequestType: GET
//Input: input *GetHsmProviderCommandInput
func (s *HsmProvidersService) GetHsmProviderCommand(input *GetHsmProviderCommandInput) (output *models.HsmProviderView, resp *http.Response, err error) {
	path := "/hsmProviders/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "GetHsmProviderCommand",
		HTTPMethod:  "GET",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.HsmProviderView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// GetHsmProviderCommandInput - Inputs for GetHsmProviderCommand
type GetHsmProviderCommandInput struct {
	Id string
}

//UpdateHsmProviderCommand - Update an HSM Provider
//RequestType: PUT
//Input: input *UpdateHsmProviderCommandInput
func (s *HsmProvidersService) UpdateHsmProviderCommand(input *UpdateHsmProviderCommandInput) (output *models.HsmProviderView, resp *http.Response, err error) {
	path := "/hsmProviders/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "UpdateHsmProviderCommand",
		HTTPMethod:  "PUT",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.HsmProviderView{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// UpdateHsmProviderCommandInput - Inputs for UpdateHsmProviderCommand
type UpdateHsmProviderCommandInput struct {
	Body models.HsmProviderView
	Id   string
}
