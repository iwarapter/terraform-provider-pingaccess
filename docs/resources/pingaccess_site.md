#Resource: pingaccess_site

Provides a site.

## Example Usage
```terraform
{!../pingaccess/test_cases/site.tf!}
```

## Argument Attributes

The following arguments are supported:

- [`availability_profile_id`](#availability_profile_id) - (Optional) The ID of the availability profile associated with the site.

- [`expected_hostname`](#expected_hostname) - (Optional) The name of the host expected in the site's certificate.

- [`keep_alive_timeout`](#keep_alive_timeout) - (Optional) The time, in milliseconds, that an HTTP persistent connection to the site can be idle before PingAccess closes the connection.

- [`load_balancing_strategy_id`](#load_balancing_strategy_id) - (Optional) The ID of the load balancing strategy associated with the site.

- [`max_connections`](#max_connections) - (Optional) The maximum number of HTTP persistent connections you want PingAccess to have open and maintain for the site. -1 indicates unlimited connections.

- [`max_web_socket_connections`](#max_web_socket_connections) - (Optional) The maximum number of WebSocket connections you want PingAccess to have open and maintain for the site. -1 indicates unlimited connections.

- [`name`](#name) - (Required) The name of the site.

- [`secure`](#secure) - (Optional) This field is true if the site expects HTTPS connections.

- [`send_pa_cookie`](#send_pa_cookie) - (Optional) This field is true if the PingAccess Token or OAuth Access Token should be included in the request to the site.

- [`site_authenticator_ids`](#site_authenticator_ids) - (Optional) The IDs of the site authenticators associated with the site.

- [`skip_hostname_verification`](#skip_hostname_verification) - (Optional) This field is true if the hostname verification of the site's certificate should be skipped.

- [`targets`](#targets) - (Required) The {hostname}:{port} pairs for the hosts that make up the site.

- [`trusted_certificate_group_id`](#trusted_certificate_group_id) - (Optional) The ID of the trusted certificate group associated with the site.

- [`use_proxy`](#use_proxy) - (Optional) True if a proxy should be used for HTTP or HTTPS requests.

- [`use_target_host_header`](#use_target_host_header) - (Optional) Setting this field to true causes PingAccess to adjust the Host header to the site's selected target host rather than the virtual host configured in the application.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

- [`id`](#id) - The site's ID.

## Import

PingAccess sites can be imported using the id, e.g.

```
$ terraform import pingaccess_site.demo_site 123
```
