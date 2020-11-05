# Resource: pingaccess_site_authenticator

Provides a site authenticator.

~> This resource will store any credentials in the backend state file, please ensure you use an appropriate backend with the relevant encryption/access controls etc for this.

-> The PingAccess API does not provider repeatable means of querying a sensitive value, we are unable to detect configuration drift of any sensitive fields in the `configuration` block.

## Example Usage

```hcl
resource "pingaccess_site_authenticator" "demo_site_authenticator" {
  name       = "demo-site-authenticator"
  class_name = "com.pingidentity.pa.siteauthenticators.BasicAuthTargetSiteAuthenticator"

  configuration = <<EOF
    {
      "username": "cheese",
      "password": {
        "value": "top_5ecr37"
      }
    }
    EOF
}
```

## Argument Attributes

The following arguments are supported:

- [`class_name`](#class_name) - (Required) The site authenticator's class name.

- [`configuration`](#configuration) - (Required) The site authenticator's configuration data.

- [`name`](#name) - (Required) The site authenticator's name.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

- [`id`](#id) - The sites authenticator's ID.

## Import

PingAccess site authenticator can be imported using the id, e.g.

```bash
$ terraform import pingaccess_site_authenticator.demo_site_authenticator 123
```
