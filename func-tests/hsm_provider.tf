resource "pingaccess_hsm_provider" "test" {
  count         = var.pa6 ? 1 : 0
  name          = "demo"
  class_name    = "com.pingidentity.pa.hsm.pkcs11.plugin.PKCS11HsmProvider"
  configuration = <<EOF
  {
    "slotId": "1234",
    "password": "top_secret",
    "library": "example"
  }
  EOF
}
