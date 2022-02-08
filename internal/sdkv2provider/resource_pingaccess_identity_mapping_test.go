package sdkv2provider

import (
	"encoding/json"
	"fmt"
	"regexp"
	"testing"

	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess/models"
	"github.com/iwarapter/pingaccess-sdk-go/v62/services/identityMappings"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func init() {
	resource.AddTestSweepers("identity_mapping", &resource.Sweeper{
		Name: "identity_mapping",
		F: func(r string) error {
			svc := identityMappings.New(conf)
			results, _, err := svc.GetIdentityMappingsCommand(&identityMappings.GetIdentityMappingsCommandInput{Filter: "acctest_"})
			if err != nil {
				return fmt.Errorf("unable to list identity mappings to sweep %s", err)
			}
			for _, item := range results.Items {
				_, err = svc.DeleteIdentityMappingCommand(&identityMappings.DeleteIdentityMappingCommandInput{Id: item.Id.String()})
				if err != nil {
					return fmt.Errorf("unable to sweep identity mappings %s because %s", item.Id.String(), err)
				}
			}
			return nil
		},
	})
}

func TestAccPingAccessIdentityMapping(t *testing.T) {
	resourceName := "pingaccess_identity_mapping.acc_test_idm"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV5ProviderFactories: testAccProviders,
		CheckDestroy:             testAccCheckPingAccessIdentityMappingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingAccessIdentityMappingConfig("bar", "SUB_HEADER"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessIdentityMappingExists(resourceName),
					testAccCheckPingAccessIdentityMappingAttributes(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", "acctest_bar"),
					resource.TestCheckResourceAttr(resourceName, "class_name", "com.pingidentity.pa.identitymappings.HeaderIdentityMapping"),
					resource.TestCheckResourceAttrSet(resourceName, "configuration"),
				),
			},
			{
				Config: testAccPingAccessIdentityMappingConfig("bar", "SUB"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessIdentityMappingExists(resourceName),
					testAccCheckPingAccessIdentityMappingAttributes(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", "acctest_bar"),
					resource.TestCheckResourceAttr(resourceName, "class_name", "com.pingidentity.pa.identitymappings.HeaderIdentityMapping"),
					resource.TestCheckResourceAttrSet(resourceName, "configuration"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config:      testAccPingAccessIdentityMappingConfigWrongClassName(),
				ExpectError: regexp.MustCompile(`unable to find className 'com.pingidentity.pa.identitymappings.foo' available classNames: com.pingidentity.pa.identitymappings.HeaderIdentityMapping, com.pingidentity.pa.identitymappings.JwtIdentityMapping`),
			},
			{
				Config:      testAccPingAccessIdentityMappingConfigMissingRequired(),
				ExpectError: regexp.MustCompile(`the field 'audience' is required for the class_name 'com.pingidentity.pa.identitymappings.JwtIdentityMapping'`),
			},
			{
				Config: testAccPingAccessIdentityMappingConfigInterpolatedSkipped(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessIdentityMappingExists(resourceName),
					testAccCheckPingAccessIdentityMappingAttributes(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", "acctest_interpolated"),
					resource.TestCheckResourceAttr(resourceName, "class_name", "com.pingidentity.pa.identitymappings.HeaderIdentityMapping"),
					resource.TestCheckResourceAttrSet(resourceName, "configuration"),
				),
			},
		},
	})
}

func testAccCheckPingAccessIdentityMappingDestroy(s *terraform.State) error {
	return nil
}

func testAccPingAccessIdentityMappingConfig(name, configUpdate string) string {
	block := ""
	if (paClient{apiVersion: paVersion}).Is62OrAbove() {
		block = `"exclusionList": false,
			"exclusionListAttributes": [],
			"exclusionListSubject": null,
			"headerNamePrefix": null,`
	}
	return fmt.Sprintf(`
	resource "pingaccess_identity_mapping" "acc_test_idm" {
		class_name = "com.pingidentity.pa.identitymappings.HeaderIdentityMapping"
		name = "acctest_%s"
		configuration = <<EOF
		{
			%s
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
	}`, name, block, configUpdate)
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

func testAccPingAccessIdentityMappingConfigInterpolatedSkipped() string {
	block := ""
	if (paClient{apiVersion: paVersion}).Is61OrAbove() {
		block = `"exclusionList": false,
			"exclusionListAttributes": [],
			"exclusionListSubject": null,
			"headerNamePrefix": null,`
	}
	return fmt.Sprintf(`
	resource "pingaccess_identity_mapping" "acc_test_idm" {
		class_name = "com.pingidentity.pa.identitymappings.HeaderIdentityMapping"
		name = "acctest_interpolated"
		configuration = <<EOF
		{
			%s
			"attributeHeaderMappings": [
				{
					"subject": true,
					"attributeName": "sub",
					"headerName": "${pingaccess_virtualhost.interpolate_me.id}"
				}
			],
			"headerClientCertificateMappings": []
		}
		EOF
	}

	resource "pingaccess_virtualhost" "interpolate_me" {
	   host                         = "idmfoo"
	   port                         = 80
	}
`, block)
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

		conn := testAccProvider.Meta().(paClient).IdentityMappings
		result, _, err := conn.GetIdentityMappingCommand(&identityMappings.GetIdentityMappingCommandInput{
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

		conn := testAccProvider.Meta().(paClient).IdentityMappings
		result, _, err := conn.GetIdentityMappingCommand(&identityMappings.GetIdentityMappingCommandInput{
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
	desc := &models.DescriptorsView{
		Items: []*models.DescriptorView{
			{
				ClassName: String("something"),
				ConfigurationFields: []*models.ConfigurationField{
					{
						Name: String("password"),
						Type: String("CONCEALED"),
					},
				},
				Label: nil,
				Type:  nil,
			},
		}}
	cases := []struct {
		Name            string
		IdentityMapping models.IdentityMappingView
		ExpectedDiags   diag.Diagnostics
	}{
		{
			Name: "Stuff breaks",
			IdentityMapping: models.IdentityMappingView{
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
			diags := resourcePingAccessIdentityMappingReadResult(resourceLocalData, &tc.IdentityMapping, desc)

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
