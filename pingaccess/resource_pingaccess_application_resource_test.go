package pingaccess

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/iwarapter/pingaccess-sdk-go/service/applications"
)

func TestAccPingAccessApplicationResource(t *testing.T) {
	var out applications.ApplicationView

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPingAccessApplicationResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingAccessApplicationResourceConfig("bart", "/bar"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessApplicationResourceExists("pingaccess_application.app_res_test", 3, &out),
					// testAccCheckPingAccessApplicationResourceAttributes(),
					// testAccCheckAWSPolicyAttachmentAttributes([]string{userName}, []string{roleName}, []string{groupName}, &out),
				),
			},
			{
				Config: testAccPingAccessApplicationResourceConfig("bart", "/bart"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessApplicationResourceExists("pingaccess_application.app_res_test", 6, &out),
					// testAccCheckAWSPolicyAttachmentAttributes([]string{userName2, userName3},
					// 	[]string{roleName2, roleName3}, []string{groupName2, groupName3}, &out),
				),
			},
		},
	})
}

func testAccCheckPingAccessApplicationResourceDestroy(s *terraform.State) error {
	return nil
}

func testAccPingAccessApplicationResourceConfig(name string, context string) string {
	return fmt.Sprintf(`
	resource "pingaccess_site" "app_res_test_site" {
		name                         = "app_res_test_site"
		targets                      = ["localhost:4321"]
		max_connections              = -1
		max_web_socket_connections   = -1
		availability_profile_id      = 1
	}

	resource "pingaccess_virtualhost" "app_res_test_virtualhost" {
		host                         = "localhost"
		port                         = 4000
		agent_resource_cache_ttl     = 900
		key_pair_id                  = 0
		trusted_certificate_group_id = 0
 }

	resource "pingaccess_application" "app_res_test" {
		access_validator_id = 0
		application_type 		= "Web"
		agent_id						= 0
		name								= "%s"
		context_root				= "%s"
		default_auth_type		= "Web"
		destination					= "Site"
		site_id							= "${pingaccess_site.app_res_test_site.id}"
		virtual_host_ids		= ["${pingaccess_virtualhost.app_res_test_virtualhost.id}"]
	}
	

resource "pingaccess_application_resource" "app_res_test_resource" {
  name = "woot"
  methods = [
    "*"
  ]
  path_prefixes = [
    "/woot"
  ]
  default_auth_type_override = "Web"
  audit_level = "OFF"
  anonymous = false
  enabled = true
  root_resource = false
  application_id = "${pingaccess_application.app_res_test.id}"
}
	`, name, context)
}

func testAccCheckPingAccessApplicationResourceExists(n string, c int64, out *applications.ApplicationView) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		// b, _ := json.Marshal(rs)
		// fmt.Printf("%s", b)

		if rs.Primary.ID == "" || rs.Primary.ID == "0" {
			return fmt.Errorf("No application ID is set")
		}

		conn := testAccProvider.Meta().(*PAClient).appconn
		resp, err := conn.GetApplicationCommand(&applications.GetApplicationCommandInput{
			Path: struct {
				Id string
			}{
				Id: rs.Primary.ID,
			}})

		if err != nil {
			return fmt.Errorf("Error: Application (%s) not found", n)
		}

		if resp.Name != rs.Primary.Attributes["name"] {
			return fmt.Errorf("Error: Application response (%s) didnt match state (%s)", resp.Name, rs.Primary.Attributes["name"])
		}

		return nil
	}
}
