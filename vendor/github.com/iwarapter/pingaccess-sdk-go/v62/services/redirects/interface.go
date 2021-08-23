package redirects

import (
	"net/http"

	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess/models"
)

type RedirectsAPI interface {
	GetRedirectsCommand(input *GetRedirectsCommandInput) (output *models.RedirectsView, resp *http.Response, err error)
	AddRedirectCommand(input *AddRedirectCommandInput) (output *models.RedirectView, resp *http.Response, err error)
	DeleteRedirectCommand(input *DeleteRedirectCommandInput) (resp *http.Response, err error)
	GetRedirectCommand(input *GetRedirectCommandInput) (output *models.RedirectView, resp *http.Response, err error)
	UpdateRedirectCommand(input *UpdateRedirectCommandInput) (output *models.RedirectView, resp *http.Response, err error)
}
