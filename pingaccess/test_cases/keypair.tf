resource "pingaccess_keypair" "demo_keypair" {
  alias     = "demo"
  file_data = filebase64("test_cases/provider.p12")
  password  = "top_secret"
}
