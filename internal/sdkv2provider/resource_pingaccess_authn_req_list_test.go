package sdkv2provider

import (
	"fmt"
	"testing"

	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess/models"
	"github.com/iwarapter/pingaccess-sdk-go/v62/services/authnReqLists"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func init() {
	resource.AddTestSweepers("authn_req_list", &resource.Sweeper{
		Name: "authn_req_list",
		F: func(r string) error {
			svc := authnReqLists.New(conf)
			results, _, err := svc.GetAuthnReqListsCommand(&authnReqLists.GetAuthnReqListsCommandInput{Filter: "acctest_"})
			if err != nil {
				return fmt.Errorf("unable to list authnReqLists to sweep %s", err)
			}
			for _, item := range results.Items {
				_, err = svc.DeleteAuthnReqListCommand(&authnReqLists.DeleteAuthnReqListCommandInput{Id: item.Id.String()})
				if err != nil {
					return fmt.Errorf("unable to sweep authnReqLists %s because %s", item.Id, err)
				}
			}
			return nil
		},
	})
}

func TestAccPingAccessAuthnReqList(t *testing.T) {
	resourceName := "pingaccess_authn_req_list.acc_test"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV5ProviderFactories: testAccProviders,
		CheckDestroy:             testAccCheckPingAccessAuthnReqListDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingAccessAuthnReqListConfig("foo", "bar"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessAuthnReqListExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", "acctest_foo"),
					resource.TestCheckResourceAttr(resourceName, "authn_reqs.0", "foo"),
					resource.TestCheckResourceAttr(resourceName, "authn_reqs.1", "bar"),
				),
			},
			{
				Config: testAccPingAccessAuthnReqListConfig("bar", "foo"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessAuthnReqListExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", "acctest_foo"),
					resource.TestCheckResourceAttr(resourceName, "authn_reqs.0", "bar"),
					resource.TestCheckResourceAttr(resourceName, "authn_reqs.1", "foo"),
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

func testAccCheckPingAccessAuthnReqListDestroy(s *terraform.State) error {
	return nil
}

func testAccPingAccessAuthnReqListConfig(req1, req2 string) string {
	return fmt.Sprintf(`
	resource "pingaccess_authn_req_list" "acc_test" {
	   name   = "acctest_foo"
	   authn_reqs = [
			 "%s",
			 "%s",
		 ]
	}`, req1, req2)
}

func testAccCheckPingAccessAuthnReqListExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" || rs.Primary.ID == "0" {
			return fmt.Errorf("No AuthnReqList ID is set")
		}

		conn := testAccProvider.Meta().(paClient).AuthnReqLists
		result, _, err := conn.GetAuthnReqListCommand(&authnReqLists.GetAuthnReqListCommandInput{
			Id: rs.Primary.ID,
		})

		if err != nil {
			return fmt.Errorf("Error: AuthnReqList (%s) not found", n)
		}

		if *result.Name != rs.Primary.Attributes["name"] {
			return fmt.Errorf("Error: AuthnReqList response (%s) didnt match state (%s)", *result.Name, rs.Primary.Attributes["name"])
		}

		return nil
	}
}

func Test_resourcePingAccessAuthnReqListReadData(t *testing.T) {
	cases := []struct {
		AuthnReqList models.AuthnReqListView
	}{
		{
			AuthnReqList: models.AuthnReqListView{
				Name: String("engine1"),
				AuthnReqs: &[]*string{
					String("foo"),
					String("bar"),
				},
			},
		},
	}
	for i, tc := range cases {
		t.Run(fmt.Sprintf("tc:%v", i), func(t *testing.T) {

			resourceSchema := resourcePingAccessAuthnReqListSchema()
			resourceLocalData := schema.TestResourceDataRaw(t, resourceSchema, map[string]interface{}{})
			resourcePingAccessAuthnReqListReadResult(resourceLocalData, &tc.AuthnReqList)

			if got := *resourcePingAccessAuthnReqListReadData(resourceLocalData); !cmp.Equal(got, tc.AuthnReqList) {
				t.Errorf("resourcePingAccessAuthnReqListReadData() = %v", cmp.Diff(got, tc.AuthnReqList))
			}
		})
	}
}
