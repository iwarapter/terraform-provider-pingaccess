package sdkv2provider

import (
	"context"

	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess/models"
	"github.com/iwarapter/pingaccess-sdk-go/v62/services/webSessions"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePingAccessWebSession() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePingAccessWebSessionCreate,
		ReadContext:   resourcePingAccessWebSessionRead,
		UpdateContext: resourcePingAccessWebSessionUpdate,
		DeleteContext: resourcePingAccessWebSessionDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema:      resourcePingAccessWebSessionSchema(),
		Description: `Provides configuration for Web Sessions within PingAccess.`,
	}
}

func resourcePingAccessWebSessionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"audience": {
			Type:             schema.TypeString,
			Required:         true,
			ValidateDiagFunc: validateAudience,
			Description:      "Enter a unique identifier between 1 and 32 characters that defines who the PA Token is applicable to.",
		},
		"cache_user_attributes": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
			Description: "Specify if PingAccess should cache user attribute information for use in policy decisions. When disabled, this data is encoded and stored in the session cookie.",
		},
		"client_credentials": {
			Type:        schema.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "Specify the client credentials.",
			Elem:        oAuthClientCredentialsResource(),
		},
		"cookie_domain": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The domain where the cookie is stored--for example, corp.yourcompany.com.",
		},
		"cookie_type": {
			Type:             schema.TypeString,
			Optional:         true,
			ValidateDiagFunc: validateCookieType,
			Default:          "Encrypted",
			Description:      "Specify an Encrypted JWT or a Signed JWT web session cookie. Default is Encrypted.",
		},
		"enable_refresh_user": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     true,
			Description: "Specify if you want to have PingAccess periodically refresh user data from PingFederate for use in policy decisions.",
		},
		"http_only_cookie": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     true,
			Description: "Enable the HttpOnly flag on cookies that contain the PA Token.",
		},
		"idle_timeout_in_minutes": {
			Type:        schema.TypeInt,
			Optional:    true,
			Default:     60,
			Description: "The length of time you want the PingAccess Token to remain active when no activity is detected.",
		},
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The name of the web session.",
		},
		"oidc_login_type": {
			Type:             schema.TypeString,
			Optional:         true,
			Default:          "Code",
			ValidateDiagFunc: validateOidcLoginType,
			Description:      "The web session token type.",
		},
		"pkce_challenge_type": {
			Type:             schema.TypeString,
			Optional:         true,
			Default:          "OFF",
			ValidateDiagFunc: validatePkceChallengeType,
			Description:      "Specify the code_challenge_method to use for PKCE during the Code login flow. OFF signifies to not use PKCE.",
		},
		"pfsession_state_cache_in_seconds": {
			Type:        schema.TypeInt,
			Optional:    true,
			Default:     60,
			Description: "Specify the number of seconds to cache PingFederate Session State information.",
		},
		"refresh_user_info_claims_interval": {
			Type:        schema.TypeInt,
			Optional:    true,
			Default:     60,
			Description: "Specify the maximum number of seconds to cache user attribute information when the Refresh User is enabled.",
		},
		"request_preservation_type": {
			Type:             schema.TypeString,
			Optional:         true,
			Default:          "POST",
			ValidateDiagFunc: validateRequestPreservationType,
			Description:      "Specify the types of request data to be preserved if the user is redirected to an authentication page when submitting information to a protected resource.",
		},
		"request_profile": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     true,
			Description: "Specifies whether the default scopes ('profile', 'email', 'address', and 'phone') should be specified in the access request.",
			Deprecated:  "DEPRECATED - to be removed in a future release; please use 'scopes' instead",
		},
		"same_site": {
			Type:             schema.TypeString,
			Optional:         true,
			Default:          "None",
			ValidateDiagFunc: validateWebSessionSameSite,
			Description:      "Specify the SameSite attribute to be used when setting the PingAccess Cookie. Default is None which allows the cookie to be used in a third-party context. If the cookie is not used in a third-party context then Lax is recommended.",
		},
		"scopes": {
			Type:     schema.TypeSet,
			Optional: true,
			DefaultFunc: func() (interface{}, error) {
				return []interface{}{"profile", "email", "address", "phone"}, nil
			},
			Description: "The list of scopes to be specified in the access request. If not specified, the default scopes ('profile', 'email', 'address', and 'phone') will be used. The openid scope is implied and does not need to be specified in this list.",
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"secure_cookie": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     true,
			Description: "Specify whether the PingAccess Cookie must be sent using only HTTPS connections.",
		},
		"send_requested_url_to_provider": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     true,
			Description: "Specify if you want to send the requested URL as part of the authentication request to the OpenID Connect Provider.",
		},
		"session_timeout_in_minutes": {
			Type:        schema.TypeInt,
			Optional:    true,
			Default:     480,
			Description: "The length of time you want the PA Token to remain active. Once the PA Token expires, an authenticated user must re-authenticate.",
		},
		"validate_session_is_alive": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
			Description: "Specify if PingAccess should validate sessions with the configured PingFederate instance during request processing.",
		},
		"web_storage_type": {
			Type:             schema.TypeString,
			Optional:         true,
			Default:          "SessionStorage",
			ValidateDiagFunc: validateWebStorageType,
			Description:      "Specify the type of web storage to use for request preservation data, either `SessionStorage` or `LocalStorage`. Default is SessionStorage.",
		},
	}
}

func resourcePingAccessWebSessionCreate(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(paClient).WebSessions
	input := webSessions.AddWebSessionCommandInput{
		Body: *resourcePingAccessWebSessionReadData(d),
	}

	result, _, err := svc.AddWebSessionCommand(&input)
	if err != nil {
		return diag.Errorf("unable to create WebSession: %s", err)
	}

	d.SetId(result.Id.String())
	return resourcePingAccessWebSessionReadResult(d, result, m.(paClient).CanMaskPasswords())
}

func resourcePingAccessWebSessionRead(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(paClient).WebSessions
	input := &webSessions.GetWebSessionCommandInput{
		Id: d.Id(),
	}
	result, _, err := svc.GetWebSessionCommand(input)
	if err != nil {
		return diag.Errorf("unable to read WebSession: %s", err)
	}
	return resourcePingAccessWebSessionReadResult(d, result, m.(paClient).CanMaskPasswords())
}

func resourcePingAccessWebSessionUpdate(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(paClient).WebSessions
	input := webSessions.UpdateWebSessionCommandInput{
		Body: *resourcePingAccessWebSessionReadData(d),
		Id:   d.Id(),
	}
	result, _, err := svc.UpdateWebSessionCommand(&input)
	if err != nil {
		return diag.Errorf("unable to update WebSession: %s", err)
	}
	return resourcePingAccessWebSessionReadResult(d, result, false)
}

func resourcePingAccessWebSessionDelete(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(paClient).WebSessions
	input := &webSessions.DeleteWebSessionCommandInput{
		Id: d.Id(),
	}

	_, err := svc.DeleteWebSessionCommand(input)
	if err != nil {
		return diag.Errorf("unable to delete WebSession: %s", err)
	}
	return nil
}

func resourcePingAccessWebSessionReadResult(d *schema.ResourceData, input *models.WebSessionView, trackPasswords bool) diag.Diagnostics {
	var diags diag.Diagnostics
	setResourceDataStringWithDiagnostic(d, "audience", input.Audience, &diags)
	setResourceDataStringWithDiagnostic(d, "name", input.Name, &diags)
	setResourceDataStringWithDiagnostic(d, "web_storage_type", input.WebStorageType, &diags)
	setResourceDataBoolWithDiagnostic(d, "cache_user_attributes", input.CacheUserAttributes, &diags)
	setResourceDataStringWithDiagnostic(d, "cookie_domain", input.CookieDomain, &diags)
	setResourceDataStringWithDiagnostic(d, "cookie_type", input.CookieType, &diags)
	setResourceDataBoolWithDiagnostic(d, "enable_refresh_user", input.EnableRefreshUser, &diags)
	setResourceDataBoolWithDiagnostic(d, "http_only_cookie", input.HttpOnlyCookie, &diags)
	setResourceDataIntWithDiagnostic(d, "idle_timeout_in_minutes", input.IdleTimeoutInMinutes, &diags)
	setResourceDataStringWithDiagnostic(d, "oidc_login_type", input.OidcLoginType, &diags)
	setResourceDataIntWithDiagnostic(d, "pfsession_state_cache_in_seconds", input.PfsessionStateCacheInSeconds, &diags)
	setResourceDataIntWithDiagnostic(d, "refresh_user_info_claims_interval", input.RefreshUserInfoClaimsInterval, &diags)
	setResourceDataStringWithDiagnostic(d, "request_preservation_type", input.RequestPreservationType, &diags)
	setResourceDataBoolWithDiagnostic(d, "request_profile", input.RequestProfile, &diags)
	setResourceDataStringWithDiagnostic(d, "same_site", input.SameSite, &diags)
	setResourceDataBoolWithDiagnostic(d, "secure_cookie", input.SecureCookie, &diags)
	setResourceDataBoolWithDiagnostic(d, "send_requested_url_to_provider", input.SendRequestedUrlToProvider, &diags)
	setResourceDataIntWithDiagnostic(d, "session_timeout_in_minutes", input.SessionTimeoutInMinutes, &diags)
	setResourceDataBoolWithDiagnostic(d, "validate_session_is_alive", input.ValidateSessionIsAlive, &diags)

	// Default is off however this field is not supported before PA 6.0 so we set to OFF to match schema default for 5.3 and override if provided.
	setResourceDataStringWithDiagnostic(d, "pkce_challenge_type", String("OFF"), &diags)
	setResourceDataStringWithDiagnostic(d, "pkce_challenge_type", input.PkceChallengeType, &diags)

	if input.Scopes != nil {
		if err := d.Set("scopes", *input.Scopes); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}
	if input.ClientCredentials != nil && input.ClientCredentials.ClientId != nil {
		setClientCredentials(d, input.ClientCredentials, trackPasswords, &diags)
	}
	return diags
}

func resourcePingAccessWebSessionReadData(d *schema.ResourceData) *models.WebSessionView {
	websession := &models.WebSessionView{
		Audience:          String(d.Get("audience").(string)),
		Name:              String(d.Get("name").(string)),
		ClientCredentials: expandOAuthClientCredentialsView(d.Get("client_credentials").([]interface{})),
	}

	websession.CacheUserAttributes = Bool(d.Get("cache_user_attributes").(bool))

	if v, ok := d.GetOk("cookie_domain"); ok && v.(string) != "" {
		websession.CookieDomain = String(v.(string))
	}

	if v, ok := d.GetOk("cookie_type"); ok {
		websession.CookieType = String(v.(string))
	}

	websession.EnableRefreshUser = Bool(d.Get("enable_refresh_user").(bool))
	websession.HttpOnlyCookie = Bool(d.Get("http_only_cookie").(bool))
	websession.IdleTimeoutInMinutes = Int(d.Get("idle_timeout_in_minutes").(int))

	if v, ok := d.GetOk("oidc_login_type"); ok {
		websession.OidcLoginType = String(v.(string))
	}

	if v, ok := d.GetOk("pkce_challenge_type"); ok {
		websession.PkceChallengeType = String(v.(string))
	}

	websession.PfsessionStateCacheInSeconds = Int(d.Get("pfsession_state_cache_in_seconds").(int))
	websession.RefreshUserInfoClaimsInterval = Int(d.Get("refresh_user_info_claims_interval").(int))
	websession.RequestPreservationType = String(d.Get("request_preservation_type").(string))
	websession.RequestProfile = Bool(d.Get("request_profile").(bool))

	if v, ok := d.GetOk("same_site"); ok {
		websession.SameSite = String(v.(string))
	}

	scopes := expandStringList(d.Get("scopes").(*schema.Set).List())
	websession.Scopes = &scopes
	websession.SecureCookie = Bool(d.Get("secure_cookie").(bool))
	websession.SendRequestedUrlToProvider = Bool(d.Get("send_requested_url_to_provider").(bool))
	websession.SessionTimeoutInMinutes = Int(d.Get("session_timeout_in_minutes").(int))
	websession.ValidateSessionIsAlive = Bool(d.Get("validate_session_is_alive").(bool))
	websession.WebStorageType = String(d.Get("web_storage_type").(string))

	return websession
}
