resource "pingaccess_ruleset" "demo_ruleset" {
  name             = "demo_ruleset"
  success_criteria = "SuccessIfAllSucceed"
  element_type     = "Rule"

  policy = [
    pingaccess_rule.demo_rule.id,
    pingaccess_rule.second_demo_rule.id,
  ]
}
