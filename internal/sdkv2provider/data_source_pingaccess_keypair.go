package sdkv2provider

import (
	"context"
	"strconv"

	"github.com/iwarapter/pingaccess-sdk-go/v62/services/keyPairs"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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
	sch["hsm_provider_id"] = &schema.Schema{
		Type:        schema.TypeInt,
		Computed:    true,
		Description: "The HSM Provider ID. The default value is 0 indicating an HSM is not used for this key pair.",
	}
	return &schema.Resource{
		ReadContext: dataSourcePingAccessKeyPairRead,
		Schema:      sch,
		Description: "Use this data source to get keypair information in the PingAccess instance.",
	}
}

func dataSourcePingAccessKeyPairRead(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(paClient).KeyPairs
	input := &keyPairs.GetKeyPairsCommandInput{
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
