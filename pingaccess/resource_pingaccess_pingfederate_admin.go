package pingaccess

import (
	"fmt"

	pa "github.com/iwarapter/pingaccess-sdk-go/pingaccess"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePingAccessPingFederateAdmin() *schema.Resource {
	return &schema.Resource{
		Create: resourcePingAccessPingFederateAdminCreate,
		Read:   resourcePingAccessPingFederateAdminRead,
		Update: resourcePingAccessPingFederateAdminUpdate,
		Delete: resourcePingAccessPingFederateAdminDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: resourcePingAccessPingFederateAdminSchema(),
	}
}

func resourcePingAccessPingFederateAdminSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"admin_password": requiredHiddenField(),
		"admin_username": {
			Type:     schema.TypeString,
			Required: true,
		},
		"audit_level": {
			Type:         schema.TypeString,
			Optional:     true,
			ValidateFunc: validateAuditLevel,
		},
		"base_path": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"host": {
			Type:     schema.TypeString,
			Required: true,
		},
		"port": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"secure": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"trusted_certificate_group_id": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"use_proxy": {
			Type:     schema.TypeBool,
			Optional: true,
		},
	}
}

func resourcePingAccessPingFederateAdminCreate(d *schema.ResourceData, m interface{}) error {
	d.SetId("pingfederate_admin_settings")
	return resourcePingAccessPingFederateAdminUpdate(d, m)
}

func resourcePingAccessPingFederateAdminRead(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pa.Client).PingFederate
	result, _, err := svc.GetPingFederateAdminCommand()
	if err != nil {
		return fmt.Errorf("Error reading pingfederate Admin settings: %s", err)
	}

	return resourcePingAccessPingFederateAdminReadResult(d, result)
}

func resourcePingAccessPingFederateAdminUpdate(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pa.Client).PingFederate
	input := pa.UpdatePingFederateAdminCommandInput{
		Body: *resourcePingAccessPingFederateAdminReadData(d),
	}
	result, _, err := svc.UpdatePingFederateAdminCommand(&input)
	if err != nil {
		return fmt.Errorf("Error updating pingfederate Admin settings: %s", err)
	}

	d.SetId("pingfederate_admin_settings")
	return resourcePingAccessPingFederateAdminReadResult(d, result)
}

func resourcePingAccessPingFederateAdminDelete(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pa.Client).PingFederate
	_, err := svc.DeletePingFederateCommand()
	if err != nil {
		return fmt.Errorf("Error resetting pingfederate Admin: %s", err)
	}
	return nil
}

func resourcePingAccessPingFederateAdminReadResult(d *schema.ResourceData, input *pa.PingFederateAdminView) error {
	if err := setResourceDataString(d, "admin_username", input.AdminUsername); err != nil {
		return err
	}
	if err := setResourceDataString(d, "audit_level", input.AuditLevel); err != nil {
		return err
	}
	if err := setResourceDataString(d, "base_path", input.BasePath); err != nil {
		return err
	}
	if err := setResourceDataString(d, "host", input.Host); err != nil {
		return err
	}
	if err := setResourceDataInt(d, "port", input.Port); err != nil {
		return err
	}
	if err := setResourceDataBool(d, "secure", input.Secure); err != nil {
		return err
	}
	if err := setResourceDataInt(d, "trusted_certificate_group_id", input.TrustedCertificateGroupId); err != nil {
		return err
	}
	if err := setResourceDataBool(d, "use_proxy", input.UseProxy); err != nil {
		return err
	}

	if input.AdminPassword != nil {
		pw, ok := d.GetOkExists("admin_password.0.value")
		creds := flattenHiddenFieldView(input.AdminPassword)
		if ok {
			creds[0]["value"] = pw
		}
		if err := d.Set("admin_password", creds); err != nil {
			return err
		}
	}

	return nil
}

func resourcePingAccessPingFederateAdminReadData(d *schema.ResourceData) *pa.PingFederateAdminView {
	pfAdmin := &pa.PingFederateAdminView{
		AdminUsername: String(d.Get("admin_username").(string)),
		Host:          String(d.Get("host").(string)),
		Port:          Int(d.Get("port").(int)),
	}
	if v, ok := d.GetOkExists("admin_password"); ok {
		pfAdmin.AdminPassword = expandHiddenFieldView(v.([]interface{}))
	}
	if v, ok := d.GetOkExists("audit_level"); ok {
		pfAdmin.AuditLevel = String(v.(string))
	}
	if v, ok := d.GetOkExists("base_path"); ok {
		pfAdmin.BasePath = String(v.(string))
	}
	if v, ok := d.GetOkExists("secure"); ok {
		pfAdmin.Secure = Bool(v.(bool))
	}
	if v, ok := d.GetOkExists("trusted_certificate_group_id"); ok {
		pfAdmin.TrustedCertificateGroupId = Int(v.(int))
	}
	if v, ok := d.GetOkExists("use_proxy"); ok {
		pfAdmin.UseProxy = Bool(v.(bool))
	}
	return pfAdmin
}
