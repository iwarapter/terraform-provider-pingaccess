package sdkv2provider

import (
	"context"
	"strconv"

	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess/models"
	"github.com/iwarapter/pingaccess-sdk-go/v62/services/keyPairs"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePingAccessKeyPairCsr() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePingAccessKeyPairCsrCreate,
		ReadContext:   resourcePingAccessKeyPairCsrRead,
		UpdateContext: resourcePingAccessKeyPairCsrUpdate,
		DeleteContext: resourcePingAccessKeyPairCsrDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema:      resourcePingAccessKeyPairCsrSchema(),
		Description: `Provides configuration for Keypair CSRs within PingAccess.`,
	}
}

func resourcePingAccessKeyPairCsrSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"keypair_id": {
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
			Description: "ID of the Key Pair to update.",
		},
		"file_data": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The CSR response data.",
		},
		"trusted_certificate_group_id": {
			Type:        schema.TypeInt,
			Optional:    true,
			Default:     0,
			Description: "The ID of the trusted certificate group associated with the CSR response.",
		},
		"chain_certificates": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "A list of base64-encoded certificates to add to the key pair certificate chain.",
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
	}
}

func resourcePingAccessKeyPairCsrCreate(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(paClient).KeyPairs
	input := keyPairs.ImportCSRResponseCommandInput{
		Body: *resourcePingAccessKeyPairCsrReadData(d),
		Id:   d.Get("keypair_id").(string),
	}

	result, _, err := svc.ImportCSRResponseCommand(&input)
	if err != nil {
		return diag.Errorf("unable to create KeyPairCsr: %s", err)
	}

	d.SetId(strconv.Itoa(*result.Id))
	return nil
}

func resourcePingAccessKeyPairCsrRead(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
}

func resourcePingAccessKeyPairCsrUpdate(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(paClient).KeyPairs
	input := keyPairs.ImportCSRResponseCommandInput{
		Body: *resourcePingAccessKeyPairCsrReadData(d),
		Id:   d.Get("keypair_id").(string),
	}

	_, _, err := svc.ImportCSRResponseCommand(&input)
	if err != nil {
		return diag.Errorf("unable to update KeyPairCsr: %s", err)
	}
	return nil
}

func resourcePingAccessKeyPairCsrDelete(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	d.SetId("")
	return nil
}

func resourcePingAccessKeyPairCsrReadData(d *schema.ResourceData) *models.CSRResponseImportDocView {
	csr := models.CSRResponseImportDocView{
		FileData: String(d.Get("file_data").(string)),
	}

	if v, ok := d.GetOk("chain_certificates"); ok {
		certs := expandStringList(v.([]interface{}))
		csr.ChainCertificates = &certs
	}
	if v, ok := d.GetOk("trusted_certificate_group_id"); ok {
		csr.TrustedCertGroupId = Int(v.(int))
	}

	return &csr
}
