package pingaccess

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	pa "github.com/iwarapter/pingaccess-sdk-go/pingaccess"
)

func dataSourcePingAccessCertificate() *schema.Resource {
	sch := resourcePingAccessCertificateSchema()
	//The normal certificate schema has a file data passed to it, this isnt required for the data resource
	delete(sch, "file_data")
	return &schema.Resource{
		Read:   dataSourcePingAccessCertificateRead,
		Schema: sch,
	}
}

func dataSourcePingAccessCertificateRead(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pa.Client).Certificates
	input := &pa.GetTrustedCertsInput{
		Alias: d.Get("alias").(string),
	}
	result, _, err := svc.GetTrustedCerts(input)
	if err != nil {
		return fmt.Errorf("Error reading Certificate: %s", err)
	}
	if len(result.Items) != 1 {
		return fmt.Errorf("Unable to find certificate with alias %s: %s", d.Get("alias").(string), err)
	}
	d.SetId(result.Items[0].Id.String())
	return resourcePingAccessCertificateReadResult(d, result.Items[0])
}
