package pingaccess

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/iwarapter/pingaccess-sdk-go/pingaccess"
)

func resourcePingAccessVirtualHost() *schema.Resource {
	return &schema.Resource{
		Create: resourcePingAccessVirtualHostCreate,
		Read:   resourcePingAccessVirtualHostRead,
		Update: resourcePingAccessVirtualHostUpdate,
		Delete: resourcePingAccessVirtualHostDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: resourcePingAccessVirtualHostSchema(),
	}
}

func resourcePingAccessVirtualHostSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		agentResourceCacheTtl: &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
		},
		host: &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		keyPairID: &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
		},
		port: &schema.Schema{
			Type:     schema.TypeInt,
			Required: true,
		},
		trustedCertificateGroupID: &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
		},
	}
}

func resourcePingAccessVirtualHostCreate(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pingaccess.Client).Virtualhosts
	input := pingaccess.AddVirtualHostCommandInput{
		Body: *resourcePingAccessVirtualHostReadData(d),
	}

	result, _, err := svc.AddVirtualHostCommand(&input)
	if err != nil {
		return fmt.Errorf("Error creating virtualhost: %s", err)
	}

	d.SetId(result.Id.String())
	return resourcePingAccessVirtualHostReadResult(d, &input.Body)
}

func resourcePingAccessVirtualHostRead(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pingaccess.Client).Virtualhosts
	input := &pingaccess.GetVirtualHostCommandInput{
		Id: d.Id(),
	}
	result, _, err := svc.GetVirtualHostCommand(input)
	if err != nil {
		return fmt.Errorf("Error reading virtualhost: %s", err)
	}
	return resourcePingAccessVirtualHostReadResult(d, result)
}

func resourcePingAccessVirtualHostUpdate(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pingaccess.Client).Virtualhosts
	input := pingaccess.UpdateVirtualHostCommandInput{
		Body: *resourcePingAccessVirtualHostReadData(d),
		Id:   d.Id(),
	}

	result, _, err := svc.UpdateVirtualHostCommand(&input)
	if err != nil {
		return fmt.Errorf("Error updating virtualhost: %s", err)
	}
	return resourcePingAccessVirtualHostReadResult(d, result)
}

func resourcePingAccessVirtualHostDelete(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pingaccess.Client).Virtualhosts
	input := &pingaccess.DeleteVirtualHostCommandInput{
		Id: d.Id(),
	}

	_, err := svc.DeleteVirtualHostCommand(input)
	if err != nil {
		return fmt.Errorf("Error deleting virtualhost: %s", err)
	}
	return nil
}

func resourcePingAccessVirtualHostReadResult(d *schema.ResourceData, input *pingaccess.VirtualHostView) error {
	setResourceDataString(d, "host", input.Host)
	setResourceDataInt(d, "port", input.Port)
	setResourceDataInt(d, "agent_resource_cache_ttl", input.AgentResourceCacheTTL)
	setResourceDataInt(d, "key_pair_id", input.KeyPairId)
	setResourceDataInt(d, "trusted_certificate_group_id", input.TrustedCertificateGroupId)

	return nil
}

func resourcePingAccessVirtualHostReadData(d *schema.ResourceData) *pingaccess.VirtualHostView {
	vh := &pingaccess.VirtualHostView{
		Host: String(d.Get("host").(string)),
		Port: Int(d.Get("port").(int)),
	}

	if _, ok := d.GetOkExists("agent_resource_cache_ttl"); ok {
		vh.AgentResourceCacheTTL = Int(d.Get("agent_resource_cache_ttl").(int))
	}

	if _, ok := d.GetOkExists("key_pair_id"); ok {
		vh.KeyPairId = Int(d.Get("key_pair_id").(int))
	}

	if _, ok := d.GetOkExists("trusted_certificate_group_id"); ok {
		vh.TrustedCertificateGroupId = Int(d.Get("trusted_certificate_group_id").(int))
	}

	return vh
}
