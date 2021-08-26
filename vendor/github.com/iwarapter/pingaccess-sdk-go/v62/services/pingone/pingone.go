package pingone

import (
	"net/http"

	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess"
	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess/client"
	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess/client/metadata"
	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess/config"
	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess/models"
	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess/request"
)

const (
	// ServiceName - The name of service.
	ServiceName = "Pingone"
)

//PingoneService provides the API operations for making requests to
// Pingone endpoint.
type PingoneService struct {
	*client.Client
}

//New createa a new instance of the PingoneService client.
//
// Example:
//   cfg := config.NewConfig().WithUsername("Administrator").WithPassword("2Access").WithEndpoint(paURL)
//
//   //Create a PingoneService from the configuration
//   svc := pingone.New(cfg)
//
func New(cfg *config.Config) *PingoneService {

	return &PingoneService{Client: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingaccess.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a Pingone operation
func (s *PingoneService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := s.NewRequest(op, params, data)

	return req
}

//DeletePingOne4CCommand - Resets the PingOne For Customers configuration to default values
//RequestType: DELETE
//Input:
func (s *PingoneService) DeletePingOne4CCommand() (resp *http.Response, err error) {
	path := "/pingone/customers"
	op := &request.Operation{
		Name:       "DeletePingOne4CCommand",
		HTTPMethod: "DELETE",
		HTTPPath:   path,
	}

	req := s.newRequest(op, nil, nil)

	if req.Send() == nil {
		return req.HTTPResponse, nil
	}
	return req.HTTPResponse, req.Error
}

//GetPingOne4CCommand - Get the PingOne For Customers configuration
//RequestType: GET
//Input:
func (s *PingoneService) GetPingOne4CCommand() (output *models.PingOne4CView, resp *http.Response, err error) {
	path := "/pingone/customers"
	op := &request.Operation{
		Name:       "GetPingOne4CCommand",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.PingOne4CView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdatePingOne4CCommand - Update the PingOne For Customers configuration
//RequestType: PUT
//Input: input *UpdatePingOne4CCommandInput
func (s *PingoneService) UpdatePingOne4CCommand(input *UpdatePingOne4CCommandInput) (output *models.PingOne4CView, resp *http.Response, err error) {
	path := "/pingone/customers"
	op := &request.Operation{
		Name:        "UpdatePingOne4CCommand",
		HTTPMethod:  "PUT",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.PingOne4CView{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// UpdatePingOne4CCommandInput - Inputs for UpdatePingOne4CCommand
type UpdatePingOne4CCommandInput struct {
	Body models.PingOne4CView
}

//GetPingOne4CMetadataCommand - Get the PingOne for Customers metadata
//RequestType: GET
//Input:
func (s *PingoneService) GetPingOne4CMetadataCommand() (output *models.OIDCProviderMetadata, resp *http.Response, err error) {
	path := "/pingone/customers/metadata"
	op := &request.Operation{
		Name:       "GetPingOne4CMetadataCommand",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.OIDCProviderMetadata{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}
