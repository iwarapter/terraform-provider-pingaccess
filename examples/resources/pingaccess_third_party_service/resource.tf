resource "pingaccess_third_party_service" "example" {
  name    = "example"
  targets = ["server.domain:1234"]
}
