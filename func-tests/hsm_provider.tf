resource "pingaccess_hsm_provider" "test" {
  count         = var.pa6 ? 1 : 0
  class_name    = "com.pingidentity.pa.hsm.pkcs11.plugin.PKCS11HsmProvider"
  name          = "demo"
  configuration = <<EOF
  {
    "slotId": "1234",
    "library": "foo",
    "password": "top_secret"
  }
  EOF
}
