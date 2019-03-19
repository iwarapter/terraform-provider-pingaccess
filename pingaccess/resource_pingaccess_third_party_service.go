package pingaccess

import (
	"crypto/tls"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform/helper/schema"
	pa "github.com/iwarapter/pingaccess-sdk-go/pingaccess"
)

func resourcePingAccessThirdPartyService() *schema.Resource {
	return &schema.Resource{
		Create: resourcePingAccessThirdPartyServiceCreate,
		Read:   resourcePingAccessThirdPartyServiceRead,
		Update: resourcePingAccessThirdPartyServiceUpdate,
		Delete: resourcePingAccessThirdPartyServiceDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: resourcePingAccessThirdPartyServiceSchema(),
	}
}

func resourcePingAccessThirdPartyServiceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"availability_profile_id": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  1,
		},
		"expected_hostname": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"host_value": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"load_balancing_strategy_id": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
		},
		"max_connections": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  -1,
		},
		"name": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"secure": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"skip_hostname_verification": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"targets": &schema.Schema{
			Type:     schema.TypeSet,
			Required: true,
			MinItems: 1,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"trusted_certificate_group_id": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  0,
		},
		"use_proxy": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
	}
}

func resourcePingAccessThirdPartyServiceCreate(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pa.Client).ThirdPartyServices
	input := pa.AddThirdPartyServiceCommandInput{
		Body: *resourcePingAccessThirdPartyServiceReadData(d),
	}

	result, _, err := svc.AddThirdPartyServiceCommand(&input)
	if err != nil {
		return fmt.Errorf("Error creating third party service: %s", err)
	}

	d.SetId(result.Id.String())
	return resourcePingAccessThirdPartyServiceReadResult(d, result)
}

func resourcePingAccessThirdPartyServiceRead(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pa.Client).ThirdPartyServices
	input := &pa.GetThirdPartyServiceCommandInput{
		Id: d.Id(),
	}
	result, _, err := svc.GetThirdPartyServiceCommand(input)
	if err != nil {
		return fmt.Errorf("Error reading third party service: %s", err)
	}
	return resourcePingAccessThirdPartyServiceReadResult(d, result)
}

func resourcePingAccessThirdPartyServiceUpdate(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pa.Client).ThirdPartyServices
	input := pa.UpdateThirdPartyServiceCommandInput{
		Body: *resourcePingAccessThirdPartyServiceReadData(d),
		Id:   d.Id(),
	}
	result, _, err := svc.UpdateThirdPartyServiceCommand(&input)
	if err != nil {
		return fmt.Errorf("Error updating third party service: %s", err)
	}
	return resourcePingAccessThirdPartyServiceReadResult(d, result)
}

func resourcePingAccessThirdPartyServiceDelete(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pa.Client).ThirdPartyServices
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	input := &pa.DeleteThirdPartyServiceCommandInput{
		Id: d.Id(),
	}

	_, err := svc.DeleteThirdPartyServiceCommand(input)
	if err != nil {
		return fmt.Errorf("Error deleting third party service: %s", err)
	}
	return nil
}

func resourcePingAccessThirdPartyServiceReadResult(d *schema.ResourceData, input *pa.ThirdPartyServiceView) (err error) {
	setResourceDataInt(d, "availability_profile_id", input.AvailabilityProfileId)
	setResourceDataString(d, "expected_hostname", input.ExpectedHostname)
	setResourceDataString(d, "host_value", input.HostValue)
	setResourceDataInt(d, "load_balancing_strategy_id", input.LoadBalancingStrategyId)
	setResourceDataInt(d, "max_connections", input.MaxConnections)
	setResourceDataString(d, "name", input.Name)
	setResourceDataBool(d, "secure", input.Secure)
	setResourceDataBool(d, "skip_hostname_verification", input.SkipHostnameVerification)
	setResourceDataInt(d, "trusted_certificate_group_id", input.TrustedCertificateGroupId)
	setResourceDataBool(d, "use_proxy", input.UseProxy)
	if err := d.Set("targets", input.Targets); err != nil {
		return err
	}
	return nil
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
