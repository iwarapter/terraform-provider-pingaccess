package proxies

import (
	"net/http"

	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess/models"
)

type ProxiesAPI interface {
	GetProxiesCommand(input *GetProxiesCommandInput) (output *models.HttpClientProxyView, resp *http.Response, err error)
	AddProxyCommand(input *AddProxyCommandInput) (output *models.HttpClientProxyView, resp *http.Response, err error)
	DeleteProxyCommand(input *DeleteProxyCommandInput) (resp *http.Response, err error)
	GetProxyCommand(input *GetProxyCommandInput) (output *models.HttpClientProxyView, resp *http.Response, err error)
	UpdateProxyCommand(input *UpdateProxyCommandInput) (output *models.HttpClientProxyView, resp *http.Response, err error)
}
