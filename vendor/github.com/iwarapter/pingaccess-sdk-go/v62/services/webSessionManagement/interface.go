package webSessionManagement

import (
	"net/http"

	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess/models"
)

type WebSessionManagementAPI interface {
	DeleteWebSessionManagementCommand() (resp *http.Response, err error)
	GetWebSessionManagementCommand() (output *models.WebSessionManagementView, resp *http.Response, err error)
	UpdateWebSessionManagementCommand(input *UpdateWebSessionManagementCommandInput) (output *models.WebSessionManagementView, resp *http.Response, err error)
	GetCookieTypes() (output *models.CookieTypesView, resp *http.Response, err error)
	GetWebSessionSupportedEncryptionAlgorithmsCommand() (output *models.AlgorithmsView, resp *http.Response, err error)
	GetWebSessionKeySetCommand() (output *models.KeySetView, resp *http.Response, err error)
	UpdateWebSessionKeySetCommand(input *UpdateWebSessionKeySetCommandInput) (output *models.KeySetView, resp *http.Response, err error)
	GetOidcLoginTypes() (output *models.OidcLoginTypesView, resp *http.Response, err error)
	GetOidcScopesCommand(input *GetOidcScopesCommandInput) (output *models.SupportedScopesView, resp *http.Response, err error)
	GetRequestPreservationTypes() (output *models.RequestPreservationTypesView, resp *http.Response, err error)
	GetWebSessionSupportedSigningAlgorithms() (output *models.SigningAlgorithmsView, resp *http.Response, err error)
	GetWebStorageTypes() (output *models.WebStorageTypesView, resp *http.Response, err error)
}
