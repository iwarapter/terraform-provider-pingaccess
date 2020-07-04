package identityMappings

import (
	"net/http"

	"github.com/iwarapter/pingaccess-sdk-go/pingaccess/models"
)

type IdentityMappingsAPI interface {
	GetIdentityMappingsCommand(input *GetIdentityMappingsCommandInput) (output *models.IdentityMappingsView, resp *http.Response, err error)
	AddIdentityMappingCommand(input *AddIdentityMappingCommandInput) (output *models.IdentityMappingView, resp *http.Response, err error)
	GetIdentityMappingDescriptorsCommand() (output *models.DescriptorsView, resp *http.Response, err error)
	GetIdentityMappingDescriptorCommand(input *GetIdentityMappingDescriptorCommandInput) (output *models.DescriptorView, resp *http.Response, err error)
	DeleteIdentityMappingCommand(input *DeleteIdentityMappingCommandInput) (resp *http.Response, err error)
	GetIdentityMappingCommand(input *GetIdentityMappingCommandInput) (output *models.IdentityMappingView, resp *http.Response, err error)
	UpdateIdentityMappingCommand(input *UpdateIdentityMappingCommandInput) (output *models.IdentityMappingView, resp *http.Response, err error)
}
