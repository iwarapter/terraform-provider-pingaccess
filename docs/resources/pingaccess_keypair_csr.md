# Resource: pingaccess_keypair_csr

Provides a keypair csr response.

## Example Usage

### Signing a CSR with an example tls signer
```hcl
resource "pingaccess_keypair" "demo_generate" {
  alias             = "demo2"
  city              = "London"
  common_name       = "Example"
  country           = "GB"
  key_algorithm     = "RSA"
  key_size          = 2048
  organization      = "Test"
  organization_unit = "Development"
  state             = "London"
  valid_days        = 365
}


data "pingaccess_keypair_csr" "csr" {
  id = pingaccess_keypair.demo_generate.id
}

resource "pingaccess_keypair_csr" "test" {
  keypair_id                   = pingaccess_keypair.demo_generate.id
  file_data                    = base64encode(tls_locally_signed_cert.example.cert_pem)
  chain_certificates           = [base64encode(tls_self_signed_cert.example.cert_pem)]
  trusted_certificate_group_id = data.pingaccess_trusted_certificate_group.trust_any.id
}

resource "tls_private_key" "example" {
  algorithm = "RSA"
  rsa_bits  = 2048
}

resource "tls_locally_signed_cert" "example" {
  cert_request_pem   = data.pingaccess_keypair_csr.csr.cert_request_pem
  ca_key_algorithm   = "RSA"
  ca_private_key_pem = tls_private_key.example.private_key_pem
  ca_cert_pem        = tls_self_signed_cert.example.cert_pem

  validity_period_hours = 12

  allowed_uses = [
    "key_encipherment",
    "digital_signature",
    "server_auth",
  ]
}

resource "tls_self_signed_cert" "example" {
  key_algorithm   = "RSA"
  private_key_pem = tls_private_key.example.private_key_pem

  subject {
    common_name  = "example.com"
    organization = "ACME Examples, Inc"
  }

  validity_period_hours = 12

  allowed_uses = [
    "key_encipherment",
    "digital_signature",
    "server_auth",
  ]
}
```

## Argument Attributes
The following arguments are supported:

- [`keypair_id`](#keypair_id) - (required) The Id for the key pair.
- [`file_data`](#file_data) - (required) The base64-encoded data of the keypair.
- [`chain_certificates`](#chain_certificates) - A list of base64-encoded certificates to add to the key pair certificate chain.
- [`trusted_certificate_group_id`](#trusted_certificate_group_id) - The ID of the trusted certificate group associated with the CSR response.

## Attributes Reference

None
