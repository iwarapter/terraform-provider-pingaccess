#Resource: pingaccess_pingfederate_admin

Configured the PingFederate OAuth.

!!! warning
    This resource manages a singleton within PingAccess and as such you should ONLY ever declare one of this resource type. Destroying the resource resets the PingFederate Admin configuration to default values

## Example Usage
```terraform
{!../pingaccess/test_cases/pingaccess_pingfederate_admin.tf!}
```

## Argument Attributes

The following arguments are supported:

- [`admin_password`](#admin_password) - The password for the administrator username.
- [`admin_username`](#admin_username) - The administrator username.
- [`audit_level`](#audit_level) - (Optional) Enable to record requests to the PingFederate Administrative API to the audit store. Either 'ON' or 'OFF'.
- [`base_path`](#base_path) - (Optional) The base path, if needed, for Administration API.
- [`host`](#host) - The host name or IP address for PingFederate Administration API.
- [`port`](#port) - The port number for PingFederate Administration API.
- [`secure`](#secure) - (Optional) Enable if PingFederate is expecting HTTPS connections.
- [`trusted_certificate_group_id`](#trusted_certificate_group_id) - (Optional) The group of certificates to use when authenticating to PingFederate Administrative API.
- [`use_proxy`](#use_proxy) - (Optional) True if a proxy should be used for HTTP or HTTPS requests.

### Attributes Reference

No additional attributes are provided.
