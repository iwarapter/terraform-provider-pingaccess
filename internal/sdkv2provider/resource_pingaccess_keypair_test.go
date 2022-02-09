package sdkv2provider

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/iwarapter/pingaccess-sdk-go/v62/services/keyPairs"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func init() {
	resource.AddTestSweepers("keypairs", &resource.Sweeper{
		Name: "keypairs",
		F: func(r string) error {
			svc := keyPairs.New(conf)
			results, _, err := svc.GetKeyPairsCommand(&keyPairs.GetKeyPairsCommandInput{})
			if err != nil {
				return fmt.Errorf("unable to list keypairs to sweep %s", err)
			}
			for _, item := range results.Items {
				if !strings.HasPrefix(*item.Alias, "Generated:") {
					_, err = svc.DeleteKeyPairCommand(&keyPairs.DeleteKeyPairCommandInput{Id: strconv.Itoa(*item.Id)})
					if err != nil {
						return fmt.Errorf("unable to sweep keypair %s because %s", strconv.Itoa(*item.Id), err)
					}
				}
			}
			return nil
		},
	})
}

func TestAccPingAccessKeyPair(t *testing.T) {
	resourceName := "pingaccess_keypair.test"
	resourceNameGen := "pingaccess_keypair.test_generate"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV5ProviderFactories: testAccProviders,
		CheckDestroy:             testAccCheckPingAccessKeyPairDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingAccessKeyPairConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessKeyPairExists(resourceName),
					testAccCheckPingAccessKeyPairAttributes(resourceName),
					resource.TestCheckResourceAttr(resourceName, "md5sum", "e7eca06b83d829d5bf514178b6271a4a"),
					resource.TestCheckResourceAttr(resourceName, "expires", "1723580460000"),
					resource.TestCheckResourceAttr(resourceName, "issuer_dn", "CN=(LOCAL) CA, OU=LOCAL, O=ORGANISATION, L=LOCALITY, ST=STATE, C=--"),
					resource.TestCheckResourceAttr(resourceName, "serial_number", "32:E7:2B:03:67:74:C1:99:DA:A9:88:A6:8A:36:95:0B:54:45:23:13"),
					resource.TestCheckResourceAttr(resourceName, "sha1sum", "596fc7aa20cea185da02afefbd239677d19be43b"),
					resource.TestCheckResourceAttr(resourceName, "signature_algorithm", "SHA256withRSA"),
					resource.TestCheckResourceAttr(resourceName, "status", "Valid"),
					resource.TestCheckResourceAttr(resourceName, "subject_dn", "CN=localhost.localdomain"),
					resource.TestCheckResourceAttr(resourceName, "subject_cn", "localhost.localdomain"),
					resource.TestCheckResourceAttr(resourceName, "valid_from", "1565900460000"),
				),
			},
			{
				Config: testAccPingAccessKeyPairConfigGenerate(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessKeyPairExists(resourceNameGen),
					testAccCheckPingAccessKeyPairAttributes(resourceNameGen),
					resource.TestCheckResourceAttrSet(resourceNameGen, "md5sum"),
					resource.TestCheckResourceAttrSet(resourceNameGen, "expires"),
					resource.TestCheckResourceAttr(resourceNameGen, "issuer_dn", "CN=Test, OU=Test, O=Test, L=Test, ST=Test, C=GB"),
					resource.TestCheckResourceAttrSet(resourceNameGen, "serial_number"),
					resource.TestCheckResourceAttrSet(resourceNameGen, "sha1sum"),
					resource.TestCheckResourceAttr(resourceNameGen, "signature_algorithm", "SHA256withRSA"),
					resource.TestCheckResourceAttr(resourceNameGen, "status", "Valid"),
					resource.TestCheckResourceAttr(resourceNameGen, "subject_dn", "CN=Test, OU=Test, O=Test, L=Test, ST=Test, C=GB"),
					resource.TestCheckResourceAttr(resourceNameGen, "subject_cn", "Test"),
					resource.TestCheckResourceAttrSet(resourceNameGen, "valid_from"),
				),
			},
			{
				ResourceName:      resourceNameGen,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckPingAccessKeyPairDestroy(s *terraform.State) error {
	return nil
}

func testAccPingAccessKeyPairConfig() string {
	return `
	resource "pingaccess_keypair" "test" {
		alias = "acctest_test"
		file_data = filebase64("test_cases/provider.p12")
		password = "password"
	}`
}

func testAccPingAccessKeyPairConfigGenerate() string {
	return `
	resource "pingaccess_keypair" "test_generate" {
		alias = "acctest_test2"
		city = "Test"
		common_name = "Test"
		country = "GB"
		key_algorithm = "RSA"
		key_size = 2048
		organization = "Test"
		organization_unit = "Test"
		state = "Test"
		valid_days = 365
	}`
}

func testAccCheckPingAccessKeyPairExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" || rs.Primary.ID == "0" {
			return fmt.Errorf("No KeyPair ID is set")
		}

		conn := testAccProvider.Meta().(paClient).KeyPairs
		result, _, err := conn.GetKeyPairCommand(&keyPairs.GetKeyPairCommandInput{
			Id: rs.Primary.ID,
		})

		if err != nil {
			return fmt.Errorf("Error: KeyPair (%s) not found", n)
		}

		if *result.Alias != rs.Primary.Attributes["alias"] {
			return fmt.Errorf("Error: KeyPair response (%s) didnt match state (%s)", *result.Alias, rs.Primary.Attributes["alias"])
		}

		return nil
	}
}

func testAccCheckPingAccessKeyPairAttributes(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs := s.RootModule().Resources[n]
		if rs.Primary.ID == "" || rs.Primary.ID == "0" {
			return fmt.Errorf("No KeyPair ID is set")
		}

		conn := testAccProvider.Meta().(paClient).KeyPairs
		result, _, err := conn.GetKeyPairCommand(&keyPairs.GetKeyPairCommandInput{
			Id: rs.Primary.ID,
		})

		if err != nil {
			return fmt.Errorf("Error: KeyPair (%s) not found", n)
		}

		if *result.Alias != rs.Primary.Attributes["alias"] {
			return fmt.Errorf("Error: KeyPair response (%s) didnt match state (%s)", *result.Alias, rs.Primary.Attributes["alias"])
		}

		return nil
	}
}
