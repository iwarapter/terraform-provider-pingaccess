---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "pingaccess_pingfederate_admin Resource - terraform-provider-pingaccess"
subcategory: ""
description: |-
  Manages the PingAccess PingFederate Admin configuration.
  -> This resource manages a singleton within PingAccess and as such you should ONLY ever declare one of this resource type. Deleting this resource resets the PingFederate Admin configuration to default values.
---

# pingaccess_pingfederate_admin (Resource)

Manages the PingAccess PingFederate Admin configuration.

-> This resource manages a singleton within PingAccess and as such you should ONLY ever declare one of this resource type. Deleting this resource resets the PingFederate Admin configuration to default values.

## Example Usage

```terraform
resource "pingaccess_pingfederate_admin" "example" {
  admin_username = "oauth"
  admin_password {
    value = "top_secret"
  }
  audit_level                  = "ON"
  base_path                    = "/path"
  host                         = "localhost"
  port                         = 9031
  secure                       = true
  trusted_certificate_group_id = 2
  use_proxy                    = true
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `admin_password` (Block List, Min: 1, Max: 1) The password for the administrator username. (see [below for nested schema](#nestedblock--admin_password))
- `admin_username` (String) The administrator username.
- `host` (String) The host name or IP address for PingFederate Administration API.
- `port` (Number) The port number for PingFederate Administration API.

### Optional

- `audit_level` (String) Enable to record requests to the PingFederate Administrative API to the audit store.
- `base_path` (String) The base path, if needed, for Administration API.
- `secure` (Boolean) Enable if PingFederate is expecting HTTPS connections.
- `trusted_certificate_group_id` (Number) The group of certificates to use when authenticating to PingFederate Administrative API.
- `use_proxy` (Boolean) True if a proxy should be used for HTTP or HTTPS requests.

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--admin_password"></a>
### Nested Schema for `admin_password`

Optional:

- `encrypted_value` (String) encrypted value of the field, as originally returned by the API.
- `value` (String, Sensitive) The value of the field. This field takes precedence over the encryptedValue field, if both are specified.

## Import

Import is supported using the following syntax:

```shell
# singleton resource with fixed id.
terraform import pingaccess_pingfederate_admin.example pingfederate_admin_settings
```
