package pingaccess

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/iwarapter/pingaccess-sdk-go/pingaccess"
)

func resourcePingAccessAcmeServer() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePingAccessAcmeServerCreate,
		ReadContext:   resourcePingAccessAcmeServerRead,
		DeleteContext: resourcePingAccessAcmeServerDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: resourcePingAccessAcmeServerSchema(),
	}
}

func resourcePingAccessAcmeServerSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
		"url": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
		"acme_accounts": acmeServerAccountsSchema(),
	}
}

func resourcePingAccessAcmeServerCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(*pingaccess.Client).Acme
	input := pingaccess.AddAcmeServerCommandInput{
		Body: *resourcePingAccessAcmeServerReadData(d),
	}

	result, _, err := svc.AddAcmeServerCommand(&input)
	if err != nil {
		return diag.FromErr(fmt.Errorf("unable to create AcmeServer: %s", err))
	}
	d.SetId(*result.Id)
	return resourcePingAccessAcmeServerReadResult(d, &input.Body)
}

func resourcePingAccessAcmeServerRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(*pingaccess.Client).Acme
	input := &pingaccess.GetAcmeServerCommandInput{
		AcmeServerId: d.Id(),
	}
	result, _, err := svc.GetAcmeServerCommand(input)
	if err != nil {
		return diag.FromErr(fmt.Errorf("unable to read AcmeServer: %s", err))
	}
	return resourcePingAccessAcmeServerReadResult(d, result)
}

func resourcePingAccessAcmeServerDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(*pingaccess.Client).Acme
	input := &pingaccess.DeleteAcmeServerCommandInput{
		AcmeServerId: d.Id(),
	}

	_, resp, err := svc.DeleteAcmeServerCommand(input)
	if err != nil {
		return diag.FromErr(fmt.Errorf("unable to delete AcmeServer: %s resp: %v", err, *resp))
	}
	return nil
}

func resourcePingAccessAcmeServerReadResult(d *schema.ResourceData, input *pingaccess.AcmeServerView) diag.Diagnostics {
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

func resourcePingAccessAcmeServerReadData(d *schema.ResourceData) *pingaccess.AcmeServerView {
	return &pingaccess.AcmeServerView{
		Name: String(d.Get("name").(string)),
		Url:  String(d.Get("url").(string)),
	}
}
