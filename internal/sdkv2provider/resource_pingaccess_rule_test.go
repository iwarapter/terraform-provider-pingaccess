package sdkv2provider

import (
	"fmt"
	"testing"

	"github.com/iwarapter/pingaccess-sdk-go/v62/services/rules"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func init() {
	resource.AddTestSweepers("rules", &resource.Sweeper{
		Name:         "rules",
		Dependencies: []string{"rulesets"},
		F: func(r string) error {
			svc := rules.New(conf)
			results, _, err := svc.GetRulesCommand(&rules.GetRulesCommandInput{Filter: "acctest_"})
			if err != nil {
				return fmt.Errorf("unable to list rules to sweep %s", err)
			}
			for _, item := range results.Items {
				_, err = svc.DeleteRuleCommand(&rules.DeleteRuleCommandInput{Id: item.Id.String()})
				if err != nil {
					return fmt.Errorf("unable to sweep rules %s because %s", item.Id.String(), err)
				}
			}
			return nil
		},
	})
}

func TestAccPingAccessRule(t *testing.T) {
	resourceName := "pingaccess_rule.acc_test_rule"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV5ProviderFactories: testAccProviders,
		CheckDestroy:             testAccCheckPingAccessRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingAccessRuleConfig("404"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessRuleExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", "acctest_test"),
					resource.TestCheckResourceAttr(resourceName, "class_name", "com.pingidentity.pa.policy.CIDRPolicyInterceptor"),
					resource.TestCheckResourceAttr(resourceName, "supported_destinations.0", "Agent"),
					resource.TestCheckResourceAttr(resourceName, "supported_destinations.1", "Site"),
				),
			},
			{
				Config: testAccPingAccessRuleConfig("403"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessRuleExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", "acctest_test"),
					resource.TestCheckResourceAttr(resourceName, "class_name", "com.pingidentity.pa.policy.CIDRPolicyInterceptor"),
					resource.TestCheckResourceAttr(resourceName, "supported_destinations.0", "Agent"),
					resource.TestCheckResourceAttr(resourceName, "supported_destinations.1", "Site"),
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

func testAccCheckPingAccessRuleDestroy(s *terraform.State) error {
	return nil
}

func testAccPingAccessRuleConfig(configUpdate string) string {
	return fmt.Sprintf(`
	resource "pingaccess_rule" "acc_test_rule" {
		class_name = "com.pingidentity.pa.policy.CIDRPolicyInterceptor"
		name = "acctest_test"
		supported_destinations = [
			"Site",
			"Agent"
		]
		configuration = <<EOF
		{
			"cidrNotation": "127.0.0.${pingaccess_virtualhost.unknown_value.id}/32",
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
	}

	resource "pingaccess_virtualhost" "unknown_value" {
	   host                         = "acctest-rule-config-dynamic-config"
	   port                         = 1111
	   agent_resource_cache_ttl     = 900
	   key_pair_id                  = 0
	   trusted_certificate_group_id = 0
	}`, configUpdate)
}

func testAccCheckPingAccessRuleExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" || rs.Primary.ID == "0" {
			return fmt.Errorf("No rule ID is set")
		}

		conn := testAccProvider.Meta().(paClient).Rules
		result, _, err := conn.GetRuleCommand(&rules.GetRuleCommandInput{
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
