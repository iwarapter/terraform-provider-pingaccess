package auth

import (
	"net/http"

	"github.com/iwarapter/pingaccess-sdk-go/pingaccess/models"
)

type AuthAPI interface {
	DeleteBasicAuthCommand() (resp *http.Response, err error)
	GetBasicAuthCommand() (output *models.BasicConfig, resp *http.Response, err error)
	UpdateBasicAuthCommand(input *UpdateBasicAuthCommandInput) (output *models.BasicAuthConfigView, resp *http.Response, err error)
	DeleteOAuthAuthCommand() (resp *http.Response, err error)
	GetOAuthAuthCommand() (output *models.OAuthConfigView, resp *http.Response, err error)
	UpdateOAuthAuthCommand(input *UpdateOAuthAuthCommandInput) (output *models.OAuthConfigView, resp *http.Response, err error)
	DeleteOidcAuthCommand() (resp *http.Response, err error)
	GetOidcAuthCommand() (output *models.OidcConfigView, resp *http.Response, err error)
	UpdateOidcAuthCommand(input *UpdateOidcAuthCommandInput) (output *models.OidcConfigView, resp *http.Response, err error)
	DeleteAdminBasicWebSessionCommand() (resp *http.Response, err error)
	GetAdminBasicWebSessionCommand() (output *models.AdminBasicWebSessionView, resp *http.Response, err error)
	UpdateAdminBasicWebSessionCommand(input *UpdateAdminBasicWebSessionCommandInput) (output *models.AdminBasicWebSessionView, resp *http.Response, err error)
}
