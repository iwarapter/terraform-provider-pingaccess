package pingaccess

import (
	"context"
	"strconv"

	"github.com/iwarapter/pingaccess-sdk-go/pingaccess/models"
	"github.com/iwarapter/pingaccess-sdk-go/services/certificates"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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

func resourcePingAccessCertificateCreate(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	input := certificates.ImportTrustedCertInput{
		Body: models.X509FileImportDocView{
			Alias:    String(d.Get("alias").(string)),
			FileData: String(d.Get("file_data").(string)),
		},
	}

	svc := m.(paClient).Certificates

	result, _, err := svc.ImportTrustedCert(&input)
	if err != nil {
		return diag.Errorf("unable to create Certificate: %s", err)
	}

	d.SetId(strconv.Itoa(*result.Id))
	return resourcePingAccessCertificateReadResult(d, result)
}

func resourcePingAccessCertificateRead(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(paClient).Certificates
	input := &certificates.GetTrustedCertInput{
		Id: d.Id(),
	}
	result, _, err := svc.GetTrustedCert(input)
	if err != nil {
		return diag.Errorf("unable to read Certificate: %s", err)
	}
	return resourcePingAccessCertificateReadResult(d, result)
}

func resourcePingAccessCertificateUpdate(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	input := certificates.UpdateTrustedCertInput{
		Body: models.X509FileImportDocView{
			Alias:    String(d.Get("alias").(string)),
			FileData: String(d.Get("file_data").(string)),
		},
		Id: d.Id(),
	}

	svc := m.(paClient).Certificates

	result, _, err := svc.UpdateTrustedCert(&input)
	if err != nil {
		return diag.Errorf("unable to update Certificate: %s", err)
	}

	d.SetId(strconv.Itoa(*result.Id))
	return resourcePingAccessCertificateReadResult(d, result)
}

func resourcePingAccessCertificateDelete(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(paClient).Certificates
	_, err := svc.DeleteTrustedCertCommand(&certificates.DeleteTrustedCertCommandInput{Id: d.Id()})
	if err != nil {
		return diag.Errorf("unable to delete Certificate: %s", err)
	}
	return nil
}

func resourcePingAccessCertificateReadResult(d *schema.ResourceData, rv *models.TrustedCertView) diag.Diagnostics {
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
