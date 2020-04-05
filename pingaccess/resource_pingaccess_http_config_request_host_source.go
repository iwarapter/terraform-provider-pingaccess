package pingaccess

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	pa "github.com/iwarapter/pingaccess-sdk-go/pingaccess"
)

func resourcePingAccessHTTPConfigRequestHostSource() *schema.Resource {
	return &schema.Resource{
		Create: resourcePingAccessHTTPConfigRequestHostSourceCreate,
		Read:   resourcePingAccessHTTPConfigRequestHostSourceRead,
		Update: resourcePingAccessHTTPConfigRequestHostSourceUpdate,
		Delete: resourcePingAccessHTTPConfigRequestHostSourceDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: resourcePingAccessHTTPConfigRequestHostSourceResourceSchema(),
	}
}

func resourcePingAccessHTTPConfigRequestHostSourceResourceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"header_name_list": requiredListOfString(),
		"list_value_location": &schema.Schema{
			Type:         schema.TypeString,
			Required:     true,
			ValidateFunc: validateListLocationValue,
		},
	}
}

func resourcePingAccessHTTPConfigRequestHostSourceCreate(d *schema.ResourceData, m interface{}) error {
	return resourcePingAccessHTTPConfigRequestHostSourceUpdate(d, m)
}

func resourcePingAccessHTTPConfigRequestHostSourceRead(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pa.Client).HttpConfig
	result, _, err := svc.GetHostSourceCommand()
	if err != nil {
		return fmt.Errorf("Error reading http config host source: %s", err)
	}
	return resourcePingAccessHTTPConfigRequestHostSourceReadResult(d, result)
}

func resourcePingAccessHTTPConfigRequestHostSourceUpdate(d *schema.ResourceData, m interface{}) error {

	svc := m.(*pa.Client).HttpConfig

	input := &pa.UpdateHostSourceCommandInput{Body: *resourcePingAccessHTTPConfigRequestHostSourceReadData(d)}
	result, _, err := svc.UpdateHostSourceCommand(input)
	if err != nil {
		b, _ := json.Marshal(input)
		return fmt.Errorf("Error updating http config host source: %s, %s", err, string(b))
	}

	d.SetId("http_config_host_source")
	return resourcePingAccessHTTPConfigRequestHostSourceReadResult(d, result)
}

func resourcePingAccessHTTPConfigRequestHostSourceDelete(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pa.Client).HttpConfig
	_, err := svc.DeleteHostSourceCommand()
	if err != nil {
		return fmt.Errorf("Error deleting http config host source: %s", err)
	}
	return nil
}

func resourcePingAccessHTTPConfigRequestHostSourceReadResult(d *schema.ResourceData, rv *pa.HostMultiValueSourceView) error {
	if err := d.Set("list_value_location", rv.ListValueLocation); err != nil {
		return err
	}
	if err := d.Set("header_name_list", rv.HeaderNameList); err != nil {
		return err
	}
	return nil
}

func resourcePingAccessHTTPConfigRequestHostSourceReadData(d *schema.ResourceData) (body *pa.HostMultiValueSourceView) {
	headerNameList := expandStringList(d.Get("header_name_list").([]interface{}))
	body = &pa.HostMultiValueSourceView{
		HeaderNameList:    &headerNameList,
		ListValueLocation: String(d.Get("list_value_location").(string)),
	}
	return
}
