package sdkv2provider

import (
	"fmt"
	"testing"

	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess/models"
	"github.com/iwarapter/pingaccess-sdk-go/v62/services/thirdPartyServices"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func init() {
	resource.AddTestSweepers("third_party_service", &resource.Sweeper{
		Name: "third_party_service",
		F: func(r string) error {
			svc := thirdPartyServices.New(conf)
			results, _, err := svc.GetThirdPartyServicesCommand(&thirdPartyServices.GetThirdPartyServicesCommandInput{Filter: "acctest_"})
			if err != nil {
				return fmt.Errorf("unable to list third_party_services to sweep %s", err)
			}
			for _, item := range results.Items {
				_, err = svc.DeleteThirdPartyServiceCommand(&thirdPartyServices.DeleteThirdPartyServiceCommandInput{Id: *item.Id})
				if err != nil {
					return fmt.Errorf("unable to sweep third_party_service %s because %s", *item.Id, err)
				}
			}
			return nil
		},
	})
}

func TestAccPingAccessThirdPartyService(t *testing.T) {
	resourceName := "pingaccess_third_party_service.demo_tps"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV5ProviderFactories: testAccProviders,
		CheckDestroy:             testAccCheckPingAccessThirdPartyServiceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingAccessThirdPartyServiceConfig("acctest_service", "localhost:1234"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessThirdPartyServiceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", "acctest_service"),
					resource.TestCheckResourceAttr(resourceName, "secure", "false"),
					resource.TestCheckResourceAttr(resourceName, "trusted_certificate_group_id", "0"),
					resource.TestCheckResourceAttr(resourceName, "max_connections", "-1"),
					resource.TestCheckResourceAttr(resourceName, "skip_hostname_verification", "false"),
					resource.TestCheckNoResourceAttr(resourceName, "expected_hostname"),
					resource.TestCheckResourceAttr(resourceName, "availability_profile_id", "1"),
					resource.TestCheckResourceAttr(resourceName, "load_balancing_strategy_id", "0"),
					resource.TestCheckResourceAttr(resourceName, "use_proxy", "false"),
					resource.TestCheckNoResourceAttr(resourceName, "host_value"),
				),
			},
			{
				Config: testAccPingAccessThirdPartyServiceConfig("acctest_service", "localhost:1235"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessThirdPartyServiceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", "acctest_service"),
					resource.TestCheckResourceAttr(resourceName, "secure", "false"),
					resource.TestCheckResourceAttr(resourceName, "trusted_certificate_group_id", "0"),
					resource.TestCheckResourceAttr(resourceName, "max_connections", "-1"),
					resource.TestCheckResourceAttr(resourceName, "skip_hostname_verification", "false"),
					resource.TestCheckNoResourceAttr(resourceName, "expected_hostname"),
					resource.TestCheckResourceAttr(resourceName, "availability_profile_id", "1"),
					resource.TestCheckResourceAttr(resourceName, "load_balancing_strategy_id", "0"),
					resource.TestCheckResourceAttr(resourceName, "use_proxy", "false"),
					resource.TestCheckNoResourceAttr(resourceName, "host_value"),
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

func testAccCheckPingAccessThirdPartyServiceDestroy(s *terraform.State) error {
	return nil
}

func testAccPingAccessThirdPartyServiceConfig(name, target string) string {
	return fmt.Sprintf(`
	resource "pingaccess_third_party_service" "demo_tps" {
		name = "%s"
		targets = [
			"%s"
		]
	}`, name, target)
}

func testAccCheckPingAccessThirdPartyServiceExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" || rs.Primary.ID == "0" {
			return fmt.Errorf("no third party service ID is set")
		}

		conn := testAccProvider.Meta().(paClient).ThirdPartyServices
		result, _, err := conn.GetThirdPartyServiceCommand(&thirdPartyServices.GetThirdPartyServiceCommandInput{
			Id: rs.Primary.ID,
		})

		if err != nil {
			return fmt.Errorf("error: ThirdPartyService (%s) not found", n)
		}

		if *result.Name != rs.Primary.Attributes["name"] {
			return fmt.Errorf("error: ThirdPartyService response (%s) didnt match state (%s)", *result.Name, rs.Primary.Attributes["name"])
		}

		return nil
	}
}

func Test_resourcePingAccessThirdPartyServiceReadData(t *testing.T) {
	cases := []struct {
		ThirdPartyService models.ThirdPartyServiceView
	}{
		{
			ThirdPartyService: models.ThirdPartyServiceView{
				Name:                      String("localhost"),
				Targets:                   &[]*string{String("localhost:1234")},
				AvailabilityProfileId:     Int(1),
				MaxConnections:            Int(-1),
				Secure:                    Bool(false),
				SkipHostnameVerification:  Bool(false),
				TrustedCertificateGroupId: Int(0),
				UseProxy:                  Bool(false),
				LoadBalancingStrategyId:   Int(0),
				ExpectedHostname:          nil,
				HostValue:                 nil,
			},
		},
		{
			ThirdPartyService: models.ThirdPartyServiceView{
				Name:                      String("localhost"),
				Targets:                   &[]*string{String("localhost:1234")},
				AvailabilityProfileId:     Int(0),
				ExpectedHostname:          String("localhost"),
				HostValue:                 String("localhost"),
				LoadBalancingStrategyId:   Int(1),
				MaxConnections:            Int(10),
				Secure:                    Bool(true),
				SkipHostnameVerification:  Bool(true),
				TrustedCertificateGroupId: Int(0),
				UseProxy:                  Bool(false),
			},
		},
	}
	for i, tc := range cases {
		t.Run(fmt.Sprintf("tc:%v", i), func(t *testing.T) {

			resourceSchema := resourcePingAccessThirdPartyServiceSchema()
			resourceLocalData := schema.TestResourceDataRaw(t, resourceSchema, map[string]interface{}{})
			resourcePingAccessThirdPartyServiceReadResult(resourceLocalData, &tc.ThirdPartyService)

			if got := *resourcePingAccessThirdPartyServiceReadData(resourceLocalData); !cmp.Equal(got, tc.ThirdPartyService) {
				t.Errorf("resourcePingAccessThirdPartyServiceReadData() = %v", cmp.Diff(got, tc.ThirdPartyService))
			}
		})
	}
}
