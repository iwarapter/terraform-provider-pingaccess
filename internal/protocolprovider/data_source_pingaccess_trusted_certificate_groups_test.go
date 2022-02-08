package protocol

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccPingAccessTrustedCertificateGroupsDataSource(t *testing.T) {
	resourceName := "data.pingaccess_trusted_certificate_group.test"

	resource.Test(t, resource.TestCase{
		ProtoV5ProviderFactories: map[string]func() (tfprotov5.ProviderServer, error){
			"pingaccess": func() (tfprotov5.ProviderServer, error) {
				return Server(), nil
			},
		},
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
	resource.Test(t, resource.TestCase{
		ProtoV5ProviderFactories: map[string]func() (tfprotov5.ProviderServer, error){
			"pingaccess": func() (tfprotov5.ProviderServer, error) {
				return Server(), nil
			},
		},
		CheckDestroy: testAccCheckPingAccessTrustedCertificateGroupsDataSourceDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccPingAccessTrustedCertificateGroupsDataSourceConfigNonExistent(),
				ExpectError: regexp.MustCompile(`unable to find TrustedCertificateGroup with the name 'junk' found '0' results`),
			},
		},
	})
}

func testAccCheckPingAccessTrustedCertificateGroupsDataSourceDestroy(s *terraform.State) error {
	return nil
}

func testAccPingAccessTrustedCertificateGroupsDataSourceConfig() string {
	return `
	data "pingaccess_trusted_certificate_group" "test" {
		name = "Trust Any"
	}`
}

func testAccPingAccessTrustedCertificateGroupsDataSourceConfigNonExistent() string {
	return `
	data "pingaccess_trusted_certificate_group" "test" {
		name = "junk"
	}`
}
