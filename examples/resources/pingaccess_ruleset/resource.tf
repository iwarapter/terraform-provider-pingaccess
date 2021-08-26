resource "pingaccess_ruleset" "example" {
  name             = "example"
  success_criteria = "SuccessIfAllSucceed"
  element_type     = "Rule"

  policy = [
    pingaccess_rule.example_1.id,
    pingaccess_rule.example_2.id,
  ]
}
