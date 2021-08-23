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

		Schema: resourcePingAccessWebSessionSchema(),
	}
}

func resourcePingAccessWebSessionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"audience": {
			Type:             schema.TypeString,
			Required:         true,
			ValidateDiagFunc: validateAudience,
		},
		"cache_user_attributes": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"client_credentials": {
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem:     oAuthClientCredentialsResource(),
		},
		"cookie_domain": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"cookie_type": {
			Type:             schema.TypeString,
			Optional:         true,
			ValidateDiagFunc: validateCookieType,
			Default:          "Encrypted",
		},
		"enable_refresh_user": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"http_only_cookie": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"idle_timeout_in_minutes": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  60,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"oidc_login_type": {
			Type:             schema.TypeString,
			Optional:         true,
			Default:          "Code",
			ValidateDiagFunc: validateOidcLoginType,
		},
		"pkce_challenge_type": {
			Type:             schema.TypeString,
			Optional:         true,
			Default:          "OFF",
			ValidateDiagFunc: validatePkceChallengeType,
		},
		"pfsession_state_cache_in_seconds": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  60,
		},
		"refresh_user_info_claims_interval": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  60,
		},
		"request_preservation_type": {
			Type:             schema.TypeString,
			Optional:         true,
			Default:          "POST",
			ValidateDiagFunc: validateRequestPreservationType,
		},
		"request_profile": {
			Type:       schema.TypeBool,
			Optional:   true,
			Default:    true,
			Deprecated: "DEPRECATED - to be removed in a future release; please use 'scopes' instead",
		},
		"same_site": {
			Type:             schema.TypeString,
			Optional:         true,
			Default:          "None",
			ValidateDiagFunc: validateWebSessionSameSite,
		},
		"scopes": {
			Type:     schema.TypeSet,
			Optional: true,
			DefaultFunc: func() (interface{}, error) {
				return []interface{}{"profile", "email", "address", "phone"}, nil
			},
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"secure_cookie": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"send_requested_url_to_provider": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"session_timeout_in_minutes": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  480,
		},
		"validate_session_is_alive": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"web_storage_type": {
			Type:             schema.TypeString,
			Optional:         true,
			Default:          "SessionStorage",
			ValidateDiagFunc: validateWebStorageType,
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
	setResourceDataStringWithDiagnostic(d, "pkce_challenge_type", input.PkceChallengeType, &diags)
	setResourceDataIntWithDiagnostic(d, "pfsession_state_cache_in_seconds", input.PfsessionStateCacheInSeconds, &diags)
	setResourceDataIntWithDiagnostic(d, "refresh_user_info_claims_interval", input.RefreshUserInfoClaimsInterval, &diags)
	setResourceDataStringWithDiagnostic(d, "request_preservation_type", input.RequestPreservationType, &diags)
	setResourceDataBoolWithDiagnostic(d, "request_profile", input.RequestProfile, &diags)
	setResourceDataStringWithDiagnostic(d, "same_site", input.SameSite, &diags)
	setResourceDataBoolWithDiagnostic(d, "secure_cookie", input.SecureCookie, &diags)
	setResourceDataBoolWithDiagnostic(d, "send_requested_url_to_provider", input.SendRequestedUrlToProvider, &diags)
	setResourceDataIntWithDiagnostic(d, "session_timeout_in_minutes", input.SessionTimeoutInMinutes, &diags)
	setResourceDataBoolWithDiagnostic(d, "validate_session_is_alive", input.ValidateSessionIsAlive, &diags)

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
