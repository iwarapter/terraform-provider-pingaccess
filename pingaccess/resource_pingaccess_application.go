package pingaccess

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"
	pa "github.com/iwarapter/pingaccess-sdk-go/pingaccess"
)

func resourcePingAccessApplication() *schema.Resource {
	return &schema.Resource{
		Create: resourcePingAccessApplicationCreate,
		Read:   resourcePingAccessApplicationRead,
		Update: resourcePingAccessApplicationUpdate,
		Delete: resourcePingAccessApplicationDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"access_validator_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"agent_id": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"application_type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"case_sensitive_path": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
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
			"identity_mapping_ids": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				MinItems: 0,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"web": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "0",
							// DefaultFunc: func() (interface{}, error) { return "0", nil },
						},
						"api": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "0",
							// DefaultFunc: func() (interface{}, error) { return "0", nil },
						},
					},
				},
			},
			name: &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"policy": applicationPolicySchema(),
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
			"web_session_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "0",
			},
		},
	}
}

func resourcePingAccessApplicationCreate(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pa.Client).Applications
	input := &pa.AddApplicationCommandInput{
		Body: *resourcePingAccessApplicationReadData(d),
	}
	result, _, err := svc.AddApplicationCommand(input)
	if err != nil {
		return fmt.Errorf("Error creating application: %s", err)
	}

	d.SetId(result.Id.String())
	return resourcePingAccessApplicationReadResult(d, &input.Body)
}

func resourcePingAccessApplicationRead(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pa.Client).Applications
	input := &pa.GetApplicationCommandInput{
		Id: d.Id(),
	}
	result, _, err := svc.GetApplicationCommand(input)
	if err != nil {
		return fmt.Errorf("Error reading application: %s", err)
	}
	return resourcePingAccessApplicationReadResult(d, result)
}

func resourcePingAccessApplicationUpdate(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pa.Client).Applications
	input := pa.UpdateApplicationCommandInput{
		Body: *resourcePingAccessApplicationReadData(d),
		Id:   d.Id(),
	}
	result, _, err := svc.UpdateApplicationCommand(&input)
	if err != nil {
		return fmt.Errorf("Error updating application: %s", err)
	}
	return resourcePingAccessApplicationReadResult(d, result)
}

func resourcePingAccessApplicationDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("[INFO] resourcePingAccessApplicationDelete")
	svc := m.(*pa.Client).Applications

	input := &pa.DeleteApplicationCommandInput{
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

func resourcePingAccessApplicationReadResult(d *schema.ResourceData, rv *pa.ApplicationView) (err error) {
	setResourceDataInt(d, "access_validator_id", rv.AccessValidatorId)
	setResourceDataInt(d, "agent_id", rv.AgentId)
	setResourceDataString(d, "application_type", rv.ApplicationType)
	setResourceDataBool(d, "case_sensitive_path", rv.CaseSensitivePath)
	setResourceDataString(d, "context_root", rv.ContextRoot)
	setResourceDataString(d, "default_auth_type", rv.DefaultAuthType)
	setResourceDataString(d, "description", rv.Description)
	setResourceDataString(d, "destination", rv.Destination)
	setResourceDataBool(d, "enabled", rv.Enabled)
	setResourceDataString(d, "name", rv.Name)
	setResourceDataString(d, "realm", rv.Realm)
	setResourceDataBool(d, "require_https", rv.RequireHTTPS)
	siteID := strconv.Itoa(*rv.SiteId)
	setResourceDataString(d, "site_id", &siteID)

	if rv.VirtualHostIds != nil {
		vhs := []string{}
		for _, vh := range *rv.VirtualHostIds {
			vhs = append(vhs, strconv.Itoa(*vh))
		}
		if err := d.Set("virtual_host_ids", vhs); err != nil {
			return err
		}
	}

	if rv.WebSessionId != nil {
		d.Set("web_session_id", strconv.Itoa(*rv.WebSessionId))
	}

	if rv.IdentityMappingIds != nil && (*rv.IdentityMappingIds["Web"] != 0 || *rv.IdentityMappingIds["API"] != 0) {
		if err = d.Set("identity_mapping_ids", flattenIdentityMappingIds(rv.IdentityMappingIds)); err != nil {
			return err
		}
	}
	if rv.Policy != nil && (len(*rv.Policy["API"]) > 0 || len(*rv.Policy["Web"]) > 0) {
		if err := d.Set("policy", flattenPolicy(rv.Policy)); err != nil {
			return err
		}
	}
	return nil
}

func resourcePingAccessApplicationReadData(d *schema.ResourceData) *pa.ApplicationView {
	siteID, _ := strconv.Atoi(d.Get("site_id").(string))
	virtualHostIds := expandStringList(d.Get(virtualHostIDs).(*schema.Set).List())
	vhIds := []*int{}
	for _, i := range virtualHostIds {
		text, _ := strconv.Atoi(*i)
		vhIds = append(vhIds, &text)
	}

	application := &pa.ApplicationView{
		AgentId:         Int(d.Get("agent_id").(int)),
		Name:            String(d.Get("name").(string)),
		ApplicationType: String(d.Get("application_type").(string)),
		ContextRoot:     String(d.Get("context_root").(string)),
		DefaultAuthType: String(d.Get("default_auth_type").(string)),
		Destination:     String(d.Get("destination").(string)),
		SiteId:          Int(siteID),
		VirtualHostIds:  &vhIds,
	}

	if _, ok := d.GetOkExists("access_validator_id"); ok {
		application.AccessValidatorId = Int(d.Get("access_validator_id").(int))
	}

	if _, ok := d.GetOkExists("case_sensitive_path"); ok {
		application.CaseSensitivePath = Bool(d.Get("case_sensitive_path").(bool))
	}

	if _, ok := d.GetOkExists("description"); ok {
		application.Description = String(d.Get("description").(string))
	}

	if _, ok := d.GetOkExists("enabled"); ok {
		application.Enabled = Bool(d.Get("enabled").(bool))
	}

	if _, ok := d.GetOkExists("realm"); ok {
		application.Realm = String(d.Get("realm").(string))
	}

	if _, ok := d.GetOkExists("require_https"); ok {
		application.RequireHTTPS = Bool(d.Get("require_https").(bool))
	}

	if _, ok := d.GetOkExists("web_session_id"); ok {
		webID, _ := strconv.Atoi(d.Get("web_session_id").(string))
		application.WebSessionId = Int(webID)
	}

	if val, ok := d.GetOkExists("identity_mapping_ids"); ok {
		if len(val.([]interface{})) > 0 {
			application.IdentityMappingIds = make(map[string]*int)
			idMapping := val.([]interface{})[0].(map[string]interface{})
			if idMapping["web"] != nil {
				id, _ := strconv.Atoi(idMapping["web"].(string))
				application.IdentityMappingIds["Web"] = Int(id)
			}
			if idMapping["api"] != nil {
				id, _ := strconv.Atoi(idMapping["api"].(string))
				application.IdentityMappingIds["API"] = Int(id)
			}
		}
	}

	if _, ok := d.GetOk(policy); ok {
		policySet := d.Get(policy).([]interface{})

		webPolicies := make([]*pa.PolicyItem, 0)
		apiPolicies := make([]*pa.PolicyItem, 0)

		policy := policySet[0].(map[string]interface{})
		for _, pV := range policy["web"].(*schema.Set).List() {
			p := pV.(map[string]interface{})
			webPolicies = append(webPolicies, &pa.PolicyItem{
				Id:   json.Number(p["id"].(string)),
				Type: String(p["type"].(string)),
			})
		}
		for _, pV := range policy["api"].(*schema.Set).List() {
			p := pV.(map[string]interface{})
			apiPolicies = append(apiPolicies, &pa.PolicyItem{
				Id:   json.Number(p["id"].(string)),
				Type: String(p["type"].(string)),
			})
		}
		policies := map[string]*[]*pa.PolicyItem{
			"Web": &webPolicies,
			"API": &apiPolicies,
		}
		application.Policy = policies
	}

	return application
}
