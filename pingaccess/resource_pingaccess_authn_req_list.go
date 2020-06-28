package pingaccess

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/iwarapter/pingaccess-sdk-go/pingaccess"
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
	svc := m.(*pingaccess.Client).AuthnReqLists
	input := pingaccess.AddAuthnReqListCommandInput{
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
	svc := m.(*pingaccess.Client).AuthnReqLists
	input := &pingaccess.GetAuthnReqListCommandInput{
		Id: d.Id(),
	}
	result, _, err := svc.GetAuthnReqListCommand(input)
	if err != nil {
		return diag.Errorf("unable to read AuthnReqList: %s", err)
	}
	return resourcePingAccessAuthnReqListReadResult(d, result)
}

func resourcePingAccessAuthnReqListUpdate(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(*pingaccess.Client).AuthnReqLists
	input := pingaccess.UpdateAuthnReqListCommandInput{
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
	svc := m.(*pingaccess.Client).AuthnReqLists
	input := &pingaccess.DeleteAuthnReqListCommandInput{
		Id: d.Id(),
	}

	_, err := svc.DeleteAuthnReqListCommand(input)
	if err != nil {
		return diag.Errorf("unable to delete AuthnReqList: %s", err)
	}
	return nil
}

func resourcePingAccessAuthnReqListReadResult(d *schema.ResourceData, input *pingaccess.AuthnReqListView) diag.Diagnostics {
	var diags diag.Diagnostics
	setResourceDataStringWithDiagnostic(d, "name", input.Name, &diags)
	if err := d.Set("authn_reqs", input.AuthnReqs); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	return diags
}

func resourcePingAccessAuthnReqListReadData(d *schema.ResourceData) *pingaccess.AuthnReqListView {
	auths := expandStringList(d.Get("authn_reqs").([]interface{}))
	engine := &pingaccess.AuthnReqListView{
		Name:      String(d.Get("name").(string)),
		AuthnReqs: &auths,
	}

	return engine
}
