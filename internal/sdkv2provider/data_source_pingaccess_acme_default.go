package sdkv2provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourcePingAccessAcmeDefault() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourcePingAccessAcmeDefaultRead,
		Schema:      dataSourcePingAccessAcmeDefaultSchema(),
	}
}

func dataSourcePingAccessAcmeDefaultSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"location": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}

func dataSourcePingAccessAcmeDefaultRead(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(paClient).Acme
	result, _, err := svc.GetDefaultAcmeServerCommand()
	if err != nil {
		return diag.Errorf("unable to read ACME Default: %s", err)

	}
	d.SetId(*result.Id)
	var diags diag.Diagnostics
	setResourceDataStringWithDiagnostic(d, "location", result.Location, &diags)
	return diags
}
