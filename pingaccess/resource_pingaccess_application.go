package pingaccess

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/iwarapter/pingaccess-sdk-go/pingaccess"
)

func resourcePingAccessApplication() *schema.Resource {
	return &schema.Resource{
		Create: resourcePingAccessApplicationCreate,
		Read:   resourcePingAccessApplicationRead,
		Update: resourcePingAccessApplicationUpdate,
		Delete: resourcePingAccessApplicationDelete,

		Schema: map[string]*schema.Schema{
			"access_validator_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"application_type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"agent_id": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"context_root": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"default_auth_type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"destination": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"site_id": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"virtual_host_ids": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			// AccessValidatorID  int                     `json:"accessValidatorId"`
			// AgentID            int                     `json:"agentId"`
			// ApplicationType    string                  `json:"applicationType"`
			// CaseSensitivePath  bool                    `json:"caseSensitivePath"`
			// ContextRoot        string                  `json:"contextRoot"`
			// DefaultAuthType    string                  `json:"defaultAuthType"`
			// Description        string                  `json:"description"`
			// Destination        string                  `json:"destination"`
			// Enabled            bool                    `json:"enabled"`
			// IdentityMappingIds map[string]int          `json:"identityMappingIds"`
			// Name               string                  `json:"name"`
			// Policy             map[string][]PolicyItem `json:"policy"`
			// Realm              string                  `json:"realm"`
			// RequireHTTPS       bool                    `json:"requireHTTPS"`
			// SiteID             int                     `json:"siteId"`
			// VirtualHostIds     []int                   `json:"virtualHostIds"`
			// WebSessionID       int                     `json:"webSessionId"`
		},
	}
}

func resourcePingAccessApplicationCreate(d *schema.ResourceData, m interface{}) error {
	log.Printf("[INFO] resourcePingAccessApplicationCreate")
	access_validator_id := d.Get("access_validator_id").(int)
	application_type := d.Get("application_type").(string)
	agent_id := d.Get("agent_id").(int)
	context_root := d.Get("context_root").(string)
	default_auth_type := d.Get("default_auth_type").(string)
	destination := d.Get("destination").(string)
	name := d.Get("name").(string)
	site_id := d.Get("site_id").(int)
	virtual_host_ids := expandStringList(d.Get("virtual_host_ids").(*schema.Set).List())

	//TODO fix this dirty dirty hack
	vh_ids := []*int{}
	for _, i := range virtual_host_ids {
		number := *i
		text, _ := strconv.Atoi(number)
		vh_ids = append(vh_ids, &text)
	}

	input := pingaccess.AddApplicationCommandInput{
		Body: pingaccess.ApplicationView{
			AccessValidatorId: access_validator_id,
			ApplicationType:   application_type,
			AgentId:           agent_id,
			ContextRoot:       context_root,
			DefaultAuthType:   default_auth_type,
			Destination:       destination,
			Name:              name,
			SiteId:            site_id,
			VirtualHostIds:    vh_ids,
		},
	}

	svc := m.(*pingaccess.Client).Applications

	result, _, err := svc.AddApplicationCommand(&input)
	if err != nil {
		return fmt.Errorf("Error creating application: %s", err)
	}

	d.SetId(strconv.Itoa(result.Id))
	return resourcePingAccessApplicationReadResult(d, &input.Body)
}

func resourcePingAccessApplicationRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("[INFO] resourcePingAccessApplicationRead")
	svc := m.(*pingaccess.Client).Applications

	input := &pingaccess.GetApplicationCommandInput{
		Path: struct {
			Id string
		}{
			Id: d.Id(),
		},
	}

	log.Printf("[INFO] ResourceID: %s", d.Id())
	log.Printf("[INFO] GetApplicationCommandInput: %s", input.Path.Id)
	result, _, _ := svc.GetApplicationCommand(input)
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(result)
	vh := pingaccess.ApplicationView{}
	json.NewDecoder(b).Decode(&vh)

	return resourcePingAccessApplicationReadResult(d, &vh)
}

func resourcePingAccessApplicationUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("[INFO] resourcePingAccessApplicationUpdate")
	access_validator_id := d.Get("access_validator_id").(int)
	application_type := d.Get("application_type").(string)
	agent_id := d.Get("agent_id").(int)
	context_root := d.Get("context_root").(string)
	default_auth_type := d.Get("default_auth_type").(string)
	destination := d.Get("destination").(string)
	name := d.Get("name").(string)
	site_id := d.Get("site_id").(int)
	virtual_host_ids := expandStringList(d.Get("virtual_host_ids").(*schema.Set).List())

	//TODO fix this dirty dirty hack
	vh_ids := []*int{}
	for _, i := range virtual_host_ids {
		number := *i
		text, _ := strconv.Atoi(number)
		vh_ids = append(vh_ids, &text)
	}

	input := pingaccess.UpdateApplicationCommandInput{
		Body: pingaccess.ApplicationView{
			AccessValidatorId: access_validator_id,
			ApplicationType:   application_type,
			AgentId:           agent_id,
			ContextRoot:       context_root,
			DefaultAuthType:   default_auth_type,
			Destination:       destination,
			Name:              name,
			SiteId:            site_id,
			VirtualHostIds:    vh_ids,
		},
	}
	input.Path.Id = d.Id()

	svc := m.(*pingaccess.Client).Applications

	_, _, err := svc.UpdateApplicationCommand(&input)
	if err != nil {
		return fmt.Errorf("Error updating application: %s", err)
	}
	log.Println("[DEBUG] End - resourcePingAccessApplicationUpdate")
	return nil
}

func resourcePingAccessApplicationDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("[INFO] resourcePingAccessApplicationDelete")
	svc := m.(*pingaccess.Client).Applications

	input := &pingaccess.DeleteApplicationCommandInput{
		Path: struct {
			Id string
		}{
			Id: d.Id(),
		},
	}

	log.Printf("[INFO] ResourceID: %s", d.Id())
	log.Printf("[INFO] DeleteApplicationCommandInput: %s", input.Path.Id)
	_, _, err := svc.DeleteApplicationCommand(input)
	if err != nil {
		return fmt.Errorf("Error deleting virtualhost: %s", err)
	}
	log.Println("[DEBUG] End - resourcePingAccessSiteDelete")
	return nil
}

func resourcePingAccessApplicationReadResult(d *schema.ResourceData, rv *pingaccess.ApplicationView) error {
	log.Printf("[INFO] resourcePingAccessApplicationReadResult")
	// if err := d.Set("name", rv.Name); err != nil {
	// 	return err
	// }
	// if err := d.Set("class_name", rv.ClassName); err != nil {
	// 	return err
	// }
	// if err := d.Set("supported_destinations", rv.SupportedDestinations); err != nil {
	// 	return err
	// }
	// // if err := d.Set("configuration", rv.Configuration); err != nil {
	// 	return err
	// }
	return nil
}
