resource "pingaccess_certificate" "example" {
  alias     = "example"
  file_data = filebase64("amazon_root_ca1.pem")
}
