package sdkv2provider

import (
	"context"

	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess/models"
	"github.com/iwarapter/pingaccess-sdk-go/v62/services/acme"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePingAccessAcmeServer() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePingAccessAcmeServerCreate,
		ReadContext:   resourcePingAccessAcmeServerRead,
		DeleteContext: resourcePingAccessAcmeServerDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema:      resourcePingAccessAcmeServerSchema(),
		Description: `Provides configuration for ACME Server within PingAccess.`,
	}
}

func resourcePingAccessAcmeServerSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
			Description: "A user-friendly name for the ACME server.",
		},
		"url": {
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
			Description: "The URL of the ACME directory resource on the ACME server.",
		},
		"acme_accounts": acmeServerAccountsSchema(),
	}
}

func resourcePingAccessAcmeServerCreate(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(paClient).Acme
	input := acme.AddAcmeServerCommandInput{
		Body: *resourcePingAccessAcmeServerReadData(d),
	}

	result, _, err := svc.AddAcmeServerCommand(&input)
	if err != nil {
		return diag.Errorf("unable to create AcmeServer: %s", err)
	}
	d.SetId(*result.Id)
	return resourcePingAccessAcmeServerReadResult(d, &input.Body)
}

func resourcePingAccessAcmeServerRead(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(paClient).Acme
	input := &acme.GetAcmeServerCommandInput{
		AcmeServerId: d.Id(),
	}
	result, _, err := svc.GetAcmeServerCommand(input)
	if err != nil {
		return diag.Errorf("unable to read AcmeServer: %s", err)
	}
	return resourcePingAccessAcmeServerReadResult(d, result)
}

func resourcePingAccessAcmeServerDelete(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(paClient).Acme
	input := &acme.DeleteAcmeServerCommandInput{
		AcmeServerId: d.Id(),
	}

	_, _, err := svc.DeleteAcmeServerCommand(input)
	if err != nil {
		return diag.Errorf("unable to delete AcmeServer: %s", err)
	}
	return nil
}

func resourcePingAccessAcmeServerReadResult(d *schema.ResourceData, input *models.AcmeServerView) diag.Diagnostics {
	var diags diag.Diagnostics
	setResourceDataStringWithDiagnostic(d, "name", input.Name, &diags)
	setResourceDataStringWithDiagnostic(d, "url", input.Url, &diags)
	if input.AcmeAccounts != nil && len(input.AcmeAccounts) > 0 {
		if err := d.Set("acme_accounts", flattenLinkViewList(input.AcmeAccounts)); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}
	return diags
}

func resourcePingAccessAcmeServerReadData(d *schema.ResourceData) *models.AcmeServerView {
	return &models.AcmeServerView{
		Name: String(d.Get("name").(string)),
		Url:  String(d.Get("url").(string)),
	}
}
