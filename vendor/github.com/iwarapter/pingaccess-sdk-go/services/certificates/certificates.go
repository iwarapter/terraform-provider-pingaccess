package certificates

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
	ServiceName = "Certificates"
)

//CertificatesService provides the API operations for making requests to
// Certificates endpoint.
type CertificatesService struct {
	*client.Client
}

//New createa a new instance of the CertificatesService client.
//
// Example:
//   cfg := config.NewConfig().WithUsername("Administrator").WithPassword("2FederateM0re").WithEndpoint(paURL.String())
//
//   //Create a CertificatesService from the configuration
//   svc := certificates.New(cfg)
//
func New(cfg *config.Config) *CertificatesService {

	return &CertificatesService{Client: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingaccess.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a Certificates operation
func (s *CertificatesService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := s.NewRequest(op, params, data)

	return req
}

//GetTrustedCerts - Get all Certificates
//RequestType: GET
//Input: input *GetTrustedCertsInput
func (s *CertificatesService) GetTrustedCerts(input *GetTrustedCertsInput) (output *models.TrustedCertsView, resp *http.Response, err error) {
	path := "/certificates"
	op := &request.Operation{
		Name:       "GetTrustedCerts",
		HTTPMethod: "GET",
		HTTPPath:   path,
		QueryParams: map[string]string{
			"page":          input.Page,
			"numberPerPage": input.NumberPerPage,
			"filter":        input.Filter,
			"alias":         input.Alias,
			"sortKey":       input.SortKey,
			"order":         input.Order,
		},
	}
	output = &models.TrustedCertsView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// GetTrustedCertsInput - Inputs for GetTrustedCerts
type GetTrustedCertsInput struct {
	Page          string
	NumberPerPage string
	Filter        string
	Alias         string
	SortKey       string
	Order         string
}

//ImportTrustedCert - Import a Certificate
//RequestType: POST
//Input: input *ImportTrustedCertInput
func (s *CertificatesService) ImportTrustedCert(input *ImportTrustedCertInput) (output *models.TrustedCertView, resp *http.Response, err error) {
	path := "/certificates"
	op := &request.Operation{
		Name:        "ImportTrustedCert",
		HTTPMethod:  "POST",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.TrustedCertView{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// ImportTrustedCertInput - Inputs for ImportTrustedCert
type ImportTrustedCertInput struct {
	Body models.X509FileImportDocView
}

//DeleteTrustedCertCommand - Delete a Certificate
//RequestType: DELETE
//Input: input *DeleteTrustedCertCommandInput
func (s *CertificatesService) DeleteTrustedCertCommand(input *DeleteTrustedCertCommandInput) (resp *http.Response, err error) {
	path := "/certificates/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "DeleteTrustedCertCommand",
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

// DeleteTrustedCertCommandInput - Inputs for DeleteTrustedCertCommand
type DeleteTrustedCertCommandInput struct {
	Id string
}

//GetTrustedCert - Get a Certificate
//RequestType: GET
//Input: input *GetTrustedCertInput
func (s *CertificatesService) GetTrustedCert(input *GetTrustedCertInput) (output *models.TrustedCertView, resp *http.Response, err error) {
	path := "/certificates/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "GetTrustedCert",
		HTTPMethod:  "GET",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.TrustedCertView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// GetTrustedCertInput - Inputs for GetTrustedCert
type GetTrustedCertInput struct {
	Id string
}

//UpdateTrustedCert - Update a Certificate
//RequestType: PUT
//Input: input *UpdateTrustedCertInput
func (s *CertificatesService) UpdateTrustedCert(input *UpdateTrustedCertInput) (output *models.TrustedCertView, resp *http.Response, err error) {
	path := "/certificates/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "UpdateTrustedCert",
		HTTPMethod:  "PUT",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.TrustedCertView{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// UpdateTrustedCertInput - Inputs for UpdateTrustedCert
type UpdateTrustedCertInput struct {
	Body models.X509FileImportDocView
	Id   string
}

//ExportTrustedCert - Export a Certificate
//RequestType: GET
//Input: input *ExportTrustedCertInput
func (s *CertificatesService) ExportTrustedCert(input *ExportTrustedCertInput) (output *string, resp *http.Response, err error) {
	path := "/certificates/{id}/file"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "ExportTrustedCert",
		HTTPMethod:  "GET",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = pingaccess.String("")
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// ExportTrustedCertInput - Inputs for ExportTrustedCert
type ExportTrustedCertInput struct {
	Id string
}
