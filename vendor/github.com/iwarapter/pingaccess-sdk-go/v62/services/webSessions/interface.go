package webSessions

import (
	"net/http"

	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess/models"
)

type WebSessionsAPI interface {
	GetWebSessionsCommand(input *GetWebSessionsCommandInput) (output *models.WebSessionsView, resp *http.Response, err error)
	AddWebSessionCommand(input *AddWebSessionCommandInput) (output *models.WebSessionView, resp *http.Response, err error)
	DeleteWebSessionCommand(input *DeleteWebSessionCommandInput) (resp *http.Response, err error)
	GetWebSessionCommand(input *GetWebSessionCommandInput) (output *models.WebSessionView, resp *http.Response, err error)
	UpdateWebSessionCommand(input *UpdateWebSessionCommandInput) (output *models.WebSessionView, resp *http.Response, err error)
}
