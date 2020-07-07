package pingaccess

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/iwarapter/pingaccess-sdk-go/pingaccess/models"
	"github.com/iwarapter/pingaccess-sdk-go/services/applications"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccPingAccessApplication(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPingAccessApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingAccessApplicationConfig("acc_test_bar", "/bar", "API"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessApplicationExists("pingaccess_application.acc_test"),
				),
			},
			{
				Config: testAccPingAccessApplicationConfig("acc_test_bar", "/bart", "Web"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessApplicationExists("pingaccess_application.acc_test"),
				),
			},
		},
	})
}

func testAccCheckPingAccessApplicationDestroy(s *terraform.State) error {
	return nil
}

func testAccPingAccessApplicationConfig(name, context, appType string) string {
	return fmt.Sprintf(`
	resource "pingaccess_site" "acc_test_site" {
		name                         = "acc_test_site"
		targets                      = ["localhost:4321"]
		max_connections              = -1
		max_web_socket_connections   = -1
		availability_profile_id      = 1
	}

	resource "pingaccess_virtualhost" "acc_test_virtualhost" {
		host                         = "acc-test-localhost"
		port                         = 4001
		agent_resource_cache_ttl     = 900
		key_pair_id                  = 0
		trusted_certificate_group_id = 0
 }

	resource "pingaccess_application" "acc_test" {
		access_validator_id = 0
		application_type 		= "Web"
		agent_id						= 0
		name								= "%s"
		context_root				= "%s"
		default_auth_type		= "Web"
		destination					= "Site"
		site_id							= "${pingaccess_site.acc_test_site.id}"
		spa_support_enabled = false
		virtual_host_ids		= ["${pingaccess_virtualhost.acc_test_virtualhost.id}"]
		web_session_id 			= "${pingaccess_websession.my_session.id}"

		// identity_mapping_ids {
		// 	web = 0
		// 	api = 0
		// }

		policy {
			web {
				type = "Rule"
				id = "${pingaccess_rule.acc_test_app_rule.id}"
			}
		}
	}

	resource "pingaccess_application" "acc_test_two" {
		application_type 		= "%s"
		agent_id						= 0
		name								= "api-demo"
		context_root				= "/"
		default_auth_type		= "%s"
		destination					= "Site"
		site_id							= "${pingaccess_site.acc_test_site.id}"
		virtual_host_ids		= ["${pingaccess_virtualhost.acc_test_virtualhost.id}"]
	}

	resource "pingaccess_identity_mapping" "idm_foo" {
		class_name = "com.pingidentity.pa.identitymappings.HeaderIdentityMapping"
		name       = "Foo"

		configuration = <<EOF
		{
			"attributeHeaderMappings": [
				{
					"subject": true,
					"attributeName": "FOO",
					"headerName": "FOO"
				}
			],
			"headerClientCertificateMappings": []
		}
		EOF
	}

	resource "pingaccess_rule" "acc_test_app_rule" {
		class_name = "com.pingidentity.pa.policy.CIDRPolicyInterceptor"
		name = "acc_test_app_rule"
		supported_destinations = [
			"Site",
			"Agent"
		]
		configuration = <<EOF
		{
			"cidrNotation": "127.0.0.1/32",
			"negate": false,
			"overrideIpSource": false,
			"headers": [],
			"headerValueLocation": "LAST",
			"fallbackToLastHopIp": true,
			"errorResponseCode": 403,
			"errorResponseStatusMsg": "Forbidden",
			"errorResponseTemplateFile": "policy.error.page.template.html",
			"errorResponseContentType": "text/html;charset=UTF-8",
			"rejectionHandler": null,
			"rejectionHandlingEnabled": false
		}
		EOF
	}

	resource "pingaccess_pingfederate_oauth" "app_demo_pfo" {
		client_id = "my_client"
		subject_attribute_name = "sany"
	}

	resource "pingaccess_pingfederate_runtime" "app_demo_pfr" {
		issuer = "%s"
		trusted_certificate_group_id = 2
	}

	resource "pingaccess_websession" "my_session" {
		depends_on = [pingaccess_pingfederate_runtime.app_demo_pfr, pingaccess_pingfederate_oauth.app_demo_pfo]
		name = "my-test-session"
		audience = "all"
		client_credentials {
			client_id = "websession"
			client_secret {
				value = "secret"
			}
		}
		scopes = [
			"profile",
			"address",
			"email",
			"phone"
		]
	}
	`, name, context, appType, appType, os.Getenv("PINGFEDERATE_TEST_IP"))
}

func testAccCheckPingAccessApplicationExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		// b, _ := json.Marshal(rs)
		// fmt.Printf("%s", b)

		if rs.Primary.ID == "" || rs.Primary.ID == "0" {
			return fmt.Errorf("No application ID is set")
		}

		conn := testAccProvider.Meta().(paClient).Applications
		result, _, err := conn.GetApplicationCommand(&applications.GetApplicationCommandInput{
			Id: rs.Primary.ID,
		})

		if err != nil {
			return fmt.Errorf("Error: Application (%s) not found", n)
		}

		if *result.Name != rs.Primary.Attributes["name"] {
			return fmt.Errorf("Error: Application response (%s) didnt match state (%s)", *result.Name, rs.Primary.Attributes["name"])
		}

		return nil
	}
}

func Test_flattenIdentityMappingIds(t *testing.T) {
	type args struct {
		in map[string]*int
	}
	tests := []struct {
		name string
		args args
		want []interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := flattenIdentityMappingIds(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("flattenIdentityMappingIds() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_resourcePingAccessApplicationReadData(t *testing.T) {
	cases := []struct {
		Application models.ApplicationView
	}{
		{
			Application: models.ApplicationView{
				Name:              String("engine1"),
				ApplicationType:   String("API"),
				AccessValidatorId: Int(0),
				AgentId:           Int(0),
				CaseSensitivePath: Bool(true),
				ContextRoot:       String("/"),
				DefaultAuthType:   String("API"),
				SiteId:            Int(0),
				SpaSupportEnabled: Bool(true),
				VirtualHostIds:    &[]*int{Int(1)},
				Policy: map[string]*[]*models.PolicyItem{
					"API": {
						{
							Id:   "1",
							Type: String("Rule"),
						},
						{
							Id:   "2",
							Type: String("Rule"),
						},
					},
					"Web": {},
				},
				WebSessionId: Int(0),
			},
		},
	}
	for i, tc := range cases {
		t.Run(fmt.Sprintf("tc:%v", i), func(t *testing.T) {

			resourceSchema := resourcePingAccessApplicationSchema()
			resourceLocalData := schema.TestResourceDataRaw(t, resourceSchema, map[string]interface{}{})
			resourcePingAccessApplicationReadResult(resourceLocalData, &tc.Application)

			if got := *resourcePingAccessApplicationReadData(resourceLocalData); !cmp.Equal(got, tc.Application) {
				t.Errorf("resourcePingAccessApplicationReadData() = %v", cmp.Diff(got, tc.Application))
			}
		})
	}
}
