# Resource: pingaccess_availability_profile

Provides an Availability Profile.

-> The PingAccess API does not provider repeatable means of querying a sensitive value, we are unable to detect configuration drift of any sensitive fields in the `configuration` block.

## Example Usage
```hcl
resource "pingaccess_availability_profile" "example" {
  name          = "example"
  class_name    = "com.pingidentity.pa.ha.availability.ondemand.OnDemandAvailabilityPlugin"
  configuration = <<EOF
  {
    "connectTimeout": 10000,
    "pooledConnectionTimeout": -1,
    "readTimeout": -1,
    "maxRetries": 2,
    "retryDelay": 250,
    "failedRetryTimeout": 60,
    "failureHttpStatusCodes": []
  }
  EOF
}
```

## Argument Attributes

The following arguments are supported:

- [`class_name`](#class_name) - (Required) The availability profile's class name.

- [`configuration`](#configuration) - (Required) The availability profile's configuration data.

- [`name`](#name) - (Required) The availability profile's name.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

- [`id`](#id) - The availability profile's ID.

## Import

PingAccess Availability Profile can be imported using the id, e.g.

```
$ terraform import pingaccess_availability_profile.example 123
```
