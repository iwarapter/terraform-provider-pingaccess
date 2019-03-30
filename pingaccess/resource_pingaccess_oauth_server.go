package pingaccess

import (
	"fmt"

	pa "github.com/iwarapter/pingaccess-sdk-go/pingaccess"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourcePingAccessOAuthServer() *schema.Resource {
	return &schema.Resource{
		Create: resourcePingAccessOAuthServerCreate,
		Read:   resourcePingAccessOAuthServerRead,
		Update: resourcePingAccessOAuthServerUpdate,
		Delete: resourcePingAccessOAuthServerDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: resourcePingAccessOAuthServerSchema(),
	}
}

func resourcePingAccessOAuthServerSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"audit_level": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Default:  "ON",
		},
		"cache_tokens": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
		},
		"client_credentials": oAuthClientCredentials(),
		"description": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"introspection_endpoint": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},

		"secure": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
		},
		"send_audience": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
		},
		"subject_attribute_name": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"targets": &schema.Schema{
			Type:     schema.TypeSet,
			Required: true,
			MinItems: 1,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"token_time_to_live_seconds": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  -1,
		},
		"trusted_certificate_group_id": &schema.Schema{
			Type:     schema.TypeInt,
			Required: true,
		},
		"use_proxy": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
		},
	}
}

func resourcePingAccessOAuthServerCreate(d *schema.ResourceData, m interface{}) error {
	d.SetId("oauth_server_settings")
	return resourcePingAccessOAuthServerUpdate(d, m)
}

func resourcePingAccessOAuthServerRead(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pa.Client).OAuth
	result, _, err := svc.GetAuthorizationServerCommand()
	if err != nil {
		return fmt.Errorf("Error reading oauth server settings: %s", err)
	}

	return resourcePingAccessOAuthServerReadResult(d, result)
}

func resourcePingAccessOAuthServerUpdate(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pa.Client).OAuth
	input := pa.UpdateAuthorizationServerCommandInput{
		Body: *resourcePingAccessOAuthServerReadData(d),
	}
	result, _, err := svc.UpdateAuthorizationServerCommand(&input)
	if err != nil {
		return fmt.Errorf("Error updating oauth server settings: %s", err.Error())
	}

	d.SetId("oauth_server_settings")
	return resourcePingAccessOAuthServerReadResult(d, result)
}

func resourcePingAccessOAuthServerDelete(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pa.Client).OAuth
	_, err := svc.DeleteAuthorizationServerCommand()
	if err != nil {
		return fmt.Errorf("Error resetting oauth server: %s", err)
	}
	return nil
}

func resourcePingAccessOAuthServerReadResult(d *schema.ResourceData, input *pa.AuthorizationServerView) (err error) {
	setResourceDataString(d, "audit_level", input.AuditLevel)
	setResourceDataBool(d, "cache_tokens", input.CacheTokens)
	setResourceDataString(d, "description", input.Description)
	setResourceDataString(d, "introspection_endpoint", input.IntrospectionEndpoint)
	setResourceDataBool(d, "secure", input.Secure)
	setResourceDataBool(d, "send_audience", input.SendAudience)
	setResourceDataString(d, "subject_attribute_name", input.SubjectAttributeName)
	setResourceDataInt(d, "token_time_to_live_seconds", input.TokenTimeToLiveSeconds)
	setResourceDataInt(d, "trusted_certificate_group_id", input.TrustedCertificateGroupId)
	setResourceDataBool(d, "use_proxy", input.UseProxy)

	if err := d.Set("targets", input.Targets); err != nil {
		return err
	}
	if input.ClientCredentials != nil {
		//TODO we should look at encrypting the value
		pw, ok := d.GetOkExists("client_credentials.0.client_secret.0.value")
		creds := flattenOAuthClientCredentialsView(input.ClientCredentials)
		if ok {
			creds[0]["client_secret"].([]map[string]interface{})[0]["value"] = pw
		}
		if err = d.Set("client_credentials", creds); err != nil {
			return err
		}
	}

	return nil
}

func resourcePingAccessOAuthServerReadData(d *schema.ResourceData) *pa.AuthorizationServerView {
	targets := expandStringList(d.Get("targets").(*schema.Set).List())
	oauth := &pa.AuthorizationServerView{
		IntrospectionEndpoint:     String(d.Get("introspection_endpoint").(string)),
		SubjectAttributeName:      String(d.Get("subject_attribute_name").(string)),
		Targets:                   &targets,
		TrustedCertificateGroupId: Int(d.Get("trusted_certificate_group_id").(int)),
		ClientCredentials:         expandOAuthClientCredentialsView(d.Get("client_credentials").([]interface{})),
	}

	if v, ok := d.GetOkExists("audit_level"); ok {
		oauth.AuditLevel = String(v.(string))
	}

	if v, ok := d.GetOkExists("cache_tokens"); ok {
		oauth.CacheTokens = Bool(v.(bool))
	}

	if v, ok := d.GetOkExists("description"); ok {
		oauth.Description = String(v.(string))
	}

	if v, ok := d.GetOkExists("secure"); ok {
		oauth.Secure = Bool(v.(bool))
	}

	if v, ok := d.GetOkExists("send_audience"); ok {
		oauth.SendAudience = Bool(v.(bool))
	}

	if v, ok := d.GetOkExists("token_time_to_live_seconds"); ok {
		oauth.TokenTimeToLiveSeconds = Int(v.(int))
	}

	if v, ok := d.GetOkExists("use_proxy"); ok {
		oauth.UseProxy = Bool(v.(bool))
	}

	return oauth
}
