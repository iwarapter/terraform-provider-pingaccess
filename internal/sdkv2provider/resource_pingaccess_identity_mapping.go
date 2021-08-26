package sdkv2provider

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess/models"
	"github.com/iwarapter/pingaccess-sdk-go/v62/services/identityMappings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePingAccessIdentityMapping() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePingAccessIdentityMappingCreate,
		ReadContext:   resourcePingAccessIdentityMappingRead,
		UpdateContext: resourcePingAccessIdentityMappingUpdate,
		DeleteContext: resourcePingAccessIdentityMappingDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: resourcePingAccessIdentityMappingSchema(),
		CustomizeDiff: func(ctx context.Context, d *schema.ResourceDiff, m interface{}) error {
			svc := m.(paClient).IdentityMappings
			desc, _, err := svc.GetIdentityMappingDescriptorsCommand()
			if err != nil {
				return fmt.Errorf("unable to retrieve IdentityMapping descriptors %s", err)
			}
			className := d.Get("class_name").(string)
			if err := descriptorsHasClassName(className, desc); err != nil {
				return err
			}
			return validateConfiguration(className, d, desc)
		},
		Description: `Provides configuration for Identity Mappings within PingAccess.

-> The PingAccess API does not provider repeatable means of querying a sensitive value, we are unable to detect configuration drift of any sensitive fields in the configuration block.`,
	}
}

func resourcePingAccessIdentityMappingSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"class_name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The identity mapping's class name.",
		},
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The name of the identity mapping.",
		},
		"configuration": {
			Type:             schema.TypeString,
			Required:         true,
			DiffSuppressFunc: suppressEquivalentJSONDiffs,
			Description:      "The identity mapping's configuration data.",
		},
	}
}

func resourcePingAccessIdentityMappingCreate(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(paClient).IdentityMappings
	input := identityMappings.AddIdentityMappingCommandInput{
		Body: *resourcePingAccessIdentityMappingReadData(d),
	}

	result, _, err := svc.AddIdentityMappingCommand(&input)
	if err != nil {
		return diag.Errorf("unable to create IdentityMapping: %s", err)
	}

	d.SetId(result.Id.String())
	return resourcePingAccessIdentityMappingReadResult(d, result, svc)
}

func resourcePingAccessIdentityMappingRead(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(paClient).IdentityMappings
	input := &identityMappings.GetIdentityMappingCommandInput{
		Id: d.Id(),
	}
	result, _, err := svc.GetIdentityMappingCommand(input)
	if err != nil {

		return diag.Errorf("unable to read IdentityMapping: %s", err)
	}
	return resourcePingAccessIdentityMappingReadResult(d, result, svc)
}

func resourcePingAccessIdentityMappingUpdate(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(paClient).IdentityMappings
	input := identityMappings.UpdateIdentityMappingCommandInput{
		Body: *resourcePingAccessIdentityMappingReadData(d),
		Id:   d.Id(),
	}

	result, _, err := svc.UpdateIdentityMappingCommand(&input)
	if err != nil {
		return diag.Errorf("unable to update IdentityMapping: %s", err)
	}

	d.SetId(result.Id.String())
	return resourcePingAccessIdentityMappingReadResult(d, result, svc)
}

func resourcePingAccessIdentityMappingDelete(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(paClient).IdentityMappings
	log.Printf("[INFO] ResourceID: %s", d.Id())
	_, err := svc.DeleteIdentityMappingCommand(&identityMappings.DeleteIdentityMappingCommandInput{Id: d.Id()})
	if err != nil {
		return diag.Errorf("unable to delete IdentityMapping: %s", err)
	}
	return nil
}

func resourcePingAccessIdentityMappingReadResult(d *schema.ResourceData, input *models.IdentityMappingView, svc identityMappings.IdentityMappingsAPI) diag.Diagnostics {
	var diags diag.Diagnostics
	b, _ := json.Marshal(input.Configuration)
	config := string(b)

	originalConfig := d.Get("configuration").(string)

	//Search the Identity Mappings descriptors for CONCEALED fields, and update the original value back as we cannot use the
	//encryptedValue provided by the API, whilst this gives us a stable plan - we cannot determine if a CONCEALED value
	//has changed and needs updating
	desc, _, _ := svc.GetIdentityMappingDescriptorsCommand()
	config = maskConfigFromDescriptors(desc, input.ClassName, originalConfig, config)

	setResourceDataStringWithDiagnostic(d, "name", input.Name, &diags)
	setResourceDataStringWithDiagnostic(d, "class_name", input.ClassName, &diags)
	setResourceDataStringWithDiagnostic(d, "configuration", &config, &diags)
	return diags
}

func resourcePingAccessIdentityMappingReadData(d *schema.ResourceData) *models.IdentityMappingView {
	config := d.Get("configuration").(string)
	var dat map[string]interface{}
	_ = json.Unmarshal([]byte(config), &dat)
	idMapping := &models.IdentityMappingView{
		Name:          String(d.Get("name").(string)),
		ClassName:     String(d.Get("class_name").(string)),
		Configuration: dat,
	}
	return idMapping
}
