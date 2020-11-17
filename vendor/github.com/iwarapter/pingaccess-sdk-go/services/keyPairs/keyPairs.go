package keyPairs

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
	ServiceName = "KeyPairs"
)

//KeyPairsService provides the API operations for making requests to
// KeyPairs endpoint.
type KeyPairsService struct {
	*client.Client
}

//New createa a new instance of the KeyPairsService client.
//
// Example:
//   cfg := config.NewConfig().WithUsername("Administrator").WithPassword("2FederateM0re").WithEndpoint(paURL.String())
//
//   //Create a KeyPairsService from the configuration
//   svc := keyPairs.New(cfg)
//
func New(cfg *config.Config) *KeyPairsService {

	return &KeyPairsService{Client: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingaccess.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a KeyPairs operation
func (s *KeyPairsService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := s.NewRequest(op, params, data)

	return req
}

//GetKeyPairsCommand - Get all Key Pairs
//RequestType: GET
//Input: input *GetKeyPairsCommandInput
func (s *KeyPairsService) GetKeyPairsCommand(input *GetKeyPairsCommandInput) (output *models.KeyPairsView, resp *http.Response, err error) {
	path := "/keyPairs"
	op := &request.Operation{
		Name:       "GetKeyPairsCommand",
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
	output = &models.KeyPairsView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// GetKeyPairsCommandInput - Inputs for GetKeyPairsCommand
type GetKeyPairsCommandInput struct {
	Page          string
	NumberPerPage string
	Filter        string
	Alias         string
	SortKey       string
	Order         string
}

//GenerateKeyPairCommand - Generate a Key Pair
//RequestType: POST
//Input: input *GenerateKeyPairCommandInput
func (s *KeyPairsService) GenerateKeyPairCommand(input *GenerateKeyPairCommandInput) (output *models.KeyPairView, resp *http.Response, err error) {
	path := "/keyPairs/generate"
	op := &request.Operation{
		Name:        "GenerateKeyPairCommand",
		HTTPMethod:  "POST",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.KeyPairView{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// GenerateKeyPairCommandInput - Inputs for GenerateKeyPairCommand
type GenerateKeyPairCommandInput struct {
	Body models.NewKeyPairConfigView
}

//ImportKeyPairCommand - Import a Key Pair from a PKCS12 file
//RequestType: POST
//Input: input *ImportKeyPairCommandInput
func (s *KeyPairsService) ImportKeyPairCommand(input *ImportKeyPairCommandInput) (output *models.KeyPairView, resp *http.Response, err error) {
	path := "/keyPairs/import"
	op := &request.Operation{
		Name:        "ImportKeyPairCommand",
		HTTPMethod:  "POST",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.KeyPairView{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// ImportKeyPairCommandInput - Inputs for ImportKeyPairCommand
type ImportKeyPairCommandInput struct {
	Body models.PKCS12FileImportDocView
}

//KeyAlgorithms - Get the key algorithms supported by Key Pair generation
//RequestType: GET
//Input:
func (s *KeyPairsService) KeyAlgorithms() (output *models.KeyAlgorithmsView, resp *http.Response, err error) {
	path := "/keyPairs/keyAlgorithms"
	op := &request.Operation{
		Name:       "KeyAlgorithms",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.KeyAlgorithmsView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetKeypairsCreatableGeneralNamesCommand - Get the valid General Names for creating Subject Alternative Names
//RequestType: GET
//Input:
func (s *KeyPairsService) GetKeypairsCreatableGeneralNamesCommand() (output *models.SanTypes, resp *http.Response, err error) {
	path := "/keyPairs/subjectAlternativeTypes"
	op := &request.Operation{
		Name:       "GetKeypairsCreatableGeneralNamesCommand",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.SanTypes{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//DeleteKeyPairCommand - Delete a Key Pair
//RequestType: DELETE
//Input: input *DeleteKeyPairCommandInput
func (s *KeyPairsService) DeleteKeyPairCommand(input *DeleteKeyPairCommandInput) (resp *http.Response, err error) {
	path := "/keyPairs/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "DeleteKeyPairCommand",
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

// DeleteKeyPairCommandInput - Inputs for DeleteKeyPairCommand
type DeleteKeyPairCommandInput struct {
	Id string
}

//GetKeyPairCommand - Get a Key Pair
//RequestType: GET
//Input: input *GetKeyPairCommandInput
func (s *KeyPairsService) GetKeyPairCommand(input *GetKeyPairCommandInput) (output *models.KeyPairView, resp *http.Response, err error) {
	path := "/keyPairs/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "GetKeyPairCommand",
		HTTPMethod:  "GET",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.KeyPairView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// GetKeyPairCommandInput - Inputs for GetKeyPairCommand
type GetKeyPairCommandInput struct {
	Id string
}

//PatchKeyPairCommand - Update the chainCertificates of a Key Pair
//RequestType: PATCH
//Input: input *PatchKeyPairCommandInput
func (s *KeyPairsService) PatchKeyPairCommand(input *PatchKeyPairCommandInput) (output *models.KeyPairView, resp *http.Response, err error) {
	path := "/keyPairs/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "PatchKeyPairCommand",
		HTTPMethod:  "PATCH",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.KeyPairView{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// PatchKeyPairCommandInput - Inputs for PatchKeyPairCommand
type PatchKeyPairCommandInput struct {
	Body models.ChainCertificatesDocView
	Id   string
}

//UpdateKeyPairCommand - Update a Key Pair
//RequestType: PUT
//Input: input *UpdateKeyPairCommandInput
func (s *KeyPairsService) UpdateKeyPairCommand(input *UpdateKeyPairCommandInput) (output *models.KeyPairView, resp *http.Response, err error) {
	path := "/keyPairs/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "UpdateKeyPairCommand",
		HTTPMethod:  "PUT",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.KeyPairView{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// UpdateKeyPairCommandInput - Inputs for UpdateKeyPairCommand
type UpdateKeyPairCommandInput struct {
	Body models.PKCS12FileImportDocView
	Id   string
}

//ExportKeyPairCert - Export only the Certificate from a Key Pair
//RequestType: GET
//Input: input *ExportKeyPairCertInput
func (s *KeyPairsService) ExportKeyPairCert(input *ExportKeyPairCertInput) (output *string, resp *http.Response, err error) {
	path := "/keyPairs/{id}/certificate"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "ExportKeyPairCert",
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

// ExportKeyPairCertInput - Inputs for ExportKeyPairCert
type ExportKeyPairCertInput struct {
	Id string
}

//GenerateCsrCommand - Generate a Certificate Signing Request for a Key Pair
//RequestType: GET
//Input: input *GenerateCsrCommandInput
func (s *KeyPairsService) GenerateCsrCommand(input *GenerateCsrCommandInput) (output *string, resp *http.Response, err error) {
	path := "/keyPairs/{id}/csr"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "GenerateCsrCommand",
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

// GenerateCsrCommandInput - Inputs for GenerateCsrCommand
type GenerateCsrCommandInput struct {
	Id string
}

//ImportCSRResponseCommand - Import a Certificate Signing Request response
//RequestType: POST
//Input: input *ImportCSRResponseCommandInput
func (s *KeyPairsService) ImportCSRResponseCommand(input *ImportCSRResponseCommandInput) (output *models.KeyPairView, resp *http.Response, err error) {
	path := "/keyPairs/{id}/csr"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "ImportCSRResponseCommand",
		HTTPMethod:  "POST",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.KeyPairView{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// ImportCSRResponseCommandInput - Inputs for ImportCSRResponseCommand
type ImportCSRResponseCommandInput struct {
	Body models.CSRResponseImportDocView
	Id   string
}

//ExportKeyPair - Export a Key Pair in the PKCS12 file format
//RequestType: POST
//Input: input *ExportKeyPairInput
func (s *KeyPairsService) ExportKeyPair(input *ExportKeyPairInput) (resp *http.Response, err error) {
	path := "/keyPairs/{id}/pkcs12"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "ExportKeyPair",
		HTTPMethod:  "POST",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}

	req := s.newRequest(op, input.Body, nil)

	if req.Send() == nil {
		return req.HTTPResponse, nil
	}
	return req.HTTPResponse, req.Error
}

// ExportKeyPairInput - Inputs for ExportKeyPair
type ExportKeyPairInput struct {
	Body models.ExportParameters
	Id   string
}

//DeleteChainCertificateCommand - Delete a Chain Certificate
//RequestType: DELETE
//Input: input *DeleteChainCertificateCommandInput
func (s *KeyPairsService) DeleteChainCertificateCommand(input *DeleteChainCertificateCommandInput) (resp *http.Response, err error) {
	path := "/keyPairs/{keyPairId}/chainCertificates/{chainCertificateId}"
	path = strings.Replace(path, "{keyPairId}", input.KeyPairId, -1)

	path = strings.Replace(path, "{chainCertificateId}", input.ChainCertificateId, -1)

	op := &request.Operation{
		Name:        "DeleteChainCertificateCommand",
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

// DeleteChainCertificateCommandInput - Inputs for DeleteChainCertificateCommand
type DeleteChainCertificateCommandInput struct {
	KeyPairId          string
	ChainCertificateId string
}
