package sharedSecrets

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
	ServiceName = "SharedSecrets"
)

//SharedSecretsService provides the API operations for making requests to
// SharedSecrets endpoint.
type SharedSecretsService struct {
	*client.Client
}

//New createa a new instance of the SharedSecretsService client.
//
// Example:
//   cfg := config.NewConfig().WithUsername("Administrator").WithPassword("2FederateM0re").WithEndpoint(paURL.String())
//
//   //Create a SharedSecretsService from the configuration
//   svc := sharedSecrets.New(cfg)
//
func New(cfg *config.Config) *SharedSecretsService {

	return &SharedSecretsService{Client: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingaccess.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a SharedSecrets operation
func (s *SharedSecretsService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := s.NewRequest(op, params, data)

	return req
}

//GetSharedSecretsCommand - Get all Shared Secrets
//RequestType: GET
//Input: input *GetSharedSecretsCommandInput
func (s *SharedSecretsService) GetSharedSecretsCommand(input *GetSharedSecretsCommandInput) (output *models.SharedSecretsView, resp *http.Response, err error) {
	path := "/sharedSecrets"
	op := &request.Operation{
		Name:       "GetSharedSecretsCommand",
		HTTPMethod: "GET",
		HTTPPath:   path,
		QueryParams: map[string]string{
			"sortKey": input.SortKey,
			"order":   input.Order,
		},
	}
	output = &models.SharedSecretsView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// GetSharedSecretsCommandInput - Inputs for GetSharedSecretsCommand
type GetSharedSecretsCommandInput struct {
	SortKey string
	Order   string
}

//AddSharedSecretCommand - Create a Shared Secret
//RequestType: POST
//Input: input *AddSharedSecretCommandInput
func (s *SharedSecretsService) AddSharedSecretCommand(input *AddSharedSecretCommandInput) (output *models.SharedSecretView, resp *http.Response, err error) {
	path := "/sharedSecrets"
	op := &request.Operation{
		Name:        "AddSharedSecretCommand",
		HTTPMethod:  "POST",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.SharedSecretView{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// AddSharedSecretCommandInput - Inputs for AddSharedSecretCommand
type AddSharedSecretCommandInput struct {
	Body models.SharedSecretView
}

//DeleteSharedSecretCommand - Delete a Shared Secret
//RequestType: DELETE
//Input: input *DeleteSharedSecretCommandInput
func (s *SharedSecretsService) DeleteSharedSecretCommand(input *DeleteSharedSecretCommandInput) (resp *http.Response, err error) {
	path := "/sharedSecrets/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "DeleteSharedSecretCommand",
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

// DeleteSharedSecretCommandInput - Inputs for DeleteSharedSecretCommand
type DeleteSharedSecretCommandInput struct {
	Id string
}

//GetSharedSecretCommand - Get a Shared Secret
//RequestType: GET
//Input: input *GetSharedSecretCommandInput
func (s *SharedSecretsService) GetSharedSecretCommand(input *GetSharedSecretCommandInput) (output *models.SharedSecretView, resp *http.Response, err error) {
	path := "/sharedSecrets/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "GetSharedSecretCommand",
		HTTPMethod:  "GET",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.SharedSecretView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// GetSharedSecretCommandInput - Inputs for GetSharedSecretCommand
type GetSharedSecretCommandInput struct {
	Id string
}
