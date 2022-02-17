package sdkv2provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess/models"
)

func flattenRuntimeApplication(in *models.PingFederateRuntimeApplicationView) []map[string]interface{} {
	m := make([]map[string]interface{}, 0, 1)
	s := make(map[string]interface{})
	if in.AdditionalVirtualHostIds != nil {
		s["additional_virtual_host_ids"] = *in.AdditionalVirtualHostIds
	}
	if in.CaseSensitive != nil {
		s["case_sensitive"] = *in.CaseSensitive
	}
	if in.ClientCertHeaderNames != nil {
		s["client_cert_header_names"] = *in.ClientCertHeaderNames
	}
	if in.ContextRoot != nil {
		s["context_root"] = *in.ContextRoot
	}
	if in.Policy != nil {
		s["policy"] = flattenPolicyItem(in.Policy)
	}
	if in.PrimaryVirtualHostId != nil {
		s["primary_virtual_host_id"] = *in.PrimaryVirtualHostId
	}
	m = append(m, s)
	return m
}

func flattenPolicies(d *schema.ResourceData, policies map[string]*[]*models.PolicyItem) diag.Diagnostics {
	if policies != nil {
		if len(*policies["API"]) > 0 || len(*policies["Web"]) > 0 || policyStateHasData(d) {
			if err := d.Set("policy", flattenPolicy(policies)); err != nil {
				return diag.FromErr(err)
			}
		}
	}
	return nil
}
