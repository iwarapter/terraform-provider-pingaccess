package accessTokenValidators

import (
	"net/http"
	"strings"

	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess"
	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess/client"
	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess/client/metadata"
	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess/config"
	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess/models"
	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess/request"
)

const (
	// ServiceName - The name of service.
	ServiceName = "AccessTokenValidators"
)

//AccessTokenValidatorsService provides the API operations for making requests to
// AccessTokenValidators endpoint.
type AccessTokenValidatorsService struct {
	*client.Client
}

//New createa a new instance of the AccessTokenValidatorsService client.
//
// Example:
//   cfg := config.NewConfig().WithUsername("Administrator").WithPassword("2Access").WithEndpoint(paURL)
//
//   //Create a AccessTokenValidatorsService from the configuration
//   svc := accessTokenValidators.New(cfg)
//
func New(cfg *config.Config) *AccessTokenValidatorsService {

	return &AccessTokenValidatorsService{Client: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingaccess.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a AccessTokenValidators operation
func (s *AccessTokenValidatorsService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := s.NewRequest(op, params, data)

	return req
}

//GetAccessTokenValidatorsCommand - Get all Access Token Validators
//RequestType: GET
//Input: input *GetAccessTokenValidatorsCommandInput
func (s *AccessTokenValidatorsService) GetAccessTokenValidatorsCommand(input *GetAccessTokenValidatorsCommandInput) (output *models.AccessTokenValidatorsView, resp *http.Response, err error) {
	path := "/accessTokenValidators"
	op := &request.Operation{
		Name:       "GetAccessTokenValidatorsCommand",
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
	output = &models.AccessTokenValidatorsView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// GetAccessTokenValidatorsCommandInput - Inputs for GetAccessTokenValidatorsCommand
type GetAccessTokenValidatorsCommandInput struct {
	Page          string
	NumberPerPage string
	Filter        string
	Name          string
	SortKey       string
	Order         string
}

//AddAccessTokenValidatorCommand - Create an Access Token Validator
//RequestType: POST
//Input: input *AddAccessTokenValidatorCommandInput
func (s *AccessTokenValidatorsService) AddAccessTokenValidatorCommand(input *AddAccessTokenValidatorCommandInput) (output *models.AccessTokenValidatorView, resp *http.Response, err error) {
	path := "/accessTokenValidators"
	op := &request.Operation{
		Name:        "AddAccessTokenValidatorCommand",
		HTTPMethod:  "POST",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.AccessTokenValidatorView{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// AddAccessTokenValidatorCommandInput - Inputs for AddAccessTokenValidatorCommand
type AddAccessTokenValidatorCommandInput struct {
	Body models.AccessTokenValidatorView
}

//GetAccessTokenValidatorDescriptorsCommand - Get descriptors for all Access Token Validators
//RequestType: GET
//Input:
func (s *AccessTokenValidatorsService) GetAccessTokenValidatorDescriptorsCommand() (output *models.DescriptorsView, resp *http.Response, err error) {
	path := "/accessTokenValidators/descriptors"
	op := &request.Operation{
		Name:       "GetAccessTokenValidatorDescriptorsCommand",
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

//DeleteAccessTokenValidatorCommand - Delete a Access Token Validator
//RequestType: DELETE
//Input: input *DeleteAccessTokenValidatorCommandInput
func (s *AccessTokenValidatorsService) DeleteAccessTokenValidatorCommand(input *DeleteAccessTokenValidatorCommandInput) (resp *http.Response, err error) {
	path := "/accessTokenValidators/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "DeleteAccessTokenValidatorCommand",
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

// DeleteAccessTokenValidatorCommandInput - Inputs for DeleteAccessTokenValidatorCommand
type DeleteAccessTokenValidatorCommandInput struct {
	Id string
}

//GetAccessTokenValidatorCommand - Get an Access Token Validator
//RequestType: GET
//Input: input *GetAccessTokenValidatorCommandInput
func (s *AccessTokenValidatorsService) GetAccessTokenValidatorCommand(input *GetAccessTokenValidatorCommandInput) (output *models.AccessTokenValidatorView, resp *http.Response, err error) {
	path := "/accessTokenValidators/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "GetAccessTokenValidatorCommand",
		HTTPMethod:  "GET",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.AccessTokenValidatorView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// GetAccessTokenValidatorCommandInput - Inputs for GetAccessTokenValidatorCommand
type GetAccessTokenValidatorCommandInput struct {
	Id string
}

//UpdateAccessTokenValidatorCommand - Update an Access Token Validator
//RequestType: PUT
//Input: input *UpdateAccessTokenValidatorCommandInput
func (s *AccessTokenValidatorsService) UpdateAccessTokenValidatorCommand(input *UpdateAccessTokenValidatorCommandInput) (output *models.AccessTokenValidatorView, resp *http.Response, err error) {
	path := "/accessTokenValidators/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "UpdateAccessTokenValidatorCommand",
		HTTPMethod:  "PUT",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.AccessTokenValidatorView{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// UpdateAccessTokenValidatorCommandInput - Inputs for UpdateAccessTokenValidatorCommand
type UpdateAccessTokenValidatorCommandInput struct {
	Body models.AccessTokenValidatorView
	Id   string
}
