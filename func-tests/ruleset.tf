resource "pingaccess_ruleset" "demo_ruleset" {
  name             = "demo_ruleset"
  success_criteria = "SuccessIfAllSucceed"
  element_type     = "Rule"

  policy = [
    pingaccess_rule.demo_1.id,
    pingaccess_rule.demo_2.id,
  ]
}
