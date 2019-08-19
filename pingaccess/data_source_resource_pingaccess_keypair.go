package pingaccess

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	pa "github.com/iwarapter/pingaccess-sdk-go/pingaccess"
)

func dataSourcePingAccessKeyPair() *schema.Resource {
	sch := resourcePingAccessKeyPairSchema()
	//The normal certificate schema has a file data passed to it, this isnt required for the data resource
	delete(sch, "file_data")
	delete(sch, "password")
	return &schema.Resource{
		Read:   dataSourcePingAccessKeyPairRead,
		Schema: sch,
	}
}

func dataSourcePingAccessKeyPairRead(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pa.Client).KeyPairs
	input := &pa.GetKeyPairsCommandInput{
		Alias: d.Get("alias").(string),
	}
	result, _, err := svc.GetKeyPairsCommand(input)
	if err != nil {
		return fmt.Errorf("Error reading KeyPair: %s", err)
	}
	if len(result.Items) != 1 {
		return fmt.Errorf("Unable to find keypair with alias %s: %s", d.Get("alias").(string), err)
	}
	d.SetId(result.Items[0].Id.String())
	return resourcePingAccessKeyPairReadResult(d, result.Items[0])
}
