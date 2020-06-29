resource "pingaccess_site" "demo" {
  name    = "demo-site"
  targets = ["localhost:1234"]
}
