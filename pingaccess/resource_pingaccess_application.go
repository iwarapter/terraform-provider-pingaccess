package pingaccess

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"strconv"

// 	"github.com/hashicorp/terraform/helper/schema"
// 	"github.com/iwarapter/pingaccess-sdk-go/pingaccess"
// )

// func resourcePingAccessApplication() *schema.Resource {
// 	setOfString := &schema.Schema{
// 		Type:     schema.TypeSet,
// 		Optional: true,
// 		Elem: &schema.Schema{
// 			Type: schema.TypeString,
// 		},
// 	}

// 	return &schema.Resource{
// 		Create: resourcePingAccessApplicationCreate,
// 		Read:   resourcePingAccessApplicationRead,
// 		Update: resourcePingAccessApplicationUpdate,
// 		Delete: resourcePingAccessApplicationDelete,

// 		Schema: map[string]*schema.Schema{
// 			"access_validator_id": &schema.Schema{
// 				Type:     schema.TypeInt,
// 				Optional: true,
// 			},
// 			"agent_id": &schema.Schema{
// 				Type:     schema.TypeInt,
// 				Required: true,
// 			},
// 			"context_root": &schema.Schema{
// 				Type:     schema.TypeString,
// 				Required: true,
// 			},
// 			"default_auth_type": &schema.Schema{
// 				Type:     schema.TypeString,
// 				Required: true,
// 			},
// 			name
// 			siteId
// 			VirtualHostIds
// 			// AccessValidatorID  int                     `json:"accessValidatorId"`
// 			// AgentID            int                     `json:"agentId"`
// 			// ApplicationType    string                  `json:"applicationType"`
// 			// CaseSensitivePath  bool                    `json:"caseSensitivePath"`
// 			// ContextRoot        string                  `json:"contextRoot"`
// 			// DefaultAuthType    string                  `json:"defaultAuthType"`
// 			// Description        string                  `json:"description"`
// 			// Destination        string                  `json:"destination"`
// 			// Enabled            bool                    `json:"enabled"`
// 			// IdentityMappingIds map[string]int          `json:"identityMappingIds"`
// 			// Name               string                  `json:"name"`
// 			// Policy             map[string][]PolicyItem `json:"policy"`
// 			// Realm              string                  `json:"realm"`
// 			// RequireHTTPS       bool                    `json:"requireHTTPS"`
// 			// SiteID             int                     `json:"siteId"`
// 			// VirtualHostIds     []int                   `json:"virtualHostIds"`
// 			// WebSessionID       int                     `json:"webSessionId"`
// 		},
// 	}
// }

// func resourcePingAccessApplicationCreate(d *schema.ResourceData, m interface{}) error {
// 	name := d.Get("name").(string)
// 	className := d.Get("class_name").(string)
// 	supportedDestinations := d.Get("supported_destinations").(*schema.Set).List()
// 	configuration := make(map[string]interface{})
// 	if kv, ok := d.GetOk("configuration"); ok {
// 		for k, v := range kv.(map[string]interface{}) {
// 			configuration[k] = v.(interface{})
// 		}
// 	}

// 	app := pingaccess.ApplicationView{
// 		Name:                  name,
// 		ClassName:             className,
// 		SupportedDestinations: supportedDestinations,
// 		Configuration:         configuration,
// 	}

// 	b, err := json.Marshal(app)
// 	log.Printf("LOGGING 1")
// 	log.Printf("[LOGGING] %s", b)
// 	res, err := m.(*pingaccess.Client).CreateApplication(app)
// 	b, err = json.Marshal(res)
// 	log.Printf("[LOGGING] %s", b)
// 	if err != nil {
// 		return fmt.Errorf("Error creating rule: %s", err)
// 	}

// 	log.Printf("LOGGING 2")
// 	d.SetId(strconv.Itoa(res.ID))
// 	// res, err := m.(*pingaccess.Client).GetApplications()
// 	// if err != nil {
// 	// 	// log.Printf("[INFO] all dead")
// 	// }
// 	// b, err := json.Marshal(res)
// 	// log.Printf("[INFO] %s", b)
// 	// d.SetId(name)
// 	return resourcePingAccessApplicationReadResult(d, &rv)
// }

// func resourcePingAccessApplicationRead(d *schema.ResourceData, m interface{}) error {

// 	return nil
// }

// func resourcePingAccessApplicationUpdate(d *schema.ResourceData, m interface{}) error {
// 	return resourcePingAccessApplicationRead(d, m)
// }

// func resourcePingAccessApplicationDelete(d *schema.ResourceData, m interface{}) error {
// 	return nil
// }

// func resourcePingAccessApplicationReadResult(d *schema.ResourceData, rv *pingaccess.ApplicationView) error {
// 	if err := d.Set("name", rv.Name); err != nil {
// 		return err
// 	}
// 	if err := d.Set("class_name", rv.ClassName); err != nil {
// 		return err
// 	}
// 	if err := d.Set("supported_destinations", rv.SupportedDestinations); err != nil {
// 		return err
// 	}
// 	// if err := d.Set("configuration", rv.Configuration); err != nil {
// 	// 	return err
// 	// }
// 	return nil
// }
