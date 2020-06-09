package pingaccess

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pa "github.com/iwarapter/pingaccess-sdk-go/pingaccess"
)

func resourcePingAccessThirdPartyService() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePingAccessThirdPartyServiceCreate,
		ReadContext:   resourcePingAccessThirdPartyServiceRead,
		UpdateContext: resourcePingAccessThirdPartyServiceUpdate,
		DeleteContext: resourcePingAccessThirdPartyServiceDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: resourcePingAccessThirdPartyServiceSchema(),
	}
}

func resourcePingAccessThirdPartyServiceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"availability_profile_id": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  1,
		},
		"expected_hostname": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"host_value": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"load_balancing_strategy_id": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"max_connections": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  -1,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"secure": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"skip_hostname_verification": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"targets": {
			Type:     schema.TypeSet,
			Required: true,
			MinItems: 1,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
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
	}
}

func resourcePingAccessThirdPartyServiceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(*pa.Client).ThirdPartyServices
	input := pa.AddThirdPartyServiceCommandInput{
		Body: *resourcePingAccessThirdPartyServiceReadData(d),
	}

	result, _, err := svc.AddThirdPartyServiceCommand(&input)
	if err != nil {
		return diag.Diagnostics{diag.FromErr(fmt.Errorf("unable to create ThirdPartyService: %s", err))}

	}

	d.SetId(*result.Id)
	return resourcePingAccessThirdPartyServiceReadResult(d, result)
}

func resourcePingAccessThirdPartyServiceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(*pa.Client).ThirdPartyServices
	input := &pa.GetThirdPartyServiceCommandInput{
		Id: d.Id(),
	}
	result, _, err := svc.GetThirdPartyServiceCommand(input)
	if err != nil {
		return diag.Diagnostics{diag.FromErr(fmt.Errorf("unable to read ThirdPartyService: %s", err))}
	}
	return resourcePingAccessThirdPartyServiceReadResult(d, result)
}

func resourcePingAccessThirdPartyServiceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(*pa.Client).ThirdPartyServices
	input := pa.UpdateThirdPartyServiceCommandInput{
		Body: *resourcePingAccessThirdPartyServiceReadData(d),
		Id:   d.Id(),
	}
	result, _, err := svc.UpdateThirdPartyServiceCommand(&input)
	if err != nil {
		return diag.Diagnostics{diag.FromErr(fmt.Errorf("unable to update ThirdPartyService: %s", err))}
	}
	return resourcePingAccessThirdPartyServiceReadResult(d, result)
}

func resourcePingAccessThirdPartyServiceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(*pa.Client).ThirdPartyServices
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	input := &pa.DeleteThirdPartyServiceCommandInput{
		Id: d.Id(),
	}

	_, err := svc.DeleteThirdPartyServiceCommand(input)
	if err != nil {
		return diag.Diagnostics{diag.FromErr(fmt.Errorf("unable to delete ThirdPartyService: %s", err))}
	}
	return nil
}

func resourcePingAccessThirdPartyServiceReadResult(d *schema.ResourceData, input *pa.ThirdPartyServiceView) diag.Diagnostics {
	var diags diag.Diagnostics
	setResourceDataIntWithDiagnostic(d, "availability_profile_id", input.AvailabilityProfileId, &diags)
	setResourceDataStringWithDiagnostic(d, "expected_hostname", input.ExpectedHostname, &diags)
	setResourceDataStringWithDiagnostic(d, "host_value", input.HostValue, &diags)
	setResourceDataIntWithDiagnostic(d, "load_balancing_strategy_id", input.LoadBalancingStrategyId, &diags)
	setResourceDataIntWithDiagnostic(d, "max_connections", input.MaxConnections, &diags)
	setResourceDataStringWithDiagnostic(d, "name", input.Name, &diags)
	setResourceDataBoolWithDiagnostic(d, "secure", input.Secure, &diags)
	setResourceDataBoolWithDiagnostic(d, "skip_hostname_verification", input.SkipHostnameVerification, &diags)
	setResourceDataIntWithDiagnostic(d, "trusted_certificate_group_id", input.TrustedCertificateGroupId, &diags)
	setResourceDataBoolWithDiagnostic(d, "use_proxy", input.UseProxy, &diags)
	if err := d.Set("targets", input.Targets); err != nil {
		diags = append(diags, diag.FromErr(err))
	}
	return diags
}

func resourcePingAccessThirdPartyServiceReadData(d *schema.ResourceData) *pa.ThirdPartyServiceView {
	targets := expandStringList(d.Get("targets").(*schema.Set).List())
	tps := &pa.ThirdPartyServiceView{
		Name:    String(d.Get("name").(string)),
		Targets: &targets,
	}

	if val, ok := d.GetOkExists("availability_profile_id"); ok {
		tps.AvailabilityProfileId = Int(val.(int))
	}
	if val, ok := d.GetOkExists("expected_hostname"); ok {
		tps.ExpectedHostname = String(val.(string))
	}
	if val, ok := d.GetOkExists("host_value"); ok {
		tps.HostValue = String(val.(string))
	}
	if val, ok := d.GetOkExists("load_balancing_strategy_id"); ok {
		tps.LoadBalancingStrategyId = Int(val.(int))
	}
	if val, ok := d.GetOkExists("max_connections"); ok {
		tps.MaxConnections = Int(val.(int))
	}
	if val, ok := d.GetOkExists("secure"); ok {
		tps.Secure = Bool(val.(bool))
	}
	if val, ok := d.GetOkExists("skip_hostname_verification"); ok {
		tps.SkipHostnameVerification = Bool(val.(bool))
	}
	if val, ok := d.GetOkExists("trusted_certificate_group_id"); ok {
		tps.TrustedCertificateGroupId = Int(val.(int))
	}
	if val, ok := d.GetOkExists("use_proxy"); ok {
		tps.UseProxy = Bool(val.(bool))
	}

	return tps
}
