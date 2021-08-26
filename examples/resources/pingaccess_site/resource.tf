resource "pingaccess_site" "example" {
  name    = "example"
  targets = ["localhost:1234"]
}
