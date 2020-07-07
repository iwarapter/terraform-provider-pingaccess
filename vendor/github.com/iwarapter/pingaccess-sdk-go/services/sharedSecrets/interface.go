package sharedSecrets

import (
	"net/http"

	"github.com/iwarapter/pingaccess-sdk-go/pingaccess/models"
)

type SharedSecretsAPI interface {
	GetSharedSecretsCommand(input *GetSharedSecretsCommandInput) (output *models.SharedSecretsView, resp *http.Response, err error)
	AddSharedSecretCommand(input *AddSharedSecretCommandInput) (output *models.SharedSecretView, resp *http.Response, err error)
	DeleteSharedSecretCommand(input *DeleteSharedSecretCommandInput) (resp *http.Response, err error)
	GetSharedSecretCommand(input *GetSharedSecretCommandInput) (output *models.SharedSecretView, resp *http.Response, err error)
}
