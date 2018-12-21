package pingaccess

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/iwarapter/pingaccess-sdk-go/pingaccess"
)

func TestAccPingAccessSite(t *testing.T) {
	var out pingaccess.SiteView
	var targets = []string{"\"localhost:1234\""}
	var updatedTargets = []string{"\"localhost:1235\""}

	resource.ParallelTest(t, resource.TestCase{
		// PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPingAccessSiteDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingAccessSiteConfig("bar", &targets),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessSiteExists("pingaccess_site.acc_test", 3, &out),
					// testAccCheckPingAccessSiteAttributes(),
					// testAccCheckAWSPolicyAttachmentAttributes([]string{userName}, []string{roleName}, []string{groupName}, &out),
				),
			},
			{
				Config: testAccPingAccessSiteConfig("bar", &updatedTargets),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessSiteExists("pingaccess_site.acc_test", 6, &out),
					// testAccCheckAWSPolicyAttachmentAttributes([]string{userName2, userName3},
					// 	[]string{roleName2, roleName3}, []string{groupName2, groupName3}, &out),
				),
			},
		},
	})
}

func testAccCheckPingAccessSiteDestroy(s *terraform.State) error {
	return nil
}

func testAccPingAccessSiteConfig(name string, targets *[]string) string {
	return fmt.Sprintf(`
	resource "pingaccess_site" "acc_test" {
		name                         = "%s"
		targets                      = [%s]
		max_connections              = -1
		max_web_socket_connections   = -1
		availability_profile_id      = 1
	}`, name, strings.Join(*targets, ","))
}

func testAccCheckPingAccessSiteExists(n string, c int64, out *pingaccess.SiteView) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" || rs.Primary.ID == "0" {
			return fmt.Errorf("No site ID is set")
		}

		conn := testAccProvider.Meta().(*pingaccess.Client).Sites
		result, _, err := conn.GetSiteCommand(&pingaccess.GetSiteCommandInput{
			Path: struct {
				Id string
			}{
				Id: rs.Primary.ID,
			}})

		if err != nil {
			return fmt.Errorf("Error: Site (%s) not found", n)
		}

		if result.Name != rs.Primary.Attributes["name"] {
			return fmt.Errorf("Error: Site response (%s) didnt match state (%s)", result.Name, rs.Primary.Attributes["name"])
		}

		return nil
	}
}
