package pingaccess

import (
	"fmt"

	pa "github.com/iwarapter/pingaccess-sdk-go/pingaccess"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourcePingAccessPingFederateRuntime() *schema.Resource {
	return &schema.Resource{
		Create: resourcePingAccessPingFederateRuntimeCreate,
		Read:   resourcePingAccessPingFederateRuntimeRead,
		Update: resourcePingAccessPingFederateRuntimeUpdate,
		Delete: resourcePingAccessPingFederateRuntimeDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: resourcePingAccessPingFederateRuntimeSchema(),
	}
}

func resourcePingAccessPingFederateRuntimeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"audit_level": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Default:  "ON",
		},
		"back_channel_base_path": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"back_channel_secure": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"base_path": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"expected_hostname": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"host": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"port": &schema.Schema{
			Type:     schema.TypeInt,
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
			Default:  true,
		},
		"targets": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			MinItems: 0,
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
			Default:  false,
		},
		"use_slo": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
	}
}

func resourcePingAccessPingFederateRuntimeCreate(d *schema.ResourceData, m interface{}) error {
	d.SetId("pingfederate_runtime_settings")
	return resourcePingAccessPingFederateRuntimeUpdate(d, m)
}

func resourcePingAccessPingFederateRuntimeRead(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pa.Client).PingFederate
	result, _, err := svc.GetPingFederateCommand()
	if err != nil {
		return fmt.Errorf("Error reading pingfederate runtime settings: %s", err)
	}

	return resourcePingAccessPingFederateRuntimeReadResult(d, result)
}

func resourcePingAccessPingFederateRuntimeUpdate(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pa.Client).PingFederate
	input := pa.UpdatePingFederateCommandInput{
		Body: *resourcePingAccessPingFederateRuntimeReadData(d),
	}
	result, _, err := svc.UpdatePingFederateCommand(&input)
	if err != nil {
		return fmt.Errorf("Error updating pingfederate runtime settings: %s", err)
	}

	d.SetId("pingfederate_runtime_settings")
	return resourcePingAccessPingFederateRuntimeReadResult(d, result)
}

func resourcePingAccessPingFederateRuntimeDelete(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pa.Client).PingFederate
	_, err := svc.DeletePingFederateCommand()
	if err != nil {
		return fmt.Errorf("Error resetting pingfederate runtime: %s", err)
	}
	return nil
}

func resourcePingAccessPingFederateRuntimeReadResult(d *schema.ResourceData, input *pa.PingFederateRuntimeView) error {
	setResourceDataString(d, "audit_level", input.AuditLevel)
	setResourceDataString(d, "back_channel_base_path", input.BackChannelBasePath)
	setResourceDataBool(d, "back_channel_secure", input.BackChannelSecure)
	setResourceDataString(d, "base_path", input.BasePath)
	setResourceDataString(d, "expected_hostname", input.ExpectedHostname)
	setResourceDataString(d, "host", input.Host)
	setResourceDataInt(d, "port", input.Port)
	setResourceDataBool(d, "secure", input.Secure)
	setResourceDataBool(d, "skip_hostname_verification", input.SkipHostnameVerification)
	setResourceDataInt(d, "trusted_certificate_group_id", input.TrustedCertificateGroupId)
	setResourceDataBool(d, "use_proxy", input.UseProxy)
	setResourceDataBool(d, "use_slo", input.UseSlo)

	if err := d.Set("targets", input.Targets); err != nil {
		return err
	}

	return nil
}

func resourcePingAccessPingFederateRuntimeReadData(d *schema.ResourceData) *pa.PingFederateRuntimeView {
	pfRuntime := &pa.PingFederateRuntimeView{
		Host: String(d.Get("host").(string)),
		Port: Int(d.Get("port").(int)),
	}

	if v, ok := d.GetOkExists("audit_level"); ok {
		pfRuntime.AuditLevel = String(v.(string))
	}

	if v, ok := d.GetOkExists("back_channel_base_path"); ok {
		pfRuntime.BackChannelBasePath = String(v.(string))
	}

	if v, ok := d.GetOkExists("back_channel_secure"); ok {
		pfRuntime.BackChannelSecure = Bool(v.(bool))
	}

	if v, ok := d.GetOkExists("base_path"); ok {
		pfRuntime.BasePath = String(v.(string))
	}

	if v, ok := d.GetOkExists("expected_hostname"); ok {
		pfRuntime.ExpectedHostname = String(v.(string))
	}

	if v, ok := d.GetOkExists("secure"); ok {
		pfRuntime.Secure = Bool(v.(bool))
	}

	if v, ok := d.GetOkExists("skip_hostname_verification"); ok {
		pfRuntime.SkipHostnameVerification = Bool(v.(bool))
	}

	if v, ok := d.GetOkExists("targets"); ok {
		targets := expandStringList(v.(*schema.Set).List())
		pfRuntime.Targets = &targets
	}

	if v, ok := d.GetOkExists("trusted_certificate_group_id"); ok {
		pfRuntime.TrustedCertificateGroupId = Int(v.(int))
	}

	if v, ok := d.GetOkExists("use_proxy"); ok {
		pfRuntime.UseProxy = Bool(v.(bool))
	}

	if v, ok := d.GetOkExists("use_slo"); ok {
		pfRuntime.UseSlo = Bool(v.(bool))
	}

	return pfRuntime
}
