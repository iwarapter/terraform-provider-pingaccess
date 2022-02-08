package sdkv2provider

import (
	"fmt"
	"testing"

	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess/models"
	"github.com/iwarapter/pingaccess-sdk-go/v62/services/httpsListeners"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccPingAccessHTTPSListener(t *testing.T) {
	resourceName := "pingaccess_https_listener.acc_test"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV5ProviderFactories: testAccProviders,
		CheckDestroy:             testAccCheckPingAccessHTTPSListenerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingAccessHTTPSListenerConfig(true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessHTTPSListenerExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", "ADMIN"),
					resource.TestCheckResourceAttr(resourceName, "key_pair_id", "1"),
					resource.TestCheckResourceAttr(resourceName, "use_server_cipher_suite_order", "true"),
				),
			},
			{
				Config: testAccPingAccessHTTPSListenerConfig(false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessHTTPSListenerExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", "ADMIN"),
					resource.TestCheckResourceAttr(resourceName, "key_pair_id", "1"),
					resource.TestCheckResourceAttr(resourceName, "use_server_cipher_suite_order", "false"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckPingAccessHTTPSListenerDestroy(s *terraform.State) error {
	return nil
}

func testAccPingAccessHTTPSListenerConfig(cipher bool) string {
	return fmt.Sprintf(`
resource "pingaccess_https_listener" "acc_test" {
	name   						  = "ADMIN"
	key_pair_id 				  = 1
	use_server_cipher_suite_order = %t
}`, cipher)
}

func testAccCheckPingAccessHTTPSListenerExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" || rs.Primary.ID == "0" {
			return fmt.Errorf("No listener ID is set")
		}

		conn := testAccProvider.Meta().(paClient).HttpsListeners
		result, _, err := conn.GetHttpsListenerCommand(&httpsListeners.GetHttpsListenerCommandInput{
			Id: rs.Primary.ID,
		})

		if err != nil {
			return fmt.Errorf("Error: listener (%s) not found", n)
		}

		if *result.Name != rs.Primary.Attributes["name"] {
			return fmt.Errorf("Error: listener response (%s) didnt match state (%s)", *result.Name, rs.Primary.Attributes["name"])
		}

		return nil
	}
}

func Test_resourcePingAccessHTTPSListenerReadData(t *testing.T) {
	cases := []struct {
		listener models.HttpsListenerView
	}{
		{
			listener: models.HttpsListenerView{
				Name:                      String("ADMIN"),
				KeyPairId:                 Int(1),
				UseServerCipherSuiteOrder: Bool(true),
			},
		},
	}
	for i, tc := range cases {
		t.Run(fmt.Sprintf("tc:%v", i), func(t *testing.T) {

			resourceSchema := resourcePingAccessHTTPSListenerSchema()
			resourceLocalData := schema.TestResourceDataRaw(t, resourceSchema, map[string]interface{}{})
			resourcePingAccessHTTPSListenerReadResult(resourceLocalData, &tc.listener)

			if got := *resourcePingAccessHTTPSListenerReadData(resourceLocalData); !cmp.Equal(got, tc.listener) {
				t.Errorf("resourcePingAccessHTTPSListenerReadData() = %v", cmp.Diff(got, tc.listener))
			}
		})
	}
}
