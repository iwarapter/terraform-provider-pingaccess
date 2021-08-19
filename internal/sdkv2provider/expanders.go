package sdkv2provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/iwarapter/pingaccess-sdk-go/pingaccess/models"
)

func expandRuntimeApplication(in []interface{}) *models.PingFederateRuntimeApplicationView {
	app := &models.PingFederateRuntimeApplicationView{}
	for _, raw := range in {
		if raw == nil {
			return app
		}
		l := raw.(map[string]interface{})
		if val, ok := l["additional_virtual_host_ids"]; ok {
			vhosts := expandIntList(val.(*schema.Set).List())
			app.AdditionalVirtualHostIds = &vhosts
		}
		if val, ok := l["case_sensitive"]; ok {
			app.CaseSensitive = Bool(val.(bool))
		}
		if val, ok := l["client_cert_header_names"]; ok {
			names := expandStringList(val.(*schema.Set).List())
			app.ClientCertHeaderNames = &names
		}
		if val, ok := l["context_root"]; ok {
			app.ContextRoot = String(val.(string))
		}
		if val, ok := l["primary_virtual_host_id"]; ok {
			app.PrimaryVirtualHostId = Int(val.(int))
		}
		if val, ok := l["policy"]; ok {
			app.Policy = expandPolicyItem(val.([]interface{}))
		}
	}
	return app
}
