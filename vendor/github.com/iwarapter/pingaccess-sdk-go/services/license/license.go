package license

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
	ServiceName = "License"
)

//LicenseService provides the API operations for making requests to
// License endpoint.
type LicenseService struct {
	*client.Client
}

//New createa a new instance of the LicenseService client.
//
// Example:
//   cfg := config.NewConfig().WithUsername("Administrator").WithPassword("2Access").WithEndpoint(paURL)
//
//   //Create a LicenseService from the configuration
//   svc := license.New(cfg)
//
func New(cfg *config.Config) *LicenseService {

	return &LicenseService{Client: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingaccess.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a License operation
func (s *LicenseService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := s.NewRequest(op, params, data)

	return req
}

//GetLicenseCommand - Get the License File
//RequestType: GET
//Input:
func (s *LicenseService) GetLicenseCommand() (output *models.LicenseView, resp *http.Response, err error) {
	path := "/license"
	op := &request.Operation{
		Name:       "GetLicenseCommand",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.LicenseView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//ImportLicenseCommand - Import a License
//RequestType: POST
//Input: input *ImportLicenseCommandInput
func (s *LicenseService) ImportLicenseCommand(input *ImportLicenseCommandInput) (output *models.LicenseView, resp *http.Response, err error) {
	path := "/license"
	op := &request.Operation{
		Name:        "ImportLicenseCommand",
		HTTPMethod:  "POST",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.LicenseView{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// ImportLicenseCommandInput - Inputs for ImportLicenseCommand
type ImportLicenseCommandInput struct {
	Body models.LicenseImportDocView
}
