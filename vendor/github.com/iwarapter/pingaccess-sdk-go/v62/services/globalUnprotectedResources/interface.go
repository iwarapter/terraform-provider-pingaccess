package globalUnprotectedResources

import (
	"net/http"

	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess/models"
)

type GlobalUnprotectedResourcesAPI interface {
	GetGlobalUnprotectedResourcesCommand(input *GetGlobalUnprotectedResourcesCommandInput) (output *models.GlobalUnprotectedResourcesView, resp *http.Response, err error)
	AddGlobalUnprotectedResourceCommand(input *AddGlobalUnprotectedResourceCommandInput) (output *models.GlobalUnprotectedResourceView, resp *http.Response, err error)
	DeleteGlobalUnprotectedResourceCommand(input *DeleteGlobalUnprotectedResourceCommandInput) (resp *http.Response, err error)
	GetGlobalUnprotectedResourceCommand(input *GetGlobalUnprotectedResourceCommandInput) (output *models.GlobalUnprotectedResourceView, resp *http.Response, err error)
	UpdateGlobalUnprotectedResourceCommand(input *UpdateGlobalUnprotectedResourceCommandInput) (output *models.GlobalUnprotectedResourceView, resp *http.Response, err error)
}
