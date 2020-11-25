package sdkv2provider

import (
	"context"

	"github.com/iwarapter/pingaccess-sdk-go/pingaccess/models"
	"github.com/iwarapter/pingaccess-sdk-go/services/pingfederate"

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
	}
}

func resourcePingAccessPingFederateOAuthSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"access_validator_id": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  1,
		},
		"cache_tokens": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"client_id": {
			Type:     schema.TypeString,
			Required: true,
		},
		"client_secret": hiddenField(),
		"name": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "PingFederate",
		},
		"send_audience": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"subject_attribute_name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"token_time_to_live_seconds": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  -1,
		},
		"use_token_introspection": {
			Type:     schema.TypeBool,
			Optional: true,
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

	return resourcePingAccessPingFederateOAuthReadResult(d, result)
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
	return resourcePingAccessPingFederateOAuthReadResult(d, result)
}

func resourcePingAccessPingFederateOAuthDelete(_ context.Context, _ *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(paClient).Pingfederate
	_, err := svc.DeletePingFederateAccessTokensCommand()
	if err != nil {
		return diag.Errorf("unable to reset PingFederateOAuth: %s", err)
	}
	return nil
}

func resourcePingAccessPingFederateOAuthReadResult(d *schema.ResourceData, input *models.PingFederateAccessTokenView) diag.Diagnostics {
	var diags diag.Diagnostics
	setResourceDataIntWithDiagnostic(d, "access_validator_id", input.AccessValidatorId, &diags)
	setResourceDataBoolWithDiagnostic(d, "cache_tokens", input.CacheTokens, &diags)
	setResourceDataStringWithDiagnostic(d, "client_id", input.ClientId, &diags)
	setResourceDataStringWithDiagnostic(d, "name", input.Name, &diags)
	setResourceDataBoolWithDiagnostic(d, "send_audience", input.SendAudience, &diags)
	setResourceDataStringWithDiagnostic(d, "subject_attribute_name", input.SubjectAttributeName, &diags)
	setResourceDataIntWithDiagnostic(d, "token_time_to_live_seconds", input.TokenTimeToLiveSeconds, &diags)
	setResourceDataBoolWithDiagnostic(d, "use_token_introspection", input.UseTokenIntrospection, &diags)

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

	return diags
}

func resourcePingAccessPingFederateOAuthReadData(d *schema.ResourceData) *models.PingFederateAccessTokenView {
	oauth := &models.PingFederateAccessTokenView{
		ClientId:             String(d.Get("client_id").(string)),
		SubjectAttributeName: String(d.Get("subject_attribute_name").(string)),
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
