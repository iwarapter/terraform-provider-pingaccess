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
	resource.AddTestSweepers("pingaccess_virtualhosts", &resource.Sweeper{
		Name: "pingaccess_virtualhosts",
		F:    testSweepVirtualhosts,
	})
}

func testSweepVirtualhosts(r string) error {
	url, _ := url.Parse("https://localhost:9000")
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	conn := pingaccess.NewClient("Administrator", "2Access2", url, "/pa-admin-api/v3", nil).Virtualhosts
	result, _, _ := conn.GetVirtualHostsCommand(&pingaccess.GetVirtualHostsCommandInput{})
	for _, v := range result.Items {
		conn.DeleteVirtualHostCommand(&pingaccess.DeleteVirtualHostCommandInput{Id: v.Id.String()})
	}
	return nil
}

func TestAccPingAccessVirtualHost(t *testing.T) {
	var out pingaccess.VirtualHostView

	resource.ParallelTest(t, resource.TestCase{
		// PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPingAccessVirtualHostDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingAccessVirtualHostConfig("cheese", 3000),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessVirtualHostExists("pingaccess_virtualhost.acc_test", 3, &out),
				),
			},
			{
				Config: testAccPingAccessVirtualHostConfig("cheese", 3001),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessVirtualHostExists("pingaccess_virtualhost.acc_test", 6, &out),
				),
			},
		},
	})
}

func testAccCheckPingAccessVirtualHostDestroy(s *terraform.State) error {
	return nil
}

func testAccPingAccessVirtualHostConfig(host string, port int) string {
	return fmt.Sprintf(`
	resource "pingaccess_virtualhost" "acc_test" {
	   host                         = "tf-acc-test-%s"
	   port                         = %d
	   agent_resource_cache_ttl     = 900
	   key_pair_id                  = 0
	   trusted_certificate_group_id = 0
	}`, host, port)
}

func testAccCheckPingAccessVirtualHostExists(n string, c int64, out *pingaccess.VirtualHostView) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" || rs.Primary.ID == "0" {
			return fmt.Errorf("No virtualhost ID is set")
		}

		conn := testAccProvider.Meta().(*pingaccess.Client).Virtualhosts
		result, _, err := conn.GetVirtualHostCommand(&pingaccess.GetVirtualHostCommandInput{
			Id: rs.Primary.ID,
		})

		if err != nil {
			return fmt.Errorf("Error: VirtualHost (%s) not found", n)
		}

		if *result.Host != rs.Primary.Attributes["host"] {
			return fmt.Errorf("Error: VirtualHost response (%s) didnt match state (%s)", *result.Host, rs.Primary.Attributes["host"])
		}

		return nil
	}
}
