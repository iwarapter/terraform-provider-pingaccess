package pingaccess

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/iwarapter/pingaccess-sdk-go/pingaccess"
)

func resourcePingAccessAccessTokenValidator() *schema.Resource {
	return &schema.Resource{
		Create: resourcePingAccessAccessTokenValidatorCreate,
		Read:   resourcePingAccessAccessTokenValidatorRead,
		Update: resourcePingAccessAccessTokenValidatorUpdate,
		Delete: resourcePingAccessAccessTokenValidatorDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: resourcePingAccessAccessTokenValidatorSchema(),
	}
}

func resourcePingAccessAccessTokenValidatorSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
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
	}
}

func resourcePingAccessAccessTokenValidatorCreate(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pingaccess.Client).AccessTokenValidators
	input := pingaccess.AddAccessTokenValidatorCommandInput{
		Body: *resourcePingAccessAccessTokenValidatorReadData(d),
	}

	result, _, err := svc.AddAccessTokenValidatorCommand(&input)
	if err != nil {
		return fmt.Errorf("Error creating AccessTokenValidator: %s", err)
	}

	d.SetId(result.Id.String())
	return resourcePingAccessAccessTokenValidatorReadResult(d, result, svc)
}

func resourcePingAccessAccessTokenValidatorRead(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pingaccess.Client).AccessTokenValidators

	input := &pingaccess.GetAccessTokenValidatorCommandInput{
		Id: d.Id(),
	}

	result, _, err := svc.GetAccessTokenValidatorCommand(input)
	if err != nil {
		return fmt.Errorf("Error reading AccessTokenValidator: %s", err)
	}

	return resourcePingAccessAccessTokenValidatorReadResult(d, result, svc)
}

func resourcePingAccessAccessTokenValidatorUpdate(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pingaccess.Client).AccessTokenValidators
	input := pingaccess.UpdateAccessTokenValidatorCommandInput{
		Body: *resourcePingAccessAccessTokenValidatorReadData(d),
		Id: d.Id(),
	}

	result, _, err := svc.UpdateAccessTokenValidatorCommand(&input)
	if err != nil {
		return fmt.Errorf("Error creating AccessTokenValidator: %s", err)
	}

	d.SetId(result.Id.String())
	return resourcePingAccessAccessTokenValidatorReadResult(d, result, svc)
}

func resourcePingAccessAccessTokenValidatorDelete(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pingaccess.Client).AccessTokenValidators
	_, err := svc.DeleteAccessTokenValidatorCommand(&pingaccess.DeleteAccessTokenValidatorCommandInput{Id: d.Id()})
	if err != nil {
		return fmt.Errorf("Error deleting AccessTokenValidator: %s", err)
	}
	return nil
}

func resourcePingAccessAccessTokenValidatorReadResult(d *schema.ResourceData, input *pingaccess.AccessTokenValidatorView, svc *pingaccess.AccessTokenValidatorsService) error {
	setResourceDataString(d, "name", input.Name)
	setResourceDataString(d, "class_name", input.ClassName)
	b, _ := json.Marshal(input.Configuration)
	config := string(b)

	originalConfig := d.Get("configuration").(string)

	//Search the Access Token Validators descriptors for CONCEALED fields, and update the original value back as we cannot use the
	//encryptedValue provided by the API, whilst this gives us a stable plan - we cannot determine if a CONCEALED value
	//has changed and needs updating
	desc, _, _ := svc.GetAccessTokenValidatorDescriptorsCommand()
	config = maskConfigFromDescriptors(desc, input.ClassName, originalConfig, config)

	if err := d.Set("configuration", config); err != nil {
		return err
	}
	return nil
}

func resourcePingAccessAccessTokenValidatorReadData(d *schema.ResourceData) *pingaccess.AccessTokenValidatorView {
	config := d.Get("configuration").(string)
	var dat map[string]interface{}
	_ = json.Unmarshal([]byte(config), &dat)
	atv := &pingaccess.AccessTokenValidatorView{
		Name:          String(d.Get("name").(string)),
		ClassName:     String(d.Get("class_name").(string)),
		Configuration: dat,
	}
	return atv
}