package pingaccess

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pa "github.com/iwarapter/pingaccess-sdk-go/pingaccess"
)

func resourcePingAccessKeyPair() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePingAccessKeyPairCreate,
		ReadContext:   resourcePingAccessKeyPairRead,
		UpdateContext: resourcePingAccessKeyPairUpdate,
		DeleteContext: resourcePingAccessKeyPairDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
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
		"alias": {
			Type:     schema.TypeString,
			Required: true,
		},
		"chain_certificates": {
			Type:     schema.TypeSet,
			Computed: true,
			Elem: &schema.Resource{
				Schema: sch,
			},
		},
		"file_data": {
			Type:          schema.TypeString,
			Optional:      true,
			ConflictsWith: []string{"city", "common_name", "country", "key_algorithm", "key_size", "organization", "organization_unit", "state", "valid_days"},
			RequiredWith:  []string{"file_data", "password"},
		},
		"password": {
			Type:          schema.TypeString,
			Sensitive:     true,
			Optional:      true,
			ConflictsWith: []string{"city", "common_name", "country", "key_algorithm", "key_size", "organization", "organization_unit", "state", "valid_days"},
			RequiredWith:  []string{"file_data", "password"},
		},
		"city": {
			Type:          schema.TypeString,
			Optional:      true,
			ConflictsWith: []string{"file_data", "password"},
			RequiredWith:  []string{"city", "common_name", "country", "key_algorithm", "key_size", "organization", "organization_unit", "state", "valid_days"},
		},
		"common_name": {
			Type:          schema.TypeString,
			Optional:      true,
			ConflictsWith: []string{"file_data", "password"},
			RequiredWith:  []string{"city", "common_name", "country", "key_algorithm", "key_size", "organization", "organization_unit", "state", "valid_days"},
		},
		"country": {
			Type:          schema.TypeString,
			Optional:      true,
			ConflictsWith: []string{"file_data", "password"},
			RequiredWith:  []string{"city", "common_name", "country", "key_algorithm", "key_size", "organization", "organization_unit", "state", "valid_days"},
		},
		"key_algorithm": {
			Type:          schema.TypeString,
			Optional:      true,
			ConflictsWith: []string{"file_data", "password"},
			RequiredWith:  []string{"city", "common_name", "country", "key_algorithm", "key_size", "organization", "organization_unit", "state", "valid_days"},
		},
		"key_size": {
			Type:          schema.TypeInt,
			Optional:      true,
			ConflictsWith: []string{"file_data", "password"},
			RequiredWith:  []string{"city", "common_name", "country", "key_algorithm", "key_size", "organization", "organization_unit", "state", "valid_days"},
		},
		"organization": {
			Type:          schema.TypeString,
			Optional:      true,
			ConflictsWith: []string{"file_data", "password"},
			RequiredWith:  []string{"city", "common_name", "country", "key_algorithm", "key_size", "organization", "organization_unit", "state", "valid_days"},
		},
		"organization_unit": {
			Type:          schema.TypeString,
			Optional:      true,
			ConflictsWith: []string{"file_data", "password"},
			RequiredWith:  []string{"city", "common_name", "country", "key_algorithm", "key_size", "organization", "organization_unit", "state", "valid_days"},
		},
		"state": {
			Type:          schema.TypeString,
			Optional:      true,
			ConflictsWith: []string{"file_data", "password"},
			RequiredWith:  []string{"city", "common_name", "country", "key_algorithm", "key_size", "organization", "organization_unit", "state", "valid_days"},
		},
		//"subject_alternative_names": {},
		"valid_days": {
			Type:          schema.TypeInt,
			Optional:      true,
			ConflictsWith: []string{"file_data", "password"},
			RequiredWith:  []string{"city", "common_name", "country", "key_algorithm", "key_size", "organization", "organization_unit", "state", "valid_days"},
		},
		"csr_pending": {
			Type:     schema.TypeBool,
			Computed: true,
		},
		"expires": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"hsm_provider_id": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  "0",
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

func resourcePingAccessKeyPairCreate(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(*pa.Client).KeyPairs
	if _, ok := d.GetOk("file_data"); ok {
		input := pa.ImportKeyPairCommandInput{
			Body: pa.PKCS12FileImportDocView{
				Alias:             String(d.Get("alias").(string)),
				FileData:          String(d.Get("file_data").(string)),
				Password:          String(d.Get("password").(string)),
				ChainCertificates: &[]*string{},
				HsmProviderId:     Int(d.Get("hsm_provider_id").(int)),
			},
		}
		result, _, err := svc.ImportKeyPairCommand(&input)
		if err != nil {
			return diag.Errorf("unable to create KeyPair: %s", err)
		}

		d.SetId(strconv.Itoa(*result.Id))
		return resourcePingAccessKeyPairReadResult(d, result)
	}

	input := pa.GenerateKeyPairCommandInput{
		Body: pa.NewKeyPairConfigView{
			Alias:            String(d.Get("alias").(string)),
			City:             String(d.Get("city").(string)),
			CommonName:       String(d.Get("common_name").(string)),
			Country:          String(d.Get("country").(string)),
			KeyAlgorithm:     String(d.Get("key_algorithm").(string)),
			KeySize:          Int(d.Get("key_size").(int)),
			Organization:     String(d.Get("organization").(string)),
			OrganizationUnit: String(d.Get("organization_unit").(string)),
			State:            String(d.Get("state").(string)),
			ValidDays:        Int(d.Get("valid_days").(int)),
			HsmProviderId:    Int(d.Get("hsm_provider_id").(int)),
		},
	}

	result, _, err := svc.GenerateKeyPairCommand(&input)
	if err != nil {
		return diag.Errorf("unable to generate KeyPair: %s", err)
	}

	d.SetId(strconv.Itoa(*result.Id))
	return resourcePingAccessKeyPairReadResult(d, result)
}

func resourcePingAccessKeyPairRead(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(*pa.Client).KeyPairs
	input := &pa.GetKeyPairCommandInput{
		Id: d.Id(),
	}
	result, _, err := svc.GetKeyPairCommand(input)
	if err != nil {
		return diag.Errorf("unable to read KeyPair: %s", err)
	}
	return resourcePingAccessKeyPairReadResult(d, result)
}

func resourcePingAccessKeyPairUpdate(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	input := pa.UpdateKeyPairCommandInput{
		Body: pa.PKCS12FileImportDocView{
			Alias:             String(d.Get("alias").(string)),
			FileData:          String(d.Get("file_data").(string)),
			Password:          String(d.Get("password").(string)),
			ChainCertificates: &[]*string{},
		},
		Id: d.Id(),
	}
	if v, ok := d.GetOk("hsm_provider_id"); ok {
		hsmID, _ := strconv.Atoi(v.(string))
		input.Body.HsmProviderId = Int(hsmID)
	}

	svc := m.(*pa.Client).KeyPairs

	result, _, err := svc.UpdateKeyPairCommand(&input)
	if err != nil {
		return diag.Errorf("unable to update KeyPair: %s", err)
	}

	d.SetId(strconv.Itoa(*result.Id))
	return resourcePingAccessKeyPairReadResult(d, result)
}

func resourcePingAccessKeyPairDelete(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(*pa.Client).KeyPairs
	_, err := svc.DeleteKeyPairCommand(&pa.DeleteKeyPairCommandInput{Id: d.Id()})
	if err != nil {
		return diag.Errorf("unable to delete KeyPair: %s", err)
	}
	return nil
}

func resourcePingAccessKeyPairReadResult(d *schema.ResourceData, rv *pa.KeyPairView) diag.Diagnostics {
	var diags diag.Diagnostics
	setResourceDataStringWithDiagnostic(d, "alias", rv.Alias, &diags)
	if rv.ChainCertificates != nil {
		if err := d.Set("chain_certificates", flattenChainCertificates(rv.ChainCertificates)); err != nil {
			diags = append(diags, diag.FromErr(err)...)

		}
	}
	setResourceDataBoolWithDiagnostic(d, "csr_pending", rv.CsrPending, &diags)
	setResourceDataIntWithDiagnostic(d, "expires", rv.Expires, &diags)
	setResourceDataIntWithDiagnostic(d, "hsm_provider_id", rv.HsmProviderId, &diags)
	setResourceDataStringWithDiagnostic(d, "issuer_dn", rv.IssuerDn, &diags)
	setResourceDataStringWithDiagnostic(d, "md5sum", rv.Md5sum, &diags)
	setResourceDataStringWithDiagnostic(d, "serial_number", rv.SerialNumber, &diags)
	setResourceDataStringWithDiagnostic(d, "sha1sum", rv.Sha1sum, &diags)
	setResourceDataStringWithDiagnostic(d, "signature_algorithm", rv.SignatureAlgorithm, &diags)
	setResourceDataStringWithDiagnostic(d, "status", rv.Status, &diags)
	setResourceDataStringWithDiagnostic(d, "subject_dn", rv.SubjectDn, &diags)
	setResourceDataStringWithDiagnostic(d, "subject_cn", rv.SubjectCn, &diags)
	setResourceDataIntWithDiagnostic(d, "valid_from", rv.ValidFrom, &diags)
	return diags
}
