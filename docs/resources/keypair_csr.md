---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "pingaccess_keypair_csr Resource - terraform-provider-pingaccess"
subcategory: ""
description: |-
  Provides configuration for Keypair CSRs within PingAccess.
---

# pingaccess_keypair_csr (Resource)

Provides configuration for Keypair CSRs within PingAccess.

## Example Usage

```terraform
# Generate our keypair
resource "pingaccess_keypair" "example_generate" {
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

# Get the generated keypair CSR
data "pingaccess_keypair_csr" "csr_example" {
  id = pingaccess_keypair.example_generate.id
}

# Import the signed certificate for our keypair
resource "pingaccess_keypair_csr" "test" {
  keypair_id                   = pingaccess_keypair.example_generate.id
  file_data                    = base64encode(tls_locally_signed_cert.example.cert_pem)
  chain_certificates           = [base64encode(tls_self_signed_cert.example.cert_pem)]
  trusted_certificate_group_id = data.pingaccess_trusted_certificate_group.trust_any.id
}

resource "tls_private_key" "example" {
  algorithm = "RSA"
  rsa_bits  = 2048
}

resource "tls_locally_signed_cert" "example" {
  cert_request_pem   = data.pingaccess_keypair_csr.csr_example.cert_request_pem
  ca_key_algorithm   = "RSA"
  ca_private_key_pem = tls_private_key.example.private_key_pem
  ca_cert_pem        = tls_self_signed_cert.example.cert_pem

  validity_period_hours = 12

  allowed_uses = ["key_encipherment", "digital_signature", "server_auth"]
}

resource "tls_self_signed_cert" "example" {
  key_algorithm   = "RSA"
  private_key_pem = tls_private_key.example.private_key_pem

  subject {
    common_name  = "example.com"
    organization = "ACME Examples, Inc"
  }

  validity_period_hours = 12

  allowed_uses = ["key_encipherment", "digital_signature", "server_auth"]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `file_data` (String) The CSR response data.
- `keypair_id` (String) ID of the Key Pair to update.

### Optional

- `chain_certificates` (List of String) A list of base64-encoded certificates to add to the key pair certificate chain.
- `trusted_certificate_group_id` (Number) The ID of the trusted certificate group associated with the CSR response.

### Read-Only

- `id` (String) The ID of this resource.
