package pingaccess

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pa "github.com/iwarapter/pingaccess-sdk-go/pingaccess"
)

func resourcePingAccessApplication() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePingAccessApplicationCreate,
		ReadContext:   resourcePingAccessApplicationRead,
		UpdateContext: resourcePingAccessApplicationUpdate,
		DeleteContext: resourcePingAccessApplicationDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: resourcePingAccessApplicationSchema(),
	}
}

func resourcePingAccessApplicationSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"access_validator_id": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"agent_id": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"application_type": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
		"case_sensitive_path": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"context_root": {
			Type:     schema.TypeString,
			Required: true,
		},
		"default_auth_type": {
			Type:     schema.TypeString,
			Required: true,
		},
		"description": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"destination": {
			Type:     schema.TypeString,
			Required: true,
		},
		"enabled": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"identity_mapping_ids": {
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
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"policy": applicationPolicySchema(),
		"realm": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"require_https": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"site_id": {
			Type:     schema.TypeString,
			Required: true,
		},
		"spa_support_enabled": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"virtual_host_ids": {
			Type:     schema.TypeSet,
			Required: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"web_session_id": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "0",
		},
	}
}

func resourcePingAccessApplicationCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(*pa.Client).Applications
	input := &pa.AddApplicationCommandInput{
		Body: *resourcePingAccessApplicationReadData(d),
	}
	result, _, err := svc.AddApplicationCommand(input)
	if err != nil {
		return diag.Diagnostics{diag.FromErr(fmt.Errorf("unable to create Application: %s", err))}
	}

	d.SetId(result.Id.String())
	return resourcePingAccessApplicationReadResult(d, &input.Body)
}

func resourcePingAccessApplicationRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(*pa.Client).Applications
	input := &pa.GetApplicationCommandInput{
		Id: d.Id(),
	}
	result, _, err := svc.GetApplicationCommand(input)
	if err != nil {
		return diag.Diagnostics{diag.FromErr(fmt.Errorf("unable to read Application: %s", err))}
	}
	return resourcePingAccessApplicationReadResult(d, result)
}

func resourcePingAccessApplicationUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(*pa.Client).Applications
	input := pa.UpdateApplicationCommandInput{
		Body: *resourcePingAccessApplicationReadData(d),
		Id:   d.Id(),
	}
	result, _, err := svc.UpdateApplicationCommand(&input)
	if err != nil {
		return diag.Diagnostics{diag.FromErr(fmt.Errorf("unable to update Application: %s", err))}
	}
	return resourcePingAccessApplicationReadResult(d, result)
}

func resourcePingAccessApplicationDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[INFO] resourcePingAccessApplicationDelete")
	svc := m.(*pa.Client).Applications

	input := &pa.DeleteApplicationCommandInput{
		Id: d.Id(),
	}

	log.Printf("[INFO] ResourceID: %s", d.Id())
	log.Printf("[INFO] DeleteApplicationCommandInput: %s", input.Id)
	_, _, err := svc.DeleteApplicationCommand(input)
	if err != nil {
		return diag.Diagnostics{diag.FromErr(fmt.Errorf("unable to delete Application: %s", err))}
	}
	log.Println("[DEBUG] End - resourcePingAccessSiteDelete")
	return nil
}

func resourcePingAccessApplicationReadResult(d *schema.ResourceData, rv *pa.ApplicationView) diag.Diagnostics {
	var diags diag.Diagnostics
	setResourceDataIntWithDiagnostic(d, "access_validator_id", rv.AccessValidatorId, &diags)
	setResourceDataIntWithDiagnostic(d, "agent_id", rv.AgentId, &diags)
	setResourceDataStringWithDiagnostic(d, "application_type", rv.ApplicationType, &diags)
	setResourceDataBoolWithDiagnostic(d, "case_sensitive_path", rv.CaseSensitivePath, &diags)
	setResourceDataStringWithDiagnostic(d, "context_root", rv.ContextRoot, &diags)
	setResourceDataStringWithDiagnostic(d, "default_auth_type", rv.DefaultAuthType, &diags)
	setResourceDataStringWithDiagnostic(d, "description", rv.Description, &diags)
	setResourceDataStringWithDiagnostic(d, "destination", rv.Destination, &diags)
	setResourceDataBoolWithDiagnostic(d, "enabled", rv.Enabled, &diags)
	setResourceDataStringWithDiagnostic(d, "name", rv.Name, &diags)
	setResourceDataStringWithDiagnostic(d, "realm", rv.Realm, &diags)
	setResourceDataBoolWithDiagnostic(d, "require_https", rv.RequireHTTPS, &diags)
	siteID := strconv.Itoa(*rv.SiteId)
	setResourceDataStringWithDiagnostic(d, "site_id", &siteID, &diags)
	setResourceDataBoolWithDiagnostic(d, "spa_support_enabled", rv.SpaSupportEnabled, &diags)

	if rv.VirtualHostIds != nil {
		var vhs []string
		for _, vh := range *rv.VirtualHostIds {
			vhs = append(vhs, strconv.Itoa(*vh))
		}
		if err := d.Set("virtual_host_ids", vhs); err != nil {
			diags = append(diags, diag.FromErr(err))
		}
	}

	if rv.WebSessionId != nil {
		setResourceDataStringWithDiagnostic(d, "web_session_id", String(strconv.Itoa(*rv.WebSessionId)), &diags)
	}

	if rv.IdentityMappingIds != nil && (*rv.IdentityMappingIds["Web"] != 0 || *rv.IdentityMappingIds["API"] != 0) {
		if err := d.Set("identity_mapping_ids", flattenIdentityMappingIds(rv.IdentityMappingIds)); err != nil {
			diags = append(diags, diag.FromErr(err))
		}
	}
	if rv.Policy != nil && (len(*rv.Policy["API"]) > 0 || len(*rv.Policy["Web"]) > 0) {
		if err := d.Set("policy", flattenPolicy(rv.Policy)); err != nil {
			diags = append(diags, diag.FromErr(err))
		}
	}
	return diags
}

func resourcePingAccessApplicationReadData(d *schema.ResourceData) *pa.ApplicationView {
	siteID, _ := strconv.Atoi(d.Get("site_id").(string))
	virtualHostIds := expandStringList(d.Get("virtual_host_ids").(*schema.Set).List())
	vhIds := []*int{}
	for _, i := range virtualHostIds {
		text, _ := strconv.Atoi(*i)
		vhIds = append(vhIds, &text)
	}

	application := &pa.ApplicationView{
		AgentId:           Int(d.Get("agent_id").(int)),
		Name:              String(d.Get("name").(string)),
		ApplicationType:   String(d.Get("application_type").(string)),
		ContextRoot:       String(d.Get("context_root").(string)),
		DefaultAuthType:   String(d.Get("default_auth_type").(string)),
		SiteId:            Int(siteID),
		SpaSupportEnabled: Bool(d.Get("spa_support_enabled").(bool)),
		VirtualHostIds:    &vhIds,
	}

	if _, ok := d.GetOk("access_validator_id"); ok {
		application.AccessValidatorId = Int(d.Get("access_validator_id").(int))
	}

	if _, ok := d.GetOk("case_sensitive_path"); ok {
		application.CaseSensitivePath = Bool(d.Get("case_sensitive_path").(bool))
	}

	if _, ok := d.GetOk("description"); ok {
		application.Description = String(d.Get("description").(string))
	}

	if v, ok := d.GetOk("destination"); ok {
		application.Destination = String(v.(string))
	}

	if _, ok := d.GetOk("enabled"); ok {
		application.Enabled = Bool(d.Get("enabled").(bool))
	}

	if _, ok := d.GetOk("realm"); ok {
		application.Realm = String(d.Get("realm").(string))
	}

	if _, ok := d.GetOk("require_https"); ok {
		application.RequireHTTPS = Bool(d.Get("require_https").(bool))
	}

	if _, ok := d.GetOk("web_session_id"); ok {
		webID, _ := strconv.Atoi(d.Get("web_session_id").(string))
		application.WebSessionId = Int(webID)
	}

	if val, ok := d.GetOk("identity_mapping_ids"); ok {
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

	if val, ok := d.GetOk("policy"); ok {
		application.Policy = expandPolicy(val.([]interface{}))
	}

	return application
}
