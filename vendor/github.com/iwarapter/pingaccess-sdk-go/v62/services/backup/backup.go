package backup

import (
	"net/http"

	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess"
	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess/client"
	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess/client/metadata"
	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess/config"
	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess/request"
)

const (
	// ServiceName - The name of service.
	ServiceName = "Backup"
)

//BackupService provides the API operations for making requests to
// Backup endpoint.
type BackupService struct {
	*client.Client
}

//New createa a new instance of the BackupService client.
//
// Example:
//   cfg := config.NewConfig().WithUsername("Administrator").WithPassword("2Access").WithEndpoint(paURL)
//
//   //Create a BackupService from the configuration
//   svc := backup.New(cfg)
//
func New(cfg *config.Config) *BackupService {

	return &BackupService{Client: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingaccess.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a Backup operation
func (s *BackupService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := s.NewRequest(op, params, data)

	return req
}

//BackupCommand - Create a local database backup
//RequestType: GET
//Input:
func (s *BackupService) BackupCommand() (resp *http.Response, err error) {
	path := "/backup"
	op := &request.Operation{
		Name:       "BackupCommand",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}

	req := s.newRequest(op, nil, nil)

	if req.Send() == nil {
		return req.HTTPResponse, nil
	}
	return req.HTTPResponse, req.Error
}
