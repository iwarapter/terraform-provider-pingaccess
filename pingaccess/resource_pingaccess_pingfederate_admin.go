package pingaccess

import (
	"context"
	"fmt"

	pa "github.com/iwarapter/pingaccess-sdk-go/pingaccess"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePingAccessPingFederateAdmin() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePingAccessPingFederateAdminCreate,
		ReadContext:   resourcePingAccessPingFederateAdminRead,
		UpdateContext: resourcePingAccessPingFederateAdminUpdate,
		DeleteContext: resourcePingAccessPingFederateAdminDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
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
			Type:             schema.TypeString,
			Optional:         true,
			ValidateDiagFunc: validateAuditLevel,
			Default: "ON",
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
			Default: 0,
		},
		"use_proxy": {
			Type:     schema.TypeBool,
			Optional: true,
			Default: false,
		},
	}
}

func resourcePingAccessPingFederateAdminCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	d.SetId("pingfederate_admin_settings")
	return resourcePingAccessPingFederateAdminUpdate(ctx, d, m)
}

func resourcePingAccessPingFederateAdminRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(*pa.Client).PingFederate
	result, _, err := svc.GetPingFederateAdminCommand()
	if err != nil {
		return diag.Diagnostics{diag.FromErr(fmt.Errorf("unable to read PingFederateAdmin: %s", err))}
	}

	return resourcePingAccessPingFederateAdminReadResult(d, result)
}

func resourcePingAccessPingFederateAdminUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(*pa.Client).PingFederate
	input := pa.UpdatePingFederateAdminCommandInput{
		Body: *resourcePingAccessPingFederateAdminReadData(d),
	}
	result, _, err := svc.UpdatePingFederateAdminCommand(&input)
	if err != nil {
		return diag.Diagnostics{diag.FromErr(fmt.Errorf("unable to update PingFederateAdmin: %s", err))}
	}

	d.SetId("pingfederate_admin_settings")
	return resourcePingAccessPingFederateAdminReadResult(d, result)
}

func resourcePingAccessPingFederateAdminDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(*pa.Client).PingFederate
	_, err := svc.DeletePingFederateCommand()
	if err != nil {
		return diag.Diagnostics{diag.FromErr(fmt.Errorf("unable to reset PingFederateAdmin: %s", err))}
	}
	return nil
}

func resourcePingAccessPingFederateAdminReadResult(d *schema.ResourceData, input *pa.PingFederateAdminView) diag.Diagnostics {
	var diags diag.Diagnostics
	setResourceDataStringWithDiagnostic(d, "admin_username", input.AdminUsername, &diags)
	setResourceDataStringWithDiagnostic(d, "audit_level", input.AuditLevel, &diags)
	setResourceDataStringWithDiagnostic(d, "base_path", input.BasePath, &diags)
	setResourceDataStringWithDiagnostic(d, "host", input.Host, &diags)
	setResourceDataIntWithDiagnostic(d, "port", input.Port, &diags)
	setResourceDataBoolWithDiagnostic(d, "secure", input.Secure, &diags)
	setResourceDataIntWithDiagnostic(d, "trusted_certificate_group_id", input.TrustedCertificateGroupId, &diags)
	setResourceDataBoolWithDiagnostic(d, "use_proxy", input.UseProxy, &diags)

	if input.AdminPassword != nil {
		pw, ok := d.GetOk("admin_password.0.value")
		creds := flattenHiddenFieldView(input.AdminPassword)
		if ok {
			creds[0]["value"] = pw
		}
		if err := d.Set("admin_password", creds); err != nil {
			diags = append(diags, diag.FromErr(err))
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
	if v, ok := d.GetOk("admin_password"); ok {
		pfAdmin.AdminPassword = expandHiddenFieldView(v.([]interface{}))
	}
	if v, ok := d.GetOk("audit_level"); ok {
		pfAdmin.AuditLevel = String(v.(string))
	}
	if v, ok := d.GetOk("base_path"); ok {
		pfAdmin.BasePath = String(v.(string))
	}
	if v, ok := d.GetOk("secure"); ok {
		pfAdmin.Secure = Bool(v.(bool))
	}
	//if v, ok := d.Get("trusted_certificate_group_id"); ok {
		pfAdmin.TrustedCertificateGroupId = Int(d.Get("trusted_certificate_group_id").(int))
	//}
	if v, ok := d.GetOk("use_proxy"); ok {
		pfAdmin.UseProxy = Bool(v.(bool))
	}
	return pfAdmin
}
