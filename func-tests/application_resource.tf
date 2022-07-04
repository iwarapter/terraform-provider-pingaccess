resource "pingaccess_application_resource" "demo_application_resource" {
  name = "demo_resource"

  methods = [
    "*",
  ]

  path_patterns {
    pattern = "/as/token.oauth2"
    type    = "WILDCARD"
  }

  path_patterns {
    pattern = "/foo"
    type    = "WILDCARD"
  }

  audit_level    = "OFF"
  anonymous      = false
  enabled        = true
  root_resource  = false
  application_id = pingaccess_application.demo_application.id

  policy {
    web {
      type = "Rule"
      id   = pingaccess_rule.demo_1.id
    }

    web {
      type = "Rule"
      id   = pingaccess_rule.demo_2.id
    }
  }
}
