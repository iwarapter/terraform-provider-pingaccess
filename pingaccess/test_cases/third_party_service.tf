resource "pingaccess_third_party_service" "demo_third_party_service" {
  name = "demo"

  targets = [
    "server.domain:1234",
  ]
}
