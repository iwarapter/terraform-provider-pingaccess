package sdkv2provider

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/iwarapter/pingaccess-sdk-go/v62/services/highAvailability"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func init() {
	resource.AddTestSweepers("load_balancing_strategy", &resource.Sweeper{
		Name: "load_balancing_strategy",
		F: func(r string) error {
			svc := highAvailability.New(conf)
			results, _, err := svc.GetLoadBalancingStrategiesCommand(&highAvailability.GetLoadBalancingStrategiesCommandInput{Filter: "acctest_"})
			if err != nil {
				return fmt.Errorf("unable to list load_balancing_strategy to sweep %s", err)
			}
			for _, item := range results.Items {
				_, err = svc.DeleteLoadBalancingStrategyCommand(&highAvailability.DeleteLoadBalancingStrategyCommandInput{Id: item.Id.String()})
				if err != nil {
					return fmt.Errorf("unable to sweep load_balancing_strategy %s because %s", item.Id.String(), err)
				}
			}
			return nil
		},
	})
}

func TestAccPingAccessLoadBalancingStrategy(t *testing.T) {
	resourceName := "pingaccess_load_balancing_strategy.test"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV5ProviderFactories: testAccProviders,
		CheckDestroy:             testAccCheckPingAccessLoadBalancingStrategyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingAccessLoadBalancingStrategyConfig("foo"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessLoadBalancingStrategyExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", "acctest_example"),
					resource.TestCheckResourceAttr(resourceName, "class_name", "com.pingidentity.pa.ha.lb.header.HeaderBasedLoadBalancingPlugin"),
					resource.TestCheckResourceAttr(resourceName, "configuration", "{\"fallbackToFirstAvailableHost\":false,\"headerName\":\"foo\"}"),
				),
			},
			{
				Config: testAccPingAccessLoadBalancingStrategyConfig("bar"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessLoadBalancingStrategyExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", "acctest_example"),
					resource.TestCheckResourceAttr(resourceName, "class_name", "com.pingidentity.pa.ha.lb.header.HeaderBasedLoadBalancingPlugin"),
					resource.TestCheckResourceAttr(resourceName, "configuration", "{\"fallbackToFirstAvailableHost\":false,\"headerName\":\"bar\"}"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config:      testAccPingAccessLoadBalancingStrategyConfigInvalidClassName(),
				ExpectError: regexp.MustCompile(`unable to find className 'com.pingidentity.pa.ha.lb.header.foo' available classNames: com.pingidentity.pa.ha.lb.header.HeaderBasedLoadBalancingPlugin`),
			},
		},
	})
}

func testAccCheckPingAccessLoadBalancingStrategyDestroy(s *terraform.State) error {
	return nil
}

func testAccPingAccessLoadBalancingStrategyConfig(configUpdate string) string {
	return fmt.Sprintf(`
	resource "pingaccess_load_balancing_strategy" "test" {
		name = "acctest_example"
		class_name = "com.pingidentity.pa.ha.lb.header.HeaderBasedLoadBalancingPlugin"
		configuration = <<EOF
		{
			"headerName": "%s",
			"fallbackToFirstAvailableHost": false
		}
		EOF
	}`, configUpdate)
}

func testAccPingAccessLoadBalancingStrategyConfigInvalidClassName() string {
	return `
	resource "pingaccess_load_balancing_strategy" "test" {
		class_name = "com.pingidentity.pa.ha.lb.header.foo"
		name = "foo"

		configuration = <<EOF
		{
			"headerName": "foo",
			"fallbackToFirstAvailableHost": false
		}
		EOF
	}`
}

func testAccCheckPingAccessLoadBalancingStrategyExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" || rs.Primary.ID == "0" {
			return fmt.Errorf("No load_balancing_strategy ID is set")
		}

		conn := testAccProvider.Meta().(paClient).HighAvailability
		result, _, err := conn.GetLoadBalancingStrategyCommand(&highAvailability.GetLoadBalancingStrategyCommandInput{
			Id: rs.Primary.ID,
		})

		if err != nil {
			return fmt.Errorf("Error: LoadBalancingStrategy (%s) not found", n)
		}

		if *result.Name != rs.Primary.Attributes["name"] {
			return fmt.Errorf("Error: LoadBalancingStrategy response (%s) didnt match state (%s)", *result.Name, rs.Primary.Attributes["name"])
		}

		return nil
	}
}
