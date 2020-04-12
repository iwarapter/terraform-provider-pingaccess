package pingaccess

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/iwarapter/pingaccess-sdk-go/pingaccess"
	"net/http"
)

func resourcePingAccessHsmProvider() *schema.Resource {
	return &schema.Resource{
		Create: resourcePingAccessHsmProviderCreate,
		Read:   resourcePingAccessHsmProviderRead,
		Update: resourcePingAccessHsmProviderUpdate,
		Delete: resourcePingAccessHsmProviderDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: resourcePingAccessHsmProviderSchema(),
	}
}

func resourcePingAccessHsmProviderSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"class_name": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"name": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"configuration": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
			//DiffSuppressFunc: suppressEquivalentConfigurationDiffs,
		},
		//"hidden_fields": setOfString(),
	}
}

func resourcePingAccessHsmProviderCreate(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pingaccess.Client).HsmProviders
	input := pingaccess.AddHsmProviderCommandInput{
		Body: *resourcePingAccessHsmProviderReadData(d),
	}

	result, _, err := svc.AddHsmProviderCommand(&input)
	if err != nil {
		return fmt.Errorf("Error creating HsmProvider: %s", err.Error())
	}

	d.SetId(result.Id.String())
	//descriptors, _, err := svc.GetHsmProviderDescriptorsCommand()

	//	for _,d := range descriptors.Items {
	//d.ClassName
	//	}

	return resourcePingAccessHsmProviderReadResult(d, result, svc)
}

func resourcePingAccessHsmProviderRead(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pingaccess.Client).HsmProviders
	input := &pingaccess.GetHsmProviderCommandInput{
		Id: d.Id(),
	}
	result, _, err := svc.GetHsmProviderCommand(input)
	if err != nil {
		return fmt.Errorf("Error reading HsmProvider: %s", err.Error())
	}
	return resourcePingAccessHsmProviderReadResult(d, result, svc)
}

func resourcePingAccessHsmProviderUpdate(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pingaccess.Client).HsmProviders
	input := pingaccess.UpdateHsmProviderCommandInput{
		Body: *resourcePingAccessHsmProviderReadData(d),
		Id:   d.Id(),
	}
	result, _, err := svc.UpdateHsmProviderCommand(&input)
	if err != nil {
		return fmt.Errorf("Error updating HsmProvider: %s", err.Error())
	} //d.SetId(result.Id.String())
	return resourcePingAccessHsmProviderReadResult(d, result, svc)
}

func resourcePingAccessHsmProviderDelete(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pingaccess.Client).HsmProviders
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	_, err := svc.DeleteHsmProviderCommand(&pingaccess.DeleteHsmProviderCommandInput{Id: d.Id()})
	if err != nil {
		return fmt.Errorf("Error deleting HsmProvider: %s", err.Error())
	}
	return nil
}

func resourcePingAccessHsmProviderReadResult(d *schema.ResourceData, input *pingaccess.HsmProviderView, svc *pingaccess.HsmProvidersService) error {
	setResourceDataString(d, "name", input.Name)
	setResourceDataString(d, "class_name", input.ClassName)

	b, _ := json.Marshal(input.Configuration)
	config := string(b)

	originalConfig := d.Get("configuration").(string)

	//Search the HSM descriptors for CONCEALED fields, and update the original value back as we cannot use the
	// encryptedValue provided by the API, whilst this gives us a stable plan - we cannot determine if a CONCEALED value
	// has changed and needs updating
	desc, _, _ := svc.GetHsmProviderDescriptorsCommand()

	config = maskConfigFromDescriptors(desc, input.ClassName, originalConfig, config)

	if err := d.Set("configuration", config); err != nil {
		return err
	}
	return nil
}

func resourcePingAccessHsmProviderReadData(d *schema.ResourceData) *pingaccess.HsmProviderView {
	config := d.Get("configuration").(string)
	var dat map[string]interface{}
	_ = json.Unmarshal([]byte(config), &dat)
	hsmProvider := &pingaccess.HsmProviderView{
		Name:          String(d.Get("name").(string)),
		ClassName:     String(d.Get("class_name").(string)),
		Configuration: dat,
	}
	return hsmProvider
}
