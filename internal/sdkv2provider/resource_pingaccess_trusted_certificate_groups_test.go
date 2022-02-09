package sdkv2provider

import (
	"fmt"
	"testing"

	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess/models"
	"github.com/iwarapter/pingaccess-sdk-go/v62/services/trustedCertificateGroups"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func init() {
	resource.AddTestSweepers("trusted_certificate_group", &resource.Sweeper{
		Name: "trusted_certificate_group",
		F: func(r string) error {
			svc := trustedCertificateGroups.New(conf)
			results, _, err := svc.GetTrustedCertificateGroupsCommand(&trustedCertificateGroups.GetTrustedCertificateGroupsCommandInput{Filter: "acctest_"})
			if err != nil {
				return fmt.Errorf("unable to list trusted_certificate_groups to sweep %s", err)
			}
			for _, item := range results.Items {
				_, err = svc.DeleteTrustedCertificateGroupCommand(&trustedCertificateGroups.DeleteTrustedCertificateGroupCommandInput{Id: item.Id.String()})
				if err != nil {
					return fmt.Errorf("unable to sweep trusted_certificate_group %s because %s", item.Id.String(), err)
				}
			}
			return nil
		},
	})
}

func TestAccPingAccessTrustedCertificateGroups(t *testing.T) {
	resourceName := "pingaccess_trusted_certificate_group.demo_tcg"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV5ProviderFactories: testAccProviders,
		CheckDestroy:             testAccCheckPingAccessTrustedCertificateGroupsDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingAccessTrustedCertificateGroupsConfig("acctest_service", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessTrustedCertificateGroupsExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", "acctest_service"),
					resource.TestCheckResourceAttr(resourceName, "use_java_trust_store", "true"),
					resource.TestCheckResourceAttr(resourceName, "skip_certificate_date_check", "true"),
				),
			},
			{
				Config: testAccPingAccessTrustedCertificateGroupsConfig("acctest_service", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessTrustedCertificateGroupsExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", "acctest_service"),
					resource.TestCheckResourceAttr(resourceName, "use_java_trust_store", "true"),
					resource.TestCheckResourceAttr(resourceName, "skip_certificate_date_check", "false"),
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

func testAccCheckPingAccessTrustedCertificateGroupsDestroy(s *terraform.State) error {
	return nil
}

func testAccPingAccessTrustedCertificateGroupsConfig(name, skipDateCheck string) string {
	return fmt.Sprintf(`
	resource "pingaccess_trusted_certificate_group" "demo_tcg" {
		name = "%s"
		use_java_trust_store = true
		skip_certificate_date_check = %s
	}`, name, skipDateCheck)
}

func testAccCheckPingAccessTrustedCertificateGroupsExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" || rs.Primary.ID == "0" {
			return fmt.Errorf("No third party service ID is set")
		}

		conn := testAccProvider.Meta().(paClient).TrustedCertificateGroups
		result, _, err := conn.GetTrustedCertificateGroupCommand(&trustedCertificateGroups.GetTrustedCertificateGroupCommandInput{
			Id: rs.Primary.ID,
		})

		if err != nil {
			return fmt.Errorf("Error: TrustedCertificateGroups (%s) not found", n)
		}

		if *result.Name != rs.Primary.Attributes["name"] {
			return fmt.Errorf("Error: TrustedCertificateGroups response (%s) didnt match state (%s)", *result.Name, rs.Primary.Attributes["name"])
		}

		return nil
	}
}

func Test_resourcePingAccessTrustedCertificateGroupsReadData(t *testing.T) {
	cases := []struct {
		TrustedCertificateGroups models.TrustedCertificateGroupView
	}{
		{
			TrustedCertificateGroups: models.TrustedCertificateGroupView{
				Name:                       String("localhost"),
				CertIds:                    &[]*int{Int(1)},
				UseJavaTrustStore:          Bool(false),
				SystemGroup:                Bool(false),
				IgnoreAllCertificateErrors: Bool(false),
				SkipCertificateDateCheck:   Bool(false),
			},
		},
	}
	for i, tc := range cases {
		t.Run(fmt.Sprintf("tc:%v", i), func(t *testing.T) {

			resourceSchema := resourcePingAccessTrustedCertificateGroupsSchema()
			resourceLocalData := schema.TestResourceDataRaw(t, resourceSchema, map[string]interface{}{})
			resourcePingAccessTrustedCertificateGroupsReadResult(resourceLocalData, &tc.TrustedCertificateGroups)

			if got := *resourcePingAccessTrustedCertificateGroupsReadData(resourceLocalData); !cmp.Equal(got, tc.TrustedCertificateGroups) {
				t.Errorf("resourcePingAccessTrustedCertificateGroupsReadData() = %v", cmp.Diff(got, tc.TrustedCertificateGroups))
			}
		})
	}
}
