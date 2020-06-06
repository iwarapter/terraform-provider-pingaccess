package pingaccess

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccPingAccessTrustedCertificateGroupsDataSource(t *testing.T) {
	resourceName := "data.pingaccess_trusted_certificate_group.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPingAccessTrustedCertificateGroupsDataSourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingAccessTrustedCertificateGroupsDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "ignore_all_certificate_errors", "true"),
					resource.TestCheckResourceAttr(resourceName, "skip_certificate_date_check", "false"),
					resource.TestCheckResourceAttr(resourceName, "system_group", "true"),
					resource.TestCheckResourceAttr(resourceName, "use_java_trust_store", "false"),
				),
			},
		},
	})
}

func TestAccPingAccessTrustedCertificateGroupsDataSource_NotFound(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPingAccessTrustedCertificateGroupsDataSourceDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccPingAccessTrustedCertificateGroupsDataSourceConfigNonExistent(),
				ExpectError: regexp.MustCompile(`Unable to find trusted certificate group with the name junk: `),
			},
		},
	})
}

func testAccCheckPingAccessTrustedCertificateGroupsDataSourceDestroy(s *terraform.State) error {
	return nil
}

func testAccPingAccessTrustedCertificateGroupsDataSourceConfig() string {
	return fmt.Sprintf(`
	data "pingaccess_trusted_certificate_group" "test" {
		name = "Trust Any"
	}`)
}

func testAccPingAccessTrustedCertificateGroupsDataSourceConfigNonExistent() string {
	return `
	data "pingaccess_trusted_certificate_group" "test" {
		name = "junk"
	}`
}
