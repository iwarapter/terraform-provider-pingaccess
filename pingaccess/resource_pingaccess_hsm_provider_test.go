package pingaccess

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	pa "github.com/iwarapter/pingaccess-sdk-go/pingaccess"
)

func TestAccPingAccessHsmProvider(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccPingAccessHsmProviderConfig("bar", "foo"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessHsmProviderExists("pingaccess_hsm_provider.acc_test_idm_bar"),
					testAccCheckPingAccessHsmProviderAttributes("pingaccess_hsm_provider.acc_test_idm_bar", "foo"),
				),
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
				//TODO The pingaccess AWS CloudHsm can be created but not deleted without a working HSM,
				// and the current TestStep is unable to plan/apply without attempting to destroy atm it seems.
			},
		},
	})
}

func testAccPingAccessHsmProviderConfig(name, configUpdate string) string {
	return fmt.Sprintf(`
	resource "pingaccess_hsm_provider" "acc_test_idm_%s" {
		class_name = "com.pingidentity.pa.hsm.cloudhsm.plugin.AwsCloudHsmProvider"
		name = "%s"
		configuration = <<EOF
		{
			"user": true,
			"password": "sub",
			"partition": "%s"
		}
		EOF
	}`, name, name, configUpdate)
}

func testAccCheckPingAccessHsmProviderExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" || rs.Primary.ID == "0" {
			return fmt.Errorf("No HsmProvider ID is set")
		}

		conn := testAccProvider.Meta().(*pa.Client).HsmProviders
		result, _, err := conn.GetHsmProviderCommand(&pa.GetHsmProviderCommandInput{
			Id: rs.Primary.ID,
		})

		if err != nil {
			return fmt.Errorf("Error: HsmProvider (%s) not found", n)
		}

		if *result.Name != rs.Primary.Attributes["name"] {
			return fmt.Errorf("Error: HsmProvider response (%s) didnt match state (%s)", *result.Name, rs.Primary.Attributes["name"])
		}

		return nil
	}
}

func testAccCheckPingAccessHsmProviderAttributes(n, partition string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs := s.RootModule().Resources[n]
		if rs.Primary.ID == "" || rs.Primary.ID == "0" {
			return fmt.Errorf("No HsmProvider ID is set")
		}

		conn := testAccProvider.Meta().(*pa.Client).HsmProviders
		result, _, err := conn.GetHsmProviderCommand(&pa.GetHsmProviderCommandInput{
			Id: rs.Primary.ID,
		})

		if err != nil {
			return fmt.Errorf("Error: HsmProvider (%s) not found", n)
		}

		if *result.Name != rs.Primary.Attributes["name"] {
			return fmt.Errorf("Error: HsmProvider response (%s) didnt match state (%s)", *result.Name, rs.Primary.Attributes["name"])
		}

		resultMapping := result.Configuration["partition"].(string)
		if resultMapping != partition {
			return fmt.Errorf("Error: HsmProvider response (%s) didnt match state (%s)", resultMapping, partition)
		}

		return nil
	}
}
