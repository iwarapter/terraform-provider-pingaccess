package pingaccess

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pa "github.com/iwarapter/pingaccess-sdk-go/pingaccess"
)

func dataSourcePingAccessAcmeDefault() *schema.Resource {
	return &schema.Resource{
		Read:   dataSourcePingAccessAcmeDefaultRead,
		Schema: dataSourcePingAccessAcmeDefaultSchema(),
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

func dataSourcePingAccessAcmeDefaultRead(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pa.Client).Acme
	result, resp, err := svc.GetDefaultAcmeServerCommand()
	if err != nil {
		return fmt.Errorf("Error reading ACME Default: %s\n%v", err.Error(), resp)
	}
	d.SetId(*result.Id)
	return setResourceDataString(d, "location", result.Location)
}
