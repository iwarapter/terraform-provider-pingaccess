package httpConfig

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
	ServiceName = "HttpConfig"
)

//HttpConfigService provides the API operations for making requests to
// HttpConfig endpoint.
type HttpConfigService struct {
	*client.Client
}

//New createa a new instance of the HttpConfigService client.
//
// Example:
//   cfg := config.NewConfig().WithUsername("Administrator").WithPassword("2Access").WithEndpoint(paURL)
//
//   //Create a HttpConfigService from the configuration
//   svc := httpConfig.New(cfg)
//
func New(cfg *config.Config) *HttpConfigService {

	return &HttpConfigService{Client: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingaccess.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a HttpConfig operation
func (s *HttpConfigService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := s.NewRequest(op, params, data)

	return req
}

//DeleteHttpMonitoringCommand - Resets the HTTP monitoring auditLevel to default value
//RequestType: DELETE
//Input:
func (s *HttpConfigService) DeleteHttpMonitoringCommand() (resp *http.Response, err error) {
	path := "/httpConfig/monitoring"
	op := &request.Operation{
		Name:       "DeleteHttpMonitoringCommand",
		HTTPMethod: "DELETE",
		HTTPPath:   path,
	}

	req := s.newRequest(op, nil, nil)

	if req.Send() == nil {
		return req.HTTPResponse, nil
	}
	return req.HTTPResponse, req.Error
}

//GetHttpMonitoringCommand - Get the HTTP monitoring auditLevel
//RequestType: GET
//Input:
func (s *HttpConfigService) GetHttpMonitoringCommand() (output *models.HttpMonitoringView, resp *http.Response, err error) {
	path := "/httpConfig/monitoring"
	op := &request.Operation{
		Name:       "GetHttpMonitoringCommand",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.HttpMonitoringView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateHttpMonitoringCommand - Update the HTTP monitoring auditLevel
//RequestType: PUT
//Input: input *UpdateHttpMonitoringCommandInput
func (s *HttpConfigService) UpdateHttpMonitoringCommand(input *UpdateHttpMonitoringCommandInput) (output *models.HttpMonitoringView, resp *http.Response, err error) {
	path := "/httpConfig/monitoring"
	op := &request.Operation{
		Name:        "UpdateHttpMonitoringCommand",
		HTTPMethod:  "PUT",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.HttpMonitoringView{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// UpdateHttpMonitoringCommandInput - Inputs for UpdateHttpMonitoringCommand
type UpdateHttpMonitoringCommandInput struct {
	Body models.HttpMonitoringView
}

//DeleteHostSourceCommand - Resets the HTTP request Host Source type to default values
//RequestType: DELETE
//Input:
func (s *HttpConfigService) DeleteHostSourceCommand() (resp *http.Response, err error) {
	path := "/httpConfig/request/hostSource"
	op := &request.Operation{
		Name:       "DeleteHostSourceCommand",
		HTTPMethod: "DELETE",
		HTTPPath:   path,
	}

	req := s.newRequest(op, nil, nil)

	if req.Send() == nil {
		return req.HTTPResponse, nil
	}
	return req.HTTPResponse, req.Error
}

//GetHostSourceCommand - Get the HTTP request Host Source type
//RequestType: GET
//Input:
func (s *HttpConfigService) GetHostSourceCommand() (output *models.HostMultiValueSourceView, resp *http.Response, err error) {
	path := "/httpConfig/request/hostSource"
	op := &request.Operation{
		Name:       "GetHostSourceCommand",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.HostMultiValueSourceView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateHostSourceCommand - Update the HTTP request Host Source type
//RequestType: PUT
//Input: input *UpdateHostSourceCommandInput
func (s *HttpConfigService) UpdateHostSourceCommand(input *UpdateHostSourceCommandInput) (output *models.HostMultiValueSourceView, resp *http.Response, err error) {
	path := "/httpConfig/request/hostSource"
	op := &request.Operation{
		Name:        "UpdateHostSourceCommand",
		HTTPMethod:  "PUT",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.HostMultiValueSourceView{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// UpdateHostSourceCommandInput - Inputs for UpdateHostSourceCommand
type UpdateHostSourceCommandInput struct {
	Body models.HostMultiValueSourceView
}

//DeleteIpSourceCommand - Resets the HTTP request IP Source type to default values
//RequestType: DELETE
//Input:
func (s *HttpConfigService) DeleteIpSourceCommand() (resp *http.Response, err error) {
	path := "/httpConfig/request/ipSource"
	op := &request.Operation{
		Name:       "DeleteIpSourceCommand",
		HTTPMethod: "DELETE",
		HTTPPath:   path,
	}

	req := s.newRequest(op, nil, nil)

	if req.Send() == nil {
		return req.HTTPResponse, nil
	}
	return req.HTTPResponse, req.Error
}

//GetIpSourceCommand - Get the HTTP request IP Source type
//RequestType: GET
//Input:
func (s *HttpConfigService) GetIpSourceCommand() (output *models.IpMultiValueSourceView, resp *http.Response, err error) {
	path := "/httpConfig/request/ipSource"
	op := &request.Operation{
		Name:       "GetIpSourceCommand",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.IpMultiValueSourceView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateIpSourceCommand - Update the HTTP request IP Source type
//RequestType: PUT
//Input: input *UpdateIpSourceCommandInput
func (s *HttpConfigService) UpdateIpSourceCommand(input *UpdateIpSourceCommandInput) (output *models.IpMultiValueSourceView, resp *http.Response, err error) {
	path := "/httpConfig/request/ipSource"
	op := &request.Operation{
		Name:        "UpdateIpSourceCommand",
		HTTPMethod:  "PUT",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.IpMultiValueSourceView{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// UpdateIpSourceCommandInput - Inputs for UpdateIpSourceCommand
type UpdateIpSourceCommandInput struct {
	Body models.IpMultiValueSourceView
}

//DeleteProtoSourceCommand - Resets the HTTP request Protocol Source type to default values
//RequestType: DELETE
//Input:
func (s *HttpConfigService) DeleteProtoSourceCommand() (resp *http.Response, err error) {
	path := "/httpConfig/request/protocolSource"
	op := &request.Operation{
		Name:       "DeleteProtoSourceCommand",
		HTTPMethod: "DELETE",
		HTTPPath:   path,
	}

	req := s.newRequest(op, nil, nil)

	if req.Send() == nil {
		return req.HTTPResponse, nil
	}
	return req.HTTPResponse, req.Error
}

//GetProtoSourceCommand - Get the HTTP request Protocol Source type
//RequestType: GET
//Input:
func (s *HttpConfigService) GetProtoSourceCommand() (output *models.ProtocolSourceView, resp *http.Response, err error) {
	path := "/httpConfig/request/protocolSource"
	op := &request.Operation{
		Name:       "GetProtoSourceCommand",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.ProtocolSourceView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateProtocolSourceCommand - Update the HTTP request Protocol Source type
//RequestType: PUT
//Input: input *UpdateProtocolSourceCommandInput
func (s *HttpConfigService) UpdateProtocolSourceCommand(input *UpdateProtocolSourceCommandInput) (output *models.ProtocolSourceView, resp *http.Response, err error) {
	path := "/httpConfig/request/protocolSource"
	op := &request.Operation{
		Name:        "UpdateProtocolSourceCommand",
		HTTPMethod:  "PUT",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.ProtocolSourceView{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// UpdateProtocolSourceCommandInput - Inputs for UpdateProtocolSourceCommand
type UpdateProtocolSourceCommandInput struct {
	Body models.ProtocolSourceView
}
