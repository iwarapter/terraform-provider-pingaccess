package tokenProvider

import (
	"net/http"

	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess/models"
)

type TokenProviderAPI interface {
	DeleteTokenProviderSettingCommand() (resp *http.Response, err error)
	GetTokenProviderSettingCommand() (output *models.TokenProviderSettingView, resp *http.Response, err error)
	UpdateTokenProviderSettingCommand(input *UpdateTokenProviderSettingCommandInput) (output *models.TokenProviderSettingView, resp *http.Response, err error)
}
