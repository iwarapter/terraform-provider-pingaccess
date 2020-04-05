package pingaccess

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/iwarapter/pingaccess-sdk-go/pingaccess"
)

func TestAccPingAccessIdentityMapping(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPingAccessIdentityMappingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingAccessIdentityMappingConfig("bar", "SUB_HEADER"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessIdentityMappingExists("pingaccess_identity_mapping.acc_test_idm_bar"),
					testAccCheckPingAccessIdentityMappingAttributes("pingaccess_identity_mapping.acc_test_idm_bar"),
				),
			},
			{
				Config: testAccPingAccessIdentityMappingConfig("bar", "SUB"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessIdentityMappingExists("pingaccess_identity_mapping.acc_test_idm_bar"),
					testAccCheckPingAccessIdentityMappingAttributes("pingaccess_identity_mapping.acc_test_idm_bar"),
				),
			},
		},
	})
}

func testAccCheckPingAccessIdentityMappingDestroy(s *terraform.State) error {
	return nil
}

func testAccPingAccessIdentityMappingConfig(name, configUpdate string) string {
	return fmt.Sprintf(`
	resource "pingaccess_identity_mapping" "acc_test_idm_%s" {
		class_name = "com.pingidentity.pa.identitymappings.HeaderIdentityMapping"
		name = "%s"
		configuration = <<EOF
		{
			"attributeHeaderMappings": [
				{
					"subject": true,
					"attributeName": "sub",
					"headerName": "%s"
				}
			],
			"headerClientCertificateMappings": []
		}
		EOF
	}`, name, name, configUpdate)
}

func testAccCheckPingAccessIdentityMappingExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" || rs.Primary.ID == "0" {
			return fmt.Errorf("No IdentityMapping ID is set")
		}

		conn := testAccProvider.Meta().(*pingaccess.Client).IdentityMappings
		result, _, err := conn.GetIdentityMappingCommand(&pingaccess.GetIdentityMappingCommandInput{
			Id: rs.Primary.ID,
		})

		if err != nil {
			return fmt.Errorf("Error: IdentityMapping (%s) not found", n)
		}

		if *result.Name != rs.Primary.Attributes["name"] {
			return fmt.Errorf("Error: IdentityMapping response (%s) didnt match state (%s)", *result.Name, rs.Primary.Attributes["name"])
		}

		return nil
	}
}

func testAccCheckPingAccessIdentityMappingAttributes(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, _ := s.RootModule().Resources[n]
		if rs.Primary.ID == "" || rs.Primary.ID == "0" {
			return fmt.Errorf("No IdentityMapping ID is set")
		}

		conn := testAccProvider.Meta().(*pingaccess.Client).IdentityMappings
		result, _, err := conn.GetIdentityMappingCommand(&pingaccess.GetIdentityMappingCommandInput{
			Id: rs.Primary.ID,
		})

		if err != nil {
			return fmt.Errorf("Error: IdentityMapping (%s) not found", n)
		}

		if *result.Name != rs.Primary.Attributes["name"] {
			return fmt.Errorf("Error: IdentityMapping response (%s) didnt match state (%s)", *result.Name, rs.Primary.Attributes["name"])
		}

		resultMapping := result.Configuration["attributeHeaderMappings"].([]interface{})[0].(map[string]interface{})["headerName"]
		var x map[string]interface{}
		json.Unmarshal([]byte(rs.Primary.Attributes["configuration"]), &x)
		stateMapping := x["attributeHeaderMappings"].([]interface{})[0].(map[string]interface{})["headerName"]

		if resultMapping != stateMapping {
			return fmt.Errorf("Error: IdentityMapping response (%s) didnt match state (%s)", resultMapping, stateMapping)
		}

		return nil
	}
}
