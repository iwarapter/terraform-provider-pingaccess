package pingaccess

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pa "github.com/iwarapter/pingaccess-sdk-go/pingaccess"
)

func resourcePingAccessHTTPConfigRequestHostSource() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePingAccessHTTPConfigRequestHostSourceCreate,
		ReadContext:   resourcePingAccessHTTPConfigRequestHostSourceRead,
		UpdateContext: resourcePingAccessHTTPConfigRequestHostSourceUpdate,
		DeleteContext: resourcePingAccessHTTPConfigRequestHostSourceDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: resourcePingAccessHTTPConfigRequestHostSourceResourceSchema(),
	}
}

func resourcePingAccessHTTPConfigRequestHostSourceResourceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"header_name_list": requiredListOfString(),
		"list_value_location": {
			Type:             schema.TypeString,
			Required:         true,
			ValidateDiagFunc: validateListLocationValue,
		},
	}
}

func resourcePingAccessHTTPConfigRequestHostSourceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourcePingAccessHTTPConfigRequestHostSourceUpdate(ctx, d, m)
}

func resourcePingAccessHTTPConfigRequestHostSourceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(*pa.Client).HttpConfig
	result, _, err := svc.GetHostSourceCommand()
	if err != nil {
		return diag.FromErr(fmt.Errorf("unable to read HttpConfigHostSource: %s", err))
	}
	return resourcePingAccessHTTPConfigRequestHostSourceReadResult(d, result)
}

func resourcePingAccessHTTPConfigRequestHostSourceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(*pa.Client).HttpConfig
	input := &pa.UpdateHostSourceCommandInput{Body: *resourcePingAccessHTTPConfigRequestHostSourceReadData(d)}
	result, _, err := svc.UpdateHostSourceCommand(input)
	if err != nil {
		return diag.FromErr(fmt.Errorf("unable to update HttpConfigHostSource: %s", err))
	}

	d.SetId("http_config_host_source")
	return resourcePingAccessHTTPConfigRequestHostSourceReadResult(d, result)
}

func resourcePingAccessHTTPConfigRequestHostSourceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(*pa.Client).HttpConfig
	_, err := svc.DeleteHostSourceCommand()
	if err != nil {
		return diag.FromErr(fmt.Errorf("unable to delete HttpConfigHostSource: %s", err))

	}
	return nil
}

func resourcePingAccessHTTPConfigRequestHostSourceReadResult(d *schema.ResourceData, rv *pa.HostMultiValueSourceView) diag.Diagnostics {
	var diags diag.Diagnostics
	setResourceDataStringWithDiagnostic(d, "list_value_location", rv.ListValueLocation, &diags)
	if err := d.Set("header_name_list", rv.HeaderNameList); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	return diags
}

func resourcePingAccessHTTPConfigRequestHostSourceReadData(d *schema.ResourceData) (body *pa.HostMultiValueSourceView) {
	headerNameList := expandStringList(d.Get("header_name_list").([]interface{}))
	body = &pa.HostMultiValueSourceView{
		HeaderNameList:    &headerNameList,
		ListValueLocation: String(d.Get("list_value_location").(string)),
	}
	return
}
