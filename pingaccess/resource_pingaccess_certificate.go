package pingaccess

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pa "github.com/iwarapter/pingaccess-sdk-go/pingaccess"
)

func resourcePingAccessCertificate() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePingAccessCertificateCreate,
		ReadContext:   resourcePingAccessCertificateRead,
		UpdateContext: resourcePingAccessCertificateUpdate,
		DeleteContext: resourcePingAccessCertificateDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: resourcePingAccessCertificateSchema(),
	}
}

func resourcePingAccessCertificateSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"alias": {
			Type:     schema.TypeString,
			Required: true,
		},
		"file_data": {
			Type:     schema.TypeString,
			Required: true,
		},
		"expires": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"issuer_dn": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"md5sum": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"serial_number": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"sha1sum": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"signature_algorithm": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"status": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"subject_cn": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"subject_dn": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"valid_from": {
			Type:     schema.TypeInt,
			Computed: true,
		},
	}
}

func resourcePingAccessCertificateCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	input := pa.ImportTrustedCertInput{
		Body: pa.X509FileImportDocView{
			Alias:    String(d.Get("alias").(string)),
			FileData: String(d.Get("file_data").(string)),
		},
	}

	svc := m.(*pa.Client).Certificates

	result, _, err := svc.ImportTrustedCert(&input)
	if err != nil {
		return diag.Diagnostics{diag.FromErr(fmt.Errorf("unable to create Certificate: %s", err))}
	}

	d.SetId(result.Id.String())
	return resourcePingAccessCertificateReadResult(d, result)
}

func resourcePingAccessCertificateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(*pa.Client).Certificates
	input := &pa.GetTrustedCertInput{
		Id: d.Id(),
	}
	result, _, err := svc.GetTrustedCert(input)
	if err != nil {
		return diag.Diagnostics{diag.FromErr(fmt.Errorf("unable to read Certificate: %s", err))}
	}
	return resourcePingAccessCertificateReadResult(d, result)
}

func resourcePingAccessCertificateUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	input := pa.UpdateTrustedCertInput{
		Body: pa.X509FileImportDocView{
			Alias:    String(d.Get("alias").(string)),
			FileData: String(d.Get("file_data").(string)),
		},
		Id: d.Id(),
	}

	svc := m.(*pa.Client).Certificates

	result, _, err := svc.UpdateTrustedCert(&input)
	if err != nil {
		return diag.Diagnostics{diag.FromErr(fmt.Errorf("unable to update Certificate: %s", err))}
	}

	d.SetId(result.Id.String())
	return resourcePingAccessCertificateReadResult(d, result)
}

func resourcePingAccessCertificateDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(*pa.Client).Certificates
	_, err := svc.DeleteTrustedCertCommand(&pa.DeleteTrustedCertCommandInput{Id: d.Id()})
	if err != nil {
		return diag.Diagnostics{diag.FromErr(fmt.Errorf("unable to delete Certificate: %s", err))}
	}
	return nil
}

func resourcePingAccessCertificateReadResult(d *schema.ResourceData, rv *pa.TrustedCertView) diag.Diagnostics {
	var diags diag.Diagnostics
	setResourceDataStringWithDiagnostic(d, "alias", rv.Alias, &diags)
	setResourceDataIntWithDiagnostic(d, "expires", rv.Expires, &diags)
	setResourceDataStringWithDiagnostic(d, "issuer_dn", rv.IssuerDn, &diags)
	setResourceDataStringWithDiagnostic(d, "md5sum", rv.Md5sum, &diags)
	setResourceDataStringWithDiagnostic(d, "serial_number", rv.SerialNumber, &diags)
	setResourceDataStringWithDiagnostic(d, "sha1sum", rv.Sha1sum, &diags)
	setResourceDataStringWithDiagnostic(d, "signature_algorithm", rv.SignatureAlgorithm, &diags)
	setResourceDataStringWithDiagnostic(d, "status", rv.Status, &diags)
	setResourceDataStringWithDiagnostic(d, "subject_dn", rv.SubjectDn, &diags)
	setResourceDataIntWithDiagnostic(d, "valid_from", rv.ValidFrom, &diags)
	return diags
}
