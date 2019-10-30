package pingaccess

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/iwarapter/pingaccess-sdk-go/pingaccess"
)

func resourcePingAccessEngineListener() *schema.Resource {
	return &schema.Resource{
		Create: resourcePingAccessEngineListenerCreate,
		Read:   resourcePingAccessEngineListenerRead,
		Update: resourcePingAccessEngineListenerUpdate,
		Delete: resourcePingAccessEngineListenerDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: resourcePingAccessEngineListenerSchema(),
	}
}

func resourcePingAccessEngineListenerSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"port": &schema.Schema{
			Type:     schema.TypeInt,
			Required: true,
		},
		"secure": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
		},
	}
}

func resourcePingAccessEngineListenerCreate(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pingaccess.Client).EngineListeners
	input := pingaccess.AddEngineListenerCommandInput{
		Body: *resourcePingAccessEngineListenerReadData(d),
	}

	result, _, err := svc.AddEngineListenerCommand(&input)
	if err != nil {
		return fmt.Errorf("Error creating EngineListener: %s", err)
	}

	d.SetId(result.Id.String())
	return resourcePingAccessEngineListenerReadResult(d, &input.Body)
}

func resourcePingAccessEngineListenerRead(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pingaccess.Client).EngineListeners
	input := &pingaccess.GetEngineListenerCommandInput{
		Id: d.Id(),
	}
	result, _, err := svc.GetEngineListenerCommand(input)
	if err != nil {
		return fmt.Errorf("Error reading EngineListener: %s", err)
	}
	return resourcePingAccessEngineListenerReadResult(d, result)
}

func resourcePingAccessEngineListenerUpdate(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pingaccess.Client).EngineListeners
	input := pingaccess.UpdateEngineListenerCommandInput{
		Body: *resourcePingAccessEngineListenerReadData(d),
		Id:   d.Id(),
	}

	result, _, err := svc.UpdateEngineListenerCommand(&input)
	if err != nil {
		return fmt.Errorf("Error updating EngineListener: %s", err.Error())
	}
	return resourcePingAccessEngineListenerReadResult(d, result)
}

func resourcePingAccessEngineListenerDelete(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pingaccess.Client).EngineListeners
	input := &pingaccess.DeleteEngineListenerCommandInput{
		Id: d.Id(),
	}

	_, err := svc.DeleteEngineListenerCommand(input)
	if err != nil {
		return fmt.Errorf("Error deleting EngineListener: %s", err)
	}
	return nil
}

func resourcePingAccessEngineListenerReadResult(d *schema.ResourceData, input *pingaccess.EngineListenerView) error {
	setResourceDataString(d, "name", input.Name)
	setResourceDataInt(d, "port", input.Port)
	setResourceDataBool(d, "secure", input.Secure)
	return nil
}

func resourcePingAccessEngineListenerReadData(d *schema.ResourceData) *pingaccess.EngineListenerView {
	engine := &pingaccess.EngineListenerView{
		Name: String(d.Get("name").(string)),
		Port: Int(d.Get("port").(int)),
	}

	if v, ok := d.GetOkExists("secure"); ok {
		engine.Secure = Bool(v.(bool))
	}

	return engine
}
