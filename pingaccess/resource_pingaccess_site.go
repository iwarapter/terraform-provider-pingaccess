package pingaccess

import (
	"context"

	"github.com/iwarapter/pingaccess-sdk-go/pingaccess/models"
	"github.com/iwarapter/pingaccess-sdk-go/services/sites"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePingAccessSite() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePingAccessSiteCreate,
		ReadContext:   resourcePingAccessSiteRead,
		UpdateContext: resourcePingAccessSiteUpdate,
		DeleteContext: resourcePingAccessSiteDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: resourcePingAccessSiteSchema(),
	}
}

func resourcePingAccessSiteSchema() map[string]*schema.Schema {
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
		"keep_alive_timeout": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  0,
		},
		"load_balancing_strategy_id": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  0,
		},
		"max_connections": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  -1,
		},
		"max_web_socket_connections": {
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
		"send_pa_cookie": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"site_authenticator_ids": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeInt,
			},
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
		"use_target_host_header": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
	}
}

func resourcePingAccessSiteCreate(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(paClient).Sites
	input := sites.AddSiteCommandInput{
		Body: *resourcePingAccessSiteReadData(d),
	}

	result, _, err := svc.AddSiteCommand(&input)
	if err != nil {
		return diag.Errorf("unable to create Site: %s", err)
	}

	d.SetId(result.Id.String())
	return resourcePingAccessSiteReadResult(d, result)
}

func resourcePingAccessSiteRead(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(paClient).Sites
	input := &sites.GetSiteCommandInput{
		Id: d.Id(),
	}
	result, _, err := svc.GetSiteCommand(input)
	if err != nil {
		return diag.Errorf("unable to read Site: %s", err)
	}
	return resourcePingAccessSiteReadResult(d, result)
}

func resourcePingAccessSiteUpdate(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(paClient).Sites
	input := sites.UpdateSiteCommandInput{
		Body: *resourcePingAccessSiteReadData(d),
		Id:   d.Id(),
	}
	result, _, err := svc.UpdateSiteCommand(&input)
	if err != nil {
		return diag.Errorf("unable to update Site: %s", err)
	}
	return resourcePingAccessSiteReadResult(d, result)
}

func resourcePingAccessSiteDelete(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(paClient).Sites
	input := &sites.DeleteSiteCommandInput{
		Id: d.Id(),
	}

	_, err := svc.DeleteSiteCommand(input)
	if err != nil {
		return diag.Errorf("unable to delete Site: %s", err)
	}
	return nil
}

func resourcePingAccessSiteReadResult(d *schema.ResourceData, input *models.SiteView) diag.Diagnostics {
	var diags diag.Diagnostics
	setResourceDataIntWithDiagnostic(d, "availability_profile_id", input.AvailabilityProfileId, &diags)
	setResourceDataStringWithDiagnostic(d, "expected_hostname", input.ExpectedHostname, &diags)
	setResourceDataIntWithDiagnostic(d, "keep_alive_timeout", input.KeepAliveTimeout, &diags)
	setResourceDataIntWithDiagnostic(d, "load_balancing_strategy_id", input.LoadBalancingStrategyId, &diags)
	setResourceDataIntWithDiagnostic(d, "max_connections", input.MaxConnections, &diags)
	setResourceDataIntWithDiagnostic(d, "max_web_socket_connections", input.MaxWebSocketConnections, &diags)
	setResourceDataStringWithDiagnostic(d, "name", input.Name, &diags)
	setResourceDataBoolWithDiagnostic(d, "secure", input.Secure, &diags)
	setResourceDataBoolWithDiagnostic(d, "send_pa_cookie", input.SendPaCookie, &diags)
	if err := d.Set("site_authenticator_ids", input.SiteAuthenticatorIds); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	setResourceDataBoolWithDiagnostic(d, "skip_hostname_verification", input.SkipHostnameVerification, &diags)
	if err := d.Set("targets", input.Targets); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	setResourceDataIntWithDiagnostic(d, "trusted_certificate_group_id", input.TrustedCertificateGroupId, &diags)
	setResourceDataBoolWithDiagnostic(d, "use_proxy", input.UseProxy, &diags)
	setResourceDataBoolWithDiagnostic(d, "use_target_host_header", input.UseTargetHostHeader, &diags)
	return diags
}

func resourcePingAccessSiteReadData(d *schema.ResourceData) *models.SiteView {
	targets := expandStringList(d.Get("targets").(*schema.Set).List())
	site := &models.SiteView{
		Name:    String(d.Get("name").(string)),
		Targets: &targets,
	}
	site.AvailabilityProfileId = Int(d.Get("availability_profile_id").(int))
	site.ExpectedHostname = String(d.Get("expected_hostname").(string))
	site.KeepAliveTimeout = Int(d.Get("keep_alive_timeout").(int))
	site.LoadBalancingStrategyId = Int(d.Get("load_balancing_strategy_id").(int))
	site.MaxConnections = Int(d.Get("max_connections").(int))
	site.MaxWebSocketConnections = Int(d.Get("max_web_socket_connections").(int))
	site.Secure = Bool(d.Get("secure").(bool))
	site.SendPaCookie = Bool(d.Get("send_pa_cookie").(bool))
	siteAuthenticatorIds := expandIntList(d.Get("site_authenticator_ids").(*schema.Set).List())
	site.SiteAuthenticatorIds = &siteAuthenticatorIds
	site.SkipHostnameVerification = Bool(d.Get("skip_hostname_verification").(bool))
	site.TrustedCertificateGroupId = Int(d.Get("trusted_certificate_group_id").(int))
	site.UseProxy = Bool(d.Get("use_proxy").(bool))
	site.UseTargetHostHeader = Bool(d.Get("use_target_host_header").(bool))

	return site
}
