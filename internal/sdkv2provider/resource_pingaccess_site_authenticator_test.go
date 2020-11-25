package sdkv2provider

import (
	"fmt"
	"net/http"
	"regexp"
	"testing"

	"github.com/iwarapter/pingaccess-sdk-go/pingaccess/models"
	"github.com/iwarapter/pingaccess-sdk-go/services/siteAuthenticators"

	"github.com/google/go-cmp/cmp"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccPingAccessSiteAuthenticator(t *testing.T) {
	resourceName := "pingaccess_site_authenticator.acc_test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV5ProviderFactories: testAccProviders,
		CheckDestroy:             testAccCheckPingAccessSiteAuthenticatorDestroy,
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
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"configuration"}, //we cant verify passwords
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

		conn := testAccProvider.Meta().(paClient).SiteAuthenticators
		result, _, err := conn.GetSiteAuthenticatorCommand(&siteAuthenticators.GetSiteAuthenticatorCommandInput{
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

type siteAuthsMock struct {
	siteAuthenticators.SiteAuthenticatorsAPI
}

func (s siteAuthsMock) GetSiteAuthenticatorDescriptorsCommand() (output *models.DescriptorsView, resp *http.Response, err error) {
	return &models.DescriptorsView{
		Items: []*models.DescriptorView{
			{
				ClassName: String("something"),
				ConfigurationFields: []*models.ConfigurationField{
					{
						Name: String("password"),
						Type: String("CONCEALED"),
					},
				},
				Label: nil,
				Type:  nil,
			},
		}}, nil, err
}

func Test_resourcePingAccessSiteAuthenticatorReadData(t *testing.T) {
	cases := []struct {
		SiteAuthenticator models.SiteAuthenticatorView
	}{
		{
			SiteAuthenticator: models.SiteAuthenticatorView{
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
			resourcePingAccessSiteAuthenticatorReadResult(resourceLocalData, &tc.SiteAuthenticator, siteAuthsMock{})

			if got := *resourcePingAccessSiteAuthenticatorReadData(resourceLocalData); !cmp.Equal(got, tc.SiteAuthenticator) {
				t.Errorf("resourcePingAccessSiteAuthenticatorReadData() = %v", cmp.Diff(got, tc.SiteAuthenticator))
			}
		})
	}
}
