package sdkv2provider

import (
	"context"
	"strconv"

	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess/models"
	"github.com/iwarapter/pingaccess-sdk-go/v62/services/trustedCertificateGroups"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePingAccessTrustedCertificateGroups() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePingAccessTrustedCertificateGroupsCreate,
		ReadContext:   resourcePingAccessTrustedCertificateGroupsRead,
		UpdateContext: resourcePingAccessTrustedCertificateGroupsUpdate,
		DeleteContext: resourcePingAccessTrustedCertificateGroupsDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: resourcePingAccessTrustedCertificateGroupsSchema(),
	}
}

func resourcePingAccessTrustedCertificateGroupsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"cert_ids": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"ignore_all_certificate_errors": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"skip_certificate_date_check": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"system_group": {
			Type:     schema.TypeBool,
			Computed: true,
		},
		"use_java_trust_store": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
	}
}

func resourcePingAccessTrustedCertificateGroupsCreate(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(paClient).TrustedCertificateGroups
	input := trustedCertificateGroups.AddTrustedCertificateGroupCommandInput{
		Body: *resourcePingAccessTrustedCertificateGroupsReadData(d),
	}

	result, _, err := svc.AddTrustedCertificateGroupCommand(&input)
	if err != nil {
		return diag.Errorf("unable to create TrustedCertificateGroup: %s", err)
	}

	d.SetId(result.Id.String())
	return resourcePingAccessTrustedCertificateGroupsReadResult(d, result)
}

func resourcePingAccessTrustedCertificateGroupsRead(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(paClient).TrustedCertificateGroups
	input := &trustedCertificateGroups.GetTrustedCertificateGroupCommandInput{
		Id: d.Id(),
	}
	result, _, err := svc.GetTrustedCertificateGroupCommand(input)
	if err != nil {
		return diag.Errorf("unable to read TrustedCertificateGroup: %s", err)
	}
	return resourcePingAccessTrustedCertificateGroupsReadResult(d, result)
}

func resourcePingAccessTrustedCertificateGroupsUpdate(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(paClient).TrustedCertificateGroups
	input := trustedCertificateGroups.UpdateTrustedCertificateGroupCommandInput{
		Body: *resourcePingAccessTrustedCertificateGroupsReadData(d),
		Id:   d.Id(),
	}
	result, _, err := svc.UpdateTrustedCertificateGroupCommand(&input)
	if err != nil {
		return diag.Errorf("unable to update TrustedCertificateGroup: %s", err)
	}
	return resourcePingAccessTrustedCertificateGroupsReadResult(d, result)
}

func resourcePingAccessTrustedCertificateGroupsDelete(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(paClient).TrustedCertificateGroups
	input := &trustedCertificateGroups.DeleteTrustedCertificateGroupCommandInput{
		Id: d.Id(),
	}

	_, err := svc.DeleteTrustedCertificateGroupCommand(input)
	if err != nil {
		return diag.Errorf("unable to delete TrustedCertificateGroup: %s", err)
	}
	return nil
}

func resourcePingAccessTrustedCertificateGroupsReadResult(d *schema.ResourceData, input *models.TrustedCertificateGroupView) diag.Diagnostics {
	var diags diag.Diagnostics
	setResourceDataBoolWithDiagnostic(d, "ignore_all_certificate_errors", input.IgnoreAllCertificateErrors, &diags)
	setResourceDataStringWithDiagnostic(d, "name", input.Name, &diags)
	setResourceDataBoolWithDiagnostic(d, "skip_certificate_date_check", input.SkipCertificateDateCheck, &diags)
	setResourceDataBoolWithDiagnostic(d, "system_group", input.SystemGroup, &diags)
	setResourceDataBoolWithDiagnostic(d, "use_java_trust_store", input.UseJavaTrustStore, &diags)

	if input.CertIds != nil {
		var certs []string
		for _, cert := range *input.CertIds {
			certs = append(certs, strconv.Itoa(*cert))
		}
		if err := d.Set("cert_ids", certs); err != nil {
			diags = append(diags, diag.FromErr(err)...)

		}
	}
	return diags
}

func resourcePingAccessTrustedCertificateGroupsReadData(d *schema.ResourceData) *models.TrustedCertificateGroupView {

	tcg := &models.TrustedCertificateGroupView{
		Name: String(d.Get("name").(string)),
	}

	if v, ok := d.GetOk("cert_ids"); ok {
		certs := expandStringList(v.([]interface{}))
		var certIDs []*int
		for _, i := range certs {
			text, _ := strconv.Atoi(*i)
			certIDs = append(certIDs, &text)
		}
		tcg.CertIds = &certIDs
	}
	tcg.IgnoreAllCertificateErrors = Bool(d.Get("ignore_all_certificate_errors").(bool))
	tcg.SkipCertificateDateCheck = Bool(d.Get("skip_certificate_date_check").(bool))
	tcg.SystemGroup = Bool(d.Get("system_group").(bool))
	tcg.UseJavaTrustStore = Bool(d.Get("use_java_trust_store").(bool))

	return tcg
}
