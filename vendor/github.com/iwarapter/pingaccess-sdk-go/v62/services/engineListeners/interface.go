package engineListeners

import (
	"net/http"

	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess/models"
)

type EngineListenersAPI interface {
	GetEngineListenersCommand(input *GetEngineListenersCommandInput) (output *models.EngineListenersView, resp *http.Response, err error)
	AddEngineListenerCommand(input *AddEngineListenerCommandInput) (output *models.EngineListenerView, resp *http.Response, err error)
	DeleteEngineListenerCommand(input *DeleteEngineListenerCommandInput) (resp *http.Response, err error)
	GetEngineListenerCommand(input *GetEngineListenerCommandInput) (output *models.EngineListenerView, resp *http.Response, err error)
	UpdateEngineListenerCommand(input *UpdateEngineListenerCommandInput) (output *models.EngineListenerView, resp *http.Response, err error)
}
