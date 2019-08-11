package pingaccess

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
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
		"class_name": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"name": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"configuration": &schema.Schema{
			Type:             schema.TypeString,
			Required:         true,
			DiffSuppressFunc: suppressEquivalentConfigurationDiffs,
		},
	}
}

func resourcePingAccessAccessTokenValidatorCreate(d *schema.ResourceData, m interface{}) error {
	name := d.Get(name).(string)
	className := d.Get(className).(string)
	config := d.Get(configuration).(string)
	var dat map[string]interface{}
	_ = json.Unmarshal([]byte(config), &dat)

	input := pingaccess.AddAccessTokenValidatorCommandInput{
		Body: pingaccess.AccessTokenValidatorView{
			Name:          String(name),
			ClassName:     String(className),
			Configuration: dat,
		},
	}

	svc := m.(*pingaccess.Client).AccessTokenValidators

	result, _, err := svc.AddAccessTokenValidatorCommand(&input)
	if err != nil {
		return fmt.Errorf("Error creating AccessTokenValidator: %s", err)
	}

	d.SetId(result.Id.String())
	return resourcePingAccessAccessTokenValidatorReadResult(d, result)
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

	return resourcePingAccessAccessTokenValidatorReadResult(d, result)
}

func resourcePingAccessAccessTokenValidatorUpdate(d *schema.ResourceData, m interface{}) error {
	name := d.Get(name).(string)
	className := d.Get(className).(string)
	config := d.Get(configuration).(string)
	var dat map[string]interface{}
	_ = json.Unmarshal([]byte(config), &dat)

	input := pingaccess.UpdateAccessTokenValidatorCommandInput{
		Body: pingaccess.AccessTokenValidatorView{
			Name:          String(name),
			ClassName:     String(className),
			Configuration: dat,
		},
		Id: d.Id(),
	}

	svc := m.(*pingaccess.Client).AccessTokenValidators

	result, _, err := svc.UpdateAccessTokenValidatorCommand(&input)
	if err != nil {
		return fmt.Errorf("Error creating AccessTokenValidator: %s", err)
	}

	d.SetId(result.Id.String())
	return resourcePingAccessAccessTokenValidatorReadResult(d, result)
}

func resourcePingAccessAccessTokenValidatorDelete(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pingaccess.Client).AccessTokenValidators
	_, err := svc.DeleteAccessTokenValidatorCommand(&pingaccess.DeleteAccessTokenValidatorCommandInput{Id: d.Id()})
	if err != nil {
		return fmt.Errorf("Error deleting AccessTokenValidator: %s", err)
	}
	return nil
}

func resourcePingAccessAccessTokenValidatorReadResult(d *schema.ResourceData, rv *pingaccess.AccessTokenValidatorView) error {
	if err := d.Set("name", rv.Name); err != nil {
		return err
	}
	if err := d.Set("class_name", rv.ClassName); err != nil {
		return err
	}
	b, _ := json.Marshal(rv.Configuration)
	if err := d.Set("configuration", string(b)); err != nil {
		return err
	}
	return nil
}
