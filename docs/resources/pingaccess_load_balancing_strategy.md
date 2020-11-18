# Resource: pingaccess_load_balancing_strategy

Provides a load balancing strategy.

-> The PingAccess API does not provider repeatable means of querying a sensitive value, we are unable to detect configuration drift of any sensitive fields in the `configuration` block.

## Example Usage
```hcl
resource "pingaccess_load_balancing_strategy" "test" {
  name = "example"
  class_name = "com.pingidentity.pa.ha.lb.header.HeaderBasedLoadBalancingPlugin"
  configuration = <<EOF
  {
    "headerName": "example",
    "fallbackToFirstAvailableHost": false
  }
  EOF
}
```

## Argument Attributes

The following arguments are supported:

- [`class_name`](#class_name) - (Required) The load balancing strategies class name.

- [`configuration`](#configuration) - (Required) The load balancing strategies configuration data.

- [`name`](#name) - (Required) The load balancing strategies name.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

- [`id`](#id) - The load balancing strategies ID.

## Import

PingAccess load balancing strategy can be imported using the id, e.g.

```
$ terraform import pingaccess_load_balancing_strategy.example 123
```
