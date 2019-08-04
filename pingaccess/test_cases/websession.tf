resource "pingaccess_websession" "demo_session" {
  name     = "demo-session"
  audience = "example"

  client_credentials {
    client_id = "websession"

    client_secret {
      value = "changeme"
    }
  }

  scopes = [
    "profile",
    "address",
    "email",
    "phone",
  ]
}
