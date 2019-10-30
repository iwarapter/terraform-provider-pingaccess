package pingaccess

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/iwarapter/pingaccess-sdk-go/pingaccess"
	pa "github.com/iwarapter/pingaccess-sdk-go/pingaccess"
)

func TestAccPingAccessHTTPSListener(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPingAccessHTTPSListenerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingAccessHTTPSListenerConfig(true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessHTTPSListenerExists("pingaccess_https_listener.acc_test"),
				),
			},
			{
				Config: testAccPingAccessHTTPSListenerConfig(false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessHTTPSListenerExists("pingaccess_https_listener.acc_test"),
				),
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
	   name   									     = "ADMIN"
		 key_pair_id 									 = 1
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
			return fmt.Errorf("No HttpsListener ID is set")
		}

		conn := testAccProvider.Meta().(*pingaccess.Client).HttpsListeners
		result, _, err := conn.GetHttpsListenerCommand(&pingaccess.GetHttpsListenerCommandInput{
			Id: rs.Primary.ID,
		})

		if err != nil {
			return fmt.Errorf("Error: HttpsListener (%s) not found", n)
		}

		if *result.Name != rs.Primary.Attributes["name"] {
			return fmt.Errorf("Error: HttpsListener response (%s) didnt match state (%s)", *result.Name, rs.Primary.Attributes["name"])
		}

		return nil
	}
}

func Test_resourcePingAccessHTTPSListenerReadData(t *testing.T) {
	cases := []struct {
		HttpsListener pa.HttpsListenerView
	}{
		{
			HttpsListener: pa.HttpsListenerView{
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
			resourcePingAccessHTTPSListenerReadResult(resourceLocalData, &tc.HttpsListener)

			if got := *resourcePingAccessHTTPSListenerReadData(resourceLocalData); !cmp.Equal(got, tc.HttpsListener) {
				t.Errorf("resourcePingAccessHTTPSListenerReadData() = %v", cmp.Diff(got, tc.HttpsListener))
			}
		})
	}
}
