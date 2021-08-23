package auth

import (
	"net/http"

	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess/models"
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
	GetAuthOidcScopesCommand(input *GetAuthOidcScopesCommandInput) (output *models.SupportedScopesView, resp *http.Response, err error)
	DeleteAdminTokenProviderCommand() (resp *http.Response, err error)
	GetAdminTokenProviderCommand() (output *models.AdminTokenProviderView, resp *http.Response, err error)
	UpdateAdminTokenProviderCommand(input *UpdateAdminTokenProviderCommandInput) (output *models.AdminTokenProviderView, resp *http.Response, err error)
	GetAdminTokenProviderMetadataCommand() (output *models.OIDCProviderMetadata, resp *http.Response, err error)
	DeleteAdminBasicWebSessionCommand() (resp *http.Response, err error)
	GetAdminBasicWebSessionCommand() (output *models.AdminBasicWebSessionView, resp *http.Response, err error)
	UpdateAdminBasicWebSessionCommand(input *UpdateAdminBasicWebSessionCommandInput) (output *models.AdminBasicWebSessionView, resp *http.Response, err error)
}
