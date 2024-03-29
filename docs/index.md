# PingAccess Provider

The PingAccess provider is used to interact with the many resources supported by the PingAccess admin API. The provider needs to be configured with the proper credentials before it can be used.

Use the navigation to the left to read about the available resources.

## Example Usage
Terraform 0.13 and later:
```hcl
# Configure the PingAccess Provider
terraform {
  required_providers {
    pingaccess = {
      source = "iwarapter/pingaccess"
      version = "0.5.0"
    }
  }
}

provider "pingaccess" {
  username = "Administrator"
  password = "2Access"
  base_url = "https://localhost:9000"
  context  = "/pa-admin-api/v3"
}

# Create a site
resource "pingaccess_site" "site" {
  # ...
}
```
Terraform 0.12 and earlier:
```hcl
# Configure the PingAccess Provider
provider "pingaccess" {
  username = "Administrator"
  password = "2Access"
  base_url = "https://localhost:9000"
  context  = "/pa-admin-api/v3"
}

# Create a site
resource "pingaccess_site" "site" {
  # ...
}
```

## Authentication

The PingAccess provider offers a flexible means of providing credentials for authentication. The following methods are supported, in this order, and explained below:

- Static credentials

- Environment variables

### Static credentials
Static credentials can be provided by adding an `username` and `password` in-line in the PingAccess provider block:

Usage:
```terraform
provider "pingaccess" {
  username = "Administrator"
  password = "2Access"
  base_url = "https://localhost:9000"
  context  = "/pa-admin-api/v3"
}
```

### Environment variables
You can provide your credentials via the `PINGACCESS_USERNAME`, `PINGACCESS_PASSWORD`, `PINGACCESS_CONTEXT` and `PINGACCESS_BASEURL` environment variables.

```terraform
provider "pingaccess" {}
```

Usage:
```bash
$ export PINGACCESS_USERNAME="Administrator"
$ export PINGACCESS_PASSWORD="top_secret"
$ export PINGACCESS_CONTEXT="/pa-admin-api/v3"
$ export PINGACCESS_BASEURL="https://myadmin.server:9000"
$ terraform plan
```


## Argument Reference

In addition to [generic `provider` arguments](https://www.terraform.io/docs/configuration/providers.html)
(e.g. `alias` and `version`), the following arguments are supported in the PingAccess
 `provider` block:

## Schema

### Optional

- **username** (String) This is the PingAccess administrative username. It must be provided, but
  it can also be sourced from the `PINGACCESS_USERNAME` environment variable.

- **password** (String, Sensitive) This is the PingAccess administrative password. It must be provided, but
  it can also be sourced from the `PINGACCESS_PASSWORD` environment variable.

- **base_url** (String) This is the PingAccess base url (protocol:server:port). It must be provided, but
  it can also be sourced from the `PINGACCESS_BASEURL` environment variable.

- **context** (String) This is the PingAccess context path for the admin API, defaults to `/pf-admin-api/v1`
and can be sourced from the `PINGACCESS_CONTEXT` environment variable.
