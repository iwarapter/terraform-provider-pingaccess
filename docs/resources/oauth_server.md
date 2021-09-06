---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "pingaccess_oauth_server Resource - terraform-provider-pingaccess"
subcategory: ""
description: |-
  Manages the PingAccess Third Party OAuth Server configuration.
  -> This resource manages a singleton within PingAccess and as such you should ONLY ever declare one of this resource type. Deleting this resource resets the Third Party OAuth Server configuration to default values.
---

# pingaccess_oauth_server (Resource)

Manages the PingAccess Third Party OAuth Server configuration.

-> This resource manages a singleton within PingAccess and as such you should ONLY ever declare one of this resource type. Deleting this resource resets the Third Party OAuth Server configuration to default values.

## Example Usage

```terraform
resource "pingaccess_oauth_server" "oauth_server" {
  targets                      = ["localhost:9031"]
  subject_attribute_name       = "san"
  trusted_certificate_group_id = 1
  introspection_endpoint       = "/introspect"
  secure                       = true
  client_credentials {
    client_id = "oauth"
    client_secret {
      value = "top_secret"
    }
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **client_credentials** (Block List, Min: 1, Max: 1) Specify the client credentials. (see [below for nested schema](#nestedblock--client_credentials))
- **introspection_endpoint** (String) The third-party OAuth 2.0 Authorization Server's token introspection endpoint.
- **subject_attribute_name** (String) The attribute you want to use from the OAuth access token as the subject for auditing purposes.
- **targets** (Set of String) One or more server hostname:port pairs used to access third-party OAuth 2.0 Authorization Server.
- **trusted_certificate_group_id** (Number) The group of certificates to use when authenticating to third-party OAuth 2.0 Authorization Server.

### Optional

- **audit_level** (String) Enable to record requests to third-party OAuth 2.0 Authorization Server to the audit store.
- **cache_tokens** (Boolean) Enable to retain token details for subsequent requests.
- **description** (String) The description of the third-party OAuth 2.0 Authorization Server.
- **secure** (Boolean) Enable if third-party OAuth 2.0 Authorization Server is expecting HTTPS connections.
- **send_audience** (Boolean) Enable to send the URI the user requested as the 'aud' OAuth parameter for PingAccess to the OAuth 2.0 Authorization server.
- **token_time_to_live_seconds** (Number) Defines the number of seconds to cache the access token. -1 means no limit. This value should be less than the PingFederate Token Lifetime.
- **use_proxy** (Boolean) True if a proxy should be used for HTTP or HTTPS requests.

### Read-Only

- **id** (String) The ID of this resource.

<a id="nestedblock--client_credentials"></a>
### Nested Schema for `client_credentials`

Required:

- **client_id** (String) Specify the client ID.

Optional:

- **client_secret** (Block List, Max: 1) Specify the client secret. (see [below for nested schema](#nestedblock--client_credentials--client_secret))
- **credentials_type** (String) Specify the credential type.
- **key_pair_id** (Number) Specify the ID of a key pair to use for mutual TLS.

<a id="nestedblock--client_credentials--client_secret"></a>
### Nested Schema for `client_credentials.client_secret`

Optional:

- **encrypted_value** (String) encrypted value of the field, as originally returned by the API.
- **value** (String, Sensitive) The value of the field. This field takes precedence over the encryptedValue field, if both are specified.

## Import

Import is supported using the following syntax:

```shell
# singleton resource with fixed id.
terraform import pingaccess_oauth_server.oauth_server oauth_server_settings
```