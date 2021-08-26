resource "pingaccess_oauth_server" "oauth_server" {
  targets                      = ["localhost:9031"]
  subject_attribute_name       = "san"
  trusted_certificate_group_id = 1
  introspection_endpoint       = "/introspect"
  secure                       = true
  client_credentials {
    client_id = "oauth"
    client_secret {
      value = "top_secret"
    }
  }
}
