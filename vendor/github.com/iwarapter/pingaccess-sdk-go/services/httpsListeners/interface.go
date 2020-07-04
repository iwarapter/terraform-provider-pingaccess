package httpsListeners

import (
	"net/http"

	"github.com/iwarapter/pingaccess-sdk-go/pingaccess/models"
)

type HttpsListenersAPI interface {
	GetHttpsListenersCommand(input *GetHttpsListenersCommandInput) (output *models.HttpsListenersView, resp *http.Response, err error)
	GetHttpsListenerCommand(input *GetHttpsListenerCommandInput) (output *models.HttpsListenerView, resp *http.Response, err error)
	UpdateHttpsListener(input *UpdateHttpsListenerInput) (output *models.HttpsListenerView, resp *http.Response, err error)
}
