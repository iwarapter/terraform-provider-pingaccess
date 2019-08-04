Provides a site authenticator.

!!! warning
    This resource will store any credentials in the backend state file, please ensure you use an appropriate backend with the relevent encryption/access controls etc for this.

## Example Usage

```terraform
{!../pingaccess/test_cases/site_authenticator.tf!}
```

## Argument Attributes

The following arguments are supported:

- [`class_name`](#class_name) - (Required) The site authenticator's class name.

- [`configuration`](#configuration) - (Required) The site authenticator's configuration data.

- [`name`](#name) - (Required) The site authenticator's name.

- [`hidden_fields`](#hidden_fields) - (Optional) This is a configuration field, to help manage credentials within the resource. As the PingAccess API doesn't return the password used and the encryptedValue field changes on every call we need to store the password field and mask it so we can correctly identify when it has changed (according to the resource not the actual API).


## Attributes Reference

In addition to all arguments above, the following attributes are exported:

- [`id`](#id) - The sites authenticator's ID.

## Import

PingAccess site authenticator can be imported using the id, e.g.

```
$ terraform import pingaccess_site_authenticator.demo_site_authenticator 123
```