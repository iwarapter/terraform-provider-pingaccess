package applications

import (
	"net/http"

	"github.com/iwarapter/pingaccess-sdk-go/pingaccess/models"
)

type ApplicationsAPI interface {
	GetApplicationsCommand(input *GetApplicationsCommandInput) (output *models.ApplicationsView, resp *http.Response, err error)
	AddApplicationCommand(input *AddApplicationCommandInput) (output *models.ApplicationView, resp *http.Response, err error)
	DeleteReservedApplicationCommand() (resp *http.Response, err error)
	GetReservedApplicationCommand() (output *models.ReservedApplicationView, resp *http.Response, err error)
	UpdateReservedApplicationCommand(input *UpdateReservedApplicationCommandInput) (output *models.ReservedApplicationView, resp *http.Response, err error)
	GetResourcesCommand(input *GetResourcesCommandInput) (output *models.ResourcesView, resp *http.Response, err error)
	GetApplicationsResourcesMethodsCommand() (output *models.MethodsView, resp *http.Response, err error)
	GetApplicationResourceResponseGeneratorDescriptorsCommand() (output *models.DescriptorsView, resp *http.Response, err error)
	GetApplicationResourceResponseGeneratorDescriptorCommand(input *GetApplicationResourceResponseGeneratorDescriptorCommandInput) (output *models.DescriptorView, resp *http.Response, err error)
	DeleteApplicationResourceCommand(input *DeleteApplicationResourceCommandInput) (resp *http.Response, err error)
	GetApplicationResourceCommand(input *GetApplicationResourceCommandInput) (output *models.ResourceView, resp *http.Response, err error)
	UpdateApplicationResourceCommand(input *UpdateApplicationResourceCommandInput) (output *models.ResourceView, resp *http.Response, err error)
	DeleteApplicationCommand(input *DeleteApplicationCommandInput) (resp *http.Response, err error)
	GetApplicationCommand(input *GetApplicationCommandInput) (output *models.ApplicationView, resp *http.Response, err error)
	UpdateApplicationCommand(input *UpdateApplicationCommandInput) (output *models.ApplicationView, resp *http.Response, err error)
	GetResourceMatchingEvaluationOrderCommand(input *GetResourceMatchingEvaluationOrderCommandInput) (output *models.ResourceMatchingEvaluationOrderView, resp *http.Response, err error)
	GetApplicationResourcesCommand(input *GetApplicationResourcesCommandInput) (output *models.ResourcesView, resp *http.Response, err error)
	AddApplicationResourceCommand(input *AddApplicationResourceCommandInput) (output *models.ResourceView, resp *http.Response, err error)
	GetResourceAutoOrderCommand(input *GetResourceAutoOrderCommandInput) (output *models.ResourceOrderView, resp *http.Response, err error)
}
