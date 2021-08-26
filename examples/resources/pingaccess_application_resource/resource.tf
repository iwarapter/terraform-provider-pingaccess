resource "pingaccess_application_resource" "app_res_test_root_resource" {
  name           = "Root Resource"
  methods        = ["*"]
  path_prefixes  = ["/*"]
  audit_level    = "ON"
  anonymous      = false
  enabled        = true
  root_resource  = true
  application_id = pingaccess_application.example.id

  path_patterns {
    pattern = "/*"
    type    = "WILDCARD"
  }

  policy {
    web {
      type = "Rule"
      id   = pingaccess_rule.example_one.id
    }
    web {
      type = "Rule"
      id   = pingaccess_rule.example_two.id
    }
  }
}
