package sdkv2provider

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/iwarapter/pingaccess-sdk-go/pingaccess/models"
	"github.com/iwarapter/pingaccess-sdk-go/services/accessTokenValidators"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePingAccessAccessTokenValidator() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePingAccessAccessTokenValidatorCreate,
		ReadContext:   resourcePingAccessAccessTokenValidatorRead,
		UpdateContext: resourcePingAccessAccessTokenValidatorUpdate,
		DeleteContext: resourcePingAccessAccessTokenValidatorDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: resourcePingAccessAccessTokenValidatorSchema(),
		CustomizeDiff: func(ctx context.Context, d *schema.ResourceDiff, m interface{}) error {
			svc := m.(paClient).AccessTokenValidators
			desc, _, err := svc.GetAccessTokenValidatorDescriptorsCommand()
			if err != nil {
				return fmt.Errorf("unable to retrieve AccessTokenValidator descriptors %s", err)
			}
			className := d.Get("class_name").(string)
			if err := descriptorsHasClassName(className, desc); err != nil {
				return err
			}
			return validateConfiguration(className, d, desc)
		},
	}
}

func resourcePingAccessAccessTokenValidatorSchema() map[string]*schema.Schema {
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

func resourcePingAccessAccessTokenValidatorCreate(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(paClient).AccessTokenValidators
	input := accessTokenValidators.AddAccessTokenValidatorCommandInput{
		Body: *resourcePingAccessAccessTokenValidatorReadData(d),
	}

	result, _, err := svc.AddAccessTokenValidatorCommand(&input)
	if err != nil {
		return diag.Errorf("unable to create AccessTokenValidator: %s", err)
	}

	d.SetId(result.Id.String())
	return resourcePingAccessAccessTokenValidatorReadResult(d, result, svc)
}

func resourcePingAccessAccessTokenValidatorRead(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(paClient).AccessTokenValidators

	input := &accessTokenValidators.GetAccessTokenValidatorCommandInput{
		Id: d.Id(),
	}

	result, _, err := svc.GetAccessTokenValidatorCommand(input)
	if err != nil {
		return diag.Errorf("unable to read AccessTokenValidator: %s", err)
	}

	return resourcePingAccessAccessTokenValidatorReadResult(d, result, svc)
}

func resourcePingAccessAccessTokenValidatorUpdate(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(paClient).AccessTokenValidators
	input := accessTokenValidators.UpdateAccessTokenValidatorCommandInput{
		Body: *resourcePingAccessAccessTokenValidatorReadData(d),
		Id:   d.Id(),
	}

	result, _, err := svc.UpdateAccessTokenValidatorCommand(&input)
	if err != nil {
		return diag.Errorf("unable to update AccessTokenValidator: %s", err)
	}

	d.SetId(result.Id.String())
	return resourcePingAccessAccessTokenValidatorReadResult(d, result, svc)
}

func resourcePingAccessAccessTokenValidatorDelete(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(paClient).AccessTokenValidators
	_, err := svc.DeleteAccessTokenValidatorCommand(&accessTokenValidators.DeleteAccessTokenValidatorCommandInput{Id: d.Id()})
	if err != nil {
		return diag.Errorf("unable to delete AccessTokenValidator: %s", err)
	}
	return nil
}

func resourcePingAccessAccessTokenValidatorReadResult(d *schema.ResourceData, input *models.AccessTokenValidatorView, svc accessTokenValidators.AccessTokenValidatorsAPI) diag.Diagnostics {
	var diags diag.Diagnostics
	b, _ := json.Marshal(input.Configuration)
	config := string(b)

	originalConfig := d.Get("configuration").(string)

	//Search the Access Token Validators descriptors for CONCEALED fields, and update the original value back as we cannot use the
	//encryptedValue provided by the API, whilst this gives us a stable plan - we cannot determine if a CONCEALED value
	//has changed and needs updating
	desc, _, _ := svc.GetAccessTokenValidatorDescriptorsCommand()
	config = maskConfigFromDescriptors(desc, input.ClassName, originalConfig, config)

	setResourceDataStringWithDiagnostic(d, "name", input.Name, &diags)
	setResourceDataStringWithDiagnostic(d, "class_name", input.ClassName, &diags)
	setResourceDataStringWithDiagnostic(d, "configuration", &config, &diags)
	return diags
}

func resourcePingAccessAccessTokenValidatorReadData(d *schema.ResourceData) *models.AccessTokenValidatorView {
	config := d.Get("configuration").(string)
	var dat map[string]interface{}
	_ = json.Unmarshal([]byte(config), &dat)
	atv := &models.AccessTokenValidatorView{
		Name:          String(d.Get("name").(string)),
		ClassName:     String(d.Get("class_name").(string)),
		Configuration: dat,
	}
	return atv
}
