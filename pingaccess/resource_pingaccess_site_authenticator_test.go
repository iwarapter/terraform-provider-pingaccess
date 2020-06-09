package pingaccess

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"regexp"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	pa "github.com/iwarapter/pingaccess-sdk-go/pingaccess"
)

func TestAccPingAccessSiteAuthenticator(t *testing.T) {
	resourceName := "pingaccess_site_authenticator.acc_test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPingAccessSiteAuthenticatorDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingAccessSiteAuthenticatorConfig("acc_test_bar", "bar"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessSiteAuthenticatorExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", "acc_test_bar"),
					resource.TestCheckResourceAttr(resourceName, "class_name", "com.pingidentity.pa.siteauthenticators.BasicAuthTargetSiteAuthenticator"),
					resource.TestCheckResourceAttr(resourceName, "configuration", "{\"password\":{\"value\":\"bar\"},\"username\":\"cheese\"}"),
					resource.TestCheckResourceAttr(resourceName+"_two", "name", "another"),
					resource.TestCheckResourceAttr(resourceName+"_two", "class_name", "com.pingidentity.pa.siteauthenticators.BasicAuthTargetSiteAuthenticator"),
					resource.TestCheckResourceAttr(resourceName+"_two", "configuration", "{\"password\":\"bar\",\"username\":\"cheese\"}"),
				),
			},
			{
				Config: testAccPingAccessSiteAuthenticatorConfig("acc_test_bar", "foo"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessSiteAuthenticatorExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", "acc_test_bar"),
					resource.TestCheckResourceAttr(resourceName, "class_name", "com.pingidentity.pa.siteauthenticators.BasicAuthTargetSiteAuthenticator"),
					resource.TestCheckResourceAttr(resourceName, "configuration", "{\"password\":{\"value\":\"foo\"},\"username\":\"cheese\"}"),
					resource.TestCheckResourceAttr(resourceName+"_two", "name", "another"),
					resource.TestCheckResourceAttr(resourceName+"_two", "class_name", "com.pingidentity.pa.siteauthenticators.BasicAuthTargetSiteAuthenticator"),
					resource.TestCheckResourceAttr(resourceName+"_two", "configuration", "{\"password\":\"foo\",\"username\":\"cheese\"}"),
				),
			},
			{
				Config:      testAccPingAccessSiteAuthenticatorConfigInvalidClassName(),
				ExpectError: regexp.MustCompile(`unable to find className 'com.pingidentity.pa.siteauthenticators.foo' available classNames: com.pingidentity.pa.siteauthenticators.BasicAuthTargetSiteAuthenticator, com.pingidentity.pa.siteauthenticators.MutualTlsSiteAuthenticator, com.pingidentity.pa.siteauthenticators.TokenMediatorSiteAuthenticator`),
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

func testAccPingAccessSiteAuthenticatorConfigInvalidClassName() string {
	return `
	resource "pingaccess_site_authenticator" "acc_test" {
		name          = "break"
		class_name		= "com.pingidentity.pa.siteauthenticators.foo"
		configuration = <<EOF
		{
			"username": "cheese",
			"password": {
				"value": "breaking"
			}
		}
		EOF
	}`
}

func testAccCheckPingAccessSiteAuthenticatorExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" || rs.Primary.ID == "0" {
			return fmt.Errorf("No site_authenticator ID is set")
		}

		conn := testAccProvider.Meta().(*pa.Client).SiteAuthenticators
		result, _, err := conn.GetSiteAuthenticatorCommand(&pa.GetSiteAuthenticatorCommandInput{
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

	descs := pa.DescriptorsView{
		Items: []*pa.DescriptorView{
			{
				ClassName: String("something"),
				ConfigurationFields: []*pa.ConfigurationField{
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
	c := pa.NewClient("", "", url, "", server.Client())

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
			resourceSchema := resourcePingAccessSiteAuthenticatorSchema()
			resourceLocalData := schema.TestResourceDataRaw(t, resourceSchema, map[string]interface{}{})
			resourcePingAccessSiteAuthenticatorReadResult(resourceLocalData, &tc.SiteAuthenticator, c.SiteAuthenticators)

			if got := *resourcePingAccessSiteAuthenticatorReadData(resourceLocalData); !cmp.Equal(got, tc.SiteAuthenticator) {
				t.Errorf("resourcePingAccessSiteAuthenticatorReadData() = %v", cmp.Diff(got, tc.SiteAuthenticator))
			}
		})
	}
}
