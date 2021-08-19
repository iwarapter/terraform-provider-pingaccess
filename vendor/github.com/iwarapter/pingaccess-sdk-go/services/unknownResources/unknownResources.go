package unknownResources

import (
	"net/http"

	"github.com/iwarapter/pingaccess-sdk-go/pingaccess"
	"github.com/iwarapter/pingaccess-sdk-go/pingaccess/client"
	"github.com/iwarapter/pingaccess-sdk-go/pingaccess/client/metadata"
	"github.com/iwarapter/pingaccess-sdk-go/pingaccess/config"
	"github.com/iwarapter/pingaccess-sdk-go/pingaccess/models"
	"github.com/iwarapter/pingaccess-sdk-go/pingaccess/request"
)

const (
	// ServiceName - The name of service.
	ServiceName = "UnknownResources"
)

//UnknownResourcesService provides the API operations for making requests to
// UnknownResources endpoint.
type UnknownResourcesService struct {
	*client.Client
}

//New createa a new instance of the UnknownResourcesService client.
//
// Example:
//   cfg := config.NewConfig().WithUsername("Administrator").WithPassword("2Access").WithEndpoint(paURL)
//
//   //Create a UnknownResourcesService from the configuration
//   svc := unknownResources.New(cfg)
//
func New(cfg *config.Config) *UnknownResourcesService {

	return &UnknownResourcesService{Client: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingaccess.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a UnknownResources operation
func (s *UnknownResourcesService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := s.NewRequest(op, params, data)

	return req
}

//Delete - Resets the global settings for unknown resources to default values
//RequestType: DELETE
//Input:
func (s *UnknownResourcesService) Delete() (resp *http.Response, err error) {
	path := "/unknownResources/settings"
	op := &request.Operation{
		Name:       "Delete",
		HTTPMethod: "DELETE",
		HTTPPath:   path,
	}

	req := s.newRequest(op, nil, nil)

	if req.Send() == nil {
		return req.HTTPResponse, nil
	}
	return req.HTTPResponse, req.Error
}

//Get - Retrieves the global settings for unknown resources
//RequestType: GET
//Input:
func (s *UnknownResourcesService) Get() (output *models.UnknownResourceSettingsView, resp *http.Response, err error) {
	path := "/unknownResources/settings"
	op := &request.Operation{
		Name:       "Get",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.UnknownResourceSettingsView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//Update - Updates the global settings for unknown resources
//RequestType: PUT
//Input: input *UpdateInput
func (s *UnknownResourcesService) Update(input *UpdateInput) (output *models.UnknownResourceSettingsView, resp *http.Response, err error) {
	path := "/unknownResources/settings"
	op := &request.Operation{
		Name:        "Update",
		HTTPMethod:  "PUT",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.UnknownResourceSettingsView{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// UpdateInput - Inputs for Update
type UpdateInput struct {
	Body models.UnknownResourceSettingsView
}
