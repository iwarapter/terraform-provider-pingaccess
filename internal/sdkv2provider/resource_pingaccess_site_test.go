package sdkv2provider

import (
	"fmt"
	"strings"
	"testing"

	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess/models"
	"github.com/iwarapter/pingaccess-sdk-go/v62/services/sites"

	"github.com/google/go-cmp/cmp"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func init() {
	resource.AddTestSweepers("sites", &resource.Sweeper{
		Name: "sites",
		F: func(r string) error {
			svc := sites.New(conf)
			results, _, err := svc.GetSitesCommand(&sites.GetSitesCommandInput{Filter: "acctest_"})
			if err != nil {
				return fmt.Errorf("unable to list sites to sweep %s", err)
			}
			for _, item := range results.Items {
				_, err = svc.DeleteSiteCommand(&sites.DeleteSiteCommandInput{Id: item.Id.String()})
				if err != nil {
					return fmt.Errorf("unable to sweep site %s because %s", item.Id.String(), err)
				}
			}
			return nil
		},
	})
}

func TestAccPingAccessSite(t *testing.T) {
	resourceName := "pingaccess_site.acc_test"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV5ProviderFactories: testAccProviders,
		CheckDestroy:             testAccCheckPingAccessSiteDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingAccessSiteConfig("acctest_bar", []string{"localhost:1234"}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessSiteExists("pingaccess_site.acc_test"),
					resource.TestCheckResourceAttr(resourceName, "availability_profile_id", "1"),
					resource.TestCheckResourceAttr(resourceName, "keep_alive_timeout", "0"),
					resource.TestCheckResourceAttr(resourceName, "load_balancing_strategy_id", "0"),
					resource.TestCheckResourceAttr(resourceName, "max_connections", "-1"),
					resource.TestCheckResourceAttr(resourceName, "max_web_socket_connections", "-1"),
					resource.TestCheckResourceAttr(resourceName, "name", "acctest_bar"),
					resource.TestCheckResourceAttr(resourceName, "secure", "false"),
					resource.TestCheckResourceAttr(resourceName, "send_pa_cookie", "true"),
					resource.TestCheckResourceAttr(resourceName, "skip_hostname_verification", "false"),
					resource.TestCheckResourceAttr(resourceName, "targets.0", "localhost:1234"),
					resource.TestCheckResourceAttr(resourceName, "trusted_certificate_group_id", "0"),
					resource.TestCheckResourceAttr(resourceName, "use_proxy", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_target_host_header", "false"),
				),
			},
			{
				Config: testAccPingAccessSiteConfig("acctest_foo", []string{"localhost:1235"}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessSiteExists("pingaccess_site.acc_test"),
					resource.TestCheckResourceAttr(resourceName, "availability_profile_id", "1"),
					resource.TestCheckResourceAttr(resourceName, "keep_alive_timeout", "0"),
					resource.TestCheckResourceAttr(resourceName, "load_balancing_strategy_id", "0"),
					resource.TestCheckResourceAttr(resourceName, "max_connections", "-1"),
					resource.TestCheckResourceAttr(resourceName, "max_web_socket_connections", "-1"),
					resource.TestCheckResourceAttr(resourceName, "name", "acctest_foo"),
					resource.TestCheckResourceAttr(resourceName, "secure", "false"),
					resource.TestCheckResourceAttr(resourceName, "send_pa_cookie", "true"),
					resource.TestCheckResourceAttr(resourceName, "skip_hostname_verification", "false"),
					resource.TestCheckResourceAttr(resourceName, "targets.0", "localhost:1235"),
					resource.TestCheckResourceAttr(resourceName, "trusted_certificate_group_id", "0"),
					resource.TestCheckResourceAttr(resourceName, "use_proxy", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_target_host_header", "false"),
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

func testAccCheckPingAccessSiteDestroy(s *terraform.State) error {
	return nil
}

func testAccPingAccessSiteConfig(name string, targets []string) string {
	return fmt.Sprintf(`
	resource "pingaccess_site" "acc_test" {
		name                         = "%s"
		targets                      = ["%s"]
		max_connections              = -1
		max_web_socket_connections   = -1
		availability_profile_id      = 1
		use_target_host_header     	 = false
	}`, name, strings.Join(targets, ","))
}

func testAccCheckPingAccessSiteExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("unable to find resource: %s", n)
		}

		if rs.Primary.ID == "" || rs.Primary.ID == "0" {
			return fmt.Errorf("no site ID is set")
		}

		conn := testAccProvider.Meta().(paClient).Sites
		result, _, err := conn.GetSiteCommand(&sites.GetSiteCommandInput{
			Id: rs.Primary.ID,
		})

		if err != nil {
			return fmt.Errorf("unable to find Site (%s) with ID: %s", n, rs.Primary.ID)
		}

		if *result.Name != rs.Primary.Attributes["name"] {
			return fmt.Errorf("error Site response (%s) didnt match state (%s)", *result.Name, rs.Primary.Attributes["name"])
		}

		return nil
	}
}

func Test_resourcePingAccessSiteReadData(t *testing.T) {
	cases := []struct {
		Site models.SiteView
	}{
		{
			Site: models.SiteView{
				Name:                      String("demo"),
				Targets:                   &[]*string{String("localhost:1234")},
				AvailabilityProfileId:     Int(1),
				ExpectedHostname:          String("true"),
				KeepAliveTimeout:          Int(1),
				LoadBalancingStrategyId:   Int(1),
				MaxConnections:            Int(1),
				MaxWebSocketConnections:   Int(1),
				Secure:                    Bool(true),
				SendPaCookie:              Bool(true),
				SiteAuthenticatorIds:      &[]*int{Int(1)},
				SkipHostnameVerification:  Bool(true),
				TrustedCertificateGroupId: Int(1),
				UseProxy:                  Bool(true),
				UseTargetHostHeader:       Bool(true),
			},
		},
		{
			Site: models.SiteView{
				Name:                      String("demo"),
				Targets:                   &[]*string{String("localhost:1234")},
				AvailabilityProfileId:     Int(0),
				ExpectedHostname:          String(""),
				KeepAliveTimeout:          Int(0),
				LoadBalancingStrategyId:   Int(0),
				MaxConnections:            Int(0),
				MaxWebSocketConnections:   Int(0),
				Secure:                    Bool(false),
				SendPaCookie:              Bool(false),
				SiteAuthenticatorIds:      &[]*int{},
				SkipHostnameVerification:  Bool(false),
				TrustedCertificateGroupId: Int(0),
				UseProxy:                  Bool(false),
				UseTargetHostHeader:       Bool(false),
			},
		},
	}
	for i, tc := range cases {
		t.Run(fmt.Sprintf("tc:%v", i), func(t *testing.T) {

			resourceSchema := resourcePingAccessSiteSchema()
			resourceLocalData := schema.TestResourceDataRaw(t, resourceSchema, map[string]interface{}{})
			resourcePingAccessSiteReadResult(resourceLocalData, &tc.Site)

			if got := *resourcePingAccessSiteReadData(resourceLocalData); !cmp.Equal(got, tc.Site) {
				t.Errorf("resourcePingAccessSiteReadData() = %v", cmp.Diff(got, tc.Site))
			}
		})
	}
}
