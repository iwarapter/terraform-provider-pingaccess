#Resource: pingaccess_ruleset

Provides a ruleset.

## Example Usage
```terraform
{!../func-tests//ruleset.tf!}
```

## Argument Attributes

The following arguments are supported:

- [`element_type`](#element_type) - (Required) ['Rule' or 'RuleSet']

- [`name`](#name) - (Required) The ruleset's name.

- [`policy`](#policy) - (Required) The list of policy ids assigned to the ruleset.

- [`success_criteria`](#success_criteria) - (Required) ['SuccessIfAllSucceed' or 'SuccessIfAnyOneSucceeds']: The ruleset's success criteria.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

- [`id`](#id) - The ruleset's ID.

## Import

PingAccess ruleset can be imported using the id, e.g.

```
$ terraform import pingaccess_ruleset.demo_ruleset 123
```
