package sdkv2provider

import (
	"context"

	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess/models"
	"github.com/iwarapter/pingaccess-sdk-go/v62/services/sites"

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
		Schema:      resourcePingAccessSiteSchema(),
		Description: `Provides configuration for Sites within PingAccess.`,
	}
}

func resourcePingAccessSiteSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"availability_profile_id": {
			Type:        schema.TypeInt,
			Optional:    true,
			Default:     1,
			Description: "The ID of the availability profile associated with the site.",
		},
		"expected_hostname": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The name of the host expected in the site's certificate.",
		},
		"keep_alive_timeout": {
			Type:        schema.TypeInt,
			Optional:    true,
			Default:     0,
			Description: "The time, in milliseconds, that an HTTP persistent connection to the site can be idle before PingAccess closes the connection.",
		},
		"load_balancing_strategy_id": {
			Type:        schema.TypeInt,
			Optional:    true,
			Default:     0,
			Description: "The ID of the load balancing strategy associated with the site.",
		},
		"max_connections": {
			Type:        schema.TypeInt,
			Optional:    true,
			Default:     -1,
			Description: "The maximum number of HTTP persistent connections you want PingAccess to have open and maintain for the site. -1 indicates unlimited connections.",
		},
		"max_web_socket_connections": {
			Type:        schema.TypeInt,
			Optional:    true,
			Default:     -1,
			Description: "The maximum number of WebSocket connections you want PingAccess to have open and maintain for the site. -1 indicates unlimited connections.",
		},
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The name of the site.",
		},
		"secure": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
			Description: "This field is true if the site expects HTTPS connections.",
		},
		"send_pa_cookie": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     true,
			Description: "This field is true if the PingAccess Token or OAuth Access Token should be included in the request to the site.",
		},
		"site_authenticator_ids": {
			Type:        schema.TypeSet,
			Optional:    true,
			Description: "The IDs of the site authenticators associated with the site.",
			Elem: &schema.Schema{
				Type: schema.TypeInt,
			},
		},
		"skip_hostname_verification": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
			Description: "This field is true if the hostname verification of the site's certificate should be skipped.",
		},
		"targets": {
			Type:        schema.TypeSet,
			Required:    true,
			MinItems:    1,
			Description: "The {hostname}:{port} pairs for the hosts that make up the site.",
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"trusted_certificate_group_id": {
			Type:        schema.TypeInt,
			Optional:    true,
			Default:     0,
			Description: "The ID of the trusted certificate group associated with the site.",
		},
		"use_proxy": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
			Description: "True if a proxy should be used for HTTP or HTTPS requests.",
		},
		"use_target_host_header": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     true,
			Description: "Setting this field to true causes PingAccess to adjust the Host header to the site's selected target host rather than the virtual host configured in the application.",
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
