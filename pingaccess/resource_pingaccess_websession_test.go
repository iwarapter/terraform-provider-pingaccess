package pingaccess

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	pa "github.com/iwarapter/pingaccess-sdk-go/pingaccess"
)

func TestAccPingAccessWebSession(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPingAccessWebSessionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingAccessWebSessionConfig("woot", "password"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessWebSessionExists("pingaccess_websession.demo_session"),
				),
			},
			{
				Config: testAccPingAccessWebSessionConfig("woot", "changeme"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessWebSessionExists("pingaccess_websession.demo_session"),
				),
			},
		},
	})
}

func testAccCheckPingAccessWebSessionDestroy(s *terraform.State) error {
	return nil
}

func testAccPingAccessWebSessionConfig(audience, password string) string {
	return fmt.Sprintf(`
	resource "pingaccess_websession" "demo_session" {
		name = "acc-test-demo-session"
		audience = "%s"
		client_credentials {
			client_id = "websession"
			client_secret {
				value = "%s"
			}
		}
		scopes = [
			"profile",
			"address",
			"email",
			"phone"
		]
	}`, audience, password)
}

func testAccCheckPingAccessWebSessionExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" || rs.Primary.ID == "0" {
			return fmt.Errorf("No websession ID is set")
		}

		conn := testAccProvider.Meta().(*pa.Client).WebSessions
		result, _, err := conn.GetWebSessionCommand(&pa.GetWebSessionCommandInput{
			Id: rs.Primary.ID,
		})

		if err != nil {
			return fmt.Errorf("Error: WebSession (%s) not found", n)
		}

		if *result.Name != rs.Primary.Attributes["name"] {
			return fmt.Errorf("Error: WebSession response (%s) didnt match state (%s)", *result.Name, rs.Primary.Attributes["name"])
		}

		return nil
	}
}

func Test_resourcePingAccessWebSessionReadData(t *testing.T) {
	cases := []struct {
		WebSession pa.WebSessionView
	}{
		{
			WebSession: pa.WebSessionView{
				Audience: String("localhost"),
				Name:     String("localhost"),
				ClientCredentials: &pa.OAuthClientCredentialsView{
					ClientId:     String("client"),
					ClientSecret: &pa.HiddenFieldView{},
				},
				WebStorageType:                String("SessionStorage"),
				CacheUserAttributes:           Bool(true),
				CookieDomain:                  String("true"),
				CookieType:                    String("Signed"),
				EnableRefreshUser:             Bool(true),
				HttpOnlyCookie:                Bool(true),
				IdleTimeoutInMinutes:          Int(0),
				OidcLoginType:                 String("Code"),
				PkceChallengeType:             String("OFF"),
				PfsessionStateCacheInSeconds:  Int(0),
				RefreshUserInfoClaimsInterval: Int(0),
				RequestPreservationType:       String("true"),
				RequestProfile:                Bool(true),
				SameSite:                      String("None"),
				Scopes:                        &[]*string{String("true")},
				SecureCookie:                  Bool(true),
				SendRequestedUrlToProvider:    Bool(true),
				SessionTimeoutInMinutes:       Int(0),
				ValidateSessionIsAlive:        Bool(true),
			},
		},
		{
			WebSession: pa.WebSessionView{
				Audience: String("localhost"),
				Name:     String("localhost"),
				ClientCredentials: &pa.OAuthClientCredentialsView{
					ClientId:     String("client"),
					ClientSecret: &pa.HiddenFieldView{},
				},
				WebStorageType:                String("SessionStorage"),
				CacheUserAttributes:           Bool(false),
				CookieType:                    String("Encrypted"),
				EnableRefreshUser:             Bool(false),
				HttpOnlyCookie:                Bool(false),
				IdleTimeoutInMinutes:          Int(1),
				OidcLoginType:                 String("POST"),
				PkceChallengeType:             String("SHA256"),
				PfsessionStateCacheInSeconds:  Int(1),
				RefreshUserInfoClaimsInterval: Int(1),
				RequestPreservationType:       String("All"),
				RequestProfile:                Bool(false),
				SameSite:                      String("Disabled"),
				Scopes:                        &[]*string{},
				SecureCookie:                  Bool(false),
				SendRequestedUrlToProvider:    Bool(false),
				SessionTimeoutInMinutes:       Int(1),
				ValidateSessionIsAlive:        Bool(false),
			},
		},
	}
	for i, tc := range cases {
		t.Run(fmt.Sprintf("tc:%v", i), func(t *testing.T) {

			resourceSchema := resourcePingAccessWebSessionSchema()
			resourceLocalData := schema.TestResourceDataRaw(t, resourceSchema, map[string]interface{}{})

			resourcePingAccessWebSessionReadResult(resourceLocalData, &tc.WebSession)
			if got := *resourcePingAccessWebSessionReadData(resourceLocalData); !cmp.Equal(got, tc.WebSession) {
				t.Errorf("resourcePingAccessWebSessionReadData() = %v", cmp.Diff(got, tc.WebSession))
			}
		})
	}
}
