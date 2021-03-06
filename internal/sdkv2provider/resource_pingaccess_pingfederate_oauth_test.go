package sdkv2provider

import (
	"fmt"
	"testing"

	"github.com/iwarapter/pingaccess-sdk-go/pingaccess/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestAccPingAccessPingFederateOAuth(t *testing.T) {
	resourceName := "pingaccess_pingfederate_oauth.demo_pfo"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV5ProviderFactories: testAccProviders,
		CheckDestroy:             testAccCheckPingAccessPingFederateOAuthDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingAccessPingFederateOAuthConfig("my_client", "san"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessPingFederateOAuthExists(resourceName),
				),
			},
			{
				Config: testAccPingAccessPingFederateOAuthConfig("my_client", "sany"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessPingFederateOAuthExists(resourceName),
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

func testAccCheckPingAccessPingFederateOAuthDestroy(s *terraform.State) error {
	return nil
}

func testAccPingAccessPingFederateOAuthConfig(client, san string) string {
	return fmt.Sprintf(`
	resource "pingaccess_pingfederate_oauth" "demo_pfo" {
		client_id = "%s"
		subject_attribute_name = "%s"
	}`, client, san)
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

		if *result.ClientId != rs.Primary.Attributes["client_id"] {
			return fmt.Errorf("Error: PingFederateOAuth response (%s) didnt match state (%s)", *result.ClientId, rs.Primary.Attributes["client_id"])
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
	}
	for i, tc := range cases {
		t.Run(fmt.Sprintf("tc:%v", i), func(t *testing.T) {

			resourceSchema := resourcePingAccessPingFederateOAuthSchema()
			resourceLocalData := schema.TestResourceDataRaw(t, resourceSchema, map[string]interface{}{})
			resourcePingAccessPingFederateOAuthReadResult(resourceLocalData, &tc.PingFederateAccessTokenView)

			if got := *resourcePingAccessPingFederateOAuthReadData(resourceLocalData); !cmp.Equal(got, tc.PingFederateAccessTokenView) {
				t.Errorf("resourcePingAccessPingFederateOAuthReadData() = %v", cmp.Diff(got, tc.PingFederateAccessTokenView))
			}
		})
	}
}
