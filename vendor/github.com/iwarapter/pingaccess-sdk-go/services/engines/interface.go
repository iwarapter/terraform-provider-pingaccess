package engines

import (
	"net/http"

	"github.com/iwarapter/pingaccess-sdk-go/pingaccess/models"
)

type EnginesAPI interface {
	GetEnginesCommand(input *GetEnginesCommandInput) (output *models.EnginesView, resp *http.Response, err error)
	AddEngineCommand(input *AddEngineCommandInput) (output *models.EngineView, resp *http.Response, err error)
	GetEngineCertificatesCommand(input *GetEngineCertificatesCommandInput) (output *models.EngineCertificateView, resp *http.Response, err error)
	GetEngineCertificateCommand(input *GetEngineCertificateCommandInput) (output *models.EngineCertificateView, resp *http.Response, err error)
	GetEngineStatusCommand() (output *models.EngineHealthStatusView, resp *http.Response, err error)
	DeleteEngineCommand(input *DeleteEngineCommandInput) (resp *http.Response, err error)
	GetEngineCommand(input *GetEngineCommandInput) (output *models.EngineView, resp *http.Response, err error)
	UpdateEngineCommand(input *UpdateEngineCommandInput) (output *models.EngineView, resp *http.Response, err error)
	GetEngineConfigFileCommand(input *GetEngineConfigFileCommandInput) (resp *http.Response, err error)
}
