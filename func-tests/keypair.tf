resource "pingaccess_keypair" "demo_keypair" {
  alias     = "demo"
  file_data = filebase64("provider.p12")
  password  = "password"
}
