package pingaccess

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pa "github.com/iwarapter/pingaccess-sdk-go/pingaccess"
)

func dataSourcePingAccessKeyPair() *schema.Resource {
	sch := resourcePingAccessKeyPairSchema()
	//The normal certificate schema has a file data passed to it, this isnt required for the data resource
	delete(sch, "file_data")
	delete(sch, "password")
	delete(sch, "city")
	delete(sch, "common_name")
	delete(sch, "country")
	delete(sch, "key_algorithm")
	delete(sch, "key_size")
	delete(sch, "organization")
	delete(sch, "organization_unit")
	delete(sch, "state")
	delete(sch, "valid_days")
	return &schema.Resource{
		ReadContext: dataSourcePingAccessKeyPairRead,
		Schema:      sch,
	}
}

func dataSourcePingAccessKeyPairRead(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(*pa.Client).KeyPairs
	input := &pa.GetKeyPairsCommandInput{
		Alias: d.Get("alias").(string),
	}
	result, _, err := svc.GetKeyPairsCommand(input)
	if err != nil {
		return diag.Errorf("unable to read KeyPair: %s", err)
	}
	if len(result.Items) != 1 {
		return diag.Errorf("unable to find KeyPair with alias '%s' found '%d' results", d.Get("alias").(string), len(result.Items))
	}
	d.SetId(strconv.Itoa(*result.Items[0].Id))
	return resourcePingAccessKeyPairReadResult(d, result.Items[0])
}
