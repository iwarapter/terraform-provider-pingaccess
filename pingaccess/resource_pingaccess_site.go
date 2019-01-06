package pingaccess

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/iwarapter/pingaccess-sdk-go/pingaccess"
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
			targets: &schema.Schema{
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
	log.Println("[DEBUG] Start - resourcePingAccessSiteCreate")
	availability_profile_id := d.Get(availabilityProfileID).(int)
	expected_hostname := d.Get(expectedHostname).(string)
	keep_alive_timeout := d.Get(keepAliveTimeout).(int)
	load_balancing_strategy_id := d.Get(loadBalancingStrategyID).(int)
	max_connections := d.Get(maxConnections).(int)
	max_web_socket_connections := d.Get(maxWebSocketConnections).(int)
	name := d.Get(name).(string)
	secure := d.Get(secure).(bool)
	send_pa_cookie := d.Get(sendPaCookie).(bool)
	site_authenticator_ids := expandIntList(d.Get(siteAuthenticatorIDs).(*schema.Set).List())
	skip_hostname_verification := d.Get(skipHostnameVerification).(bool)
	targets := expandStringList(d.Get(targets).(*schema.Set).List())
	trusted_certificate_group_id := d.Get(trustedCertificateGroupID).(int)
	use_proxy := d.Get(useProxy).(bool)
	use_target_host_header := d.Get(useTargetHostHeader).(bool)

	input := pingaccess.AddSiteCommandInput{
		Body: pingaccess.SiteView{
			AvailabilityProfileId:   Int(availability_profile_id),
			ExpectedHostname:        String(expected_hostname),
			KeepAliveTimeout:        Int(keep_alive_timeout),
			LoadBalancingStrategyId: Int(load_balancing_strategy_id),
			MaxConnections:          Int(max_connections),
			MaxWebSocketConnections: Int(max_web_socket_connections),
			Name:                      String(name),
			Secure:                    Bool(secure),
			SendPaCookie:              Bool(send_pa_cookie),
			SiteAuthenticatorIds:      &site_authenticator_ids,
			SkipHostnameVerification:  Bool(skip_hostname_verification),
			Targets:                   &targets,
			TrustedCertificateGroupId: Int(trusted_certificate_group_id),
			UseProxy:                  Bool(use_proxy),
			UseTargetHostHeader:       Bool(use_target_host_header),
		},
	}

	svc := m.(*pingaccess.Client).Sites

	result, _, err := svc.AddSiteCommand(&input)
	if err != nil {
		return fmt.Errorf("Error creating virtualhost: %s", err)
	}

	d.SetId(result.Id.String())
	log.Println("[DEBUG] End - resourcePingAccessSiteCreate")
	return resourcePingAccessSiteReadResult(d, result)
}

func resourcePingAccessSiteRead(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] Start - resourcePingAccessSiteRead")
	svc := m.(*pingaccess.Client).Sites

	input := &pingaccess.GetSiteCommandInput{
		Id: d.Id(),
	}

	log.Printf("[INFO] ResourceID: %s", d.Id())
	log.Printf("[INFO] GetSiteCommandInput: %s", input.Id)
	result, _, _ := svc.GetSiteCommand(input)

	log.Println("[DEBUG] End - resourcePingAccessSiteRead")
	return resourcePingAccessSiteReadResult(d, result)
}

func resourcePingAccessSiteUpdate(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] Start - resourcePingAccessSiteUpdate")
	availability_profile_id := d.Get(availabilityProfileID).(int)
	expected_hostname := d.Get(expectedHostname).(string)
	keep_alive_timeout := d.Get(keepAliveTimeout).(int)
	load_balancing_strategy_id := d.Get(loadBalancingStrategyID).(int)
	max_connections := d.Get(maxConnections).(int)
	max_web_socket_connections := d.Get(maxWebSocketConnections).(int)
	name := d.Get(name).(string)
	secure := d.Get(secure).(bool)
	send_pa_cookie := d.Get(sendPaCookie).(bool)
	site_authenticator_ids := expandIntList(d.Get(siteAuthenticatorIDs).(*schema.Set).List())
	skip_hostname_verification := d.Get(skipHostnameVerification).(bool)
	targets := expandStringList(d.Get(targets).(*schema.Set).List())
	trusted_certificate_group_id := d.Get(trustedCertificateGroupID).(int)
	use_proxy := d.Get(useProxy).(bool)
	use_target_host_header := d.Get(useTargetHostHeader).(bool)

	input := pingaccess.UpdateSiteCommandInput{
		Body: pingaccess.SiteView{
			AvailabilityProfileId:   Int(availability_profile_id),
			ExpectedHostname:        String(expected_hostname),
			KeepAliveTimeout:        Int(keep_alive_timeout),
			LoadBalancingStrategyId: Int(load_balancing_strategy_id),
			MaxConnections:          Int(max_connections),
			MaxWebSocketConnections: Int(max_web_socket_connections),
			Name:                      String(name),
			Secure:                    Bool(secure),
			SendPaCookie:              Bool(send_pa_cookie),
			SiteAuthenticatorIds:      &site_authenticator_ids,
			SkipHostnameVerification:  Bool(skip_hostname_verification),
			Targets:                   &targets,
			TrustedCertificateGroupId: Int(trusted_certificate_group_id),
			UseProxy:                  Bool(use_proxy),
			UseTargetHostHeader:       Bool(use_target_host_header),
		},
	}
	input.Id = d.Id()

	svc := m.(*pingaccess.Client).Sites
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	_, _, err := svc.UpdateSiteCommand(&input)
	if err != nil {
		return fmt.Errorf("Error updating virtualhost: %s", err)
	}
	log.Println("[DEBUG] End - resourcePingAccessSiteUpdate")
	return nil
}

func resourcePingAccessSiteDelete(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] Start - resourcePingAccessSiteDelete")
	svc := m.(*pingaccess.Client).Sites
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	input := &pingaccess.DeleteSiteCommandInput{
		Id: d.Id(),
	}

	log.Printf("[INFO] ResourceID: %s", d.Id())
	log.Printf("[INFO] DeleteSiteCommandInput: %s", input.Id)
	_, err := svc.DeleteSiteCommand(input)
	if err != nil {
		return fmt.Errorf("Error deleting virtualhost: %s", err)
	}
	log.Println("[DEBUG] End - resourcePingAccessSiteDelete")
	return nil
}

func resourcePingAccessSiteReadResult(d *schema.ResourceData, input *pingaccess.SiteView) error {
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
