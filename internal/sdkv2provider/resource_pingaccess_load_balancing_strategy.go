package sdkv2provider

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/iwarapter/pingaccess-sdk-go/v62/services/highAvailability"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePingAccessLoadBalancingStrategy() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePingAccessLoadBalancingStrategyCreate,
		ReadContext:   resourcePingAccessLoadBalancingStrategyRead,
		UpdateContext: resourcePingAccessLoadBalancingStrategyUpdate,
		DeleteContext: resourcePingAccessLoadBalancingStrategyDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: resourcePingAccessLoadBalancingStrategySchema(),
		CustomizeDiff: func(ctx context.Context, d *schema.ResourceDiff, m interface{}) error {
			svc := m.(paClient).HighAvailability
			desc, _, err := svc.GetLoadBalancingStrategyDescriptorsCommand()
			if err != nil {
				return fmt.Errorf("unable to retrieve LoadBalancingStrategy descriptors %s", err)
			}
			className := d.Get("class_name").(string)
			if err := descriptorsHasClassName(className, desc); err != nil {
				return err
			}
			return validateConfiguration(className, d, desc)
		},
	}
}

func resourcePingAccessLoadBalancingStrategySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"class_name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"configuration": {
			Type:             schema.TypeString,
			Required:         true,
			DiffSuppressFunc: suppressEquivalentJSONDiffs,
		},
	}
}

func resourcePingAccessLoadBalancingStrategyCreate(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(paClient).HighAvailability
	input := highAvailability.AddLoadBalancingStrategyCommandInput{
		Body: *resourcePingAccessLoadBalancingStrategyReadData(d),
	}

	result, _, err := svc.AddLoadBalancingStrategyCommand(&input)
	if err != nil {
		return diag.Errorf("unable to create LoadBalancingStrategy: %s", err)
	}

	d.SetId(result.Id.String())
	return resourcePingAccessLoadBalancingStrategyReadResult(d, result, svc)
}

func resourcePingAccessLoadBalancingStrategyRead(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(paClient).HighAvailability

	input := &highAvailability.GetLoadBalancingStrategyCommandInput{
		Id: d.Id(),
	}

	result, _, err := svc.GetLoadBalancingStrategyCommand(input)
	if err != nil {
		return diag.Errorf("unable to read LoadBalancingStrategy: %s", err)
	}

	return resourcePingAccessLoadBalancingStrategyReadResult(d, result, svc)
}

func resourcePingAccessLoadBalancingStrategyUpdate(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(paClient).HighAvailability
	input := highAvailability.UpdateLoadBalancingStrategyCommandInput{
		Body: *resourcePingAccessLoadBalancingStrategyReadData(d),
		Id:   d.Id(),
	}

	result, _, err := svc.UpdateLoadBalancingStrategyCommand(&input)
	if err != nil {
		return diag.Errorf("unable to update LoadBalancingStrategy: %s", err)
	}

	d.SetId(result.Id.String())
	return resourcePingAccessLoadBalancingStrategyReadResult(d, result, svc)
}

func resourcePingAccessLoadBalancingStrategyDelete(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(paClient).HighAvailability
	_, err := svc.DeleteLoadBalancingStrategyCommand(&highAvailability.DeleteLoadBalancingStrategyCommandInput{Id: d.Id()})
	if err != nil {
		return diag.Errorf("unable to delete LoadBalancingStrategy: %s", err)
	}
	return nil
}

func resourcePingAccessLoadBalancingStrategyReadResult(d *schema.ResourceData, input *models.LoadBalancingStrategyView, svc highAvailability.HighAvailabilityAPI) diag.Diagnostics {
	var diags diag.Diagnostics
	b, _ := json.Marshal(input.Configuration)
	config := string(b)

	originalConfig := d.Get("configuration").(string)

	//Search the Load Balancing Strategy descriptors for CONCEALED fields, and update the original value back as we cannot use the
	//encryptedValue provided by the API, whilst this gives us a stable plan - we cannot determine if a CONCEALED value
	//has changed and needs updating
	desc, _, _ := svc.GetLoadBalancingStrategyDescriptorsCommand()
	config = maskConfigFromDescriptors(desc, input.ClassName, originalConfig, config)

	setResourceDataStringWithDiagnostic(d, "name", input.Name, &diags)
	setResourceDataStringWithDiagnostic(d, "class_name", input.ClassName, &diags)
	setResourceDataStringWithDiagnostic(d, "configuration", &config, &diags)
	return diags
}

func resourcePingAccessLoadBalancingStrategyReadData(d *schema.ResourceData) *models.LoadBalancingStrategyView {
	config := d.Get("configuration").(string)
	var dat map[string]interface{}
	_ = json.Unmarshal([]byte(config), &dat)
	atv := &models.LoadBalancingStrategyView{
		Name:          String(d.Get("name").(string)),
		ClassName:     String(d.Get("class_name").(string)),
		Configuration: dat,
	}
	return atv
}
