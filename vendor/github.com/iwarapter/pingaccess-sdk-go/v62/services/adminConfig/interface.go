package adminConfig

import (
	"net/http"

	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess/models"
)

type AdminConfigAPI interface {
	DeleteAdminConfigurationCommand() (resp *http.Response, err error)
	GetAdminConfigurationCommand() (output *models.AdminConfigurationView, resp *http.Response, err error)
	UpdateAdminConfigurationCommand(input *UpdateAdminConfigurationCommandInput) (output *models.AdminConfigurationView, resp *http.Response, err error)
	GetReplicaAdminsCommand() (output *models.ReplicaAdminsView, resp *http.Response, err error)
	AddReplicaAdminCommand(input *AddReplicaAdminCommandInput) (output *models.ReplicaAdminView, resp *http.Response, err error)
	DeleteReplicaAdminCommand(input *DeleteReplicaAdminCommandInput) (resp *http.Response, err error)
	GetReplicaAdminCommand(input *GetReplicaAdminCommandInput) (output *models.ReplicaAdminView, resp *http.Response, err error)
	UpdateAdminReplicaCommand(input *UpdateAdminReplicaCommandInput) (output *models.ReplicaAdminView, resp *http.Response, err error)
	GetAdminReplicaFileCommand(input *GetAdminReplicaFileCommandInput) (resp *http.Response, err error)
}
