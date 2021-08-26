resource "pingaccess_application" "example" {
  access_validator_id = 0
  application_type    = "Web"
  agent_id            = 0
  name                = "example"
  context_root        = "/example"
  destination         = "Site"
  site_id             = pingaccess_site.example.id
  spa_support_enabled = false
  virtual_host_ids    = [pingaccess_virtualhost.example.id]
  web_session_id      = pingaccess_websession.example.id

  policy {
    web {
      type = "Rule"
      id   = pingaccess_rule.example.id
    }
  }
}
