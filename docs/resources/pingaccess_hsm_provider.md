# Resource: pingaccess_hsm_provider

Provides a HSM provider.

-> The PingAccess API does not provider repeatable means of querying a sensitive value, we are unable to detect configuration drift of any sensitive fields in the `configuration` block.

## Example Usage
```hcl
resource "pingaccess_hsm_provider" "test" {
  class_name    = "com.pingidentity.pa.hsm.cloudhsm.plugin.AwsCloudHsmProvider"
  name          = "demo"
  configuration = <<EOF
  {
    "user": "bob",
    "password": "top_secret",
    "partition": "p1"
  }
  EOF
}
```

## Argument Attributes

The following arguments are supported:

- [`class_name`](#class_name) - (Required) The HSM provider's class name.

- [`configuration`](#configuration) - (Required) The HSM provider's configuration data.

- [`name`](#name) - (Required) The HSM provider's name.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

- [`id`](#id) - The HSM provider's ID.

## Import

PingAccess HSM provider can be imported using the id, e.g.

```
$ terraform import pingaccess_hsm_provider.demo_hsm_provider 123
```
