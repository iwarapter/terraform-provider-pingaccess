Provides a identity mapping.

### Example Usage
```terraform
{!../pingaccess/test_cases/identity_mapping.tf!}
```
### Argument Attributes
The following arguments are supported:

- [`class_name`](#class_name) - (Required) The identity mapping's class name.
- [`configuration`](#configuration) - (Required) The identity mapping's configuration data.
- [`name`](#name) - (Required) The name of the identity mapping.

### Attributes Reference
In addition to all arguments above, the following attributes are exported:

- [`id`](#id) - The identity mapping's ID.

### Import
PingAccess identity mapping can be imported using the id, e.g.

```bash
$ terraform import pingaccess_identity_mapping.demo_identity_mapping 123
```