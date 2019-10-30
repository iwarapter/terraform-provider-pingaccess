package pingaccess

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/iwarapter/pingaccess-sdk-go/pingaccess"
)

func resourcePingAccessAuthnReqList() *schema.Resource {
	return &schema.Resource{
		Create: resourcePingAccessAuthnReqListCreate,
		Read:   resourcePingAccessAuthnReqListRead,
		Update: resourcePingAccessAuthnReqListUpdate,
		Delete: resourcePingAccessAuthnReqListDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: resourcePingAccessAuthnReqListSchema(),
	}
}

func resourcePingAccessAuthnReqListSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"authn_reqs": requiredListOfString(),
	}
}

func resourcePingAccessAuthnReqListCreate(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pingaccess.Client).AuthnReqLists
	input := pingaccess.AddAuthnReqListCommandInput{
		Body: *resourcePingAccessAuthnReqListReadData(d),
	}

	result, _, err := svc.AddAuthnReqListCommand(&input)
	if err != nil {
		return fmt.Errorf("Error creating AuthnReqList: %s", err)
	}

	d.SetId(result.Id.String())
	return resourcePingAccessAuthnReqListReadResult(d, &input.Body)
}

func resourcePingAccessAuthnReqListRead(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pingaccess.Client).AuthnReqLists
	input := &pingaccess.GetAuthnReqListCommandInput{
		Id: d.Id(),
	}
	result, _, err := svc.GetAuthnReqListCommand(input)
	if err != nil {
		return fmt.Errorf("Error reading AuthnReqList: %s", err)
	}
	return resourcePingAccessAuthnReqListReadResult(d, result)
}

func resourcePingAccessAuthnReqListUpdate(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pingaccess.Client).AuthnReqLists
	input := pingaccess.UpdateAuthnReqListCommandInput{
		Body: *resourcePingAccessAuthnReqListReadData(d),
		Id:   d.Id(),
	}

	result, _, err := svc.UpdateAuthnReqListCommand(&input)
	if err != nil {
		return fmt.Errorf("Error updating AuthnReqList: %s", err.Error())
	}
	return resourcePingAccessAuthnReqListReadResult(d, result)
}

func resourcePingAccessAuthnReqListDelete(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pingaccess.Client).AuthnReqLists
	input := &pingaccess.DeleteAuthnReqListCommandInput{
		Id: d.Id(),
	}

	_, err := svc.DeleteAuthnReqListCommand(input)
	if err != nil {
		return fmt.Errorf("Error deleting AuthnReqList: %s", err)
	}
	return nil
}

func resourcePingAccessAuthnReqListReadResult(d *schema.ResourceData, input *pingaccess.AuthnReqListView) error {
	setResourceDataString(d, "name", input.Name)
	if err := d.Set("authn_reqs", input.AuthnReqs); err != nil {
		return err
	}
	return nil
}

func resourcePingAccessAuthnReqListReadData(d *schema.ResourceData) *pingaccess.AuthnReqListView {
	auths := expandStringList(d.Get("authn_reqs").([]interface{}))
	engine := &pingaccess.AuthnReqListView{
		Name:      String(d.Get("name").(string)),
		AuthnReqs: &auths,
	}

	return engine
}
