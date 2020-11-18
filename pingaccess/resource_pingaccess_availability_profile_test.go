package pingaccess

import (
	"fmt"
	"github.com/iwarapter/pingaccess-sdk-go/services/highAvailability"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccPingAccessAvailabilityProfile(t *testing.T) {
	resourceName := "pingaccess_availability_profile.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPingAccessAvailabilityProfileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingAccessAvailabilityProfileConfig("10000"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessAvailabilityProfileExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", "foo"),
					resource.TestCheckResourceAttr(resourceName, "class_name", "com.pingidentity.pa.ha.availability.ondemand.OnDemandAvailabilityPlugin"),
					resource.TestCheckResourceAttr(resourceName, "configuration", "{\"connectTimeout\":10000,\"failedRetryTimeout\":60,\"failureHttpStatusCodes\":[],\"maxRetries\":2,\"pooledConnectionTimeout\":-1,\"readTimeout\":-1,\"retryDelay\":250}"),
				),
			},
			{
				Config: testAccPingAccessAvailabilityProfileConfig("5000"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessAvailabilityProfileExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", "foo"),
					resource.TestCheckResourceAttr(resourceName, "class_name", "com.pingidentity.pa.ha.availability.ondemand.OnDemandAvailabilityPlugin"),
					resource.TestCheckResourceAttr(resourceName, "configuration", "{\"connectTimeout\":5000,\"failedRetryTimeout\":60,\"failureHttpStatusCodes\":[],\"maxRetries\":2,\"pooledConnectionTimeout\":-1,\"readTimeout\":-1,\"retryDelay\":250}"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config:      testAccPingAccessAvailabilityProfileConfigInvalidClassName(),
				ExpectError: regexp.MustCompile(`unable to find className 'com.pingidentity.pa.AvailabilityProfiles.foo' available classNames: com.pingidentity.pa.ha.availability.ondemand.OnDemandAvailabilityPlugin`),
			},
		},
	})
}

func testAccCheckPingAccessAvailabilityProfileDestroy(s *terraform.State) error {
	return nil
}

func testAccPingAccessAvailabilityProfileConfig(configUpdate string) string {
	return fmt.Sprintf(`
	resource "pingaccess_availability_profile" "test" {
		class_name = "com.pingidentity.pa.ha.availability.ondemand.OnDemandAvailabilityPlugin"
		name = "foo"

		configuration = <<EOF
		{
			"connectTimeout": %s,
			"pooledConnectionTimeout": -1,
			"readTimeout": -1,
			"maxRetries": 2,
			"retryDelay": 250,
			"failedRetryTimeout": 60,
			"failureHttpStatusCodes": []
		}
		EOF
	}`, configUpdate)
}

func testAccPingAccessAvailabilityProfileConfigInvalidClassName() string {
	return `
	resource "pingaccess_availability_profile" "test" {
		class_name = "com.pingidentity.pa.AvailabilityProfiles.foo"
		name = "foo"

		configuration = <<EOF
		{
			"description": null,
			"path": "/foo",
			"subjectAttributeName": "foo",
			"issuer": null,
			"audience": null
		}
		EOF
	}`
}

func testAccCheckPingAccessAvailabilityProfileExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" || rs.Primary.ID == "0" {
			return fmt.Errorf("No availability_profile ID is set")
		}

		conn := testAccProvider.Meta().(paClient).HighAvailability
		result, _, err := conn.GetAvailabilityProfileCommand(&highAvailability.GetAvailabilityProfileCommandInput{
			Id: rs.Primary.ID,
		})

		if err != nil {
			return fmt.Errorf("Error: AvailabilityProfile (%s) not found", n)
		}

		if *result.Name != rs.Primary.Attributes["name"] {
			return fmt.Errorf("Error: AvailabilityProfile response (%s) didnt match state (%s)", *result.Name, rs.Primary.Attributes["name"])
		}

		return nil
	}
}
