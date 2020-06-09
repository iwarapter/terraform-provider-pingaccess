package pingaccess

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/iwarapter/pingaccess-sdk-go/pingaccess"
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
			svc := m.(*pingaccess.Client).AccessTokenValidators
			desc, _, _ := svc.GetAccessTokenValidatorDescriptorsCommand()
			className := d.Get("class_name").(string)
			return descriptorsHasClassName(className, desc)
			//if err := descriptorsHasClassName(className, desc); err != nil {
			//	return err
			//}
			//return validateConfiguration(className, d, desc)
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

func resourcePingAccessAccessTokenValidatorCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(*pingaccess.Client).AccessTokenValidators
	input := pingaccess.AddAccessTokenValidatorCommandInput{
		Body: *resourcePingAccessAccessTokenValidatorReadData(d),
	}

	result, _, err := svc.AddAccessTokenValidatorCommand(&input)
	if err != nil {
		return diag.Diagnostics{diag.FromErr(fmt.Errorf("unable to create AccessTokenValidator: %s", err))}
	}

	d.SetId(result.Id.String())
	return resourcePingAccessAccessTokenValidatorReadResult(d, result, svc)
}

func resourcePingAccessAccessTokenValidatorRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(*pingaccess.Client).AccessTokenValidators

	input := &pingaccess.GetAccessTokenValidatorCommandInput{
		Id: d.Id(),
	}

	result, _, err := svc.GetAccessTokenValidatorCommand(input)
	if err != nil {
		return diag.Diagnostics{diag.FromErr(fmt.Errorf("unable to read AccessTokenValidator: %s", err))}
	}

	return resourcePingAccessAccessTokenValidatorReadResult(d, result, svc)
}

func resourcePingAccessAccessTokenValidatorUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(*pingaccess.Client).AccessTokenValidators
	input := pingaccess.UpdateAccessTokenValidatorCommandInput{
		Body: *resourcePingAccessAccessTokenValidatorReadData(d),
		Id:   d.Id(),
	}

	result, _, err := svc.UpdateAccessTokenValidatorCommand(&input)
	if err != nil {
		return diag.Diagnostics{diag.FromErr(fmt.Errorf("unable to update AccessTokenValidator: %s", err))}
	}

	d.SetId(result.Id.String())
	return resourcePingAccessAccessTokenValidatorReadResult(d, result, svc)
}

func resourcePingAccessAccessTokenValidatorDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(*pingaccess.Client).AccessTokenValidators
	_, err := svc.DeleteAccessTokenValidatorCommand(&pingaccess.DeleteAccessTokenValidatorCommandInput{Id: d.Id()})
	if err != nil {
		return diag.Diagnostics{diag.FromErr(fmt.Errorf("unable to delete AccessTokenValidator: %s", err))}
	}
	return nil
}

func resourcePingAccessAccessTokenValidatorReadResult(d *schema.ResourceData, input *pingaccess.AccessTokenValidatorView, svc *pingaccess.AccessTokenValidatorsService) diag.Diagnostics {
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

func resourcePingAccessAccessTokenValidatorReadData(d *schema.ResourceData) *pingaccess.AccessTokenValidatorView {
	config := d.Get("configuration").(string)
	var dat map[string]interface{}
	_ = json.Unmarshal([]byte(config), &dat)
	atv := &pingaccess.AccessTokenValidatorView{
		Name:          String(d.Get("name").(string)),
		ClassName:     String(d.Get("class_name").(string)),
		Configuration: dat,
	}
	return atv
}
