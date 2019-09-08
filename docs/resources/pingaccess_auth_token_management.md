Provides a auth token management.

## Example Usage
```terraform
{!../pingaccess/test_cases/auth_token_management.tf!}
```

## Argument Attributes

The following arguments are supported:

- [`key_roll_enabled`](#key_roll_enabled) - The issuer value to include in auth tokens. PingAccess inserts this value as the iss claim within the auth tokens.

- [`key_roll_period_in_hours`](#key_roll_period_in_hours) - This field is true if key rollover is enabled. When false, PingAccess will not rollover keys at the configured interval.

- [`issuer`](#issuer) - The interval (in hours) at which PingAccess will roll the keys. Key rollover updates keys at regular intervals to ensure the security of signed auth tokens.

- [`signing_algorithm`](#signing_algorithm) - The signing algorithm used when creating signed auth tokens.

### Attributes Reference

No additional attributes are provided.
