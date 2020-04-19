resource "pingaccess_certificate" "demo_certificate" {
  alias     = "demo"
  file_data = base64encode(file("test_cases/amazon_root_ca1.pem"))
}
