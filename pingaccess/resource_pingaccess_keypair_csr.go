package pingaccess

import (
	"context"
	"strconv"

	"github.com/iwarapter/pingaccess-sdk-go/pingaccess/models"
	"github.com/iwarapter/pingaccess-sdk-go/services/keyPairs"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePingAccessKeyPairCsr() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePingAccessKeyPairCsrCreate,
		ReadContext:   resourcePingAccessKeyPairCsrRead,
		DeleteContext: resourcePingAccessKeyPairCsrDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: resourcePingAccessKeyPairCsrSchema(),
	}
}

func resourcePingAccessKeyPairCsrSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"keypair_id": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
		"file_data": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
	}
}

func resourcePingAccessKeyPairCsrCreate(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(paClient).KeyPairs
	input := keyPairs.ImportCSRResponseCommandInput{
		Body: models.CSRResponseImportDocView{
			FileData: String(d.Get("file_data").(string)),
		},
		Id: d.Get("keypair_id").(string),
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

func resourcePingAccessKeyPairCsrDelete(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	d.SetId("")
	return nil
}
