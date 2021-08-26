package sdkv2provider

import (
	"context"
	"strconv"

	"github.com/iwarapter/pingaccess-sdk-go/v62/services/certificates"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourcePingAccessCertificate() *schema.Resource {
	sch := resourcePingAccessCertificateSchema()
	//The normal certificate schema has a file data passed to it, this isnt required for the data resource
	delete(sch, "file_data")
	return &schema.Resource{
		ReadContext: dataSourcePingAccessCertificateRead,
		Schema:      sch,
		Description: "Use this data source to get certificate information in the PingAccess instance.",
	}
}

func dataSourcePingAccessCertificateRead(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(paClient).Certificates
	input := &certificates.GetTrustedCertsInput{
		Alias: d.Get("alias").(string),
	}
	result, _, err := svc.GetTrustedCerts(input)
	if err != nil {
		return diag.Errorf("unable to read Certificate: %s", err)

	}
	if len(result.Items) != 1 {
		return diag.Errorf("unable to find Certificate with alias '%s' found '%d' results", d.Get("alias").(string), len(result.Items))
	}
	d.SetId(strconv.Itoa(*result.Items[0].Id))
	return resourcePingAccessCertificateReadResult(d, result.Items[0])
}
