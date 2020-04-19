#Resource: pingaccess_trusted_certificate_group

Provides a trusted certificate group.

## Example Usage
```terraform
{!../pingaccess/test_cases/trusted_certificate_group.tf!}
```

## Argument Attributes

The following arguments are supported:

- [`cert_ids`](#cert_ids) - The IDs of the certificates that are in the trusted certificate group.

- [`ignore_all_certificate_errors`](#ignore_all_certificate_errors) -  This field is read-only and is only set to true
for the Trust Any certificate group.

- [`name`](#name) -  The name of trusted certificate group.

- [`skip_certificate_date_check`](#skip_certificate_date_check) -  This field is true if certificates that have expired or are not yet valid but have passed the other certificate checks should be trusted.

- [`use_java_trust_store`](#use_java_trust_store) -  This field is true if the certificates in the group should also include all certificates in the Java Trust Store.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

- [`id`](#id) - The trusted certificate group's ID.

- [`system_group`](#system_group) -  This field is read-only and indicates the trusted certificate group cannot be modified.

## Import

PingAccess trusted certificate group can be imported using the id, e.g.

```
$ terraform import pingaccess_trusted_certificate_group.demo_trusted_certificate_group 123
```
