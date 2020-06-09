package pingaccess

import (
	"context"
	"fmt"

	pa "github.com/iwarapter/pingaccess-sdk-go/pingaccess"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePingAccessOAuthServer() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePingAccessOAuthServerCreate,
		ReadContext:   resourcePingAccessOAuthServerRead,
		UpdateContext: resourcePingAccessOAuthServerUpdate,
		DeleteContext: resourcePingAccessOAuthServerDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: resourcePingAccessOAuthServerSchema(),
	}
}

func resourcePingAccessOAuthServerSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"audit_level": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "ON",
		},
		"cache_tokens": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"client_credentials": oAuthClientCredentials(),
		"description": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"introspection_endpoint": {
			Type:     schema.TypeString,
			Required: true,
		},

		"secure": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"send_audience": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"subject_attribute_name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"targets": {
			Type:     schema.TypeSet,
			Required: true,
			MinItems: 1,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"token_time_to_live_seconds": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  -1,
		},
		"trusted_certificate_group_id": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"use_proxy": {
			Type:     schema.TypeBool,
			Optional: true,
		},
	}
}

func resourcePingAccessOAuthServerCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	d.SetId("oauth_server_settings")
	return resourcePingAccessOAuthServerUpdate(ctx, d, m)
}

func resourcePingAccessOAuthServerRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(*pa.Client).OAuth
	result, _, err := svc.GetAuthorizationServerCommand()
	if err != nil {
		return diag.Diagnostics{diag.FromErr(fmt.Errorf("unable to read OAuthServerSettings: %s", err))}
	}

	return resourcePingAccessOAuthServerReadResult(d, result)
}

func resourcePingAccessOAuthServerUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(*pa.Client).OAuth
	input := pa.UpdateAuthorizationServerCommandInput{
		Body: *resourcePingAccessOAuthServerReadData(d),
	}
	result, _, err := svc.UpdateAuthorizationServerCommand(&input)
	if err != nil {
		return diag.Diagnostics{diag.FromErr(fmt.Errorf("unable to update OAuthServerSettings: %s", err))}
	}

	d.SetId("oauth_server_settings")
	return resourcePingAccessOAuthServerReadResult(d, result)
}

func resourcePingAccessOAuthServerDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(*pa.Client).OAuth
	_, err := svc.DeleteAuthorizationServerCommand()
	if err != nil {
		return diag.Diagnostics{diag.FromErr(fmt.Errorf("unable to reset OAuthServerSettings: %s", err))}
	}
	return nil
}

func resourcePingAccessOAuthServerReadResult(d *schema.ResourceData, input *pa.AuthorizationServerView) diag.Diagnostics {
	var diags diag.Diagnostics
	setResourceDataStringWithDiagnostic(d, "audit_level", input.AuditLevel, &diags)
	setResourceDataBoolWithDiagnostic(d, "cache_tokens", input.CacheTokens, &diags)
	setResourceDataStringWithDiagnostic(d, "description", input.Description, &diags)
	setResourceDataStringWithDiagnostic(d, "introspection_endpoint", input.IntrospectionEndpoint, &diags)
	setResourceDataBoolWithDiagnostic(d, "secure", input.Secure, &diags)
	setResourceDataBoolWithDiagnostic(d, "send_audience", input.SendAudience, &diags)
	setResourceDataStringWithDiagnostic(d, "subject_attribute_name", input.SubjectAttributeName, &diags)
	setResourceDataIntWithDiagnostic(d, "token_time_to_live_seconds", input.TokenTimeToLiveSeconds, &diags)
	setResourceDataIntWithDiagnostic(d, "trusted_certificate_group_id", input.TrustedCertificateGroupId, &diags)
	setResourceDataBoolWithDiagnostic(d, "use_proxy", input.UseProxy, &diags)

	if err := d.Set("targets", input.Targets); err != nil {
		diags = append(diags, diag.FromErr(err))
	}
	if input.ClientCredentials != nil {
		//TODO we should look at encrypting the value
		pw, ok := d.GetOk("client_credentials.0.client_secret.0.value")
		creds := flattenOAuthClientCredentialsView(input.ClientCredentials)
		if ok {
			creds[0]["client_secret"].([]map[string]interface{})[0]["value"] = pw
		}
		if err := d.Set("client_credentials", creds); err != nil {
			diags = append(diags, diag.FromErr(err))
		}
	}

	return diags
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

	if v, ok := d.GetOk("audit_level"); ok {
		oauth.AuditLevel = String(v.(string))
	}

	if v, ok := d.GetOk("cache_tokens"); ok {
		oauth.CacheTokens = Bool(v.(bool))
	}

	if v, ok := d.GetOk("description"); ok {
		oauth.Description = String(v.(string))
	}

	if v, ok := d.GetOk("secure"); ok {
		oauth.Secure = Bool(v.(bool))
	}

	if v, ok := d.GetOk("send_audience"); ok {
		oauth.SendAudience = Bool(v.(bool))
	}

	if v, ok := d.GetOk("token_time_to_live_seconds"); ok {
		oauth.TokenTimeToLiveSeconds = Int(v.(int))
	}

	if v, ok := d.GetOk("use_proxy"); ok {
		oauth.UseProxy = Bool(v.(bool))
	}

	return oauth
}
