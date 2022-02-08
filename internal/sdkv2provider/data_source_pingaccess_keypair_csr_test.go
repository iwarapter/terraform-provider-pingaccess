package sdkv2provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccPingAccessKeyPairCsrDataSource(t *testing.T) {
	resourceName := "data.pingaccess_keypair_csr.test"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV5ProviderFactories: testAccProviders,
		CheckDestroy:             testAccCheckPingAccessKeyPairCsrDataSourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingAccessKeyPairCsrDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "cert_request_pem"),
				),
			},
		},
	})
}

func TestAccPingAccessKeyPairCsrDataSource_NotFound(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV5ProviderFactories: testAccProviders,
		CheckDestroy:             testAccCheckPingAccessKeyPairCsrDataSourceDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccPingAccessKeyPairCsrDataSourceConfigNonExistent(),
				ExpectError: regexp.MustCompile(`unable to find KeyPairCSR with id 'junk'`),
			},
		},
	})
}

func testAccCheckPingAccessKeyPairCsrDataSourceDestroy(s *terraform.State) error {
	return nil
}

func testAccPingAccessKeyPairCsrDataSourceConfig() string {
	return `
	data "pingaccess_keypair_csr" "test" {
		id = "1"
	}`
}

func testAccPingAccessKeyPairCsrDataSourceConfigNonExistent() string {
	return `
	data "pingaccess_keypair_csr" "test" {
		id = "junk"
	}`
}
