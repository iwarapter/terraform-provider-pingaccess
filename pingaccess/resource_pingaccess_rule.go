package pingaccess

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
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
			"class_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"supported_destinations": setOfString(),
			"configuration": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: suppressEquivalentJsonDiffs,
			},
		},
		CustomizeDiff: customdiff.ComputedIf("configuration", func(diff *schema.ResourceDiff, meta interface{}) bool {
			return diff.HasChange("configuration")
		}),
	}
}

func resourcePingAccessRuleCreate(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pingaccess.Client).Rules
	input := pingaccess.AddRuleCommandInput{
		Body: *resourcePingAccessRuleReadData(d),
	}

	result, _, err := svc.AddRuleCommand(&input)
	if err != nil {
		return fmt.Errorf("Error creating rule: %s", err)
	}

	d.SetId(result.Id.String())
	return resourcePingAccessRuleReadResult(d, result, svc)
}

func resourcePingAccessRuleRead(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pingaccess.Client).Rules
	input := &pingaccess.GetRuleCommandInput{
		Id: d.Id(),
	}

	result, _, err := svc.GetRuleCommand(input)
	if err != nil {
		return fmt.Errorf("Error reading rule: %s", err)
	}

	return resourcePingAccessRuleReadResult(d, result, svc)
}

func resourcePingAccessRuleUpdate(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pingaccess.Client).Rules
	input := pingaccess.UpdateRuleCommandInput{
		Body: *resourcePingAccessRuleReadData(d),
		Id: d.Id(),
	}

	result, _, err := svc.UpdateRuleCommand(&input)
	if err != nil {
		return fmt.Errorf("Error updating rule: %s", err)
	}

	d.SetId(result.Id.String())
	return resourcePingAccessRuleReadResult(d, result, svc)
}

func resourcePingAccessRuleDelete(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pingaccess.Client).Rules
	_, err := svc.DeleteRuleCommand(&pingaccess.DeleteRuleCommandInput{Id: d.Id()})
	if err != nil {
		return fmt.Errorf("Error deleting rule: %s", err)
	}
	return nil
}

func resourcePingAccessRuleReadResult(d *schema.ResourceData, input *pingaccess.RuleView, svc *pingaccess.RulesService) error {
	setResourceDataString(d, "name", input.Name)
	setResourceDataString(d, "class_name", input.ClassName)
	if err := d.Set("supported_destinations", input.SupportedDestinations); err != nil {
		return err
	}
	b, _ := json.Marshal(input.Configuration)
	config := string(b)

	originalConfig := d.Get("configuration").(string)

	//Search the Rules descriptors for CONCEALED fields, and update the original value back as we cannot use the
	//encryptedValue provided by the API, whilst this gives us a stable plan - we cannot determine if a CONCEALED value
	//has changed and needs updating
	desc, _, _ := svc.GetRuleDescriptorsCommand()
	config = maskConfigFromRuleDescriptors(desc, input.ClassName, originalConfig, config)

	if err := d.Set("configuration", config); err != nil {
		return err
	}
	return nil
}

func resourcePingAccessRuleReadData(d *schema.ResourceData) *pingaccess.RuleView {
	config := d.Get("configuration").(string)
	var dat map[string]interface{}
	_ = json.Unmarshal([]byte(config), &dat)
	supDests := expandStringList(d.Get("supported_destinations").(*schema.Set).List())
	rule := &pingaccess.RuleView{
		Name:          String(d.Get("name").(string)),
		ClassName:     String(d.Get("class_name").(string)),
		Configuration: dat,
		SupportedDestinations: &supDests,
	}
	return rule
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
