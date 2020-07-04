package sites

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
	ServiceName = "Sites"
)

//SitesService provides the API operations for making requests to
// Sites endpoint.
type SitesService struct {
	*client.Client
}

//New createa a new instance of the SitesService client.
//
// Example:
//   cfg := config.NewConfig().WithUsername("Administrator").WithPassword("2FederateM0re").WithEndpoint(paURL.String())
//
//   //Create a SitesService from the configuration
//   svc := sites.New(cfg)
//
func New(cfg *config.Config) *SitesService {

	return &SitesService{Client: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingaccess.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a Sites operation
func (s *SitesService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := s.NewRequest(op, params, data)

	return req
}

//GetSitesCommand - Get all Sites
//RequestType: GET
//Input: input *GetSitesCommandInput
func (s *SitesService) GetSitesCommand(input *GetSitesCommandInput) (output *models.SitesView, resp *http.Response, err error) {
	path := "/sites"
	op := &request.Operation{
		Name:       "GetSitesCommand",
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
	output = &models.SitesView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// GetSitesCommandInput - Inputs for GetSitesCommand
type GetSitesCommandInput struct {
	Page          string
	NumberPerPage string
	Filter        string
	Name          string
	SortKey       string
	Order         string
}

//AddSiteCommand - Create a Site
//RequestType: POST
//Input: input *AddSiteCommandInput
func (s *SitesService) AddSiteCommand(input *AddSiteCommandInput) (output *models.SiteView, resp *http.Response, err error) {
	path := "/sites"
	op := &request.Operation{
		Name:        "AddSiteCommand",
		HTTPMethod:  "POST",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.SiteView{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// AddSiteCommandInput - Inputs for AddSiteCommand
type AddSiteCommandInput struct {
	Body models.SiteView
}

//DeleteSiteCommand - Delete a Site
//RequestType: DELETE
//Input: input *DeleteSiteCommandInput
func (s *SitesService) DeleteSiteCommand(input *DeleteSiteCommandInput) (resp *http.Response, err error) {
	path := "/sites/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "DeleteSiteCommand",
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

// DeleteSiteCommandInput - Inputs for DeleteSiteCommand
type DeleteSiteCommandInput struct {
	Id string
}

//GetSiteCommand - Get a Site
//RequestType: GET
//Input: input *GetSiteCommandInput
func (s *SitesService) GetSiteCommand(input *GetSiteCommandInput) (output *models.SiteView, resp *http.Response, err error) {
	path := "/sites/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "GetSiteCommand",
		HTTPMethod:  "GET",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.SiteView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// GetSiteCommandInput - Inputs for GetSiteCommand
type GetSiteCommandInput struct {
	Id string
}

//UpdateSiteCommand - Update a Site
//RequestType: PUT
//Input: input *UpdateSiteCommandInput
func (s *SitesService) UpdateSiteCommand(input *UpdateSiteCommandInput) (output *models.SiteView, resp *http.Response, err error) {
	path := "/sites/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "UpdateSiteCommand",
		HTTPMethod:  "PUT",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.SiteView{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// UpdateSiteCommandInput - Inputs for UpdateSiteCommand
type UpdateSiteCommandInput struct {
	Body models.SiteView
	Id   string
}
