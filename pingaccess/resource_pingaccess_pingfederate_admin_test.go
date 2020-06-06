package pingaccess

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pa "github.com/iwarapter/pingaccess-sdk-go/pingaccess"
)

func TestAccPingAccessPingFederateAdmin(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPingAccessPingFederateAdminDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingAccessPingFederateAdminConfig(os.Getenv("PINGFEDERATE_TEST_IP"), true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessPingFederateAdminExists("pingaccess_pingfederate_admin.demo"),
				),
			},
			{
				Config: testAccPingAccessPingFederateAdminConfig(os.Getenv("PINGFEDERATE_TEST_IP"), true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessPingFederateAdminExists("pingaccess_pingfederate_admin.demo"),
				),
			},
		},
	})
}

func testAccCheckPingAccessPingFederateAdminDestroy(s *terraform.State) error {
	return nil
}

func testAccPingAccessPingFederateAdminConfig(issuer string, proxy bool) string {
	return fmt.Sprintf(`
	resource "pingaccess_pingfederate_admin" "demo" {
		admin_username = "Administrator"
		admin_password {
			value = "2FederateM0re"
		}
		base_path = "/path"
		audit_level = "ON"
		host = "%s"
		port = 9031
		secure = false
		trusted_certificate_group_id = 2
		use_proxy = %v
	}`, issuer, proxy)
}

func testAccCheckPingAccessPingFederateAdminExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" || rs.Primary.ID == "0" {
			return fmt.Errorf("No third party service ID is set")
		}

		conn := testAccProvider.Meta().(*pa.Client).PingFederate
		result, _, err := conn.GetPingFederateAdminCommand()

		if err != nil {
			return fmt.Errorf("Error: PingFederateAdmin (%s) not found", n)
		}

		if *result.Host != rs.Primary.Attributes["host"] {
			return fmt.Errorf("Error: PingFederateAdmin response (%s) didnt match state (%s)", *result.Host, rs.Primary.Attributes["host"])
		}

		return nil
	}
}

func Test_resourcePingAccessPingFederateAdminReadData(t *testing.T) {
	cases := []struct {
		PingFederateAdmin pa.PingFederateAdminView
	}{
		{
			PingFederateAdmin: pa.PingFederateAdminView{
				AdminPassword: &pa.HiddenFieldView{
					Value: String("secret"),
				},
				AdminUsername: String("admin"),
				Host:          String("localhost"),
				Port:          Int(9031),
			},
		},
		{
			PingFederateAdmin: pa.PingFederateAdminView{
				AdminPassword: &pa.HiddenFieldView{
					Value: String("secret"),
				},
				AdminUsername:             String("admin"),
				Host:                      String("localhost"),
				Port:                      Int(9031),
				TrustedCertificateGroupId: Int(0),
				UseProxy:                  Bool(true),
				Secure:                    Bool(true),
				AuditLevel:                String("ON"),
				BasePath:                  String("/"),
			},
		},
	}
	for i, tc := range cases {
		t.Run(fmt.Sprintf("tc:%v", i), func(t *testing.T) {

			resourceSchema := resourcePingAccessPingFederateAdminSchema()
			resourceLocalData := schema.TestResourceDataRaw(t, resourceSchema, map[string]interface{}{})
			resourcePingAccessPingFederateAdminReadResult(resourceLocalData, &tc.PingFederateAdmin)

			if got := *resourcePingAccessPingFederateAdminReadData(resourceLocalData); !cmp.Equal(got, tc.PingFederateAdmin) {
				t.Errorf("resourcePingAccessPingFederateAdminReadData() = %v", cmp.Diff(got, tc.PingFederateAdmin))
			}
		})
	}
}
