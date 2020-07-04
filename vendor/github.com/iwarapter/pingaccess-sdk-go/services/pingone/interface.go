package pingone

import (
	"net/http"

	"github.com/iwarapter/pingaccess-sdk-go/pingaccess/models"
)

type PingoneAPI interface {
	DeletePingOne4CCommand() (resp *http.Response, err error)
	GetPingOne4CCommand() (output *models.PingOne4CView, resp *http.Response, err error)
	UpdatePingOne4CCommand(input *UpdatePingOne4CCommandInput) (output *models.PingOne4CView, resp *http.Response, err error)
	GetPingOne4CMetadataCommand() (output *models.OIDCProviderMetadata, resp *http.Response, err error)
}
