
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
