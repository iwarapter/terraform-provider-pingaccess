package pingaccess

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/iwarapter/pingaccess-sdk-go/pingaccess"
)

func resourcePingAccessIdentityMapping() *schema.Resource {
	return &schema.Resource{
		Create: resourcePingAccessIdentityMappingCreate,
		Read:   resourcePingAccessIdentityMappingRead,
		Update: resourcePingAccessIdentityMappingUpdate,
		Delete: resourcePingAccessIdentityMappingDelete,
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
			"configuration": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: suppressEquivalentJsonDiffs,
			},
		},
	}
}

func resourcePingAccessIdentityMappingCreate(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pingaccess.Client).IdentityMappings
	input := pingaccess.AddIdentityMappingCommandInput{
		Body: *resourcePingAccessIdentityMappingReadData(d),
	}

	result, _, err := svc.AddIdentityMappingCommand(&input)
	if err != nil {
		return fmt.Errorf("Error creating IdentityMapping: %s", err)
	}

	d.SetId(result.Id.String())
	return resourcePingAccessIdentityMappingReadResult(d, result, svc)
}

func resourcePingAccessIdentityMappingRead(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pingaccess.Client).IdentityMappings
	input := &pingaccess.GetIdentityMappingCommandInput{
		Id: d.Id(),
	}
	result, _, err := svc.GetIdentityMappingCommand(input)
	if err != nil {
		return fmt.Errorf("Error reading IdentityMapping: %s", err)
	}
	return resourcePingAccessIdentityMappingReadResult(d, result, svc)
}

func resourcePingAccessIdentityMappingUpdate(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pingaccess.Client).IdentityMappings
	input := pingaccess.UpdateIdentityMappingCommandInput{
		Body: *resourcePingAccessIdentityMappingReadData(d),
		Id: d.Id(),
	}

	result, _, err := svc.UpdateIdentityMappingCommand(&input)
	if err != nil {
		return fmt.Errorf("Error creating IdentityMapping: %s", err)
	}

	d.SetId(result.Id.String())
	return resourcePingAccessIdentityMappingReadResult(d, result, svc)
}

func resourcePingAccessIdentityMappingDelete(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pingaccess.Client).IdentityMappings
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	log.Printf("[INFO] ResourceID: %s", d.Id())
	_, err := svc.DeleteIdentityMappingCommand(&pingaccess.DeleteIdentityMappingCommandInput{Id: d.Id()})
	if err != nil {
		return fmt.Errorf("Error deleting IdentityMapping: %s", err)
	}
	return nil
}

func resourcePingAccessIdentityMappingReadResult(d *schema.ResourceData, input *pingaccess.IdentityMappingView, svc *pingaccess.IdentityMappingsService) error {
	setResourceDataString(d, "name", input.Name)
	setResourceDataString(d, "class_name", input.ClassName)

	b, _ := json.Marshal(input.Configuration)
	config := string(b)

	originalConfig := d.Get("configuration").(string)

	//Search the Identity Mappings descriptors for CONCEALED fields, and update the original value back as we cannot use the
	//encryptedValue provided by the API, whilst this gives us a stable plan - we cannot determine if a CONCEALED value
	//has changed and needs updating
	desc, _, _ := svc.GetIdentityMappingDescriptorsCommand()
	config = maskConfigFromDescriptors(desc, input.ClassName, originalConfig, config)

	if err := d.Set("configuration", config); err != nil {
		return err
	}
	return nil
}

func resourcePingAccessIdentityMappingReadData(d *schema.ResourceData) *pingaccess.IdentityMappingView {
	config := d.Get("configuration").(string)
	var dat map[string]interface{}
	_ = json.Unmarshal([]byte(config), &dat)
	idMapping := &pingaccess.IdentityMappingView{
		Name:          String(d.Get("name").(string)),
		ClassName:     String(d.Get("class_name").(string)),
		Configuration: dat,
	}
	return idMapping
}