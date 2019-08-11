Provides a access token validator.

## Example Usage
```terraform
{!../pingaccess/test_cases/access_token_validator.tf!}
```

## Argument Attributes

The following arguments are supported:

- [`class_name`](#class_name) - (Required) The access token validator's class name.

- [`configuration`](#configuration) - (Required) The access token validator's configuration data.

- [`name`](#name) - (Required) The access token validator's name.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

- [`id`](#id) - The access token validator's ID.

## Import

PingAccess access token validator can be imported using the id, e.g.

```
$ terraform import pingaccess_access_token_validator.demo_access_token_validator 123
```
