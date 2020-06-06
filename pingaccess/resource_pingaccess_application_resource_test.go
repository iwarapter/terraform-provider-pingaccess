package pingaccess

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	pa "github.com/iwarapter/pingaccess-sdk-go/pingaccess"
)

func init() {
	resource.AddTestSweepers("pingaccess_application_resource", &resource.Sweeper{
		Name: "pingaccess_application_resource",
		F:    testSweepApplicationResources,
	})
}

func testSweepApplicationResources(r string) error {
	url, _ := url.Parse("https://localhost:9000")
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	conn := pa.NewClient("Administrator", "2Access2", url, "/pa-admin-api/v3", nil).Applications
	result, _, _ := conn.GetResourcesCommand(&pa.GetResourcesCommandInput{Filter: "acc_test_"})
	for _, v := range result.Items {
		conn.DeleteApplicationResourceCommand(&pa.DeleteApplicationResourceCommandInput{ApplicationId: strconv.Itoa(*v.ApplicationId), ResourceId: v.Id.String()})
	}
	return nil
}

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

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPingAccessApplicationResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingAccessApplicationResourceConfig("acc_test_bart", "/bar", policy1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessApplicationResourceExists("pingaccess_application_resource.app_res_test_resource"),
					testAccCheckPingAccessApplicationResourceAttributes("pingaccess_application_resource.app_res_test_resource", "bart", "/bar"),
				),
			},
			{
				Config: testAccPingAccessApplicationResourceConfig("acc_test_bart", "/bart", policy2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessApplicationResourceExists("pingaccess_application_resource.app_res_test_resource"),
					testAccCheckPingAccessApplicationResourceAttributes("pingaccess_application_resource.app_res_test_resource", "bart", "/bart"),
				),
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
		application_type 		= "Web"
		agent_id						= 0
		name								= "%s"
		context_root				= "/bar"
		default_auth_type		= "Web"
		destination					= "Site"
		site_id							= "${pingaccess_site.app_res_test_site.id}"
		virtual_host_ids		= ["${pingaccess_virtualhost.app_res_test_virtualhost.id}"]

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
	application_id = "${pingaccess_application.app_res_test.id}"

	policy {
		%s
	}
}

resource "pingaccess_application_resource" "app_res_test_root_resource" {
  name = "Root Resource"
  methods = [
    "*"
	]

  path_prefixes = [
		"/*"
  ]

	policy {}

  audit_level = "ON"
  anonymous = false
  enabled = true
  root_resource = true
  application_id = "${pingaccess_application.app_res_test.id}"
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

		conn := testAccProvider.Meta().(*pa.Client).Applications
		result, _, err := conn.GetApplicationResourceCommand(&pa.GetApplicationResourceCommandInput{
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

		conn := testAccProvider.Meta().(*pa.Client).Applications
		result, _, err := conn.GetApplicationResourceCommand(&pa.GetApplicationResourceCommandInput{
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
		Resource pa.ResourceView
	}{
		{
			Resource: pa.ResourceView{
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
				Policy: map[string]*[]*pa.PolicyItem{
					"Web": &[]*pa.PolicyItem{
						&pa.PolicyItem{
							Id:   json.Number("1"),
							Type: String("Rule"),
						},
						&pa.PolicyItem{
							Id:   json.Number("2"),
							Type: String("RuleSet"),
						},
					},
					"API": &[]*pa.PolicyItem{},
				},
				RootResource: Bool(false),
				Unprotected:  Bool(false),
			},
		},
		{
			Resource: pa.ResourceView{
				ApplicationId: Int(0),
				Methods:       &[]*string{String("GET")},
				Name:          String("false"),
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
