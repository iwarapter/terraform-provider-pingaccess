package pingaccess

import (
	"fmt"
	"testing"

	"github.com/iwarapter/pingaccess-sdk-go/pingaccess/models"
	"github.com/iwarapter/pingaccess-sdk-go/services/engineListeners"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccPingAccessEngineListener(t *testing.T) {
	resourceName := "pingaccess_engine_listener.acc_test"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPingAccessEngineListenerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingAccessEngineListenerConfig("cheese", true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessEngineListenerExists(resourceName),
				),
			},
			{
				Config: testAccPingAccessEngineListenerConfig("cheese", false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessEngineListenerExists(resourceName),
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

func testAccCheckPingAccessEngineListenerDestroy(s *terraform.State) error {
	return nil
}

func testAccPingAccessEngineListenerConfig(host string, secure bool) string {
	return fmt.Sprintf(`
	resource "pingaccess_engine_listener" "acc_test" {
	   name   = "engine-%s"
	   port   = 443
	   secure	= %t
	}`, host, secure)
}

func testAccCheckPingAccessEngineListenerExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" || rs.Primary.ID == "0" {
			return fmt.Errorf("No EngineListener ID is set")
		}

		conn := testAccProvider.Meta().(paClient).EngineListeners
		result, _, err := conn.GetEngineListenerCommand(&engineListeners.GetEngineListenerCommandInput{
			Id: rs.Primary.ID,
		})

		if err != nil {
			return fmt.Errorf("Error: EngineListener (%s) not found", n)
		}

		if *result.Name != rs.Primary.Attributes["name"] {
			return fmt.Errorf("Error: EngineListener response (%s) didnt match state (%s)", *result.Name, rs.Primary.Attributes["name"])
		}

		return nil
	}
}

func Test_resourcePingAccessEngineListenerReadData(t *testing.T) {
	cases := []struct {
		EngineListener models.EngineListenerView
	}{
		{
			EngineListener: models.EngineListenerView{
				Name:                      String("engine1"),
				Port:                      Int(9999),
				Secure:                    Bool(true),
				TrustedCertificateGroupId: Int(0),
			},
		},
		{
			EngineListener: models.EngineListenerView{
				Name:                      String("engine2"),
				Port:                      Int(9999),
				Secure:                    Bool(false),
				TrustedCertificateGroupId: Int(2),
			},
		},
	}
	for i, tc := range cases {
		t.Run(fmt.Sprintf("tc:%v", i), func(t *testing.T) {

			resourceSchema := resourcePingAccessEngineListenerSchema()
			resourceLocalData := schema.TestResourceDataRaw(t, resourceSchema, map[string]interface{}{})
			resourcePingAccessEngineListenerReadResult(resourceLocalData, &tc.EngineListener)

			if got := *resourcePingAccessEngineListenerReadData(resourceLocalData); !cmp.Equal(got, tc.EngineListener) {
				t.Errorf("resourcePingAccessEngineListenerReadData() = %v", cmp.Diff(got, tc.EngineListener))
			}
		})
	}
}
