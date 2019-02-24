package pingaccess

import (
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"
)

func applicationPolicySchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"web": applicationPolicyItemSchema(),
				"api": applicationPolicyItemSchema(),
			},
		},
	}
}

func applicationPolicyItemSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeSet,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"type": {
					Type:     schema.TypeString,
					Required: true,
				},
				"id": {
					Type:     schema.TypeString,
					Required: true,
				},
			},
		},
	}
}

func flattenIdentityMappingIds(in map[string]*int) []interface{} {
	// NOTE: the top level structure to set is a map
	m := make(map[string]interface{})
	if in["Web"] != nil {
		m["web"] = strconv.Itoa(*in["Web"])
	}
	if in["API"] != nil {
		m["api"] = strconv.Itoa(*in["API"])
	}
	if m["api"] == "0" && m["web"] == "0" {
		return []interface{}{}
	}
	return []interface{}{m}
}
