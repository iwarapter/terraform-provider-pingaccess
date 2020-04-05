package pingaccess

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	pa "github.com/iwarapter/pingaccess-sdk-go/pingaccess"
)

func dataSourcePingAccessTrustedCertificateGroups() *schema.Resource {
	return &schema.Resource{
		Read:   dataSourcePingAccessTrustedCertificateGroupsRead,
		Schema: dataSourcePingAccessTrustedCertificateGroupsSchema(),
	}
}

func dataSourcePingAccessTrustedCertificateGroupsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"cert_ids": &schema.Schema{
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"ignore_all_certificate_errors": &schema.Schema{
			Type:     schema.TypeBool,
			Computed: true,
		},
		"name": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"skip_certificate_date_check": &schema.Schema{
			Type:     schema.TypeBool,
			Computed: true,
		},
		"system_group": &schema.Schema{
			Type:     schema.TypeBool,
			Computed: true,
		},
		"use_java_trust_store": &schema.Schema{
			Type:     schema.TypeBool,
			Computed: true,
		},
	}
}

func dataSourcePingAccessTrustedCertificateGroupsRead(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pa.Client).TrustedCertificateGroups
	input := &pa.GetTrustedCertificateGroupsCommandInput{
		Name: d.Get("name").(string),
	}
	result, _, err := svc.GetTrustedCertificateGroupsCommand(input)
	if err != nil {
		return fmt.Errorf("Error reading TrustedCertificateGroups: %s", err)
	}
	if len(result.Items) != 1 {
		return fmt.Errorf("Unable to find trusted certificate group with the name %s: %s", d.Get("name").(string), err)
	}
	d.SetId(result.Items[0].Id.String())
	return resourcePingAccessTrustedCertificateGroupsReadResult(d, result.Items[0])
}
