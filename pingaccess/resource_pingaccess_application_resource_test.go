package pingaccess

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/iwarapter/pingaccess-sdk-go/pingaccess"
)

func init() {
	resource.AddTestSweepers("pingaccess_application_resource", &resource.Sweeper{
		Name: "pingaccess_application_resource",
		F:    testSweepApplicationResources,
	})
}

func testSweepApplicationResources(r string) error {
	url, _ := url.Parse("https://localhost:9000")
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	conn := pingaccess.NewClient("Administrator", "2Access2", url, "/pa-admin-api/v3", nil).Applications
	result, _, _ := conn.GetResourcesCommand(&pingaccess.GetResourcesCommandInput{Filter: "acc_test_"})
	for _, v := range result.Items {
		conn.DeleteApplicationResourceCommand(&pingaccess.DeleteApplicationResourceCommandInput{ApplicationId: strconv.Itoa(*v.ApplicationId), ResourceId: v.Id.String()})
	}
	return nil
}

func TestAccPingAccessApplicationResource(t *testing.T) {
	var out pingaccess.ApplicationView

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPingAccessApplicationResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingAccessApplicationResourceConfig("acc_test_bart", "/bar"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessApplicationResourceExists("pingaccess_application_resource.app_res_test_resource", 3, &out),
					testAccCheckPingAccessApplicationResourceAttributes("pingaccess_application_resource.app_res_test_resource", "bart", "/bar"),
				),
			},
			{
				Config: testAccPingAccessApplicationResourceConfig("acc_test_bart", "/bart"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessApplicationResourceExists("pingaccess_application_resource.app_res_test_resource", 6, &out),
					testAccCheckPingAccessApplicationResourceAttributes("pingaccess_application_resource.app_res_test_resource", "bart", "/bart"),
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
		name                         = "acc_test_app_res_test_site"
		targets                      = ["localhost:4321"]
		max_connections              = -1
		max_web_socket_connections   = -1
		availability_profile_id      = 1
	}

	resource "pingaccess_virtualhost" "app_res_test_virtualhost" {
		host                         = "acc-test-localhost"
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
		context_root				= "/bar"
		default_auth_type		= "Web"
		destination					= "Site"
		site_id							= "${pingaccess_site.app_res_test_site.id}"
		virtual_host_ids		= ["${pingaccess_virtualhost.app_res_test_virtualhost.id}"]
	}

resource "pingaccess_application_resource" "app_res_test_resource" {
  name = "acc_test_woot"
  methods = [
    "*"
  ]

  path_patterns {
    pattern = "/as/token.oauth2"
    type    = "WILDCARD"
	}
	
	path_patterns {
    pattern = "%s"
    type    = "WILDCARD"
  }

  path_prefixes = [
		"/as/token.oauth2",
		"%s"
  ]
  audit_level = "OFF"
  anonymous = false
  enabled = true
  root_resource = false
  application_id = "${pingaccess_application.app_res_test.id}"
}
	`, name, context, context)
}

func testAccCheckPingAccessApplicationResourceExists(n string, c int64, out *pingaccess.ApplicationView) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		b, _ := json.Marshal(rs)
		log.Printf("[INFO] RS: %s", b)
		if rs.Primary.ID == "" || rs.Primary.ID == "0" {
			return fmt.Errorf("No application resource ID is set")
		}

		conn := testAccProvider.Meta().(*pingaccess.Client).Applications
		result, _, err := conn.GetApplicationResourceCommand(&pingaccess.GetApplicationResourceCommandInput{
			ApplicationId: rs.Primary.Attributes["application_id"],
			ResourceId:    rs.Primary.ID,
		})

		if err != nil {
			return fmt.Errorf("Error: Application (%s) not found", n)
		}

		if *result.Name != rs.Primary.Attributes["name"] {
			return fmt.Errorf("Error: Application Resource response (%s) didnt match state (%s)", *result.Name, rs.Primary.Attributes["name"])
		}

		return nil
	}
}

func testAccCheckPingAccessApplicationResourceAttributes(n, name, context string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		conn := testAccProvider.Meta().(*pingaccess.Client).Applications
		result, _, err := conn.GetApplicationResourceCommand(&pingaccess.GetApplicationResourceCommandInput{
			ApplicationId: rs.Primary.Attributes["application_id"],
			ResourceId:    rs.Primary.ID,
		})

		if err != nil {
			return fmt.Errorf("Error: Application (%s) not found", n)
		}

		if *result.Name != rs.Primary.Attributes["name"] {
			return fmt.Errorf("Error: Application Resource response (%s) didnt match state (%s)", *result.Name, rs.Primary.Attributes["name"])
		}

		if strconv.Itoa(len(*result.Methods)) != rs.Primary.Attributes["methods.#"] {
			return fmt.Errorf("Error: Application Resource response (%s) didnt match state (%s)", strconv.Itoa(len(*result.Methods)), rs.Primary.Attributes["methods.#"])
		}

		return nil
	}
}
