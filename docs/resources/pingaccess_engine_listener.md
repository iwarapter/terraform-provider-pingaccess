#Resource: pingaccess_engine_listener

Provides a engine listener.

## Example Usage
```terraform
{!../func-tests//engine_listener.tf!}
```

## Argument Attributes

The following arguments are supported:

- [`name`](#name) - (Required) The name of the engine listener.
- [`port`](#port) - (Required) The port the engine listener listens on.
- [`secure`](#secure) - Indicator if the engine listener should listen to HTTPS connections.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

- [`id`](#id) - The engine listener's ID.

## Import

PingAccess engine listeners can be imported using the id, e.g.

```shell
$ terraform import pingaccess_engine_listener.demo 123
```
