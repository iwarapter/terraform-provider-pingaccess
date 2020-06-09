package pingaccess

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"regexp"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	pa "github.com/iwarapter/pingaccess-sdk-go/pingaccess"
)

func TestAccPingAccessIdentityMapping(t *testing.T) {
	resourceName := "pingaccess_identity_mapping.acc_test_idm"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPingAccessIdentityMappingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingAccessIdentityMappingConfig("bar", "SUB_HEADER"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessIdentityMappingExists(resourceName),
					testAccCheckPingAccessIdentityMappingAttributes(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", "bar"),
					resource.TestCheckResourceAttr(resourceName, "class_name", "com.pingidentity.pa.identitymappings.HeaderIdentityMapping"),
					resource.TestCheckResourceAttrSet(resourceName, "configuration"),
				),
			},
			{
				Config: testAccPingAccessIdentityMappingConfig("bar", "SUB"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessIdentityMappingExists(resourceName),
					testAccCheckPingAccessIdentityMappingAttributes(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", "bar"),
					resource.TestCheckResourceAttr(resourceName, "class_name", "com.pingidentity.pa.identitymappings.HeaderIdentityMapping"),
					resource.TestCheckResourceAttrSet(resourceName, "configuration"),
				),
			},
			{
				Config:      testAccPingAccessIdentityMappingConfigWrongClassName(),
				ExpectError: regexp.MustCompile(`unable to find className 'com.pingidentity.pa.identitymappings.foo' available className's: com.pingidentity.pa.identitymappings.HeaderIdentityMapping, com.pingidentity.pa.identitymappings.JwtIdentityMapping`),
			},
			{
				Config:      testAccPingAccessIdentityMappingConfigMissingRequired(),
				ExpectError: regexp.MustCompile(`configuration validation failed against the class descriptor definition\nthe field 'headerName' is required for the class_name 'com.pingidentity.pa.identitymappings.JwtIdentityMapping'\nthe field 'audience' is required for the class_name 'com.pingidentity.pa.identitymappings.JwtIdentityMapping'`),
			},
		},
	})
}

func testAccCheckPingAccessIdentityMappingDestroy(s *terraform.State) error {
	return nil
}

func testAccPingAccessIdentityMappingConfig(name, configUpdate string) string {
	return fmt.Sprintf(`
	resource "pingaccess_identity_mapping" "acc_test_idm" {
		class_name = "com.pingidentity.pa.identitymappings.HeaderIdentityMapping"
		name = "%s"
		configuration = <<EOF
		{
			"attributeHeaderMappings": [
				{
					"subject": true,
					"attributeName": "sub",
					"headerName": "%s"
				}
			],
			"headerClientCertificateMappings": []
		}
		EOF
	}`, name, configUpdate)
}

func testAccPingAccessIdentityMappingConfigWrongClassName() string {
	return `
	resource "pingaccess_identity_mapping" "acc_test_idm" {
		class_name = "com.pingidentity.pa.identitymappings.foo"
		name = "wrong_class_name"
		configuration = <<EOF
		{
			"attributeHeaderMappings": [
				{
					"subject": true,
					"attributeName": "sub",
					"headerName": "sub"
				}
			],
			"headerClientCertificateMappings": []
		}
		EOF
	}`
}

func testAccPingAccessIdentityMappingConfigMissingRequired() string {
	return `
	resource "pingaccess_identity_mapping" "acc_test_idm" {
		class_name = "com.pingidentity.pa.identitymappings.JwtIdentityMapping"
		name = "missing_required"
		configuration = "{}"
	}`
}

func testAccCheckPingAccessIdentityMappingExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" || rs.Primary.ID == "0" {
			return fmt.Errorf("No IdentityMapping ID is set")
		}

		conn := testAccProvider.Meta().(*pa.Client).IdentityMappings
		result, _, err := conn.GetIdentityMappingCommand(&pa.GetIdentityMappingCommandInput{
			Id: rs.Primary.ID,
		})

		if err != nil {
			return fmt.Errorf("Error: IdentityMapping (%s) not found", n)
		}

		if *result.Name != rs.Primary.Attributes["name"] {
			return fmt.Errorf("Error: IdentityMapping response (%s) didnt match state (%s)", *result.Name, rs.Primary.Attributes["name"])
		}

		return nil
	}
}

func testAccCheckPingAccessIdentityMappingAttributes(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs := s.RootModule().Resources[n]
		if rs.Primary.ID == "" || rs.Primary.ID == "0" {
			return fmt.Errorf("No IdentityMapping ID is set")
		}

		conn := testAccProvider.Meta().(*pa.Client).IdentityMappings
		result, _, err := conn.GetIdentityMappingCommand(&pa.GetIdentityMappingCommandInput{
			Id: rs.Primary.ID,
		})

		if err != nil {
			return fmt.Errorf("Error: IdentityMapping (%s) not found", n)
		}

		if *result.Name != rs.Primary.Attributes["name"] {
			return fmt.Errorf("Error: IdentityMapping response (%s) didnt match state (%s)", *result.Name, rs.Primary.Attributes["name"])
		}

		resultMapping := result.Configuration["attributeHeaderMappings"].([]interface{})[0].(map[string]interface{})["headerName"]
		var x map[string]interface{}
		json.Unmarshal([]byte(rs.Primary.Attributes["configuration"]), &x)
		stateMapping := x["attributeHeaderMappings"].([]interface{})[0].(map[string]interface{})["headerName"]

		if resultMapping != stateMapping {
			return fmt.Errorf("Error: IdentityMapping response (%s) didnt match state (%s)", resultMapping, stateMapping)
		}

		return nil
	}
}

func Test_resourcePingAccessIdentityMappingReadData(t *testing.T) {
	descs := pa.DescriptorsView{
		Items: []*pa.DescriptorView{
			{
				ClassName: String("something"),
				ConfigurationFields: []*pa.ConfigurationField{
					{
						Name: String("password"),
						Type: String("CONCEALED"),
					},
				},
				Label: nil,
				Type:  nil,
			},
		}}

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		equals(t, req.URL.String(), "/identityMappings/descriptors")
		// Send response to be tested
		b, _ := json.Marshal(descs)
		rw.Write(b)
	}))
	// Close the server when test finishes
	defer server.Close()

	// Use Client & URL from our local test server
	url, _ := url.Parse(server.URL)
	c := pa.NewClient("", "", url, "", server.Client())

	cases := []struct {
		Name            string
		IdentityMapping pa.IdentityMappingView
		ExpectedDiags   diag.Diagnostics
	}{
		{
			Name: "Stuff breaks",
			IdentityMapping: pa.IdentityMappingView{
				Name:      String("foo"),
				ClassName: String("foo"),
			},
			ExpectedDiags: diag.Diagnostics{},
		},
	}
	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {

			resourceSchema := resourcePingAccessIdentityMappingSchema()
			resourceLocalData := schema.TestResourceDataRaw(t, resourceSchema, map[string]interface{}{})
			diags := resourcePingAccessIdentityMappingReadResult(resourceLocalData, &tc.IdentityMapping, c.IdentityMappings)

			checkDiagnostics(t, tc.Name, diags, tc.ExpectedDiags)

			if got := *resourcePingAccessIdentityMappingReadData(resourceLocalData); !cmp.Equal(got, tc.IdentityMapping) {
				t.Errorf("resourcePingAccessIdentityMappingReadData() = %v", cmp.Diff(got, tc.IdentityMapping))
			}
		})
	}
}

func checkDiagnostics(t *testing.T, tn string, got, expected diag.Diagnostics) {
	if len(got) != len(expected) {
		t.Fatalf("%s: wrong number of diags, expected %d, got %d", tn, len(expected), len(got))
	}
	for j := range got {
		if got[j].Severity != expected[j].Severity {
			t.Fatalf("%s: expected severity %v, got %v", tn, expected[j].Severity, got[j].Severity)
		}
		if !got[j].AttributePath.Equals(expected[j].AttributePath) {
			t.Fatalf("%s: attribute paths do not match expected: %v, got %v", tn, expected[j].AttributePath, got[j].AttributePath)
		}
	}
}
