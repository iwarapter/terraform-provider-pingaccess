# Data Source: pingaccess_keypair_csr

Use this data source to get the CSR of a keypair in Ping Access.

## Example Usage
```hcl
data "pingaccess_keypair_csr" "csr" {
  id = pingaccess_keypair.demo_generate.id
}
```
## Argument Attributes
The following arguments are supported:

- [`id`](#id) - (required) The ID for the keypair.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

- [`cert_request_pem`](#cert_request_pem) - The keypairs's Certificate Signing Response.
