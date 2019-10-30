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

func TestAccPingAccessEngineListener(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPingAccessEngineListenerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingAccessEngineListenerConfig("cheese", true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessEngineListenerExists("pingaccess_engine_listener.acc_test"),
				),
			},
			{
				Config: testAccPingAccessEngineListenerConfig("cheese", false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessEngineListenerExists("pingaccess_engine_listener.acc_test"),
				),
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

		conn := testAccProvider.Meta().(*pingaccess.Client).EngineListeners
		result, _, err := conn.GetEngineListenerCommand(&pingaccess.GetEngineListenerCommandInput{
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
		EngineListener pa.EngineListenerView
	}{
		{
			EngineListener: pa.EngineListenerView{
				Name: String("engine1"),
				Port: Int(9999),
			},
		},
		{
			EngineListener: pa.EngineListenerView{
				Name:   String("engine2"),
				Port:   Int(9999),
				Secure: Bool(true),
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
