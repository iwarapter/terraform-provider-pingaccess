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
		"description": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"issuer": {
			Type:     schema.TypeString,
			Required: true,
		},
		"skip_hostname_verification": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"sts_token_exchange_endpoint": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"trusted_certificate_group_id": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"use_proxy": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"use_slo": {
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
	result, _, err := svc.GetPingFederateRuntimeCommand()
	if err != nil {
		return fmt.Errorf("Error reading pingfederate runtime settings: %s", err)
	}

	return resourcePingAccessPingFederateRuntimeReadResult(d, result)
}

func resourcePingAccessPingFederateRuntimeUpdate(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pa.Client).PingFederate
	input := pa.UpdatePingFederateRuntimeCommandInput{
		Body: *resourcePingAccessPingFederateRuntimeReadData(d),
	}
	result, _, err := svc.UpdatePingFederateRuntimeCommand(&input)
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

func resourcePingAccessPingFederateRuntimeReadResult(d *schema.ResourceData, input *pa.PingFederateMetadataRuntimeView) error {
	if err := setResourceDataString(d, "description", input.Description); err != nil {
		return err
	}
	if err := setResourceDataString(d, "issuer", input.Issuer); err != nil {
		return err
	}
	if err := setResourceDataBool(d, "skip_hostname_verification", input.SkipHostnameVerification); err != nil {
		return err
	}
	if err := setResourceDataString(d, "sts_token_exchange_endpoint", input.StsTokenExchangeEndpoint); err != nil {
		return err
	}
	if err := setResourceDataInt(d, "trusted_certificate_group_id", input.TrustedCertificateGroupId); err != nil {
		return err
	}
	if err := setResourceDataBool(d, "use_proxy", input.UseProxy); err != nil {
		return err
	}
	if err := setResourceDataBool(d, "use_slo", input.UseSlo); err != nil {
		return err
	}

	return nil
}

func resourcePingAccessPingFederateRuntimeReadData(d *schema.ResourceData) *pa.PingFederateMetadataRuntimeView {
	pfRuntime := &pa.PingFederateMetadataRuntimeView{
		Issuer: String(d.Get("issuer").(string)),
	}

	if v, ok := d.GetOkExists("description"); ok {
		pfRuntime.Description = String(v.(string))
	}

	if v, ok := d.GetOkExists("skip_hostname_verification"); ok {
		pfRuntime.SkipHostnameVerification = Bool(v.(bool))
	}

	if v, ok := d.GetOkExists("sts_token_exchange_endpoint"); ok {
		pfRuntime.StsTokenExchangeEndpoint = String(v.(string))
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
