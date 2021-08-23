package oauthKeyManagement

import (
	"net/http"

	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess/models"
)

type OauthKeyManagementAPI interface {
	DeleteOAuthKeyManagementCommand() (resp *http.Response, err error)
	GetOAuthKeyManagementCommand() (output *models.OAuthKeyManagementView, resp *http.Response, err error)
	UpdateOAuthKeyManagementCommand(input *UpdateOAuthKeyManagementCommandInput) (output *models.OAuthKeyManagementView, resp *http.Response, err error)
	GetOAuthKeySetCommand() (output *models.KeySetView, resp *http.Response, err error)
	UpdateOAuthKeySetCommand(input *UpdateOAuthKeySetCommandInput) (output *models.KeySetView, resp *http.Response, err error)
}
