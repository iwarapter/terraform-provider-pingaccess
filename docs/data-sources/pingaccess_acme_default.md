# Data Source: pingaccess_acme_default

Use this data source to get the ID of the default ACME server within PingAccess.

## Example Usage
```hcl
data "pingaccess_acme_default" "default" {}
```
## Argument Attributes
This data source requires no attributes to be set.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

- [`id`](#id) - The ACME servers's ID.

- [`location`](#location) - An absolute path to the associated resource.
