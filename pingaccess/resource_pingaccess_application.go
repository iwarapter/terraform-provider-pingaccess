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
	// PolicyItem := &schema.Resource{
	// 	Schema: map[string]*schema.Schema{
	// 		"type": {
	// 			Type:     schema.TypeString,
	// 			Optional: true,
	// 		},
	// 		"id": {
	// 			Type:     schema.TypeString,
	// 			Optional: true,
	// 		},
	// 	},
	// }

	return &schema.Resource{
		Create: resourcePingAccessApplicationCreate,
		Read:   resourcePingAccessApplicationRead,
		Update: resourcePingAccessApplicationUpdate,
		Delete: resourcePingAccessApplicationDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			accessValidatorID: &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			agentID: &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			applicationType: &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			caseSensitivePath: &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			contextRoot: &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			defaultAuthType: &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			description: &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			destination: &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			enabled: &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			identityMappingIds: &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 2,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"web": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"api": {
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			name: &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			policy: &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"web": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": {
										Type:     schema.TypeString,
										Required: true,
									},
									"id": {
										Type:     schema.TypeString,
										Required: true,
									},
								},
							},
						},
						"api": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": {
										Type:     schema.TypeString,
										Required: true,
									},
									"id": {
										Type:     schema.TypeString,
										Required: true,
									},
								},
							},
						},
					},
				},
			},
			realm: &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			requireHTTPS: &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			siteID: &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			virtualHostIDs: &schema.Schema{
				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			webSessionId: &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourcePingAccessApplicationCreate(d *schema.ResourceData, m interface{}) error {
	log.Printf("[INFO] resourcePingAccessApplicationCreate")
	access_validator_id := d.Get("access_validator_id").(int)
	application_type := d.Get(applicationType).(string)
	agent_id := d.Get(agentID).(int)
	case_sensitive_path := d.Get(caseSensitivePath).(bool)
	context_root := d.Get(contextRoot).(string)
	default_auth_type := d.Get(defaultAuthType).(string)
	description := d.Get(description).(string)
	destination := d.Get(destination).(string)
	enabled := d.Get(enabled).(bool)
	// identity_mapping_ids := magic
	name := d.Get(name).(string)
	// policy := magic
	realm := d.Get(realm).(string)
	require_https := d.Get(requireHTTPS).(bool)
	site_id := d.Get(siteID).(string)
	virtual_host_ids := expandStringList(d.Get(virtualHostIDs).(*schema.Set).List())
	web_session_id := d.Get(webSessionId).(int)

	//TODO fix this dirty dirty hack
	vh_ids := []*int{}
	for _, i := range virtual_host_ids {
		text, _ := strconv.Atoi(*i)
		vh_ids = append(vh_ids, &text)
	}

	siteID, _ := strconv.Atoi(site_id)

	input := &pingaccess.AddApplicationCommandInput{
		Body: pingaccess.ApplicationView{
			AccessValidatorId: Int(access_validator_id),
			ApplicationType:   String(application_type),
			AgentId:           Int(agent_id),
			CaseSensitivePath: Bool(case_sensitive_path),
			ContextRoot:       String(context_root),
			DefaultAuthType:   String(default_auth_type),
			Description:       String(description),
			Destination:       String(destination),
			Enabled:           Bool(enabled),
			//IdentityMappingIds: magic
			Name: String(name),
			//Policy: magic
			Realm:          String(realm),
			RequireHTTPS:   Bool(require_https),
			SiteId:         Int(siteID),
			VirtualHostIds: &vh_ids,
			WebSessionId:   Int(web_session_id),
		},
	}

	if _, ok := d.GetOk(policy); ok {
		policySet := d.Get(policy).([]interface{})

		webPolicies := make([]*pingaccess.PolicyItem, 0)
		apiPolicies := make([]*pingaccess.PolicyItem, 0)

		policy := policySet[0].(map[string]interface{})
		for _, pV := range policy["web"].(*schema.Set).List() {
			p := pV.(map[string]interface{})
			webPolicies = append(webPolicies, &pingaccess.PolicyItem{
				Id:   json.Number(p["id"].(string)),
				Type: String(p["type"].(string)),
			})
		}
		for _, pV := range policy["api"].(*schema.Set).List() {
			p := pV.(map[string]interface{})
			apiPolicies = append(apiPolicies, &pingaccess.PolicyItem{
				Id:   json.Number(p["id"].(string)),
				Type: String(p["type"].(string)),
			})
		}
		policies := map[string]*[]*pingaccess.PolicyItem{
			"Web": &webPolicies,
			"API": &apiPolicies,
		}
		input.Body.Policy = policies
	}

	svc := m.(*pingaccess.Client).Applications

	result, _, err := svc.AddApplicationCommand(input)
	if err != nil {
		return fmt.Errorf("Error creating application: %s", err)
	}

	d.SetId(result.Id.String())
	return resourcePingAccessApplicationReadResult(d, &input.Body)
}

func resourcePingAccessApplicationRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("[INFO] resourcePingAccessApplicationRead")
	svc := m.(*pingaccess.Client).Applications

	input := &pingaccess.GetApplicationCommandInput{
		Id: d.Id(),
	}

	log.Printf("[INFO] ResourceID: %s", d.Id())
	log.Printf("[INFO] GetApplicationCommandInput: %s", input.Id)
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
	application_type := d.Get(applicationType).(string)
	agent_id := d.Get(agentID).(int)
	case_sensitive_path := d.Get(caseSensitivePath).(bool)
	context_root := d.Get(contextRoot).(string)
	default_auth_type := d.Get(defaultAuthType).(string)
	description := d.Get(description).(string)
	destination := d.Get(destination).(string)
	enabled := d.Get(enabled).(bool)
	// identity_mapping_ids := magic
	name := d.Get(name).(string)
	// policy := magic
	realm := d.Get(realm).(string)
	require_https := d.Get(requireHTTPS).(bool)
	site_id := d.Get(siteID).(string)
	virtual_host_ids := expandStringList(d.Get(virtualHostIDs).(*schema.Set).List())
	web_session_id := d.Get(webSessionId).(int)

	//TODO fix this dirty dirty hack
	vh_ids := []*int{}
	for _, i := range virtual_host_ids {
		text, _ := strconv.Atoi(*i)
		vh_ids = append(vh_ids, &text)
	}

	siteID, _ := strconv.Atoi(site_id)

	input := pingaccess.UpdateApplicationCommandInput{
		Body: pingaccess.ApplicationView{
			AccessValidatorId: Int(access_validator_id),
			ApplicationType:   String(application_type),
			AgentId:           Int(agent_id),
			CaseSensitivePath: Bool(case_sensitive_path),
			ContextRoot:       String(context_root),
			DefaultAuthType:   String(default_auth_type),
			Description:       String(description),
			Destination:       String(destination),
			Enabled:           Bool(enabled),
			//IdentityMappingIds: magic
			Name: String(name),
			//Policy: magic
			Realm:          String(realm),
			RequireHTTPS:   Bool(require_https),
			SiteId:         Int(siteID),
			VirtualHostIds: &vh_ids,
			WebSessionId:   Int(web_session_id),
		},
	}
	input.Id = d.Id()

	if _, ok := d.GetOk(policy); ok {
		policySet := d.Get(policy).([]interface{})

		webPolicies := make([]*pingaccess.PolicyItem, 0)
		apiPolicies := make([]*pingaccess.PolicyItem, 0)

		policy := policySet[0].(map[string]interface{})
		for _, pV := range policy["web"].(*schema.Set).List() {
			p := pV.(map[string]interface{})
			webPolicies = append(webPolicies, &pingaccess.PolicyItem{
				Id:   json.Number(p["id"].(string)),
				Type: String(p["type"].(string)),
			})
		}
		for _, pV := range policy["api"].(*schema.Set).List() {
			p := pV.(map[string]interface{})
			apiPolicies = append(apiPolicies, &pingaccess.PolicyItem{
				Id:   json.Number(p["id"].(string)),
				Type: String(p["type"].(string)),
			})
		}
		policies := map[string]*[]*pingaccess.PolicyItem{
			"Web": &webPolicies,
			"API": &apiPolicies,
		}
		input.Body.Policy = policies
	}

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
		Id: d.Id(),
	}

	log.Printf("[INFO] ResourceID: %s", d.Id())
	log.Printf("[INFO] DeleteApplicationCommandInput: %s", input.Id)
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
