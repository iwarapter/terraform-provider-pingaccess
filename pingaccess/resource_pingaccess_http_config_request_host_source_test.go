package pingaccess

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/iwarapter/pingaccess-sdk-go/pingaccess"
	pa "github.com/iwarapter/pingaccess-sdk-go/pingaccess"
)

func TestAccPingAccessHTTPConfigRequestHostSource(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPingAccessHTTPConfigRequestHostSourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingAccessHTTPConfigRequestHostSourceConfig("X-Forwarded-Host", "FIRST"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessHTTPConfigRequestHostSourceExists("pingaccess_http_config_request_host_source.http_config_request_host_source"),
				),
			},
			{
				Config: testAccPingAccessHTTPConfigRequestHostSourceConfig("MagicHostHeader", "LAST"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessHTTPConfigRequestHostSourceExists("pingaccess_http_config_request_host_source.http_config_request_host_source"),
				),
			},
		},
	})
}

func testAccCheckPingAccessHTTPConfigRequestHostSourceDestroy(s *terraform.State) error {
	return nil
}

func testAccPingAccessHTTPConfigRequestHostSourceConfig(header, location string) string {
	return fmt.Sprintf(`
	resource "pingaccess_http_config_request_host_source" "http_config_request_host_source" {
		header_name_list = [
			"Host",
			"%s"
		]
		list_value_location = "%s"
	}`, header, location)
}

func testAccCheckPingAccessHTTPConfigRequestHostSourceExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" || rs.Primary.ID == "0" {
			return fmt.Errorf("No http config request host source response ID is set")
		}

		conn := testAccProvider.Meta().(*pingaccess.Client).HttpConfig
		result, _, err := conn.GetHostSourceCommand()

		if err != nil {
			return fmt.Errorf("Error: http config request host source response (%s) not found", n)
		}

		if *result.ListValueLocation != rs.Primary.Attributes["list_value_location"] {
			return fmt.Errorf("Error: http config request host source response (%s) didnt match state (%s)", *result.ListValueLocation, rs.Primary.Attributes["list_value_location"])
		}

		return nil
	}
}

func Test_resourcePingAccessHTTPConfigRequestHostSourceReadData(t *testing.T) {
	cases := []struct {
		Resource pa.HostMultiValueSourceView
	}{
		{
			Resource: pa.HostMultiValueSourceView{
				HeaderNameList:    &[]*string{String("woot")},
				ListValueLocation: String("FIRST"),
			},
		},
	}
	for i, tc := range cases {
		t.Run(fmt.Sprintf("tc:%v", i), func(t *testing.T) {

			resourceSchema := resourcePingAccessHTTPConfigRequestHostSourceResourceSchema()
			resourceLocalData := schema.TestResourceDataRaw(t, resourceSchema, map[string]interface{}{})
			resourcePingAccessHTTPConfigRequestHostSourceReadResult(resourceLocalData, &tc.Resource)

			if got := *resourcePingAccessHTTPConfigRequestHostSourceReadData(resourceLocalData); !cmp.Equal(got, tc.Resource) {
				t.Errorf("resourcePingAccessHTTPConfigRequestHostSourceReadData() = %v", cmp.Diff(got, tc.Resource))
			}
		})
	}
}
