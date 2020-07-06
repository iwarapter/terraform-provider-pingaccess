package pingaccess

import (
	"bytes"
	"encoding/json"
	"fmt"
	"hash/crc32"
	"log"
	"strconv"
	"strings"

	"github.com/iwarapter/pingaccess-sdk-go/pingaccess/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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

func computedListOfString() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Computed: true,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	}
}

func acmeServerAccountsSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Computed: true,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"location": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"id": {
					Type:     schema.TypeString,
					Computed: true,
				},
			},
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
					Computed: true,
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

func requiredHiddenField() *schema.Schema {
	sch := hiddenField()
	sch.Optional = false
	sch.Required = true
	return sch
}

func expandHiddenFieldView(in []interface{}) *models.HiddenFieldView {
	hf := &models.HiddenFieldView{}
	for _, raw := range in {
		if raw == nil {
			return hf
		}
		l := raw.(map[string]interface{})
		if val, ok := l["value"]; ok {
			hf.Value = String(val.(string))
		}
		if val, ok := l["encrypted_value"]; ok {
			hf.EncryptedValue = String(val.(string))
		}
	}
	return hf
}

func flattenHiddenFieldView(in *models.HiddenFieldView) []map[string]interface{} {
	m := make([]map[string]interface{}, 0, 1)
	s := make(map[string]interface{})
	if in.Value != nil {
		s["value"] = *in.Value //TODO this is bad don't do this.
	}
	if in.EncryptedValue != nil {
		s["encrypted_value"] = *in.EncryptedValue
	}
	m = append(m, s)
	return m
}

func expandOAuthClientCredentialsView(in []interface{}) *models.OAuthClientCredentialsView {
	hf := &models.OAuthClientCredentialsView{}
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

func flattenOAuthClientCredentialsView(in *models.OAuthClientCredentialsView) []map[string]interface{} {
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
	return []interface{}{m}
}

func expandPolicyItem(in []interface{}) []*models.PolicyItem {
	var policies []*models.PolicyItem
	for _, raw := range in {
		policy := &models.PolicyItem{}
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

func flattenPolicyItem(in []*models.PolicyItem) []interface{} {
	var m []interface{}
	for _, v := range in {
		s := make(map[string]interface{})
		s["id"] = v.Id.String()
		s["type"] = *v.Type
		m = append(m, s)
	}
	return m
}

func expandPolicy(in []interface{}) map[string]*[]*models.PolicyItem {
	ca := map[string]*[]*models.PolicyItem{}

	webPolicies := make([]*models.PolicyItem, 0)
	apiPolicies := make([]*models.PolicyItem, 0)
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

func flattenPolicy(in map[string]*[]*models.PolicyItem) []interface{} {
	// m := make([]map[string]interface{}, 0, 1)
	var m []interface{}
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
	vs := make([]*string, 0, len(configured))
	for _, v := range configured {
		val := v.(string)
		if val != "" {
			vs = append(vs, &val)
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

func flattenChainCertificates(in []*models.ChainCertificateView) *schema.Set {
	var m []interface{}
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
	return hashString(buf.String())
}

func flattenLinkViewList(in []*models.LinkView) []interface{} {
	var m []interface{}
	for _, v := range in {
		m = append(m, flattenLinkView(v))
	}
	return m
}

func flattenLinkView(in *models.LinkView) map[string]interface{} {
	s := make(map[string]interface{})
	if in.Id != nil {
		s["id"] = *in.Id
	}
	if in.Location != nil {
		s["location"] = *in.Location
	}
	return s
}

// Searches a given set of RuleDescriptors for a matching className, when found it will check all fields types for
// a CONCEALED flag or COMPOSITE if CONCEALED, we massage the configuration to to remove the encryptedValue returned by
// current API and set the value back to the original defined. For COMPOSITE fields we then iterate recursively on its
// fields.
//
// TODO This has a drawback that we cannot detect drift in CONCEALED fields due to the way the PingAccess API works.
func maskConfigFromRuleDescriptors(desc *models.RuleDescriptorsView, input *string, originalConfig string, config string) string {
	for _, value := range desc.Items {
		if *value.ClassName == *input {
			config = maskConfigFromRuleDescriptor(value, String(""), originalConfig, config)
		}
	}
	return config
}

func maskConfigFromRuleDescriptor(desc *models.RuleDescriptorView, input *string, originalConfig string, config string) string {
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
func maskConfigFromDescriptors(desc *models.DescriptorsView, input *string, originalConfig string, config string) string {
	for _, value := range desc.Items {
		if *value.ClassName == *input {
			config = maskConfigFromDescriptor(value, String(""), originalConfig, config)
		}
	}
	return config
}

func maskConfigFromDescriptor(desc *models.DescriptorView, input *string, originalConfig string, config string) string {
	for _, c := range desc.ConfigurationFields {
		config = maskConfigFromConfigurationField(c, input, originalConfig, config)
	}
	return config
}

func maskConfigFromConfigurationField(field *models.ConfigurationField, input *string, originalConfig string, config string) string {
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

//Checks all the fields in the descriptor to ensure all required fields are set
//
func descriptorsHasClassName(className string, desc *models.DescriptorsView) error {
	var classes []string
	for _, value := range desc.Items {
		classes = append(classes, *value.ClassName)
		if *value.ClassName == className {
			return nil
		}
	}
	return fmt.Errorf("unable to find className '%s' available classNames: %s", className, strings.Join(classes, ", "))
}

//Checks the class name specified exists in the DescriptorsView
//
func validateConfiguration(className string, d *schema.ResourceDiff, desc *models.DescriptorsView) error {
	var diags diag.Diagnostics
	conf := d.Get("configuration").(string)
	if conf == "" {
		log.Println("[INFO] configuration is in a potentially unknown state, gracefully skipping configuration validation")
		return nil
	}
	for _, value := range desc.Items {
		if *value.ClassName == className {
			for _, f := range value.ConfigurationFields {
				if *f.Required {
					if v := gjson.Get(conf, *f.Name); !v.Exists() {
						diags = append(diags, diag.Errorf("the field '%s' is required for the class_name '%s'", *f.Name, className)...)
					}
				}
			}
		}
	}
	if diags.HasError() {
		msgs := []string{
			"configuration validation failed against the class descriptor definition",
		}
		for _, diagnostic := range diags {
			msgs = append(msgs, diagnostic.Summary)
		}
		return fmt.Errorf(strings.Join(msgs, "\n"))
	}
	return nil
}

//Checks the class name specified exists in the RuleDescriptorsView
//
func ruleDescriptorsHasClassName(className string, desc *models.RuleDescriptorsView) error {
	var classes []string
	for _, value := range desc.Items {
		classes = append(classes, *value.ClassName)
		if *value.ClassName == className {
			return nil
		}
	}
	return fmt.Errorf("unable to find className '%s' available classNames: %s", className, strings.Join(classes, ", "))
}

//Checks all the fields in the Rule descriptor to ensure all required fields are set
//
func validateRulesConfiguration(className string, d *schema.ResourceDiff, desc *models.RuleDescriptorsView) error {
	var diags diag.Diagnostics
	conf := d.Get("configuration").(string)
	if conf == "" {
		log.Println("[INFO] configuration is in a potentially unknown state, gracefully skipping configuration validation")
		return nil
	}
	for _, value := range desc.Items {
		if *value.ClassName == className {
			for _, f := range value.ConfigurationFields {
				if *f.Required {
					if v := gjson.Get(conf, *f.Name); !v.Exists() {
						diags = append(diags, diag.Errorf("the field '%s' is required for the class_name '%s'", *f.Name, className)...)
					}
				}
			}
		}
	}
	if diags.HasError() {
		msgs := []string{
			"configuration validation failed against the class descriptor definition",
		}
		for _, diagnostic := range diags {
			msgs = append(msgs, diagnostic.Summary)
		}
		return fmt.Errorf(strings.Join(msgs, "\n"))
	}
	return nil
}

// String hashes a string to a unique hashcode.
//
// crc32 returns a uint32, but for our use we need
// and non negative integer. Here we cast to an integer
// and invert it if the result is negative.
func hashString(s string) int {
	v := int(crc32.ChecksumIEEE([]byte(s)))
	if v >= 0 {
		return v
	}
	if -v >= 0 {
		return -v
	}
	// v == MinInt
	return 0
}
