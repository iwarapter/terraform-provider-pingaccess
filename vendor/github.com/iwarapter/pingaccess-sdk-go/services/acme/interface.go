package acme

import (
	"net/http"

	"github.com/iwarapter/pingaccess-sdk-go/pingaccess/models"
)

type AcmeAPI interface {
	GetAcmeServersCommand(input *GetAcmeServersCommandInput) (output *models.AcmeServersView, resp *http.Response, err error)
	AddAcmeServerCommand(input *AddAcmeServerCommandInput) (output *models.AcmeServerView, resp *http.Response, err error)
	GetDefaultAcmeServerCommand() (output *models.LinkView, resp *http.Response, err error)
	UpdateDefaultAcmeServerCommand(input *UpdateDefaultAcmeServerCommandInput) (output *models.LinkView, resp *http.Response, err error)
	DeleteAcmeServerCommand(input *DeleteAcmeServerCommandInput) (output *models.AcmeServerView, resp *http.Response, err error)
	GetAcmeServerCommand(input *GetAcmeServerCommandInput) (output *models.AcmeServerView, resp *http.Response, err error)
	GetAcmeAccountsCommand(input *GetAcmeAccountsCommandInput) (output *models.AcmeAccountView, resp *http.Response, err error)
	AddAcmeAccountCommand(input *AddAcmeAccountCommandInput) (output *models.AcmeAccountView, resp *http.Response, err error)
	DeleteAcmeAccountCommand(input *DeleteAcmeAccountCommandInput) (output *models.AcmeAccountView, resp *http.Response, err error)
	GetAcmeAccountCommand(input *GetAcmeAccountCommandInput) (output *models.AcmeAccountView, resp *http.Response, err error)
	GetAcmeCertificateRequestsCommand(input *GetAcmeCertificateRequestsCommandInput) (output *models.AcmeCertificateRequestView, resp *http.Response, err error)
	AddAcmeCertificateRequestCommand(input *AddAcmeCertificateRequestCommandInput) (output *models.AcmeCertificateRequestView, resp *http.Response, err error)
	DeleteAcmeCertificateRequestCommand(input *DeleteAcmeCertificateRequestCommandInput) (output *models.AcmeCertificateRequestView, resp *http.Response, err error)
	GetAcmeCertificateRequestCommand(input *GetAcmeCertificateRequestCommandInput) (output *models.AcmeCertificateRequestView, resp *http.Response, err error)
}
