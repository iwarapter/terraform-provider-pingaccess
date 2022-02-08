package sdkv2provider

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccPingAccessCertificateDataSource(t *testing.T) {
	resourceName := "data.pingaccess_certificate.test"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV5ProviderFactories: testAccProviders,
		CheckDestroy:             testAccCheckPingAccessCertificateDataSourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingAccessCertificateDataSourceConfig("acctest_foobar"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "md5sum", "a0d4ef0bf7b5d849952aecf5c4fc8187"),
					resource.TestCheckResourceAttr(resourceName, "expires", "2221603200000"),
					resource.TestCheckResourceAttr(resourceName, "issuer_dn", "CN=Amazon Root CA 3, O=Amazon, C=US"),
					resource.TestCheckResourceAttr(resourceName, "serial_number", "06:6C:9F:D5:74:97:36:66:3F:3B:0B:9A:D9:E8:9E:76:03:F2:4A"),
					resource.TestCheckResourceAttr(resourceName, "sha1sum", "0d44dd8c3c8c1a1a58756481e90f2e2affb3d26e"),
					resource.TestCheckResourceAttr(resourceName, "signature_algorithm", "SHA256withECDSA"),
					resource.TestCheckResourceAttr(resourceName, "status", "Valid"),
					resource.TestCheckResourceAttr(resourceName, "subject_dn", "CN=Amazon Root CA 3, O=Amazon, C=US"),
					resource.TestCheckResourceAttr(resourceName, "valid_from", "1432598400000"),
				),
			},
		},
	})
}

func TestAccPingAccessCertificateDataSource_NotFound(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV5ProviderFactories: testAccProviders,
		CheckDestroy:             testAccCheckPingAccessCertificateDataSourceDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccPingAccessCertificateDataSourceConfigNonExistent(),
				ExpectError: regexp.MustCompile(` unable to find Certificate with alias 'junk' found '0' results`),
			},
		},
	})
}

func testAccCheckPingAccessCertificateDataSourceDestroy(s *terraform.State) error {
	return nil
}

func testAccPingAccessCertificateDataSourceConfig(alias string) string {
	return fmt.Sprintf(`
	resource "pingaccess_certificate" "example" {
		alias = "%s"
		file_data = base64encode(file("test_cases/amazon_root_ca3.pem"))
	}

	data "pingaccess_certificate" "test" {
		alias = pingaccess_certificate.example.alias

		depends_on = [
			pingaccess_certificate.example
		]
	}`, alias)
}

func testAccPingAccessCertificateDataSourceConfigNonExistent() string {
	return `
	data "pingaccess_certificate" "test" {
		alias = "junk"
	}`
}
