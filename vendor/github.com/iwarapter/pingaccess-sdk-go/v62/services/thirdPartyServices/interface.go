package thirdPartyServices

import (
	"net/http"

	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess/models"
)

type ThirdPartyServicesAPI interface {
	GetThirdPartyServicesCommand(input *GetThirdPartyServicesCommandInput) (output *models.ThirdPartyServicesView, resp *http.Response, err error)
	AddThirdPartyServiceCommand(input *AddThirdPartyServiceCommandInput) (output *models.ThirdPartyServiceView, resp *http.Response, err error)
	DeleteThirdPartyServiceCommand(input *DeleteThirdPartyServiceCommandInput) (resp *http.Response, err error)
	GetThirdPartyServiceCommand(input *GetThirdPartyServiceCommandInput) (output *models.ThirdPartyServiceView, resp *http.Response, err error)
	UpdateThirdPartyServiceCommand(input *UpdateThirdPartyServiceCommandInput) (output *models.ThirdPartyServiceView, resp *http.Response, err error)
}
