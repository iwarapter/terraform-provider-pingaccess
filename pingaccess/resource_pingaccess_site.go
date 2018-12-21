package pingaccess

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/iwarapter/pingaccess-sdk-go/pingaccess"
)

func resourcePingAccessSite() *schema.Resource {
	return &schema.Resource{
		Create: resourcePingAccessSiteCreate,
		Read:   resourcePingAccessSiteRead,
		Update: resourcePingAccessSiteUpdate,
		Delete: resourcePingAccessSiteDelete,

		Schema: map[string]*schema.Schema{
			"availability_profile_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"expected_hostname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"keep_alive_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"load_balancing_strategy_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"max_connections": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"max_web_socket_connections": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"secure": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"send_pa_cookie": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"site_authenticator_ids": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"skip_hostname_verification": &schema.Schema{
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
			"trusted_certificate_group_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"use_proxy": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"use_target_host_header": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
		},
	}
}

func resourcePingAccessSiteCreate(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] Start - resourcePingAccessSiteCreate")
	availability_profile_id := d.Get("availability_profile_id").(int)
	expected_hostname := d.Get("expected_hostname").(string)
	keep_alive_timeout := d.Get("keep_alive_timeout").(int)
	load_balancing_strategy_id := d.Get("load_balancing_strategy_id").(int)
	max_connections := d.Get("max_connections").(int)
	max_web_socket_connections := d.Get("max_web_socket_connections").(int)
	name := d.Get("name").(string)
	secure := d.Get("secure").(bool)
	send_pa_cookie := d.Get("send_pa_cookie").(bool)
	site_authenticator_ids := expandIntList(d.Get("site_authenticator_ids").(*schema.Set).List())
	skip_hostname_verification := d.Get("skip_hostname_verification").(bool)
	targets := expandStringList(d.Get("targets").(*schema.Set).List())
	trusted_certificate_group_id := d.Get("trusted_certificate_group_id").(int)
	use_proxy := d.Get("use_proxy").(bool)
	use_target_host_header := d.Get("use_target_host_header").(bool)

	input := pingaccess.AddSiteCommandInput{
		Body: pingaccess.SiteView{
			AvailabilityProfileId:   availability_profile_id,
			ExpectedHostname:        expected_hostname,
			KeepAliveTimeout:        keep_alive_timeout,
			LoadBalancingStrategyId: load_balancing_strategy_id,
			MaxConnections:          max_connections,
			MaxWebSocketConnections: max_web_socket_connections,
			Name:                      name,
			Secure:                    secure,
			SendPaCookie:              send_pa_cookie,
			SiteAuthenticatorIds:      site_authenticator_ids,
			SkipHostnameVerification:  skip_hostname_verification,
			Targets:                   targets,
			TrustedCertificateGroupId: trusted_certificate_group_id,
			UseProxy:                  use_proxy,
			UseTargetHostHeader:       use_target_host_header,
		},
	}

	svc := m.(*pingaccess.Client).Sites

	result, _, err := svc.AddSiteCommand(&input)
	if err != nil {
		return fmt.Errorf("Error creating virtualhost: %s", err)
	}

	d.SetId(strconv.Itoa(result.Id))
	log.Println("[DEBUG] End - resourcePingAccessSiteCreate")
	return resourcePingAccessSiteReadResult(d, &input.Body)
}

func resourcePingAccessSiteRead(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] Start - resourcePingAccessSiteRead")
	svc := m.(*pingaccess.Client).Sites

	input := &pingaccess.GetSiteCommandInput{
		Path: struct {
			Id string
		}{
			Id: d.Id(),
		},
	}

	log.Printf("[INFO] ResourceID: %s", d.Id())
	log.Printf("[INFO] GetSiteCommandInput: %s", input.Path.Id)
	result, _, _ := svc.GetSiteCommand(input)
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(result)
	vh := pingaccess.SiteView{}
	json.NewDecoder(b).Decode(&vh)

	log.Println("[DEBUG] End - resourcePingAccessSiteRead")
	return resourcePingAccessSiteReadResult(d, &vh)
}

func resourcePingAccessSiteUpdate(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] Start - resourcePingAccessSiteUpdate")
	availability_profile_id := d.Get("availability_profile_id").(int)
	expected_hostname := d.Get("expected_hostname").(string)
	keep_alive_timeout := d.Get("keep_alive_timeout").(int)
	load_balancing_strategy_id := d.Get("load_balancing_strategy_id").(int)
	max_connections := d.Get("max_connections").(int)
	max_web_socket_connections := d.Get("max_web_socket_connections").(int)
	name := d.Get("name").(string)
	secure := d.Get("secure").(bool)
	send_pa_cookie := d.Get("send_pa_cookie").(bool)
	site_authenticator_ids := expandIntList(d.Get("site_authenticator_ids").(*schema.Set).List())
	skip_hostname_verification := d.Get("skip_hostname_verification").(bool)
	targets := expandStringList(d.Get("targets").(*schema.Set).List())
	trusted_certificate_group_id := d.Get("trusted_certificate_group_id").(int)
	use_proxy := d.Get("use_proxy").(bool)
	use_target_host_header := d.Get("use_target_host_header").(bool)

	input := pingaccess.UpdateSiteCommandInput{
		Body: pingaccess.SiteView{
			AvailabilityProfileId:   availability_profile_id,
			ExpectedHostname:        expected_hostname,
			KeepAliveTimeout:        keep_alive_timeout,
			LoadBalancingStrategyId: load_balancing_strategy_id,
			MaxConnections:          max_connections,
			MaxWebSocketConnections: max_web_socket_connections,
			Name:                      name,
			Secure:                    secure,
			SendPaCookie:              send_pa_cookie,
			SiteAuthenticatorIds:      site_authenticator_ids,
			SkipHostnameVerification:  skip_hostname_verification,
			Targets:                   targets,
			TrustedCertificateGroupId: trusted_certificate_group_id,
			UseProxy:                  use_proxy,
			UseTargetHostHeader:       use_target_host_header,
		},
	}
	input.Path.Id = d.Id()

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
		Path: struct {
			Id string
		}{
			Id: d.Id(),
		},
	}

	log.Printf("[INFO] ResourceID: %s", d.Id())
	log.Printf("[INFO] DeleteSiteCommandInput: %s", input.Path.Id)
	_, err := svc.DeleteSiteCommand(input)
	if err != nil {
		return fmt.Errorf("Error deleting virtualhost: %s", err)
	}
	log.Println("[DEBUG] End - resourcePingAccessSiteDelete")
	return nil
}

func resourcePingAccessSiteReadResult(d *schema.ResourceData, input *pingaccess.SiteView) error {
	log.Println("[DEBUG] Start - resourcePingAccessSiteReadResult")
	if err := d.Set("availability_profile_id", input.AvailabilityProfileId); err != nil {
		return err
	}
	if err := d.Set("expected_hostname", input.ExpectedHostname); err != nil {
		return err
	}
	if err := d.Set("keep_alive_timeout", input.KeepAliveTimeout); err != nil {
		return err
	}
	if err := d.Set("load_balancing_strategy_id", input.LoadBalancingStrategyId); err != nil {
		return err
	}
	if err := d.Set("max_connections", input.MaxConnections); err != nil {
		return err
	}
	if err := d.Set("max_web_socket_connections", input.MaxWebSocketConnections); err != nil {
		return err
	}
	if err := d.Set("name", input.Name); err != nil {
		return err
	}
	if err := d.Set("secure", input.Secure); err != nil {
		return err
	}
	if err := d.Set("send_pa_cookie", input.SendPaCookie); err != nil {
		return err
	}
	if err := d.Set("site_authenticator_ids", input.SiteAuthenticatorIds); err != nil {
		return err
	}
	if err := d.Set("skip_hostname_verification", input.SkipHostnameVerification); err != nil {
		return err
	}
	if err := d.Set("targets", input.Targets); err != nil {
		return err
	}
	if err := d.Set("trusted_certificate_group_id", input.TrustedCertificateGroupId); err != nil {
		return err
	}
	if err := d.Set("use_proxy", input.UseProxy); err != nil {
		return err
	}
	if err := d.Set("use_target_host_header", input.UseTargetHostHeader); err != nil {
		return err
	}
	log.Println("[DEBUG] End - resourcePingAccessSiteReadResult")
	return nil
}
