package version

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
	ServiceName = "Version"
)

//VersionService provides the API operations for making requests to
// Version endpoint.
type VersionService struct {
	*client.Client
}

//New createa a new instance of the VersionService client.
//
// Example:
//   cfg := config.NewConfig().WithUsername("Administrator").WithPassword("2Access").WithEndpoint(paURL)
//
//   //Create a VersionService from the configuration
//   svc := version.New(cfg)
//
func New(cfg *config.Config) *VersionService {

	return &VersionService{Client: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingaccess.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a Version operation
func (s *VersionService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := s.NewRequest(op, params, data)

	return req
}

//VersionCommand - Get the PingAccess version number
//RequestType: GET
//Input:
func (s *VersionService) VersionCommand() (output *models.VersionDocClass, resp *http.Response, err error) {
	path := "/version"
	op := &request.Operation{
		Name:       "VersionCommand",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.VersionDocClass{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}
