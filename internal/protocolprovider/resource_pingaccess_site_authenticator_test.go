package protocol

import (
	"context"
	"fmt"
	"regexp"
	"testing"

	"github.com/iwarapter/terraform-provider-pingaccess/internal/sdkv2provider"

	"github.com/hashicorp/terraform-plugin-go/tfprotov5"

	tfmux "github.com/hashicorp/terraform-plugin-mux"
	"github.com/iwarapter/pingaccess-sdk-go/v62/services/siteAuthenticators"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func init() {
	resource.AddTestSweepers("site_authenticator", &resource.Sweeper{
		Name: "site_authenticator",
		F: func(r string) error {
			svc := siteAuthenticators.New(conf)
			results, _, err := svc.GetSiteAuthenticatorsCommand(&siteAuthenticators.GetSiteAuthenticatorsCommandInput{Filter: "acctest_"})
			if err != nil {
				return fmt.Errorf("unable to list site_authenticators to sweep %s", err)
			}
			for _, item := range results.Items {
				_, err = svc.DeleteSiteAuthenticatorCommand(&siteAuthenticators.DeleteSiteAuthenticatorCommandInput{Id: item.Id.String()})
				if err != nil {
					return fmt.Errorf("unable to sweep site_authenticator %s because %s", item.Id.String(), err)
				}
			}
			return nil
		},
	})
}

func TestAccPingAccessSiteAuthenticator(t *testing.T) {
	resourceName := "pingaccess_site_authenticator.test"

	resource.Test(t, resource.TestCase{
		ProtoV5ProviderFactories: map[string]func() (tfprotov5.ProviderServer, error){
			"pingaccess": func() (tfprotov5.ProviderServer, error) {
				return Server(), nil
			},
		},
		CheckDestroy: testAccCheckPingAccessSiteAuthenticatorDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingAccessSiteAuthenticatorConfig("acctest_foo", `<<EOF
{
	"username": "cheese",
	"password": {
		"value": "secret"
	}
}
EOF`),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessSiteAuthenticatorExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", "acctest_foo"),
					resource.TestCheckResourceAttr(resourceName, "class_name", "com.pingidentity.pa.siteauthenticators.BasicAuthTargetSiteAuthenticator"),
					resource.TestCheckResourceAttr(resourceName, "configuration", "{\n\t\"username\": \"cheese\",\n\t\"password\": {\n\t\t\"value\": \"secret\"\n\t}\n}\n"),
				),
			},
			{
				Config: testAccPingAccessSiteAuthenticatorConfig("acctest_foo", `<<EOF
{
	"username": "cheese",
	"password": {
		"value": "secret2"
	}
}
EOF`),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessSiteAuthenticatorExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", "acctest_foo"),
					resource.TestCheckResourceAttr(resourceName, "class_name", "com.pingidentity.pa.siteauthenticators.BasicAuthTargetSiteAuthenticator"),
					resource.TestCheckResourceAttr(resourceName, "configuration", "{\n\t\"username\": \"cheese\",\n\t\"password\": {\n\t\t\"value\": \"secret2\"\n\t}\n}\n"),
				),
			},
			{
				Config: testAccPingAccessSiteAuthenticatorConfig("acctest_foo", `<<EOF
{
	"username": "cheese",
	"password": "secret2"
}
EOF`),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessSiteAuthenticatorExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", "acctest_foo"),
					resource.TestCheckResourceAttr(resourceName, "class_name", "com.pingidentity.pa.siteauthenticators.BasicAuthTargetSiteAuthenticator"),
					resource.TestCheckResourceAttr(resourceName, "configuration", "{\n\t\"username\": \"cheese\",\n\t\"password\": \"secret2\"\n}\n"),
				),
			},
			//TODO we dont support json style imports
			//{
			//	ResourceName:      resourceName,
			//	ImportState:       true,
			//	ImportStateVerify: true,
			//},
			{
				Config: testAccPingAccessSiteAuthenticatorConfigInvalidClassName(`<<EOF
		{
			"description": null,
			"path": "/foo",
			"subjectAttributeName": "foo",
			"issuer": null,
			"audience": null
		}
		EOF`),
				ExpectError: regexp.MustCompile(`unable to find className 'com.pingidentity.pa.siteauthenticators.foo'\navailable classNames:\ncom.pingidentity.pa.siteauthenticators.BasicAuthTargetSiteAuthenticator`),
			},
		},
	})
}

func TestAccPingAccessSiteAuthenticatorWithDynamicPsuedoType(t *testing.T) {
	resourceName := "pingaccess_site_authenticator.test"

	resource.Test(t, resource.TestCase{
		ProtoV5ProviderFactories: map[string]func() (tfprotov5.ProviderServer, error){
			"pingaccess": func() (tfprotov5.ProviderServer, error) {
				return Server(), nil
			},
		},
		CheckDestroy: testAccCheckPingAccessSiteAuthenticatorDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingAccessSiteAuthenticatorConfig("acctest_bar", `{
			"username": "cheese",
			"password": {
				"value": "secret"
			}
		}`),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessSiteAuthenticatorExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", "acctest_bar"),
					resource.TestCheckResourceAttr(resourceName, "class_name", "com.pingidentity.pa.siteauthenticators.BasicAuthTargetSiteAuthenticator"),
					resource.TestCheckResourceAttr(resourceName, "configuration.username", "cheese"),
					resource.TestCheckResourceAttr(resourceName, "configuration.password.value", "secret"),
				),
			},
			{
				Config: testAccPingAccessSiteAuthenticatorConfig("acctest_bar", `{
			"username": "cheese",
			"password": {
				"value": "secret2"
			}
		}`),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessSiteAuthenticatorExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", "acctest_bar"),
					resource.TestCheckResourceAttr(resourceName, "class_name", "com.pingidentity.pa.siteauthenticators.BasicAuthTargetSiteAuthenticator"),
					resource.TestCheckResourceAttr(resourceName, "configuration.username", "cheese"),
					resource.TestCheckResourceAttr(resourceName, "configuration.password.value", "secret2"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccPingAccessSiteAuthenticatorConfig("acctest_bar", `{
			"username": "cheese",
		}`),
				ExpectError: regexp.MustCompile(`the field 'password' is required for the class_name\n'com.pingidentity.pa.siteauthenticators.BasicAuthTargetSiteAuthenticator'`),
			},
			{
				Config: testAccPingAccessSiteAuthenticatorConfigInvalidClassName(`{
			"username": "cheese",
			"password": {
				"value": "secret"
			}
		}`),
				ExpectError: regexp.MustCompile(`unable to find className 'com.pingidentity.pa.siteauthenticators.foo'\navailable classNames:\ncom.pingidentity.pa.siteauthenticators.BasicAuthTargetSiteAuthenticator`),
			},
		},
	})
}

func TestAccPingAccessSiteAuthenticatorIssue72(t *testing.T) {
	resource1 := "pingaccess_site_authenticator.test1"
	resource2 := "pingaccess_site_authenticator.test2"

	resource.Test(t, resource.TestCase{
		ProtoV5ProviderFactories: map[string]func() (tfprotov5.ProviderServer, error){
			//"pingaccess": func() (tfprotov5.ProviderServer, error) {
			//	return Server(), nil
			//},
			"pingaccess": func() (tfprotov5.ProviderServer, error) {
				ctx := context.Background()
				sdkv2 := sdkv2provider.Provider()
				factory, err := tfmux.NewSchemaServerFactory(ctx, sdkv2.GRPCProvider, Server)
				if err != nil {
					return nil, err
				}
				return factory.Server(), nil
			},
		},
		CheckDestroy: testAccCheckPingAccessSiteAuthenticatorDestroy,
		Steps: []resource.TestStep{
			{
				Config: `
resource "pingaccess_site_authenticator" "test1" {
  name = "issue 72 heredoc"
  class_name = "com.pingidentity.pa.siteauthenticators.MutualTlsSiteAuthenticator"
  configuration = <<EOF
  {
    "keyPairId": "${pingaccess_keypair.demo_keypair.id}"
  }
  EOF
}

resource "pingaccess_site_authenticator" "test2" {
  name = "issue 72 dynamic"
  class_name = "com.pingidentity.pa.siteauthenticators.MutualTlsSiteAuthenticator"
  configuration = {
    "keyPairId" = pingaccess_keypair.demo_keypair.id
  }
}

resource "pingaccess_keypair" "demo_keypair" {
  alias             = "demo2"
  city              = "London"
  common_name       = "Example"
  country           = "GB"
  key_algorithm     = "RSA"
  key_size          = 2048
  organization      = "Test"
  organization_unit = "Development"
  state             = "London"
  valid_days        = 365
}

`,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessSiteAuthenticatorExists(resource1),
					resource.TestCheckResourceAttr(resource1, "name", "issue 72 heredoc"),
					resource.TestCheckResourceAttr(resource1, "class_name", "com.pingidentity.pa.siteauthenticators.MutualTlsSiteAuthenticator"),
					resource.TestCheckResourceAttrSet(resource1, "configuration"),

					testAccCheckPingAccessSiteAuthenticatorExists(resource2),
					resource.TestCheckResourceAttr(resource2, "name", "issue 72 dynamic"),
					resource.TestCheckResourceAttr(resource2, "class_name", "com.pingidentity.pa.siteauthenticators.MutualTlsSiteAuthenticator"),
					resource.TestCheckResourceAttrSet(resource2, "configuration.keyPairId"),
				),
			},
		},
	})
}

func testAccCheckPingAccessSiteAuthenticatorDestroy(s *terraform.State) error {
	return nil
}

func testAccPingAccessSiteAuthenticatorConfig(name, configUpdate string) string {
	return fmt.Sprintf(`
	resource "pingaccess_site_authenticator" "test" {
		class_name = "com.pingidentity.pa.siteauthenticators.BasicAuthTargetSiteAuthenticator"
		name = "%s"

		configuration = %s
	}
`, name, configUpdate)
}

func testAccPingAccessSiteAuthenticatorConfigInvalidClassName(configUpdate string) string {
	return fmt.Sprintf(`
	resource "pingaccess_site_authenticator" "test" {
		class_name		= "com.pingidentity.pa.siteauthenticators.foo"
		name = "acctest_foo"
		configuration = %s
	}`, configUpdate)
}

func testAccCheckPingAccessSiteAuthenticatorExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" || rs.Primary.ID == "0" {
			return fmt.Errorf("no site_authenticator ID is set")
		}

		conn := siteAuthenticators.New(conf)
		result, _, err := conn.GetSiteAuthenticatorCommand(&siteAuthenticators.GetSiteAuthenticatorCommandInput{
			Id: rs.Primary.ID,
		})

		if err != nil {
			return fmt.Errorf("error: SiteAuthenticator (%s) not found", n)
		}

		if *result.Name != rs.Primary.Attributes["name"] {
			return fmt.Errorf("error: SiteAuthenticator response (%s) didnt match state (%s)", *result.Name, rs.Primary.Attributes["name"])
		}

		return nil
	}
}
