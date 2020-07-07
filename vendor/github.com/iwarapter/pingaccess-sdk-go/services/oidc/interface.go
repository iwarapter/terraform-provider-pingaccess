package oidc

import (
	"net/http"

	"github.com/iwarapter/pingaccess-sdk-go/pingaccess/models"
)

type OidcAPI interface {
	DeleteOIDCProviderCommand() (resp *http.Response, err error)
	GetOIDCProviderCommand() (output *models.OIDCProviderView, resp *http.Response, err error)
	UpdateOIDCProviderCommand(input *UpdateOIDCProviderCommandInput) (output *models.OIDCProviderView, resp *http.Response, err error)
	GetOIDCProviderPluginDescriptorsCommand() (output *models.DescriptorsView, resp *http.Response, err error)
	GetOIDCProviderPluginDescriptorCommand(input *GetOIDCProviderPluginDescriptorCommandInput) (output *models.DescriptorView, resp *http.Response, err error)
	GetOIDCProviderMetadataCommand() (output *models.OIDCProviderMetadata, resp *http.Response, err error)
}
