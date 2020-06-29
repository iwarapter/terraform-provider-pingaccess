resource "pingaccess_application" "demo_application" {
  application_type = "Web"
  name             = "demo"
  context_root     = "/"
  destination      = "Site"
  site_id          = pingaccess_site.demo.id
  virtual_host_ids = [pingaccess_virtualhost.demo.id]

  policy {
    web {
      type = "Rule"
      id   = pingaccess_rule.demo_rule.id
    }
  }
}
