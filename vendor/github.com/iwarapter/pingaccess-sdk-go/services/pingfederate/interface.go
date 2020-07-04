package pingfederate

import (
	"net/http"

	"github.com/iwarapter/pingaccess-sdk-go/pingaccess/models"
)

type PingfederateAPI interface {
	DeletePingFederateCommand() (resp *http.Response, err error)
	GetPingFederateCommand() (output *models.PingFederateRuntimeView, resp *http.Response, err error)
	UpdatePingFederateCommand(input *UpdatePingFederateCommandInput) (output *models.PingFederateRuntimeView, resp *http.Response, err error)
	DeletePingFederateAccessTokensCommand() (resp *http.Response, err error)
	GetPingFederateAccessTokensCommand() (output *models.PingFederateAccessTokenView, resp *http.Response, err error)
	UpdatePingFederateAccessTokensCommand(input *UpdatePingFederateAccessTokensCommandInput) (output *models.PingFederateAccessTokenView, resp *http.Response, err error)
	DeletePingFederateAdminCommand() (resp *http.Response, err error)
	GetPingFederateAdminCommand() (output *models.PingFederateAdminView, resp *http.Response, err error)
	UpdatePingFederateAdminCommand(input *UpdatePingFederateAdminCommandInput) (output *models.PingFederateAdminView, resp *http.Response, err error)
	GetLegacyPingFederateMetadataCommand() (output *models.OIDCProviderMetadata, resp *http.Response, err error)
	DeletePingFederateRuntimeCommand() (resp *http.Response, err error)
	GetPingFederateRuntimeCommand() (output *models.PingFederateMetadataRuntimeView, resp *http.Response, err error)
	UpdatePingFederateRuntimeCommand(input *UpdatePingFederateRuntimeCommandInput) (output *models.PingFederateMetadataRuntimeView, resp *http.Response, err error)
	GetPingFederateMetadataCommand() (output *models.OIDCProviderMetadata, resp *http.Response, err error)
}
