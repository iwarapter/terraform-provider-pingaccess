package sdkv2provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccPingAccessAcmeDefaultDataSource(t *testing.T) {
	if !(paClient{apiVersion: paVersion}).Is60OrAbove() {
		t.Skipf("This test only runs against PingAccess 6.0 and above, not: %s", paVersion)
	}
	resourceName := "data.pingaccess_acme_default.test"
	resource.Test(t, resource.TestCase{
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
