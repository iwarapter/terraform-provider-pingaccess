package rules

import (
	"net/http"

	"github.com/iwarapter/pingaccess-sdk-go/pingaccess/models"
)

type RulesAPI interface {
	GetRulesCommand(input *GetRulesCommandInput) (output *models.RulesView, resp *http.Response, err error)
	AddRuleCommand(input *AddRuleCommandInput) (output *models.RuleView, resp *http.Response, err error)
	GetRuleDescriptorsCommand() (output *models.RuleDescriptorsView, resp *http.Response, err error)
	GetRuleDescriptorCommand(input *GetRuleDescriptorCommandInput) (output *models.RuleDescriptorView, resp *http.Response, err error)
	DeleteRuleCommand(input *DeleteRuleCommandInput) (resp *http.Response, err error)
	GetRuleCommand(input *GetRuleCommandInput) (output *models.RuleView, resp *http.Response, err error)
	UpdateRuleCommand(input *UpdateRuleCommandInput) (output *models.RuleView, resp *http.Response, err error)
}
