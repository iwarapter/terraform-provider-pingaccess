# Resource: pingaccess_acme_server

Provides a acme server.

## Example Usage
```hcl
resource "pingaccess_acme_server" "example" {
  name = "example"
  url  = "https://acme-staging-v02.api.letsencrypt.org/directory"
}
```

## Argument Attributes

The following arguments are supported:

- [`name`](#name) - (Required) A user-friendly name for the ACME server.

- [`url`](#url) - (Required) The URL of the ACME directory resource on the ACME server.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

- [`id`](#id) - The acme server's ID.

- [`acme_accounts`](#acme_accounts) - An array of references to accounts.

## Import

PingAccess acme server can be imported using the id, e.g.

```
$ terraform import pingaccess_acme_server.demo_acme_server 123
```
