package sdkv2provider

import (
	"fmt"
	"testing"

	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess/models"
	"github.com/iwarapter/pingaccess-sdk-go/v62/services/webSessions"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func init() {
	resource.AddTestSweepers("websessions", &resource.Sweeper{
		Name: "websessions",
		F: func(r string) error {
			svc := webSessions.New(conf)
			results, _, err := svc.GetWebSessionsCommand(&webSessions.GetWebSessionsCommandInput{Filter: "acctest_"})
			if err != nil {
				return fmt.Errorf("unable to list websessions to sweep %s", err)
			}
			for _, item := range results.Items {
				_, err = svc.DeleteWebSessionCommand(&webSessions.DeleteWebSessionCommandInput{Id: item.Id.String()})
				if err != nil {
					return fmt.Errorf("unable to sweep websession %s because %s", item.Id.String(), err)
				}
			}
			return nil
		},
	})
}

func TestAccPingAccessWebSession(t *testing.T) {
	resourceName := "pingaccess_websession.demo_session"

	canMask := !(paClient{apiVersion: paVersion}).Is61OrAbove()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV5ProviderFactories: testAccProviders,
		CheckDestroy:             testAccCheckPingAccessWebSessionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingAccessWebSessionConfig("woot", "password"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessWebSessionExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "audience", "woot"),
					resource.TestCheckResourceAttr(resourceName, "name", "acctest_demo-session"),
					resource.TestCheckResourceAttr(resourceName, "client_credentials.0.client_secret.0.value", "password"),
					resource.TestCheckResourceAttrSet(resourceName, "client_credentials.0.client_secret.0.encrypted_value"),
				),
			},
			{
				Config: testAccPingAccessWebSessionConfig("woot", "changeme"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessWebSessionExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "audience", "woot"),
					resource.TestCheckResourceAttr(resourceName, "name", "acctest_demo-session"),
					resource.TestCheckResourceAttr(resourceName, "client_credentials.0.client_secret.0.value", "changeme"),
					resource.TestCheckResourceAttrSet(resourceName, "client_credentials.0.client_secret.0.encrypted_value"),
				),
			},
			{ // we change the password directly and check the provider detects it
				Config: testAccPingAccessWebSessionConfig("woot", "changeme"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessWebSessionExists(resourceName),
					testAccCheckPingAccessWebSessionCanTrackPasswordChanges(resourceName, canMask),
					resource.TestCheckResourceAttr(resourceName, "audience", "woot"),
					resource.TestCheckResourceAttr(resourceName, "name", "acctest_demo-session"),
					resource.TestCheckResourceAttr(resourceName, "client_credentials.0.client_secret.0.value", "changeme"),
					resource.TestCheckResourceAttrSet(resourceName, "client_credentials.0.client_secret.0.encrypted_value"),
				),
				ExpectNonEmptyPlan: canMask,
			},
			{
				Config: testAccPingAccessWebSessionConfig("woot", "changeme"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessWebSessionExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "client_credentials.0.client_secret.0.value", "changeme"),
					resource.TestCheckResourceAttrSet(resourceName, "client_credentials.0.client_secret.0.encrypted_value"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"client_credentials.0.client_secret.0.value", "client_credentials.0.client_secret.0.encrypted_value"}, //we cant verify passwords
			},
		},
	})
}

func TestAccPingAccessWebSessionPrivateKeyJwt(t *testing.T) {
	resourceName := "pingaccess_websession.demo_session"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV5ProviderFactories: testAccProviders,
		CheckDestroy:             testAccCheckPingAccessWebSessionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingAccessWebSessionPrivateKeyJwtConfig("woot"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessWebSessionExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "audience", "woot"),
					resource.TestCheckResourceAttr(resourceName, "name", "acctest_demo-session-two"),
					resource.TestCheckResourceAttr(resourceName, "client_credentials.0.client_id", "websession"),
					resource.TestCheckResourceAttr(resourceName, "client_credentials.0.credentials_type", "PRIVATE_KEY_JWT"),
				),
			},
			{
				Config: testAccPingAccessWebSessionPrivateKeyJwtConfig("bar"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessWebSessionExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "audience", "bar"),
					resource.TestCheckResourceAttr(resourceName, "name", "acctest_demo-session-two"),
					resource.TestCheckResourceAttr(resourceName, "client_credentials.0.client_id", "websession"),
					resource.TestCheckResourceAttr(resourceName, "client_credentials.0.credentials_type", "PRIVATE_KEY_JWT"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
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
		name = "acctest_demo-session"
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

func testAccPingAccessWebSessionPrivateKeyJwtConfig(audience string) string {
	return fmt.Sprintf(`
	resource "pingaccess_websession" "demo_session" {
		name = "acctest_demo-session-two"
		audience = "%s"
		client_credentials {
			client_id = "websession"
			credentials_type = "PRIVATE_KEY_JWT"
		}
		scopes = ["profile","address","email","phone"]
	}`, audience)
}

func testAccCheckPingAccessWebSessionExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not found: %s", n)
		}

		if rs.Primary.ID == "" || rs.Primary.ID == "0" {
			return fmt.Errorf("no websession ID is set")
		}

		conn := testAccProvider.Meta().(paClient).WebSessions
		result, _, err := conn.GetWebSessionCommand(&webSessions.GetWebSessionCommandInput{
			Id: rs.Primary.ID,
		})

		if err != nil {
			return fmt.Errorf("error: WebSession (%s) not found", n)
		}

		if *result.Name != rs.Primary.Attributes["name"] {
			return fmt.Errorf("error: WebSession response (%s) didnt match state (%s)", *result.Name, rs.Primary.Attributes["name"])
		}

		return nil
	}
}

func testAccCheckPingAccessWebSessionCanTrackPasswordChanges(n string, changePassword bool) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not found: %s", n)
		}

		if rs.Primary.ID == "" || rs.Primary.ID == "0" {
			return fmt.Errorf("no websession ID is set")
		}

		conn := testAccProvider.Meta().(paClient).WebSessions
		result, _, err := conn.GetWebSessionCommand(&webSessions.GetWebSessionCommandInput{
			Id: rs.Primary.ID,
		})

		if err != nil {
			return fmt.Errorf("unable to retrieve webSession %s", err)
		}
		if changePassword {
			result.ClientCredentials.ClientSecret.Value = String("i have been changed")
			updated, _, err := conn.UpdateWebSessionCommand(&webSessions.UpdateWebSessionCommandInput{
				Body: *result,
				Id:   rs.Primary.ID,
			})

			if err != nil {
				return fmt.Errorf("failed to update webSession %s", err)
			}

			if *updated.ClientCredentials.ClientSecret.EncryptedValue == *result.ClientCredentials.ClientSecret.EncryptedValue {
				return fmt.Errorf("the encryptedValue for webSession should of changed after an update")
			}
		}

		return nil
	}
}

func Test_resourcePingAccessWebSessionReadData(t *testing.T) {
	cases := []struct {
		WebSession models.WebSessionView
	}{
		{
			WebSession: models.WebSessionView{
				Audience: String("localhost"),
				Name:     String("localhost"),
				ClientCredentials: &models.OAuthClientCredentialsView{
					ClientId:        String("client"),
					ClientSecret:    &models.HiddenFieldView{},
					KeyPairId:       Int(0),
					CredentialsType: String("SECRET"),
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
			WebSession: models.WebSessionView{
				Audience: String("localhost"),
				Name:     String("localhost"),
				ClientCredentials: &models.OAuthClientCredentialsView{
					ClientId:        String("client"),
					ClientSecret:    &models.HiddenFieldView{},
					KeyPairId:       Int(1),
					CredentialsType: String("PRIVATE_KEY_JWT"),
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

			resourcePingAccessWebSessionReadResult(resourceLocalData, &tc.WebSession, false)
			if got := *resourcePingAccessWebSessionReadData(resourceLocalData); !cmp.Equal(got, tc.WebSession) {
				t.Errorf("resourcePingAccessWebSessionReadData() = %v", cmp.Diff(got, tc.WebSession))
			}
		})
	}
}
