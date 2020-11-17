# Resource: pingaccess_http_config_request_host_source

Provides HTTP request Host Source type resource.

-> This resource manages a singleton within PingAccess and as such you should ONLY ever declare one of this resource type. Destroying the resource will resets the HTTP request Host Source type to default values

## Example Usage
```hcl
resource "pingaccess_http_config_request_host_source" "test" {
  header_name_list = [
    "Host",
    "X-Forwarded-Host"
  ]
  list_value_location = "LAST"
}
```

## Argument Attributes

The following arguments are supported:

- [`header_name_list`](#header_name_list) - An array of header names used to identify the host source name.

- [`list_value_location`](#list_value_location) - The location in a matching header value list to use as the source. Either FIRST or LAST.

## Attributes Reference

No additional attributes are provided.

## Import

PingAccess HTTP request Host Source resources can be imported using the id, e.g.

```bash
$ terraform import pingaccess_http_config_request_host_source.example 123
```
