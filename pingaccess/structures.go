package pingaccess

import (
	"strconv"
)

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
