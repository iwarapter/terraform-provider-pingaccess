package pingaccess

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/iwarapter/pingaccess-sdk-go/pingaccess"
	pa "github.com/iwarapter/pingaccess-sdk-go/pingaccess"
)

func init() {
	resource.AddTestSweepers("pingaccess_site_authenticator", &resource.Sweeper{
		Name:         "pingaccess_site_authenticator",
		F:            testSweepSiteAuthenticators,
		Dependencies: []string{"pingaccess_site"},
	})
}

func testSweepSiteAuthenticators(r string) error {
	url, _ := url.Parse("https://localhost:9000")
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	conn := pingaccess.NewClient("Administrator", "2Access2", url, "/pa-admin-api/v3", nil).SiteAuthenticators
	result, _, _ := conn.GetSiteAuthenticatorsCommand(&pingaccess.GetSiteAuthenticatorsCommandInput{Filter: "acc_test_"})
	for _, v := range result.Items {
		log.Printf("Sweeper: Deleting %s", *v.Name)
		conn.DeleteSiteAuthenticatorCommand(&pingaccess.DeleteSiteAuthenticatorCommandInput{Id: v.Id.String()})
	}
	return nil
}

func TestAccPingAccessSiteAuthenticator(t *testing.T) {
	var out pingaccess.SiteAuthenticatorView

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPingAccessSiteAuthenticatorDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingAccessSiteAuthenticatorConfig("acc_test_bar", "bar"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessSiteAuthenticatorExists("pingaccess_site_authenticator.acc_test", 3, &out),
				),
			},
			{
				Config: testAccPingAccessSiteAuthenticatorConfig("acc_test_bar", "foo"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessSiteAuthenticatorExists("pingaccess_site_authenticator.acc_test", 6, &out),
				),
			},
		},
	})
}

func testAccCheckPingAccessSiteAuthenticatorDestroy(s *terraform.State) error {
	return nil
}

func testAccPingAccessSiteAuthenticatorConfig(name, password string) string {
	return fmt.Sprintf(`
	resource "pingaccess_site_authenticator" "acc_test" {
		name          = "%s"
		class_name		= "com.pingidentity.pa.siteauthenticators.BasicAuthTargetSiteAuthenticator"
		configuration = <<EOF
		{
			"username": "cheese",
			"password": {
				"value": "%s"
			}
		}
		EOF
	}

	resource "pingaccess_site_authenticator" "acc_test_two" {
		name          = "another"
		class_name		= "com.pingidentity.pa.siteauthenticators.BasicAuthTargetSiteAuthenticator"
		configuration = <<EOF
		{
			"username": "cheese",
			"password": "%s"
		}
		EOF
	}`, name, password, password)
}

func testAccCheckPingAccessSiteAuthenticatorExists(n string, c int64, out *pingaccess.SiteAuthenticatorView) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" || rs.Primary.ID == "0" {
			return fmt.Errorf("No site_authenticator ID is set")
		}

		conn := testAccProvider.Meta().(*pingaccess.Client).SiteAuthenticators
		result, _, err := conn.GetSiteAuthenticatorCommand(&pingaccess.GetSiteAuthenticatorCommandInput{
			Id: rs.Primary.ID,
		})

		if err != nil {
			return fmt.Errorf("Error: SiteAuthenticator (%s) not found", n)
		}

		if *result.Name != rs.Primary.Attributes["name"] {
			return fmt.Errorf("Error: SiteAuthenticator response (%s) didnt match state (%s)", *result.Name, rs.Primary.Attributes["name"])
		}

		return nil
	}
}

func Test_resourcePingAccessSiteAuthenticatorReadData(t *testing.T) {
	cases := []struct {
		SiteAuthenticator pa.SiteAuthenticatorView
	}{
		{
			SiteAuthenticator: pa.SiteAuthenticatorView{
				Name:      String("demo"),
				ClassName: String("something"),
				Configuration: map[string]interface{}{
					"foo": "bar",
					"password": map[string]interface{}{
						"value": "top-secret",
					},
				},
			},
		},
	}
	for i, tc := range cases {
		t.Run(fmt.Sprintf("tc:%v", i), func(t *testing.T) {

			descs := pingaccess.DescriptorsView{
				Items: []*pingaccess.DescriptorView{
					{
						ClassName: String("something"),
						ConfigurationFields: []*pingaccess.ConfigurationField{
							{
								Name: String("password"),
								Type: String("CONCEALED"),
							},
						},
						Label: nil,
						Type:  nil,
					},
				}}

			server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
				// Test request parameters
				equals(t, req.URL.String(), "/siteAuthenticators/descriptors")
				// Send response to be tested
				b, _ := json.Marshal(descs)
				rw.Write(b)
			}))
			// Close the server when test finishes
			defer server.Close()

			// Use Client & URL from our local test server
			url, _ := url.Parse(server.URL)
			c := pingaccess.NewClient("", "", url, "", server.Client())

			resourceSchema := resourcePingAccessSiteAuthenticatorSchema()
			resourceLocalData := schema.TestResourceDataRaw(t, resourceSchema, map[string]interface{}{})
			//resourceLocalData.Set("hidden_fields", []string{"password"})
			resourcePingAccessSiteAuthenticatorReadResult(resourceLocalData, &tc.SiteAuthenticator, c.SiteAuthenticators)

			if got := *resourcePingAccessSiteAuthenticatorReadData(resourceLocalData); !cmp.Equal(got, tc.SiteAuthenticator) {
				t.Errorf("resourcePingAccessSiteAuthenticatorReadData() = %v", cmp.Diff(got, tc.SiteAuthenticator))
			}
		})
	}
}
