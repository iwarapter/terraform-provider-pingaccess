package pingaccess

import (
	"fmt"
	"testing"

	"github.com/iwarapter/pingaccess-sdk-go/pingaccess/models"
	"github.com/iwarapter/pingaccess-sdk-go/services/virtualhosts"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccPingAccessVirtualHost(t *testing.T) {
	resourceName := "pingaccess_virtualhost.acc_test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPingAccessVirtualHostDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingAccessVirtualHostConfig("cheese", 3000),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessVirtualHostExists(resourceName),
				),
			},
			{
				Config: testAccPingAccessVirtualHostConfig("cheese", 3001),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessVirtualHostExists(resourceName),
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

func testAccCheckPingAccessVirtualHostExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" || rs.Primary.ID == "0" {
			return fmt.Errorf("No virtualhost ID is set")
		}

		conn := testAccProvider.Meta().(paClient).Virtualhosts
		result, _, err := conn.GetVirtualHostCommand(&virtualhosts.GetVirtualHostCommandInput{
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

func Test_resourcePingAccessVirtualhostReadData(t *testing.T) {
	cases := []struct {
		VirtualHost models.VirtualHostView
	}{
		{
			VirtualHost: models.VirtualHostView{
				Host:                      String("localhost"),
				Port:                      Int(9999),
				AgentResourceCacheTTL:     Int(0),
				KeyPairId:                 Int(0),
				TrustedCertificateGroupId: Int(0),
			},
		},
		{
			VirtualHost: models.VirtualHostView{
				Host:                      String("localhost"),
				Port:                      Int(9999),
				AgentResourceCacheTTL:     Int(30),
				KeyPairId:                 Int(30),
				TrustedCertificateGroupId: Int(30),
			},
		},
	}
	for i, tc := range cases {
		t.Run(fmt.Sprintf("tc:%v", i), func(t *testing.T) {

			resourceSchema := resourcePingAccessVirtualHostSchema()
			resourceLocalData := schema.TestResourceDataRaw(t, resourceSchema, map[string]interface{}{})
			resourcePingAccessVirtualHostReadResult(resourceLocalData, &tc.VirtualHost)

			if got := *resourcePingAccessVirtualHostReadData(resourceLocalData); !cmp.Equal(got, tc.VirtualHost) {
				t.Errorf("resourcePingAccessVirtualHostReadData() = %v", cmp.Diff(got, tc.VirtualHost))
			}
		})
	}
}
