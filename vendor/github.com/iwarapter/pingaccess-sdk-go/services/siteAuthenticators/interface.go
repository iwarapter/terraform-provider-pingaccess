package siteAuthenticators

import (
	"net/http"

	"github.com/iwarapter/pingaccess-sdk-go/pingaccess/models"
)

type SiteAuthenticatorsAPI interface {
	GetSiteAuthenticatorsCommand(input *GetSiteAuthenticatorsCommandInput) (output *models.SiteAuthenticatorsView, resp *http.Response, err error)
	AddSiteAuthenticatorCommand(input *AddSiteAuthenticatorCommandInput) (output *models.SiteAuthenticatorView, resp *http.Response, err error)
	GetSiteAuthenticatorDescriptorsCommand() (output *models.DescriptorsView, resp *http.Response, err error)
	GetSiteAuthenticatorDescriptorCommand(input *GetSiteAuthenticatorDescriptorCommandInput) (output *models.DescriptorView, resp *http.Response, err error)
	DeleteSiteAuthenticatorCommand(input *DeleteSiteAuthenticatorCommandInput) (resp *http.Response, err error)
	GetSiteAuthenticatorCommand(input *GetSiteAuthenticatorCommandInput) (output *models.SiteAuthenticatorView, resp *http.Response, err error)
	UpdateSiteAuthenticatorCommand(input *UpdateSiteAuthenticatorCommandInput) (output *models.SiteAuthenticatorView, resp *http.Response, err error)
}
