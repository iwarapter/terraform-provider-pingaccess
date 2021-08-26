resource "pingaccess_trusted_certificate_group" "example" {
  name                        = "example"
  use_java_trust_store        = true
  skip_certificate_date_check = false
  cert_ids                    = [pingaccess_certificate.example.id]
}
