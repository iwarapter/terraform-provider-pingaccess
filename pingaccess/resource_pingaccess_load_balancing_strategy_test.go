package pingaccess

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/iwarapter/pingaccess-sdk-go/pingaccess"
)

var rrConfig = map[string]interface{}{
	"stickySessionEnabled": true,
	"cookieName":           "PA_S",
}
var headerConfig = map[string]interface{}{
	"fallbackToFirstAvailableHost": true,
	"headerName":                   "COOKIE-D",
}

func init() {
	resource.AddTestSweepers("pingaccess_loadbalancingstrategies", &resource.Sweeper{
		Name:         "pingaccess_loadbalancingstrategies",
		F:            testSweepLoadBalancingStrategies,
		Dependencies: []string{"pingaccess_application", "pingaccess_application_resource"},
	})
}

func testSweepLoadBalancingStrategies(r string) error {
	url, _ := url.Parse("https://localhost:9000")
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	conn := pingaccess.NewClient("Administrator", "2Access2", url, "/pa-admin-api/v3", nil).HighAvailability
	result, _, _ := conn.GetLoadBalancingStrategiesCommand(&pingaccess.GetLoadBalancingStrategiesCommandInput{Filter: "acc-test-"})
	for _, v := range result.Items {
		log.Printf("Sweeper: Deleting %s", *v.Name)
		conn.DeleteLoadBalancingStrategyCommand(&pingaccess.DeleteLoadBalancingStrategyCommandInput{Id: v.Id.String()})
	}
	return nil
}

func TestAccPingAccessLoadBalancingStrategy(t *testing.T) {
	var out pingaccess.LoadBalancingStrategyView
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPingAccessLoadBalancingStrategyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingAccessLoadBalancingStrategyConfig("RR Load balancer", cookieBasedClassName, rrConfig),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessLoadBalancingStrategyExists("pingaccess_load_balancing_strategy.acc_test", 3, &out),
				),
			},
			{
				Config: testAccPingAccessLoadBalancingStrategyConfig("Header based Load balancer", headerBasedClassName, headerConfig),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessLoadBalancingStrategyExists("pingaccess_load_balancing_strategy.acc_test", 6, &out),
				),
			},
		},
	})
}

func testAccCheckPingAccessLoadBalancingStrategyDestroy(s *terraform.State) error {
	return nil
}

func testAccPingAccessLoadBalancingStrategyConfig(name string, className string, configuration map[string]interface{}) string {
	jsonConfig, _ := json.Marshal(configuration)
	return fmt.Sprintf(`
	resource "pingaccess_load_balancing_strategy" "acc_test" {
	    name                         = "tf-acc-test-%s"
	    class_name                    = "%s"
	    configuration = <<EOF
		%v
		EOF
	}`, name, className, string(jsonConfig))
}

func testAccCheckPingAccessLoadBalancingStrategyExists(n string, c int64, out *pingaccess.LoadBalancingStrategyView) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" || rs.Primary.ID == "0" {
			return fmt.Errorf("No load balancing strategy ID is set")
		}

		conn := testAccProvider.Meta().(*pingaccess.Client).HighAvailability
		result, _, err := conn.GetLoadBalancingStrategyCommand(&pingaccess.GetLoadBalancingStrategyCommandInput{
			Id: rs.Primary.ID,
		})
		if err != nil {
			return fmt.Errorf("Error: Load Balancing Strategy (%s) not found", n)
		}

		if *result.ClassName != rs.Primary.Attributes["class_name"] {
			return fmt.Errorf("Error: Load Balancing Strategy response (%s) didnt match state (%s)", *result.ClassName, rs.Primary.Attributes["class_name"])
		}
		return nil
	}
}
