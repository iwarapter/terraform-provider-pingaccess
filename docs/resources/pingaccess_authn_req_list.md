#Resource: pingaccess_authn_req_list

Provides a AuthN Req List.

## Example Usage
```terraform
{!../func-tests//authn_req_list.tf!}
```

## Argument Attributes

The following arguments are supported:

- [`name`](#name) - (Required) The name of the AuthN Req List.
- [`authn_reqs`](#authn_reqs) - (Required) The list of AuthN requirements.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

- [`id`](#id) - The AuthN Req List's ID.

## Import

PingAccess authn_req_list can be imported using the id, e.g.

```shell
$ terraform import authn_req_list.demo 123
```
