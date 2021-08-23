package httpConfig

import (
	"net/http"

	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess/models"
)

type HttpConfigAPI interface {
	DeleteHttpMonitoringCommand() (resp *http.Response, err error)
	GetHttpMonitoringCommand() (output *models.HttpMonitoringView, resp *http.Response, err error)
	UpdateHttpMonitoringCommand(input *UpdateHttpMonitoringCommandInput) (output *models.HttpMonitoringView, resp *http.Response, err error)
	DeleteHostSourceCommand() (resp *http.Response, err error)
	GetHostSourceCommand() (output *models.HostMultiValueSourceView, resp *http.Response, err error)
	UpdateHostSourceCommand(input *UpdateHostSourceCommandInput) (output *models.HostMultiValueSourceView, resp *http.Response, err error)
	DeleteIpSourceCommand() (resp *http.Response, err error)
	GetIpSourceCommand() (output *models.IpMultiValueSourceView, resp *http.Response, err error)
	UpdateIpSourceCommand(input *UpdateIpSourceCommandInput) (output *models.IpMultiValueSourceView, resp *http.Response, err error)
	DeleteProtoSourceCommand() (resp *http.Response, err error)
	GetProtoSourceCommand() (output *models.ProtocolSourceView, resp *http.Response, err error)
	UpdateProtocolSourceCommand(input *UpdateProtocolSourceCommandInput) (output *models.ProtocolSourceView, resp *http.Response, err error)
}
