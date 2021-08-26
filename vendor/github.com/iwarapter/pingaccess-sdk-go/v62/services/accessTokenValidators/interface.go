package accessTokenValidators

import (
	"net/http"

	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess/models"
)

type AccessTokenValidatorsAPI interface {
	GetAccessTokenValidatorsCommand(input *GetAccessTokenValidatorsCommandInput) (output *models.AccessTokenValidatorsView, resp *http.Response, err error)
	AddAccessTokenValidatorCommand(input *AddAccessTokenValidatorCommandInput) (output *models.AccessTokenValidatorView, resp *http.Response, err error)
	GetAccessTokenValidatorDescriptorsCommand() (output *models.DescriptorsView, resp *http.Response, err error)
	DeleteAccessTokenValidatorCommand(input *DeleteAccessTokenValidatorCommandInput) (resp *http.Response, err error)
	GetAccessTokenValidatorCommand(input *GetAccessTokenValidatorCommandInput) (output *models.AccessTokenValidatorView, resp *http.Response, err error)
	UpdateAccessTokenValidatorCommand(input *UpdateAccessTokenValidatorCommandInput) (output *models.AccessTokenValidatorView, resp *http.Response, err error)
}
