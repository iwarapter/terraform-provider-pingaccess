# Resource: pingaccess_pingfederate_oauth

Configured the PingFederate OAuth.

-> This resource manages a singleton within PingAccess and as such you should ONLY ever declare one of this resource type. Destroying the resource resets the PingAccess OAuth Client configuration to default values

## Example Usage
```hcl
resource "pingaccess_pingfederate_oauth" "demo" {
  access_validator_id    = 1
  cache_tokens           = true
  subject_attribute_name = "san"
  name                   = "PingFederate"
  client_id              = "oauth"
  client_secret {
    value = "top_secret"
  }
  send_audience              = true
  token_time_to_live_seconds = 300
  use_token_introspection    = true
}
```

## Argument Attributes

The following arguments are supported:

- [`access_validator_id`](#access_validator_id) - (Optional) The Access Validator Id.
- [`cache_tokens`](#cache_tokens) - (Optional) Enable to retain token details for subsequent requests.
- [`client_id`](#client_id) - The Client ID which PingAccess should use when requesting PingFederate to validate access tokens. The client must have Access Token Validation grant type allowed.
- [`client_secret`](#client_secret) - (Optional) The Client Secret for the Client ID.
- [`name`](#name) - (Optional) The unique Access Validator name.
- [`send_audience`](#send_audience) - (Optional) Enable to send the URI the user requested as the 'aud' OAuth parameter for PingAccess to use to select an Access Token Manager.
- [`subject_attribute_name`](#subject_attribute_name) - The attribute you want to use from the OAuth access token as the subject for auditing purposes.
- [`token_time_to_live_seconds`](#token_time_to_live_seconds) - (Optional) Defines the number of seconds to cache the access token. -1 means no limit. This value should be less than the PingFederate Token Lifetime.
- [`use_token_introspection`](#use_token_introspection) - (Optional) Specify if token introspection is enabled.

## Attributes Reference

No additional attributes are provided.
