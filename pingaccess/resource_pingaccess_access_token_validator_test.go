package pingaccess

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/iwarapter/pingaccess-sdk-go/pingaccess"
)

func TestAccPingAccessAccessTokenValidator(t *testing.T) {
	resourceName := "pingaccess_access_token_validator.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPingAccessAccessTokenValidatorDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingAccessAccessTokenValidatorConfig("/bar"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessAccessTokenValidatorExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", "foo"),
					resource.TestCheckResourceAttr(resourceName, "class_name", "com.pingidentity.pa.accesstokenvalidators.JwksEndpoint"),
					resource.TestCheckResourceAttr(resourceName, "configuration", "{\"audience\":null,\"description\":null,\"issuer\":null,\"path\":\"/bar\",\"subjectAttributeName\":\"foo\"}"),
				),
			},
			{
				Config: testAccPingAccessAccessTokenValidatorConfig("/bar/foo"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessAccessTokenValidatorExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", "foo"),
					resource.TestCheckResourceAttr(resourceName, "class_name", "com.pingidentity.pa.accesstokenvalidators.JwksEndpoint"),
					resource.TestCheckResourceAttr(resourceName, "configuration", "{\"audience\":null,\"description\":null,\"issuer\":null,\"path\":\"/bar/foo\",\"subjectAttributeName\":\"foo\"}"),
				),
			},
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
		name = "foo"
		
		configuration = <<EOF
		{
			"description": null,
			"path": "%s",
			"subjectAttributeName": "foo",
			"issuer": null,
			"audience": null
		}
		EOF
	}`, configUpdate)
}

func testAccCheckPingAccessAccessTokenValidatorExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" || rs.Primary.ID == "0" {
			return fmt.Errorf("No access_token_validator ID is set")
		}

		conn := testAccProvider.Meta().(*pingaccess.Client).AccessTokenValidators
		result, _, err := conn.GetAccessTokenValidatorCommand(&pingaccess.GetAccessTokenValidatorCommandInput{
			Id: rs.Primary.ID,
		})

		if err != nil {
			return fmt.Errorf("Error: AccessTokenValidator (%s) not found", n)
		}

		if *result.Name != rs.Primary.Attributes["name"] {
			return fmt.Errorf("Error: AccessTokenValidator response (%s) didnt match state (%s)", *result.Name, rs.Primary.Attributes["name"])
		}

		return nil
	}
}
