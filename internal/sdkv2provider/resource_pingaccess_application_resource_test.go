package sdkv2provider

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"testing"

	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess/models"
	"github.com/iwarapter/pingaccess-sdk-go/v62/services/applications"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccPingAccessApplicationResource(t *testing.T) {
	policy1 := `web {
		type = "Rule"
		id = "${pingaccess_rule.acc_test_resource_rule.id}"
	}

	web {
		type = "Rule"
		id = "${pingaccess_rule.acc_test_resource_rule_two.id}"
	}
	`

	policy2 := `web {
		type = "Rule"
		id = "${pingaccess_rule.acc_test_resource_rule_two.id}"
	}
	web {
		type = "Rule"
		id = "${pingaccess_rule.acc_test_resource_rule.id}"
	}`
	resourceName := "pingaccess_application_resource.app_res_test_resource"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV5ProviderFactories: testAccProviders,
		CheckDestroy:             testAccCheckPingAccessApplicationResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingAccessApplicationResourceConfig("acc_test_bart", "/bar", policy1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessApplicationResourceExists(resourceName),
					testAccCheckPingAccessApplicationResourceAttributes(resourceName, "bart", "/bar"),
				),
			},
			{
				Config: testAccPingAccessApplicationResourceConfig("acc_test_bart", "/bart", policy2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessApplicationResourceExists(resourceName),
					testAccCheckPingAccessApplicationResourceAttributes(resourceName, "bart", "/bart"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: func(d *terraform.State) (string, error) {
					rs, ok := d.RootModule().Resources[resourceName]
					if !ok {
						return "", fmt.Errorf("unable to find resource %s", resourceName)
					}
					return fmt.Sprintf("%s/%s", rs.Primary.Attributes["application_id"], rs.Primary.ID), nil
				},
			},
		},
	})
}

func testAccCheckPingAccessApplicationResourceDestroy(s *terraform.State) error {
	return nil
}

func testAccPingAccessApplicationResourceConfig(name string, context string, policy string) string {
	return fmt.Sprintf(`
	resource "pingaccess_site" "app_res_test_site" {
		name                         = "acc_test_app_res_test_site"
		targets                      = ["localhost:4321"]
		max_connections              = -1
		max_web_socket_connections   = -1
		availability_profile_id      = 1
	}

	resource "pingaccess_virtualhost" "app_res_test_virtualhost" {
		host                         = "acc-test-localhost"
		port                         = 4000
		agent_resource_cache_ttl     = 900
		key_pair_id                  = 0
		trusted_certificate_group_id = 0
 	}

	resource "pingaccess_application" "app_res_test" {
		access_validator_id = 0
		application_type 	= "Web"
		agent_id			= 0
		name				= "%s"
		context_root		= "/bar"
		default_auth_type	= "Web"
		destination			= "Site"
		site_id				= pingaccess_site.app_res_test_site.id
		virtual_host_ids	= [pingaccess_virtualhost.app_res_test_virtualhost.id]

		identity_mapping_ids {
			web = 0
			api = 0
		}
	}

resource "pingaccess_application_resource" "app_res_test_resource" {
  name = "acc_test_woot"
  methods = [
    "*"
  ]

  path_patterns {
    pattern = "/as/token.oauth2"
    type    = "WILDCARD"
  }

  path_patterns {
    pattern = "%s"
    type    = "WILDCARD"
  }

  path_prefixes = [
	"/as/token.oauth2",
	"%s"
  ]
  audit_level = "OFF"
  anonymous = false
  enabled = true
  root_resource = false
  application_id = pingaccess_application.app_res_test.id

  policy {
	%s
  }

  resource_type = "Standard"

}

resource "pingaccess_application_resource" "app_res_virtual_resource" {
  name = "acc_test_woot_virtual"
  methods = [
    "*"
  ]

  path_patterns {
    pattern = "/virtual"
    type    = "WILDCARD"
  }

  path_prefixes = [
	"/virtual"
  ]
  audit_level = "OFF"
  anonymous = false
  enabled = true
  root_resource = false
  application_id = pingaccess_application.app_res_test.id

  resource_type = "Virtual"

  resource_type_configuration {
      response_generator {
          class_name = "com.pingidentity.pa.resources.responsegenerator.TemplateResponseGenerator"
          configuration = <<EOF
          {
              "template": "<h1>Hello from Virtual Resource</h1>",
              "responseCode": 200,
              "mediaType": "text/html;charset=utf-8"
          }
          EOF
      }
  }

  depends_on = [pingaccess_application_resource.app_res_test_resource]
}

resource "pingaccess_application_resource" "app_res_test_root_resource" {
  name = "Root Resource"
  methods = [
    "*"
	]

  path_prefixes = [
		"/*"
  ]

  path_patterns {
    pattern = "/*"
    type    = "WILDCARD"
  }

  policy {}

  audit_level = "ON"
  anonymous = false
  enabled = true
  root_resource = true
  application_id = pingaccess_application.app_res_test.id

  resource_type = "Standard"
}

resource "pingaccess_rule" "acc_test_resource_rule" {
	class_name = "com.pingidentity.pa.policy.CIDRPolicyInterceptor"
	name = "acc_test_resource_rule"
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

resource "pingaccess_rule" "acc_test_resource_rule_two" {
	class_name = "com.pingidentity.pa.policy.CIDRPolicyInterceptor"
	name = "acc_test_resource_rule_two"
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
	`, name, context, context, policy)
}

func testAccCheckPingAccessApplicationResourceExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		b, _ := json.Marshal(rs)
		log.Printf("[INFO] RS: %s", b)
		if rs.Primary.ID == "" || rs.Primary.ID == "0" {
			return fmt.Errorf("No application resource ID is set")
		}

		conn := testAccProvider.Meta().(paClient).Applications
		result, _, err := conn.GetApplicationResourceCommand(&applications.GetApplicationResourceCommandInput{
			ApplicationId: rs.Primary.Attributes["application_id"],
			ResourceId:    rs.Primary.ID,
		})

		if err != nil {
			return fmt.Errorf("Error: Application (%s) not found", n)
		}

		if *result.Name != rs.Primary.Attributes["name"] {
			return fmt.Errorf("Error: Application Resource response (%s) didnt match state (%s)", *result.Name, rs.Primary.Attributes["name"])
		}

		return nil
	}
}

func testAccCheckPingAccessApplicationResourceAttributes(n, name, context string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		conn := testAccProvider.Meta().(paClient).Applications
		result, _, err := conn.GetApplicationResourceCommand(&applications.GetApplicationResourceCommandInput{
			ApplicationId: rs.Primary.Attributes["application_id"],
			ResourceId:    rs.Primary.ID,
		})

		if err != nil {
			return fmt.Errorf("Error: Application (%s) not found", n)
		}

		if *result.Name != rs.Primary.Attributes["name"] {
			return fmt.Errorf("Error: Application Resource response (%s) didnt match state (%s)", *result.Name, rs.Primary.Attributes["name"])
		}

		if strconv.Itoa(len(*result.Methods)) != rs.Primary.Attributes["methods.#"] {
			return fmt.Errorf("Error: Application Resource response (%s) didnt match state (%s)", strconv.Itoa(len(*result.Methods)), rs.Primary.Attributes["methods.#"])
		}

		if len(*result.Policy["Web"]) != 2 {
			return fmt.Errorf("Expected 2 web policies, got: %v", *result.Policy["Web"])
		}

		return nil
	}
}

func Test_resourcePingAccessApplicationResourceReadData(t *testing.T) {
	cases := []struct {
		Resource models.ResourceView
	}{

		{
			Resource: models.ResourceView{
				Anonymous:               Bool(false),
				ApplicationId:           Int(0),
				AuditLevel:              String("false"),
				DefaultAuthTypeOverride: String("false"),
				Enabled:                 Bool(false),
				Methods:                 &[]*string{String("false")},
				Name:                    String("false"),
				// PathPatterns: []*pa.PathPatternView{
				// 	&pa.PathPatternView{
				// 		Pattern: String("/*"),
				// 		Type:    String("WILDCARD"),
				// 	},
				// },
				PathPrefixes: &[]*string{String("false")},
				Policy: map[string]*[]*models.PolicyItem{
					"Web": {
						{
							Id:   json.Number("1"),
							Type: String("Rule"),
						},
						{
							Id:   json.Number("2"),
							Type: String("RuleSet"),
						},
					},
					"API": {},
				},
				RootResource: Bool(false),
				Unprotected:  Bool(false),
				ResourceType: String("Standard"),
			},
		},

		{
			Resource: models.ResourceView{
				Anonymous:               Bool(false),
				ApplicationId:           Int(0),
				AuditLevel:              String("false"),
				DefaultAuthTypeOverride: String("false"),
				Enabled:                 Bool(false),
				Methods:                 &[]*string{String("false")},
				Name:                    String("false"),
				// PathPatterns: []*pa.PathPatternView{
				// 	&pa.PathPatternView{
				// 		Pattern: String("/*"),
				// 		Type:    String("WILDCARD"),
				// 	},
				// },
				PathPrefixes: &[]*string{String("false")},
				Policy: map[string]*[]*models.PolicyItem{
					"Web": {
						{
							Id:   json.Number("1"),
							Type: String("Rule"),
						},
						{
							Id:   json.Number("2"),
							Type: String("RuleSet"),
						},
					},
					"API": {},
				},
				RootResource: Bool(false),
				Unprotected:  Bool(false),
				ResourceType: String("Virtual"),

				ResourceTypeConfiguration: &models.ResourceTypeConfigurationView{
					ResponseGenerator: &models.ResponseGeneratorView{
						ClassName: String("com.pingidentity.pa.resources.responsegenerator.TemplateResponseGenerator"),
						Configuration: map[string]interface{}{
							"test": "value",
						},
					},
				},
			},
		},

		{
			Resource: models.ResourceView{
				ApplicationId: Int(0),
				Methods:       &[]*string{String("GET")},
				Name:          String("false"),
				Anonymous:     Bool(false),
				AuditLevel:    String("OFF"),
				Enabled:       Bool(false),
				RootResource:  Bool(true),
				Unprotected:   Bool(true),
				ResourceType:  String("Standard"),
			},
		},
	}
	for i, tc := range cases {
		t.Run(fmt.Sprintf("tc:%v", i), func(t *testing.T) {

			resourceSchema := resourcePingAccessApplicationResourceSchema()
			resourceLocalData := schema.TestResourceDataRaw(t, resourceSchema, map[string]interface{}{})
			resourcePingAccessApplicationResourceReadResult(resourceLocalData, &tc.Resource)

			if got := *resourcePingAccessApplicationResourceReadData(resourceLocalData); !cmp.Equal(got, tc.Resource) {
				t.Errorf("resourcePingAccessApplicationResourceReadData() = %v", cmp.Diff(got, tc.Resource))
			}
		})
	}
}
