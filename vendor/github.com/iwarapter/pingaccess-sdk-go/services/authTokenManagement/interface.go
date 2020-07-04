package authTokenManagement

import (
	"net/http"

	"github.com/iwarapter/pingaccess-sdk-go/pingaccess/models"
)

type AuthTokenManagementAPI interface {
	DeleteAuthTokenManagementCommand() (resp *http.Response, err error)
	GetAuthTokenManagementCommand() (output *models.AuthTokenManagementView, resp *http.Response, err error)
	UpdateAuthTokenManagementCommand(input *UpdateAuthTokenManagementCommandInput) (output *models.AuthTokenManagementView, resp *http.Response, err error)
	GetAuthTokenKeySetCommand() (output *models.KeySetView, resp *http.Response, err error)
	UpdateAuthTokenKeySetCommand(input *UpdateAuthTokenKeySetCommandInput) (output *models.KeySetView, resp *http.Response, err error)
}
