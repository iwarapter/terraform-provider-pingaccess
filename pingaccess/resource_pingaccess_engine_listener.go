package pingaccess

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pa "github.com/iwarapter/pingaccess-sdk-go/pingaccess"
)

func resourcePingAccessEngineListener() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePingAccessEngineListenerCreate,
		ReadContext:   resourcePingAccessEngineListenerRead,
		UpdateContext: resourcePingAccessEngineListenerUpdate,
		DeleteContext: resourcePingAccessEngineListenerDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: resourcePingAccessEngineListenerSchema(),
	}
}

func resourcePingAccessEngineListenerSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"port": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"secure": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"trusted_certificate_group_id": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  0,
		},
	}
}

func resourcePingAccessEngineListenerCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(*pa.Client).EngineListeners
	input := pa.AddEngineListenerCommandInput{
		Body: *resourcePingAccessEngineListenerReadData(d),
	}

	result, _, err := svc.AddEngineListenerCommand(&input)
	if err != nil {
		return diag.FromErr(fmt.Errorf("unable to create EngineListener: %s", err))
	}

	d.SetId(result.Id.String())
	return resourcePingAccessEngineListenerReadResult(d, &input.Body)
}

func resourcePingAccessEngineListenerRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(*pa.Client).EngineListeners
	input := &pa.GetEngineListenerCommandInput{
		Id: d.Id(),
	}
	result, _, err := svc.GetEngineListenerCommand(input)
	if err != nil {
		return diag.FromErr(fmt.Errorf("unable to read EngineListener: %s", err))
	}
	return resourcePingAccessEngineListenerReadResult(d, result)
}

func resourcePingAccessEngineListenerUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(*pa.Client).EngineListeners
	input := pa.UpdateEngineListenerCommandInput{
		Body: *resourcePingAccessEngineListenerReadData(d),
		Id:   d.Id(),
	}

	result, _, err := svc.UpdateEngineListenerCommand(&input)
	if err != nil {
		return diag.FromErr(fmt.Errorf("unable to update EngineListener: %s", err))
	}
	return resourcePingAccessEngineListenerReadResult(d, result)
}

func resourcePingAccessEngineListenerDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(*pa.Client).EngineListeners
	input := &pa.DeleteEngineListenerCommandInput{
		Id: d.Id(),
	}

	_, err := svc.DeleteEngineListenerCommand(input)
	if err != nil {
		return diag.FromErr(fmt.Errorf("unable to delete EngineListener: %s", err))

	}
	return nil
}

func resourcePingAccessEngineListenerReadResult(d *schema.ResourceData, input *pa.EngineListenerView) diag.Diagnostics {
	var diags diag.Diagnostics
	setResourceDataStringWithDiagnostic(d, "name", input.Name, &diags)
	setResourceDataIntWithDiagnostic(d, "port", input.Port, &diags)
	setResourceDataBoolWithDiagnostic(d, "secure", input.Secure, &diags)
	setResourceDataIntWithDiagnostic(d, "trusted_certificate_group_id", input.TrustedCertificateGroupId, &diags)
	return diags
}

func resourcePingAccessEngineListenerReadData(d *schema.ResourceData) *pa.EngineListenerView {
	engine := &pa.EngineListenerView{
		Name: String(d.Get("name").(string)),
		Port: Int(d.Get("port").(int)),
	}
	engine.Secure = Bool(d.Get("secure").(bool))
	engine.TrustedCertificateGroupId = Int(d.Get("trusted_certificate_group_id").(int))

	return engine
}
