resource "pingaccess_application" "demo_application" {
  application_type  = "Web"
  name              = "demo"
  context_root      = "/"
  default_auth_type = "Web"
  destination       = "Site"
  site_id           = pingaccess_site.demo_site.id
  virtual_host_ids  = [pingaccess_virtualhost.demo_virtualhost.id]

  policy {
    web {
      type = "Rule"
      id   = pingaccess_rule.demo_rule.id
    }
  }
}
