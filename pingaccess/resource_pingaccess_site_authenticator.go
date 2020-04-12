package pingaccess

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/iwarapter/pingaccess-sdk-go/pingaccess"
	"net/http"
)

func resourcePingAccessSiteAuthenticator() *schema.Resource {
	return &schema.Resource{
		Create: resourcePingAccessSiteAuthenticatorCreate,
		Read:   resourcePingAccessSiteAuthenticatorRead,
		Update: resourcePingAccessSiteAuthenticatorUpdate,
		Delete: resourcePingAccessSiteAuthenticatorDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: resourcePingAccessSiteAuthenticatorSchema(),
	}
}

func resourcePingAccessSiteAuthenticatorSchema() map[string]*schema.Schema {
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
			Type:     schema.TypeString,
			Required: true,
			DiffSuppressFunc: suppressEquivalentJsonDiffs,
		},
	}
}

func resourcePingAccessSiteAuthenticatorCreate(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pingaccess.Client).SiteAuthenticators
	input := pingaccess.AddSiteAuthenticatorCommandInput{
		Body: *resourcePingAccessSiteAuthenticatorReadData(d),
	}

	result, _, err := svc.AddSiteAuthenticatorCommand(&input)
	if err != nil {
		return fmt.Errorf("Error creating site authenticator: %s", err.Error())
	}

	d.SetId(result.Id.String())
	return resourcePingAccessSiteAuthenticatorReadResult(d, result, svc)
}

func resourcePingAccessSiteAuthenticatorRead(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pingaccess.Client).SiteAuthenticators
	input := &pingaccess.GetSiteAuthenticatorCommandInput{
		Id: d.Id(),
	}
	result, _, err := svc.GetSiteAuthenticatorCommand(input)
	if err != nil {
		return fmt.Errorf("Error reading site authenticator: %s", err.Error())
	}
	return resourcePingAccessSiteAuthenticatorReadResult(d, result, svc)
}

func resourcePingAccessSiteAuthenticatorUpdate(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pingaccess.Client).SiteAuthenticators
	input := pingaccess.UpdateSiteAuthenticatorCommandInput{
		Body: *resourcePingAccessSiteAuthenticatorReadData(d),
		Id:   d.Id(),
	}
	result, _, err := svc.UpdateSiteAuthenticatorCommand(&input)
	if err != nil {
		return fmt.Errorf("Error updating site authenticator: %s", err.Error())
	}
	return resourcePingAccessSiteAuthenticatorReadResult(d, result, svc)
}

func resourcePingAccessSiteAuthenticatorDelete(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pingaccess.Client).SiteAuthenticators
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	input := &pingaccess.DeleteSiteAuthenticatorCommandInput{
		Id: d.Id(),
	}

	_, err := svc.DeleteSiteAuthenticatorCommand(input)
	if err != nil {
		return fmt.Errorf("Error deleting site authenticator: %s", err.Error())
	}
	return nil
}

func resourcePingAccessSiteAuthenticatorReadResult(d *schema.ResourceData, input *pingaccess.SiteAuthenticatorView, svc *pingaccess.SiteAuthenticatorsService) error {
	setResourceDataString(d, "name", input.Name)
	setResourceDataString(d, "class_name", input.ClassName)

	b, _ := json.Marshal(input.Configuration)
	config := string(b)

	originalConfig := d.Get("configuration").(string)

	//Search the Site Authenticator descriptors for CONCEALED fields, and update the original value back as we cannot use the
	//encryptedValue provided by the API, whilst this gives us a stable plan - we cannot determine if a CONCEALED value
	//has changed and needs updating
	desc, _, _ := svc.GetSiteAuthenticatorDescriptorsCommand()
	config = maskConfigFromDescriptors(desc, input.ClassName, originalConfig, config)

	if err := d.Set("configuration", config); err != nil {
		return err
	}
	return nil
}

func resourcePingAccessSiteAuthenticatorReadData(d *schema.ResourceData) *pingaccess.SiteAuthenticatorView {
	config := d.Get("configuration").(string)
	var dat map[string]interface{}
	_ = json.Unmarshal([]byte(config), &dat)
	siteAuthenticator := &pingaccess.SiteAuthenticatorView{
		Name:          String(d.Get("name").(string)),
		ClassName:     String(d.Get("class_name").(string)),
		Configuration: dat,
	}
	return siteAuthenticator
}
