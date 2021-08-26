package sdkv2provider

import (
	"context"

	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess/models"
	"github.com/iwarapter/pingaccess-sdk-go/v62/services/thirdPartyServices"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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
		Schema:      resourcePingAccessThirdPartyServiceSchema(),
		Description: `Provides configuration for Third Party Services within PingAccess.`,
	}
}

func resourcePingAccessThirdPartyServiceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"availability_profile_id": {
			Type:        schema.TypeInt,
			Optional:    true,
			Default:     1,
			Description: "The ID of the availability profile associated with the third-party service.",
		},
		"expected_hostname": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The name of the host expected in the third-party service's certificate.",
		},
		"host_value": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The Host header field value in the requests sent to a Third-Party Services. When set, PingAccess will use the hostValue as the Host header field value. Otherwise, the target value will be used.",
		},
		"load_balancing_strategy_id": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "The ID of the load balancing strategy associated with the third-party service.",
		},
		"max_connections": {
			Type:        schema.TypeInt,
			Optional:    true,
			Default:     -1,
			Description: "The maximum number of HTTP persistent connections you want PingAccess to have open and maintain for the third-party service. -1 indicates unlimited connections.",
		},
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The name of the third-party service.",
		},
		"secure": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
			Description: "This field is true if the third-party service expects HTTPS connections.",
		},
		"skip_hostname_verification": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
			Description: "This field is true if the hostname verification of the third-party service's certificate should be skipped.",
		},
		"targets": {
			Type:        schema.TypeSet,
			Required:    true,
			MinItems:    1,
			Description: "The {hostname}:{port} pairs for the hosts that make up the third-party service.",
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"trusted_certificate_group_id": {
			Type:        schema.TypeInt,
			Optional:    true,
			Default:     0,
			Description: "The ID of the trusted certificate group associated with the third-party service.",
		},
		"use_proxy": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
			Description: "True if a proxy should be used for HTTP or HTTPS requests.",
		},
	}
}

func resourcePingAccessThirdPartyServiceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(paClient).ThirdPartyServices
	input := thirdPartyServices.AddThirdPartyServiceCommandInput{
		Body: *resourcePingAccessThirdPartyServiceReadData(d),
	}

	result, _, err := svc.AddThirdPartyServiceCommand(&input)
	if err != nil {
		return diag.Errorf("unable to create ThirdPartyService: %s", err)

	}

	d.SetId(*result.Id)
	return resourcePingAccessThirdPartyServiceReadResult(d, result)
}

func resourcePingAccessThirdPartyServiceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(paClient).ThirdPartyServices
	input := &thirdPartyServices.GetThirdPartyServiceCommandInput{
		Id: d.Id(),
	}
	result, _, err := svc.GetThirdPartyServiceCommand(input)
	if err != nil {
		return diag.Errorf("unable to read ThirdPartyService: %s", err)
	}
	return resourcePingAccessThirdPartyServiceReadResult(d, result)
}

func resourcePingAccessThirdPartyServiceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(paClient).ThirdPartyServices
	input := thirdPartyServices.UpdateThirdPartyServiceCommandInput{
		Body: *resourcePingAccessThirdPartyServiceReadData(d),
		Id:   d.Id(),
	}
	result, _, err := svc.UpdateThirdPartyServiceCommand(&input)
	if err != nil {
		return diag.Errorf("unable to update ThirdPartyService: %s", err)
	}
	return resourcePingAccessThirdPartyServiceReadResult(d, result)
}

func resourcePingAccessThirdPartyServiceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(paClient).ThirdPartyServices

	input := &thirdPartyServices.DeleteThirdPartyServiceCommandInput{
		Id: d.Id(),
	}

	_, err := svc.DeleteThirdPartyServiceCommand(input)
	if err != nil {
		return diag.Errorf("unable to delete ThirdPartyService: %s", err)
	}
	return nil
}

func resourcePingAccessThirdPartyServiceReadResult(d *schema.ResourceData, input *models.ThirdPartyServiceView) diag.Diagnostics {
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
		diags = append(diags, diag.FromErr(err)...)
	}
	return diags
}

func resourcePingAccessThirdPartyServiceReadData(d *schema.ResourceData) *models.ThirdPartyServiceView {
	targets := expandStringList(d.Get("targets").(*schema.Set).List())
	tps := &models.ThirdPartyServiceView{
		Name:    String(d.Get("name").(string)),
		Targets: &targets,
	}
	tps.AvailabilityProfileId = Int(d.Get("availability_profile_id").(int))
	if val, ok := d.GetOk("expected_hostname"); ok {
		tps.ExpectedHostname = String(val.(string))
	}
	if val, ok := d.GetOk("host_value"); ok {
		tps.HostValue = String(val.(string))
	}
	tps.LoadBalancingStrategyId = Int(d.Get("load_balancing_strategy_id").(int))
	tps.MaxConnections = Int(d.Get("max_connections").(int))
	tps.Secure = Bool(d.Get("secure").(bool))
	tps.SkipHostnameVerification = Bool(d.Get("skip_hostname_verification").(bool))
	tps.TrustedCertificateGroupId = Int(d.Get("trusted_certificate_group_id").(int))
	tps.UseProxy = Bool(d.Get("use_proxy").(bool))

	return tps
}
