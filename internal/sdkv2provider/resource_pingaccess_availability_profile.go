package sdkv2provider

import (
	"context"
	"encoding/json"

	"github.com/iwarapter/pingaccess-sdk-go/v62/services/highAvailability"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePingAccessAvailabilityProfile() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePingAccessAvailabilityProfileCreate,
		ReadContext:   resourcePingAccessAvailabilityProfileRead,
		UpdateContext: resourcePingAccessAvailabilityProfileUpdate,
		DeleteContext: resourcePingAccessAvailabilityProfileDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: resourcePingAccessAvailabilityProfileSchema(),
		CustomizeDiff: func(ctx context.Context, d *schema.ResourceDiff, m interface{}) error {
			className := d.Get("class_name").(string)
			if err := descriptorsHasClassName(className, m.(paClient).HighAvailabilityDescriptors); err != nil {
				return err
			}
			return validateConfiguration(className, d, m.(paClient).HighAvailabilityDescriptors)
		},
		Description: `Provides configuration for Availability Profiles within PingAccess.

-> The PingAccess API does not provider repeatable means of querying a sensitive value, we are unable to detect configuration drift of any sensitive fields in the configuration block.`,
	}
}

func resourcePingAccessAvailabilityProfileSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"class_name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The class name of the availability profile.",
		},
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The name of the availability profile.",
		},
		"configuration": {
			Type:             schema.TypeString,
			Required:         true,
			DiffSuppressFunc: suppressEquivalentJSONDiffs,
			Description:      "The availability profile's configuration data.",
		},
	}
}

func resourcePingAccessAvailabilityProfileCreate(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(paClient).HighAvailability
	input := highAvailability.AddAvailabilityProfileCommandInput{
		Body: *resourcePingAccessAvailabilityProfileReadData(d),
	}

	result, _, err := svc.AddAvailabilityProfileCommand(&input)
	if err != nil {
		return diag.Errorf("unable to create AvailabilityProfile: %s", err)
	}

	d.SetId(result.Id.String())
	return resourcePingAccessAvailabilityProfileReadResult(d, result, m.(paClient).HighAvailabilityDescriptors)
}

func resourcePingAccessAvailabilityProfileRead(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(paClient).HighAvailability

	input := &highAvailability.GetAvailabilityProfileCommandInput{
		Id: d.Id(),
	}

	result, _, err := svc.GetAvailabilityProfileCommand(input)
	if err != nil {
		return diag.Errorf("unable to read AvailabilityProfile: %s", err)
	}

	return resourcePingAccessAvailabilityProfileReadResult(d, result, m.(paClient).HighAvailabilityDescriptors)
}

func resourcePingAccessAvailabilityProfileUpdate(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(paClient).HighAvailability
	input := highAvailability.UpdateAvailabilityProfileCommandInput{
		Body: *resourcePingAccessAvailabilityProfileReadData(d),
		Id:   d.Id(),
	}

	result, _, err := svc.UpdateAvailabilityProfileCommand(&input)
	if err != nil {
		return diag.Errorf("unable to update AvailabilityProfile: %s", err)
	}

	d.SetId(result.Id.String())
	return resourcePingAccessAvailabilityProfileReadResult(d, result, m.(paClient).HighAvailabilityDescriptors)
}

func resourcePingAccessAvailabilityProfileDelete(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(paClient).HighAvailability
	_, err := svc.DeleteAvailabilityProfileCommand(&highAvailability.DeleteAvailabilityProfileCommandInput{Id: d.Id()})
	if err != nil {
		return diag.Errorf("unable to delete AvailabilityProfile: %s", err)
	}
	return nil
}

func resourcePingAccessAvailabilityProfileReadResult(d *schema.ResourceData, input *models.AvailabilityProfileView, desc *models.DescriptorsView) diag.Diagnostics {
	var diags diag.Diagnostics
	b, _ := json.Marshal(input.Configuration)
	config := string(b)

	originalConfig := d.Get("configuration").(string)

	//Search the Availability Profiles descriptors for CONCEALED fields, and update the original value back as we cannot use the
	//encryptedValue provided by the API, whilst this gives us a stable plan - we cannot determine if a CONCEALED value
	//has changed and needs updating
	config = maskConfigFromDescriptors(desc, input.ClassName, originalConfig, config)

	setResourceDataStringWithDiagnostic(d, "name", input.Name, &diags)
	setResourceDataStringWithDiagnostic(d, "class_name", input.ClassName, &diags)
	setResourceDataStringWithDiagnostic(d, "configuration", &config, &diags)
	return diags
}

func resourcePingAccessAvailabilityProfileReadData(d *schema.ResourceData) *models.AvailabilityProfileView {
	config := d.Get("configuration").(string)
	var dat map[string]interface{}
	_ = json.Unmarshal([]byte(config), &dat)
	atv := &models.AvailabilityProfileView{
		Name:          String(d.Get("name").(string)),
		ClassName:     String(d.Get("class_name").(string)),
		Configuration: dat,
	}
	return atv
}
