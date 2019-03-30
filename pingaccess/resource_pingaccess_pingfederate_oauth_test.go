package pingaccess

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform/helper/schema"
	pa "github.com/iwarapter/pingaccess-sdk-go/pingaccess"
)

// func TestAccPingAccessPingFederateOAuth(t *testing.T) {
// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:     func() { testAccPreCheck(t) },
// 		Providers:    testAccProviders,
// 		CheckDestroy: testAccCheckPingAccessPingFederateOAuthDestroy,
// 		Steps: []resource.TestStep{
// 			{
// 				Config: testAccPingAccessPingFederateOAuthConfig("my_client", "san"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckPingAccessPingFederateOAuthExists("pingaccess_pingfederate_oauth.demo_pfo"),
// 				),
// 			},
// 			{
// 				Config: testAccPingAccessPingFederateOAuthConfig("my_client", "sany"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckPingAccessPingFederateOAuthExists("pingaccess_pingfederate_oauth.demo_pfo"),
// 				),
// 			},
// 		},
// 	})
// }

// func testAccCheckPingAccessPingFederateOAuthDestroy(s *terraform.State) error {
// 	return nil
// }

// func testAccPingAccessPingFederateOAuthConfig(client, san string) string {
// 	return fmt.Sprintf(`
// 	resource "pingaccess_pingfederate_oauth" "demo_pfo" {
// 		client_id = "%s"
// 		subject_attribute_name = "%s"
// 	}`, client, san)
// }

// func testAccCheckPingAccessPingFederateOAuthExists(n string) resource.TestCheckFunc {
// 	return func(s *terraform.State) error {
// 		rs, ok := s.RootModule().Resources[n]
// 		if !ok {
// 			return fmt.Errorf("Not found: %s", n)
// 		}

// 		if rs.Primary.ID == "" || rs.Primary.ID == "0" {
// 			return fmt.Errorf("No third party service ID is set")
// 		}

// 		conn := testAccProvider.Meta().(*pa.Client).PingFederate
// 		result, _, err := conn.GetPingFederateAccessTokensCommand()

// 		if err != nil {
// 			return fmt.Errorf("Error: PingFederateOAuth (%s) not found", n)
// 		}

// 		if *result.ClientId != rs.Primary.Attributes["client_id"] {
// 			return fmt.Errorf("Error: PingFederateOAuth response (%s) didnt match state (%s)", *result.ClientId, rs.Primary.Attributes["client_id"])
// 		}

// 		return nil
// 	}
// }

func Test_resourcePingAccessPingFederateOAuthReadData(t *testing.T) {
	cases := []struct {
		PingFederateAccessTokenView pa.PingFederateAccessTokenView
	}{
		{
			PingFederateAccessTokenView: pa.PingFederateAccessTokenView{
				ClientId:               String("client_1"),
				SubjectAttributeName:   String("san"),
				AccessValidatorId:      Int(1),
				Name:                   String("PingFederate"),
				TokenTimeToLiveSeconds: Int(-1),
			},
		},
		{
			PingFederateAccessTokenView: pa.PingFederateAccessTokenView{
				ClientId:               String("client_1"),
				SubjectAttributeName:   String("san"),
				AccessValidatorId:      Int(1),
				CacheTokens:            Bool(true),
				Name:                   String("PingFederate"),
				SendAudience:           Bool(true),
				TokenTimeToLiveSeconds: Int(30),
				UseTokenIntrospection:  Bool(true),
				ClientSecret: &pa.HiddenFieldView{
					Value: String("password"),
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
