package pingaccess

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pa "github.com/iwarapter/pingaccess-sdk-go/pingaccess"
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
	}
}

func resourcePingAccessHTTPSListenerSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:             schema.TypeString,
			Required:         true,
			ValidateDiagFunc: validateHTTPListenerName,
			ForceNew:         true,
		},
		"key_pair_id": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"use_server_cipher_suite_order": {
			Type:     schema.TypeBool,
			Required: true,
		},
	}
}

func resourcePingAccessHTTPSListenerCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(*pa.Client).HttpsListeners
	input := pa.GetHttpsListenersCommandInput{}

	result, _, err := svc.GetHttpsListenersCommand(&input)
	if err != nil {
		return diag.FromErr(fmt.Errorf("unable to retrieving listener: %s", err))
	}

	name := d.Get("name").(string)
	for _, listener := range result.Items {
		if *listener.Name == name {
			d.SetId(listener.Id.String())
			return resourcePingAccessHTTPSListenerReadResult(d, listener)
		}
	}
	return diag.FromErr(fmt.Errorf("unable to manage listener: %s", err))
}

func resourcePingAccessHTTPSListenerRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(*pa.Client).HttpsListeners
	input := &pa.GetHttpsListenerCommandInput{
		Id: d.Id(),
	}
	result, _, err := svc.GetHttpsListenerCommand(input)
	if err != nil {
		return diag.FromErr(fmt.Errorf("unable to read listener: %s", err))
	}
	return resourcePingAccessHTTPSListenerReadResult(d, result)
}

func resourcePingAccessHTTPSListenerUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(*pa.Client).HttpsListeners
	input := pa.UpdateHttpsListenerInput{
		Body: *resourcePingAccessHTTPSListenerReadData(d),
		Id:   d.Id(),
	}

	result, _, err := svc.UpdateHttpsListener(&input)
	if err != nil {
		return diag.FromErr(fmt.Errorf("unable to update listener: %s", err))
	}
	return resourcePingAccessHTTPSListenerReadResult(d, result)
}

func resourcePingAccessHTTPSListenerDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
}

func resourcePingAccessHTTPSListenerReadResult(d *schema.ResourceData, input *pa.HttpsListenerView) diag.Diagnostics {
	var diags diag.Diagnostics
	setResourceDataStringWithDiagnostic(d, "name", input.Name, &diags)
	setResourceDataIntWithDiagnostic(d, "key_pair_id", input.KeyPairId, &diags)
	setResourceDataBoolWithDiagnostic(d, "use_server_cipher_suite_order", input.UseServerCipherSuiteOrder, &diags)
	return diags
}

func resourcePingAccessHTTPSListenerReadData(d *schema.ResourceData) *pa.HttpsListenerView {
	engine := &pa.HttpsListenerView{
		Name:                      String(d.Get("name").(string)),
		KeyPairId:                 Int(d.Get("key_pair_id").(int)),
		UseServerCipherSuiteOrder: Bool(d.Get("use_server_cipher_suite_order").(bool)),
	}

	return engine
}
