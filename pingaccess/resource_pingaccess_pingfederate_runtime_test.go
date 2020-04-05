package pingaccess

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	pa "github.com/iwarapter/pingaccess-sdk-go/pingaccess"
)

// func TestAccPingAccessPingFederateRuntime(t *testing.T) {
// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:     func() { testAccPreCheck(t) },
// 		Providers:    testAccProviders,
// 		CheckDestroy: testAccCheckPingAccessPingFederateRuntimeDestroy,
// 		Steps: []resource.TestStep{
// 			{
// 				Config: testAccPingAccessPingFederateRuntimeConfig("localhost", "9030"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckPingAccessPingFederateRuntimeExists("pingaccess_pingfederate_runtime.demo_pfr"),
// 				),
// 			},
// 			{
// 				Config: testAccPingAccessPingFederateRuntimeConfig("localhost", "9031"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckPingAccessPingFederateRuntimeExists("pingaccess_pingfederate_runtime.demo_pfr"),
// 				),
// 			},
// 		},
// 	})
// }

// func testAccCheckPingAccessPingFederateRuntimeDestroy(s *terraform.State) error {
// 	return nil
// }

// func testAccPingAccessPingFederateRuntimeConfig(host, port string) string {
// 	return fmt.Sprintf(`
// 	resource "pingaccess_pingfederate_runtime" "demo_pfr" {
// 		host = "%s"
// 		port = %s
// 		targets = []
//   	skip_hostname_verification = false
// 		expected_hostname = "hosty"
// 		back_channel_secure = true
// 		use_slo = true,
// 		base_path = "/woot"
// 		audit_level = "OFF"
// 		secure = true
// 		trusted_certificate_group_id = 1
// 		use_proxy = true
// 	}`, host, port)
// }

// func testAccCheckPingAccessPingFederateRuntimeExists(n string) resource.TestCheckFunc {
// 	return func(s *terraform.State) error {
// 		rs, ok := s.RootModule().Resources[n]
// 		if !ok {
// 			return fmt.Errorf("Not found: %s", n)
// 		}

// 		if rs.Primary.ID == "" || rs.Primary.ID == "0" {
// 			return fmt.Errorf("No third party service ID is set")
// 		}

// 		conn := testAccProvider.Meta().(*pa.Client).PingFederate
// 		result, _, err := conn.GetPingFederateCommand()

// 		if err != nil {
// 			return fmt.Errorf("Error: PingFederateRuntime (%s) not found", n)
// 		}

// 		if *result.Host != rs.Primary.Attributes["host"] {
// 			return fmt.Errorf("Error: PingFederateRuntime response (%s) didnt match state (%s)", *result.Host, rs.Primary.Attributes["host"])
// 		}

// 		return nil
// 	}
// }

func Test_resourcePingAccessPingFederateRuntimeReadData(t *testing.T) {
	cases := []struct {
		PingFederateRuntime pa.PingFederateRuntimeView
	}{
		{
			PingFederateRuntime: pa.PingFederateRuntimeView{
				Host:                     String("localhost"),
				Port:                     Int(9031),
				AuditLevel:               String("ON"),
				BackChannelSecure:        Bool(false),
				SkipHostnameVerification: Bool(true),
				Targets:                  &[]*string{},
				UseProxy:                 Bool(false),
				UseSlo:                   Bool(false),
				Secure:                   Bool(false),
			},
		},
		{
			PingFederateRuntime: pa.PingFederateRuntimeView{
				Host:                      String("localhost"),
				Port:                      Int(9031),
				AuditLevel:                String("none"),
				BackChannelBasePath:       String("/path"),
				BackChannelSecure:         Bool(true),
				BasePath:                  String("/path"),
				ExpectedHostname:          String("hosty"),
				Secure:                    Bool(true),
				SkipHostnameVerification:  Bool(true),
				Targets:                   &[]*string{String("localhost")},
				TrustedCertificateGroupId: Int(0),
				UseProxy:                  Bool(true),
				UseSlo:                    Bool(true),
			},
		},
	}
	for i, tc := range cases {
		t.Run(fmt.Sprintf("tc:%v", i), func(t *testing.T) {

			resourceSchema := resourcePingAccessPingFederateRuntimeSchema()
			resourceLocalData := schema.TestResourceDataRaw(t, resourceSchema, map[string]interface{}{})
			resourcePingAccessPingFederateRuntimeReadResult(resourceLocalData, &tc.PingFederateRuntime)

			if got := *resourcePingAccessPingFederateRuntimeReadData(resourceLocalData); !cmp.Equal(got, tc.PingFederateRuntime) {
				t.Errorf("resourcePingAccessPingFederateRuntimeReadData() = %v", cmp.Diff(got, tc.PingFederateRuntime))
			}
		})
	}
}
