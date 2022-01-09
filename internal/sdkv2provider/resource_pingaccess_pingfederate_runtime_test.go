package sdkv2provider

import (
	"fmt"
	"net/url"
	"os"
	"testing"

	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestAccPingAccessPingFederateRuntimeIssuer(t *testing.T) {
	if !(paClient{apiVersion: paVersion}).Is60OrAbove() {
		t.Skipf("This test only runs against PingAccess 6.0 and above, not: %s", paVersion)
	}
	resourceName := "pingaccess_pingfederate_runtime.demo"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV5ProviderFactories: testAccProviders,
		CheckDestroy:             testAccCheckPingAccessPingFederateRuntimeDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingAccessPingFederateRuntimeConfig(os.Getenv("PINGFEDERATE_TEST_IP"), "foo"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessPingFederateRuntimeExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "description", "foo"),
					resource.TestCheckResourceAttr(resourceName, "skip_hostname_verification", "true"),
					resource.TestCheckResourceAttr(resourceName, "sts_token_exchange_endpoint", "https://foo/bar"),
					resource.TestCheckResourceAttr(resourceName, "use_slo", "false"),
					resource.TestCheckResourceAttr(resourceName, "trusted_certificate_group_id", "2"),
					resource.TestCheckResourceAttr(resourceName, "use_proxy", "true"),
				),
			},
			{
				Config: testAccPingAccessPingFederateRuntimeConfig(os.Getenv("PINGFEDERATE_TEST_IP"), "bar"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessPingFederateRuntimeExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "description", "bar"),
					resource.TestCheckResourceAttr(resourceName, "skip_hostname_verification", "true"),
					resource.TestCheckResourceAttr(resourceName, "sts_token_exchange_endpoint", "https://foo/bar"),
					resource.TestCheckResourceAttr(resourceName, "use_slo", "false"),
					resource.TestCheckResourceAttr(resourceName, "trusted_certificate_group_id", "2"),
					resource.TestCheckResourceAttr(resourceName, "use_proxy", "true"),
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

func TestAccPingAccessPingFederateRuntimeDeprecatedRuntime(t *testing.T) {
	if !(paClient{apiVersion: paVersion}).Is62OrAbove() {
		t.Skipf("This test only runs against PingAccess 5.3, 6.0 or 6.1, not: %s", paVersion)
	}
	resourceName := "pingaccess_pingfederate_runtime.demo"
	u, _ := url.Parse(os.Getenv("PINGFEDERATE_TEST_IP"))

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV5ProviderFactories: testAccProviders,
		CheckDestroy:             testAccCheckPingAccessPingFederateRuntimeDestroy,
		Steps: []resource.TestStep{

			{
				Config: testAccPingAccessPingFederateDeprecatedRuntimeConfig(u.Hostname(), u.Port(), "ON"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessPingFederateDeprecatedRuntimeExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "skip_hostname_verification", "true"),
					resource.TestCheckResourceAttr(resourceName, "use_slo", "false"),
					resource.TestCheckResourceAttr(resourceName, "trusted_certificate_group_id", "2"),
					resource.TestCheckResourceAttr(resourceName, "use_proxy", "true"),
					resource.TestCheckResourceAttr(resourceName, "audit_level", "ON"),
					resource.TestCheckResourceAttr(resourceName, "back_channel_secure", "true"),
					resource.TestCheckResourceAttr(resourceName, "host", u.Hostname()),
					resource.TestCheckResourceAttr(resourceName, "port", u.Port()),
					resource.TestCheckResourceAttr(resourceName, "secure", "false"),
				),
			},
			{
				Config: testAccPingAccessPingFederateDeprecatedRuntimeConfig(u.Hostname(), u.Port(), "OFF"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessPingFederateDeprecatedRuntimeExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "skip_hostname_verification", "true"),
					resource.TestCheckResourceAttr(resourceName, "use_slo", "false"),
					resource.TestCheckResourceAttr(resourceName, "trusted_certificate_group_id", "2"),
					resource.TestCheckResourceAttr(resourceName, "use_proxy", "true"),
					resource.TestCheckResourceAttr(resourceName, "audit_level", "OFF"),
					resource.TestCheckResourceAttr(resourceName, "back_channel_secure", "true"),
					resource.TestCheckResourceAttr(resourceName, "host", u.Hostname()),
					resource.TestCheckResourceAttr(resourceName, "port", u.Port()),
					resource.TestCheckResourceAttr(resourceName, "secure", "false"),
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

func TestAccPingAccessPingFederateRuntimeNewConfig(t *testing.T) {
	if !(paClient{apiVersion: paVersion}).Is61OrAbove() {
		t.Skipf("This test only runs against PingAccess 6.1 or above, not: %s", paVersion)
	}
	resourceName := "pingaccess_pingfederate_runtime.demo"
	u, _ := url.Parse(os.Getenv("PINGFEDERATE_TEST_IP"))

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV5ProviderFactories: testAccProviders,
		CheckDestroy:             testAccCheckPingAccessPingFederateRuntimeDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingAccessPingFederateNewRuntimeConfig(u.Hostname(), u.Port(), "ON"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessPingFederateDeprecatedRuntimeExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "skip_hostname_verification", "true"),
					resource.TestCheckResourceAttr(resourceName, "use_slo", "false"),
					resource.TestCheckResourceAttr(resourceName, "trusted_certificate_group_id", "2"),
					resource.TestCheckResourceAttr(resourceName, "use_proxy", "true"),
					resource.TestCheckResourceAttr(resourceName, "audit_level", "ON"),
					resource.TestCheckResourceAttr(resourceName, "back_channel_secure", "true"),
					resource.TestCheckResourceAttr(resourceName, "targets.0", u.Hostname()+":"+u.Port()),
					resource.TestCheckResourceAttr(resourceName, "secure", "false"),
				),
			},
			{
				Config: testAccPingAccessPingFederateNewRuntimeConfig(u.Hostname(), u.Port(), "OFF"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessPingFederateDeprecatedRuntimeExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "skip_hostname_verification", "true"),
					resource.TestCheckResourceAttr(resourceName, "use_slo", "false"),
					resource.TestCheckResourceAttr(resourceName, "trusted_certificate_group_id", "2"),
					resource.TestCheckResourceAttr(resourceName, "use_proxy", "true"),
					resource.TestCheckResourceAttr(resourceName, "audit_level", "OFF"),
					resource.TestCheckResourceAttr(resourceName, "back_channel_secure", "true"),
					resource.TestCheckResourceAttr(resourceName, "targets.0", u.Hostname()+":"+u.Port()),
					resource.TestCheckResourceAttr(resourceName, "secure", "false"),
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

func testAccCheckPingAccessPingFederateRuntimeDestroy(_ *terraform.State) error {
	return nil
}

func testAccPingAccessPingFederateRuntimeConfig(issuer, configChange string) string {
	return fmt.Sprintf(`
	resource "pingaccess_pingfederate_runtime" "demo" {
		description = "%s"
		issuer = "%s"
		skip_hostname_verification = true
		sts_token_exchange_endpoint = "https://foo/bar"
		use_slo = false
		trusted_certificate_group_id = 2
		use_proxy = true
	}`, configChange, issuer)
}

func testAccPingAccessPingFederateDeprecatedRuntimeConfig(host, port, configChange string) string {
	return fmt.Sprintf(`
	resource "pingaccess_pingfederate_runtime" "demo" {
		host = "%s"
		port = %s
		audit_level = "%s"
		skip_hostname_verification = true
		use_slo = false
		trusted_certificate_group_id = 2
		use_proxy = true
  		back_channel_secure = true
	}`, host, port, configChange)
}

func testAccPingAccessPingFederateNewRuntimeConfig(host, port, configChange string) string {
	return fmt.Sprintf(`
	resource "pingaccess_pingfederate_runtime" "demo" {
		targets = ["%s:%s"]
		audit_level = "%s"
		skip_hostname_verification = true
		use_slo = false
		trusted_certificate_group_id = 2
		use_proxy = true
		application {
	    	primary_virtual_host_id = 1
  		}
  		back_channel_secure = true
	}`, host, port, configChange)
}

func testAccCheckPingAccessPingFederateRuntimeExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not found: %s", n)
		}

		conn := testAccProvider.Meta().(paClient).Pingfederate
		result, _, err := conn.GetPingFederateRuntimeCommand()

		if err != nil {
			return fmt.Errorf("error: PingFederateRuntime (%s) not found", n)
		}

		if *result.Issuer != rs.Primary.Attributes["issuer"] {
			return fmt.Errorf("error: PingFederateRuntime response (%s) didnt match state (%s)", *result.Issuer, rs.Primary.Attributes["issuer"])
		}

		return nil
	}
}

func testAccCheckPingAccessPingFederateDeprecatedRuntimeExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not found: %s", n)
		}

		conn := testAccProvider.Meta().(paClient).Pingfederate
		result, resp, err := conn.GetPingFederateCommand()

		if err != nil {
			return fmt.Errorf("error: PingFederateDeprecatedRuntime (%s) not found", n)
		}

		if result == nil {
			return fmt.Errorf("error: PingFederateDeprecatedRuntime (%s) not response %v", err, resp)
		}

		if *result.Host != rs.Primary.Attributes["host"] {
			return fmt.Errorf("error: PingFederateDeprecatedRuntime response (%s) didnt match state (%s)", *result.Host, rs.Primary.Attributes["host"])
		}

		return nil
	}
}

func Test_resourcePingAccessPingFederateRuntimeReadData(t *testing.T) {
	cases := []struct {
		PingFederateRuntime models.PingFederateMetadataRuntimeView
	}{
		{
			PingFederateRuntime: models.PingFederateMetadataRuntimeView{
				Issuer:                    String("localhost"),
				SkipHostnameVerification:  Bool(true),
				UseProxy:                  Bool(false),
				UseSlo:                    Bool(false),
				TrustedCertificateGroupId: Int(0),
			},
		},
		{
			PingFederateRuntime: models.PingFederateMetadataRuntimeView{
				Issuer:                    String("localhost"),
				Description:               String("foo"),
				SkipHostnameVerification:  Bool(true),
				StsTokenExchangeEndpoint:  String("https://foo/bar"),
				TrustedCertificateGroupId: Int(0),
				UseProxy:                  Bool(true),
				UseSlo:                    Bool(true),
			},
		},
	}
	for i, tc := range cases {
		t.Run(fmt.Sprintf("tc:%v", i), func(t *testing.T) {

			resourceSchema := resourcePingAccessPingFederateRuntimeSchema()
			resourceLocalData := schema.TestResourceDataRaw(t, resourceSchema, map[string]interface{}{})
			resourcePingAccessPingFederateRuntimeReadResult(resourceLocalData, &tc.PingFederateRuntime)

			if got := *resourcePingAccessPingFederateRuntimeReadData(resourceLocalData); !cmp.Equal(got, tc.PingFederateRuntime) {
				t.Errorf("resourcePingAccessPingFederateRuntimeReadData() = %v", cmp.Diff(got, tc.PingFederateRuntime))
			}
		})
	}
}

func Test_resourcePingAccessPingFederateDeprecatedRuntimeReadData(t *testing.T) {
	cases := []struct {
		PingFederateRuntime models.PingFederateRuntimeView
	}{
		{
			PingFederateRuntime: models.PingFederateRuntimeView{
				Host:                      String("localhost"),
				Port:                      Int(9031),
				SkipHostnameVerification:  Bool(true),
				UseProxy:                  Bool(false),
				UseSlo:                    Bool(false),
				AuditLevel:                String("ON"),
				TrustedCertificateGroupId: Int(0),
			},
		},
		{
			PingFederateRuntime: models.PingFederateRuntimeView{
				Host:                      String("localhost"),
				Port:                      Int(9031),
				AuditLevel:                String("ON"),
				BackChannelBasePath:       String("/foo"),
				BackChannelSecure:         Bool(true),
				BasePath:                  String("/bar"),
				ExpectedHostname:          String("hosty"),
				Secure:                    Bool(true),
				Targets:                   &[]*string{String("t1:9031")},
				SkipHostnameVerification:  Bool(true),
				TrustedCertificateGroupId: Int(2),
				UseProxy:                  Bool(true),
				UseSlo:                    Bool(true),
			},
		},
		{
			PingFederateRuntime: models.PingFederateRuntimeView{
				//Defaults
				AuditLevel:        String("ON"),
				BackChannelSecure: Bool(true),
				Host:              String(""),
				Port:              Int(0),
				UseProxy:          Bool(false),
				UseSlo:            Bool(false),
				//test
				Targets:                   &[]*string{String("t1:9031")},
				SkipHostnameVerification:  Bool(true),
				TrustedCertificateGroupId: Int(2),
				Application: &models.PingFederateRuntimeApplicationView{
					PrimaryVirtualHostId:     Int(1),
					AdditionalVirtualHostIds: &[]*int{},
					CaseSensitive:            Bool(true),
					ClientCertHeaderNames:    &[]*string{},
					ContextRoot:              String("/"),
					Policy: []*models.PolicyItem{
						{
							Id:   "1",
							Type: String("Rule"),
						},
						{
							Id:   "2",
							Type: String("Rule"),
						},
					},
				},
			},
		},
	}
	for i, tc := range cases {
		t.Run(fmt.Sprintf("tc:%v", i), func(t *testing.T) {

			resourceSchema := resourcePingAccessPingFederateRuntimeSchema()
			resourceLocalData := schema.TestResourceDataRaw(t, resourceSchema, map[string]interface{}{})
			resourcePingAccessPingFederateDeprecatedRuntimeReadResult(resourceLocalData, &tc.PingFederateRuntime)

			if got := *resourcePingAccessPingFederateDeprecatedRuntimeReadData(resourceLocalData); !cmp.Equal(got, tc.PingFederateRuntime) {
				t.Errorf("resourcePingAccessPingFederateRuntimeReadData() = %v", cmp.Diff(got, tc.PingFederateRuntime))
			}
		})
	}
}
