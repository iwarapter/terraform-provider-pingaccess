package pingaccess

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/iwarapter/pingaccess-sdk-go/pingaccess"
)

func resourcePingAccessAcmeServer() *schema.Resource {
	return &schema.Resource{
		Create: resourcePingAccessAcmeServerCreate,
		Read:   resourcePingAccessAcmeServerRead,
		//Update: resourcePingAccessAcmeServerUpdate,
		Delete: resourcePingAccessAcmeServerDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: resourcePingAccessAcmeServerSchema(),
	}
}

func resourcePingAccessAcmeServerSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
		"url": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
		"acme_accounts": acmeServerAccountsSchema(),
	}
}

func resourcePingAccessAcmeServerCreate(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pingaccess.Client).Acme
	input := pingaccess.AddAcmeServerCommandInput{
		Body: *resourcePingAccessAcmeServerReadData(d),
	}

	result, _, err := svc.AddAcmeServerCommand(&input)
	if err != nil {
		return fmt.Errorf("Error creating AcmeServer: %s", err)
	}
	d.SetId(*result.Id)
	return resourcePingAccessAcmeServerReadResult(d, &input.Body)
}

func resourcePingAccessAcmeServerRead(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pingaccess.Client).Acme
	input := &pingaccess.GetAcmeServerCommandInput{
		AcmeServerId: d.Id(),
	}
	result, _, err := svc.GetAcmeServerCommand(input)
	if err != nil {
		return fmt.Errorf("Error reading AcmeServer: %s", err)
	}
	return resourcePingAccessAcmeServerReadResult(d, result)
}

func resourcePingAccessAcmeServerDelete(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pingaccess.Client).Acme
	input := &pingaccess.DeleteAcmeServerCommandInput{
		AcmeServerId: d.Id(),
	}

	_, _, err := svc.DeleteAcmeServerCommand(input)
	if err != nil {
		return fmt.Errorf("Error deleting AcmeServer: %s", err)
	}
	return nil
}

func resourcePingAccessAcmeServerReadResult(d *schema.ResourceData, input *pingaccess.AcmeServerView) error {
	setResourceDataString(d, "name", input.Name)
	setResourceDataString(d, "url", input.Url)
	if input.AcmeAccounts != nil && len(input.AcmeAccounts) > 0 {
		if err := d.Set("acme_accounts", flattenLinkViewList(input.AcmeAccounts)); err != nil {
			return err
		}
	}
	return nil
}

func resourcePingAccessAcmeServerReadData(d *schema.ResourceData) *pingaccess.AcmeServerView {
	return &pingaccess.AcmeServerView{
		Name: String(d.Get("name").(string)),
		Url:  String(d.Get("url").(string)),
	}
}
