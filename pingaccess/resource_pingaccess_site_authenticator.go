package pingaccess

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/tidwall/gjson"

	"github.com/hashicorp/terraform/helper/schema"
	pa "github.com/iwarapter/pingaccess-sdk-go/pingaccess"
	"github.com/tidwall/sjson"
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
		"hidden_fields": setOfString(),
	}
}

func resourcePingAccessSiteAuthenticatorCreate(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pa.Client).SiteAuthenticators
	input := pa.AddSiteAuthenticatorCommandInput{
		Body: *resourcePingAccessSiteAuthenticatorReadData(d),
	}

	result, _, err := svc.AddSiteAuthenticatorCommand(&input)
	if err != nil {
		return fmt.Errorf("Error creating site: %s", err)
	}

	d.SetId(result.Id.String())
	return resourcePingAccessSiteAuthenticatorReadResult(d, result)
}

func resourcePingAccessSiteAuthenticatorRead(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pa.Client).SiteAuthenticators
	input := &pa.GetSiteAuthenticatorCommandInput{
		Id: d.Id(),
	}
	result, _, err := svc.GetSiteAuthenticatorCommand(input)
	if err != nil {
		return fmt.Errorf("Error reading site: %s", err)
	}
	return resourcePingAccessSiteAuthenticatorReadResult(d, result)
}

func resourcePingAccessSiteAuthenticatorUpdate(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pa.Client).SiteAuthenticators
	input := pa.UpdateSiteAuthenticatorCommandInput{
		Body: *resourcePingAccessSiteAuthenticatorReadData(d),
		Id:   d.Id(),
	}
	result, _, err := svc.UpdateSiteAuthenticatorCommand(&input)
	if err != nil {
		return fmt.Errorf("Error updating virtualhost: %s", err)
	}
	return resourcePingAccessSiteAuthenticatorReadResult(d, result)
}

func resourcePingAccessSiteAuthenticatorDelete(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pa.Client).SiteAuthenticators
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	input := &pa.DeleteSiteAuthenticatorCommandInput{
		Id: d.Id(),
	}

	_, err := svc.DeleteSiteAuthenticatorCommand(input)
	if err != nil {
		return fmt.Errorf("Error deleting virtualhost: %s", err)
	}
	return nil
}

func resourcePingAccessSiteAuthenticatorReadResult(d *schema.ResourceData, input *pa.SiteAuthenticatorView) error {
	setResourceDataString(d, "name", input.Name)
	setResourceDataString(d, "class_name", input.ClassName)

	b, _ := json.Marshal(input.Configuration)
	config := string(b)

	originalConfig := d.Get("configuration").(string)
	if _, ok := d.GetOkExists("hidden_fields"); ok {
		hiddenFields := expandStringList(d.Get("hidden_fields").(*schema.Set).List())
		for _, f := range hiddenFields {
			path := fmt.Sprintf("%s.value", *f)
			v := gjson.Get(originalConfig, path)
			if v.Exists() {
				config, _ = sjson.Set(config, path, v.String())
			}
			path = fmt.Sprintf("%s.encryptedValue", *f)
			config, _ = sjson.Delete(config, path)
		}
	}
	if err := d.Set("configuration", config); err != nil {
		return err
	}
	return nil
}

func resourcePingAccessSiteAuthenticatorReadData(d *schema.ResourceData) *pa.SiteAuthenticatorView {
	config := d.Get("configuration").(string)
	var dat map[string]interface{}
	_ = json.Unmarshal([]byte(config), &dat)
	siteAuthenticator := &pa.SiteAuthenticatorView{
		Name:          String(d.Get("name").(string)),
		ClassName:     String(d.Get("class_name").(string)),
		Configuration: dat,
	}
	return siteAuthenticator
}
