package authnReqLists

import (
	"net/http"

	"github.com/iwarapter/pingaccess-sdk-go/pingaccess/models"
)

type AuthnReqListsAPI interface {
	GetAuthnReqListsCommand(input *GetAuthnReqListsCommandInput) (output *models.AuthnReqListsView, resp *http.Response, err error)
	AddAuthnReqListCommand(input *AddAuthnReqListCommandInput) (output *models.AuthnReqListView, resp *http.Response, err error)
	DeleteAuthnReqListCommand(input *DeleteAuthnReqListCommandInput) (resp *http.Response, err error)
	GetAuthnReqListCommand(input *GetAuthnReqListCommandInput) (output *models.AuthnReqListView, resp *http.Response, err error)
	UpdateAuthnReqListCommand(input *UpdateAuthnReqListCommandInput) (output *models.AuthnReqListView, resp *http.Response, err error)
}
