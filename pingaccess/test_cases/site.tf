resource "pingaccess_site" "demo_site" {
  name    = "demo-site"
  targets = ["localhost:1234"]
}
