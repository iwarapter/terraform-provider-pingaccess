package protocol

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-go/tfprotov5"

	"github.com/iwarapter/pingaccess-sdk-go/services/accessTokenValidators"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func init() {
	resource.AddTestSweepers("access_token_validators", &resource.Sweeper{
		Name: "access_token_validators",
		F: func(r string) error {
			svc := accessTokenValidators.New(conf)
			results, _, err := svc.GetAccessTokenValidatorsCommand(&accessTokenValidators.GetAccessTokenValidatorsCommandInput{Filter: "acctest_"})
			if err != nil {
				return fmt.Errorf("unable to list access token validators to sweep %s", err)
			}
			for _, item := range results.Items {
				_, err = svc.DeleteAccessTokenValidatorCommand(&accessTokenValidators.DeleteAccessTokenValidatorCommandInput{Id: item.Id.String()})
				if err != nil {
					return fmt.Errorf("unable to sweep access token validator %s because %s", item.Id.String(), err)
				}
			}
			return nil
		},
	})
}

func TestAccPingAccessAccessTokenValidator(t *testing.T) {
	resourceName := "pingaccess_access_token_validator.test"

	resource.ParallelTest(t, resource.TestCase{
		ProtoV5ProviderFactories: map[string]func() (tfprotov5.ProviderServer, error){
			"pingaccess": func() (tfprotov5.ProviderServer, error) {
				return Server(), nil
			},
		},
		CheckDestroy: testAccCheckPingAccessAccessTokenValidatorDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingAccessAccessTokenValidatorConfig("acctest_foo", `<<EOF
{
   "description": null,
   "path": "/bar",
   "subjectAttributeName": "demo",
   "issuer": null,
   "audience": null
}
EOF`),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessAccessTokenValidatorExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", "acctest_foo"),
					resource.TestCheckResourceAttr(resourceName, "class_name", "com.pingidentity.pa.accesstokenvalidators.JwksEndpoint"),
					resource.TestCheckResourceAttr(resourceName, "configuration", "{\n   \"description\": null,\n   \"path\": \"/bar\",\n   \"subjectAttributeName\": \"demo\",\n   \"issuer\": null,\n   \"audience\": null\n}\n"),
				),
			},
			{
				Config: testAccPingAccessAccessTokenValidatorConfig("acctest_foo", `<<EOF
{
   "description": null,
   "path": "/foo/bar",
   "subjectAttributeName": "demo",
   "issuer": null,
   "audience": null
}
EOF`),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessAccessTokenValidatorExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", "acctest_foo"),
					resource.TestCheckResourceAttr(resourceName, "class_name", "com.pingidentity.pa.accesstokenvalidators.JwksEndpoint"),
					resource.TestCheckResourceAttr(resourceName, "configuration", "{\n   \"description\": null,\n   \"path\": \"/foo/bar\",\n   \"subjectAttributeName\": \"demo\",\n   \"issuer\": null,\n   \"audience\": null\n}\n"),
				),
			},
			//TODO we dont support json style imports
			//{
			//	ResourceName:      resourceName,
			//	ImportState:       true,
			//	ImportStateVerify: true,
			//},
			{
				Config: testAccPingAccessAccessTokenValidatorConfigInvalidClassName(`<<EOF
		{
			"description": null,
			"path": "/foo",
			"subjectAttributeName": "foo",
			"issuer": null,
			"audience": null
		}
		EOF`),
				ExpectError: regexp.MustCompile(`unable to find className 'com.pingidentity.pa.accesstokenvalidators.foo'\navailable classNames: com.pingidentity.pa.accesstokenvalidators.JwksEndpoint`),
			},
		},
	})
}

func TestAccPingAccessAccessTokenValidatorWithDynamicPsuedoType(t *testing.T) {
	resourceName := "pingaccess_access_token_validator.test"

	resource.ParallelTest(t, resource.TestCase{
		ProtoV5ProviderFactories: map[string]func() (tfprotov5.ProviderServer, error){
			"pingaccess": func() (tfprotov5.ProviderServer, error) {
				return Server(), nil
			},
		},
		CheckDestroy: testAccCheckPingAccessAccessTokenValidatorDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingAccessAccessTokenValidatorConfig("acctest_bar", `{
			"path" = "/bar/foo"
			"subjectAttributeName" = "foo"
		}`),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessAccessTokenValidatorExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", "acctest_bar"),
					resource.TestCheckResourceAttr(resourceName, "class_name", "com.pingidentity.pa.accesstokenvalidators.JwksEndpoint"),
					resource.TestCheckResourceAttr(resourceName, "configuration.subjectAttributeName", "foo"),
					resource.TestCheckResourceAttr(resourceName, "configuration.path", "/bar/foo"),
				),
			},
			{
				Config: testAccPingAccessAccessTokenValidatorConfig("acctest_bar", `{
			"path" = "/foo/bar"
			"subjectAttributeName" = "foo"
		}`),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessAccessTokenValidatorExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", "acctest_bar"),
					resource.TestCheckResourceAttr(resourceName, "class_name", "com.pingidentity.pa.accesstokenvalidators.JwksEndpoint"),
					resource.TestCheckResourceAttr(resourceName, "configuration.subjectAttributeName", "foo"),
					resource.TestCheckResourceAttr(resourceName, "configuration.path", "/foo/bar"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccPingAccessAccessTokenValidatorConfig("acctest_bar", `{
			"subjectAttributeName" = "foo"
		}`),
				ExpectError: regexp.MustCompile(`the field 'path' is required for the class_name\n'com.pingidentity.pa.accesstokenvalidators.JwksEndpoint'`),
			},
			{
				Config: testAccPingAccessAccessTokenValidatorConfigInvalidClassName(`{
			"path" = "/foo"
			"subjectAttributeName" = "foo"
		}`),
				ExpectError: regexp.MustCompile(`unable to find className 'com.pingidentity.pa.accesstokenvalidators.foo'\navailable classNames: com.pingidentity.pa.accesstokenvalidators.JwksEndpoint`),
			},
			{
				Config: testAccPingAccessAccessTokenValidatorConfig("acctest_bar", `{
			"description"          = null
			"path" = "/bar/foo"
			"subjectAttributeName" = "foo"
		}`),
				ExpectError: regexp.MustCompile(`configuration fields cannot be null, remove 'description' or set a non-null\nvalue`),
			},
		},
	})
}

func testAccCheckPingAccessAccessTokenValidatorDestroy(s *terraform.State) error {
	return nil
}

func testAccPingAccessAccessTokenValidatorConfig(name, configUpdate string) string {
	return fmt.Sprintf(`
	resource "pingaccess_access_token_validator" "test" {
		class_name = "com.pingidentity.pa.accesstokenvalidators.JwksEndpoint"
		name = "%s"

		configuration = %s
	}
`, name, configUpdate)
}

func testAccPingAccessAccessTokenValidatorConfigInvalidClassName(configUpdate string) string {
	return fmt.Sprintf(`
	resource "pingaccess_access_token_validator" "test" {
		class_name = "com.pingidentity.pa.accesstokenvalidators.foo"
		name = "acctest_foo"
		configuration = %s
	}`, configUpdate)
}

func testAccCheckPingAccessAccessTokenValidatorExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" || rs.Primary.ID == "0" {
			return fmt.Errorf("no access_token_validator ID is set")
		}

		conn := accessTokenValidators.New(conf)
		result, _, err := conn.GetAccessTokenValidatorCommand(&accessTokenValidators.GetAccessTokenValidatorCommandInput{
			Id: rs.Primary.ID,
		})

		if err != nil {
			return fmt.Errorf("error: AccessTokenValidator (%s) not found", n)
		}

		if *result.Name != rs.Primary.Attributes["name"] {
			return fmt.Errorf("error: AccessTokenValidator response (%s) didnt match state (%s)", *result.Name, rs.Primary.Attributes["name"])
		}

		return nil
	}
}
