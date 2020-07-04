package config

import (
	"net/http"

	"github.com/iwarapter/pingaccess-sdk-go/pingaccess/models"
)

type ConfigAPI interface {
	ConfigExportCommand() (output *models.ExportData, resp *http.Response, err error)
	GetConfigExportWorkflowsCommand() (output *models.ConfigStatusesView, resp *http.Response, err error)
	AddConfigExportWorkflowCommand() (output *models.ConfigStatusView, resp *http.Response, err error)
	GetConfigExportWorkflowCommand(input *GetConfigExportWorkflowCommandInput) (output *models.ConfigStatusView, resp *http.Response, err error)
	GetConfigExportWorkflowDataCommand(input *GetConfigExportWorkflowDataCommandInput) (output *models.ExportData, resp *http.Response, err error)
	ConfigImportCommand(input *ConfigImportCommandInput) (resp *http.Response, err error)
	GetConfigImportWorkflowsCommand() (output *models.ConfigStatusesView, resp *http.Response, err error)
	AddConfigImportWorkflowCommand(input *AddConfigImportWorkflowCommandInput) (resp *http.Response, err error)
	GetConfigImportWorkflowCommand(input *GetConfigImportWorkflowCommandInput) (output *models.ConfigStatusView, resp *http.Response, err error)
}
