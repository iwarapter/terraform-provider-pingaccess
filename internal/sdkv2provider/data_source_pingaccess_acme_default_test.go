package sdkv2provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccPingAccessAcmeDefaultDataSource(t *testing.T) {
	resourceName := "data.pingaccess_acme_default.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV5ProviderFactories: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: `data "pingaccess_acme_default" "test" {}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "location"),
				),
			},
		},
	})
}
