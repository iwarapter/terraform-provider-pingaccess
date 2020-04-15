#Resource: pingaccess_https_listener

Provides a https listener.

## Example Usage
```terraform
{!../pingaccess/test_cases/https_listener.tf!}
```

## Argument Attributes

The following arguments are supported:

- [`name`](#name) - (Required) The name of the HTTPS listener.
- [`key_pair_id`](#key_pair_id) - (Required) The ID of the default key pair used by the HTTPS listener.
- [`use_server_cipher_suite_order`](#use_server_cipher_suite_order) - (Required) Enable server cipher suite ordering for the HTTPS listener.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

- [`id`](#id) - The https listener's ID.

## Import

PingAccess https_listeners can be imported using the id, e.g.

```shell
$ terraform import pingaccess_https_listener.demo 123
```