package sdkv2provider

import (
	"context"

	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess/models"
	"github.com/iwarapter/pingaccess-sdk-go/v62/services/oauth"

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
		Description: `Manages the PingAccess Third Party OAuth Server configuration.

-> This resource manages a singleton within PingAccess and as such you should ONLY ever declare one of this resource type. Deleting this resource resets the Third Party OAuth Server configuration to default values.`,
	}
}

func resourcePingAccessOAuthServerSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"audit_level": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "ON",
			Description: "Enable to record requests to third-party OAuth 2.0 Authorization Server to the audit store.",
		},
		"cache_tokens": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Enable to retain token details for subsequent requests.",
		},
		"client_credentials": {
			Type:        schema.TypeList,
			Required:    true,
			MaxItems:    1,
			Description: "Specify the client credentials.",
			Elem:        oAuthClientCredentialsResource(),
		},
		"description": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The description of the third-party OAuth 2.0 Authorization Server.",
		},
		"introspection_endpoint": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The third-party OAuth 2.0 Authorization Server's token introspection endpoint.",
		},

		"secure": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Enable if third-party OAuth 2.0 Authorization Server is expecting HTTPS connections.",
		},
		"send_audience": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Enable to send the URI the user requested as the 'aud' OAuth parameter for PingAccess to the OAuth 2.0 Authorization server.",
		},
		"subject_attribute_name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The attribute you want to use from the OAuth access token as the subject for auditing purposes.",
		},
		"targets": {
			Type:        schema.TypeSet,
			Required:    true,
			MinItems:    1,
			Description: "One or more server hostname:port pairs used to access third-party OAuth 2.0 Authorization Server.",
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"token_time_to_live_seconds": {
			Type:        schema.TypeInt,
			Optional:    true,
			Default:     -1,
			Description: "Defines the number of seconds to cache the access token. -1 means no limit. This value should be less than the PingFederate Token Lifetime.",
		},
		"trusted_certificate_group_id": {
			Type:        schema.TypeInt,
			Required:    true,
			Description: "The group of certificates to use when authenticating to third-party OAuth 2.0 Authorization Server.",
		},
		"use_proxy": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "True if a proxy should be used for HTTP or HTTPS requests.",
		},
	}
}

func resourcePingAccessOAuthServerCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	d.SetId("oauth_server_settings")
	return resourcePingAccessOAuthServerUpdate(ctx, d, m)
}

func resourcePingAccessOAuthServerRead(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(paClient).Oauth
	result, _, err := svc.GetAuthorizationServerCommand()
	if err != nil {
		return diag.Errorf("unable to read OAuthServerSettings: %s", err)
	}
	return resourcePingAccessOAuthServerReadResult(d, result, m.(paClient).CanMaskPasswords())
}

func resourcePingAccessOAuthServerUpdate(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(paClient).Oauth
	input := oauth.UpdateAuthorizationServerCommandInput{
		Body: *resourcePingAccessOAuthServerReadData(d),
	}
	result, _, err := svc.UpdateAuthorizationServerCommand(&input)
	if err != nil {
		return diag.Errorf("unable to update OAuthServerSettings: %s", err)
	}

	d.SetId("oauth_server_settings")
	return resourcePingAccessOAuthServerReadResult(d, result, false)
}

func resourcePingAccessOAuthServerDelete(_ context.Context, _ *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(paClient).Oauth
	_, err := svc.DeleteAuthorizationServerCommand()
	if err != nil {
		return diag.Errorf("unable to reset OAuthServerSettings: %s", err)
	}
	return nil
}

func resourcePingAccessOAuthServerReadResult(d *schema.ResourceData, input *models.AuthorizationServerView, trackPasswords bool) diag.Diagnostics {
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
		diags = append(diags, diag.FromErr(err)...)
	}
	if input.ClientCredentials != nil && input.ClientCredentials.ClientSecret != nil {
		setClientCredentials(d, input.ClientCredentials, trackPasswords, &diags)
	}

	return diags
}

func resourcePingAccessOAuthServerReadData(d *schema.ResourceData) *models.AuthorizationServerView {
	targets := expandStringList(d.Get("targets").(*schema.Set).List())
	oauth := &models.AuthorizationServerView{
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
