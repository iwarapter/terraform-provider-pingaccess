package pingaccess

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/iwarapter/pingaccess-sdk-go/pingaccess"
)

func resourcePingAccessLoadBalancingStrategy() *schema.Resource {
	return &schema.Resource{
		Create: resourcePingAccessLoadBalancingStrategyCreate,
		Read:   resourcePingAccessLoadBalancingStrategyRead,
		Update: resourcePingAccessLoadBalancingStrategyUpdate,
		Delete: resourcePingAccessLoadBalancingStrategyDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: resourcePingAccessLoadBalancingStrategySchema(),
	}
}

func resourcePingAccessLoadBalancingStrategySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		className: &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		configuration: &schema.Schema{
			Type:             schema.TypeString,
			Required:         true,
			DiffSuppressFunc: suppressEquivalentConfigurationDiffs,
		},
		name: &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
	}
}

func resourcePingAccessLoadBalancingStrategyCreate(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pingaccess.Client).HighAvailability
	input := pingaccess.AddLoadBalancingStrategyCommandInput{
		Body: *resourcePingAccessLoadBalancingStrategyReadData(d),
	}
	result, _, err := svc.AddLoadBalancingStrategyCommand(&input)
	if err != nil {
		return fmt.Errorf("Error creating LoadBalancerStrategy: %s", err)
	}
	d.SetId(result.Id.String())
	return resourcePingAccessLoadBalancingStrategyReadResult(d, &input.Body)
}

func resourcePingAccessLoadBalancingStrategyRead(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pingaccess.Client).HighAvailability
	input := &pingaccess.GetLoadBalancingStrategyCommandInput{
		Id: d.Id(),
	}
	result, _, err := svc.GetLoadBalancingStrategyCommand(input)
	if err != nil {
		return fmt.Errorf("Error reading Load Balancing Strategy: %s", err)
	}
	return resourcePingAccessLoadBalancingStrategyReadResult(d, result)
}

func resourcePingAccessLoadBalancingStrategyUpdate(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pingaccess.Client).HighAvailability
	input := pingaccess.UpdateLoadBalancingStrategyCommandInput{
		Body: *resourcePingAccessLoadBalancingStrategyReadData(d),
		Id:   d.Id(),
	}

	result, _, err := svc.UpdateLoadBalancingStrategyCommand(&input)
	if err != nil {
		return fmt.Errorf("Error updating Load Balancing Strategy: %s", err.Error())
	}
	return resourcePingAccessLoadBalancingStrategyReadResult(d, result)
}

func resourcePingAccessLoadBalancingStrategyDelete(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pingaccess.Client).HighAvailability
	input := &pingaccess.DeleteLoadBalancingStrategyCommandInput{
		Id: d.Id(),
	}

	_, err := svc.DeleteLoadBalancingStrategyCommand(input)
	if err != nil {
		return fmt.Errorf("Error deleting Load Balancing Strategy: %s", err)
	}
	return nil
}

func resourcePingAccessLoadBalancingStrategyReadResult(d *schema.ResourceData, input *pingaccess.LoadBalancingStrategyView) error {
	setResourceDataString(d, "class_name", input.ClassName)
	b, _ := json.Marshal(input.Configuration)
	if err := d.Set("configuration", string(b)); err != nil {
		return err
	}
	setResourceDataString(d, "name", input.Name)
	return nil
}

func resourcePingAccessLoadBalancingStrategyReadData(d *schema.ResourceData) *pingaccess.LoadBalancingStrategyView {
	config := d.Get(configuration).(string)
	var dat map[string]interface{}
	_ = json.Unmarshal([]byte(config), &dat)
	lbs := &pingaccess.LoadBalancingStrategyView{
		Name:          String(d.Get(name).(string)),
		ClassName:     String(d.Get(className).(string)),
		Configuration: dat,
	}
	return lbs
}
