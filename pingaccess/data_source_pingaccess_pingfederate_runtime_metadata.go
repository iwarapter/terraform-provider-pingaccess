package pingaccess

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	pa "github.com/iwarapter/pingaccess-sdk-go/pingaccess"
)

func dataSourcePingAccessPingFederateRuntimeMetadata() *schema.Resource {
	return &schema.Resource{
		Read:   dataSourcePingAccessPingFederateRuntimeMetadataRead,
		Schema: dataSourcePingAccessPingFederateRuntimeMetadataSchema(),
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

func dataSourcePingAccessPingFederateRuntimeMetadataRead(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pa.Client).PingFederate
	result, resp, err := svc.GetPingFederateMetadataCommand()
	if err != nil {
		return fmt.Errorf("Error reading PingFederate Runtime Metadata: %s\n%v", err.Error(), resp)
	}
	d.SetId("pingfederate_runtime_metadata")
	d.Set("authorization_endpoint", result.Authorization_endpoint)
	d.Set("backchannel_authentication_endpoint", result.Backchannel_authentication_endpoint)
	d.Set("claim_types_supported", result.Claim_types_supported)
	d.Set("claims_parameter_supported", result.Claims_parameter_supported)
	d.Set("claims_supported", result.Claims_supported)
	d.Set("code_challenge_methods_supported", result.Code_challenge_methods_supported)
	d.Set("end_session_endpoint", result.End_session_endpoint)
	d.Set("grant_types_supported", result.Grant_types_supported)
	d.Set("id_token_signing_alg_values_supported", result.Id_token_signing_alg_values_supported)
	d.Set("introspection_endpoint", result.Introspection_endpoint)
	d.Set("issuer", result.Issuer)
	d.Set("jwks_uri", result.Jwks_uri)
	d.Set("ping_end_session_endpoint", result.Ping_end_session_endpoint)
	d.Set("ping_revoked_sris_endpoint", result.Ping_revoked_sris_endpoint)
	d.Set("request_object_signing_alg_values_supported", result.Request_object_signing_alg_values_supported)
	d.Set("request_parameter_supported", result.Request_parameter_supported)
	d.Set("request_uri_parameter_supported", result.Request_uri_parameter_supported)
	d.Set("response_modes_supported", result.Response_modes_supported)
	d.Set("response_types_supported", result.Response_types_supported)
	d.Set("revocation_endpoint", result.Revocation_endpoint)
	d.Set("scopes_supported", result.Scopes_supported)
	d.Set("subject_types_supported", result.Subject_types_supported)
	d.Set("token_endpoint", result.Token_endpoint)
	d.Set("token_endpoint_auth_methods_supported", result.Token_endpoint_auth_methods_supported)
	d.Set("userinfo_endpoint", result.Userinfo_endpoint)
	d.Set("userinfo_signing_alg_values_supported", result.Userinfo_signing_alg_values_supported)
	return nil
}
