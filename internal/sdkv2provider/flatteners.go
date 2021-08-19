package sdkv2provider

import "github.com/iwarapter/pingaccess-sdk-go/pingaccess/models"

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
