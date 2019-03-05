package pingaccess

import (
	"crypto/tls"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform/helper/schema"
	pa "github.com/iwarapter/pingaccess-sdk-go/pingaccess"
)

func resourcePingAccessWebSession() *schema.Resource {
	return &schema.Resource{
		Create: resourcePingAccessWebSessionCreate,
		Read:   resourcePingAccessWebSessionRead,
		Update: resourcePingAccessWebSessionUpdate,
		Delete: resourcePingAccessWebSessionDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: resourcePingAccessWebSessionSchema(),
	}
}

func resourcePingAccessWebSessionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"audience": &schema.Schema{
			Type:         schema.TypeString,
			Required:     true,
			ValidateFunc: validateAudience,
		},
		"cache_user_attributes": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"client_credentials": oAuthClientCredentials(),
		"cookie_domain": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"cookie_type": &schema.Schema{
			Type:         schema.TypeString,
			Optional:     true,
			ValidateFunc: validateCookieType,
			Default:      "Signed",
		},
		"enable_refresh_user": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"http_only_cookie": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"idle_timeout_in_minutes": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  60,
		},
		"name": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"oidc_login_type": &schema.Schema{
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "Code",
			ValidateFunc: validateOidcLoginType,
		},
		"pfsession_state_cache_in_seconds": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  60,
		},
		"refresh_user_info_claims_interval": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  60,
		},
		"request_preservation_type": &schema.Schema{
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "POST",
			ValidateFunc: validateRequestPreservationType,
		},
		"request_profile": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"scopes": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			DefaultFunc: func() (interface{}, error) {
				return []interface{}{"profile", "email", "address", "phone"}, nil
			},
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"secure_cookie": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"send_requested_url_to_provider": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"session_timeout_in_minutes": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  480,
		},
		"validate_session_is_alive": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"web_storage_type": &schema.Schema{
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "SessionStorage",
			ValidateFunc: validateWebStorageType,
		},
	}
}

func scopesDefaultFunc() schema.SchemaDefaultFunc {
	return func() (interface{}, error) {

		// d := schema.NewSet(stringHash, []interface{}{"profile", "email", "address", "phone"})
		// log.Printf("%v", d)
		// log.Printf("%v", d.List())
		// return d.List(), nil
		return []interface{}{"profile", "email", "address", "phone"}, nil
	}
}

func resourcePingAccessWebSessionCreate(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pa.Client).WebSessions
	input := pa.AddWebSessionCommandInput{
		Body: *resourcePingAccessWebSessionReadData(d),
	}

	result, _, err := svc.AddWebSessionCommand(&input)
	if err != nil {
		return fmt.Errorf("Error creating websession: %s", err)
	}

	d.SetId(result.Id.String())
	return resourcePingAccessWebSessionReadResult(d, result)
}

func resourcePingAccessWebSessionRead(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pa.Client).WebSessions
	input := &pa.GetWebSessionCommandInput{
		Id: d.Id(),
	}
	result, _, _ := svc.GetWebSessionCommand(input)
	return resourcePingAccessWebSessionReadResult(d, result)
}

func resourcePingAccessWebSessionUpdate(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pa.Client).WebSessions
	input := pa.UpdateWebSessionCommandInput{
		Body: *resourcePingAccessWebSessionReadData(d),
		Id:   d.Id(),
	}
	result, _, err := svc.UpdateWebSessionCommand(&input)
	if err != nil {
		return fmt.Errorf("Error updating websession: %s", err)
	}
	return resourcePingAccessWebSessionReadResult(d, result)
}

func resourcePingAccessWebSessionDelete(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pa.Client).WebSessions
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	input := &pa.DeleteWebSessionCommandInput{
		Id: d.Id(),
	}

	_, err := svc.DeleteWebSessionCommand(input)
	if err != nil {
		return fmt.Errorf("Error deleting websession: %s", err)
	}
	return nil
}

func resourcePingAccessWebSessionReadResult(d *schema.ResourceData, input *pa.WebSessionView) (err error) {
	setResourceDataString(d, "audience", input.Audience)
	setResourceDataString(d, "name", input.Name)
	setResourceDataString(d, "web_storage_type", input.WebStorageType)
	setResourceDataBool(d, "cache_user_attributes", input.CacheUserAttributes)
	setResourceDataString(d, "cookie_domain", input.CookieDomain)
	setResourceDataString(d, "cookie_type", input.CookieType)
	setResourceDataBool(d, "enable_refresh_user", input.EnableRefreshUser)
	setResourceDataBool(d, "http_only_cookie", input.HttpOnlyCookie)
	setResourceDataInt(d, "idle_timeout_in_minutes", input.IdleTimeoutInMinutes)
	setResourceDataString(d, "oidc_login_type", input.OidcLoginType)
	setResourceDataInt(d, "pfsession_state_cache_in_seconds", input.PfsessionStateCacheInSeconds)
	setResourceDataInt(d, "refresh_user_info_claims_interval", input.RefreshUserInfoClaimsInterval)
	setResourceDataString(d, "request_preservation_type", input.RequestPreservationType)
	setResourceDataBool(d, "request_profile", input.RequestProfile)
	setResourceDataBool(d, "secure_cookie", input.SecureCookie)
	setResourceDataBool(d, "send_requested_url_to_provider", input.SendRequestedUrlToProvider)
	setResourceDataInt(d, "session_timeout_in_minutes", input.SessionTimeoutInMinutes)
	setResourceDataBool(d, "validate_session_is_alive", input.ValidateSessionIsAlive)

	if input.Scopes != nil {
		if err = d.Set("scopes", *input.Scopes); err != nil {
			return err
		}
	}
	if input.ClientCredentials != nil {
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

func resourcePingAccessWebSessionReadData(d *schema.ResourceData) *pa.WebSessionView {
	websession := &pa.WebSessionView{
		Audience:          String(d.Get("audience").(string)),
		Name:              String(d.Get("name").(string)),
		ClientCredentials: expandOAuthClientCredentialsView(d.Get("client_credentials").([]interface{})),
	}

	if _, ok := d.GetOkExists("cache_user_attributes"); ok {
		websession.CacheUserAttributes = Bool(d.Get("cache_user_attributes").(bool))
	}

	if _, ok := d.GetOkExists("cookie_domain"); ok {
		websession.CookieDomain = String(d.Get("cookie_domain").(string))
	}

	if _, ok := d.GetOkExists("cookie_type"); ok {
		websession.CookieType = String(d.Get("cookie_type").(string))
	}

	if _, ok := d.GetOkExists("enable_refresh_user"); ok {
		websession.EnableRefreshUser = Bool(d.Get("enable_refresh_user").(bool))
	}

	if _, ok := d.GetOkExists("http_only_cookie"); ok {
		websession.HttpOnlyCookie = Bool(d.Get("http_only_cookie").(bool))
	}

	if _, ok := d.GetOkExists("idle_timeout_in_minutes"); ok {
		websession.IdleTimeoutInMinutes = Int(d.Get("idle_timeout_in_minutes").(int))
	}

	if _, ok := d.GetOkExists("oidc_login_type"); ok {
		websession.OidcLoginType = String(d.Get("oidc_login_type").(string))
	}

	if _, ok := d.GetOkExists("pfsession_state_cache_in_seconds"); ok {
		websession.PfsessionStateCacheInSeconds = Int(d.Get("pfsession_state_cache_in_seconds").(int))
	}

	if _, ok := d.GetOkExists("refresh_user_info_claims_interval"); ok {
		websession.RefreshUserInfoClaimsInterval = Int(d.Get("refresh_user_info_claims_interval").(int))
	}

	if _, ok := d.GetOkExists("request_preservation_type"); ok {
		websession.RequestPreservationType = String(d.Get("request_preservation_type").(string))
	}

	if _, ok := d.GetOkExists("request_profile"); ok {
		websession.RequestProfile = Bool(d.Get("request_profile").(bool))
	}

	if _, ok := d.GetOkExists("scopes"); ok {
		scopes := expandStringList(d.Get("scopes").(*schema.Set).List())
		websession.Scopes = &scopes
	}

	if _, ok := d.GetOkExists("secure_cookie"); ok {
		websession.SecureCookie = Bool(d.Get("secure_cookie").(bool))
	}

	if _, ok := d.GetOkExists("send_requested_url_to_provider"); ok {
		websession.SendRequestedUrlToProvider = Bool(d.Get("send_requested_url_to_provider").(bool))
	}

	if _, ok := d.GetOkExists("session_timeout_in_minutes"); ok {
		websession.SessionTimeoutInMinutes = Int(d.Get("session_timeout_in_minutes").(int))
	}

	if _, ok := d.GetOkExists("validate_session_is_alive"); ok {
		websession.ValidateSessionIsAlive = Bool(d.Get("validate_session_is_alive").(bool))
	}

	if _, ok := d.GetOkExists("web_storage_type"); ok {
		websession.WebStorageType = String(d.Get("web_storage_type").(string))
	}

	return websession
}
