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
	resource.AddTestSweepers("pingaccess_rule", &resource.Sweeper{
		Name:         "pingaccess_rule",
		F:            testSweepRules,
		Dependencies: []string{"pingaccess_ruleset"},
	})
}

func testSweepRules(r string) error {
	url, _ := url.Parse("https://localhost:9000")
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	conn := pingaccess.NewClient("Administrator", "2Access2", url, "/pa-admin-api/v3", nil).Rules
	result, _, _ := conn.GetRulesCommand(&pingaccess.GetRulesCommandInput{Filter: "acc_test_"})
	for _, v := range result.Items {
		conn.DeleteRuleCommand(&pingaccess.DeleteRuleCommandInput{Id: v.Id.String()})
	}
	return nil
}

func TestAccPingAccessRule(t *testing.T) {
	var out pingaccess.RuleView

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPingAccessRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingAccessRuleConfig("bar", "404"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessRuleExists("pingaccess_rule.acc_test_rule_bar", 3, &out),
					// testAccCheckPingAccessRuleAttributes(),
					// testAccCheckAWSPolicyAttachmentAttributes([]string{userName}, []string{roleName}, []string{groupName}, &out),
				),
			},
			{
				Config: testAccPingAccessRuleConfig("bar", "403"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessRuleExists("pingaccess_rule.acc_test_rule_bar", 6, &out),
					// testAccCheckAWSPolicyAttachmentAttributes([]string{userName2, userName3},
					// 	[]string{roleName2, roleName3}, []string{groupName2, groupName3}, &out),
				),
			},
		},
	})
}

func testAccCheckPingAccessRuleDestroy(s *terraform.State) error {
	return nil
}

func testAccPingAccessRuleConfig(name, configUpdate string) string {
	return fmt.Sprintf(`
	resource "pingaccess_rule" "acc_test_rule_%s" {
		class_name = "com.pingidentity.pa.policy.CIDRPolicyInterceptor"
		name = "%s"
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
			"errorResponseCode": %s,
			"errorResponseStatusMsg": "Forbidden",
			"errorResponseTemplateFile": "policy.error.page.template.html",
			"errorResponseContentType": "text/html;charset=UTF-8",
			"rejectionHandler": null,
			"rejectionHandlingEnabled": false
		}
		EOF
	}`, name, name, configUpdate)
}

func testAccCheckPingAccessRuleExists(n string, c int64, out *pingaccess.RuleView) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" || rs.Primary.ID == "0" {
			return fmt.Errorf("No rule ID is set")
		}

		conn := testAccProvider.Meta().(*pingaccess.Client).Rules
		result, _, err := conn.GetRuleCommand(&pingaccess.GetRuleCommandInput{
			Id: rs.Primary.ID,
		})

		if err != nil {
			return fmt.Errorf("Error: Rule (%s) not found", n)
		}

		if *result.Name != rs.Primary.Attributes["name"] {
			return fmt.Errorf("Error: Rule response (%s) didnt match state (%s)", *result.Name, rs.Primary.Attributes["name"])
		}

		return nil
	}
}
