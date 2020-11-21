package pingaccess

import (
	"fmt"
	"github.com/iwarapter/pingaccess-sdk-go/services/oauth"
	"regexp"
	"testing"

	"github.com/iwarapter/pingaccess-sdk-go/pingaccess/models"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccPingAccessOAuthServer(t *testing.T) {
	resourceName := "pingaccess_oauth_server.demo_pfr"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		ProtoV5ProviderFactories: testAccProviders,
		CheckDestroy: testAccCheckPingAccessOAuthServerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingAccessOAuthServerConfig("/introspect", "top"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessOAuthServerExists(resourceName),
				),
			},
			{
				Config: testAccPingAccessOAuthServerConfig("/introspect", "secret"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessOAuthServerExists(resourceName),
				),
			},
			{ // we change the password directly and check the provider detects it
				Config: testAccPingAccessOAuthServerConfig("/introspect", "secret"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessOAuthServerExists(resourceName),
					testAccCheckPingAccessOAuthServerCanTrackPasswordChanges(resourceName, true),
					resource.TestCheckResourceAttr(resourceName, "introspection_endpoint", "/introspect"),
					resource.TestCheckResourceAttr(resourceName, "client_credentials.0.client_secret.0.value", "secret"),
					resource.TestCheckResourceAttrSet(resourceName, "client_credentials.0.client_secret.0.encrypted_value"),

				),
				ExpectNonEmptyPlan: true,
			},
			{
				Config: testAccPingAccessOAuthServerConfig("/introspect", "secret"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessOAuthServerExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "introspection_endpoint", "/introspect"),
					resource.TestCheckResourceAttr(resourceName, "client_credentials.0.client_secret.0.value", "secret"),
					resource.TestCheckResourceAttrSet(resourceName, "client_credentials.0.client_secret.0.encrypted_value"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{"client_credentials.0.client_secret.0.value"},
			},
			{
				Config: testAccPingAccessOAuthServerConfig("https://thing/introspect", "secret"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessOAuthServerExists(resourceName),
				),
				ExpectError: regexp.MustCompile(`unable to update OAuthServerSettings: Save Failed:\nintrospectionEndpoint contains 1 validation failures:\n\tIntrospection endpoint must be a valid relative path`),
			},
		},
	})
}

func testAccCheckPingAccessOAuthServerDestroy(s *terraform.State) error {
	return nil
}

func testAccPingAccessOAuthServerConfig(introspectionURL, clientPassword string) string {
	return fmt.Sprintf(`
	resource "pingaccess_oauth_server" "demo_pfr" {
		targets = ["localhost:9031"]
		subject_attribute_name = "san"
		trusted_certificate_group_id = 1
		introspection_endpoint = "%s"
		client_credentials {
			client_id = "oauth"
			client_secret {
				value = "%s"
			}
		}
		secure = true
	}`, introspectionURL, clientPassword)
}

func testAccCheckPingAccessOAuthServerExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" || rs.Primary.ID == "0" {
			return fmt.Errorf("No third party service ID is set")
		}

		conn := testAccProvider.Meta().(paClient).Oauth
		result, _, err := conn.GetAuthorizationServerCommand()

		if err != nil {
			return fmt.Errorf("Error: OAuthServer (%s) not found", n)
		}

		if *result.IntrospectionEndpoint != rs.Primary.Attributes["introspection_endpoint"] {
			return fmt.Errorf("Error: OAuthServer response (%s) didnt match state (%s)", *result.IntrospectionEndpoint, rs.Primary.Attributes["introspection_endpoint"])
		}

		return nil
	}
}

func testAccCheckPingAccessOAuthServerCanTrackPasswordChanges(n string, changePassword bool) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not found: %s", n)
		}

		if rs.Primary.ID == "" || rs.Primary.ID == "0" {
			return fmt.Errorf("no OAuth Server ID is set")
		}

		conn := testAccProvider.Meta().(paClient).Oauth
		result, _, err := conn.GetAuthorizationServerCommand()

		if err != nil {
			return fmt.Errorf("unable to retrieve OAuth Server %s", err)
		}
		if changePassword {
			result.ClientCredentials.ClientSecret.Value = String("i have been changed")
			updated, _, err := conn.UpdateAuthorizationServerCommand(&oauth.UpdateAuthorizationServerCommandInput{
				Body: *result,
			})

			if err != nil {
				return fmt.Errorf("failed to update OAuth Server %s", err)
			}

			if *updated.ClientCredentials.ClientSecret.EncryptedValue == *result.ClientCredentials.ClientSecret.EncryptedValue {
				return fmt.Errorf("the encryptedValue for OAuth Server should of changed after an update")
			}
		}
		return nil
	}
}

func Test_resourcePingAccessOAuthServerReadData(t *testing.T) {
	cases := []struct {
		OAuthServer models.AuthorizationServerView
	}{
		{
			OAuthServer: models.AuthorizationServerView{
				IntrospectionEndpoint:     String("/introspection"),
				SubjectAttributeName:      String("alt"),
				Targets:                   &[]*string{String("localhost")},
				TrustedCertificateGroupId: Int(0),
				ClientCredentials: &models.OAuthClientCredentialsView{
					ClientId: String("client"),
					ClientSecret: &models.HiddenFieldView{
						Value:          String("Secrets"),
						EncryptedValue: String("foo"),
					},
				},
				AuditLevel:             String("ON"),
				TokenTimeToLiveSeconds: Int(-1),
			},
		},
		{
			OAuthServer: models.AuthorizationServerView{
				IntrospectionEndpoint:     String("/introspection"),
				SubjectAttributeName:      String("alt"),
				Targets:                   &[]*string{String("localhost")},
				TrustedCertificateGroupId: Int(0),
				ClientCredentials: &models.OAuthClientCredentialsView{
					ClientId:     String("client"),
					ClientSecret: &models.HiddenFieldView{},
				},
				AuditLevel:             String("none"),
				Secure:                 Bool(true),
				UseProxy:               Bool(true),
				CacheTokens:            Bool(true),
				Description:            String("foo"),
				SendAudience:           Bool(true),
				TokenTimeToLiveSeconds: Int(60),
			},
		},
	}
	for i, tc := range cases {
		t.Run(fmt.Sprintf("tc:%v", i), func(t *testing.T) {

			resourceSchema := resourcePingAccessOAuthServerSchema()
			resourceLocalData := schema.TestResourceDataRaw(t, resourceSchema, map[string]interface{}{})
			resourcePingAccessOAuthServerReadResult(resourceLocalData, &tc.OAuthServer, false)

			if got := *resourcePingAccessOAuthServerReadData(resourceLocalData); !cmp.Equal(got, tc.OAuthServer) {
				t.Errorf("resourcePingAccessOAuthServerReadData() = %v", cmp.Diff(got, tc.OAuthServer))
			}

			resourcePingAccessOAuthServerReadResult(resourceLocalData, &tc.OAuthServer, false)
		})
	}
}
