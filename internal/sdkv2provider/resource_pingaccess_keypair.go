package sdkv2provider

import (
	"context"
	"crypto/ecdsa"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"regexp"
	"strconv"
	"time"

	models60 "github.com/iwarapter/pingaccess-sdk-go/v60/pingaccess/models"
	keyPairs60 "github.com/iwarapter/pingaccess-sdk-go/v60/services/keyPairs"
	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess/models"
	"github.com/iwarapter/pingaccess-sdk-go/v62/services/keyPairs"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePingAccessKeyPair() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePingAccessKeyPairCreate,
		ReadContext:   resourcePingAccessKeyPairRead,
		UpdateContext: resourcePingAccessKeyPairUpdate,
		DeleteContext: resourcePingAccessKeyPairDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourcePingAccessKeyPairImport,
		},
		Schema: resourcePingAccessKeyPairSchema(),
		Description: `Provides configuration for Keypairs within PingAccess.

-> The PingAccess API does not provider repeatable means of querying a sensitive value, we are unable to detect configuration drift of any sensitive fields in the configuration block.`,
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
			Type:        schema.TypeString,
			Required:    true,
			Description: "A unique alias name to identify the key pair. Special characters and spaces are allowed.",
		},
		"chain_certificates": {
			Type:        schema.TypeSet,
			Computed:    true,
			Description: "The complete list of certificates in the key pair certificate chain.",
			Elem: &schema.Resource{
				Schema: sch,
			},
		},
		"file_data": {
			Type:          schema.TypeString,
			Optional:      true,
			ConflictsWith: []string{"city", "common_name", "country", "key_algorithm", "key_size", "organization", "organization_unit", "state", "valid_days"},
			RequiredWith:  []string{"file_data", "password"},
			Description:   "Base-64 encoded PKCS12 or PEM file data. For PEM, the private key must precede the certificates, and certificates must be ordered from leaf to root. In BCFIPS mode, only PEM with PBES2 and AES or Triple DES encryption is accepted and 128-bit salt is required.",
		},
		"password": {
			Type:          schema.TypeString,
			Sensitive:     true,
			Optional:      true,
			ConflictsWith: []string{"city", "common_name", "country", "key_algorithm", "key_size", "organization", "organization_unit", "state", "valid_days"},
			RequiredWith:  []string{"file_data", "password"},
			Description:   "The password used to protect the private key. In FIPS mode, the password must be at least 14 characters.",
		},
		"city": {
			Type:          schema.TypeString,
			Optional:      true,
			ConflictsWith: []string{"file_data", "password"},
			RequiredWith:  []string{"city", "common_name", "country", "key_algorithm", "key_size", "organization", "organization_unit", "state", "valid_days"},
			Description:   "The city or other primary location (L) where the company operates.",
		},
		"common_name": {
			Type:          schema.TypeString,
			Optional:      true,
			ConflictsWith: []string{"file_data", "password"},
			RequiredWith:  []string{"city", "common_name", "country", "key_algorithm", "key_size", "organization", "organization_unit", "state", "valid_days"},
			Description:   "The common name (CN) identifying the certificate.",
		},
		"country": {
			Type:          schema.TypeString,
			Optional:      true,
			ConflictsWith: []string{"file_data", "password"},
			RequiredWith:  []string{"city", "common_name", "country", "key_algorithm", "key_size", "organization", "organization_unit", "state", "valid_days"},
			Description:   "The country (C) where the company is based, using two capital letters.",
		},
		"key_algorithm": {
			Type:          schema.TypeString,
			Optional:      true,
			ConflictsWith: []string{"file_data", "password"},
			RequiredWith:  []string{"city", "common_name", "country", "key_algorithm", "key_size", "organization", "organization_unit", "state", "valid_days"},
			Description:   "The key algorithm to use to generate a key.",
		},
		"key_size": {
			Type:          schema.TypeInt,
			Optional:      true,
			ConflictsWith: []string{"file_data", "password"},
			RequiredWith:  []string{"city", "common_name", "country", "key_algorithm", "key_size", "organization", "organization_unit", "state", "valid_days"},
			Description:   "The number of bits used in the key. Choices depend on selected key algorithm.",
		},
		"organization": {
			Type:          schema.TypeString,
			Optional:      true,
			ConflictsWith: []string{"file_data", "password"},
			RequiredWith:  []string{"city", "common_name", "country", "key_algorithm", "key_size", "organization", "organization_unit", "state", "valid_days"},
			Description:   "The organization (O) or company name creating the certificate.",
		},
		"organization_unit": {
			Type:          schema.TypeString,
			Optional:      true,
			ConflictsWith: []string{"file_data", "password"},
			RequiredWith:  []string{"city", "common_name", "country", "key_algorithm", "key_size", "organization", "organization_unit", "state", "valid_days"},
			Description:   "The specific unit within the organization (OU).",
		},
		"state": {
			Type:          schema.TypeString,
			Optional:      true,
			ConflictsWith: []string{"file_data", "password"},
			RequiredWith:  []string{"city", "common_name", "country", "key_algorithm", "key_size", "organization", "organization_unit", "state", "valid_days"},
			Description:   "The state (ST) or other political unit encompassing the location.",
		},
		//"subject_alternative_names": {},
		"valid_days": {
			Type:          schema.TypeInt,
			Optional:      true,
			ConflictsWith: []string{"file_data", "password"},
			RequiredWith:  []string{"city", "common_name", "country", "key_algorithm", "key_size", "organization", "organization_unit", "state", "valid_days"},
			Description:   "The number of days the certificate is valid.",
		},
		"csr_pending": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "True if a CSR is generated for this key pair.",
		},
		"expires": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: "The date at which the certificate expires as the number of milliseconds since January 1, 1970, 00:00:00 GMT.",
		},
		"hsm_provider_id": {
			Type:        schema.TypeInt,
			Optional:    true,
			Default:     "0",
			Description: "The HSM Provider ID. The default value is 0 indicating an HSM is not used for this key pair.",
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

func resourcePingAccessKeyPairCreate(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(paClient).KeyPairs
	//Import a keypair
	if _, ok := d.GetOk("file_data"); ok {
		//Import a keypair using 6.2+ API
		if m.(paClient).Is62OrAbove() {
			input := keyPairs.ImportKeyPairCommandInput{
				Body: models.PKCS12FileImportDocView{
					Alias:             String(d.Get("alias").(string)),
					FileData:          String(d.Get("file_data").(string)),
					Password:          &models.HiddenFieldView{Value: String(d.Get("password").(string))},
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
		//Import a keypair using 6.0 API
		input := keyPairs60.ImportKeyPairCommandInput{
			Body: models60.PKCS12FileImportDocView{
				Alias:             String(d.Get("alias").(string)),
				FileData:          String(d.Get("file_data").(string)),
				Password:          String(d.Get("password").(string)),
				ChainCertificates: &[]*string{},
				HsmProviderId:     Int(d.Get("hsm_provider_id").(int)),
			},
		}
		svc60 := m.(paClient).KeyPairsV60
		result, _, err := svc60.ImportKeyPairCommand(&input)
		if err != nil {
			return diag.Errorf("unable to create KeyPair: %s", err)
		}

		d.SetId(strconv.Itoa(*result.Id))
		return resourcePingAccessKeyPairReadResult(d, &models.KeyPairView{
			Alias:              result.Alias,
			CsrPending:         result.CsrPending,
			Expires:            result.Expires,
			HsmProviderId:      result.HsmProviderId,
			Id:                 result.Id,
			IssuerDn:           result.IssuerDn,
			Md5sum:             result.Md5sum,
			SerialNumber:       result.SerialNumber,
			Sha1sum:            result.Sha1sum,
			SignatureAlgorithm: result.SignatureAlgorithm,
			Status:             result.Status,
			SubjectCn:          result.SubjectCn,
			SubjectDn:          result.SubjectDn,
			ValidFrom:          result.ValidFrom,
		})
	}

	input := keyPairs.GenerateKeyPairCommandInput{
		Body: models.NewKeyPairConfigView{
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
	svc := m.(paClient).KeyPairs
	input := &keyPairs.GetKeyPairCommandInput{
		Id: d.Id(),
	}
	result, _, err := svc.GetKeyPairCommand(input)
	if err != nil {
		return diag.Errorf("unable to read KeyPair: %s", err)
	}
	return resourcePingAccessKeyPairReadResult(d, result)
}

func resourcePingAccessKeyPairUpdate(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	input := keyPairs.UpdateKeyPairCommandInput{
		Body: models.PKCS12FileImportDocView{
			Alias:             String(d.Get("alias").(string)),
			FileData:          String(d.Get("file_data").(string)),
			Password:          &models.HiddenFieldView{Value: String(d.Get("password").(string))},
			ChainCertificates: &[]*string{},
		},
		Id: d.Id(),
	}
	if v, ok := d.GetOk("hsm_provider_id"); ok {
		hsmID, _ := strconv.Atoi(v.(string))
		input.Body.HsmProviderId = Int(hsmID)
	}

	svc := m.(paClient).KeyPairs

	result, _, err := svc.UpdateKeyPairCommand(&input)
	if err != nil {
		return diag.Errorf("unable to update KeyPair: %s", err)
	}

	d.SetId(strconv.Itoa(*result.Id))
	return resourcePingAccessKeyPairReadResult(d, result)
}

func resourcePingAccessKeyPairDelete(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(paClient).KeyPairs
	_, err := svc.DeleteKeyPairCommand(&keyPairs.DeleteKeyPairCommandInput{Id: d.Id()})
	if err != nil {
		return diag.Errorf("unable to delete KeyPair: %s", err)
	}
	return nil
}

func resourcePingAccessKeyPairImport(_ context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	svc := m.(paClient).KeyPairs
	is6 := m.(paClient).Is60OrAbove()
	input := keyPairs.GetKeyPairCommandInput{
		Id: d.Id(),
	}
	result, _, err := svc.GetKeyPairCommand(&input)
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve read keypair for import %s", err)
	}
	certInput := keyPairs.ExportKeyPairCertInput{
		Id: d.Id(),
	}
	certResult, _, err := svc.ExportKeyPairCert(&certInput)
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve read keypair cert for import %s", err)
	}
	p, _ := pem.Decode([]byte(*certResult))
	cert, err := x509.ParseCertificate(p.Bytes)
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve parse keypair cert for import %s", err)
	}

	diags := resourcePingAccessKeyPairReadResult(d, result)
	if !is6 {
		//hsm providers are only available in 6+ so we just set to state to match the default for 5.3
		setResourceDataIntWithDiagnostic(d, "hsm_provider_id", Int(0), &diags)
	}

	//import based on upload
	//TODO unable to properly support upload style imports - https://discuss.hashicorp.com/t/importer-functions-reading-file-config/17624/2

	//import based on generate
	if m, err := extractMatch("CN=([^,]+)", *result.SubjectDn); err == nil {
		setResourceDataStringWithDiagnostic(d, "common_name", String(m), &diags)
	}
	if m, err := extractMatch("OU=([^,]+)", *result.SubjectDn); err == nil {
		setResourceDataStringWithDiagnostic(d, "organization_unit", String(m), &diags)
	}
	if m, err := extractMatch("O=([^,]+)", *result.SubjectDn); err == nil {
		setResourceDataStringWithDiagnostic(d, "organization", String(m), &diags)
	}
	if m, err := extractMatch("L=([^,]+)", *result.SubjectDn); err == nil {
		setResourceDataStringWithDiagnostic(d, "city", String(m), &diags)
	}
	if m, err := extractMatch("ST=([^,]+)", *result.SubjectDn); err == nil {
		setResourceDataStringWithDiagnostic(d, "state", String(m), &diags)
	}
	if m, err := extractMatch("C=([^,]+)", *result.SubjectDn); err == nil {
		setResourceDataStringWithDiagnostic(d, "country", String(m), &diags)
	}
	from := time.Unix(int64(*result.ValidFrom/1000), 0)  //time.Parse(time.RFC3339, *result.ValidFrom)
	expires := time.Unix(int64(*result.Expires/1000), 0) //, _ := time.Parse(time.RFC3339, *result.Expires)
	setResourceDataIntWithDiagnostic(d, "valid_days", Int(int(expires.Sub(from).Hours()/24)), &diags)
	setResourceDataStringWithDiagnostic(d, "key_algorithm", String(cert.PublicKeyAlgorithm.String()), &diags)
	switch pubKey := cert.PublicKey.(type) {
	case *rsa.PublicKey:
		setResourceDataIntWithDiagnostic(d, "key_size", Int(pubKey.N.BitLen()), &diags)
	case *ecdsa.PublicKey:
		setResourceDataIntWithDiagnostic(d, "key_size", Int(pubKey.Curve.Params().BitSize), &diags)
	default:
		return nil, fmt.Errorf("unable to parse certificate keysize, unsupported public key")
	}
	if diags.HasError() {
		return nil, fmt.Errorf("unable to import  %s", err)
	}
	return []*schema.ResourceData{d}, nil
}

func resourcePingAccessKeyPairReadResult(d *schema.ResourceData, rv *models.KeyPairView) diag.Diagnostics {
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

func extractMatch(re, source string) (string, error) {
	reg := regexp.MustCompile(re)
	matches := reg.FindStringSubmatch(source)
	if len(matches) == 2 {
		return matches[1], nil
	}
	return "", fmt.Errorf("unable to find match, matches: %v", matches)
}
