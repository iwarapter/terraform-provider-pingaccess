#Generating a Keypair
resource "pingaccess_keypair" "example_generate" {
  alias             = "example1"
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

#Importing a Keypair
resource "pingaccess_keypair" "example_keypair" {
  alias     = "example"
  file_data = filebase64("provider.p12")
  password  = "password"
}
