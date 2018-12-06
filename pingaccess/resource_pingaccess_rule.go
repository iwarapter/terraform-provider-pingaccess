package pingaccess

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"strconv"

// 	"github.com/hashicorp/terraform/helper/schema"
// 	"github.com/iwarapter/pingaccess-sdk-go/pingaccess"
// )

// func resourcePingAccessRule() *schema.Resource {
// 	setOfString := &schema.Schema{
// 		Type:     schema.TypeSet,
// 		Optional: true,
// 		Elem: &schema.Schema{
// 			Type: schema.TypeString,
// 		},
// 	}

// 	return &schema.Resource{
// 		Create: resourcePingAccessRuleCreate,
// 		Read:   resourcePingAccessRuleRead,
// 		Update: resourcePingAccessRuleUpdate,
// 		Delete: resourcePingAccessRuleDelete,

// 		Schema: map[string]*schema.Schema{
// 			"class_name": &schema.Schema{
// 				Type:     schema.TypeString,
// 				Required: true,
// 			},
// 			"name": &schema.Schema{
// 				Type:     schema.TypeString,
// 				Required: true,
// 			},
// 			"supported_destinations": setOfString,
// 			"configuration": &schema.Schema{
// 				Type:     schema.TypeMap,
// 				Required: true,
// 			},
// 		},
// 	}
// }

// func resourcePingAccessRuleCreate(d *schema.ResourceData, m interface{}) error {
// 	name := d.Get("name").(string)
// 	className := d.Get("class_name").(string)
// 	supportedDestinations := d.Get("supported_destinations").(*schema.Set).List()
// 	configuration := make(map[string]interface{})
// 	if kv, ok := d.GetOk("configuration"); ok {
// 		for k, v := range kv.(map[string]interface{}) {
// 			configuration[k] = v.(interface{})
// 		}
// 	}

// 	rv := pingaccess.RuleView{
// 		Name:                  name,
// 		ClassName:             className,
// 		SupportedDestinations: supportedDestinations,
// 		Configuration:         configuration,
// 	}

// 	b, err := json.Marshal(rv)
// 	log.Printf("LOGGING 1")
// 	log.Printf("[LOGGING] %s", b)
// 	res, err := m.(*pingaccess.Client).CreateRule(rv)
// 	b, err = json.Marshal(res)
// 	log.Printf("[LOGGING] %s", b)
// 	if err != nil {
// 		return fmt.Errorf("Error creating rule: %s", err)
// 	}

// 	log.Printf("LOGGING 2")
// 	d.SetId(strconv.Itoa(res.ID))
// 	// res, err := m.(*pingaccess.Client).GetRules()
// 	// if err != nil {
// 	// 	// log.Printf("[INFO] all dead")
// 	// }
// 	// b, err := json.Marshal(res)
// 	// log.Printf("[INFO] %s", b)
// 	// d.SetId(name)
// 	return resourcePingAccessRuleReadResult(d, &rv)
// }

// func resourcePingAccessRuleRead(d *schema.ResourceData, m interface{}) error {

// 	return nil
// }

// func resourcePingAccessRuleUpdate(d *schema.ResourceData, m interface{}) error {
// 	return resourcePingAccessRuleRead(d, m)
// }

// func resourcePingAccessRuleDelete(d *schema.ResourceData, m interface{}) error {
// 	return nil
// }

// func resourcePingAccessRuleReadResult(d *schema.ResourceData, rv *pingaccess.RuleView) error {
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
