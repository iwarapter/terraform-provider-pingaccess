package pingaccess

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	pa "github.com/iwarapter/pingaccess-sdk-go/pingaccess"
)

func TestAccPingAccessAuthnReqList(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPingAccessAuthnReqListDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingAccessAuthnReqListConfig("foo", "bar"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessAuthnReqListExists("pingaccess_authn_req_list.acc_test"),
				),
			},
			{
				Config: testAccPingAccessAuthnReqListConfig("bar", "foo"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessAuthnReqListExists("pingaccess_authn_req_list.acc_test"),
				),
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
	   name   = "foo"
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

		conn := testAccProvider.Meta().(*pa.Client).AuthnReqLists
		result, _, err := conn.GetAuthnReqListCommand(&pa.GetAuthnReqListCommandInput{
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
		AuthnReqList pa.AuthnReqListView
	}{
		{
			AuthnReqList: pa.AuthnReqListView{
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
