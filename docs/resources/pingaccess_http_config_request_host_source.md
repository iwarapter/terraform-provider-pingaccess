#Resource: pingaccess_http_config_request_host_source

Provides HTTP request Host Source type resource.

!!! warning
    This resource manages a singleton within PingAccess and as such you should ONLY ever declare one of this resource type. Destroying the resource will resets the HTTP request Host Source type to default values

## Example Usage
```terraform
{!../func-tests//http_config_request_host_source.tf!}
```

## Argument Attributes

The following arguments are supported:

- [`header_name_list`](#header_name_list) - An array of header names used to identify the host source name.

- [`list_value_location`](#list_value_location) - The location in a matching header value list to use as the source. Either FIRST or LAST.

### Attributes Reference

No additional attributes are provided.
