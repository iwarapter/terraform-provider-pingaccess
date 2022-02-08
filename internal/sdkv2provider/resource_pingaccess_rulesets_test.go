package sdkv2provider

import (
	"fmt"
	"testing"

	"github.com/iwarapter/pingaccess-sdk-go/v62/services/rulesets"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func init() {
	resource.AddTestSweepers("rulesets", &resource.Sweeper{
		Name: "rulesets",
		F: func(r string) error {
			svc := rulesets.New(conf)
			results, _, err := svc.GetRuleSetsCommand(&rulesets.GetRuleSetsCommandInput{Filter: "acctest_"})
			if err != nil {
				return fmt.Errorf("unable to list rulesets to sweep %s", err)
			}
			for _, item := range results.Items {
				_, err = svc.DeleteRuleSetCommand(&rulesets.DeleteRuleSetCommandInput{Id: item.Id.String()})
				if err != nil {
					return fmt.Errorf("unable to sweep ruleset %s because %s", item.Id.String(), err)
				}
			}
			return nil
		},
	})
}

func TestAccPingAccessRuleSet(t *testing.T) {
	resourceName := "pingaccess_ruleset.ruleset_one"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV5ProviderFactories: testAccProviders,
		CheckDestroy:             testAccCheckPingAccessRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingAccessRuleSetConfig("SuccessIfAllSucceed"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessRuleSetExists(resourceName),
					testAccCheckPingAccessRuleSetAttributes(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", "acctest_ruleset-one"),
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
					resource.TestCheckResourceAttr(resourceName, "name", "acctest_ruleset-one"),
					resource.TestCheckResourceAttr(resourceName, "success_criteria", "SuccessIfAnyOneSucceeds"),
					resource.TestCheckResourceAttr(resourceName, "element_type", "Rule"),
					resource.TestCheckResourceAttrSet(resourceName, "policy.0"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccPingAccessRuleSetConfig(configUpdate string) string {
	return fmt.Sprintf(`
	resource "pingaccess_ruleset" "ruleset_one" {
		name             = "acctest_ruleset-one"
		success_criteria = "%s"
		element_type     = "Rule"
		policy = [pingaccess_rule.ruleset_rule_one.id]
	}

	resource "pingaccess_rule" "ruleset_rule_one" {
		class_name = "com.pingidentity.pa.policy.CIDRPolicyInterceptor"
		name = "acctest_ruleset-rule-one"
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

		conn := testAccProvider.Meta().(paClient).Rulesets
		result, _, err := conn.GetRuleSetCommand(&rulesets.GetRuleSetCommandInput{
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

		conn := testAccProvider.Meta().(paClient).Rulesets
		result, _, err := conn.GetRuleSetCommand(&rulesets.GetRuleSetCommandInput{
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
