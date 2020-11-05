# Resource: pingaccess_access_token_validator

Provides a access token validator.

-> The PingAccess API does not provider repeatable means of querying a sensitive value, we are unable to detect configuration drift of any sensitive fields in the `configuration` block.

## Example Usage
```hcl
resource "pingaccess_access_token_validator" "example" {
  class_name = "com.pingidentity.pa.accesstokenvalidators.JwksEndpoint"
  name       = "demo"

  configuration = <<EOF
	{
		"description": null,
		"path": "/bar",
		"subjectAttributeName": "demo",
		"issuer": null,
		"audience": null
	}
	EOF
}
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
