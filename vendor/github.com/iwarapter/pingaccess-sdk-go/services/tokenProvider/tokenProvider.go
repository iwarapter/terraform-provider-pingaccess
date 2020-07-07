package tokenProvider

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
	ServiceName = "TokenProvider"
)

//TokenProviderService provides the API operations for making requests to
// TokenProvider endpoint.
type TokenProviderService struct {
	*client.Client
}

//New createa a new instance of the TokenProviderService client.
//
// Example:
//   cfg := config.NewConfig().WithUsername("Administrator").WithPassword("2FederateM0re").WithEndpoint(paURL.String())
//
//   //Create a TokenProviderService from the configuration
//   svc := tokenProvider.New(cfg)
//
func New(cfg *config.Config) *TokenProviderService {

	return &TokenProviderService{Client: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingaccess.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a TokenProvider operation
func (s *TokenProviderService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := s.NewRequest(op, params, data)

	return req
}

//DeleteTokenProviderSettingCommand - Resets the Token Provider settings to default values
//RequestType: DELETE
//Input:
func (s *TokenProviderService) DeleteTokenProviderSettingCommand() (resp *http.Response, err error) {
	path := "/tokenProvider/settings"
	op := &request.Operation{
		Name:       "DeleteTokenProviderSettingCommand",
		HTTPMethod: "DELETE",
		HTTPPath:   path,
	}

	req := s.newRequest(op, nil, nil)

	if req.Send() == nil {
		return req.HTTPResponse, nil
	}
	return req.HTTPResponse, req.Error
}

//GetTokenProviderSettingCommand - Get the Token Provider settings
//RequestType: GET
//Input:
func (s *TokenProviderService) GetTokenProviderSettingCommand() (output *models.TokenProviderSettingView, resp *http.Response, err error) {
	path := "/tokenProvider/settings"
	op := &request.Operation{
		Name:       "GetTokenProviderSettingCommand",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.TokenProviderSettingView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateTokenProviderSettingCommand - Update the Token Provider setting
//RequestType: PUT
//Input: input *UpdateTokenProviderSettingCommandInput
func (s *TokenProviderService) UpdateTokenProviderSettingCommand(input *UpdateTokenProviderSettingCommandInput) (output *models.TokenProviderSettingView, resp *http.Response, err error) {
	path := "/tokenProvider/settings"
	op := &request.Operation{
		Name:        "UpdateTokenProviderSettingCommand",
		HTTPMethod:  "PUT",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.TokenProviderSettingView{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// UpdateTokenProviderSettingCommandInput - Inputs for UpdateTokenProviderSettingCommand
type UpdateTokenProviderSettingCommandInput struct {
	Body models.TokenProviderSettingView
}
