package pingaccess

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	pa "github.com/iwarapter/pingaccess-sdk-go/pingaccess"
)

func TestAccPingAccessAcmeServer(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPingAccessAcmeServerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingAccessAcmeServerConfig("https://host.docker.internal:14000/dir"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessAcmeServerExists("pingaccess_acme_server.acc_test"),
				),
			},
			{
				Config: testAccPingAccessAcmeServerConfig("https://host.docker.internal:14000/dir2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessAcmeServerExists("pingaccess_acme_server.acc_test"),
				),
				//ExpectNonEmptyPlan: true,
			},
		},
	})
}

func testAccCheckPingAccessAcmeServerDestroy(s *terraform.State) error {
	return nil
}

func testAccPingAccessAcmeServerConfig(url string) string {
	return fmt.Sprintf(`
	resource "pingaccess_acme_server" "acc_test" {
	   	name 				= "foo"
	   	url 				= "%s"
	}`, url)
}

func testAccCheckPingAccessAcmeServerExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" || rs.Primary.ID == "0" {
			return fmt.Errorf("No AcmeServer ID is set")
		}

		conn := testAccProvider.Meta().(*pa.Client).Acme
		result, _, err := conn.GetAcmeServerCommand(&pa.GetAcmeServerCommandInput{
			AcmeServerId: rs.Primary.ID,
		})

		if err != nil {
			return fmt.Errorf("Error: AcmeServer (%s) not found", n)
		}

		if *result.Name != rs.Primary.Attributes["name"] {
			return fmt.Errorf("Error: AcmeServer response (%s) didnt match state (%s)", *result.Name, rs.Primary.Attributes["name"])
		}
		return nil
	}
}

func Test_resourcePingAccessAcmeServerReadData(t *testing.T) {
	cases := []struct {
		AcmeServer pa.AcmeServerView
	}{
		{
			AcmeServer: pa.AcmeServerView{
				Name: String("engine1"),
				Url:  String("foo"),
			},
		},
	}
	for i, tc := range cases {
		t.Run(fmt.Sprintf("tc:%v", i), func(t *testing.T) {

			resourceSchema := resourcePingAccessAcmeServerSchema()
			resourceLocalData := schema.TestResourceDataRaw(t, resourceSchema, map[string]interface{}{})
			resourcePingAccessAcmeServerReadResult(resourceLocalData, &tc.AcmeServer)

			if got := *resourcePingAccessAcmeServerReadData(resourceLocalData); !cmp.Equal(got, tc.AcmeServer) {
				t.Errorf("resourcePingAccessAcmeServerReadData() = %v", cmp.Diff(got, tc.AcmeServer))
			}
		})
	}
}
