Create a Load Balancing Strategy.

## Example Usage
```terraform
resource "pingaccess_load_balancing_strategy" "demo_strategy" {
	   name                         = "Round robin policy"
	   class_name                    = "com.pingidentity.pa.ha.lb.roundrobin.CookieBasedRoundRobinPlugin"
	   configuration     = {
        "stickySessionEnabled": true,
        	"cookieName": "PA_S",
        }
	}
```

## Argument Attributes

The following arguments are supported:

- [`className`](#className) - (Required) The class name of the load balancing strategy.
- [`configuration`](#configuration) - (Required) The load balancing strategy's configuration data.
- [`name`](#name) - (Required) The name of the load balancing strategy.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

- [`id`](#id) - The strategy's ID.

## Import

PingAccess Load balancing strategy can be imported using the id, e.g.

```shell
$ terraform import pingaccess_load_balancing_strategy.demo_loadstrategy 123
```