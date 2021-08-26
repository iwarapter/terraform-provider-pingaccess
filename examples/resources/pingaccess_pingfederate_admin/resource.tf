resource "pingaccess_pingfederate_admin" "example" {
  admin_username = "oauth"
  admin_password {
    value = "top_secret"
  }
  audit_level                  = "ON"
  base_path                    = "/path"
  host                         = "localhost"
  port                         = 9031
  secure                       = true
  trusted_certificate_group_id = 2
  use_proxy                    = true
}
