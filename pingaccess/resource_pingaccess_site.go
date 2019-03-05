package pingaccess

import (
	"crypto/tls"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform/helper/schema"
	pa "github.com/iwarapter/pingaccess-sdk-go/pingaccess"
)

func resourcePingAccessSite() *schema.Resource {
	return &schema.Resource{
		Create: resourcePingAccessSiteCreate,
		Read:   resourcePingAccessSiteRead,
		Update: resourcePingAccessSiteUpdate,
		Delete: resourcePingAccessSiteDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: resourcePingAccessSiteSchema(),
	}
}

func resourcePingAccessSiteSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		availabilityProfileID: &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
		},
		expectedHostname: &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		keepAliveTimeout: &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
		},
		loadBalancingStrategyID: &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
		},
		maxConnections: &schema.Schema{
			Type:     schema.TypeInt,
			Required: true,
		},
		maxWebSocketConnections: &schema.Schema{
			Type:     schema.TypeInt,
			Required: true,
		},
		name: &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		secure: &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
		},
		sendPaCookie: &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		siteAuthenticatorIDs: &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeInt,
			},
		},
		skipHostnameVerification: &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
		},
		"targets": &schema.Schema{
			Type:     schema.TypeSet,
			Required: true,
			MinItems: 1,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		trustedCertificateGroupID: &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
		},
		useProxy: &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
		},
		"use_target_host_header": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
	}
}

func resourcePingAccessSiteCreate(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pa.Client).Sites
	input := pa.AddSiteCommandInput{
		Body: *resourcePingAccessSiteReadData(d),
	}

	result, _, err := svc.AddSiteCommand(&input)
	if err != nil {
		return fmt.Errorf("Error creating site: %s", err)
	}

	d.SetId(result.Id.String())
	return resourcePingAccessSiteReadResult(d, result)
}

func resourcePingAccessSiteRead(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pa.Client).Sites
	input := &pa.GetSiteCommandInput{
		Id: d.Id(),
	}
	result, _, err := svc.GetSiteCommand(input)
	if err != nil {
		return fmt.Errorf("Error reading site: %s", err)
	}
	return resourcePingAccessSiteReadResult(d, result)
}

func resourcePingAccessSiteUpdate(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pa.Client).Sites
	input := pa.UpdateSiteCommandInput{
		Body: *resourcePingAccessSiteReadData(d),
		Id:   d.Id(),
	}
	result, _, err := svc.UpdateSiteCommand(&input)
	if err != nil {
		return fmt.Errorf("Error updating virtualhost: %s", err)
	}
	return resourcePingAccessSiteReadResult(d, result)
}

func resourcePingAccessSiteDelete(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pa.Client).Sites
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	input := &pa.DeleteSiteCommandInput{
		Id: d.Id(),
	}

	_, err := svc.DeleteSiteCommand(input)
	if err != nil {
		return fmt.Errorf("Error deleting virtualhost: %s", err)
	}
	return nil
}

func resourcePingAccessSiteReadResult(d *schema.ResourceData, input *pa.SiteView) error {
	setResourceDataInt(d, "availability_profile_id", input.AvailabilityProfileId)
	setResourceDataString(d, "expected_hostname", input.ExpectedHostname)
	setResourceDataInt(d, "keep_alive_timeout", input.KeepAliveTimeout)
	setResourceDataInt(d, "load_balancing_strategy_id", input.LoadBalancingStrategyId)
	setResourceDataInt(d, "max_connections", input.MaxConnections)
	setResourceDataInt(d, "max_web_socket_connections", input.MaxWebSocketConnections)
	setResourceDataString(d, "name", input.Name)
	setResourceDataBool(d, "secure", input.Secure)
	setResourceDataBool(d, "send_pa_cookie", input.SendPaCookie)
	if err := d.Set("site_authenticator_ids", input.SiteAuthenticatorIds); err != nil {
		return err
	}
	setResourceDataBool(d, "skip_hostname_verification", input.SkipHostnameVerification)
	if err := d.Set("targets", input.Targets); err != nil {
		return err
	}
	setResourceDataInt(d, "trusted_certificate_group_id", input.TrustedCertificateGroupId)
	setResourceDataBool(d, "use_proxy", input.UseProxy)
	setResourceDataBool(d, "use_target_host_header", input.UseTargetHostHeader)
	return nil
}

func resourcePingAccessSiteReadData(d *schema.ResourceData) *pa.SiteView {
	targets := expandStringList(d.Get(targets).(*schema.Set).List())
	site := &pa.SiteView{
		Name:    String(d.Get("name").(string)),
		Targets: &targets,
	}

	if _, ok := d.GetOkExists("availability_profile_id"); ok {
		site.AvailabilityProfileId = Int(d.Get("availability_profile_id").(int))
	}

	if _, ok := d.GetOkExists("expected_hostname"); ok {
		site.ExpectedHostname = String(d.Get("expected_hostname").(string))
	}

	if _, ok := d.GetOkExists("keep_alive_timeout"); ok {
		site.KeepAliveTimeout = Int(d.Get("keep_alive_timeout").(int))
	}

	if _, ok := d.GetOkExists("load_balancing_strategy_id"); ok {
		site.LoadBalancingStrategyId = Int(d.Get("load_balancing_strategy_id").(int))
	}

	if _, ok := d.GetOkExists("max_connections"); ok {
		site.MaxConnections = Int(d.Get("max_connections").(int))
	}

	if _, ok := d.GetOkExists("max_web_socket_connections"); ok {
		site.MaxWebSocketConnections = Int(d.Get("max_web_socket_connections").(int))
	}

	if _, ok := d.GetOkExists("secure"); ok {
		site.Secure = Bool(d.Get("secure").(bool))
	}

	if _, ok := d.GetOkExists("send_pa_cookie"); ok {
		site.SendPaCookie = Bool(d.Get("send_pa_cookie").(bool))
	}

	if _, ok := d.GetOkExists("site_authenticator_ids"); ok {
		siteAuthenticatorIds := expandIntList(d.Get(siteAuthenticatorIDs).(*schema.Set).List())
		site.SiteAuthenticatorIds = &siteAuthenticatorIds
	}

	if _, ok := d.GetOkExists("skip_hostname_verification"); ok {
		site.SkipHostnameVerification = Bool(d.Get("skip_hostname_verification").(bool))
	}

	if _, ok := d.GetOkExists("trusted_certificate_group_id"); ok {
		site.TrustedCertificateGroupId = Int(d.Get("trusted_certificate_group_id").(int))
	}

	if _, ok := d.GetOkExists("use_proxy"); ok {
		site.UseProxy = Bool(d.Get("use_proxy").(bool))
	}

	if _, ok := d.GetOkExists("use_target_host_header"); ok {
		site.UseTargetHostHeader = Bool(d.Get("use_target_host_header").(bool))
	}

	return site
}
