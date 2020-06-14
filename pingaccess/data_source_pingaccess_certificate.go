package pingaccess

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pa "github.com/iwarapter/pingaccess-sdk-go/pingaccess"
)

func dataSourcePingAccessCertificate() *schema.Resource {
	sch := resourcePingAccessCertificateSchema()
	//The normal certificate schema has a file data passed to it, this isnt required for the data resource
	delete(sch, "file_data")
	return &schema.Resource{
		ReadContext: dataSourcePingAccessCertificateRead,
		Schema:      sch,
	}
}

func dataSourcePingAccessCertificateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(*pa.Client).Certificates
	input := &pa.GetTrustedCertsInput{
		Alias: d.Get("alias").(string),
	}
	result, _, err := svc.GetTrustedCerts(input)
	if err != nil {
		return diag.FromErr(fmt.Errorf("unable to read Certificate: %s", err))

	}
	if len(result.Items) != 1 {
		return diag.FromErr(fmt.Errorf("unable to find Certificate with alias '%s' found '%d' results", d.Get("alias").(string), len(result.Items)))
	}
	d.SetId(strconv.Itoa(*result.Items[0].Id))
	return resourcePingAccessCertificateReadResult(d, result.Items[0])
}
