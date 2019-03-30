package pingaccess

import (
	"fmt"

	pa "github.com/iwarapter/pingaccess-sdk-go/pingaccess"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourcePingAccessPingFederateOAuth() *schema.Resource {
	return &schema.Resource{
		Create: resourcePingAccessPingFederateOAuthCreate,
		Read:   resourcePingAccessPingFederateOAuthRead,
		Update: resourcePingAccessPingFederateOAuthUpdate,
		Delete: resourcePingAccessPingFederateOAuthDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: resourcePingAccessPingFederateOAuthSchema(),
	}
}

func resourcePingAccessPingFederateOAuthSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"access_validator_id": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  1,
		},
		"cache_tokens": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
		},
		"client_id": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"client_secret": hiddenField(),
		"name": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Default:  "PingFederate",
		},
		"send_audience": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
		},
		"subject_attribute_name": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"token_time_to_live_seconds": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  -1,
		},
		"use_token_introspection": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
		},
	}
}

func resourcePingAccessPingFederateOAuthCreate(d *schema.ResourceData, m interface{}) error {
	d.SetId("pingfederate_oauth_settings")
	return resourcePingAccessPingFederateOAuthUpdate(d, m)
}

func resourcePingAccessPingFederateOAuthRead(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pa.Client).PingFederate
	result, _, err := svc.GetPingFederateAccessTokensCommand()
	if err != nil {
		return fmt.Errorf("Error reading pingfederate oauth settings: %s", err)
	}

	return resourcePingAccessPingFederateOAuthReadResult(d, result)
}

func resourcePingAccessPingFederateOAuthUpdate(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pa.Client).PingFederate
	input := pa.UpdatePingFederateAccessTokensCommandInput{
		Body: *resourcePingAccessPingFederateOAuthReadData(d),
	}
	result, _, err := svc.UpdatePingFederateAccessTokensCommand(&input)
	if err != nil {
		return fmt.Errorf("Error updating pingfederate oauth settings: %s", err)
	}

	d.SetId("pingfederate_oauth_settings")
	return resourcePingAccessPingFederateOAuthReadResult(d, result)
}

func resourcePingAccessPingFederateOAuthDelete(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pa.Client).PingFederate
	_, err := svc.DeletePingFederateAccessTokensCommand()
	if err != nil {
		return fmt.Errorf("Error resetting pingfederate oauth: %s", err)
	}
	return nil
}

func resourcePingAccessPingFederateOAuthReadResult(d *schema.ResourceData, input *pa.PingFederateAccessTokenView) error {
	setResourceDataInt(d, "access_validator_id", input.AccessValidatorId)
	setResourceDataBool(d, "cache_tokens", input.CacheTokens)
	setResourceDataString(d, "client_id", input.ClientId)
	setResourceDataString(d, "name", input.Name)
	setResourceDataBool(d, "send_audience", input.SendAudience)
	setResourceDataString(d, "subject_attribute_name", input.SubjectAttributeName)
	setResourceDataInt(d, "token_time_to_live_seconds", input.TokenTimeToLiveSeconds)
	setResourceDataBool(d, "use_token_introspection", input.UseTokenIntrospection)

	if input.ClientSecret != nil {
		pw, ok := d.GetOkExists("client_secret.0.value")
		creds := flattenHiddenFieldView(input.ClientSecret)
		if ok {
			creds[0]["value"] = pw
		}
		if err := d.Set("client_secret", creds); err != nil {
			return err
		}
		// if err := d.Set("client_secret", flattenHiddenFieldView(input.ClientSecret)); err != nil {
		// 	return err
		// }
	}

	return nil
}

func resourcePingAccessPingFederateOAuthReadData(d *schema.ResourceData) *pa.PingFederateAccessTokenView {
	oauth := &pa.PingFederateAccessTokenView{
		ClientId:             String(d.Get("client_id").(string)),
		SubjectAttributeName: String(d.Get("subject_attribute_name").(string)),
	}

	if v, ok := d.GetOkExists("access_validator_id"); ok {
		oauth.AccessValidatorId = Int(v.(int))
	}

	if v, ok := d.GetOkExists("cache_tokens"); ok {
		oauth.CacheTokens = Bool(v.(bool))
	}

	if v, ok := d.GetOkExists("name"); ok {
		oauth.Name = String(v.(string))
	}

	if v, ok := d.GetOkExists("send_audience"); ok {
		oauth.SendAudience = Bool(v.(bool))
	}

	if v, ok := d.GetOkExists("token_time_to_live_seconds"); ok {
		oauth.TokenTimeToLiveSeconds = Int(v.(int))
	}

	if v, ok := d.GetOkExists("use_token_introspection"); ok {
		oauth.UseTokenIntrospection = Bool(v.(bool))
	}

	if v, ok := d.GetOkExists("client_secret"); ok {
		oauth.ClientSecret = expandHiddenFieldView(v.([]interface{}))
	}

	return oauth
}
