package acme

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
	ServiceName = "Acme"
)

//AcmeService provides the API operations for making requests to
// Acme endpoint.
type AcmeService struct {
	*client.Client
}

//New createa a new instance of the AcmeService client.
//
// Example:
//   cfg := config.NewConfig().WithUsername("Administrator").WithPassword("2FederateM0re").WithEndpoint(paURL.String())
//
//   //Create a AcmeService from the configuration
//   svc := acme.New(cfg)
//
func New(cfg *config.Config) *AcmeService {

	return &AcmeService{Client: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingaccess.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a Acme operation
func (s *AcmeService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := s.NewRequest(op, params, data)

	return req
}

//GetAcmeServersCommand - Get all ACME Servers
//RequestType: GET
//Input: input *GetAcmeServersCommandInput
func (s *AcmeService) GetAcmeServersCommand(input *GetAcmeServersCommandInput) (output *models.AcmeServersView, resp *http.Response, err error) {
	path := "/acme/servers"
	op := &request.Operation{
		Name:       "GetAcmeServersCommand",
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
	output = &models.AcmeServersView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// GetAcmeServersCommandInput - Inputs for GetAcmeServersCommand
type GetAcmeServersCommandInput struct {
	Page          string
	NumberPerPage string
	Filter        string
	Name          string
	SortKey       string
	Order         string
}

//AddAcmeServerCommand - Add an ACME Server
//RequestType: POST
//Input: input *AddAcmeServerCommandInput
func (s *AcmeService) AddAcmeServerCommand(input *AddAcmeServerCommandInput) (output *models.AcmeServerView, resp *http.Response, err error) {
	path := "/acme/servers"
	op := &request.Operation{
		Name:        "AddAcmeServerCommand",
		HTTPMethod:  "POST",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.AcmeServerView{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// AddAcmeServerCommandInput - Inputs for AddAcmeServerCommand
type AddAcmeServerCommandInput struct {
	Body models.AcmeServerView
}

//GetDefaultAcmeServerCommand - Get the default ACME Server
//RequestType: GET
//Input:
func (s *AcmeService) GetDefaultAcmeServerCommand() (output *models.LinkView, resp *http.Response, err error) {
	path := "/acme/servers/default"
	op := &request.Operation{
		Name:       "GetDefaultAcmeServerCommand",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.LinkView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateDefaultAcmeServerCommand - Update the default ACME Server
//RequestType: PUT
//Input: input *UpdateDefaultAcmeServerCommandInput
func (s *AcmeService) UpdateDefaultAcmeServerCommand(input *UpdateDefaultAcmeServerCommandInput) (output *models.LinkView, resp *http.Response, err error) {
	path := "/acme/servers/default"
	op := &request.Operation{
		Name:        "UpdateDefaultAcmeServerCommand",
		HTTPMethod:  "PUT",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.LinkView{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// UpdateDefaultAcmeServerCommandInput - Inputs for UpdateDefaultAcmeServerCommand
type UpdateDefaultAcmeServerCommandInput struct {
	Body models.LinkView
}

//DeleteAcmeServerCommand - Delete an ACME Server
//RequestType: DELETE
//Input: input *DeleteAcmeServerCommandInput
func (s *AcmeService) DeleteAcmeServerCommand(input *DeleteAcmeServerCommandInput) (output *models.AcmeServerView, resp *http.Response, err error) {
	path := "/acme/servers/{acmeServerId}"
	path = strings.Replace(path, "{acmeServerId}", input.AcmeServerId, -1)

	op := &request.Operation{
		Name:        "DeleteAcmeServerCommand",
		HTTPMethod:  "DELETE",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.AcmeServerView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// DeleteAcmeServerCommandInput - Inputs for DeleteAcmeServerCommand
type DeleteAcmeServerCommandInput struct {
	AcmeServerId string
}

//GetAcmeServerCommand - Get an ACME Server
//RequestType: GET
//Input: input *GetAcmeServerCommandInput
func (s *AcmeService) GetAcmeServerCommand(input *GetAcmeServerCommandInput) (output *models.AcmeServerView, resp *http.Response, err error) {
	path := "/acme/servers/{acmeServerId}"
	path = strings.Replace(path, "{acmeServerId}", input.AcmeServerId, -1)

	op := &request.Operation{
		Name:        "GetAcmeServerCommand",
		HTTPMethod:  "GET",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.AcmeServerView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// GetAcmeServerCommandInput - Inputs for GetAcmeServerCommand
type GetAcmeServerCommandInput struct {
	AcmeServerId string
}

//GetAcmeAccountsCommand - Get all ACME Accounts
//RequestType: GET
//Input: input *GetAcmeAccountsCommandInput
func (s *AcmeService) GetAcmeAccountsCommand(input *GetAcmeAccountsCommandInput) (output *models.AcmeAccountView, resp *http.Response, err error) {
	path := "/acme/servers/{acmeServerId}/accounts"
	path = strings.Replace(path, "{acmeServerId}", input.AcmeServerId, -1)

	op := &request.Operation{
		Name:       "GetAcmeAccountsCommand",
		HTTPMethod: "GET",
		HTTPPath:   path,
		QueryParams: map[string]string{
			"page":          input.Page,
			"numberPerPage": input.NumberPerPage,
			"sortKey":       input.SortKey,
			"order":         input.Order,
		},
	}
	output = &models.AcmeAccountView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// GetAcmeAccountsCommandInput - Inputs for GetAcmeAccountsCommand
type GetAcmeAccountsCommandInput struct {
	Page          string
	NumberPerPage string
	SortKey       string
	Order         string

	AcmeServerId string
}

//AddAcmeAccountCommand - Add an ACME Account
//RequestType: POST
//Input: input *AddAcmeAccountCommandInput
func (s *AcmeService) AddAcmeAccountCommand(input *AddAcmeAccountCommandInput) (output *models.AcmeAccountView, resp *http.Response, err error) {
	path := "/acme/servers/{acmeServerId}/accounts"
	path = strings.Replace(path, "{acmeServerId}", input.AcmeServerId, -1)

	op := &request.Operation{
		Name:        "AddAcmeAccountCommand",
		HTTPMethod:  "POST",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.AcmeAccountView{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// AddAcmeAccountCommandInput - Inputs for AddAcmeAccountCommand
type AddAcmeAccountCommandInput struct {
	Body         models.AcmeAccountView
	AcmeServerId string
}

//DeleteAcmeAccountCommand - Delete an ACME Account
//RequestType: DELETE
//Input: input *DeleteAcmeAccountCommandInput
func (s *AcmeService) DeleteAcmeAccountCommand(input *DeleteAcmeAccountCommandInput) (output *models.AcmeAccountView, resp *http.Response, err error) {
	path := "/acme/servers/{acmeServerId}/accounts/{acmeAccountId}"
	path = strings.Replace(path, "{acmeServerId}", input.AcmeServerId, -1)

	path = strings.Replace(path, "{acmeAccountId}", input.AcmeAccountId, -1)

	op := &request.Operation{
		Name:        "DeleteAcmeAccountCommand",
		HTTPMethod:  "DELETE",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.AcmeAccountView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// DeleteAcmeAccountCommandInput - Inputs for DeleteAcmeAccountCommand
type DeleteAcmeAccountCommandInput struct {
	AcmeServerId  string
	AcmeAccountId string
}

//GetAcmeAccountCommand - Get an ACME Account
//RequestType: GET
//Input: input *GetAcmeAccountCommandInput
func (s *AcmeService) GetAcmeAccountCommand(input *GetAcmeAccountCommandInput) (output *models.AcmeAccountView, resp *http.Response, err error) {
	path := "/acme/servers/{acmeServerId}/accounts/{acmeAccountId}"
	path = strings.Replace(path, "{acmeServerId}", input.AcmeServerId, -1)

	path = strings.Replace(path, "{acmeAccountId}", input.AcmeAccountId, -1)

	op := &request.Operation{
		Name:        "GetAcmeAccountCommand",
		HTTPMethod:  "GET",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.AcmeAccountView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// GetAcmeAccountCommandInput - Inputs for GetAcmeAccountCommand
type GetAcmeAccountCommandInput struct {
	AcmeServerId  string
	AcmeAccountId string
}

//GetAcmeCertificateRequestsCommand - Get all ACME Certificate Requests
//RequestType: GET
//Input: input *GetAcmeCertificateRequestsCommandInput
func (s *AcmeService) GetAcmeCertificateRequestsCommand(input *GetAcmeCertificateRequestsCommandInput) (output *models.AcmeCertificateRequestView, resp *http.Response, err error) {
	path := "/acme/servers/{acmeServerId}/accounts/{acmeAccountId}/certificateRequests"
	path = strings.Replace(path, "{acmeServerId}", input.AcmeServerId, -1)

	path = strings.Replace(path, "{acmeAccountId}", input.AcmeAccountId, -1)

	op := &request.Operation{
		Name:       "GetAcmeCertificateRequestsCommand",
		HTTPMethod: "GET",
		HTTPPath:   path,
		QueryParams: map[string]string{
			"keyPairId":     input.KeyPairId,
			"page":          input.Page,
			"numberPerPage": input.NumberPerPage,
			"sortKey":       input.SortKey,
			"order":         input.Order,
		},
	}
	output = &models.AcmeCertificateRequestView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// GetAcmeCertificateRequestsCommandInput - Inputs for GetAcmeCertificateRequestsCommand
type GetAcmeCertificateRequestsCommandInput struct {
	KeyPairId     string
	Page          string
	NumberPerPage string
	SortKey       string
	Order         string

	AcmeServerId  string
	AcmeAccountId string
}

//AddAcmeCertificateRequestCommand - Initiate the ACME protocol
//RequestType: POST
//Input: input *AddAcmeCertificateRequestCommandInput
func (s *AcmeService) AddAcmeCertificateRequestCommand(input *AddAcmeCertificateRequestCommandInput) (output *models.AcmeCertificateRequestView, resp *http.Response, err error) {
	path := "/acme/servers/{acmeServerId}/accounts/{acmeAccountId}/certificateRequests"
	path = strings.Replace(path, "{acmeServerId}", input.AcmeServerId, -1)

	path = strings.Replace(path, "{acmeAccountId}", input.AcmeAccountId, -1)

	op := &request.Operation{
		Name:        "AddAcmeCertificateRequestCommand",
		HTTPMethod:  "POST",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.AcmeCertificateRequestView{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// AddAcmeCertificateRequestCommandInput - Inputs for AddAcmeCertificateRequestCommand
type AddAcmeCertificateRequestCommandInput struct {
	Body          models.AcmeCertificateRequestView
	AcmeServerId  string
	AcmeAccountId string
}

//DeleteAcmeCertificateRequestCommand - Delete an ACME Certificate Request
//RequestType: DELETE
//Input: input *DeleteAcmeCertificateRequestCommandInput
func (s *AcmeService) DeleteAcmeCertificateRequestCommand(input *DeleteAcmeCertificateRequestCommandInput) (output *models.AcmeCertificateRequestView, resp *http.Response, err error) {
	path := "/acme/servers/{acmeServerId}/accounts/{acmeAccountId}/certificateRequests/{acmeCertificateRequestId}"
	path = strings.Replace(path, "{acmeServerId}", input.AcmeServerId, -1)

	path = strings.Replace(path, "{acmeAccountId}", input.AcmeAccountId, -1)

	path = strings.Replace(path, "{acmeCertificateRequestId}", input.AcmeCertificateRequestId, -1)

	op := &request.Operation{
		Name:        "DeleteAcmeCertificateRequestCommand",
		HTTPMethod:  "DELETE",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.AcmeCertificateRequestView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// DeleteAcmeCertificateRequestCommandInput - Inputs for DeleteAcmeCertificateRequestCommand
type DeleteAcmeCertificateRequestCommandInput struct {
	AcmeServerId             string
	AcmeAccountId            string
	AcmeCertificateRequestId string
}

//GetAcmeCertificateRequestCommand - Get an ACME Certificate Request
//RequestType: GET
//Input: input *GetAcmeCertificateRequestCommandInput
func (s *AcmeService) GetAcmeCertificateRequestCommand(input *GetAcmeCertificateRequestCommandInput) (output *models.AcmeCertificateRequestView, resp *http.Response, err error) {
	path := "/acme/servers/{acmeServerId}/accounts/{acmeAccountId}/certificateRequests/{acmeCertificateRequestId}"
	path = strings.Replace(path, "{acmeServerId}", input.AcmeServerId, -1)

	path = strings.Replace(path, "{acmeAccountId}", input.AcmeAccountId, -1)

	path = strings.Replace(path, "{acmeCertificateRequestId}", input.AcmeCertificateRequestId, -1)

	op := &request.Operation{
		Name:        "GetAcmeCertificateRequestCommand",
		HTTPMethod:  "GET",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.AcmeCertificateRequestView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// GetAcmeCertificateRequestCommandInput - Inputs for GetAcmeCertificateRequestCommand
type GetAcmeCertificateRequestCommandInput struct {
	AcmeServerId             string
	AcmeAccountId            string
	AcmeCertificateRequestId string
}
