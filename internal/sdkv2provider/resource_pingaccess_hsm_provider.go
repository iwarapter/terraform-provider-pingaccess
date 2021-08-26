package sdkv2provider

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess/models"
	"github.com/iwarapter/pingaccess-sdk-go/v62/services/hsmProviders"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePingAccessHsmProvider() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePingAccessHsmProviderCreate,
		ReadContext:   resourcePingAccessHsmProviderRead,
		UpdateContext: resourcePingAccessHsmProviderUpdate,
		DeleteContext: resourcePingAccessHsmProviderDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: resourcePingAccessHsmProviderSchema(),
		CustomizeDiff: func(ctx context.Context, d *schema.ResourceDiff, m interface{}) error {
			svc := m.(paClient).HsmProviders
			desc, _, err := svc.GetHsmProviderDescriptorsCommand()
			if err != nil {
				return fmt.Errorf("unable to retrieve HsmProvider descriptors %s", err)
			}
			className := d.Get("class_name").(string)
			if err := descriptorsHasClassName(className, desc); err != nil {
				return err
			}
			return validateConfiguration(className, d, desc)
		},
		Description: `Provides configuration for HSM Providers within PingAccess.

-> The PingAccess API does not provider repeatable means of querying a sensitive value, we are unable to detect configuration drift of any sensitive fields in the configuration block.`,
	}
}

func resourcePingAccessHsmProviderSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"class_name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The HSM provider's class name.",
		},
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The HSM provider's name.",
		},
		"configuration": {
			Type:             schema.TypeString,
			Required:         true,
			DiffSuppressFunc: suppressEquivalentJSONDiffs,
			Description:      "The HSM provider's configuration data.",
		},
	}
}

func resourcePingAccessHsmProviderCreate(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(paClient).HsmProviders
	input := hsmProviders.AddHsmProviderCommandInput{
		Body: *resourcePingAccessHsmProviderReadData(d),
	}

	result, _, err := svc.AddHsmProviderCommand(&input)
	if err != nil {
		return diag.Errorf("unable to create HsmProvider: %s", err)
	}

	d.SetId(result.Id.String())
	return resourcePingAccessHsmProviderReadResult(d, result, svc)
}

func resourcePingAccessHsmProviderRead(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(paClient).HsmProviders
	input := &hsmProviders.GetHsmProviderCommandInput{
		Id: d.Id(),
	}
	result, _, err := svc.GetHsmProviderCommand(input)
	if err != nil {
		return diag.Errorf("unable to read HsmProvider: %s", err)
	}
	return resourcePingAccessHsmProviderReadResult(d, result, svc)
}

func resourcePingAccessHsmProviderUpdate(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(paClient).HsmProviders
	input := hsmProviders.UpdateHsmProviderCommandInput{
		Body: *resourcePingAccessHsmProviderReadData(d),
		Id:   d.Id(),
	}
	result, _, err := svc.UpdateHsmProviderCommand(&input)
	if err != nil {
		return diag.Errorf("unable to update HsmProvider: %s", err)
	}
	d.SetId(result.Id.String())
	return resourcePingAccessHsmProviderReadResult(d, result, svc)
}

func resourcePingAccessHsmProviderDelete(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(paClient).HsmProviders
	_, err := svc.DeleteHsmProviderCommand(&hsmProviders.DeleteHsmProviderCommandInput{Id: d.Id()})
	if err != nil {
		return diag.Errorf("unable to delete HsmProvider: %s", err)
	}
	return nil
}

func resourcePingAccessHsmProviderReadResult(d *schema.ResourceData, input *models.HsmProviderView, svc hsmProviders.HsmProvidersAPI) diag.Diagnostics {
	var diags diag.Diagnostics
	b, _ := json.Marshal(input.Configuration)
	config := string(b)

	originalConfig := d.Get("configuration").(string)

	//Search the HSM descriptors for CONCEALED fields, and update the original value back as we cannot use the
	// encryptedValue provided by the API, whilst this gives us a stable plan - we cannot determine if a CONCEALED value
	// has changed and needs updating
	desc, _, _ := svc.GetHsmProviderDescriptorsCommand()

	config = maskConfigFromDescriptors(desc, input.ClassName, originalConfig, config)

	setResourceDataStringWithDiagnostic(d, "name", input.Name, &diags)
	setResourceDataStringWithDiagnostic(d, "class_name", input.ClassName, &diags)
	setResourceDataStringWithDiagnostic(d, "configuration", &config, &diags)
	return diags
}

func resourcePingAccessHsmProviderReadData(d *schema.ResourceData) *models.HsmProviderView {
	config := d.Get("configuration").(string)
	var dat map[string]interface{}
	_ = json.Unmarshal([]byte(config), &dat)
	hsmProvider := &models.HsmProviderView{
		Name:          String(d.Get("name").(string)),
		ClassName:     String(d.Get("class_name").(string)),
		Configuration: dat,
	}
	return hsmProvider
}
