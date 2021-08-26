package sdkv2provider

import (
	"context"

	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess/models"
	"github.com/iwarapter/pingaccess-sdk-go/v62/services/pingfederate"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePingAccessPingFederateOAuth() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePingAccessPingFederateOAuthCreate,
		ReadContext:   resourcePingAccessPingFederateOAuthRead,
		UpdateContext: resourcePingAccessPingFederateOAuthUpdate,
		DeleteContext: resourcePingAccessPingFederateOAuthDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: resourcePingAccessPingFederateOAuthSchema(),
		Description: `Manages the PingFederate OAuth Client configuration.

-> This resource manages a singleton within PingAccess and as such you should ONLY ever declare one of this resource type. Deleting this resource resets the PingFederate OAuth Client configuration to default values.`,
	}
}

func resourcePingAccessPingFederateOAuthSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"access_validator_id": {
			Type:        schema.TypeInt,
			Optional:    true,
			Default:     1,
			Description: "The Access Validator Id. This field is read-only.",
		},
		"cache_tokens": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Enable to retain token details for subsequent requests.",
		},
		"client_credentials": {
			Type:          schema.TypeList,
			Optional:      true,
			MaxItems:      1,
			Description:   "Specify the credentials for the OAuth client configured in PingFederate.",
			Elem:          oAuthClientCredentialsResource(),
			ConflictsWith: []string{"client_id", "client_secret"},
		},
		"client_id": {
			Type:          schema.TypeString,
			Optional:      true,
			RequiredWith:  []string{"client_secret"},
			Description:   "The Client ID which PingAccess should use when requesting PingFederate to validate access tokens. The client must have Access Token Validation grant type allowed.",
			Deprecated:    "DEPRECATED - to be removed in a future release; please use 'client_credentials' instead.",
			ConflictsWith: []string{"client_credentials"},
		},
		"client_secret": {
			Type:          schema.TypeList,
			Optional:      true,
			MaxItems:      1,
			Description:   "The Client Secret for the Client ID.",
			Deprecated:    "DEPRECATED - to be removed in a future release; please use 'client_credentials' instead.",
			Elem:          hiddenFieldResource(),
			ConflictsWith: []string{"client_credentials"},
		},
		"name": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The unique Access Validator name.",
			Default:     "PingFederate",
		},
		"send_audience": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Enable to send the URI the user requested as the 'aud' OAuth parameter for PingAccess to use to select an Access Token Manager.",
		},
		"subject_attribute_name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The attribute you want to use from the OAuth access token as the subject for auditing purposes.",
		},
		"token_time_to_live_seconds": {
			Type:        schema.TypeInt,
			Optional:    true,
			Default:     -1,
			Description: "Defines the number of seconds to cache the access token. -1 means no limit. This value should be less than the PingFederate Token Lifetime.",
		},
		"use_token_introspection": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Specify if token introspection is enabled.",
		},
	}
}

func resourcePingAccessPingFederateOAuthCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	d.SetId("pingfederate_oauth_settings")
	return resourcePingAccessPingFederateOAuthUpdate(ctx, d, m)
}

func resourcePingAccessPingFederateOAuthRead(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(paClient).Pingfederate
	result, _, err := svc.GetPingFederateAccessTokensCommand()
	if err != nil {
		return diag.Errorf("unable to read PingFederateOAuth: %s", err)
	}

	return resourcePingAccessPingFederateOAuthReadResult(d, result, m.(paClient).CanMaskPasswords())
}

func resourcePingAccessPingFederateOAuthUpdate(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(paClient).Pingfederate
	input := pingfederate.UpdatePingFederateAccessTokensCommandInput{
		Body: *resourcePingAccessPingFederateOAuthReadData(d),
	}
	result, _, err := svc.UpdatePingFederateAccessTokensCommand(&input)
	if err != nil {
		return diag.Errorf("unable to update PingFederateOAuth: %s", err)
	}

	d.SetId("pingfederate_oauth_settings")
	return resourcePingAccessPingFederateOAuthReadResult(d, result, false)
}

func resourcePingAccessPingFederateOAuthDelete(_ context.Context, _ *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(paClient).Pingfederate
	_, err := svc.DeletePingFederateAccessTokensCommand()
	if err != nil {
		return diag.Errorf("unable to reset PingFederateOAuth: %s", err)
	}
	return nil
}

func resourcePingAccessPingFederateOAuthReadResult(d *schema.ResourceData, input *models.PingFederateAccessTokenView, trackPasswords bool) diag.Diagnostics {
	var diags diag.Diagnostics
	setResourceDataIntWithDiagnostic(d, "access_validator_id", input.AccessValidatorId, &diags)
	setResourceDataBoolWithDiagnostic(d, "cache_tokens", input.CacheTokens, &diags)
	setResourceDataStringWithDiagnostic(d, "name", input.Name, &diags)
	setResourceDataBoolWithDiagnostic(d, "send_audience", input.SendAudience, &diags)
	setResourceDataStringWithDiagnostic(d, "subject_attribute_name", input.SubjectAttributeName, &diags)
	setResourceDataIntWithDiagnostic(d, "token_time_to_live_seconds", input.TokenTimeToLiveSeconds, &diags)
	setResourceDataBoolWithDiagnostic(d, "use_token_introspection", input.UseTokenIntrospection, &diags)

	if input.ClientCredentials != nil && input.ClientCredentials.ClientId != nil {
		setClientCredentials(d, input.ClientCredentials, trackPasswords, &diags)
	} else {
		setResourceDataStringWithDiagnostic(d, "client_id", input.ClientId, &diags)
		if input.ClientSecret != nil {
			pw, ok := d.GetOk("client_secret.0.value")
			creds := flattenHiddenFieldView(input.ClientSecret)
			if ok {
				creds[0]["value"] = pw
			}
			if err := d.Set("client_secret", creds); err != nil {
				diags = append(diags, diag.FromErr(err)...)
			}
		}
	}

	return diags
}

func resourcePingAccessPingFederateOAuthReadData(d *schema.ResourceData) *models.PingFederateAccessTokenView {
	oauth := &models.PingFederateAccessTokenView{
		SubjectAttributeName: String(d.Get("subject_attribute_name").(string)),
	}

	if v, ok := d.GetOk("client_id"); ok {
		oauth.ClientId = String(v.(string))
	}

	if v, ok := d.GetOk("client_credentials"); ok {
		oauth.ClientCredentials = expandOAuthClientCredentialsView(v.([]interface{}))
	}

	if v, ok := d.GetOk("access_validator_id"); ok {
		oauth.AccessValidatorId = Int(v.(int))
	}

	if v, ok := d.GetOk("cache_tokens"); ok {
		oauth.CacheTokens = Bool(v.(bool))
	}

	if v, ok := d.GetOk("name"); ok {
		oauth.Name = String(v.(string))
	}

	if v, ok := d.GetOk("send_audience"); ok {
		oauth.SendAudience = Bool(v.(bool))
	}

	if v, ok := d.GetOk("token_time_to_live_seconds"); ok {
		oauth.TokenTimeToLiveSeconds = Int(v.(int))
	}

	if v, ok := d.GetOk("use_token_introspection"); ok {
		oauth.UseTokenIntrospection = Bool(v.(bool))
	}

	if v, ok := d.GetOk("client_secret"); ok {
		oauth.ClientSecret = expandHiddenFieldView(v.([]interface{}))
	}

	return oauth
}
