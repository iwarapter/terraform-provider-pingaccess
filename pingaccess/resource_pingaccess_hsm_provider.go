package pingaccess

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/iwarapter/pingaccess-sdk-go/pingaccess"
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
			svc := m.(*pingaccess.Client).HsmProviders
			desc, _, _ := svc.GetHsmProviderDescriptorsCommand()
			className := d.Get("class_name").(string)
			if err := descriptorsHasClassName(className, desc); err != nil {
				return err
			}
			return validateConfiguration(className, d, desc)
		},
	}
}

func resourcePingAccessHsmProviderSchema() map[string]*schema.Schema {
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
			Type:     schema.TypeString,
			Required: true,
		},
	}
}

func resourcePingAccessHsmProviderCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(*pingaccess.Client).HsmProviders
	input := pingaccess.AddHsmProviderCommandInput{
		Body: *resourcePingAccessHsmProviderReadData(d),
	}

	result, _, err := svc.AddHsmProviderCommand(&input)
	if err != nil {
		return diag.FromErr(fmt.Errorf("unable to create HsmProvider: %s", err))
	}

	d.SetId(result.Id.String())
	return resourcePingAccessHsmProviderReadResult(d, result, svc)
}

func resourcePingAccessHsmProviderRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(*pingaccess.Client).HsmProviders
	input := &pingaccess.GetHsmProviderCommandInput{
		Id: d.Id(),
	}
	result, _, err := svc.GetHsmProviderCommand(input)
	if err != nil {
		return diag.FromErr(fmt.Errorf("unable to read HsmProvider: %s", err))
	}
	return resourcePingAccessHsmProviderReadResult(d, result, svc)
}

func resourcePingAccessHsmProviderUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(*pingaccess.Client).HsmProviders
	input := pingaccess.UpdateHsmProviderCommandInput{
		Body: *resourcePingAccessHsmProviderReadData(d),
		Id:   d.Id(),
	}
	result, _, err := svc.UpdateHsmProviderCommand(&input)
	if err != nil {
		return diag.FromErr(fmt.Errorf("unable to update HsmProvider: %s", err))
	}
	d.SetId(result.Id.String())
	return resourcePingAccessHsmProviderReadResult(d, result, svc)
}

func resourcePingAccessHsmProviderDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(*pingaccess.Client).HsmProviders
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	_, err := svc.DeleteHsmProviderCommand(&pingaccess.DeleteHsmProviderCommandInput{Id: d.Id()})
	if err != nil {
		return diag.FromErr(fmt.Errorf("unable to delete HsmProvider: %s", err))
	}
	return nil
}

func resourcePingAccessHsmProviderReadResult(d *schema.ResourceData, input *pingaccess.HsmProviderView, svc pingaccess.HsmProvidersAPI) diag.Diagnostics {
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

func resourcePingAccessHsmProviderReadData(d *schema.ResourceData) *pingaccess.HsmProviderView {
	config := d.Get("configuration").(string)
	var dat map[string]interface{}
	_ = json.Unmarshal([]byte(config), &dat)
	hsmProvider := &pingaccess.HsmProviderView{
		Name:          String(d.Get("name").(string)),
		ClassName:     String(d.Get("class_name").(string)),
		Configuration: dat,
	}
	return hsmProvider
}
