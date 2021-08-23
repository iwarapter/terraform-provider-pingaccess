package trustedCertificateGroups

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
	ServiceName = "TrustedCertificateGroups"
)

//TrustedCertificateGroupsService provides the API operations for making requests to
// TrustedCertificateGroups endpoint.
type TrustedCertificateGroupsService struct {
	*client.Client
}

//New createa a new instance of the TrustedCertificateGroupsService client.
//
// Example:
//   cfg := config.NewConfig().WithUsername("Administrator").WithPassword("2Access").WithEndpoint(paURL)
//
//   //Create a TrustedCertificateGroupsService from the configuration
//   svc := trustedCertificateGroups.New(cfg)
//
func New(cfg *config.Config) *TrustedCertificateGroupsService {

	return &TrustedCertificateGroupsService{Client: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingaccess.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a TrustedCertificateGroups operation
func (s *TrustedCertificateGroupsService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := s.NewRequest(op, params, data)

	return req
}

//GetTrustedCertificateGroupsCommand - Get all Trusted Certificate Groups
//RequestType: GET
//Input: input *GetTrustedCertificateGroupsCommandInput
func (s *TrustedCertificateGroupsService) GetTrustedCertificateGroupsCommand(input *GetTrustedCertificateGroupsCommandInput) (output *models.TrustedCertificateGroupsView, resp *http.Response, err error) {
	path := "/trustedCertificateGroups"
	op := &request.Operation{
		Name:       "GetTrustedCertificateGroupsCommand",
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
	output = &models.TrustedCertificateGroupsView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// GetTrustedCertificateGroupsCommandInput - Inputs for GetTrustedCertificateGroupsCommand
type GetTrustedCertificateGroupsCommandInput struct {
	Page          string
	NumberPerPage string
	Filter        string
	Name          string
	SortKey       string
	Order         string
}

//AddTrustedCertificateGroupCommand - Create a Trusted Certificate Group
//RequestType: POST
//Input: input *AddTrustedCertificateGroupCommandInput
func (s *TrustedCertificateGroupsService) AddTrustedCertificateGroupCommand(input *AddTrustedCertificateGroupCommandInput) (output *models.TrustedCertificateGroupView, resp *http.Response, err error) {
	path := "/trustedCertificateGroups"
	op := &request.Operation{
		Name:        "AddTrustedCertificateGroupCommand",
		HTTPMethod:  "POST",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.TrustedCertificateGroupView{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// AddTrustedCertificateGroupCommandInput - Inputs for AddTrustedCertificateGroupCommand
type AddTrustedCertificateGroupCommandInput struct {
	Body models.TrustedCertificateGroupView
}

//DeleteTrustedCertificateGroupCommand - Delete a Trusted Certificate Group
//RequestType: DELETE
//Input: input *DeleteTrustedCertificateGroupCommandInput
func (s *TrustedCertificateGroupsService) DeleteTrustedCertificateGroupCommand(input *DeleteTrustedCertificateGroupCommandInput) (resp *http.Response, err error) {
	path := "/trustedCertificateGroups/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "DeleteTrustedCertificateGroupCommand",
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

// DeleteTrustedCertificateGroupCommandInput - Inputs for DeleteTrustedCertificateGroupCommand
type DeleteTrustedCertificateGroupCommandInput struct {
	Id string
}

//GetTrustedCertificateGroupCommand - Get a Trusted Certificate Group
//RequestType: GET
//Input: input *GetTrustedCertificateGroupCommandInput
func (s *TrustedCertificateGroupsService) GetTrustedCertificateGroupCommand(input *GetTrustedCertificateGroupCommandInput) (output *models.TrustedCertificateGroupView, resp *http.Response, err error) {
	path := "/trustedCertificateGroups/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "GetTrustedCertificateGroupCommand",
		HTTPMethod:  "GET",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.TrustedCertificateGroupView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// GetTrustedCertificateGroupCommandInput - Inputs for GetTrustedCertificateGroupCommand
type GetTrustedCertificateGroupCommandInput struct {
	Id string
}

//UpdateTrustedCertificateGroupCommand - Update a TrustedCertificateGroup
//RequestType: PUT
//Input: input *UpdateTrustedCertificateGroupCommandInput
func (s *TrustedCertificateGroupsService) UpdateTrustedCertificateGroupCommand(input *UpdateTrustedCertificateGroupCommandInput) (output *models.TrustedCertificateGroupView, resp *http.Response, err error) {
	path := "/trustedCertificateGroups/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "UpdateTrustedCertificateGroupCommand",
		HTTPMethod:  "PUT",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.TrustedCertificateGroupView{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// UpdateTrustedCertificateGroupCommandInput - Inputs for UpdateTrustedCertificateGroupCommand
type UpdateTrustedCertificateGroupCommandInput struct {
	Body models.TrustedCertificateGroupView
	Id   string
}
