package pingaccess

import (
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/iwarapter/pingaccess-sdk-go/pingaccess"
)

//Provider does stuff
//
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"region": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: descriptions["region"],
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			// "pingaccess_rule":        resourcePingAccessRule(),
			"pingaccess_virtualhost": resourcePingAccessVirtualHost(),
			"pingaccess_site":        resourcePingAccessSite(),
		},
		ConfigureFunc: providerConfigure,
	}
}

var descriptions map[string]string

func init() {
	descriptions = map[string]string{
		"region": "The region where AWS operations will take place. Examples\n" +
			"are us-east-1, us-west-2, etc.",
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	// config := &Client{
	// 	Username: d.Get("region").(string),
	// }
	config := &pingaccess.Config{
		Username: "Administrator",
		Password: "2Access2",
		BaseURL:  "https://localhost:9000/pa-admin-api/v3",
	}

	return config, nil
}

// Takes the result of flatmap.Expand for an array of strings
// and returns a []*string
func expandStringList(configured []interface{}) []*string {
	log.Printf("[INFO] expandStringList %d", len(configured))
	vs := make([]*string, 0, len(configured))
	for _, v := range configured {
		val := v.(string)
		if val != "" {
			vs = append(vs, &val)
			log.Printf("[DEBUG] Appending: %s", val)
		}
	}
	return vs
}

// Takes the result of flatmap.Expand for an array of strings
// and returns a []*int
func expandIntList(configured []interface{}) []*int {
	vs := make([]*int, 0, len(configured))
	for _, v := range configured {
		_, ok := v.(*int)
		if ok {
			vs = append(vs, v.(*int))
		}
	}
	return vs
}
