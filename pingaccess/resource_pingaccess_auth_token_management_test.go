package pingaccess

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	pa "github.com/iwarapter/pingaccess-sdk-go/pingaccess"
)

func TestAccPingAccessAuthTokenManagement(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPingAccessAuthTokenManagementDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingAccessAuthTokenManagementConfig("PingAccessAuthToken"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessAuthTokenManagementExists("pingaccess_auth_token_management.demo"),
				),
			},
			{
				Config: testAccPingAccessAuthTokenManagementConfig("PingAccessAuthToken2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessAuthTokenManagementExists("pingaccess_auth_token_management.demo"),
				),
			},
		},
	})
}

func testAccCheckPingAccessAuthTokenManagementDestroy(s *terraform.State) error {
	return nil
}

func testAccPingAccessAuthTokenManagementConfig(issuer string) string {
	return fmt.Sprintf(`
	resource "pingaccess_auth_token_management" "demo" {
		key_roll_enabled = true
    key_roll_period_in_hours = 24
  	issuer = "%s"
  	signing_algorithm = "P-256"
	}`, issuer)
}

func testAccCheckPingAccessAuthTokenManagementExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" || rs.Primary.ID == "0" {
			return fmt.Errorf("No auth token management ID is set")
		}

		conn := testAccProvider.Meta().(*pa.Client).AuthTokenManagements
		result, _, err := conn.GetAuthTokenManagementCommand()

		if err != nil {
			return fmt.Errorf("Error: AuthTokenManagement (%s) not found", n)
		}

		if *result.Issuer != rs.Primary.Attributes["issuer"] {
			return fmt.Errorf("Error: AuthTokenManagement response (%s) didnt match state (%s)", *result.Issuer, rs.Primary.Attributes["issuer"])
		}

		return nil
	}
}

func Test_resourcePingAccessAuthTokenManagementReadData(t *testing.T) {
	cases := []struct {
		AuthTokenManagementView pa.AuthTokenManagementView
	}{
		{
			AuthTokenManagementView: pa.AuthTokenManagementView{
				Issuer:               String("PingAccessAuthTokenDemo"),
				KeyRollEnabled:       Bool(false),
				KeyRollPeriodInHours: Int(23),
				SigningAlgorithm:     String("P-512"),
			},
		},
	}
	for i, tc := range cases {
		t.Run(fmt.Sprintf("tc:%v", i), func(t *testing.T) {

			resourceSchema := resourcePingAccessAuthTokenManagementSchema()
			resourceLocalData := schema.TestResourceDataRaw(t, resourceSchema, map[string]interface{}{})
			resourcePingAccessAuthTokenManagementReadResult(resourceLocalData, &tc.AuthTokenManagementView)

			if got := *resourcePingAccessAuthTokenManagementReadData(resourceLocalData); !cmp.Equal(got, tc.AuthTokenManagementView) {
				t.Errorf("resourcePingAccessAuthTokenManagementReadData() = %v", cmp.Diff(got, tc.AuthTokenManagementView))
			}

			resourcePingAccessAuthTokenManagementReadResult(resourceLocalData, &tc.AuthTokenManagementView)
		})
	}
}
