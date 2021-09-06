---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "pingaccess_auth_token_management Resource - terraform-provider-pingaccess"
subcategory: ""
description: |-
  Manages the PingAccess Auth Token Management configuration.
  -> This resource manages a singleton within PingAccess and as such you should ONLY ever declare one of this resource type. Deleting this resource resets the Auth Token Management configuration to default values.
---

# pingaccess_auth_token_management (Resource)

Manages the PingAccess Auth Token Management configuration.

-> This resource manages a singleton within PingAccess and as such you should ONLY ever declare one of this resource type. Deleting this resource resets the Auth Token Management configuration to default values.

## Example Usage

```terraform
resource "pingaccess_auth_token_management" "example" {
  key_roll_enabled         = true
  key_roll_period_in_hours = 24
  issuer                   = "PingAccessAuthToken"
  signing_algorithm        = "P-256"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- **issuer** (String) The issuer value to include in auth tokens. PingAccess inserts this value as the iss claim within the auth tokens.
- **key_roll_enabled** (Boolean) This field is true if key rollover is enabled. When false, PingAccess will not rollover keys at the configured interval.
- **key_roll_period_in_hours** (Number) The interval (in hours) at which PingAccess will roll the keys. Key rollover updates keys at regular intervals to ensure the security of signed auth tokens.
- **signing_algorithm** (String) The signing algorithm used when creating signed auth tokens.

### Read-Only

- **id** (String) The ID of this resource.

## Import

Import is supported using the following syntax:

```shell
# singleton resource with fixed id.
terraform import pingaccess_auth_token_management.example auth_token_management
```