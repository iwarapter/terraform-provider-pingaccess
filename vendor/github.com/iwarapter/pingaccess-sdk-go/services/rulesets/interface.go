package rulesets

import (
	"net/http"

	"github.com/iwarapter/pingaccess-sdk-go/pingaccess/models"
)

type RulesetsAPI interface {
	GetRuleSetsCommand(input *GetRuleSetsCommandInput) (output *models.RuleSetsView, resp *http.Response, err error)
	AddRuleSetCommand(input *AddRuleSetCommandInput) (output *models.RuleSetView, resp *http.Response, err error)
	GetRuleSetElementTypesCommand() (output *models.RuleSetElementTypesView, resp *http.Response, err error)
	GetRuleSetSuccessCriteriaCommand() (output *models.RuleSetSuccessCriteriaView, resp *http.Response, err error)
	DeleteRuleSetCommand(input *DeleteRuleSetCommandInput) (resp *http.Response, err error)
	GetRuleSetCommand(input *GetRuleSetCommandInput) (output *models.RuleSetView, resp *http.Response, err error)
	UpdateRuleSetCommand(input *UpdateRuleSetCommandInput) (output *models.RuleSetView, resp *http.Response, err error)
}
