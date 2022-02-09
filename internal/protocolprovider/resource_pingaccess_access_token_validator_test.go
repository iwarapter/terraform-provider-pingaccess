package protocol

import (
	"context"
	"fmt"
	"regexp"
	"testing"

	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/hashicorp/terraform-plugin-go/tfprotov5"

	"github.com/iwarapter/pingaccess-sdk-go/v62/services/accessTokenValidators"

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

	resource.Test(t, resource.TestCase{
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

	resource.Test(t, resource.TestCase{
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

func Test_resourcePingAccessAccessTokenValidator_UpgradeResourceState(t *testing.T) {
	r := resourcePingAccessAccessTokenValidator{
		client:                accessTokenValidators.New(&config.Config{Endpoint: String("foo")}),
		genericPluginResource: genericPluginResource{},
	}
	tests := []struct {
		name string
		ctx  context.Context
		req  *tfprotov5.UpgradeResourceStateRequest
		resp *tfprotov5.UpgradeResourceStateResponse
	}{
		{
			"we can upgrade state",
			context.Background(),
			&tfprotov5.UpgradeResourceStateRequest{
				TypeName: "pingaccess_access_token_validator",
				Version:  0,
				RawState: &tfprotov5.RawState{
					JSON: []byte(`{
            "class_name": "com.pingidentity.pa.accesstokenvalidators.JwksEndpoint",
            "configuration": "{\"audience\":null,\"description\":null,\"issuer\":null,\"path\":\"/bar\",\"subjectAttributeName\":\"demo\"}",
            "id": "3",
            "name": "demo"
          }`),
				},
			},
			&tfprotov5.UpgradeResourceStateResponse{
				UpgradedState: &tfprotov5.DynamicValue{MsgPack: []byte{132, 170, 99, 108, 97, 115, 115, 95, 110, 97, 109, 101, 217, 54, 99, 111, 109, 46, 112, 105, 110, 103, 105, 100, 101, 110, 116, 105, 116, 121, 46, 112, 97, 46, 97, 99, 99, 101, 115, 115, 116, 111, 107, 101, 110, 118, 97, 108, 105, 100, 97, 116, 111, 114, 115, 46, 74, 119, 107, 115, 69, 110, 100, 112, 111, 105, 110, 116, 173, 99, 111, 110, 102, 105, 103, 117, 114, 97, 116, 105, 111, 110, 146, 196, 8, 34, 115, 116, 114, 105, 110, 103, 34, 217, 94, 123, 34, 97, 117, 100, 105, 101, 110, 99, 101, 34, 58, 110, 117, 108, 108, 44, 34, 100, 101, 115, 99, 114, 105, 112, 116, 105, 111, 110, 34, 58, 110, 117, 108, 108, 44, 34, 105, 115, 115, 117, 101, 114, 34, 58, 110, 117, 108, 108, 44, 34, 112, 97, 116, 104, 34, 58, 34, 47, 98, 97, 114, 34, 44, 34, 115, 117, 98, 106, 101, 99, 116, 65, 116, 116, 114, 105, 98, 117, 116, 101, 78, 97, 109, 101, 34, 58, 34, 100, 101, 109, 111, 34, 125, 162, 105, 100, 161, 51, 164, 110, 97, 109, 101, 164, 100, 101, 109, 111}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := r.UpgradeResourceState(tt.ctx, tt.req)
			require.NoError(t, err)
			assert.Equal(t, tt.resp, resp)
		})
	}
}
