package pingaccess

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/iwarapter/pingaccess-sdk-go/pingaccess"
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
		Schema: resourcePingAccessRuleSetSchema(),
	}
}

func resourcePingAccessRuleSetSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"element_type": {
			Type:             schema.TypeString,
			Required:         true,
			ValidateDiagFunc: validateRuleOrRuleSet,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"policy": {
			Type:     schema.TypeSet,
			Required: true,
			MinItems: 1,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"success_criteria": {
			Type:             schema.TypeString,
			Required:         true,
			ValidateDiagFunc: validateSuccessIfAllSucceedOrSuccessIfAnyOneSucceeds,
		},
	}
}

func resourcePingAccessRuleSetCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// log.Printf("[INFO] resourcePingAccessRuleSetCreate")
	name := d.Get("name").(string)
	elementType := d.Get("element_type").(string)
	policy := expandStringList(d.Get("policy").(*schema.Set).List())
	successCriteria := d.Get("success_criteria").(string)

	//TODO generalise this into a helper function
	var polIds []*int
	for _, i := range policy {
		text, _ := strconv.Atoi(*i)
		polIds = append(polIds, &text)
	}

	input := pingaccess.AddRuleSetCommandInput{
		Body: pingaccess.RuleSetView{
			Name:            String(name),
			ElementType:     String(elementType),
			Policy:          &polIds,
			SuccessCriteria: String(successCriteria),
		},
	}

	svc := m.(*pingaccess.Client).Rulesets

	result, _, err := svc.AddRuleSetCommand(&input)
	if err != nil {
		return diag.Diagnostics{diag.FromErr(fmt.Errorf("unable to create RuleSet: %s", err))}
	}

	d.SetId(result.Id.String())
	return resourcePingAccessRuleSetReadResult(d, result)
}

func resourcePingAccessRuleSetRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(*pingaccess.Client).Rulesets
	result, _, err := svc.GetRuleSetCommand(&pingaccess.GetRuleSetCommandInput{
		Id: d.Id(),
	})
	if err != nil {
		return diag.Diagnostics{diag.FromErr(fmt.Errorf("unable to read RuleSet: %s", err))}
	}
	return resourcePingAccessRuleSetReadResult(d, result)
}

func resourcePingAccessRuleSetUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// log.Printf("[INFO] resourcePingAccessRuleSetUpdate")
	name := d.Get("name").(string)
	elementType := d.Get("element_type").(string)
	policy := expandStringList(d.Get("policy").(*schema.Set).List())
	successCriteria := d.Get("success_criteria").(string)

	//TODO generalise this into a helper function
	var polIds []*int
	for _, i := range policy {
		text, _ := strconv.Atoi(*i)
		polIds = append(polIds, &text)
	}

	input := pingaccess.UpdateRuleSetCommandInput{
		Body: pingaccess.RuleSetView{
			Name:            String(name),
			ElementType:     String(elementType),
			Policy:          &polIds,
			SuccessCriteria: String(successCriteria),
		},
		Id: d.Id(),
	}

	svc := m.(*pingaccess.Client).Rulesets

	result, _, err := svc.UpdateRuleSetCommand(&input)
	if err != nil {
		return diag.Diagnostics{diag.FromErr(fmt.Errorf("unable to update RuleSet: %s", err))}
	}
	d.SetId(result.Id.String())
	return resourcePingAccessRuleSetReadResult(d, result)
}

func resourcePingAccessRuleSetDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(*pingaccess.Client).Rulesets

	_, err := svc.DeleteRuleSetCommand(&pingaccess.DeleteRuleSetCommandInput{Id: d.Id()})
	if err != nil {
		return diag.Diagnostics{diag.FromErr(fmt.Errorf("unable to delete RuleSet: %s", err))}
	}
	return nil
}

func resourcePingAccessRuleSetReadResult(d *schema.ResourceData, input *pingaccess.RuleSetView) diag.Diagnostics {
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
			diags = append(diags, diag.FromErr(err))
		}
	}
	return diags
}
