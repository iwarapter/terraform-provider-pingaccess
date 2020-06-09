package pingaccess

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccPingAccessPingFederateRuntimeMetadataDataSource(t *testing.T) {
	resourceName := "data.pingaccess_pingfederate_runtime_metadata.test"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				//We run two steps so the first enables the PF runtime, the second the DS can then query it.
				Config: testAccPingAccessPingFederateRuntimeMetadataConfig(),
			},
			{
				Config: testAccPingAccessPingFederateRuntimeMetadataConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "authorization_endpoint"),
					resource.TestCheckResourceAttrSet(resourceName, "backchannel_authentication_endpoint"),
					resource.TestCheckResourceAttrSet(resourceName, "claim_types_supported.0"),
					resource.TestCheckResourceAttrSet(resourceName, "claims_parameter_supported"),
					resource.TestCheckNoResourceAttr(resourceName, "claims_supported.0"),
					resource.TestCheckNoResourceAttr(resourceName, "code_challenge_methods_supported.0"),
					resource.TestCheckNoResourceAttr(resourceName, "end_session_endpoint"),
					resource.TestCheckResourceAttrSet(resourceName, "grant_types_supported.0"),
					resource.TestCheckResourceAttrSet(resourceName, "id_token_signing_alg_values_supported.0"),
					resource.TestCheckResourceAttrSet(resourceName, "introspection_endpoint"),
					resource.TestCheckResourceAttrSet(resourceName, "issuer"),
					resource.TestCheckResourceAttrSet(resourceName, "jwks_uri"),
					resource.TestCheckResourceAttrSet(resourceName, "ping_end_session_endpoint"),
					resource.TestCheckResourceAttrSet(resourceName, "ping_revoked_sris_endpoint"),
					resource.TestCheckResourceAttrSet(resourceName, "request_object_signing_alg_values_supported.0"),
					resource.TestCheckResourceAttrSet(resourceName, "request_parameter_supported"),
					resource.TestCheckResourceAttrSet(resourceName, "request_uri_parameter_supported"),
					resource.TestCheckResourceAttrSet(resourceName, "response_modes_supported.0"),
					resource.TestCheckResourceAttrSet(resourceName, "response_types_supported.0"),
					resource.TestCheckResourceAttrSet(resourceName, "revocation_endpoint"),
					resource.TestCheckResourceAttrSet(resourceName, "scopes_supported.0"),
					resource.TestCheckResourceAttrSet(resourceName, "subject_types_supported.0"),
					resource.TestCheckResourceAttrSet(resourceName, "token_endpoint"),
					resource.TestCheckResourceAttrSet(resourceName, "token_endpoint_auth_methods_supported.0"),
					resource.TestCheckResourceAttrSet(resourceName, "userinfo_endpoint"),
					resource.TestCheckNoResourceAttr(resourceName, "userinfo_signing_alg_values_supported.0"),
				),
			},
		},
	})
}

func testAccPingAccessPingFederateRuntimeMetadataConfig() string {
	return fmt.Sprintf(`data "pingaccess_pingfederate_runtime_metadata" "test" {}
resource "pingaccess_pingfederate_runtime" "app_demo_pfr" {
	issuer = "https://%s:9031"
	trusted_certificate_group_id = 2
}`, os.Getenv("PINGFEDERATE_TEST_IP"))
}
