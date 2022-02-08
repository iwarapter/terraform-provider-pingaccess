package sdkv2provider

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/iwarapter/pingaccess-sdk-go/v62/services/certificates"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func init() {
	resource.AddTestSweepers("certificate", &resource.Sweeper{
		Name:         "certificate",
		Dependencies: []string{"keypairs", "trusted_certificate_group"},
		F: func(r string) error {
			svc := certificates.New(conf)
			results, _, err := svc.GetTrustedCerts(&certificates.GetTrustedCertsInput{})
			if err != nil {
				return fmt.Errorf("unable to list certificates to sweep %s", err)
			}
			for _, item := range results.Items {
				_, err = svc.DeleteTrustedCertCommand(&certificates.DeleteTrustedCertCommandInput{Id: strconv.Itoa(*item.Id)})
				if err != nil {
					return fmt.Errorf("unable to sweep certificates %d because %s", *item.Id, err)
				}
			}
			return nil
		},
	})
}

func TestAccPingAccessCertificate(t *testing.T) {
	resourceName := "pingaccess_certificate.test"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV5ProviderFactories: testAccProviders,
		CheckDestroy:             testAccCheckPingAccessCertificateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingAccessCertificateConfig("bar"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessCertificateExists(resourceName),
					testAccCheckPingAccessCertificateAttributes(resourceName),
				),
			},
			{
				Config: testAccPingAccessCertificateConfig("foo"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessCertificateExists(resourceName),
					testAccCheckPingAccessCertificateAttributes(resourceName),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckPingAccessCertificateDestroy(s *terraform.State) error {
	return nil
}

func testAccPingAccessCertificateConfig(name string) string {
	return fmt.Sprintf(`
	resource "pingaccess_trusted_certificate_group" "test" {
		name = "acctest_tcg"
		cert_ids = [
			pingaccess_certificate.test.id
		]
	}

	resource "pingaccess_certificate" "test" {
		alias = "acctest_%s"
		file_data = base64encode(chomp(file("test_cases/amazon_root_ca1.pem")))
	}`, name)
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

		conn := testAccProvider.Meta().(paClient).Certificates
		result, _, err := conn.GetTrustedCert(&certificates.GetTrustedCertInput{
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
		rs := s.RootModule().Resources[n]
		if rs.Primary.ID == "" || rs.Primary.ID == "0" {
			return fmt.Errorf("No Certificate ID is set")
		}

		conn := testAccProvider.Meta().(paClient).Certificates
		result, _, err := conn.GetTrustedCert(&certificates.GetTrustedCertInput{
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
