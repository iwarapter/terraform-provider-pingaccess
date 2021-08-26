package sdkv2provider

import (
	"context"

	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess/models"
	"github.com/iwarapter/pingaccess-sdk-go/v62/services/httpConfig"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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
		Description: `Manages the PingAccess HTTP Request Host Source configuration.

-> This resource manages a singleton within PingAccess and as such you should ONLY ever declare one of this resource type. Deleting this resource resets the HTTP Request Host Source configuration to default values.`,
	}
}

func resourcePingAccessHTTPConfigRequestHostSourceResourceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"header_name_list": {
			Type:        schema.TypeList,
			Required:    true,
			Description: "An array of header names used to identify the host source name.",
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"list_value_location": {
			Type:             schema.TypeString,
			Required:         true,
			Description:      "The location in a matching header value list to use as the source.",
			ValidateDiagFunc: validateListLocationValue,
		},
	}
}

func resourcePingAccessHTTPConfigRequestHostSourceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourcePingAccessHTTPConfigRequestHostSourceUpdate(ctx, d, m)
}

func resourcePingAccessHTTPConfigRequestHostSourceRead(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(paClient).HttpConfig
	result, _, err := svc.GetHostSourceCommand()
	if err != nil {
		return diag.Errorf("unable to read HttpConfigHostSource: %s", err)
	}
	return resourcePingAccessHTTPConfigRequestHostSourceReadResult(d, result)
}

func resourcePingAccessHTTPConfigRequestHostSourceUpdate(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(paClient).HttpConfig
	input := &httpConfig.UpdateHostSourceCommandInput{Body: *resourcePingAccessHTTPConfigRequestHostSourceReadData(d)}
	result, _, err := svc.UpdateHostSourceCommand(input)
	if err != nil {
		return diag.Errorf("unable to update HttpConfigHostSource: %s", err)
	}

	d.SetId("http_config_host_source")
	return resourcePingAccessHTTPConfigRequestHostSourceReadResult(d, result)
}

func resourcePingAccessHTTPConfigRequestHostSourceDelete(_ context.Context, _ *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(paClient).HttpConfig
	_, err := svc.DeleteHostSourceCommand()
	if err != nil {
		return diag.Errorf("unable to delete HttpConfigHostSource: %s", err)

	}
	return nil
}

func resourcePingAccessHTTPConfigRequestHostSourceReadResult(d *schema.ResourceData, rv *models.HostMultiValueSourceView) diag.Diagnostics {
	var diags diag.Diagnostics
	setResourceDataStringWithDiagnostic(d, "list_value_location", rv.ListValueLocation, &diags)
	if err := d.Set("header_name_list", rv.HeaderNameList); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	return diags
}

func resourcePingAccessHTTPConfigRequestHostSourceReadData(d *schema.ResourceData) (body *models.HostMultiValueSourceView) {
	headerNameList := expandStringList(d.Get("header_name_list").([]interface{}))
	body = &models.HostMultiValueSourceView{
		HeaderNameList:    &headerNameList,
		ListValueLocation: String(d.Get("list_value_location").(string)),
	}
	return
}
