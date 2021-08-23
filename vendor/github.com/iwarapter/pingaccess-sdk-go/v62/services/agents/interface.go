package agents

import (
	"net/http"

	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess/models"
)

type AgentsAPI interface {
	GetAgentsCommand(input *GetAgentsCommandInput) (output *models.AgentsView, resp *http.Response, err error)
	AddAgentCommand(input *AddAgentCommandInput) (output *models.AgentView, resp *http.Response, err error)
	GetAgentCertificatesCommand(input *GetAgentCertificatesCommandInput) (output *models.AgentCertificatesView, resp *http.Response, err error)
	GetAgentCertificateCommand(input *GetAgentCertificateCommandInput) (output *models.AgentCertificateView, resp *http.Response, err error)
	GetAgentFileCommand(input *GetAgentFileCommandInput) (resp *http.Response, err error)
	DeleteAgentCommand(input *DeleteAgentCommandInput) (resp *http.Response, err error)
	GetAgentCommand(input *GetAgentCommandInput) (output *models.AgentView, resp *http.Response, err error)
	UpdateAgentCommand(input *UpdateAgentCommandInput) (output *models.AgentView, resp *http.Response, err error)
}
