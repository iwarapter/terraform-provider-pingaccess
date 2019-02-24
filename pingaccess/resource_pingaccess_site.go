package pingaccess

import (
	"crypto/tls"
	"fmt"
	"log"
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

		Schema: map[string]*schema.Schema{
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
			useTargetHostHeader: &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
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
		return fmt.Errorf("Error creating virtualhost: %s", err)
	}

	d.SetId(result.Id.String())
	return resourcePingAccessSiteReadResult(d, result)
}

func resourcePingAccessSiteRead(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pa.Client).Sites
	input := &pa.GetSiteCommandInput{
		Id: d.Id(),
	}
	result, _, _ := svc.GetSiteCommand(input)
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
	log.Println("[DEBUG] Start - resourcePingAccessSiteReadResult")
	if err := d.Set(availabilityProfileID, input.AvailabilityProfileId); err != nil {
		return err
	}
	if err := d.Set(expectedHostname, input.ExpectedHostname); err != nil {
		return err
	}
	if err := d.Set(keepAliveTimeout, input.KeepAliveTimeout); err != nil {
		return err
	}
	if err := d.Set(loadBalancingStrategyID, input.LoadBalancingStrategyId); err != nil {
		return err
	}
	if err := d.Set(maxConnections, input.MaxConnections); err != nil {
		return err
	}
	if err := d.Set(maxWebSocketConnections, input.MaxWebSocketConnections); err != nil {
		return err
	}
	if err := d.Set(name, input.Name); err != nil {
		return err
	}
	if err := d.Set(secure, input.Secure); err != nil {
		return err
	}
	if err := d.Set(sendPaCookie, input.SendPaCookie); err != nil {
		return err
	}
	if err := d.Set(siteAuthenticatorIDs, input.SiteAuthenticatorIds); err != nil {
		return err
	}
	if err := d.Set(skipHostnameVerification, input.SkipHostnameVerification); err != nil {
		return err
	}
	if err := d.Set(targets, input.Targets); err != nil {
		return err
	}
	if err := d.Set(trustedCertificateGroupID, input.TrustedCertificateGroupId); err != nil {
		return err
	}
	if err := d.Set(useProxy, input.UseProxy); err != nil {
		return err
	}
	if err := d.Set(useTargetHostHeader, input.UseTargetHostHeader); err != nil {
		return err
	}
	log.Println("[DEBUG] End - resourcePingAccessSiteReadResult")
	return nil
}

func resourcePingAccessSiteReadData(d *schema.ResourceData) *pa.SiteView {
	targets := expandStringList(d.Get(targets).(*schema.Set).List())
	site := &pa.SiteView{
		Name:    String(d.Get("name").(string)),
		Targets: &targets,
	}

	if _, ok := d.GetOk("availability_profile_id"); ok {
		site.AvailabilityProfileId = Int(d.Get("availability_profile_id").(int))
	}

	if _, ok := d.GetOk("expected_hostname"); ok {
		site.ExpectedHostname = String(d.Get("expected_hostname").(string))
	}

	if _, ok := d.GetOk("keep_alive_timeout"); ok {
		site.KeepAliveTimeout = Int(d.Get("keep_alive_timeout").(int))
	}

	if _, ok := d.GetOk("load_balancing_strategy_id"); ok {
		site.KeepAliveTimeout = Int(d.Get("load_balancing_strategy_id").(int))
	}

	if _, ok := d.GetOk("max_connections"); ok {
		site.MaxConnections = Int(d.Get("max_connections").(int))
	}

	if _, ok := d.GetOk("max_web_socket_connections"); ok {
		site.MaxWebSocketConnections = Int(d.Get("max_web_socket_connections").(int))
	}

	if _, ok := d.GetOk("secure"); ok {
		site.Secure = Bool(d.Get("secure").(bool))
	}

	if _, ok := d.GetOk("send_pa_cookie"); ok {
		site.SendPaCookie = Bool(d.Get("send_pa_cookie").(bool))
	}

	if _, ok := d.GetOk("site_authenticator_ids"); ok {
		siteAuthenticatorIds := expandIntList(d.Get(siteAuthenticatorIDs).(*schema.Set).List())
		site.SiteAuthenticatorIds = &siteAuthenticatorIds
	}

	if _, ok := d.GetOk("skip_hostname_verification"); ok {
		site.SkipHostnameVerification = Bool(d.Get("skip_hostname_verification").(bool))
	}

	if _, ok := d.GetOk("trusted_certificate_group_id"); ok {
		site.TrustedCertificateGroupId = Int(d.Get("trusted_certificate_group_id").(int))
	}

	if _, ok := d.GetOk("use_proxy"); ok {
		site.UseProxy = Bool(d.Get("use_proxy").(bool))
	}

	if _, ok := d.GetOk("use_target_host_header"); ok {
		site.UseTargetHostHeader = Bool(d.Get("use_target_host_header").(bool))
	}

	return site
}
