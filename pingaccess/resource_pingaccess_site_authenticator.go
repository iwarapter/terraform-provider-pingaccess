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

func resourcePingAccessSiteAuthenticator() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePingAccessSiteAuthenticatorCreate,
		ReadContext:   resourcePingAccessSiteAuthenticatorRead,
		UpdateContext: resourcePingAccessSiteAuthenticatorUpdate,
		DeleteContext: resourcePingAccessSiteAuthenticatorDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: resourcePingAccessSiteAuthenticatorSchema(),
	}
}

func resourcePingAccessSiteAuthenticatorSchema() map[string]*schema.Schema {
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

func resourcePingAccessSiteAuthenticatorCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(*pingaccess.Client).SiteAuthenticators
	input := pingaccess.AddSiteAuthenticatorCommandInput{
		Body: *resourcePingAccessSiteAuthenticatorReadData(d),
	}

	result, _, err := svc.AddSiteAuthenticatorCommand(&input)
	if err != nil {
		return diag.Diagnostics{diag.FromErr(fmt.Errorf("unable to create SiteAuthenticator: %s", err))}
	}

	d.SetId(result.Id.String())
	return resourcePingAccessSiteAuthenticatorReadResult(d, result, svc)
}

func resourcePingAccessSiteAuthenticatorRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(*pingaccess.Client).SiteAuthenticators
	input := &pingaccess.GetSiteAuthenticatorCommandInput{
		Id: d.Id(),
	}
	result, _, err := svc.GetSiteAuthenticatorCommand(input)
	if err != nil {
		return diag.Diagnostics{diag.FromErr(fmt.Errorf("unable to read SiteAuthenticator: %s", err))}
	}
	return resourcePingAccessSiteAuthenticatorReadResult(d, result, svc)
}

func resourcePingAccessSiteAuthenticatorUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(*pingaccess.Client).SiteAuthenticators
	input := pingaccess.UpdateSiteAuthenticatorCommandInput{
		Body: *resourcePingAccessSiteAuthenticatorReadData(d),
		Id:   d.Id(),
	}
	result, _, err := svc.UpdateSiteAuthenticatorCommand(&input)
	if err != nil {
		return diag.Diagnostics{diag.FromErr(fmt.Errorf("unable to update SiteAuthenticator: %s", err))}
	}
	return resourcePingAccessSiteAuthenticatorReadResult(d, result, svc)
}

func resourcePingAccessSiteAuthenticatorDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(*pingaccess.Client).SiteAuthenticators
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	input := &pingaccess.DeleteSiteAuthenticatorCommandInput{
		Id: d.Id(),
	}

	_, err := svc.DeleteSiteAuthenticatorCommand(input)
	if err != nil {
		return diag.Diagnostics{diag.FromErr(fmt.Errorf("unable to delete SiteAuthenticator: %s", err))}
	}
	return nil
}

func resourcePingAccessSiteAuthenticatorReadResult(d *schema.ResourceData, input *pingaccess.SiteAuthenticatorView, svc *pingaccess.SiteAuthenticatorsService) diag.Diagnostics {
	var diags diag.Diagnostics
	b, _ := json.Marshal(input.Configuration)
	config := string(b)

	originalConfig := d.Get("configuration").(string)

	//Search the Site Authenticator descriptors for CONCEALED fields, and update the original value back as we cannot use the
	//encryptedValue provided by the API, whilst this gives us a stable plan - we cannot determine if a CONCEALED value
	//has changed and needs updating
	desc, _, _ := svc.GetSiteAuthenticatorDescriptorsCommand()
	config = maskConfigFromDescriptors(desc, input.ClassName, originalConfig, config)

	setResourceDataStringWithDiagnostic(d, "name", input.Name, &diags)
	setResourceDataStringWithDiagnostic(d, "class_name", input.ClassName, &diags)
	setResourceDataStringWithDiagnostic(d, "configuration", &config, &diags)
	return diags
}

func resourcePingAccessSiteAuthenticatorReadData(d *schema.ResourceData) *pingaccess.SiteAuthenticatorView {
	config := d.Get("configuration").(string)
	var dat map[string]interface{}
	_ = json.Unmarshal([]byte(config), &dat)
	siteAuthenticator := &pingaccess.SiteAuthenticatorView{
		Name:          String(d.Get("name").(string)),
		ClassName:     String(d.Get("class_name").(string)),
		Configuration: dat,
	}
	return siteAuthenticator
}
