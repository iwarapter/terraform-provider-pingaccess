package pingaccess

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	pa "github.com/iwarapter/pingaccess-sdk-go/pingaccess"
)

func init() {
	resource.AddTestSweepers("pingaccess_websession", &resource.Sweeper{
		Name: "pingaccess_websession",
		F:    testSweepWebSession,
	})
}

func testSweepWebSession(r string) error {
	url, _ := url.Parse("https://localhost:9000")
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	conn := pa.NewClient("Administrator", "2Access2", url, "/pa-admin-api/v3", nil).WebSessions
	result, _, _ := conn.GetWebSessionsCommand(&pa.GetWebSessionsCommandInput{Filter: "acc-test-"})
	for _, v := range result.Items {
		conn.DeleteWebSessionCommand(&pa.DeleteWebSessionCommandInput{Id: v.Id.String()})
	}
	return nil
}

func TestAccPingAccessWebSession(t *testing.T) {
	var out pa.WebSessionView

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPingAccessWebSessionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingAccessWebSessionConfig("woot"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessWebSessionExists("pingaccess_websession.demo_session", 3, &out),
				),
			},
			{
				Config: testAccPingAccessWebSessionConfig("foo"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessWebSessionExists("pingaccess_websession.demo_session", 6, &out),
				),
			},
		},
	})
}

func testAccCheckPingAccessWebSessionDestroy(s *terraform.State) error {
	return nil
}

func testAccPingAccessWebSessionConfig(audience string) string {
	return fmt.Sprintf(`
	resource "pingaccess_websession" "demo_session" {
		name = "acc-test-demo-session"
		audience = "%s"
		client_credentials {
			client_id = "websession",
			client_secret {
				value = "password"
			}
		}
		scopes = [
			"profile",
			"address",
			"email",
			"phone"
		]
	}`, name)
}

func testAccCheckPingAccessWebSessionExists(n string, c int64, out *pa.WebSessionView) resource.TestCheckFunc {
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
				CookieType:                    String("true"),
				EnableRefreshUser:             Bool(true),
				HttpOnlyCookie:                Bool(true),
				IdleTimeoutInMinutes:          Int(0),
				OidcLoginType:                 String("true"),
				PfsessionStateCacheInSeconds:  Int(0),
				RefreshUserInfoClaimsInterval: Int(0),
				RequestPreservationType:       String("true"),
				RequestProfile:                Bool(true),
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
				CookieDomain:                  String(""),
				CookieType:                    String(""),
				EnableRefreshUser:             Bool(false),
				HttpOnlyCookie:                Bool(false),
				IdleTimeoutInMinutes:          Int(1),
				OidcLoginType:                 String(""),
				PfsessionStateCacheInSeconds:  Int(1),
				RefreshUserInfoClaimsInterval: Int(1),
				RequestPreservationType:       String(""),
				RequestProfile:                Bool(false),
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
