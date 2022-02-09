package sdkv2provider

import (
	"context"
	"strconv"

	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess/models"
	"github.com/iwarapter/pingaccess-sdk-go/v62/services/rulesets"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePingAccessRuleSet() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePingAccessRuleSetCreate,
		ReadContext:   resourcePingAccessRuleSetRead,
		UpdateContext: resourcePingAccessRuleSetUpdate,
		DeleteContext: resourcePingAccessRuleSetDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema:      resourcePingAccessRuleSetSchema(),
		Description: `Provides configuration for Rulesets within PingAccess.`,
	}
}

func resourcePingAccessRuleSetSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"element_type": {
			Type:             schema.TypeString,
			Required:         true,
			ValidateDiagFunc: validateRuleOrRuleSet,
			Description:      "The rule set's element type (what it contains). Can be either `Rule` or `Ruleset`.",
		},
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The rule set's name.",
		},
		"policy": {
			Type:        schema.TypeList,
			Required:    true,
			MinItems:    1,
			Description: "The list of policy ids assigned to the rule set.",
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"success_criteria": {
			Type:             schema.TypeString,
			Required:         true,
			ValidateDiagFunc: validateSuccessIfAllSucceedOrSuccessIfAnyOneSucceeds,
			Description:      "The rule set's success criteria. Can be either `SuccessIfAllSucceed` or `SuccessIfAnyOneSucceeds`.",
		},
	}
}

func resourcePingAccessRuleSetCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	name := d.Get("name").(string)
	elementType := d.Get("element_type").(string)
	policy := expandStringList(d.Get("policy").([]interface{}))
	successCriteria := d.Get("success_criteria").(string)

	//TODO generalise this into a helper function
	var polIds []*int
	for _, i := range policy {
		text, _ := strconv.Atoi(*i)
		polIds = append(polIds, &text)
	}

	input := rulesets.AddRuleSetCommandInput{
		Body: models.RuleSetView{
			Name:            String(name),
			ElementType:     String(elementType),
			Policy:          &polIds,
			SuccessCriteria: String(successCriteria),
		},
	}

	svc := m.(paClient).Rulesets

	result, _, err := svc.AddRuleSetCommand(&input)
	if err != nil {
		return diag.Errorf("unable to create RuleSet: %s", err)
	}

	d.SetId(result.Id.String())
	return resourcePingAccessRuleSetReadResult(d, result)
}

func resourcePingAccessRuleSetRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(paClient).Rulesets
	result, _, err := svc.GetRuleSetCommand(&rulesets.GetRuleSetCommandInput{
		Id: d.Id(),
	})
	if err != nil {
		return diag.Errorf("unable to read RuleSet: %s", err)
	}
	return resourcePingAccessRuleSetReadResult(d, result)
}

func resourcePingAccessRuleSetUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	name := d.Get("name").(string)
	elementType := d.Get("element_type").(string)
	policy := expandStringList(d.Get("policy").([]interface{}))
	successCriteria := d.Get("success_criteria").(string)

	//TODO generalise this into a helper function
	var polIds []*int
	for _, i := range policy {
		text, _ := strconv.Atoi(*i)
		polIds = append(polIds, &text)
	}

	input := rulesets.UpdateRuleSetCommandInput{
		Body: models.RuleSetView{
			Name:            String(name),
			ElementType:     String(elementType),
			Policy:          &polIds,
			SuccessCriteria: String(successCriteria),
		},
		Id: d.Id(),
	}

	svc := m.(paClient).Rulesets

	result, _, err := svc.UpdateRuleSetCommand(&input)
	if err != nil {
		return diag.Errorf("unable to update RuleSet: %s", err)
	}
	d.SetId(result.Id.String())
	return resourcePingAccessRuleSetReadResult(d, result)
}

func resourcePingAccessRuleSetDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(paClient).Rulesets

	_, err := svc.DeleteRuleSetCommand(&rulesets.DeleteRuleSetCommandInput{Id: d.Id()})
	if err != nil {
		return diag.Errorf("unable to delete RuleSet: %s", err)
	}
	return nil
}

func resourcePingAccessRuleSetReadResult(d *schema.ResourceData, input *models.RuleSetView) diag.Diagnostics {
	var diags diag.Diagnostics
	setResourceDataStringWithDiagnostic(d, "name", input.Name, &diags)
	setResourceDataStringWithDiagnostic(d, "success_criteria", input.SuccessCriteria, &diags)
	setResourceDataStringWithDiagnostic(d, "element_type", input.ElementType, &diags)
	if input.Policy != nil {
		var polIds []string
		for _, p := range *input.Policy {
			polIds = append(polIds, strconv.Itoa(*p))
		}
		if err := d.Set("policy", polIds); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}
	return diags
}
