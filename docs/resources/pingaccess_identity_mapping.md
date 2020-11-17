# Resource: pingaccess_identity_mapping

Provides an identity mapping.

-> The PingAccess API does not provider repeatable means of querying a sensitive value, we are unable to detect configuration drift of any sensitive fields in the `configuration` block.

## Example Usage
```hcl
resource "pingaccess_identity_mapping" "demo_identity_mapping" {
  class_name = "com.pingidentity.pa.identitymappings.HeaderIdentityMapping"
  name       = "demo_identity_mapping"

  configuration = <<EOF
  {
    "attributeHeaderMappings": [
      {
        "subject": true,
        "attributeName": "sub",
        "headerName": "sub"
      }
    ],
    "headerClientCertificateMappings": []
  }
  EOF
}
```

## Argument Attributes
The following arguments are supported:

- [`class_name`](#class_name) - (Required) The identity mapping's class name.
- [`configuration`](#configuration) - (Required) The identity mapping's configuration data.
- [`name`](#name) - (Required) The name of the identity mapping.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:

- [`id`](#id) - The identity mapping's ID.

### Import

PingAccess identity mapping can be imported using the id, e.g.

```bash
$ terraform import pingaccess_identity_mapping.demo_identity_mapping 123
```
