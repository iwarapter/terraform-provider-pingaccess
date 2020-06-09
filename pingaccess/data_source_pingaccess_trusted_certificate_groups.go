package pingaccess

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pa "github.com/iwarapter/pingaccess-sdk-go/pingaccess"
)

func dataSourcePingAccessTrustedCertificateGroups() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourcePingAccessTrustedCertificateGroupsRead,
		Schema:      dataSourcePingAccessTrustedCertificateGroupsSchema(),
	}
}

func dataSourcePingAccessTrustedCertificateGroupsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"cert_ids": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"ignore_all_certificate_errors": {
			Type:     schema.TypeBool,
			Computed: true,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"skip_certificate_date_check": {
			Type:     schema.TypeBool,
			Computed: true,
		},
		"system_group": {
			Type:     schema.TypeBool,
			Computed: true,
		},
		"use_java_trust_store": {
			Type:     schema.TypeBool,
			Computed: true,
		},
	}
}

func dataSourcePingAccessTrustedCertificateGroupsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(*pa.Client).TrustedCertificateGroups
	input := &pa.GetTrustedCertificateGroupsCommandInput{
		Name: d.Get("name").(string),
	}
	result, _, err := svc.GetTrustedCertificateGroupsCommand(input)
	if err != nil {
		return diag.Diagnostics{diag.FromErr(fmt.Errorf("unable to read TrustedCertificateGroup: %s", err))}
	}
	if result == nil {
		return diag.Diagnostics{diag.FromErr(fmt.Errorf("unable to find TrustedCertificateGroup with the name '%s', result was nil", d.Get("name").(string)))}
	}
	if len(result.Items) != 1 {
		return diag.Diagnostics{diag.FromErr(fmt.Errorf("unable to find TrustedCertificateGroup with the name '%s' found '%d' results", d.Get("name").(string), len(result.Items)))}
	}
	d.SetId(result.Items[0].Id.String())
	return resourcePingAccessTrustedCertificateGroupsReadResult(d, result.Items[0])
}
