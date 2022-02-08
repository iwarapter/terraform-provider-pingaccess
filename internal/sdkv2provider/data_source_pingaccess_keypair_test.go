package sdkv2provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccPingAccessKeyPairDataSource(t *testing.T) {
	resourceName := "data.pingaccess_keypair.test"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV5ProviderFactories: testAccProviders,
		CheckDestroy:             testAccCheckPingAccessKeyPairDataSourceDestroy,
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
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV5ProviderFactories: testAccProviders,
		CheckDestroy:             testAccCheckPingAccessKeyPairDataSourceDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccPingAccessKeyPairDataSourceConfigNonExistent(),
				ExpectError: regexp.MustCompile(`unable to find KeyPair with alias 'junk' found '0' results`),
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
