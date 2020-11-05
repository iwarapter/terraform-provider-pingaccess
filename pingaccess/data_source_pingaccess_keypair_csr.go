package pingaccess

import (
	"context"
	"strings"

	"github.com/iwarapter/pingaccess-sdk-go/services/keyPairs"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourcePingAccessKeyPairCsr() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourcePingAccessKeyPairCsrRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"cert_request_pem": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourcePingAccessKeyPairCsrRead(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(paClient).KeyPairs
	var diags diag.Diagnostics
	input := &keyPairs.GenerateCsrCommandInput{
		Id: d.Get("id").(string),
	}
	result, _, err := svc.GenerateCsrCommand(input)
	if err != nil {
		return diag.Errorf("unable to find KeyPairCSR with id '%s'", d.Get("id").(string))
	}
	d.SetId(d.Get("id").(string))
	*result = strings.ReplaceAll(*result, " NEW ", " ")
	setResourceDataStringWithDiagnostic(d, "cert_request_pem", result, &diags)
	return diags
}
