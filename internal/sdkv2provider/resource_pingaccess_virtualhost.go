package sdkv2provider

import (
	"context"

	"github.com/iwarapter/pingaccess-sdk-go/pingaccess/models"
	"github.com/iwarapter/pingaccess-sdk-go/services/virtualhosts"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePingAccessVirtualHost() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePingAccessVirtualHostCreate,
		ReadContext:   resourcePingAccessVirtualHostRead,
		UpdateContext: resourcePingAccessVirtualHostUpdate,
		DeleteContext: resourcePingAccessVirtualHostDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: resourcePingAccessVirtualHostSchema(),
	}
}

func resourcePingAccessVirtualHostSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"agent_resource_cache_ttl": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  0,
		},
		"host": {
			Type:     schema.TypeString,
			Required: true,
		},
		"key_pair_id": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  0,
		},
		"port": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"trusted_certificate_group_id": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  0,
		},
	}
}

func resourcePingAccessVirtualHostCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(paClient).Virtualhosts
	input := virtualhosts.AddVirtualHostCommandInput{
		Body: *resourcePingAccessVirtualHostReadData(d),
	}

	result, _, err := svc.AddVirtualHostCommand(&input)
	if err != nil {
		return diag.Errorf("unable to create VirtualHost: %s", err)
	}

	d.SetId(result.Id.String())
	return resourcePingAccessVirtualHostReadResult(d, &input.Body)
}

func resourcePingAccessVirtualHostRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(paClient).Virtualhosts
	input := &virtualhosts.GetVirtualHostCommandInput{
		Id: d.Id(),
	}
	result, _, err := svc.GetVirtualHostCommand(input)
	if err != nil {
		return diag.Errorf("unable to read VirtualHost: %s", err)
	}
	return resourcePingAccessVirtualHostReadResult(d, result)
}

func resourcePingAccessVirtualHostUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(paClient).Virtualhosts
	input := virtualhosts.UpdateVirtualHostCommandInput{
		Body: *resourcePingAccessVirtualHostReadData(d),
		Id:   d.Id(),
	}

	result, _, err := svc.UpdateVirtualHostCommand(&input)
	if err != nil {
		return diag.Errorf("unable to update VirtualHost: %s", err)
	}
	return resourcePingAccessVirtualHostReadResult(d, result)
}

func resourcePingAccessVirtualHostDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(paClient).Virtualhosts
	input := &virtualhosts.DeleteVirtualHostCommandInput{
		Id: d.Id(),
	}

	_, err := svc.DeleteVirtualHostCommand(input)
	if err != nil {
		return diag.Errorf("unable to delete VirtualHost: %s", err)
	}
	return nil
}

func resourcePingAccessVirtualHostReadResult(d *schema.ResourceData, input *models.VirtualHostView) diag.Diagnostics {
	var diags diag.Diagnostics
	setResourceDataStringWithDiagnostic(d, "host", input.Host, &diags)
	setResourceDataIntWithDiagnostic(d, "port", input.Port, &diags)
	setResourceDataIntWithDiagnostic(d, "agent_resource_cache_ttl", input.AgentResourceCacheTTL, &diags)
	setResourceDataIntWithDiagnostic(d, "key_pair_id", input.KeyPairId, &diags)
	setResourceDataIntWithDiagnostic(d, "trusted_certificate_group_id", input.TrustedCertificateGroupId, &diags)

	return diags
}

func resourcePingAccessVirtualHostReadData(d *schema.ResourceData) *models.VirtualHostView {
	vh := &models.VirtualHostView{
		Host: String(d.Get("host").(string)),
		Port: Int(d.Get("port").(int)),
	}
	vh.AgentResourceCacheTTL = Int(d.Get("agent_resource_cache_ttl").(int))
	vh.KeyPairId = Int(d.Get("key_pair_id").(int))
	vh.TrustedCertificateGroupId = Int(d.Get("trusted_certificate_group_id").(int))

	return vh
}
