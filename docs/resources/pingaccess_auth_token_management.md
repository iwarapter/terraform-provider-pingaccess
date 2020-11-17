# Resource: pingaccess_auth_token_management

Provides a auth token management.

->  This resource manages a singleton within PingAccess and as such you should ONLY ever declare one of this resource type. Destroying the resource will resets the Auth Token Management configuration to default values

## Example Usage
```hcl
resource "pingaccess_auth_token_management" "demo" {
  key_roll_enabled         = true
  key_roll_period_in_hours = 24
  issuer                   = "PingAccessAuthToken"
  signing_algorithm        = "P-256"
}
```

## Argument Attributes

The following arguments are supported:

- [`key_roll_enabled`](#key_roll_enabled) - The issuer value to include in auth tokens. PingAccess inserts this value as the iss claim within the auth tokens.

- [`key_roll_period_in_hours`](#key_roll_period_in_hours) - This field is true if key rollover is enabled. When false, PingAccess will not rollover keys at the configured interval.

- [`issuer`](#issuer) - The interval (in hours) at which PingAccess will roll the keys. Key rollover updates keys at regular intervals to ensure the security of signed auth tokens.

- [`signing_algorithm`](#signing_algorithm) - The signing algorithm used when creating signed auth tokens.

## Attributes Reference

No additional attributes are provided.

## Import

-> The resource ID is fixed as `auth_token_management` because this is a singleton resource.

Auth Token Management can be imported using the id, e.g.

```bash
$ terraform import pingaccess_auth_token_management.demo auth_token_management
```
