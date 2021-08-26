package highAvailability

import (
	"net/http"

	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess/models"
)

type HighAvailabilityAPI interface {
	GetAvailabilityProfilesCommand(input *GetAvailabilityProfilesCommandInput) (output *models.AvailabilityProfilesView, resp *http.Response, err error)
	AddAvailabilityProfileCommand(input *AddAvailabilityProfileCommandInput) (output *models.AvailabilityProfileView, resp *http.Response, err error)
	GetAvailabilityProfileDescriptorsCommand() (output *models.DescriptorsView, resp *http.Response, err error)
	GetAvailabilityProfileDescriptorCommand(input *GetAvailabilityProfileDescriptorCommandInput) (output *models.DescriptorView, resp *http.Response, err error)
	DeleteAvailabilityProfileCommand(input *DeleteAvailabilityProfileCommandInput) (resp *http.Response, err error)
	GetAvailabilityProfileCommand(input *GetAvailabilityProfileCommandInput) (output *models.AvailabilityProfileView, resp *http.Response, err error)
	UpdateAvailabilityProfileCommand(input *UpdateAvailabilityProfileCommandInput) (output *models.AvailabilityProfileView, resp *http.Response, err error)
	GetLoadBalancingStrategiesCommand(input *GetLoadBalancingStrategiesCommandInput) (output *models.LoadBalancingStrategiesView, resp *http.Response, err error)
	AddLoadBalancingStrategyCommand(input *AddLoadBalancingStrategyCommandInput) (output *models.LoadBalancingStrategyView, resp *http.Response, err error)
	GetLoadBalancingStrategyDescriptorsCommand() (output *models.DescriptorsView, resp *http.Response, err error)
	GetLoadBalancingStrategyDescriptorCommand(input *GetLoadBalancingStrategyDescriptorCommandInput) (output *models.DescriptorView, resp *http.Response, err error)
	DeleteLoadBalancingStrategyCommand(input *DeleteLoadBalancingStrategyCommandInput) (resp *http.Response, err error)
	GetLoadBalancingStrategyCommand(input *GetLoadBalancingStrategyCommandInput) (output *models.LoadBalancingStrategyView, resp *http.Response, err error)
	UpdateLoadBalancingStrategyCommand(input *UpdateLoadBalancingStrategyCommandInput) (output *models.LoadBalancingStrategyView, resp *http.Response, err error)
}
