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
