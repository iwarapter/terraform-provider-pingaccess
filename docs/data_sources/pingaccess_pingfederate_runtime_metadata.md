#Data Source: pingaccess_pingfederate_runtime_metadata

Use this data source to get the PingFederate Runtime metadata.

### Example Usage
```terraform
data "pingaccess_pingfederate_runtime_metadata" "meta" {}
```
### Argument Attributes
This data source requires no attributes to be set.

### Attributes Reference

In addition to all arguments above, the following attributes are exported:

- [`authorization_endpoint`](#authorization_endpoint) - URL of the OpenID Connect provider's authorization endpoint.
- [`backchannel_authentication_endpoint`](#backchannel_authentication_endpoint) - The backchannel authentication endpoint.
- [`claim_types_supported`](#claim_types_supported) - JSON array containing a list of the claim types that the OpenID Connect provider supports.
- [`claims_parameter_supported`](#claims_parameter_supported) - boolean value specifying whether the OpenID Connect provider supports use of the claims parameter, with true indicating support.
- [`claims_supported`](#claims_supported) - JSON array containing a list of the claim names of the claims that the OpenID Connect provider MAY be able to supply values for.
- [`code_challenge_methods_supported`](#code_challenge_methods_supported) - Proof Key for Code Exchange (PKCE) code challenge methods supported by this OpenID Connect provider.
- [`end_session_endpoint`](#end_session_endpoint) - URL at the OpenID Connect provider to which a relying party can perform a redirect to request that the end-user be logged out at the OpenID Connect provider.
- [`grant_types_supported`](#grant_types_supported) - JSON array containing a list of the OAuth 2.0 grant type values that this OpenID Connect provider supports.
- [`id_token_signing_alg_values_supported`](#id_token_signing_alg_values_supported) - JSON array containing a list of the JWS signing algorithms supported by the OpenID Connect provider for the id token to encode the claims in a JWT.
- [`introspection_endpoint`](#introspection_endpoint) - URL of the OpenID Connect provider's OAuth 2.0 introspection endpoint.
- [`issuer`](#issuer) - OpenID Connect provider's issuer identifier URL.
- [`jwks_uri`](#jwks_uri) - URL of the OpenID Connect provider's JWK Set document.
- [`ping_end_session_endpoint`](#ping_end_session_endpoint) - PingFederate logout endpoint. (Not applicable if PingFederate is not the OpenID Connect provider)
- [`ping_revoked_sris_endpoint`](#ping_revoked_sris_endpoint) - PingFederate session revocation endpoint. (Not applicable if PingFederate is not the OpenID Connect provider)
- [`request_object_signing_alg_values_supported`](#request_object_signing_alg_values_supported) - JSON array containing a list of the JWS signing algorithms supported by the OpenID Connect provider for request objects.
- [`request_parameter_supported`](#request_parameter_supported) - boolean value specifying whether the OpenID Connect provider supports use of the request parameter, with true indicating support.
- [`request_uri_parameter_supported`](#request_uri_parameter_supported) - boolean value specifying whether the OpenID Connect provider supports use of the request_uri parameter, with true indicating support.
- [`response_modes_supported`](#response_modes_supported) - JSON array containing a list of the OAuth 2.0 "response_mode" values that this OpenID Connect provider supports.
- [`response_types_supported`](#response_types_supported) - JSON array containing a list of the OAuth 2.0 "response_type" values that this OpenID Connect provider supports.
- [`revocation_endpoint`](#revocation_endpoint) - URL of the OpenID Connect provider's OAuth 2.0 revocation endpoint.
- [`scopes_supported`](#scopes_supported) - JSON array containing a list of the OAuth 2.0 "scope" values that this OpenID Connect provider supports.
- [`subject_types_supported`](#subject_types_supported) - JSON array containing a list of the Subject Identifier types that this OpenID Connect provider supports.
- [`token_endpoint`](#token_endpoint) - URL of the OpenID Connect provider's token endpoint.
- [`token_endpoint_auth_methods_supported`](#token_endpoint_auth_methods_supported) - JSON array containing a list of client authentication methods supported by this token endpoint.
- [`userinfo_endpoint`](#userinfo_endpoint) - URL of the OpenID Connect provider's userInfo endpoint.
- [`userinfo_signing_alg_values_supported`](#userinfo_signing_alg_values_supported) - JSON array containing a list of the JWS signing algorithms supported by the userInfo endpoint to encode the claims in a JWT.
