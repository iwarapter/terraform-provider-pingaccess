package pingaccess

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/iwarapter/pingaccess-sdk-go/pingaccess"
)

func init() {
	resource.AddTestSweepers("pingaccess_ruleset", &resource.Sweeper{
		Name: "pingaccess_ruleset",
		F:    testSweepRuleSets,
	})
}

func testSweepRuleSets(r string) error {
	url, _ := url.Parse("https://localhost:9000")
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	conn := pingaccess.NewClient("Administrator", "2Access2", url, "/pa-admin-api/v3", nil).Rulesets
	result, _, _ := conn.GetRuleSetsCommand(&pingaccess.GetRuleSetsCommandInput{Filter: "acc_test_"})
	for _, v := range result.Items {
		conn.DeleteRuleSetCommand(&pingaccess.DeleteRuleSetCommandInput{Id: v.Id.String()})
	}
	return nil
}

func TestAccPingAccessRuleSet(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		// PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPingAccessRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingAccessRuleSetConfig("SuccessIfAllSucceed"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessRuleSetExists("pingaccess_ruleset.ruleset_one"),
					testAccCheckPingAccessRuleSetAttributes("pingaccess_ruleset.ruleset_one"),
				),
			},
			{
				Config: testAccPingAccessRuleSetConfig("SuccessIfAnyOneSucceeds"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessRuleSetExists("pingaccess_ruleset.ruleset_one"),
					testAccCheckPingAccessRuleSetAttributes("pingaccess_ruleset.ruleset_one"),
				),
			},
		},
	})
}

func testAccCheckPingAccessRuleSetDestroy(s *terraform.State) error {
	return nil
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
		rs, _ := s.RootModule().Resources[n]

		conn := testAccProvider.Meta().(*pingaccess.Client).Rulesets
		result, _, err := conn.GetRuleSetCommand(&pingaccess.GetRuleSetCommandInput{
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

		conn := testAccProvider.Meta().(*pingaccess.Client).Rulesets
		result, _, err := conn.GetRuleSetCommand(&pingaccess.GetRuleSetCommandInput{
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
