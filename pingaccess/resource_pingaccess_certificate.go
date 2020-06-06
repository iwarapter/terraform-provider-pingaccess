package pingaccess

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pa "github.com/iwarapter/pingaccess-sdk-go/pingaccess"
)

func resourcePingAccessCertificate() *schema.Resource {
	return &schema.Resource{
		Create: resourcePingAccessCertificateCreate,
		Read:   resourcePingAccessCertificateRead,
		Update: resourcePingAccessCertificateUpdate,
		Delete: resourcePingAccessCertificateDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: resourcePingAccessCertificateSchema(),
	}
}

func resourcePingAccessCertificateSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"alias": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"file_data": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"expires": &schema.Schema{
			Type:     schema.TypeInt,
			Computed: true,
		},
		"issuer_dn": &schema.Schema{
			Type:     schema.TypeString,
			Computed: true,
		},
		"md5sum": &schema.Schema{
			Type:     schema.TypeString,
			Computed: true,
		},
		"serial_number": &schema.Schema{
			Type:     schema.TypeString,
			Computed: true,
		},
		"sha1sum": &schema.Schema{
			Type:     schema.TypeString,
			Computed: true,
		},
		"signature_algorithm": &schema.Schema{
			Type:     schema.TypeString,
			Computed: true,
		},
		"status": &schema.Schema{
			Type:     schema.TypeString,
			Computed: true,
		},
		// "subject_alternative_names": setOfString(),
		"subject_cn": &schema.Schema{
			Type:     schema.TypeString,
			Computed: true,
		},
		"subject_dn": &schema.Schema{
			Type:     schema.TypeString,
			Computed: true,
		},
		"valid_from": &schema.Schema{
			Type:     schema.TypeInt,
			Computed: true,
		},
	}
}

func resourcePingAccessCertificateCreate(d *schema.ResourceData, m interface{}) error {
	input := pa.ImportTrustedCertInput{
		Body: pa.X509FileImportDocView{
			Alias:    String(d.Get("alias").(string)),
			FileData: String(d.Get("file_data").(string)),
		},
	}

	svc := m.(*pa.Client).Certificates

	result, _, err := svc.ImportTrustedCert(&input)
	if err != nil {
		return fmt.Errorf("Error creating Certificate: %s", err)
	}

	d.SetId(result.Id.String())
	return resourcePingAccessCertificateReadResult(d, result)
}

func resourcePingAccessCertificateRead(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pa.Client).Certificates
	input := &pa.GetTrustedCertInput{
		Id: d.Id(),
	}
	result, _, err := svc.GetTrustedCert(input)
	if err != nil {
		return fmt.Errorf("Error reading Certificate: %s", err)
	}
	return resourcePingAccessCertificateReadResult(d, result)
}

func resourcePingAccessCertificateUpdate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error creating Certificate: %s", err)
	}

	d.SetId(result.Id.String())
	return resourcePingAccessCertificateReadResult(d, result)
}

func resourcePingAccessCertificateDelete(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pa.Client).Certificates
	_, err := svc.DeleteTrustedCertCommand(&pa.DeleteTrustedCertCommandInput{Id: d.Id()})
	if err != nil {
		return fmt.Errorf("Error deleting virtualhost: %s", err)
	}
	return nil
}

func resourcePingAccessCertificateReadResult(d *schema.ResourceData, rv *pa.TrustedCertView) error {
	if err := d.Set("alias", rv.Alias); err != nil {
		return err
	}
	if err := d.Set("expires", rv.Expires); err != nil {
		return err
	}
	if err := d.Set("issuer_dn", rv.IssuerDn); err != nil {
		return err
	}
	if err := d.Set("md5sum", rv.Md5sum); err != nil {
		return err
	}
	if err := d.Set("serial_number", rv.SerialNumber); err != nil {
		return err
	}
	if err := d.Set("sha1sum", rv.Sha1sum); err != nil {
		return err
	}
	if err := d.Set("signature_algorithm", rv.SignatureAlgorithm); err != nil {
		return err
	}
	if err := d.Set("status", rv.Status); err != nil {
		return err
	}
	if err := d.Set("subject_dn", rv.SubjectDn); err != nil {
		return err
	}
	if err := d.Set("valid_from", rv.ValidFrom); err != nil {
		return err
	}
	// "subject_alternative_names": setOfString(),
	// "subject_cn": &schema.Schema{
	// 	Type:     schema.TypeString,
	// 	Computed: true,
	// },
	return nil
}
