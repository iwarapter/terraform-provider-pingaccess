package pingaccess

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/iwarapter/pingaccess-sdk-go/pingaccess"
	"github.com/tidwall/sjson"
)

func resourcePingAccessRule() *schema.Resource {
	return &schema.Resource{
		Create: resourcePingAccessRuleCreate,
		Read:   resourcePingAccessRuleRead,
		Update: resourcePingAccessRuleUpdate,
		Delete: resourcePingAccessRuleDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			className: &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			name: &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			supportedDestinations: setOfString(),
			configuration: &schema.Schema{
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: suppressEquivalentConfigurationDiffs,
			},
			"ignrored_configuration_fields": setOfString(),
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

	result, _, err := svc.GetRuleCommand(input)
	if err != nil {
		return fmt.Errorf("Error reading rule: %s", err)
	}

	return resourcePingAccessRuleReadResult(d, result)
}

func resourcePingAccessRuleUpdate(d *schema.ResourceData, m interface{}) error {
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
	return resourcePingAccessRuleReadResult(d, result)
}

func resourcePingAccessRuleDelete(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pingaccess.Client).Rules
	_, err := svc.DeleteRuleCommand(&pingaccess.DeleteRuleCommandInput{Id: d.Id()})
	if err != nil {
		return fmt.Errorf("Error deleting rule: %s", err)
	}
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
	if _, ok := d.GetOkExists("ignrored_configuration_fields"); ok {
		for _, f := range expandStringList(d.Get("ignrored_configuration_fields").(*schema.Set).List()) {
			old, _ = sjson.Delete(old, *f)
			new, _ = sjson.Delete(new, *f)
		}
	}
	var o1 interface{}
	var o2 interface{}
	_ = json.Unmarshal([]byte(old), &o1)
	_ = json.Unmarshal([]byte(new), &o2)
	return reflect.DeepEqual(o1, o2)
}
