#Data Source: pingaccess_trusted_certificate_group

Use this data source to get the ID of a trusted certificate group in Ping Access, you can reference it by name without having to hard code the IDs as input.

### Example Usage
```terraform
data "pingaccess_trusted_certificate_group" "trust_any" {
  name = "Trust Any"
}
```
### Argument Attributes
The following arguments are supported:

- [`name`](#name) - (required) The name for the trusted certificate group.

### Attributes Reference

In addition to all arguments above, the following attributes are exported:

- [`id`](#id) - The certificates's ID.

- [`cert_ids`](#cert_ids) - The IDs of the certificates that are in the trusted certificate group.

- [`ignore_all_certificate_errors`](#ignore_all_certificate_errors) -  This field is read-only and is only set to true 
for the Trust Any certificate group.

- [`skip_certificate_date_check`](#skip_certificate_date_check) -  This field is true if certificates that have expired or are not yet valid but have passed the other certificate checks should be trusted.

- [`system_group`](#system_group) -  This field is read-only and indicates the trusted certificate group cannot be modified.

- [`use_java_trust_store`](#use_java_trust_store) -  This field is true if the certificates in the group should also include all certificates in the Java Trust Store.
