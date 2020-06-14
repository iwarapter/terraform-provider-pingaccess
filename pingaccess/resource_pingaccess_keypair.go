package pingaccess

import (
	"context"
	"fmt"
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
			Type:     schema.TypeString,
			Required: true,
		},
		"password": {
			Type:      schema.TypeString,
			Required:  true,
			Sensitive: true,
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
			Type:     schema.TypeString,
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

func resourcePingAccessKeyPairCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
		return diag.FromErr(fmt.Errorf("unable to create KeyPair: %s", err))
	}

	d.SetId(strconv.Itoa(*result.Id))
	return resourcePingAccessKeyPairReadResult(d, result)
}

func resourcePingAccessKeyPairRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(*pa.Client).KeyPairs
	input := &pa.GetKeyPairCommandInput{
		Id: d.Id(),
	}
	result, _, err := svc.GetKeyPairCommand(input)
	if err != nil {
		return diag.FromErr(fmt.Errorf("unable to read KeyPair: %s", err))
	}
	return resourcePingAccessKeyPairReadResult(d, result)
}

func resourcePingAccessKeyPairUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
		return diag.FromErr(fmt.Errorf("unable to update KeyPair: %s", err))
	}

	d.SetId(strconv.Itoa(*result.Id))
	return resourcePingAccessKeyPairReadResult(d, result)
}

func resourcePingAccessKeyPairDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(*pa.Client).KeyPairs
	_, err := svc.DeleteKeyPairCommand(&pa.DeleteKeyPairCommandInput{Id: d.Id()})
	if err != nil {
		return diag.FromErr(fmt.Errorf("unable to delete KeyPair: %s", err))
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
	setResourceDataStringWithDiagnostic(d, "hsm_provider_id", String(strconv.Itoa(*rv.HsmProviderId)), &diags)
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
