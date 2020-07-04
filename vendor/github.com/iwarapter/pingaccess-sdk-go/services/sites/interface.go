package sites

import (
	"net/http"

	"github.com/iwarapter/pingaccess-sdk-go/pingaccess/models"
)

type SitesAPI interface {
	GetSitesCommand(input *GetSitesCommandInput) (output *models.SitesView, resp *http.Response, err error)
	AddSiteCommand(input *AddSiteCommandInput) (output *models.SiteView, resp *http.Response, err error)
	DeleteSiteCommand(input *DeleteSiteCommandInput) (resp *http.Response, err error)
	GetSiteCommand(input *GetSiteCommandInput) (output *models.SiteView, resp *http.Response, err error)
	UpdateSiteCommand(input *UpdateSiteCommandInput) (output *models.SiteView, resp *http.Response, err error)
}
