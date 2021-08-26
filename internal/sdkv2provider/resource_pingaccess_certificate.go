package sdkv2provider

import (
	"context"
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"

	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess/models"
	"github.com/iwarapter/pingaccess-sdk-go/v62/services/certificates"

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
			StateContext: resourcePingAccessCertificateImport,
		},
		Schema:      resourcePingAccessCertificateSchema(),
		Description: `Provides configuration for Certificates within PingAccess.`,
	}
}

func resourcePingAccessCertificateSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"alias": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The alias for the certificate.",
		},
		"file_data": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The base64-encoded data of the certificate.",
		},
		"expires": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: "The date at which the certificate expires as the number of milliseconds since January 1, 1970, 00:00:00 GMT.",
		},
		"issuer_dn": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The issuer DN for the certificate.",
		},
		"md5sum": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: `The MD5 sum for the certificate. The value will be set to "" when in FIPS mode.`,
		},
		"serial_number": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The serial number for the certificate.",
		},
		"sha1sum": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The SHA1 sum for the certificate.",
		},
		"signature_algorithm": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The algorithm used to sign the certificate.",
		},
		"status": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "A high-level status for the certificate.",
		},
		"subject_cn": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The subject CN for the certificate.",
		},
		"subject_dn": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The subject DN for the certificate.",
		},
		"valid_from": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: "The date at which the certificate is valid from as the number of milliseconds since January 1, 1970, 00:00:00 GMT.",
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

func resourcePingAccessCertificateImport(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	diags := resourcePingAccessCertificateRead(ctx, d, m)
	if diags.HasError() {
		msg := []string{}
		for _, diagnostic := range diags {
			msg = append(msg, diagnostic.Summary)
		}
		return nil, fmt.Errorf("unable to retrieve certifcate information:\n%s", strings.Join(msg, "\n"))
	}

	svc := m.(paClient).Certificates
	input := certificates.ExportTrustedCertInput{
		Id: d.Id(),
	}
	result, _, err := svc.ExportTrustedCert(&input)
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve certifcate file data %s", err)
	}
	trimmed := strings.Replace(*result, "\r\n", "\n", -1)
	encoded := base64.StdEncoding.EncodeToString([]byte(trimmed))
	setResourceDataStringWithDiagnostic(d, "file_data", String(encoded), &diags)

	return []*schema.ResourceData{d}, nil
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
