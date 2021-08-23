package sdkv2provider

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess/models"
	"github.com/iwarapter/pingaccess-sdk-go/v62/services/rules"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePingAccessRule() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePingAccessRuleCreate,
		ReadContext:   resourcePingAccessRuleRead,
		UpdateContext: resourcePingAccessRuleUpdate,
		DeleteContext: resourcePingAccessRuleDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: resourcePingAccessRuleSchema(),
		CustomizeDiff: func(ctx context.Context, d *schema.ResourceDiff, m interface{}) error {
			svc := m.(paClient).Rules
			desc, _, err := svc.GetRuleDescriptorsCommand()
			if err != nil {
				return fmt.Errorf("unable to retrieve Rule descriptors %s", err)
			}
			className := d.Get("class_name").(string)
			if err := ruleDescriptorsHasClassName(className, desc); err != nil {
				return err
			}
			return validateRulesConfiguration(className, d, desc)
		},
	}
}

func resourcePingAccessRuleSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"class_name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"supported_destinations": setOfString(),
		"configuration": {
			Type:             schema.TypeString,
			Required:         true,
			DiffSuppressFunc: suppressEquivalentJSONDiffs,
		},
	}
}

func resourcePingAccessRuleCreate(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(paClient).Rules
	input := rules.AddRuleCommandInput{
		Body: *resourcePingAccessRuleReadData(d),
	}

	result, _, err := svc.AddRuleCommand(&input)
	if err != nil {
		return diag.Errorf("unable to create Rule: %s", err)
	}

	d.SetId(result.Id.String())
	return resourcePingAccessRuleReadResult(d, result, svc)
}

func resourcePingAccessRuleRead(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(paClient).Rules
	input := &rules.GetRuleCommandInput{
		Id: d.Id(),
	}

	result, _, err := svc.GetRuleCommand(input)
	if err != nil {
		return diag.Errorf("unable to read Rule: %s", err)
	}

	return resourcePingAccessRuleReadResult(d, result, svc)
}

func resourcePingAccessRuleUpdate(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(paClient).Rules
	input := rules.UpdateRuleCommandInput{
		Body: *resourcePingAccessRuleReadData(d),
		Id:   d.Id(),
	}

	result, _, err := svc.UpdateRuleCommand(&input)
	if err != nil {
		return diag.Errorf("unable to update Rule: %s", err)
	}

	d.SetId(result.Id.String())
	return resourcePingAccessRuleReadResult(d, result, svc)
}

func resourcePingAccessRuleDelete(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(paClient).Rules
	_, err := svc.DeleteRuleCommand(&rules.DeleteRuleCommandInput{Id: d.Id()})
	if err != nil {
		return diag.Errorf("unable to delete Rule: %s", err)
	}
	return nil
}

func resourcePingAccessRuleReadResult(d *schema.ResourceData, input *models.RuleView, svc rules.RulesAPI) diag.Diagnostics {
	var diags diag.Diagnostics
	b, _ := json.Marshal(input.Configuration)
	config := string(b)

	originalConfig := d.Get("configuration").(string)

	//Search the Rules descriptors for CONCEALED fields, and update the original value back as we cannot use the
	//encryptedValue provided by the API, whilst this gives us a stable plan - we cannot determine if a CONCEALED value
	//has changed and needs updating
	desc, _, _ := svc.GetRuleDescriptorsCommand()
	config = maskConfigFromRuleDescriptors(desc, input.ClassName, originalConfig, config)

	setResourceDataStringWithDiagnostic(d, "name", input.Name, &diags)
	setResourceDataStringWithDiagnostic(d, "class_name", input.ClassName, &diags)
	setResourceDataStringWithDiagnostic(d, "configuration", &config, &diags)
	if err := d.Set("supported_destinations", input.SupportedDestinations); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	return diags
}

func resourcePingAccessRuleReadData(d *schema.ResourceData) *models.RuleView {
	config := d.Get("configuration").(string)
	var dat map[string]interface{}
	_ = json.Unmarshal([]byte(config), &dat)
	supDests := expandStringList(d.Get("supported_destinations").(*schema.Set).List())
	rule := &models.RuleView{
		Name:                  String(d.Get("name").(string)),
		ClassName:             String(d.Get("class_name").(string)),
		Configuration:         dat,
		SupportedDestinations: &supDests,
	}
	return rule
}
