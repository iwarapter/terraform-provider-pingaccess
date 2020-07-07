package hsmProviders

import (
	"net/http"

	"github.com/iwarapter/pingaccess-sdk-go/pingaccess/models"
)

type HsmProvidersAPI interface {
	GetHsmProvidersCommand(input *GetHsmProvidersCommandInput) (output *models.HsmProviderView, resp *http.Response, err error)
	AddHsmProviderCommand(input *AddHsmProviderCommandInput) (output *models.HsmProviderView, resp *http.Response, err error)
	GetHsmProviderDescriptorsCommand() (output *models.DescriptorsView, resp *http.Response, err error)
	DeleteHsmProviderCommand(input *DeleteHsmProviderCommandInput) (resp *http.Response, err error)
	GetHsmProviderCommand(input *GetHsmProviderCommandInput) (output *models.HsmProviderView, resp *http.Response, err error)
	UpdateHsmProviderCommand(input *UpdateHsmProviderCommandInput) (output *models.HsmProviderView, resp *http.Response, err error)
}
