package pingaccess

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	pa "github.com/iwarapter/pingaccess-sdk-go/pingaccess"
)

func dataSourcePingAccessCertificate() *schema.Resource {
	return &schema.Resource{
		Read: dataSourcePingAccessCertificateRead,
		Schema: map[string]*schema.Schema{
			"alias": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"expires": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"issuer_dn": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"md5sum": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"serial_number": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"sha1sum": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"signature_algorithm": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			// "subject_alternative_names": setOfString(),
			"subject_cn": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"subject_dn": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"valid_from": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
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
