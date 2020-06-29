#Resource: pingaccess_rule

Provides a rule.

!!! tip
    The PingAccess API does not provider repeatable means of querying a sensitive value, we are unable to detect configuration drift of any sensitive fields in the `configuration` block.

## Example Usage
```terraform
{!../func-tests//rule.tf!}
```

## Argument Attributes

The following arguments are supported:

- [`class_name`](#class_name) - (Required) The rule's class name.

- [`configuration`](#configuration) - (Required) The rule's configuration data.

- [`name`](#name) - (Required) The rule's name.

- [`supported_destinations`](#supported_destinations) - (Optional) The supported destinations for this rule.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

- [`id`](#id) - The rule's ID.

## Import

PingAccess rule can be imported using the id, e.g.

```
$ terraform import pingaccess_rule.demo_rule 123
```
