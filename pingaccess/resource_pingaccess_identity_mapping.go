package pingaccess

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/iwarapter/pingaccess-sdk-go/pingaccess"
)

func resourcePingAccessIdentityMapping() *schema.Resource {
	return &schema.Resource{
		Create: resourcePingAccessIdentityMappingCreate,
		Read:   resourcePingAccessIdentityMappingRead,
		Update: resourcePingAccessIdentityMappingUpdate,
		Delete: resourcePingAccessIdentityMappingDelete,

		Schema: map[string]*schema.Schema{
			className: &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			name: &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			configuration: &schema.Schema{
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: suppressEquivalentConfigurationDiffs,
			},
		},
	}
}

func resourcePingAccessIdentityMappingCreate(d *schema.ResourceData, m interface{}) error {
	log.Printf("[INFO] resourcePingAccessIdentityMappingCreate")
	name := d.Get(name).(string)
	className := d.Get(className).(string)
	config := d.Get(configuration).(string)
	var dat map[string]interface{}
	_ = json.Unmarshal([]byte(config), &dat)

	input := pingaccess.AddIdentityMappingCommandInput{
		Body: pingaccess.IdentityMappingView{
			Name:          String(name),
			ClassName:     String(className),
			Configuration: dat,
		},
	}

	svc := m.(*pingaccess.Client).IdentityMappings

	result, _, err := svc.AddIdentityMappingCommand(&input)
	if err != nil {
		return fmt.Errorf("Error creating IdentityMapping: %s", err)
	}

	d.SetId(result.Id.String())
	return resourcePingAccessIdentityMappingReadResult(d, result)
}

func resourcePingAccessIdentityMappingRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("[INFO] resourcePingAccessIdentityMappingRead")
	svc := m.(*pingaccess.Client).IdentityMappings

	input := &pingaccess.GetIdentityMappingCommandInput{
		Id: d.Id(),
	}

	log.Printf("[INFO] ResourceID: %s", d.Id())
	log.Printf("[INFO] GetApplicationCommandInput: %s", input.Id)
	result, _, _ := svc.GetIdentityMappingCommand(input)

	return resourcePingAccessIdentityMappingReadResult(d, result)
}

func resourcePingAccessIdentityMappingUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("[INFO] resourcePingAccessIdentityMappingUpdate")
	name := d.Get(name).(string)
	className := d.Get(className).(string)
	config := d.Get(configuration).(string)
	var dat map[string]interface{}
	_ = json.Unmarshal([]byte(config), &dat)

	input := pingaccess.UpdateIdentityMappingCommandInput{
		Body: pingaccess.IdentityMappingView{
			Name:          String(name),
			ClassName:     String(className),
			Configuration: dat,
		},
		Id: d.Id(),
	}

	svc := m.(*pingaccess.Client).IdentityMappings

	result, _, err := svc.UpdateIdentityMappingCommand(&input)
	if err != nil {
		return fmt.Errorf("Error creating IdentityMapping: %s", err)
	}

	d.SetId(result.Id.String())
	log.Println("[DEBUG] End - resourcePingAccessIdentityMappingUpdate")
	return resourcePingAccessIdentityMappingReadResult(d, result)
}

func resourcePingAccessIdentityMappingDelete(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] Start - resourcePingAccessIdentityMappingDelete")
	svc := m.(*pingaccess.Client).IdentityMappings
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	log.Printf("[INFO] ResourceID: %s", d.Id())
	_, err := svc.DeleteIdentityMappingCommand(&pingaccess.DeleteIdentityMappingCommandInput{Id: d.Id()})
	if err != nil {
		return fmt.Errorf("Error deleting virtualhost: %s", err)
	}
	log.Println("[DEBUG] End - resourcePingAccessIdentityMappingDelete")
	return nil
}

func resourcePingAccessIdentityMappingReadResult(d *schema.ResourceData, rv *pingaccess.IdentityMappingView) error {
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
