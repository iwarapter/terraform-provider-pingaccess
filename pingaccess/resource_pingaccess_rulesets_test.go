package pingaccess

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	pa "github.com/iwarapter/pingaccess-sdk-go/pingaccess"
)

func TestAccPingAccessRuleSet(t *testing.T) {
	resourceName := "pingaccess_ruleset.ruleset_one"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPingAccessRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingAccessRuleSetConfig("SuccessIfAllSucceed"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessRuleSetExists(resourceName),
					testAccCheckPingAccessRuleSetAttributes(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", "tf-acc-ruleset-one"),
					resource.TestCheckResourceAttr(resourceName, "success_criteria", "SuccessIfAllSucceed"),
					resource.TestCheckResourceAttr(resourceName, "element_type", "Rule"),
					resource.TestCheckResourceAttrSet(resourceName, "policy.0"),
				),
			},
			{
				Config: testAccPingAccessRuleSetConfig("SuccessIfAnyOneSucceeds"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessRuleSetExists(resourceName),
					testAccCheckPingAccessRuleSetAttributes(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", "tf-acc-ruleset-one"),
					resource.TestCheckResourceAttr(resourceName, "success_criteria", "SuccessIfAnyOneSucceeds"),
					resource.TestCheckResourceAttr(resourceName, "element_type", "Rule"),
					resource.TestCheckResourceAttrSet(resourceName, "policy.0"),
				),
			},
		},
	})
}

func testAccPingAccessRuleSetConfig(configUpdate string) string {
	return fmt.Sprintf(`
	resource "pingaccess_ruleset" "ruleset_one" {
		name             = "tf-acc-ruleset-one"
		success_criteria = "%s"
		element_type     = "Rule"
		policy = [
			"${pingaccess_rule.ruleset_rule_one.id}"
		]
	}

	resource "pingaccess_rule" "ruleset_rule_one" {
		class_name = "com.pingidentity.pa.policy.CIDRPolicyInterceptor"
		name = "tf-acc-ruleset-rule-one"
		supported_destinations = [
			"Site",
			"Agent"
		]
		configuration = <<EOF
		{
			"cidrNotation": "127.0.0.1/32",
			"negate": false,
			"overrideIpSource": false,
			"headers": [],
			"headerValueLocation": "LAST",
			"fallbackToLastHopIp": true,
			"errorResponseCode": 403,
			"errorResponseStatusMsg": "Forbidden",
			"errorResponseTemplateFile": "policy.error.page.template.html",
			"errorResponseContentType": "text/html;charset=UTF-8",
			"rejectionHandler": null,
			"rejectionHandlingEnabled": false
		}
		EOF
	}`, configUpdate)
}

func testAccCheckPingAccessRuleSetAttributes(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs := s.RootModule().Resources[n]

		conn := testAccProvider.Meta().(*pa.Client).Rulesets
		result, _, err := conn.GetRuleSetCommand(&pa.GetRuleSetCommandInput{
			Id: rs.Primary.ID,
		})

		if err != nil {
			return fmt.Errorf("Error: RuleSet (%s) not found", n)
		}

		if *result.SuccessCriteria != rs.Primary.Attributes["success_criteria"] {
			return fmt.Errorf("Error: RuleSet response (%s) didnt match state (%s)", *result.SuccessCriteria, rs.Primary.Attributes["success_criteria"])
		}

		return nil
	}
}

func testAccCheckPingAccessRuleSetExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" || rs.Primary.ID == "0" {
			return fmt.Errorf("No RuleSet ID is set")
		}

		conn := testAccProvider.Meta().(*pa.Client).Rulesets
		result, _, err := conn.GetRuleSetCommand(&pa.GetRuleSetCommandInput{
			Id: rs.Primary.ID,
		})

		if err != nil {
			return fmt.Errorf("Error: RuleSet (%s) not found", n)
		}

		if *result.Name != rs.Primary.Attributes["name"] {
			return fmt.Errorf("Error: Rule response (%s) didnt match state (%s)", *result.Name, rs.Primary.Attributes["name"])
		}

		return nil
	}
}
