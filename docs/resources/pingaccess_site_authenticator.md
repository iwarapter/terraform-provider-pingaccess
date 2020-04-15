#Resource: pingaccess_site_authenticator

Provides a site authenticator.

!!! warning
    This resource will store any credentials in the backend state file, please ensure you use an appropriate backend with the relevant encryption/access controls etc for this.

!!! tip
    The PingAccess API does not provider repeatable means of querying a sensitive value, we are unable to detect configuration drift of any sensitive fields in the `configuration` block.

## Example Usage

```terraform
{!../pingaccess/test_cases/site_authenticator.tf!}
```

## Argument Attributes

The following arguments are supported:

- [`class_name`](#class_name) - (Required) The site authenticator's class name.

- [`configuration`](#configuration) - (Required) The site authenticator's configuration data.

- [`name`](#name) - (Required) The site authenticator's name.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

- [`id`](#id) - The sites authenticator's ID.

## Import

PingAccess site authenticator can be imported using the id, e.g.

```
$ terraform import pingaccess_site_authenticator.demo_site_authenticator 123
```