package pingaccess

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/iwarapter/pingaccess-sdk-go/service/virtualhosts"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourcePingAccessVirtualHost() *schema.Resource {
	return &schema.Resource{
		Create: resourcePingAccessVirtualHostCreate,
		Read:   resourcePingAccessVirtualHostRead,
		Update: resourcePingAccessVirtualHostUpdate,
		Delete: resourcePingAccessVirtualHostDelete,

		Schema: map[string]*schema.Schema{
			"agent_resource_cache_ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"host": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"key_pair_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"trusted_certificate_group_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourcePingAccessVirtualHostCreate(d *schema.ResourceData, m interface{}) error {
	log.Println("[INFO] resourcePingAccessVirtualHostCreate")
	agent_resource_cache_ttl := d.Get("agent_resource_cache_ttl").(int)
	host := d.Get("host").(string)
	key_pair_id := d.Get("key_pair_id").(int)
	port := d.Get("port").(int)
	trusted_certificate_group_id := d.Get("trusted_certificate_group_id").(int)

	input := virtualhosts.AddVirtualHostCommandInput{
		Body: virtualhosts.VirtualHostView{
			AgentResourceCacheTTL: agent_resource_cache_ttl,
			Host:      host,
			KeyPairId: key_pair_id,
			Port:      port,
			TrustedCertificateGroupId: trusted_certificate_group_id,
		},
	}

	svc := m.(*PAClient).vhconn
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	res, err := svc.AddVirtualHostCommand(&input) //.CreateVirtualHost(rv)
	if err != nil {
		return fmt.Errorf("Error creating virtualhost: %s", err)
	}

	d.SetId(strconv.Itoa(res.Id))
	return resourcePingAccessVirtualHostReadResult(d, &input.Body)
}

func resourcePingAccessVirtualHostRead(d *schema.ResourceData, m interface{}) error {
	log.Println("[INFO] Start - resourcePingAccessVirtualHostRead")
	svc := m.(*PAClient).vhconn
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	input := &virtualhosts.GetVirtualHostCommandInput{
		Path: struct {
			Id string
		}{
			Id: d.Id(),
		},
	}

	log.Printf("[INFO] ResourceID: %s", d.Id())
	log.Printf("[INFO] GetVirtualHostCommandInput: %s", input.Path.Id)
	resp, _ := svc.GetVirtualHostCommand(input)
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(resp)
	vh := virtualhosts.VirtualHostView{}
	json.NewDecoder(b).Decode(&vh)

	log.Println("[INFO] End - resourcePingAccessVirtualHostRead")
	return resourcePingAccessVirtualHostReadResult(d, &vh)
}

func resourcePingAccessVirtualHostUpdate(d *schema.ResourceData, m interface{}) error {
	log.Println("[INFO] Start - resourcePingAccessVirtualHostUpdate")
	agent_resource_cache_ttl := d.Get("agent_resource_cache_ttl").(int)
	host := d.Get("host").(string)
	key_pair_id := d.Get("key_pair_id").(int)
	port := d.Get("port").(int)
	trusted_certificate_group_id := d.Get("trusted_certificate_group_id").(int)

	input := virtualhosts.UpdateVirtualHostCommandInput{
		Body: virtualhosts.VirtualHostView{
			AgentResourceCacheTTL: agent_resource_cache_ttl,
			Host:      host,
			KeyPairId: key_pair_id,
			Port:      port,
			TrustedCertificateGroupId: trusted_certificate_group_id,
		},
	}
	input.Path.Id = d.Id()

	svc := m.(*PAClient).vhconn
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	_, err := svc.UpdateVirtualHostCommand(&input)
	if err != nil {
		return fmt.Errorf("Error updating virtualhost: %s", err)
	}
	log.Println("[INFO] End - resourcePingAccessVirtualHostUpdate")
	return nil
}

func resourcePingAccessVirtualHostDelete(d *schema.ResourceData, m interface{}) error {
	log.Println("[INFO] Start - resourcePingAccessVirtualHostDelete")
	svc := m.(*PAClient).vhconn
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	input := &virtualhosts.DeleteVirtualHostCommandInput{
		Path: struct {
			Id string
		}{
			Id: d.Id(),
		},
	}

	log.Printf("[INFO] ResourceID: %s", d.Id())
	log.Printf("[INFO] DeleteVirtualHostCommandInput: %s", input.Path.Id)
	err := svc.DeleteVirtualHostCommand(input)
	if err != nil {
		return fmt.Errorf("Error deleting virtualhost: %s", err)
	}
	log.Println("[INFO] End - resourcePingAccessVirtualHostDelete")
	return nil
}

func resourcePingAccessVirtualHostReadResult(d *schema.ResourceData, input *virtualhosts.VirtualHostView) error {
	log.Println("[INFO] Start - resourcePingAccessVirtualHostReadResult")
	if err := d.Set("agent_resource_cache_ttl", input.AgentResourceCacheTTL); err != nil {
		return err
	}
	if err := d.Set("host", input.Host); err != nil {
		return err
	}
	if err := d.Set("key_pair_id", input.KeyPairId); err != nil {
		return err
	}
	if err := d.Set("port", input.Port); err != nil {
		return err
	}
	if err := d.Set("trusted_certificate_group_id", input.TrustedCertificateGroupId); err != nil {
		return err
	}

	log.Println("[INFO] End - resourcePingAccessVirtualHostReadResult")
	return nil
}
