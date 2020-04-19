The PingAccess provider is used to interact with the many resources supported by the PingAccess admin API. The provider needs to be configured with the proper credentials before it can be used.

Use the navigation to the left to read about the available resources.


### Example Usage
```terraform
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

### Authentication

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
