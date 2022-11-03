package sdkv2provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess/models"
	"github.com/iwarapter/pingaccess-sdk-go/v62/services/applications"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func init() {
	resource.AddTestSweepers("applications", &resource.Sweeper{
		Name: "applications",
		F: func(r string) error {
			svc := applications.New(conf)
			results, _, err := svc.GetApplicationsCommand(&applications.GetApplicationsCommandInput{})
			if err != nil {
				return fmt.Errorf("unable to list applications to sweep %s", err)
			}
			for _, item := range results.Items {
				_, err = svc.DeleteApplicationCommand(&applications.DeleteApplicationCommandInput{Id: item.Id.String()})
				if err != nil {
					return fmt.Errorf("unable to sweep application %s because %s", item.Id.String(), err)
				}
			}
			return nil
		},
	})
}

func TestAccPingAccessApplication(t *testing.T) {
	resourceName := "pingaccess_application.acc_test"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV5ProviderFactories: testAccProviders,
		CheckDestroy:             testAccCheckPingAccessApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingAccessApplicationConfig("acc_test_bar", "/bar", "API"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessApplicationExists(resourceName),
				),
			},
			{
				Config: testAccPingAccessApplicationConfig("acc_test_bar", "/bart", "Web"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessApplicationExists(resourceName),
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

func testAccCheckPingAccessApplicationDestroy(s *terraform.State) error {
	return nil
}

func testAccPingAccessApplicationConfig(name, context, appType string) string {
	block := ""
	oauth := `client_id              = "foo"
  client_secret {
    value = "top_secret"
  }`
	if (paClient{apiVersion: paVersion}).Is62OrAbove() {
		block = `"exclusionList": false,
			"exclusionListAttributes": [],
			"exclusionListSubject": null,
			"headerNamePrefix": null,`
	}
	if (paClient{apiVersion: paVersion}).Is61OrAbove() {
		oauth = `client_credentials {
			credentials_type = "SECRET"
			client_id = "my_client"
			client_secret {
				value = "secret"
			}
		}`
	}
	return fmt.Sprintf(`
resource "pingaccess_site" "acc_test_site" {
  name                       = "acctest_site"
  targets                    = ["localhost:4321"]
  max_connections            = -1
  max_web_socket_connections = -1
  availability_profile_id    = 1
}

resource "pingaccess_virtualhost" "acc_test_virtualhost" {
  host                         = "acctest-localhost"
  port                         = 4001
  agent_resource_cache_ttl     = 900
  key_pair_id                  = 0
  trusted_certificate_group_id = 0
}

resource "pingaccess_application" "acc_test" {
  access_validator_id = 0
  application_type    = "Web"
  agent_id            = 0
  name                = "%s"
  context_root        = "%s"
  destination         = "Site"
  site_id             = "${pingaccess_site.acc_test_site.id}"
  spa_support_enabled = false
  virtual_host_ids    = ["${pingaccess_virtualhost.acc_test_virtualhost.id}"]
  web_session_id      = "${pingaccess_websession.my_session.id}"

  // identity_mapping_ids {
  // 	web = 0
  // 	api = 0
  // }

  policy {
    web {
      type = "Rule"
      id   = "${pingaccess_rule.acc_test_app_rule.id}"
    }
  }
}

resource "pingaccess_application" "acc_test_two" {
  application_type  = "%s"
  agent_id          = 0
  name              = "api-demo"
  context_root      = "/"
  default_auth_type = "%s"
  destination       = "Site"
  site_id           = "${pingaccess_site.acc_test_site.id}"
  virtual_host_ids  = ["${pingaccess_virtualhost.acc_test_virtualhost.id}"]
}

resource "pingaccess_identity_mapping" "idm_foo" {
  class_name = "com.pingidentity.pa.identitymappings.HeaderIdentityMapping"
  name       = "acctest_foo"

  configuration = <<EOF
		{
			%s
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
  name       = "acctest_app_rule"
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
		%s
  subject_attribute_name = "sany"
}

resource "pingaccess_pingfederate_runtime" "app_demo_pfr" {
  issuer                       = "%s"
  trusted_certificate_group_id = 2
}

resource "pingaccess_websession" "my_session" {
  depends_on = [pingaccess_pingfederate_runtime.app_demo_pfr, pingaccess_pingfederate_oauth.app_demo_pfo]
  name       = "acctest_session"
  audience   = "all"
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

resource "pingaccess_application" "app_res_test" {
  access_validator_id     = 0
  application_type        = "Web"
  agent_id                = 0
  name                    = "acc_test_regex"
  context_root            = "/ordering"
  default_auth_type       = "Web"
  destination             = "Site"
  manual_ordering_enabled = true
  site_id                 = pingaccess_site.acc_test_site.id
  virtual_host_ids        = [pingaccess_virtualhost.acc_test_virtualhost.id]
  identity_mapping_ids {
    web = 0
    api = 0
  }
}

resource "pingaccess_application_resource" "app_res_test_resource" {
  name    = "acc_test_regex_pattern"
  methods = ["*"]

  path_patterns {
    pattern = "/.*"
    type    = "REGEX"
  }

  audit_level    = "OFF"
  anonymous      = false
  enabled        = true
  root_resource  = false
  application_id = pingaccess_application.app_res_test.id
  resource_type  = "Standard"
}




	`, name, context, appType, appType, block, oauth, os.Getenv("PINGFEDERATE_TEST_IP"))
}

func testAccCheckPingAccessApplicationExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

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
				ManualOrderingEnabled: Bool(true),
				WebSessionId:          Int(0),
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

func Test_issue158(t *testing.T) {
	resourceName := "pingaccess_application.acc_test"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV5ProviderFactories: testAccProviders,
		CheckDestroy:             testAccCheckPingAccessApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: `
resource "pingaccess_application" "acc_test" {
  name                    = "Test"
  agent_id                = 0
  enabled                 = true
  require_https           = true
  description             = "Managed with Terraform"
  application_type        = "Web"
  destination             = "Site"
  context_root            = "/test/contextroot"
  manual_ordering_enabled = true
  case_sensitive_path     = false
  site_id                 = pingaccess_site.site_test.id
  virtual_host_ids        = [pingaccess_virtualhost.vh_test.id]
  access_validator_id     = 0
  spa_support_enabled     = true
}

resource "pingaccess_virtualhost" "vh_test" {
  host                     = "localhost"
  port                     = 443
  agent_resource_cache_ttl = 900
}

data "pingaccess_trusted_certificate_group" "trust_any" {
  name = "Trust Any"
}
resource "pingaccess_site" "site_test" {
  name                         = "Site test"
  secure                       = true
  skip_hostname_verification   = true
  targets                      = ["google.com:80"]
  trusted_certificate_group_id = data.pingaccess_trusted_certificate_group.trust_any.id
  send_pa_cookie               = false
}`,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessApplicationExists(resourceName),
				),
			},
			{
				Config: `
resource "pingaccess_application" "acc_test" {
  name                    = "Test2"
  agent_id                = 0
  enabled                 = true
  require_https           = true
  description             = "Managed with Terraform"
  application_type        = "Web"
  destination             = "Site"
  context_root            = "/test/contextroot"
  manual_ordering_enabled = true
  case_sensitive_path     = false
  site_id                 = pingaccess_site.site_test.id
  virtual_host_ids        = [pingaccess_virtualhost.vh_test.id]
  access_validator_id     = 0
  spa_support_enabled     = true
}

resource "pingaccess_virtualhost" "vh_test" {
  host                     = "localhost"
  port                     = 443
  agent_resource_cache_ttl = 900
}

data "pingaccess_trusted_certificate_group" "trust_any" {
  name = "Trust Any"
}

resource "pingaccess_site" "site_test" {
  name                         = "Site test"
  secure                       = true
  skip_hostname_verification   = true
  targets                      = ["google.com:80"]
  trusted_certificate_group_id = data.pingaccess_trusted_certificate_group.trust_any.id
  send_pa_cookie               = false
}`,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessApplicationExists(resourceName),
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

func Test_issue48(t *testing.T) {
	app := &models.ApplicationView{
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
		ManualOrderingEnabled: Bool(true),
		WebSessionId:          Int(0),
	}

	resourceSchema := resourcePingAccessApplicationSchema()
	resourceLocalData := schema.TestResourceDataRaw(t, resourceSchema, map[string]interface{}{})
	resourcePingAccessApplicationReadResult(resourceLocalData, app)

	exp := resourcePingAccessApplicationReadData(resourceLocalData)
	assert.Equal(t, exp, app)

	app.Policy["API"] = &[]*models.PolicyItem{}
	assert.NotEqual(t, exp, app) //make sure they are now different
	//when we read it should now become equal
	resourcePingAccessApplicationReadResult(resourceLocalData, app)
	exp = resourcePingAccessApplicationReadData(resourceLocalData)
	assert.Equal(t, exp, app)
}
