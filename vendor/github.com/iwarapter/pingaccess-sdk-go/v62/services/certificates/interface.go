package certificates

import (
	"net/http"

	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess/models"
)

type CertificatesAPI interface {
	GetTrustedCerts(input *GetTrustedCertsInput) (output *models.TrustedCertsView, resp *http.Response, err error)
	ImportTrustedCert(input *ImportTrustedCertInput) (output *models.TrustedCertView, resp *http.Response, err error)
	DeleteTrustedCertCommand(input *DeleteTrustedCertCommandInput) (resp *http.Response, err error)
	GetTrustedCert(input *GetTrustedCertInput) (output *models.TrustedCertView, resp *http.Response, err error)
	UpdateTrustedCert(input *UpdateTrustedCertInput) (output *models.TrustedCertView, resp *http.Response, err error)
	ExportTrustedCert(input *ExportTrustedCertInput) (output *string, resp *http.Response, err error)
}
