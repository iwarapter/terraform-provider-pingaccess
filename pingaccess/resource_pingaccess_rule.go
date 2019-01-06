package pingaccess

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/iwarapter/pingaccess-sdk-go/pingaccess"
)

func resourcePingAccessRule() *schema.Resource {
	setOfString := &schema.Schema{
		Type:     schema.TypeSet,
		Optional: true,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	}

	return &schema.Resource{
		Create: resourcePingAccessRuleCreate,
		Read:   resourcePingAccessRuleRead,
		Update: resourcePingAccessRuleUpdate,
		Delete: resourcePingAccessRuleDelete,

		Schema: map[string]*schema.Schema{
			className: &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			name: &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			supportedDestinations: setOfString,
			configuration: &schema.Schema{
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: suppressEquivalentConfigurationDiffs,
			},
		},
	}
}

func resourcePingAccessRuleCreate(d *schema.ResourceData, m interface{}) error {
	log.Printf("[INFO] resourcePingAccessRuleCreate")
	name := d.Get(name).(string)
	className := d.Get(className).(string)
	supDests := expandStringList(d.Get(supportedDestinations).(*schema.Set).List())
	config := d.Get(configuration).(string)
	var dat map[string]interface{}
	_ = json.Unmarshal([]byte(config), &dat)

	input := pingaccess.AddRuleCommandInput{
		Body: pingaccess.RuleView{
			Name:                  String(name),
			ClassName:             String(className),
			SupportedDestinations: &supDests,
			Configuration:         dat,
		},
	}

	svc := m.(*pingaccess.Client).Rules

	result, _, err := svc.AddRuleCommand(&input)
	if err != nil {
		return fmt.Errorf("Error creating rule: %s", err)
	}

	d.SetId(result.Id.String())
	return resourcePingAccessRuleReadResult(d, result)
}

func resourcePingAccessRuleRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("[INFO] resourcePingAccessRuleRead")
	svc := m.(*pingaccess.Client).Rules

	input := &pingaccess.GetRuleCommandInput{
		Id: d.Id(),
	}

	log.Printf("[INFO] ResourceID: %s", d.Id())
	log.Printf("[INFO] GetApplicationCommandInput: %s", input.Id)
	result, _, _ := svc.GetRuleCommand(input)
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(result)
	rv := pingaccess.RuleView{}
	json.NewDecoder(b).Decode(&rv)

	return resourcePingAccessRuleReadResult(d, &rv)
}

func resourcePingAccessRuleUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("[INFO] resourcePingAccessRuleUpdate")
	name := d.Get(name).(string)
	className := d.Get(className).(string)
	supDests := expandStringList(d.Get(supportedDestinations).(*schema.Set).List())
	config := d.Get(configuration).(string)
	var dat map[string]interface{}
	_ = json.Unmarshal([]byte(config), &dat)

	input := pingaccess.UpdateRuleCommandInput{
		Body: pingaccess.RuleView{
			Name:                  String(name),
			ClassName:             String(className),
			SupportedDestinations: &supDests,
			Configuration:         dat,
		},
		Id: d.Id(),
	}

	svc := m.(*pingaccess.Client).Rules

	result, _, err := svc.UpdateRuleCommand(&input)
	if err != nil {
		return fmt.Errorf("Error creating rule: %s", err)
	}

	d.SetId(result.Id.String())
	log.Println("[DEBUG] End - resourcePingAccessRuleUpdate")
	return resourcePingAccessRuleReadResult(d, result)
}

func resourcePingAccessRuleDelete(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] Start - resourcePingAccessRuleDelete")
	svc := m.(*pingaccess.Client).Rules
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	log.Printf("[INFO] ResourceID: %s", d.Id())
	_, err := svc.DeleteRuleCommand(&pingaccess.DeleteRuleCommandInput{Id: d.Id()})
	if err != nil {
		return fmt.Errorf("Error deleting virtualhost: %s", err)
	}
	log.Println("[DEBUG] End - resourcePingAccessRuleDelete")
	return nil
}

func resourcePingAccessRuleReadResult(d *schema.ResourceData, rv *pingaccess.RuleView) error {
	if err := d.Set("name", rv.Name); err != nil {
		return err
	}
	if err := d.Set("class_name", rv.ClassName); err != nil {
		return err
	}
	if err := d.Set("supported_destinations", rv.SupportedDestinations); err != nil {
		return err
	}
	b, _ := json.Marshal(rv.Configuration)
	if err := d.Set("configuration", string(b)); err != nil {
		return err
	}
	return nil
}

func suppressEquivalentConfigurationDiffs(k, old, new string, d *schema.ResourceData) bool {
	var o1 interface{}
	var o2 interface{}
	_ = json.Unmarshal([]byte(old), &o1)
	_ = json.Unmarshal([]byte(new), &o2)

	return reflect.DeepEqual(o1, o2)
}
