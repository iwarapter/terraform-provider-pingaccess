package pingaccess

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"

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

		Schema: map[string]*schema.Schema{
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
		},
	}
}

func resourcePingAccessVirtualHostCreate(d *schema.ResourceData, m interface{}) error {
	log.Println("[INFO] resourcePingAccessVirtualHostCreate")
	agent_resource_cache_ttl := d.Get(agentResourceCacheTtl).(int)
	host := d.Get(host).(string)
	key_pair_id := d.Get(keyPairID).(int)
	port := d.Get(port).(int)
	trusted_certificate_group_id := d.Get(trustedCertificateGroupID).(int)

	input := pingaccess.AddVirtualHostCommandInput{
		Body: pingaccess.VirtualHostView{
			AgentResourceCacheTTL:     Int(agent_resource_cache_ttl),
			Host:                      String(host),
			KeyPairId:                 Int(key_pair_id),
			Port:                      Int(port),
			TrustedCertificateGroupId: Int(trusted_certificate_group_id),
		},
	}

	svc := m.(*pingaccess.Client).Virtualhosts

	result, _, err := svc.AddVirtualHostCommand(&input) //.CreateVirtualHost(rv)
	if err != nil {
		return fmt.Errorf("Error creating virtualhost: %s", err)
	}

	d.SetId(result.Id.String())
	return resourcePingAccessVirtualHostReadResult(d, &input.Body)
}

func resourcePingAccessVirtualHostRead(d *schema.ResourceData, m interface{}) error {
	log.Println("[INFO] Start - resourcePingAccessVirtualHostRead")
	svc := m.(*pingaccess.Client).Virtualhosts

	input := &pingaccess.GetVirtualHostCommandInput{
		Id: d.Id(),
	}

	log.Printf("[INFO] ResourceID: %s", d.Id())
	log.Printf("[INFO] GetVirtualHostCommandInput: %s", input.Id)
	result, _, _ := svc.GetVirtualHostCommand(input)
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(result)
	vh := pingaccess.VirtualHostView{}
	json.NewDecoder(b).Decode(&vh)

	log.Println("[INFO] End - resourcePingAccessVirtualHostRead")
	return resourcePingAccessVirtualHostReadResult(d, &vh)
}

func resourcePingAccessVirtualHostUpdate(d *schema.ResourceData, m interface{}) error {
	log.Println("[INFO] Start - resourcePingAccessVirtualHostUpdate")
	agent_resource_cache_ttl := d.Get(agentResourceCacheTtl).(int)
	host := d.Get(host).(string)
	key_pair_id := d.Get(keyPairID).(int)
	port := d.Get(port).(int)
	trusted_certificate_group_id := d.Get(trustedCertificateGroupID).(int)

	input := pingaccess.UpdateVirtualHostCommandInput{
		Body: pingaccess.VirtualHostView{
			AgentResourceCacheTTL:     Int(agent_resource_cache_ttl),
			Host:                      String(host),
			KeyPairId:                 Int(key_pair_id),
			Port:                      Int(port),
			TrustedCertificateGroupId: Int(trusted_certificate_group_id),
		},
	}
	input.Id = d.Id()

	svc := m.(*pingaccess.Client).Virtualhosts

	_, _, err := svc.UpdateVirtualHostCommand(&input)
	if err != nil {
		return fmt.Errorf("Error updating virtualhost: %s", err)
	}
	log.Println("[INFO] End - resourcePingAccessVirtualHostUpdate")
	return nil
}

func resourcePingAccessVirtualHostDelete(d *schema.ResourceData, m interface{}) error {
	log.Println("[INFO] Start - resourcePingAccessVirtualHostDelete")
	svc := m.(*pingaccess.Client).Virtualhosts

	input := &pingaccess.DeleteVirtualHostCommandInput{
		Id: d.Id(),
	}

	log.Printf("[INFO] ResourceID: %s", d.Id())
	log.Printf("[INFO] DeleteVirtualHostCommandInput: %s", input.Id)
	_, err := svc.DeleteVirtualHostCommand(input)
	if err != nil {
		return fmt.Errorf("Error deleting virtualhost: %s", err)
	}
	log.Println("[INFO] End - resourcePingAccessVirtualHostDelete")
	return nil
}

func resourcePingAccessVirtualHostReadResult(d *schema.ResourceData, input *pingaccess.VirtualHostView) error {
	log.Println("[INFO] Start - resourcePingAccessVirtualHostReadResult")
	if err := d.Set(agentResourceCacheTtl, input.AgentResourceCacheTTL); err != nil {
		return err
	}
	if err := d.Set(host, input.Host); err != nil {
		return err
	}
	if err := d.Set(keyPairID, input.KeyPairId); err != nil {
		return err
	}
	if err := d.Set(port, input.Port); err != nil {
		return err
	}
	if err := d.Set(trustedCertificateGroupID, input.TrustedCertificateGroupId); err != nil {
		return err
	}

	log.Println("[INFO] End - resourcePingAccessVirtualHostReadResult")
	return nil
}
