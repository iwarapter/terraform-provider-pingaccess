package pingaccess

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pa "github.com/iwarapter/pingaccess-sdk-go/pingaccess"
)

func TestAccPingAccessPingFederateRuntime(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPingAccessPingFederateRuntimeDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingAccessPingFederateRuntimeConfig(fmt.Sprintf("https://%s:9031", os.Getenv("PINGFEDERATE_TEST_IP"))),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessPingFederateRuntimeExists("pingaccess_pingfederate_runtime.demo_pfr"),
				),
			},
			{
				Config: testAccPingAccessPingFederateRuntimeConfig(fmt.Sprintf("https://%s:9031", os.Getenv("PINGFEDERATE_TEST_IP"))),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessPingFederateRuntimeExists("pingaccess_pingfederate_runtime.demo_pfr"),
				),
			},
		},
	})
}

func testAccCheckPingAccessPingFederateRuntimeDestroy(s *terraform.State) error {
	return nil
}

func testAccPingAccessPingFederateRuntimeConfig(issuer string) string {
	return fmt.Sprintf(`
	resource "pingaccess_pingfederate_runtime" "demo_pfr" {
		description = "foo"
		issuer = "%s"
		skip_hostname_verification = true
		sts_token_exchange_endpoint = "https://foo/bar"
		use_slo = false
		trusted_certificate_group_id = 2
		use_proxy = true
	}`, issuer)
}

func testAccCheckPingAccessPingFederateRuntimeExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" || rs.Primary.ID == "0" {
			return fmt.Errorf("No third party service ID is set")
		}

		conn := testAccProvider.Meta().(*pa.Client).PingFederate
		result, _, err := conn.GetPingFederateRuntimeCommand()

		if err != nil {
			return fmt.Errorf("Error: PingFederateRuntime (%s) not found", n)
		}

		if *result.Issuer != rs.Primary.Attributes["issuer"] {
			return fmt.Errorf("Error: PingFederateRuntime response (%s) didnt match state (%s)", *result.Issuer, rs.Primary.Attributes["issuer"])
		}

		return nil
	}
}

func Test_resourcePingAccessPingFederateRuntimeReadData(t *testing.T) {
	cases := []struct {
		PingFederateRuntime pa.PingFederateMetadataRuntimeView
	}{
		{
			PingFederateRuntime: pa.PingFederateMetadataRuntimeView{
				Issuer:                   String("localhost"),
				SkipHostnameVerification: Bool(true),
				UseProxy:                 Bool(false),
				UseSlo:                   Bool(false),
			},
		},
		{
			PingFederateRuntime: pa.PingFederateMetadataRuntimeView{
				Issuer:                    String("localhost"),
				Description:               String("foo"),
				SkipHostnameVerification:  Bool(true),
				StsTokenExchangeEndpoint:  String("https://foo/bar"),
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
