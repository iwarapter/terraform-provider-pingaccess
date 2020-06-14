package pingaccess

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pa "github.com/iwarapter/pingaccess-sdk-go/pingaccess"
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

func dataSourcePingAccessAcmeDefaultRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(*pa.Client).Acme
	result, resp, err := svc.GetDefaultAcmeServerCommand()
	if err != nil {
		return diag.FromErr(fmt.Errorf("unable to read ACME Default: %s\n%v", err.Error(), resp))

	}
	d.SetId(*result.Id)
	var diags diag.Diagnostics
	setResourceDataStringWithDiagnostic(d, "location", result.Location, &diags)
	return diags
}
