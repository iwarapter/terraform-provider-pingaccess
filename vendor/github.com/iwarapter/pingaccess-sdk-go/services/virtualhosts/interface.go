package virtualhosts

import (
	"net/http"

	"github.com/iwarapter/pingaccess-sdk-go/pingaccess/models"
)

type VirtualhostsAPI interface {
	GetVirtualHostsCommand(input *GetVirtualHostsCommandInput) (output *models.VirtualHostsView, resp *http.Response, err error)
	AddVirtualHostCommand(input *AddVirtualHostCommandInput) (output *models.VirtualHostView, resp *http.Response, err error)
	DeleteVirtualHostCommand(input *DeleteVirtualHostCommandInput) (resp *http.Response, err error)
	GetVirtualHostCommand(input *GetVirtualHostCommandInput) (output *models.VirtualHostView, resp *http.Response, err error)
	UpdateVirtualHostCommand(input *UpdateVirtualHostCommandInput) (output *models.VirtualHostView, resp *http.Response, err error)
}
