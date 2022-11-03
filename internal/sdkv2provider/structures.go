package sdkv2provider

import (
	"bytes"
	"encoding/json"
	"fmt"
	"hash/crc32"
	"log"
	"strconv"
	"strings"

	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

func acmeServerAccountsSchema() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Computed:    true,
		Optional:    true,
		Description: "An array of references to accounts. This array is read-only.",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"location": {
					Type:        schema.TypeString,
					Computed:    true,
					Description: "An absolute path to the associated resource.",
				},
				"id": {
					Type:        schema.TypeString,
					Computed:    true,
					Description: "The id of the associated resource. When both id and location are specified, id takes precedence and location is ignored.",
				},
			},
		},
	}
}

func applicationPolicySchema() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Optional:    true,
		MaxItems:    1,
		Description: "A map of policy items associated with the resource. The key is 'web' or 'api' and the value is a list of Policy Items.",
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
		Type:        schema.TypeList,
		Optional:    true,
		Description: "List of Rule/RuleSets to be applied.",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"type": {
					Type:        schema.TypeString,
					Required:    true,
					Description: "If this is either a `Rule` or `RuleSet`.",
				},
				"id": {
					Type:        schema.TypeString,
					Required:    true,
					Description: "The ID of the specific rule or ruleset.",
				},
			},
		},
	}
}

func oAuthClientCredentialsResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"client_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Specify the client ID.",
			},
			"client_secret": hiddenField(),
			"key_pair_id": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     0,
				Description: "Specify the ID of a key pair to use for mutual TLS.",
			},
			"credentials_type": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "SECRET",
				Description:  "Specify the credential type.",
				ValidateFunc: validation.StringInSlice([]string{"SECRET", "CERTIFICATE", "PRIVATE_KEY_JWT"}, false),
			},
		},
	}
}

func hiddenFieldResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"encrypted_value": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "encrypted value of the field, as originally returned by the API.",
			},
			"value": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				Description: "The value of the field. This field takes precedence over the encryptedValue field, if both are specified.",
			},
		},
	}
}

func hiddenField() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Optional:    true,
		MaxItems:    1,
		Description: "Specify the client secret.",
		Elem:        hiddenFieldResource(),
	}
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
		if val, ok := l["client_secret"]; ok && len(val.([]interface{})) > 0 {
			hf.ClientSecret = expandHiddenFieldView(val.([]interface{}))
		}
		if val, ok := l["credentials_type"]; ok {
			hf.CredentialsType = String(val.(string))
		}
		if val, ok := l["key_pair_id"]; ok {
			hf.KeyPairId = Int(val.(int))
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

func flattenPathPatternView(in []*models.PathPatternView) []interface{} {
	var m []interface{}
	for _, v := range in {
		s := make(map[string]interface{})
		s["pattern"] = v.Pattern
		s["type"] = v.Type
		m = append(m, s)
	}
	return m
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
func maskConfigFromDescriptors(desc *models.DescriptorsView, className *string, originalConfig string, config string) string {
	for _, value := range desc.Items {
		if *value.ClassName == *className {
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

// Checks all the fields in the descriptor to ensure all required fields are set
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

// Checks the class name specified exists in the DescriptorsView
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

// Checks the class name specified exists in the RuleDescriptorsView
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

// Checks all the fields in the Rule descriptor to ensure all required fields are set
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

func setClientCredentials(d *schema.ResourceData, input *models.OAuthClientCredentialsView, trackPasswords bool, diags *diag.Diagnostics) {
	pw, ok := d.GetOk("client_credentials.0.client_secret.0.value")
	creds := flattenOAuthClientCredentialsView(input)

	if input.KeyPairId != nil {
		creds[0]["key_pair_id"] = *input.KeyPairId
	}
	if input.CredentialsType != nil {
		creds[0]["credentials_type"] = *input.CredentialsType
	} else {
		//set the resource state default
		creds[0]["credentials_type"] = "SECRET"
	}
	if _, hasClientSecret := creds[0]["client_secret"]; hasClientSecret && trackPasswords {
		enc, encOk := d.GetOk("client_credentials.0.client_secret.0.encrypted_value")
		creds[0]["client_secret"].([]map[string]interface{})[0]["value"] = pw
		if err := d.Set("client_credentials", creds); err != nil {
			*diags = append(*diags, diag.FromErr(err)...)
		}
		if encOk && enc.(string) != "" && enc.(string) != *input.ClientSecret.EncryptedValue {
			creds[0]["client_secret"].([]map[string]interface{})[0]["value"] = ""
		}
	} else {
		//legacy behaviour
		if ok {
			creds[0]["client_secret"].([]map[string]interface{})[0]["value"] = pw
		}
	}
	if err := d.Set("client_credentials", creds); err != nil {
		*diags = append(*diags, diag.FromErr(err)...)
	}
}

func expandResourceTypeConfiguration(in []interface{}) *models.ResourceTypeConfigurationView {
	rtcv := &models.ResourceTypeConfigurationView{}

	for _, raw := range in {
		l := raw.(map[string]interface{})
		if val, ok := l["response_generator"]; ok {
			rtcv.ResponseGenerator = expandResponseGenerator(val.([]interface{}))
		}
	}

	return rtcv
}

func expandResponseGenerator(in []interface{}) *models.ResponseGeneratorView {
	rg := &models.ResponseGeneratorView{}
	for _, raw := range in {
		if raw == nil {
			return rg
		}
		l := raw.(map[string]interface{})
		if val, ok := l["class_name"]; ok {
			rg.ClassName = String(val.(string))
		}
		if val, ok := l["configuration"]; ok {
			config := val.(string)
			var dat map[string]interface{}
			_ = json.Unmarshal([]byte(config), &dat)
			rg.Configuration = dat
		}
	}

	return rg
}

func flattenResourceTypeConfiguration(in *models.ResourceTypeConfigurationView) []interface{} {
	var m []interface{}

	s := make(map[string]interface{})
	if in.ResponseGenerator != nil {
		s["response_generator"] = flattenResponseGenerator(in.ResponseGenerator)
		m = append(m, s)
	}

	return m
}

func flattenResponseGenerator(in *models.ResponseGeneratorView) []interface{} {
	var m []interface{}
	s := make(map[string]interface{})
	if in.ClassName != nil {
		s["class_name"] = *in.ClassName
	}

	if in.Configuration != nil {
		b, _ := json.Marshal(in.Configuration)
		s["configuration"] = String(string(b))
	}
	m = append(m, s)

	return m
}
