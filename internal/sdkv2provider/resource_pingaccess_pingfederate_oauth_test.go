package sdkv2provider

import (
	"fmt"
	"testing"

	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestAccPingAccessPingFederateOAuth61OrBelow(t *testing.T) {
	if (paClient{apiVersion: paVersion}).Is61OrAbove() {
		t.Skipf("This test only runs against PingAccess 5.3 or 6.0, not: %s", paVersion)
	}
	resourceName := "pingaccess_pingfederate_oauth.demo_pfo"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV5ProviderFactories: testAccProviders,
		CheckDestroy:             testAccCheckPingAccessPingFederateOAuthDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingAccessPingFederateOAuthConfig60("my_client", "san"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessPingFederateOAuthExists(resourceName),
				),
			},
			{
				Config: testAccPingAccessPingFederateOAuthConfig60("my_client", "sany"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessPingFederateOAuthExists(resourceName),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"client_secret.0.value", "client_secret.0.encrypted_value"},
			},
		},
	})
}

func TestAccPingAccessPingFederateOAuth61OrAbove(t *testing.T) {
	if !(paClient{apiVersion: paVersion}).Is61OrAbove() {
		t.Skipf("This test only runs against PingAccess 6.1 or above, not: %s", paVersion)
	}
	resourceName := "pingaccess_pingfederate_oauth.demo_pfo"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV5ProviderFactories: testAccProviders,
		CheckDestroy:             testAccCheckPingAccessPingFederateOAuthDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingAccessPingFederateOAuthConfig61("my_client", "san"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessPingFederateOAuthExists(resourceName),
				),
			},
			{
				Config: testAccPingAccessPingFederateOAuthConfig61("my_client", "sany"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessPingFederateOAuthExists(resourceName),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"client_credentials.0.client_secret.0.value"},
			},
		},
	})
}

func testAccCheckPingAccessPingFederateOAuthDestroy(s *terraform.State) error {
	return nil
}

func testAccPingAccessPingFederateOAuthConfig60(client, san string) string {
	return fmt.Sprintf(`
resource "pingaccess_pingfederate_oauth" "demo_pfo" {
  access_validator_id    = 1
  cache_tokens           = true
  subject_attribute_name = "%s"
  name                   = "PingFederate"
  client_id              = "%s"
  client_secret {
    value = "top_secret"
  }
  send_audience              = true
  token_time_to_live_seconds = 300
  use_token_introspection    = true
}

`, san, client)
}

func testAccPingAccessPingFederateOAuthConfig61(client, san string) string {
	return fmt.Sprintf(`
resource "pingaccess_pingfederate_oauth" "demo_pfo" {
  access_validator_id    = 1
  cache_tokens           = true
  subject_attribute_name = "%s"
  name                   = "PingFederate"
  client_credentials {
    client_id              = "%s"
    client_secret {
      value = "top_secret"
    }
  }
  send_audience              = true
  token_time_to_live_seconds = 300
  use_token_introspection    = true
}

`, san, client)
}

func testAccCheckPingAccessPingFederateOAuthExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" || rs.Primary.ID == "0" {
			return fmt.Errorf("No third party service ID is set")
		}

		conn := testAccProvider.Meta().(paClient).Pingfederate
		result, _, err := conn.GetPingFederateAccessTokensCommand()

		if err != nil {
			return fmt.Errorf("Error: PingFederateOAuth (%s) not found", n)
		}

		if *result.Name != rs.Primary.Attributes["name"] {
			return fmt.Errorf("Error: PingFederateOAuth response (%s) didnt match state (%s)", *result.Name, rs.Primary.Attributes["name"])
		}

		return nil
	}
}

func Test_resourcePingAccessPingFederateOAuthReadData(t *testing.T) {
	cases := []struct {
		PingFederateAccessTokenView models.PingFederateAccessTokenView
	}{
		{
			PingFederateAccessTokenView: models.PingFederateAccessTokenView{
				ClientId:               String("client_1"),
				SubjectAttributeName:   String("san"),
				AccessValidatorId:      Int(1),
				Name:                   String("PingFederate"),
				TokenTimeToLiveSeconds: Int(-1),
			},
		},
		{
			PingFederateAccessTokenView: models.PingFederateAccessTokenView{
				ClientId:               String("client_1"),
				SubjectAttributeName:   String("san"),
				AccessValidatorId:      Int(1),
				CacheTokens:            Bool(true),
				Name:                   String("PingFederate"),
				SendAudience:           Bool(true),
				TokenTimeToLiveSeconds: Int(30),
				UseTokenIntrospection:  Bool(true),
				ClientSecret: &models.HiddenFieldView{
					Value:          String("password"),
					EncryptedValue: String("foo"),
				},
			},
		},
		{
			PingFederateAccessTokenView: models.PingFederateAccessTokenView{
				SubjectAttributeName:   String("san"),
				AccessValidatorId:      Int(1),
				CacheTokens:            Bool(true),
				Name:                   String("PingFederate"),
				SendAudience:           Bool(true),
				TokenTimeToLiveSeconds: Int(30),
				UseTokenIntrospection:  Bool(true),
				ClientCredentials: &models.OAuthClientCredentialsView{
					ClientId:        String("example"),
					ClientSecret:    &models.HiddenFieldView{Value: String("secret"), EncryptedValue: String("")},
					CredentialsType: String("PRIVATE_KEY_JWT"),
					KeyPairId:       Int(1),
				},
			},
		},
	}
	for i, tc := range cases {
		t.Run(fmt.Sprintf("tc:%v", i), func(t *testing.T) {

			resourceSchema := resourcePingAccessPingFederateOAuthSchema()
			resourceLocalData := schema.TestResourceDataRaw(t, resourceSchema, map[string]interface{}{})
			resourcePingAccessPingFederateOAuthReadResult(resourceLocalData, &tc.PingFederateAccessTokenView, false)

			if got := *resourcePingAccessPingFederateOAuthReadData(resourceLocalData); !cmp.Equal(got, tc.PingFederateAccessTokenView) {
				t.Errorf("resourcePingAccessPingFederateOAuthReadData() = %v", cmp.Diff(got, tc.PingFederateAccessTokenView))
			}
		})
	}
}
