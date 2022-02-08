package sdkv2provider

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/iwarapter/pingaccess-sdk-go/v62/services/hsmProviders"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func init() {
	resource.AddTestSweepers("hsm_provider", &resource.Sweeper{
		Name: "hsm_provider",
		F: func(r string) error {
			//svc := hsmProviders.New(conf)
			//results, _, err := svc.GetHsmProvidersCommand(&hsmProviders.GetHsmProvidersCommandInput{Filter: "acctest_"})
			//if err != nil {
			//	return fmt.Errorf("unable to list hsm providers to sweep %s", err)
			//}
			//for _, item := range results.Items {
			//	_, err = svc.DeleteHsmProviderCommand(&hsmProviders.DeleteHsmProviderCommandInput{Id: item.Id.String()})
			//	if err != nil {
			//		return fmt.Errorf("unable to sweep engine listeners %s because %s", item.Id.String(), err)
			//	}
			//}
			return nil
		},
	})
}

func TestAccPingAccessHsmProvider(t *testing.T) {
	if !(paClient{apiVersion: paVersion}).Is61OrAbove() {
		t.Skipf("This test only runs against PingAccess 6.1 or above, not: %s", paVersion)
	}
	resourceName := "pingaccess_hsm_provider.acc_test_hsm"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV5ProviderFactories: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccPingAccessHsmProviderConfig("foo"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessHsmProviderExists(resourceName),
					testAccCheckPingAccessHsmProviderAttributes(resourceName, "foo"),
					resource.TestCheckResourceAttr(resourceName, "name", "acctest_demo"),
					resource.TestCheckResourceAttr(resourceName, "class_name", "com.pingidentity.pa.hsm.pkcs11.plugin.PKCS11HsmProvider"),
					resource.TestCheckResourceAttr(resourceName, "configuration", "{\"library\":\"foo\",\"password\":\"top_secret\",\"slotId\":\"1234\"}"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"configuration"}, //we cant verify passwords
			},
			{
				Config:      testAccPingAccessHsmProviderConfigInvalidClassName(),
				ExpectError: regexp.MustCompile(`unable to find className 'com.pingidentity.pa.hsm.cloudhsm.plugin.foo' available classNames: com.pingidentity.pa.hsm.cloudhsm.plugin.AwsCloudHsmProvider`),
			},
		},
	})
}

func testAccPingAccessHsmProviderConfig(configUpdate string) string {
	return fmt.Sprintf(`
	resource "pingaccess_hsm_provider" "acc_test_hsm" {
	  class_name    = "com.pingidentity.pa.hsm.pkcs11.plugin.PKCS11HsmProvider"
	  name          = "acctest_demo"
	  configuration = <<EOF
	  {
		"slotId": "1234",
		"library": "%s",
		"password": "top_secret"
	  }
	  EOF
	}`, configUpdate)
}

func testAccPingAccessHsmProviderConfigInvalidClassName() string {
	return `
	resource "pingaccess_hsm_provider" "acc_test_hsm" {
		class_name = "com.pingidentity.pa.hsm.cloudhsm.plugin.foo"
		name = "test"
		configuration = <<EOF
		{
			"user": true,
			"password": "sub",
			"partition": "test"
		}
		EOF
	}`
}

func testAccCheckPingAccessHsmProviderExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" || rs.Primary.ID == "0" {
			return fmt.Errorf("No HsmProvider ID is set")
		}

		conn := testAccProvider.Meta().(paClient).HsmProviders
		result, _, err := conn.GetHsmProviderCommand(&hsmProviders.GetHsmProviderCommandInput{
			Id: rs.Primary.ID,
		})

		if err != nil {
			return fmt.Errorf("Error: HsmProvider (%s) not found", n)
		}

		if *result.Name != rs.Primary.Attributes["name"] {
			return fmt.Errorf("Error: HsmProvider response (%s) didnt match state (%s)", *result.Name, rs.Primary.Attributes["name"])
		}

		return nil
	}
}

func testAccCheckPingAccessHsmProviderAttributes(n, library string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs := s.RootModule().Resources[n]
		if rs.Primary.ID == "" || rs.Primary.ID == "0" {
			return fmt.Errorf("No HsmProvider ID is set")
		}

		conn := testAccProvider.Meta().(paClient).HsmProviders
		result, _, err := conn.GetHsmProviderCommand(&hsmProviders.GetHsmProviderCommandInput{
			Id: rs.Primary.ID,
		})

		if err != nil {
			return fmt.Errorf("Error: HsmProvider (%s) not found", n)
		}

		if *result.Name != rs.Primary.Attributes["name"] {
			return fmt.Errorf("error: HsmProvider response (%s) didnt match state (%s)", *result.Name, rs.Primary.Attributes["name"])
		}

		resultMapping := result.Configuration["library"].(string)
		if resultMapping != library {
			return fmt.Errorf("error: HsmProvider response (%s) didnt match state (%s)", resultMapping, library)
		}

		return nil
	}
}
