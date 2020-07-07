package pingaccess

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourcePingAccessPingFederateRuntimeMetadata() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourcePingAccessPingFederateRuntimeMetadataRead,
		Schema:      dataSourcePingAccessPingFederateRuntimeMetadataSchema(),
	}
}

func dataSourcePingAccessPingFederateRuntimeMetadataSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"authorization_endpoint": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"backchannel_authentication_endpoint": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"claim_types_supported": computedListOfString(),
		"claims_parameter_supported": {
			Type:     schema.TypeBool,
			Computed: true,
		},
		"claims_supported":                 computedListOfString(),
		"code_challenge_methods_supported": computedListOfString(),
		"end_session_endpoint": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"grant_types_supported":                 computedListOfString(),
		"id_token_signing_alg_values_supported": computedListOfString(),
		"introspection_endpoint": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"issuer": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"jwks_uri": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"ping_end_session_endpoint": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"ping_revoked_sris_endpoint": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"request_object_signing_alg_values_supported": computedListOfString(),
		"request_parameter_supported": {
			Type:     schema.TypeBool,
			Computed: true,
		},
		"request_uri_parameter_supported": {
			Type:     schema.TypeBool,
			Computed: true,
		},
		"response_modes_supported": computedListOfString(),
		"response_types_supported": computedListOfString(),
		"revocation_endpoint": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"scopes_supported":        computedListOfString(),
		"subject_types_supported": computedListOfString(),
		"token_endpoint": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"token_endpoint_auth_methods_supported": computedListOfString(),
		"userinfo_endpoint": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"userinfo_signing_alg_values_supported": computedListOfString(),
	}
}

func dataSourcePingAccessPingFederateRuntimeMetadataRead(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(paClient).Pingfederate
	result, resp, err := svc.GetPingFederateMetadataCommand()
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
