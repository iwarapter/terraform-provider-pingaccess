package pingaccess

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	pa "github.com/iwarapter/pingaccess-sdk-go/pingaccess"
)

func setOfString() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeSet,
		Optional: true,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	}
}

func requiredListOfString() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Required: true,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	}
}

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
		// DefaultFunc: func() (interface{}, error) {
		// 	return []map[string][]interface{}{}, nil
		// },
	}
}

func applicationPolicyItemSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
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

func oAuthClientCredentials() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Required: true,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"client_id": {
					Type:     schema.TypeString,
					Required: true,
				},
				"client_secret": hiddenField(),
			},
		},
	}
}

func hiddenField() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"encrypted_value": {
					Type:     schema.TypeString,
					Optional: true,
				},
				"value": {
					Type:      schema.TypeString,
					Optional:  true,
					Sensitive: true,
				},
			},
		},
	}
}

func expandHiddenFieldView(in []interface{}) *pa.HiddenFieldView {
	hf := &pa.HiddenFieldView{}
	for _, raw := range in {
		if raw == nil {
			return hf
		}
		l := raw.(map[string]interface{})
		if val, ok := l["value"]; ok {
			hf.Value = String(val.(string))
		}
		// if val, ok := l["encrypted_value"]; ok {
		// 	hf.EncryptedValue = String(val.(string))
		// }
	}
	return hf
}

func flattenHiddenFieldView(in *pa.HiddenFieldView) []map[string]interface{} {
	m := make([]map[string]interface{}, 0, 1)
	s := make(map[string]interface{})
	if in.Value != nil {
		s["value"] = *in.Value //TODO this is bad don't do this.
	}
	// if in.EncryptedValue != nil {
	// s["encrypted_value"] = *in.EncryptedValue
	// }
	m = append(m, s)
	return m
}

func expandOAuthClientCredentialsView(in []interface{}) *pa.OAuthClientCredentialsView {
	hf := &pa.OAuthClientCredentialsView{}
	for _, raw := range in {
		l := raw.(map[string]interface{})
		if val, ok := l["client_id"]; ok {
			hf.ClientId = String(val.(string))
		}
		if val, ok := l["client_secret"]; ok {
			hf.ClientSecret = expandHiddenFieldView(val.([]interface{}))
		}
	}
	return hf
}

func flattenOAuthClientCredentialsView(in *pa.OAuthClientCredentialsView) []map[string]interface{} {
	m := make([]map[string]interface{}, 0, 1)
	s := make(map[string]interface{})
	if in.ClientId != nil {
		s["client_id"] = *in.ClientId
	}
	if in.ClientSecret != nil {
		s["client_secret"] = flattenHiddenFieldView(in.ClientSecret)
	}
	m = append(m, s)
	return m
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
	log.Printf("FLATTENER: %v, %v", *in["API"], *in["Web"])
	// if m["api"] == "0" && m["web"] == "0" {
	// 	return []interface{}{}
	// }
	return []interface{}{m}
}

func expandPolicyItem(in []interface{}) []*pa.PolicyItem {
	policies := []*pa.PolicyItem{}
	for _, raw := range in {
		policy := &pa.PolicyItem{}
		l := raw.(map[string]interface{})
		if val, ok := l["id"]; ok {
			policy.Id = json.Number(val.(string))
		}
		if val, ok := l["type"]; ok {
			policy.Type = String(val.(string))
		}
		policies = append(policies, policy)
	}
	return policies
}

func flattenPolicyItem(in []*pa.PolicyItem) []interface{} {
	m := []interface{}{}
	for _, v := range in {
		s := make(map[string]interface{})
		s["id"] = v.Id.String()
		s["type"] = *v.Type
		m = append(m, s)
	}
	return m
}

func policyItemHash(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	buf.WriteString(m["id"].(string))
	if d, ok := m["type"]; ok && d.(string) != "" {
		buf.WriteString(fmt.Sprintf("%s-", d.(string)))
	}
	return hashcode.String(buf.String())
}

func stringHash(v interface{}) int {
	var buf bytes.Buffer
	buf.WriteString(v.(string))
	return hashcode.String(buf.String())
}

func expandPolicy(in []interface{}) map[string]*[]*pa.PolicyItem {
	ca := map[string]*[]*pa.PolicyItem{}

	webPolicies := make([]*pa.PolicyItem, 0)
	apiPolicies := make([]*pa.PolicyItem, 0)
	for _, raw := range in {
		if raw != nil {
			l := raw.(map[string]interface{})
			if val, ok := l["web"]; ok && len(val.([]interface{})) > 0 {
				webPolicies = expandPolicyItem(val.([]interface{}))
			}
			if val, ok := l["api"]; ok && len(val.([]interface{})) > 0 {
				apiPolicies = expandPolicyItem(val.([]interface{}))
			}
		}
	}

	ca["Web"] = &webPolicies
	ca["API"] = &apiPolicies

	return ca
}

func flattenPolicy(in map[string]*[]*pa.PolicyItem) []interface{} {
	// m := make([]map[string]interface{}, 0, 1)
	m := []interface{}{}
	s := make(map[string]interface{})
	if val, ok := in["Web"]; ok { // && len(*in["Web"]) > 0 {
		s["web"] = flattenPolicyItem(*val)
	}
	if val, ok := in["API"]; ok { //&& len(*in["API"]) > 0 {
		s["api"] = flattenPolicyItem(*val)
	}
	m = append(m, s)
	return m
	// return //[]interface{}{s}
}

// Takes the result of flatmap.Expand for an array of strings
// and returns a []string
func expandStringList(configured []interface{}) []*string {
	// log.Printf("[INFO] expandStringList %d", len(configured))
	vs := make([]*string, 0, len(configured))
	for _, v := range configured {
		val := v.(string)
		if val != "" {
			vs = append(vs, &val)
			// log.Printf("[DEBUG] Appending: %s", val)
		}
	}
	return vs
}

// Takes the result of flatmap.Expand for an array of strings
// and returns a []*int
func expandIntList(configured []interface{}) []*int {
	vs := make([]*int, 0, len(configured))
	for _, v := range configured {
		_, ok := v.(int)
		if ok {
			val := v.(int)
			vs = append(vs, &val)
		}
	}
	return vs
}

func flattenChainCertificates(in []*pa.ChainCertificateView) *schema.Set {
	m := []interface{}{}
	for _, v := range in {
		s := make(map[string]interface{})
		s["alias"] = *v.Alias
		s["expires"] = v.Expires
		s["issuer_dn"] = *v.IssuerDn
		s["md5sum"] = *v.Md5sum
		s["serial_number"] = *v.SerialNumber
		s["sha1sum"] = *v.Sha1sum
		s["signature_algorithm"] = *v.SignatureAlgorithm
		s["status"] = *v.Status
		s["subject_cn"] = *v.SubjectCn
		s["subject_dn"] = *v.SubjectDn
		s["valid_from"] = v.ValidFrom
		m = append(m, s)
	}
	return schema.NewSet(configFieldHash, m)
}

func configFieldHash(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	buf.WriteString(m["alias"].(string))
	if d, ok := m["md5sum"]; ok && d.(string) != "" {
		buf.WriteString(fmt.Sprintf("%s-", d.(string)))
	}
	return hashcode.String(buf.String())
}

// Searches a given set of RuleDescriptors for a matching className, when found it will check all fields types for
// a CONCEALED flag or COMPOSITE if CONCEALED, we massage the configuration to to remove the encryptedValue returned by
// current API and set the value back to the original defined. For COMPOSITE fields we then iterate recursively on its
// fields.
//
// TODO This has a drawback that we cannot detect drift in CONCEALED fields due to the way the PingAccess API works.
func maskConfigFromRuleDescriptors(desc *pa.RuleDescriptorsView, input *string, originalConfig string, config string) string {
	for _, value := range desc.Items {
		if *value.ClassName == *input {
			config = maskConfigFromRuleDescriptor(value, String(""), originalConfig, config)
		}
	}
	return config
}

func maskConfigFromRuleDescriptor(desc *pa.RuleDescriptorView, input *string, originalConfig string, config string) string {
	for _, c := range desc.ConfigurationFields {
		config = maskConfigFromConfigurationField(c, input, originalConfig, config)
	}
	return config
}


// Searches a given set of descriptors for a matching className, when found it will check all fields types for
// a CONCEALED flag or COMPOSITE if CONCEALED, we massage the configuration to to remove the encryptedValue returned by
// current API and set the value back to the original defined. For COMPOSITE fields we then iterate recursively on its
// fields.
//
// TODO This has a drawback that we cannot detect drift in CONCEALED fields due to the way the PingAccess API works.
func maskConfigFromDescriptors(desc *pa.DescriptorsView, input *string, originalConfig string, config string) string {
	for _, value := range desc.Items {
		if *value.ClassName == *input {
			config = maskConfigFromDescriptor(value, String(""), originalConfig, config)
		}
	}
	return config
}

func maskConfigFromDescriptor(desc *pa.DescriptorView, input *string, originalConfig string, config string) string {
	for _, c := range desc.ConfigurationFields {
		config = maskConfigFromConfigurationField(c, input, originalConfig, config)
	}
	return config
}

func maskConfigFromConfigurationField(field *pa.ConfigurationField, input *string, originalConfig string, config string) string {
	if *field.Type == "CONCEALED" {
		path := fmt.Sprintf("%s.value", *field.Name)
		v := gjson.Get(originalConfig, path)
		if v.Exists() {
			config, _ = sjson.Set(config, path, v.String())
		} else if v = gjson.Get(originalConfig, *field.Name); v.Exists() {
			config, _ = sjson.Set(config, *field.Name, v.String())
		}
		config, _ = sjson.Delete(config, fmt.Sprintf("%s.encryptedValue", *field.Name))
	} else if *field.Type == "COMPOSITE" {
		for _, value := range field.Fields {
			newInput := String(fmt.Sprintf("%s.%s", *input, *value.Name))
			if *input == "" {
				newInput = value.Name
			}
			config = maskConfigFromConfigurationField(value, newInput, originalConfig, config)
		}
	}
	return config
}
