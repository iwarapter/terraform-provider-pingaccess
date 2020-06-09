package pingaccess

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pa "github.com/iwarapter/pingaccess-sdk-go/pingaccess"
)

func dataSourcePingAccessKeyPair() *schema.Resource {
	sch := resourcePingAccessKeyPairSchema()
	//The normal certificate schema has a file data passed to it, this isnt required for the data resource
	delete(sch, "file_data")
	delete(sch, "password")
	return &schema.Resource{
		ReadContext: dataSourcePingAccessKeyPairRead,
		Schema:      sch,
	}
}

func dataSourcePingAccessKeyPairRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(*pa.Client).KeyPairs
	input := &pa.GetKeyPairsCommandInput{
		Alias: d.Get("alias").(string),
	}
	result, _, err := svc.GetKeyPairsCommand(input)
	if err != nil {
		return diag.Diagnostics{diag.FromErr(fmt.Errorf("unable to read KeyPair: %s", err))}

	}
	if len(result.Items) != 1 {
		return diag.Diagnostics{diag.FromErr(fmt.Errorf("unable to find KeyPair with alias '%s' found '%d' results", d.Get("alias").(string), len(result.Items)))}
	}
	d.SetId(result.Items[0].Id.String())
	return resourcePingAccessKeyPairReadResult(d, result.Items[0])
}
