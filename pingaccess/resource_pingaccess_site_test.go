package pingaccess

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/iwarapter/pingaccess-sdk-go/pingaccess"
	pa "github.com/iwarapter/pingaccess-sdk-go/pingaccess"
)

func init() {
	resource.AddTestSweepers("pingaccess_sites", &resource.Sweeper{
		Name:         "pingaccess_sites",
		F:            testSweepSites,
		Dependencies: []string{"pingaccess_application", "pingaccess_application_resource"},
	})
}

func testSweepSites(r string) error {
	url, _ := url.Parse("https://localhost:9000")
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	conn := pingaccess.NewClient("Administrator", "2Access2", url, "/pa-admin-api/v3", nil).Sites
	result, _, _ := conn.GetSitesCommand(&pingaccess.GetSitesCommandInput{Filter: "acc_test_"})
	for _, v := range result.Items {
		log.Printf("Sweeper: Deleting %s", *v.Name)
		conn.DeleteSiteCommand(&pingaccess.DeleteSiteCommandInput{Id: v.Id.String()})
	}
	return nil
}

func TestAccPingAccessSite(t *testing.T) {
	var out pingaccess.SiteView

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPingAccessSiteDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingAccessSiteConfig("acc_test_bar", []string{"\"localhost:1234\""}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessSiteExists("pingaccess_site.acc_test", 3, &out),
					// testAccCheckPingAccessSiteAttributes(),
					// testAccCheckAWSPolicyAttachmentAttributes([]string{userName}, []string{roleName}, []string{groupName}, &out),
				),
			},
			{
				Config: testAccPingAccessSiteConfig("acc_test_bar", []string{"\"localhost:1235\""}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessSiteExists("pingaccess_site.acc_test", 6, &out),
					// testAccCheckAWSPolicyAttachmentAttributes([]string{userName2, userName3},
					// 	[]string{roleName2, roleName3}, []string{groupName2, groupName3}, &out),
				),
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
		targets                      = [%s]
		max_connections              = -1
		max_web_socket_connections   = -1
		availability_profile_id      = 1
		use_target_host_header     	 = false
	}`, name, strings.Join(targets, ","))
}

func testAccCheckPingAccessSiteExists(n string, c int64, out *pingaccess.SiteView) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" || rs.Primary.ID == "0" {
			return fmt.Errorf("No site ID is set")
		}

		conn := testAccProvider.Meta().(*pingaccess.Client).Sites
		result, _, err := conn.GetSiteCommand(&pingaccess.GetSiteCommandInput{
			Id: rs.Primary.ID,
		})

		if err != nil {
			return fmt.Errorf("Error: Site (%s) not found", n)
		}

		if *result.Name != rs.Primary.Attributes["name"] {
			return fmt.Errorf("Error: Site response (%s) didnt match state (%s)", *result.Name, rs.Primary.Attributes["name"])
		}

		return nil
	}
}

func Test_resourcePingAccessSiteReadData(t *testing.T) {
	cases := []struct {
		Site pa.SiteView
	}{
		{
			Site: pa.SiteView{
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
			Site: pa.SiteView{
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
