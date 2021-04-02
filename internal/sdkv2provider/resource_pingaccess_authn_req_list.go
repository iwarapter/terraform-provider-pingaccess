package sdkv2provider

import (
	"context"

	"github.com/iwarapter/pingaccess-sdk-go/pingaccess/models"
	"github.com/iwarapter/pingaccess-sdk-go/services/authnReqLists"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePingAccessAuthnReqList() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePingAccessAuthnReqListCreate,
		ReadContext:   resourcePingAccessAuthnReqListRead,
		UpdateContext: resourcePingAccessAuthnReqListUpdate,
		DeleteContext: resourcePingAccessAuthnReqListDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: resourcePingAccessAuthnReqListSchema(),
	}
}

func resourcePingAccessAuthnReqListSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"authn_reqs": requiredListOfString(),
	}
}

func resourcePingAccessAuthnReqListCreate(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(paClient).AuthnReqLists
	input := authnReqLists.AddAuthnReqListCommandInput{
		Body: *resourcePingAccessAuthnReqListReadData(d),
	}

	result, _, err := svc.AddAuthnReqListCommand(&input)
	if err != nil {
		return diag.Errorf("unable to create AuthnReqList: %s", err)
	}

	d.SetId(result.Id.String())
	return resourcePingAccessAuthnReqListReadResult(d, &input.Body)
}

func resourcePingAccessAuthnReqListRead(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(paClient).AuthnReqLists
	input := &authnReqLists.GetAuthnReqListCommandInput{
		Id: d.Id(),
	}
	result, _, err := svc.GetAuthnReqListCommand(input)
	if err != nil {
		return diag.Errorf("unable to read AuthnReqList: %s", err)
	}
	return resourcePingAccessAuthnReqListReadResult(d, result)
}

func resourcePingAccessAuthnReqListUpdate(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(paClient).AuthnReqLists
	input := authnReqLists.UpdateAuthnReqListCommandInput{
		Body: *resourcePingAccessAuthnReqListReadData(d),
		Id:   d.Id(),
	}

	result, _, err := svc.UpdateAuthnReqListCommand(&input)
	if err != nil {
		return diag.Errorf("unable to update AuthnReqList: %s", err)
	}
	return resourcePingAccessAuthnReqListReadResult(d, result)
}

func resourcePingAccessAuthnReqListDelete(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(paClient).AuthnReqLists
	input := &authnReqLists.DeleteAuthnReqListCommandInput{
		Id: d.Id(),
	}

	_, err := svc.DeleteAuthnReqListCommand(input)
	if err != nil {
		return diag.Errorf("unable to delete AuthnReqList: %s", err)
	}
	return nil
}

func resourcePingAccessAuthnReqListReadResult(d *schema.ResourceData, input *models.AuthnReqListView) diag.Diagnostics {
	var diags diag.Diagnostics
	setResourceDataStringWithDiagnostic(d, "name", input.Name, &diags)
	if err := d.Set("authn_reqs", input.AuthnReqs); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	return diags
}

func resourcePingAccessAuthnReqListReadData(d *schema.ResourceData) *models.AuthnReqListView {
	auths := expandStringList(d.Get("authn_reqs").([]interface{}))
	engine := &models.AuthnReqListView{
		Name:      String(d.Get("name").(string)),
		AuthnReqs: &auths,
	}

	return engine
}
