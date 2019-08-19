package pingaccess

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccPingAccessKeyPairDataSource(t *testing.T) {
	resourceName := "data.pingaccess_keypair.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPingAccessKeyPairDataSourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingAccessKeyPairDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "issuer_dn", "CN=localhost, O=Ping Identity, C=US"),
					resource.TestCheckResourceAttr(resourceName, "signature_algorithm", "SHA256withRSA"),
					resource.TestCheckResourceAttr(resourceName, "status", "Valid"),
					resource.TestCheckResourceAttr(resourceName, "subject_dn", "CN=localhost, O=Ping Identity, C=US"),
				),
			},
		},
	})
}

func TestAccPingAccessKeyPairDataSource_NotFound(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPingAccessKeyPairDataSourceDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccPingAccessKeyPairDataSourceConfigNonExistent(),
				ExpectError: regexp.MustCompile(`Unable to find keypair with alias junk: `),
			},
		},
	})
}

func testAccCheckPingAccessKeyPairDataSourceDestroy(s *terraform.State) error {
	return nil
}

func testAccPingAccessKeyPairDataSourceConfig() string {
	return `
	data "pingaccess_keypair" "test" {
		alias = "Generated: ADMIN"
	}`
}

func testAccPingAccessKeyPairDataSourceConfigNonExistent() string {
	return `
	data "pingaccess_keypair" "test" {
		alias = "junk"
	}`
}
