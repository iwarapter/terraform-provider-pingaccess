resource "pingaccess_trusted_certificate_group" "demo_certificate_group" {
  name                        = "demo_certificate_group"
  use_java_trust_store        = true
  skip_certificate_date_check = false
}
