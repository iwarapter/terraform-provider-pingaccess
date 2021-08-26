package sdkv2provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourcePingAccessVersion() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourcePingAccessVersionRead,
		Schema:      dataSourcePingAccessVersionSchema(),
		Description: "Use this data source to get the version of the PingAccess instance.",
	}
}

func dataSourcePingAccessVersionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"version": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The version of PingAccess.",
		},
	}
}

func dataSourcePingAccessVersionRead(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(paClient).Version
	result, resp, err := svc.VersionCommand()
	if err != nil {
		return diag.Errorf("unable to read Version: %s\n%v", err, resp)
	}
	var diags diag.Diagnostics
	d.SetId("version")
	setResourceDataStringWithDiagnostic(d, "version", result.Version, &diags)
	return diags
}
