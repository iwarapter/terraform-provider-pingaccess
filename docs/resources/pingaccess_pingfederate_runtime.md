# Resource: pingaccess_pingfederate_runtime

Configured the PingFederate runtime.

-> This resource manages a singleton within PingAccess and as such you should ONLY ever declare one of this resource type. Destroying the resource resets the PingFederate configuration

## Example Usage
```hcl
resource "pingaccess_pingfederate_runtime" "demo" {
  description                  = "foo"
  issuer                       = "https://localhost:9031"
  skip_hostname_verification   = true
  use_slo                      = false
  trusted_certificate_group_id = 2
  use_proxy                    = true
}
```

## Argument Attributes

The following arguments are supported:

- [`description`](#description) - The description of the PingFederate Runtime token provider.
- [`issuer`](#issuer) - The issuer url of the PingFederate token provider.
- [`skip_hostname_verification`](#skip_hostname_verification) - Set to true if HTTP communications to PingFederate should not perform hostname verification of the certificate.
- [`sts_token_exchange_endpoint`](#sts_token_exchange_endpoint) -  The url of the PingFederate STS token-to-token exchange endpoint that is used for token mediation. Specify if it is being served from a different host/port than the issuer is. Otherwise, it is assumed to be {{issuer}}/pf/sts.wst.
- [`trusted_certificate_group_id`](#trusted_certificate_group_id) - The group of certificates to use when authenticating to PingFederate.
- [`use_proxy`](#use_proxy) - Set to true if a proxy should be used for HTTP or HTTPS requests.
- [`use_slo`](#use_slo) - Set to true if OIDC single log out should be used on the /pa/oidc/logout on the engines.

## Attributes Reference

No additional attributes are provided.
