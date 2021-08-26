package sdkv2provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourcePingAccessPingFederateRuntimeMetadata() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourcePingAccessPingFederateRuntimeMetadataRead,
		Schema:      dataSourcePingAccessPingFederateRuntimeMetadataSchema(),
		Description: "Use this data source to get the PingFederate Runtime metadata.",
	}
}

func dataSourcePingAccessPingFederateRuntimeMetadataSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"authorization_endpoint": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "URL of the OpenID Connect provider's authorization endpoint.",
		},
		"backchannel_authentication_endpoint": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The endpoint used to initiate an out-of-band authentication.",
		},
		"claim_types_supported": {
			Type:        schema.TypeList,
			Computed:    true,
			Description: "JSON array containing a list of the claim types that the OpenID Connect provider supports.",
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"claims_parameter_supported": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Boolean value specifying whether the OpenID Connect provider supports use of the claims parameter, with true indicating support.",
		},
		"claims_supported": {
			Type:        schema.TypeList,
			Computed:    true,
			Description: "JSON array containing a list of the claim names of the claims that the OpenID Connect provider MAY be able to supply values for.",
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"code_challenge_methods_supported": {
			Type:        schema.TypeList,
			Computed:    true,
			Description: "Proof Key for Code Exchange (PKCE) code challenge methods supported by this OpenID Connect provider.",
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"end_session_endpoint": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "URL at the OpenID Connect provider to which a relying party can perform a redirect to request that the end-user be logged out at the OpenID Connect provider.",
		},
		"grant_types_supported": {
			Type:        schema.TypeList,
			Computed:    true,
			Description: "JSON array containing a list of the OAuth 2.0 grant type values that this OpenID Connect provider supports.",
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"id_token_signing_alg_values_supported": {
			Type:        schema.TypeList,
			Computed:    true,
			Description: "JSON array containing a list of the JWS signing algorithms supported by the OpenID Connect provider for the id token to encode the claims in a JWT.",
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"introspection_endpoint": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "URL of the OpenID Connect provider's OAuth 2.0 introspection endpoint.",
		},
		"issuer": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "OpenID Connect provider's issuer identifier URL.",
		},
		"jwks_uri": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "URL of the OpenID Connect provider's JWK Set document.",
		},
		"ping_end_session_endpoint": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "PingFederate logout endpoint. (Not applicable if PingFederate is not the OpenID Connect provider)",
		},
		"ping_revoked_sris_endpoint": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "PingFederate session revocation endpoint. (Not applicable if PingFederate is not the OpenID Connect provider)",
		},
		"request_object_signing_alg_values_supported": {
			Type:        schema.TypeList,
			Computed:    true,
			Description: "JSON array containing a list of the JWS signing algorithms supported by the OpenID Connect provider for request objects.",
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"request_parameter_supported": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Boolean value specifying whether the OpenID Connect provider supports use of the request parameter, with true indicating support.",
		},
		"request_uri_parameter_supported": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Boolean value specifying whether the OpenID Connect provider supports use of the request_uri parameter, with true indicating support.",
		},
		"response_modes_supported": {
			Type:        schema.TypeList,
			Computed:    true,
			Description: "JSON array containing a list of the OAuth 2.0 \"response_mode\" values that this OpenID Connect provider supports.",
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"response_types_supported": {
			Type:        schema.TypeList,
			Computed:    true,
			Description: "JSON array containing a list of the OAuth 2.0 \"response_type\" values that this OpenID Connect provider supports.",
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"revocation_endpoint": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "URL of the OpenID Connect provider's OAuth 2.0 revocation endpoint.",
		},
		"scopes_supported": {
			Type:        schema.TypeList,
			Computed:    true,
			Description: "JSON array containing a list of the OAuth 2.0 \"scope\" values that this OpenID Connect provider supports.",
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"subject_types_supported": {
			Type:        schema.TypeList,
			Computed:    true,
			Description: "JSON array containing a list of the Subject Identifier types that this OpenID Connect provider supports.",
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"token_endpoint": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "URL of the OpenID Connect provider's token endpoint.",
		},
		"token_endpoint_auth_methods_supported": {
			Type:        schema.TypeList,
			Computed:    true,
			Description: "JSON array containing a list of client authentication methods supported by this token endpoint.",
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"userinfo_endpoint": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "URL of the OpenID Connect provider's userInfo endpoint.",
		},
		"userinfo_signing_alg_values_supported": {
			Type:        schema.TypeList,
			Computed:    true,
			Description: "JSON array containing a list of the JWS signing algorithms supported by the userInfo endpoint to encode the claims in a JWT.",
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
	}
}

func dataSourcePingAccessPingFederateRuntimeMetadataRead(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(paClient).Pingfederate
	result, resp, err := svc.GetPingFederateRuntimeMetadataCommand()
	if err != nil {
		return diag.Errorf("unable to read PingFederate Runtime Metadata: %s\n%v", err, resp)
	}
	var diags diag.Diagnostics
	d.SetId("pingfederate_runtime_metadata")
	setResourceDataStringWithDiagnostic(d, "authorization_endpoint", result.AuthorizationEndpoint, &diags)
	setResourceDataStringWithDiagnostic(d, "backchannel_authentication_endpoint", result.BackchannelAuthenticationEndpoint, &diags)
	if err := d.Set("claim_types_supported", result.ClaimTypesSupported); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	if err := d.Set("claims_parameter_supported", result.ClaimsParameterSupported); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	if err := d.Set("claims_supported", result.ClaimsSupported); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	if err := d.Set("code_challenge_methods_supported", result.CodeChallengeMethodsSupported); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	setResourceDataStringWithDiagnostic(d, "end_session_endpoint", result.EndSessionEndpoint, &diags)
	if err := d.Set("grant_types_supported", result.GrantTypesSupported); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	if err := d.Set("id_token_signing_alg_values_supported", result.IdTokenSigningAlgValuesSupported); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	setResourceDataStringWithDiagnostic(d, "introspection_endpoint", result.IntrospectionEndpoint, &diags)
	setResourceDataStringWithDiagnostic(d, "issuer", result.Issuer, &diags)
	setResourceDataStringWithDiagnostic(d, "jwks_uri", result.JwksUri, &diags)
	setResourceDataStringWithDiagnostic(d, "ping_end_session_endpoint", result.PingEndSessionEndpoint, &diags)
	setResourceDataStringWithDiagnostic(d, "ping_revoked_sris_endpoint", result.PingRevokedSrisEndpoint, &diags)
	if err := d.Set("request_object_signing_alg_values_supported", result.RequestObjectSigningAlgValuesSupported); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	setResourceDataBoolWithDiagnostic(d, "request_parameter_supported", result.RequestParameterSupported, &diags)
	setResourceDataBoolWithDiagnostic(d, "request_uri_parameter_supported", result.RequestUriParameterSupported, &diags)
	if err := d.Set("response_modes_supported", result.ResponseModesSupported); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	if err := d.Set("response_types_supported", result.ResponseTypesSupported); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	setResourceDataStringWithDiagnostic(d, "revocation_endpoint", result.RevocationEndpoint, &diags)
	if err := d.Set("scopes_supported", result.ScopesSupported); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	if err := d.Set("subject_types_supported", result.SubjectTypesSupported); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	setResourceDataStringWithDiagnostic(d, "token_endpoint", result.TokenEndpoint, &diags)
	if err := d.Set("token_endpoint_auth_methods_supported", result.TokenEndpointAuthMethodsSupported); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	setResourceDataStringWithDiagnostic(d, "userinfo_endpoint", result.UserinfoEndpoint, &diags)
	if err := d.Set("userinfo_signing_alg_values_supported", result.UserinfoSigningAlgValuesSupported); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	return diags
}
