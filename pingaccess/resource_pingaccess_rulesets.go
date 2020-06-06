package pingaccess

import (
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/iwarapter/pingaccess-sdk-go/pingaccess"
)

func resourcePingAccessRuleSet() *schema.Resource {
	return &schema.Resource{
		Create: resourcePingAccessRuleSetCreate,
		Read:   resourcePingAccessRuleSetRead,
		Update: resourcePingAccessRuleSetUpdate,
		Delete: resourcePingAccessRuleSetDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"element_type": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateRuleOrRuleSet,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"policy": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true,
				MinItems: 1,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"success_criteria": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateSuccessIfAllSucceedOrSuccessIfAnyOneSucceeds,
			},
		},
	}
}

func resourcePingAccessRuleSetCreate(d *schema.ResourceData, m interface{}) error {
	// log.Printf("[INFO] resourcePingAccessRuleSetCreate")
	name := d.Get("name").(string)
	elementType := d.Get("element_type").(string)
	policy := expandStringList(d.Get("policy").(*schema.Set).List())
	successCriteria := d.Get("success_criteria").(string)

	//TODO generalise this into a helper function
	pol_ids := []*int{}
	for _, i := range policy {
		text, _ := strconv.Atoi(*i)
		pol_ids = append(pol_ids, &text)
	}

	input := pingaccess.AddRuleSetCommandInput{
		Body: pingaccess.RuleSetView{
			Name:            String(name),
			ElementType:     String(elementType),
			Policy:          &pol_ids,
			SuccessCriteria: String(successCriteria),
		},
	}

	svc := m.(*pingaccess.Client).Rulesets

	result, _, err := svc.AddRuleSetCommand(&input)
	if err != nil {
		return fmt.Errorf("Error creating RuleSet: %s", err)
	}

	d.SetId(result.Id.String())
	return resourcePingAccessRuleSetReadResult(d, result)
}

func resourcePingAccessRuleSetRead(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pingaccess.Client).Rulesets
	result, _, err := svc.GetRuleSetCommand(&pingaccess.GetRuleSetCommandInput{
		Id: d.Id(),
	})
	if err != nil {
		return fmt.Errorf("Error reading RuleSet: %s", err)
	}
	return resourcePingAccessRuleSetReadResult(d, result)
}

func resourcePingAccessRuleSetUpdate(d *schema.ResourceData, m interface{}) error {
	// log.Printf("[INFO] resourcePingAccessRuleSetUpdate")
	name := d.Get("name").(string)
	elementType := d.Get("element_type").(string)
	policy := expandStringList(d.Get("policy").(*schema.Set).List())
	successCriteria := d.Get("success_criteria").(string)

	//TODO generalise this into a helper function
	pol_ids := []*int{}
	for _, i := range policy {
		text, _ := strconv.Atoi(*i)
		pol_ids = append(pol_ids, &text)
	}

	input := pingaccess.UpdateRuleSetCommandInput{
		Body: pingaccess.RuleSetView{
			Name:            String(name),
			ElementType:     String(elementType),
			Policy:          &pol_ids,
			SuccessCriteria: String(successCriteria),
		},
		Id: d.Id(),
	}

	svc := m.(*pingaccess.Client).Rulesets

	result, _, err := svc.UpdateRuleSetCommand(&input)
	if err != nil {
		return fmt.Errorf("Error updating RuleSet: %s", err)
	}
	d.SetId(result.Id.String())
	return resourcePingAccessRuleSetReadResult(d, result)
}

func resourcePingAccessRuleSetDelete(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pingaccess.Client).Rulesets

	_, err := svc.DeleteRuleSetCommand(&pingaccess.DeleteRuleSetCommandInput{Id: d.Id()})
	if err != nil {
		return fmt.Errorf("Error deleting RuleSet: %s", err)
	}
	return nil
}

func resourcePingAccessRuleSetReadResult(d *schema.ResourceData, rv *pingaccess.RuleSetView) error {
	if err := d.Set("name", rv.Name); err != nil {
		return err
	}
	if err := d.Set("success_criteria", rv.SuccessCriteria); err != nil {
		return err
	}
	if err := d.Set("element_type", rv.ElementType); err != nil {
		return err
	}

	if rv.Policy != nil {
		pol_ids := []string{}
		for _, p := range *rv.Policy {
			pol_ids = append(pol_ids, strconv.Itoa(*p))
		}
		if err := d.Set("policy", pol_ids); err != nil {
			return err
		}
	}
	return nil
}
