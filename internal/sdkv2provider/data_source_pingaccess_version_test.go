package sdkv2provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccPingAccessVersionDataSource(t *testing.T) {
	resourceName := "data.pingaccess_version.test"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV5ProviderFactories: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccPingAccessVersionConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "version"),
				),
			},
		},
	})
}

func testAccPingAccessVersionConfig() string {
	return `data "pingaccess_version" "test" {}`
}
