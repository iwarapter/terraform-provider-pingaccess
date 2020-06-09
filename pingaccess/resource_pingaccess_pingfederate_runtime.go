package pingaccess

import (
	"context"
	"fmt"

	pa "github.com/iwarapter/pingaccess-sdk-go/pingaccess"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePingAccessPingFederateRuntime() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePingAccessPingFederateRuntimeCreate,
		ReadContext:   resourcePingAccessPingFederateRuntimeRead,
		UpdateContext: resourcePingAccessPingFederateRuntimeUpdate,
		DeleteContext: resourcePingAccessPingFederateRuntimeDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: resourcePingAccessPingFederateRuntimeSchema(),
	}
}

func resourcePingAccessPingFederateRuntimeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"description": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"issuer": {
			Type:     schema.TypeString,
			Required: true,
		},
		"skip_hostname_verification": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"sts_token_exchange_endpoint": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"trusted_certificate_group_id": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"use_proxy": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"use_slo": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
	}
}

func resourcePingAccessPingFederateRuntimeCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	d.SetId("pingfederate_runtime_settings")
	return resourcePingAccessPingFederateRuntimeUpdate(ctx, d, m)
}

func resourcePingAccessPingFederateRuntimeRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(*pa.Client).PingFederate
	result, _, err := svc.GetPingFederateRuntimeCommand()
	if err != nil {
		return diag.Diagnostics{diag.FromErr(fmt.Errorf("unable to read PingFederateRuntime: %s", err))}
	}

	return resourcePingAccessPingFederateRuntimeReadResult(d, result)
}

func resourcePingAccessPingFederateRuntimeUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(*pa.Client).PingFederate
	input := pa.UpdatePingFederateRuntimeCommandInput{
		Body: *resourcePingAccessPingFederateRuntimeReadData(d),
	}
	result, _, err := svc.UpdatePingFederateRuntimeCommand(&input)
	if err != nil {
		return diag.Diagnostics{diag.FromErr(fmt.Errorf("unable to update PingFederateRuntime: %s", err))}
	}

	d.SetId("pingfederate_runtime_settings")
	return resourcePingAccessPingFederateRuntimeReadResult(d, result)
}

func resourcePingAccessPingFederateRuntimeDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(*pa.Client).PingFederate
	_, err := svc.DeletePingFederateCommand()
	if err != nil {
		return diag.Diagnostics{diag.FromErr(fmt.Errorf("unable to reset PingFederateRuntime: %s", err))}
	}
	return nil
}

func resourcePingAccessPingFederateRuntimeReadResult(d *schema.ResourceData, input *pa.PingFederateMetadataRuntimeView) diag.Diagnostics {
	var diags diag.Diagnostics
	setResourceDataStringWithDiagnostic(d, "description", input.Description, &diags)
	setResourceDataStringWithDiagnostic(d, "issuer", input.Issuer, &diags)
	setResourceDataBoolWithDiagnostic(d, "skip_hostname_verification", input.SkipHostnameVerification, &diags)
	setResourceDataStringWithDiagnostic(d, "sts_token_exchange_endpoint", input.StsTokenExchangeEndpoint, &diags)
	setResourceDataIntWithDiagnostic(d, "trusted_certificate_group_id", input.TrustedCertificateGroupId, &diags)
	setResourceDataBoolWithDiagnostic(d, "use_proxy", input.UseProxy, &diags)
	setResourceDataBoolWithDiagnostic(d, "use_slo", input.UseSlo, &diags)
	return diags
}

func resourcePingAccessPingFederateRuntimeReadData(d *schema.ResourceData) *pa.PingFederateMetadataRuntimeView {
	pfRuntime := &pa.PingFederateMetadataRuntimeView{
		Issuer: String(d.Get("issuer").(string)),
	}
	if v, ok := d.GetOk("description"); ok {
		pfRuntime.Description = String(v.(string))
	}
	if v, ok := d.GetOk("skip_hostname_verification"); ok {
		pfRuntime.SkipHostnameVerification = Bool(v.(bool))
	}
	if v, ok := d.GetOk("sts_token_exchange_endpoint"); ok {
		pfRuntime.StsTokenExchangeEndpoint = String(v.(string))
	}
	if v, ok := d.GetOkExists("trusted_certificate_group_id"); ok {
		pfRuntime.TrustedCertificateGroupId = Int(v.(int))
	}
	if v, ok := d.GetOkExists("use_proxy"); ok {
		pfRuntime.UseProxy = Bool(v.(bool))
	}
	if v, ok := d.GetOkExists("use_slo"); ok {
		pfRuntime.UseSlo = Bool(v.(bool))
	}
	return pfRuntime
}
