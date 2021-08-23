package oauth

import (
	"net/http"

	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess/models"
)

type OauthAPI interface {
	DeleteAuthorizationServerCommand() (resp *http.Response, err error)
	GetAuthorizationServerCommand() (output *models.AuthorizationServerView, resp *http.Response, err error)
	UpdateAuthorizationServerCommand(input *UpdateAuthorizationServerCommandInput) (output *models.AuthorizationServerView, resp *http.Response, err error)
}
