# Resource: pingaccess_third_party_service

Provides a third party service.

## Example Usage
```hcl
resource "pingaccess_third_party_service" "demo_third_party_service" {
  name = "demo"

  targets = [
    "server.domain:1234",
  ]
}
```

## Argument Attributes

The following arguments are supported:

- [`availability_profile_id`](#availability_profile_id) - (Optional) The ID of the availability profile associated with the third-party service.

- [`expected_hostname`](#expected_hostname) - (Optional) The name of the host expected in the third-party service's certificate.

- [`host_value`](#host_value) - (Optional) The Host header field value in the requests sent to a Third-Party Services. When set, PingAccess will use the hostValue as the Host header field value. Otherwise, the target value will be used.

- [`load_balancing_strategy_id`](#load_balancing_strategy_id) - (Optional) The ID of the load balancing strategy associated with the third-party service.

- [`max_connections`](#max_connections) - (Optional) The maximum number of HTTP persistent connections you want PingAccess to have open and maintain for the third-party service. -1 indicates unlimited connections.

- [`name`](#name) - (Required) The name of the third-party service.

- [`secure`](#secure) - (Optional) This field is true if the third-party service expects HTTPS connections.

- [`skip_hostname_verification`](#skip_hostname_verification) - (Optional) This field is true if the hostname verification of the third-party service's certificate should be skipped.

- [`targets`](#targets) - (Required) The {hostname}:{port} pairs for the hosts that make up the third-party service.

- [`trusted_certificate_group_id`](#trusted_certificate_group_id) - (Optional) The ID of the trusted certificate group associated with the third-party service.

- [`use_proxy`](#use_proxy) - (Optional) True if a proxy should be used for HTTP or HTTPS requests.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

- [`id`](#id) - The third party service's ID.

## Import

PingAccess third party service can be imported using the id, e.g.

```bash
$ terraform import pingaccess_third_party_service.demo_third_party_service 123
```
