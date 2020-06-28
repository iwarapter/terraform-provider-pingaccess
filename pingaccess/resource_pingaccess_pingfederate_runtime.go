package pingaccess

import (
	"context"
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
	deprecationMsg := "This is the deprecated pingfederate runtime configuration and is only to provide support for 5.x. If you are using 6.x please use 'issuer' configuration instead."
	return map[string]*schema.Schema{
		//New API
		"description": {
			Type:          schema.TypeString,
			Optional:      true,
			RequiredWith:  []string{"issuer"},
			ConflictsWith: []string{"audit_level", "back_channel_base_path", "back_channel_secure", "base_path", "expected_hostname", "host", "port", "secure", "targets"},
		},
		"issuer": {
			Type:          schema.TypeString,
			Optional:      true,
			ConflictsWith: []string{"audit_level", "back_channel_base_path", "back_channel_secure", "base_path", "expected_hostname", "host", "port", "secure", "targets"},
		},
		"sts_token_exchange_endpoint": {
			Type:          schema.TypeString,
			Optional:      true,
			RequiredWith:  []string{"issuer"},
			ConflictsWith: []string{"audit_level", "back_channel_base_path", "back_channel_secure", "base_path", "expected_hostname", "host", "port", "secure", "targets"},
		},

		//Deprecated API
		"audit_level": {
			Type:          schema.TypeString,
			Optional:      true,
			Default:       "ON",
			ConflictsWith: []string{"description", "issuer", "sts_token_exchange_endpoint"},
			Deprecated:    deprecationMsg,
		},
		"back_channel_base_path": {
			Type:          schema.TypeString,
			Optional:      true,
			ConflictsWith: []string{"description", "issuer", "sts_token_exchange_endpoint"},
			Deprecated:    deprecationMsg,
		},
		"back_channel_secure": {
			Type:          schema.TypeBool,
			Optional:      true,
			Default:       false,
			ConflictsWith: []string{"description", "issuer", "sts_token_exchange_endpoint"},
			Deprecated:    deprecationMsg,
		},
		"base_path": {
			Type:          schema.TypeString,
			Optional:      true,
			ConflictsWith: []string{"description", "issuer", "sts_token_exchange_endpoint"},
			Deprecated:    deprecationMsg,
		},
		"expected_hostname": {
			Type:          schema.TypeString,
			Optional:      true,
			ConflictsWith: []string{"description", "issuer", "sts_token_exchange_endpoint"},
			Deprecated:    deprecationMsg,
		},
		"host": {
			Type:          schema.TypeString,
			Optional:      true,
			RequiredWith:  []string{"host", "port"},
			ConflictsWith: []string{"description", "issuer", "sts_token_exchange_endpoint"},
			Deprecated:    deprecationMsg,
		},
		"port": {
			Type:          schema.TypeInt,
			Optional:      true,
			RequiredWith:  []string{"host", "port"},
			ConflictsWith: []string{"description", "issuer", "sts_token_exchange_endpoint"},
			Deprecated:    deprecationMsg,
		},
		"secure": {
			Type:          schema.TypeBool,
			Optional:      true,
			Default:       false,
			ConflictsWith: []string{"description", "issuer", "sts_token_exchange_endpoint"},
			Deprecated:    deprecationMsg,
		},
		"targets": {
			Type:     schema.TypeSet,
			Optional: true,
			MinItems: 0,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			ConflictsWith: []string{"description", "issuer", "sts_token_exchange_endpoint"},
			Deprecated:    deprecationMsg,
		},

		//Common
		"skip_hostname_verification": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"trusted_certificate_group_id": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  0,
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

func resourcePingAccessPingFederateRuntimeRead(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(*pa.Client).Pingfederate
	if _, ok := d.GetOk("issuer"); ok {
		result, _, err := svc.GetPingFederateRuntimeCommand()
		if err != nil {
			return diag.Errorf("unable to read PingFederateRuntime: %s", err)
		}

		return resourcePingAccessPingFederateRuntimeReadResult(d, result)
	}

	result, _, err := svc.GetPingFederateCommand()
	if err != nil {
		return diag.Errorf("unable to read deprecated PingFederateRuntime: %s", err)
	}

	return resourcePingAccessPingFederateDeprecatedRuntimeReadResult(d, result)
}

func resourcePingAccessPingFederateRuntimeUpdate(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(*pa.Client).Pingfederate

	if _, ok := d.GetOk("issuer"); ok {
		input := pa.UpdatePingFederateRuntimeCommandInput{
			Body: *resourcePingAccessPingFederateRuntimeReadData(d),
		}
		result, _, err := svc.UpdatePingFederateRuntimeCommand(&input)
		if err != nil {
			return diag.Errorf("unable to update PingFederateRuntime: %s", err)
		}

		d.SetId("pingfederate_runtime_settings")
		return resourcePingAccessPingFederateRuntimeReadResult(d, result)
	}

	input := pa.UpdatePingFederateCommandInput{
		Body: *resourcePingAccessPingFederateDeprecatedRuntimeReadData(d),
	}
	result, _, err := svc.UpdatePingFederateCommand(&input)
	if err != nil {
		return diag.Errorf("error updating deprecated PingFederateRuntime: %s", err)
	}

	d.SetId("pingfederate_runtime_settings")
	return resourcePingAccessPingFederateDeprecatedRuntimeReadResult(d, result)
}

func resourcePingAccessPingFederateRuntimeDelete(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(*pa.Client).Pingfederate
	if _, ok := d.GetOk("issuer"); ok {
		_, err := svc.DeletePingFederateRuntimeCommand()
		if err != nil {
			return diag.Errorf("unable to reset PingFederateRuntime: %s", err)
		}
		return nil
	}

	_, err := svc.DeletePingFederateCommand()
	if err != nil {
		return diag.Errorf("unable to reset deprecated PingFederateRuntime: %s", err)
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

func resourcePingAccessPingFederateDeprecatedRuntimeReadResult(d *schema.ResourceData, input *pa.PingFederateRuntimeView) diag.Diagnostics {
	var diags diag.Diagnostics
	setResourceDataStringWithDiagnostic(d, "audit_level", input.AuditLevel, &diags)
	setResourceDataStringWithDiagnostic(d, "back_channel_base_path", input.BackChannelBasePath, &diags)
	setResourceDataBoolWithDiagnostic(d, "back_channel_secure", input.BackChannelSecure, &diags)
	setResourceDataStringWithDiagnostic(d, "base_path", input.BasePath, &diags)
	setResourceDataStringWithDiagnostic(d, "expected_hostname", input.ExpectedHostname, &diags)
	setResourceDataStringWithDiagnostic(d, "host", input.Host, &diags)
	setResourceDataIntWithDiagnostic(d, "port", input.Port, &diags)
	setResourceDataBoolWithDiagnostic(d, "secure", input.Secure, &diags)
	setResourceDataBoolWithDiagnostic(d, "skip_hostname_verification", input.SkipHostnameVerification, &diags)
	setResourceDataIntWithDiagnostic(d, "trusted_certificate_group_id", input.TrustedCertificateGroupId, &diags)
	setResourceDataBoolWithDiagnostic(d, "use_proxy", input.UseProxy, &diags)
	setResourceDataBoolWithDiagnostic(d, "use_slo", input.UseSlo, &diags)
	if input.Targets != nil {
		if err := d.Set("targets", input.Targets); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}
	return diags
}

func resourcePingAccessPingFederateRuntimeReadData(d *schema.ResourceData) *pa.PingFederateMetadataRuntimeView {
	pfRuntime := &pa.PingFederateMetadataRuntimeView{
		Issuer:                    String(d.Get("issuer").(string)),
		TrustedCertificateGroupId: Int(d.Get("trusted_certificate_group_id").(int)),
		UseProxy:                  Bool(d.Get("use_proxy").(bool)),
		UseSlo:                    Bool(d.Get("use_slo").(bool)),
	}
	if v, ok := d.GetOk("description"); ok {
		pfRuntime.Description = String(v.(string))
	}
	pfRuntime.SkipHostnameVerification = Bool(d.Get("skip_hostname_verification").(bool))
	if v, ok := d.GetOk("sts_token_exchange_endpoint"); ok {
		pfRuntime.StsTokenExchangeEndpoint = String(v.(string))
	}

	return pfRuntime
}

func resourcePingAccessPingFederateDeprecatedRuntimeReadData(d *schema.ResourceData) *pa.PingFederateRuntimeView {
	pfRuntime := &pa.PingFederateRuntimeView{
		Host:                      String(d.Get("host").(string)),
		Port:                      Int(d.Get("port").(int)),
		TrustedCertificateGroupId: Int(d.Get("trusted_certificate_group_id").(int)),
		UseProxy:                  Bool(d.Get("use_proxy").(bool)),
		UseSlo:                    Bool(d.Get("use_slo").(bool)),
	}

	if v, ok := d.GetOk("audit_level"); ok {
		pfRuntime.AuditLevel = String(v.(string))
	}

	if v, ok := d.GetOk("back_channel_base_path"); ok {
		pfRuntime.BackChannelBasePath = String(v.(string))
	}

	if v, ok := d.GetOk("back_channel_secure"); ok {
		pfRuntime.BackChannelSecure = Bool(v.(bool))
	}

	if v, ok := d.GetOk("base_path"); ok {
		pfRuntime.BasePath = String(v.(string))
	}

	if v, ok := d.GetOk("expected_hostname"); ok {
		pfRuntime.ExpectedHostname = String(v.(string))
	}

	if v, ok := d.GetOk("secure"); ok {
		pfRuntime.Secure = Bool(v.(bool))
	}

	if v, ok := d.GetOk("skip_hostname_verification"); ok {
		pfRuntime.SkipHostnameVerification = Bool(v.(bool))
	}

	if v, ok := d.GetOk("targets"); ok {
		targets := expandStringList(v.(*schema.Set).List())
		pfRuntime.Targets = &targets
	}

	return pfRuntime
}
