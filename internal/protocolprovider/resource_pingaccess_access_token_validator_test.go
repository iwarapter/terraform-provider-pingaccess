package protocol

import (
	"fmt"
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
				Config: testAccPingAccessAccessTokenValidatorConfig("/bar"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessAccessTokenValidatorExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", "acctest_foo"),
					resource.TestCheckResourceAttr(resourceName, "class_name", "com.pingidentity.pa.accesstokenvalidators.JwksEndpoint"),
					//resource.TestCheckResourceAttr(resourceName, "configuration", "{\"audience\":null,\"description\":null,\"issuer\":null,\"path\":\"/bar\",\"subjectAttributeName\":\"foo\"}"),
				),
			},
			{
				Config: testAccPingAccessAccessTokenValidatorConfig("/bar/foo"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessAccessTokenValidatorExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", "acctest_foo"),
					resource.TestCheckResourceAttr(resourceName, "class_name", "com.pingidentity.pa.accesstokenvalidators.JwksEndpoint"),
					resource.TestCheckResourceAttr(resourceName, "configuration", "\t\t{\n\t\t\t\"description\": null,\n\t\t\t\"path\": \"/bar/foo\",\n\t\t\t\"subjectAttributeName\": \"foo\",\n\t\t\t\"issuer\": null,\n\t\t\t\"audience\": null\n\t\t}\n"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			//{
			//	Config:      testAccPingAccessAccessTokenValidatorConfigInvalidClassName(),
			//	ExpectError: regexp.MustCompile(`unable to find className 'com.pingidentity.pa.accesstokenvalidators.foo' available classNames: com.pingidentity.pa.accesstokenvalidators.JwksEndpoint`),
			//},
		},
	})
}

func testAccCheckPingAccessAccessTokenValidatorDestroy(s *terraform.State) error {
	return nil
}

func testAccPingAccessAccessTokenValidatorConfig(configUpdate string) string {
	return fmt.Sprintf(`
	resource "pingaccess_access_token_validator" "test" {
		class_name = "com.pingidentity.pa.accesstokenvalidators.JwksEndpoint"
		name = "acctest_foo"

		configuration = <<EOF
		{
			"description": null,
			"path": "%s",
			"subjectAttributeName": "foo",
			"issuer": null,
			"audience": null
		}
		EOF
	}
`, configUpdate)
}

//func testAccPingAccessAccessTokenValidatorConfigInvalidClassName() string {
//	return `
//	resource "pingaccess_access_token_validator" "test" {
//		class_name = "com.pingidentity.pa.accesstokenvalidators.foo"
//		name = "acctest_foo"
//
//		configuration = <<EOF
//		{
//			"description": null,
//			"path": "/foo",
//			"subjectAttributeName": "foo",
//			"issuer": null,
//			"audience": null
//		}
//		EOF
//	}`
//}

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
