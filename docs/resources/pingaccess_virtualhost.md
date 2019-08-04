Provides a virtualhost.

## Example Usage
```terraform
resource "pingaccess_virtualhost" "demo_virtualhost" {
  host = "localhost"
  port = 1234
}
```

## Argument Attributes

The following arguments are supported:

- [`access_validator_id`](#access_validator_id) - Non-zero if the application is protected by an Authorization Server. Only applies to applications of type API.
- [`pingaccess_virtualhost_host`](#pingaccess_virtualhost_host) - (Required) The host name for the Virtual Host.
- [`pingaccess_virtualhost_port`](#pingaccess_virtualhost_port) - (Required) The integer port number for the Virtual Host.
- [`pingaccess_virtualhost_agent_resource_cache_ttl`](#pingaccess_virtualhost_agent_resource_cache_ttl) - (Optional) Indicates the number of seconds the Agent can cache resources for this application.
- [`pingaccess_virtualhost_key_pair_id`](#pingaccess_virtualhost_key_pair_id) - (Optional) Key pair assigned to Virtual Host used by SNI, If no key pair is assigned to a virtual host, ENGINE HTTPS Listener key pair will be used.
- [`pingaccess_virtualhost_trusted_certificate_group_id`](#pingaccess_virtualhost_trusted_certificate_group_id) - (Optional) Trusted Certificate Group assigned to Virtual Host for client certificate authentication.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

- [`id`](#id) - The virtualhost's ID.

## Import

PingAccess Virtualhosts can be imported using the id, e.g.

```shell
$ terraform import pingaccess_virtualhost.demo_virtualhost 123
```