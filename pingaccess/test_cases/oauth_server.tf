resource "pingaccess_oauth_server" "demo_oauth_server" {
  targets                      = ["localhost:9031"]
  subject_attribute_name       = "san"
  trusted_certificate_group_id = 1
  introspection_endpoint       = "https://localhost:443/introspection"

  client_credentials {
    client_id = "oauth"

    client_secret {
      value = "top_secret"
    }
  }

  secure = true
}
