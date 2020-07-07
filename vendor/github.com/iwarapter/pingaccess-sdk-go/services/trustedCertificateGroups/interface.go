package trustedCertificateGroups

import (
	"net/http"

	"github.com/iwarapter/pingaccess-sdk-go/pingaccess/models"
)

type TrustedCertificateGroupsAPI interface {
	GetTrustedCertificateGroupsCommand(input *GetTrustedCertificateGroupsCommandInput) (output *models.TrustedCertificateGroupsView, resp *http.Response, err error)
	AddTrustedCertificateGroupCommand(input *AddTrustedCertificateGroupCommandInput) (output *models.TrustedCertificateGroupView, resp *http.Response, err error)
	DeleteTrustedCertificateGroupCommand(input *DeleteTrustedCertificateGroupCommandInput) (resp *http.Response, err error)
	GetTrustedCertificateGroupCommand(input *GetTrustedCertificateGroupCommandInput) (output *models.TrustedCertificateGroupView, resp *http.Response, err error)
	UpdateTrustedCertificateGroupCommand(input *UpdateTrustedCertificateGroupCommandInput) (output *models.TrustedCertificateGroupView, resp *http.Response, err error)
}
