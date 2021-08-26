resource "pingaccess_websession" "example" {
  name     = "example"
  audience = "demo"
  client_credentials {
    client_id = "websession"
    client_secret {
      value = "changeme"
    }
  }
  scopes = ["profile", "address", "email", "phone"]
}
