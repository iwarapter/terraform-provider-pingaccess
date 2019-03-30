package pingaccess

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	pa "github.com/iwarapter/pingaccess-sdk-go/pingaccess"
)

func TestAccPingAccessOAuthServer(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPingAccessOAuthServerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingAccessOAuthServerConfig("/introspect", "top"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessOAuthServerExists("pingaccess_oauth_server.demo_pfr"),
				),
			},
			{
				Config: testAccPingAccessOAuthServerConfig("/introspect", "secret"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessOAuthServerExists("pingaccess_oauth_server.demo_pfr"),
				),
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

		conn := testAccProvider.Meta().(*pa.Client).OAuth
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

func Test_resourcePingAccessOAuthServerReadData(t *testing.T) {
	cases := []struct {
		OAuthServer pa.AuthorizationServerView
	}{
		{
			OAuthServer: pa.AuthorizationServerView{
				IntrospectionEndpoint:     String("/introspection"),
				SubjectAttributeName:      String("alt"),
				Targets:                   &[]*string{String("localhost")},
				TrustedCertificateGroupId: Int(0),
				ClientCredentials: &pa.OAuthClientCredentialsView{
					ClientId: String("client"),
					ClientSecret: &pa.HiddenFieldView{
						Value: String("Secrets"),
					},
				},
				AuditLevel:             String("ON"),
				TokenTimeToLiveSeconds: Int(-1),
			},
		},
		{
			OAuthServer: pa.AuthorizationServerView{
				IntrospectionEndpoint:     String("/introspection"),
				SubjectAttributeName:      String("alt"),
				Targets:                   &[]*string{String("localhost")},
				TrustedCertificateGroupId: Int(0),
				ClientCredentials: &pa.OAuthClientCredentialsView{
					ClientId:     String("client"),
					ClientSecret: &pa.HiddenFieldView{},
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
			resourcePingAccessOAuthServerReadResult(resourceLocalData, &tc.OAuthServer)

			if got := *resourcePingAccessOAuthServerReadData(resourceLocalData); !cmp.Equal(got, tc.OAuthServer) {
				t.Errorf("resourcePingAccessOAuthServerReadData() = %v", cmp.Diff(got, tc.OAuthServer))
			}

			resourcePingAccessOAuthServerReadResult(resourceLocalData, &tc.OAuthServer)
		})
	}
}
