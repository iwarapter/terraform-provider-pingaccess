package pingaccess

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/iwarapter/pingaccess-sdk-go/pingaccess"
)

func resourcePingAccessHTTPSListener() *schema.Resource {
	return &schema.Resource{
		Create: resourcePingAccessHTTPSListenerCreate,
		Read:   resourcePingAccessHTTPSListenerRead,
		Update: resourcePingAccessHTTPSListenerUpdate,
		Delete: resourcePingAccessHTTPSListenerDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: resourcePingAccessHTTPSListenerSchema(),
	}
}

func resourcePingAccessHTTPSListenerSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": &schema.Schema{
			Type:         schema.TypeString,
			Required:     true,
			ValidateFunc: validateHTTPListenerName,
			ForceNew:     true,
		},
		"key_pair_id": &schema.Schema{
			Type:     schema.TypeInt,
			Required: true,
		},
		"use_server_cipher_suite_order": &schema.Schema{
			Type:     schema.TypeBool,
			Required: true,
		},
	}
}

func resourcePingAccessHTTPSListenerCreate(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pingaccess.Client).HttpsListeners
	input := pingaccess.GetHttpsListenersCommandInput{}

	result, _, err := svc.GetHttpsListenersCommand(&input)
	if err != nil {
		return fmt.Errorf("Error retrieving HttpsListener list: %s", err)
	}

	name := d.Get("name").(string)
	for _, listener := range result.Items {
		if *listener.Name == name {
			d.SetId(listener.Id.String())
			return resourcePingAccessHTTPSListenerReadResult(d, listener)
		}
	}
	return fmt.Errorf("Error managing HttpsListener: %s", name)
}

func resourcePingAccessHTTPSListenerRead(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pingaccess.Client).HttpsListeners
	input := &pingaccess.GetHttpsListenerCommandInput{
		Id: d.Id(),
	}
	result, _, err := svc.GetHttpsListenerCommand(input)
	if err != nil {
		return fmt.Errorf("Error reading HttpsListener: %s", err)
	}
	return resourcePingAccessHTTPSListenerReadResult(d, result)
}

func resourcePingAccessHTTPSListenerUpdate(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pingaccess.Client).HttpsListeners
	input := pingaccess.UpdateHttpsListenerInput{
		Body: *resourcePingAccessHTTPSListenerReadData(d),
		Id:   d.Id(),
	}

	result, _, err := svc.UpdateHttpsListener(&input)
	if err != nil {
		return fmt.Errorf("Error updating HttpsListener: %s", err.Error())
	}
	return resourcePingAccessHTTPSListenerReadResult(d, result)
}

func resourcePingAccessHTTPSListenerDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourcePingAccessHTTPSListenerReadResult(d *schema.ResourceData, input *pingaccess.HttpsListenerView) error {
	setResourceDataString(d, "name", input.Name)
	setResourceDataInt(d, "key_pair_id", input.KeyPairId)
	setResourceDataBool(d, "use_server_cipher_suite_order", input.UseServerCipherSuiteOrder)
	return nil
}

func resourcePingAccessHTTPSListenerReadData(d *schema.ResourceData) *pingaccess.HttpsListenerView {
	engine := &pingaccess.HttpsListenerView{
		Name:                      String(d.Get("name").(string)),
		KeyPairId:                 Int(d.Get("key_pair_id").(int)),
		UseServerCipherSuiteOrder: Bool(d.Get("use_server_cipher_suite_order").(bool)),
	}

	return engine
}
