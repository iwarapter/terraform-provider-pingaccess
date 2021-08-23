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
			svc := m.(paClient).HighAvailability
			desc, _, err := svc.GetAvailabilityProfileDescriptorsCommand()
			if err != nil {
				return fmt.Errorf("unable to retrieve AvailabilityProfile descriptors %s", err)
			}
			className := d.Get("class_name").(string)
			if err := descriptorsHasClassName(className, desc); err != nil {
				return err
			}
			return validateConfiguration(className, d, desc)
		},
	}
}

func resourcePingAccessAvailabilityProfileSchema() map[string]*schema.Schema {
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
	return resourcePingAccessAvailabilityProfileReadResult(d, result, svc)
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

	return resourcePingAccessAvailabilityProfileReadResult(d, result, svc)
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
	return resourcePingAccessAvailabilityProfileReadResult(d, result, svc)
}

func resourcePingAccessAvailabilityProfileDelete(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(paClient).HighAvailability
	_, err := svc.DeleteAvailabilityProfileCommand(&highAvailability.DeleteAvailabilityProfileCommandInput{Id: d.Id()})
	if err != nil {
		return diag.Errorf("unable to delete AvailabilityProfile: %s", err)
	}
	return nil
}

func resourcePingAccessAvailabilityProfileReadResult(d *schema.ResourceData, input *models.AvailabilityProfileView, svc highAvailability.HighAvailabilityAPI) diag.Diagnostics {
	var diags diag.Diagnostics
	b, _ := json.Marshal(input.Configuration)
	config := string(b)

	originalConfig := d.Get("configuration").(string)

	//Search the Availability Profiles descriptors for CONCEALED fields, and update the original value back as we cannot use the
	//encryptedValue provided by the API, whilst this gives us a stable plan - we cannot determine if a CONCEALED value
	//has changed and needs updating
	desc, _, _ := svc.GetAvailabilityProfileDescriptorsCommand()
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
