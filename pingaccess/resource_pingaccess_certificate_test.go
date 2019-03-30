package pingaccess

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	pa "github.com/iwarapter/pingaccess-sdk-go/pingaccess"
)

func TestAccPingAccessCertificate(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPingAccessCertificateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingAccessCertificateConfig("bar", "1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessCertificateExists("pingaccess_certificate.acc_test_idm_bar"),
					testAccCheckPingAccessCertificateAttributes("pingaccess_certificate.acc_test_idm_bar"),
				),
			},
			{
				Config: testAccPingAccessCertificateConfig("bar", "2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessCertificateExists("pingaccess_certificate.acc_test_idm_bar"),
					testAccCheckPingAccessCertificateAttributes("pingaccess_certificate.acc_test_idm_bar"),
				),
			},
		},
	})
}

func testAccCheckPingAccessCertificateDestroy(s *terraform.State) error {
	return nil
}

func testAccPingAccessCertificateConfig(name, configUpdate string) string {
	return fmt.Sprintf(`
	resource "pingaccess_certificate" "acc_test_idm_%s" {
		alias = "%s"
		file_data = "${base64encode(file("test_cases/amazon_root_ca%s.pem"))}"
	}`, name, name, configUpdate)
}

func testAccCheckPingAccessCertificateExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" || rs.Primary.ID == "0" {
			return fmt.Errorf("No Certificate ID is set")
		}

		conn := testAccProvider.Meta().(*pa.Client).Certificates
		result, _, err := conn.GetTrustedCert(&pa.GetTrustedCertInput{
			Id: rs.Primary.ID,
		})

		if err != nil {
			return fmt.Errorf("Error: Certificate (%s) not found", n)
		}

		if *result.Alias != rs.Primary.Attributes["alias"] {
			return fmt.Errorf("Error: Certificate response (%s) didnt match state (%s)", *result.Alias, rs.Primary.Attributes["alias"])
		}

		return nil
	}
}

func testAccCheckPingAccessCertificateAttributes(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, _ := s.RootModule().Resources[n]
		if rs.Primary.ID == "" || rs.Primary.ID == "0" {
			return fmt.Errorf("No Certificate ID is set")
		}

		conn := testAccProvider.Meta().(*pa.Client).Certificates
		result, _, err := conn.GetTrustedCert(&pa.GetTrustedCertInput{
			Id: rs.Primary.ID,
		})

		if err != nil {
			return fmt.Errorf("Error: Certificate (%s) not found", n)
		}

		if *result.Alias != rs.Primary.Attributes["alias"] {
			return fmt.Errorf("Error: Certificate response (%s) didnt match state (%s)", *result.Alias, rs.Primary.Attributes["alias"])
		}

		return nil
	}
}
