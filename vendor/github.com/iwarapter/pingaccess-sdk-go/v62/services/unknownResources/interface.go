package unknownResources

import (
	"net/http"

	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess/models"
)

type UnknownResourcesAPI interface {
	Delete() (resp *http.Response, err error)
	Get() (output *models.UnknownResourceSettingsView, resp *http.Response, err error)
	Update(input *UpdateInput) (output *models.UnknownResourceSettingsView, resp *http.Response, err error)
}
