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
    pattern = "%s"
    type    = "WILDCARD"
  }

  path_prefixes = [
    "/as/token.oauth2",
    "%s",
  ]

  audit_level    = "OFF"
  anonymous      = false
  enabled        = true
  root_resource  = false
  application_id = "${pingaccess_application.demo_application.id}"

  policy {
    web {
      type = "Rule"
      id   = "${pingaccess_rule.demo_rule_one.id}"
    }

    web {
      type = "Rule"
      id   = "${pingaccess_rule.demo_rule_two.id}"
    }
  }
}
