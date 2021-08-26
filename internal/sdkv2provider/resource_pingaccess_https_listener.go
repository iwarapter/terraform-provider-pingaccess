package sdkv2provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess/models"
	"github.com/iwarapter/pingaccess-sdk-go/v62/services/httpsListeners"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePingAccessHTTPSListener() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePingAccessHTTPSListenerCreate,
		ReadContext:   resourcePingAccessHTTPSListenerRead,
		UpdateContext: resourcePingAccessHTTPSListenerUpdate,
		DeleteContext: resourcePingAccessHTTPSListenerDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: resourcePingAccessHTTPSListenerSchema(),
		Description: `Manages the PingAccess HTTPS Listeners configuration.

-> This resource manages a fixed resources within PingAccess and can only manage the specified names. Deleting this resource just stops tracking it's configuration.`,
	}
}

func resourcePingAccessHTTPSListenerSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:         schema.TypeString,
			Required:     true,
			ValidateFunc: validation.StringInSlice([]string{"ADMIN", "AGENT", "ENGINE", "CONFIG QUERY", "SIDEBAND"}, false),
			ForceNew:     true,
			Description:  "The name of the HTTPS listener. One of `ADMIN`, `AGENT`, `ENGINE`, `CONFIG QUERY`, `SIDEBAND`",
		},
		"key_pair_id": {
			Type:        schema.TypeInt,
			Required:    true,
			Description: "The ID of the default key pair used by the HTTPS listener.",
		},
		"use_server_cipher_suite_order": {
			Type:        schema.TypeBool,
			Required:    true,
			Description: "Enable server cipher suite ordering for the HTTPS listener.",
		},
	}
}

func resourcePingAccessHTTPSListenerCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(paClient).HttpsListeners

	result, _, err := svc.GetHttpsListenersCommand(&httpsListeners.GetHttpsListenersCommandInput{})
	if err != nil {
		return diag.Errorf("unable to retrieving listener: %s", err)
	}

	name := d.Get("name").(string)
	for _, listener := range result.Items {
		if *listener.Name == name {
			d.SetId(listener.Id.String())
			return resourcePingAccessHTTPSListenerUpdate(ctx, d, m)
		}
	}
	return diag.Errorf("unable to manage listener: %s", err)
}

func resourcePingAccessHTTPSListenerRead(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(paClient).HttpsListeners
	input := &httpsListeners.GetHttpsListenerCommandInput{
		Id: d.Id(),
	}
	result, _, err := svc.GetHttpsListenerCommand(input)
	if err != nil {
		return diag.Errorf("unable to read listener: %s", err)
	}
	return resourcePingAccessHTTPSListenerReadResult(d, result)
}

func resourcePingAccessHTTPSListenerUpdate(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(paClient).HttpsListeners
	input := httpsListeners.UpdateHttpsListenerInput{
		Body: *resourcePingAccessHTTPSListenerReadData(d),
		Id:   d.Id(),
	}

	result, _, err := svc.UpdateHttpsListener(&input)
	if err != nil {
		return diag.Errorf("unable to update listener: %s", err)
	}
	return resourcePingAccessHTTPSListenerReadResult(d, result)
}

func resourcePingAccessHTTPSListenerDelete(_ context.Context, _ *schema.ResourceData, _ interface{}) diag.Diagnostics {
	return nil
}

func resourcePingAccessHTTPSListenerReadResult(d *schema.ResourceData, input *models.HttpsListenerView) diag.Diagnostics {
	var diags diag.Diagnostics
	setResourceDataStringWithDiagnostic(d, "name", input.Name, &diags)
	setResourceDataIntWithDiagnostic(d, "key_pair_id", input.KeyPairId, &diags)
	setResourceDataBoolWithDiagnostic(d, "use_server_cipher_suite_order", input.UseServerCipherSuiteOrder, &diags)
	return diags
}

func resourcePingAccessHTTPSListenerReadData(d *schema.ResourceData) *models.HttpsListenerView {
	engine := &models.HttpsListenerView{
		Name:                      String(d.Get("name").(string)),
		KeyPairId:                 Int(d.Get("key_pair_id").(int)),
		UseServerCipherSuiteOrder: Bool(d.Get("use_server_cipher_suite_order").(bool)),
	}

	return engine
}
