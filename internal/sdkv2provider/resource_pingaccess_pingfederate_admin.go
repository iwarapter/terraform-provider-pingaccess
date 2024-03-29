package sdkv2provider

import (
	"context"

	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess/models"
	"github.com/iwarapter/pingaccess-sdk-go/v62/services/pingfederate"

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
		Description: `Manages the PingAccess PingFederate Admin configuration.

-> This resource manages a singleton within PingAccess and as such you should ONLY ever declare one of this resource type. Deleting this resource resets the PingFederate Admin configuration to default values.`,
	}
}

func resourcePingAccessPingFederateAdminSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"admin_password": {
			Type:        schema.TypeList,
			Required:    true,
			MaxItems:    1,
			Description: "The password for the administrator username.",
			Elem:        hiddenFieldResource(),
		},
		"admin_username": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The administrator username.",
		},
		"audit_level": {
			Type:             schema.TypeString,
			Optional:         true,
			ValidateDiagFunc: validateAuditLevel,
			Default:          "ON",
			Description:      "Enable to record requests to the PingFederate Administrative API to the audit store.",
		},
		"base_path": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The base path, if needed, for Administration API.",
		},
		"host": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The host name or IP address for PingFederate Administration API.",
		},
		"port": {
			Type:        schema.TypeInt,
			Required:    true,
			Description: "The port number for PingFederate Administration API.",
		},
		"secure": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Enable if PingFederate is expecting HTTPS connections.",
		},
		"trusted_certificate_group_id": {
			Type:        schema.TypeInt,
			Optional:    true,
			Default:     0,
			Description: "The group of certificates to use when authenticating to PingFederate Administrative API.",
		},
		"use_proxy": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
			Description: "True if a proxy should be used for HTTP or HTTPS requests.",
		},
	}
}

func resourcePingAccessPingFederateAdminCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	d.SetId("pingfederate_admin_settings")
	return resourcePingAccessPingFederateAdminUpdate(ctx, d, m)
}

func resourcePingAccessPingFederateAdminRead(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(paClient).Pingfederate
	result, _, err := svc.GetPingFederateAdminCommand()
	if err != nil {
		return diag.Errorf("unable to read PingFederateAdmin: %s", err)
	}

	return resourcePingAccessPingFederateAdminReadResult(d, result)
}

func resourcePingAccessPingFederateAdminUpdate(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(paClient).Pingfederate
	input := pingfederate.UpdatePingFederateAdminCommandInput{
		Body: *resourcePingAccessPingFederateAdminReadData(d),
	}
	result, _, err := svc.UpdatePingFederateAdminCommand(&input)
	if err != nil {
		return diag.Errorf("unable to update PingFederateAdmin: %s", err)
	}

	d.SetId("pingfederate_admin_settings")
	return resourcePingAccessPingFederateAdminReadResult(d, result)
}

func resourcePingAccessPingFederateAdminDelete(_ context.Context, _ *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(paClient).Pingfederate
	_, err := svc.DeletePingFederateCommand()
	if err != nil {
		return diag.Errorf("unable to reset PingFederateAdmin: %s", err)
	}
	return nil
}

func resourcePingAccessPingFederateAdminReadResult(d *schema.ResourceData, input *models.PingFederateAdminView) diag.Diagnostics {
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
			diags = append(diags, diag.FromErr(err)...)
		}
	}

	return nil
}

func resourcePingAccessPingFederateAdminReadData(d *schema.ResourceData) *models.PingFederateAdminView {
	pfAdmin := &models.PingFederateAdminView{
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
