package pingaccess

import (
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	pa "github.com/iwarapter/pingaccess-sdk-go/pingaccess"
)

func resourcePingAccessKeyPair() *schema.Resource {
	return &schema.Resource{
		Create: resourcePingAccessKeyPairCreate,
		Read:   resourcePingAccessKeyPairRead,
		Update: resourcePingAccessKeyPairUpdate,
		Delete: resourcePingAccessKeyPairDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: resourcePingAccessKeyPairSchema(),
	}
}

func resourcePingAccessKeyPairSchema() map[string]*schema.Schema {
	//chain_certificates take a certificate schema with no file_data and a computed alias
	sch := resourcePingAccessCertificateSchema()
	delete(sch, "file_data")
	sch["alias"] = &schema.Schema{
		Type:     schema.TypeString,
		Computed: true,
	}

	return map[string]*schema.Schema{
		"alias": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"chain_certificates": &schema.Schema{
			Type:     schema.TypeSet,
			Computed: true,
			Elem: &schema.Resource{
				Schema: sch,
			},
		},
		"file_data": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"password": &schema.Schema{
			Type:      schema.TypeString,
			Required:  true,
			Sensitive: true,
		},
		"csr_pending": &schema.Schema{
			Type:     schema.TypeBool,
			Computed: true,
		},
		"expires": &schema.Schema{
			Type:     schema.TypeInt,
			Computed: true,
		},
		"hsm_provider_id": {
			Type:     schema.TypeString,
			Optional: true,
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

func resourcePingAccessKeyPairCreate(d *schema.ResourceData, m interface{}) error {
	input := pa.ImportKeyPairCommandInput{
		Body: pa.PKCS12FileImportDocView{
			Alias:             String(d.Get("alias").(string)),
			FileData:          String(d.Get("file_data").(string)),
			Password:          String(d.Get("password").(string)),
			ChainCertificates: &[]*string{},
		},
	}

	svc := m.(*pa.Client).KeyPairs

	result, _, err := svc.ImportKeyPairCommand(&input)
	if err != nil {
		return fmt.Errorf("Error creating KeyPair: %s", err)
	}

	d.SetId(result.Id.String())
	return resourcePingAccessKeyPairReadResult(d, result)
}

func resourcePingAccessKeyPairRead(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pa.Client).KeyPairs
	input := &pa.GetKeyPairCommandInput{
		Id: d.Id(),
	}
	result, _, err := svc.GetKeyPairCommand(input)
	if err != nil {
		return fmt.Errorf("Error reading KeyPair: %s", err)
	}
	return resourcePingAccessKeyPairReadResult(d, result)
}

func resourcePingAccessKeyPairUpdate(d *schema.ResourceData, m interface{}) error {
	input := pa.UpdateKeyPairCommandInput{
		Body: pa.PKCS12FileImportDocView{
			Alias:             String(d.Get("alias").(string)),
			FileData:          String(d.Get("file_data").(string)),
			Password:          String(d.Get("password").(string)),
			ChainCertificates: &[]*string{},
		},
		Id: d.Id(),
	}
	if _, ok := d.GetOkExists("hsm_provider_id"); ok {
		hsmId, _ := strconv.Atoi(d.Get("hsm_provider_id").(string))
		input.Body.HsmProviderId = Int(hsmId)
	}

	svc := m.(*pa.Client).KeyPairs

	result, _, err := svc.UpdateKeyPairCommand(&input)
	if err != nil {
		return fmt.Errorf("Error creating KeyPair: %s", err)
	}

	d.SetId(result.Id.String())
	return resourcePingAccessKeyPairReadResult(d, result)
}

func resourcePingAccessKeyPairDelete(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pa.Client).KeyPairs
	_, err := svc.DeleteKeyPairCommand(&pa.DeleteKeyPairCommandInput{Id: d.Id()})
	if err != nil {
		return fmt.Errorf("Error deleting key pair: %s", err)
	}
	return nil
}

func resourcePingAccessKeyPairReadResult(d *schema.ResourceData, rv *pa.KeyPairView) error {
	if err := d.Set("alias", rv.Alias); err != nil {
		return err
	}
	if rv.ChainCertificates != nil {
		if err := d.Set("chain_certificates", flattenChainCertificates(rv.ChainCertificates)); err != nil {
			return err
		}
	}
	if err := d.Set("csr_pending", rv.CsrPending); err != nil {
		return err
	}
	if err := d.Set("expires", rv.Expires); err != nil {
		return err
	}
	if err := d.Set("hsm_provider_id", rv.HsmProviderId); err != nil {

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
	if err := d.Set("subject_cn", rv.SubjectCn); err != nil {
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
