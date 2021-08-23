package rejectionHandlers

import (
	"net/http"

	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess/models"
)

type RejectionHandlersAPI interface {
	GetRejectionHandlersCommand(input *GetRejectionHandlersCommandInput) (output *models.RejectionHandlersView, resp *http.Response, err error)
	AddRejectionHandlerCommand(input *AddRejectionHandlerCommandInput) (output *models.RejectionHandlerView, resp *http.Response, err error)
	GetRejectionHandlerDescriptorsCommand() (output *models.DescriptorsView, resp *http.Response, err error)
	GetRejecitonHandlerDescriptorCommand(input *GetRejecitonHandlerDescriptorCommandInput) (output *models.DescriptorView, resp *http.Response, err error)
	DeleteRejectionHandlerCommand(input *DeleteRejectionHandlerCommandInput) (resp *http.Response, err error)
	GetRejectionHandlerCommand(input *GetRejectionHandlerCommandInput) (output *models.RejectionHandlerView, resp *http.Response, err error)
	UpdateRejectionHandlerCommand(input *UpdateRejectionHandlerCommandInput) (output *models.RejectionHandlerView, resp *http.Response, err error)
}
