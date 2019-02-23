package pingaccess

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/iwarapter/pingaccess-sdk-go/pingaccess"
)

func init() {
	resource.AddTestSweepers("pingaccess_application", &resource.Sweeper{
		Name: "pingaccess_application",
		F:    testSweepApplication,
	})
}

func testSweepApplication(r string) error {
	url, _ := url.Parse("https://localhost:9000")
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	conn := pingaccess.NewClient("Administrator", "2Access2", url, "/pa-admin-api/v3", nil).Applications
	result, _, _ := conn.GetApplicationsCommand(&pingaccess.GetApplicationsCommandInput{Filter: "acc_test_"})
	for _, v := range result.Items {
		conn.DeleteApplicationCommand(&pingaccess.DeleteApplicationCommandInput{Id: v.Id.String()})
	}
	return nil
}

func TestAccPingAccessApplication(t *testing.T) {
	var out pingaccess.ApplicationView

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPingAccessApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingAccessApplicationConfig("acc_test_bar", "/bar"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessApplicationExists("pingaccess_application.acc_test", 3, &out),
				),
			},
			{
				Config: testAccPingAccessApplicationConfig("acc_test_bar", "/bart"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessApplicationExists("pingaccess_application.acc_test", 6, &out),
				),
			},
		},
	})
}

func testAccCheckPingAccessApplicationDestroy(s *terraform.State) error {
	return nil
}

func testAccPingAccessApplicationConfig(name string, context string) string {
	return fmt.Sprintf(`
	resource "pingaccess_site" "acc_test_site" {
		name                         = "acc_test_site"
		targets                      = ["localhost:4321"]
		max_connections              = -1
		max_web_socket_connections   = -1
		availability_profile_id      = 1
	}

	resource "pingaccess_virtualhost" "acc_test_virtualhost" {
		host                         = "acc-test-localhost"
		port                         = 4001
		agent_resource_cache_ttl     = 900
		key_pair_id                  = 0
		trusted_certificate_group_id = 0
 }

	resource "pingaccess_application" "acc_test" {
		access_validator_id = 0
		application_type 		= "Web"
		agent_id						= 0
		name								= "%s"
		context_root				= "%s"
		default_auth_type		= "Web"
		destination					= "Site"
		site_id							= "${pingaccess_site.acc_test_site.id}"
		virtual_host_ids		= ["${pingaccess_virtualhost.acc_test_virtualhost.id}"]
	}`, name, context)
}

func testAccCheckPingAccessApplicationExists(n string, c int64, out *pingaccess.ApplicationView) resource.TestCheckFunc {
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

		conn := testAccProvider.Meta().(*pingaccess.Client).Applications
		result, _, err := conn.GetApplicationCommand(&pingaccess.GetApplicationCommandInput{
			Id: rs.Primary.ID,
		})

		if err != nil {
			return fmt.Errorf("Error: Application (%s) not found", n)
		}

		if *result.Name != rs.Primary.Attributes["name"] {
			return fmt.Errorf("Error: Application response (%s) didnt match state (%s)", *result.Name, rs.Primary.Attributes["name"])
		}

		return nil
	}
}

func Test_flattenIdentityMappingIds(t *testing.T) {
	type args struct {
		in map[string]*int
	}
	tests := []struct {
		name string
		args args
		want []interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := flattenIdentityMappingIds(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("flattenIdentityMappingIds() = %v, want %v", got, tt.want)
			}
		})
	}
}
