package pingaccess

import (
	"context"
	"strconv"

	"github.com/iwarapter/pingaccess-sdk-go/pingaccess/models"
	"github.com/iwarapter/pingaccess-sdk-go/services/applications"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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
			Default:  0,
		},
		"agent_id": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  0,
		},
		"application_type": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
			ValidateDiagFunc: func(value interface{}, path cty.Path) diag.Diagnostics {
				v := value.(string)
				if v != "Web" && v != "API" && v != "Dynamic" {
					return diag.Errorf("must be either 'Web', 'API' or 'Dynamic' not %s", v)
				}
				return nil
			},
		},
		"case_sensitive_path": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"context_root": {
			Type:     schema.TypeString,
			Required: true,
		},
		"default_auth_type": {
			Type:       schema.TypeString,
			Optional:   true,
			Deprecated: "This field is no longer used and should be removed.",
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

func resourcePingAccessApplicationCreate(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(paClient).Applications
	input := &applications.AddApplicationCommandInput{
		Body: *resourcePingAccessApplicationReadData(d),
	}
	result, _, err := svc.AddApplicationCommand(input)
	if err != nil {
		return diag.Errorf("unable to create Application: %s", err)
	}

	d.SetId(result.Id.String())
	return resourcePingAccessApplicationReadResult(d, &input.Body)
}

func resourcePingAccessApplicationRead(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(paClient).Applications
	input := &applications.GetApplicationCommandInput{
		Id: d.Id(),
	}
	result, _, err := svc.GetApplicationCommand(input)
	if err != nil {
		return diag.Errorf("unable to read Application: %s", err)
	}
	return resourcePingAccessApplicationReadResult(d, result)
}

func resourcePingAccessApplicationUpdate(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(paClient).Applications
	input := applications.UpdateApplicationCommandInput{
		Body: *resourcePingAccessApplicationReadData(d),
		Id:   d.Id(),
	}
	result, _, err := svc.UpdateApplicationCommand(&input)
	if err != nil {
		return diag.Errorf("unable to update Application: %s", err)
	}
	return resourcePingAccessApplicationReadResult(d, result)
}

func resourcePingAccessApplicationDelete(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(paClient).Applications

	input := &applications.DeleteApplicationCommandInput{
		Id: d.Id(),
	}

	_, _, err := svc.DeleteApplicationCommand(input)
	if err != nil {
		return diag.Errorf("unable to delete Application: %s", err)
	}
	return nil
}

func resourcePingAccessApplicationReadResult(d *schema.ResourceData, rv *models.ApplicationView) diag.Diagnostics {
	var diags diag.Diagnostics
	setResourceDataIntWithDiagnostic(d, "access_validator_id", rv.AccessValidatorId, &diags)
	setResourceDataIntWithDiagnostic(d, "agent_id", rv.AgentId, &diags)
	setResourceDataStringWithDiagnostic(d, "application_type", rv.ApplicationType, &diags)
	setResourceDataBoolWithDiagnostic(d, "case_sensitive_path", rv.CaseSensitivePath, &diags)
	setResourceDataStringWithDiagnostic(d, "context_root", rv.ContextRoot, &diags)
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
			diags = append(diags, diag.FromErr(err)...)
		}
	}

	if rv.WebSessionId != nil {
		setResourceDataStringWithDiagnostic(d, "web_session_id", String(strconv.Itoa(*rv.WebSessionId)), &diags)
	}

	if rv.IdentityMappingIds != nil && (*rv.IdentityMappingIds["Web"] != 0 || *rv.IdentityMappingIds["API"] != 0) {
		if err := d.Set("identity_mapping_ids", flattenIdentityMappingIds(rv.IdentityMappingIds)); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}
	if rv.Policy != nil {
		if len(*rv.Policy["API"]) > 0 || len(*rv.Policy["Web"]) > 0 || policyStateHasData(d) {
			if err := d.Set("policy", flattenPolicy(rv.Policy)); err != nil {
				diags = append(diags, diag.FromErr(err)...)
			}
		}
	}
	return diags
}

//https://github.com/hashicorp/terraform-plugin-sdk/issues/142
//because we cannot set the default for the policies, there was an original check to not write policy state
//if the default response was returned from the api (web and api with empty arrays), however this left an edge case
//where config and state could have rules but a manual removal of all rules would not be saved. This helper method checks
//to see if the current config/state has values that should be zero'd out.
func policyStateHasData(d *schema.ResourceData) bool {
	if v, ok := d.GetOk("policy"); ok {
		pol := expandPolicy(v.([]interface{}))
		for _, items := range pol {
			if items != nil && len(*items) > 0 {
				return true
			}
		}
	}
	return false
}

func resourcePingAccessApplicationReadData(d *schema.ResourceData) *models.ApplicationView {
	siteID, _ := strconv.Atoi(d.Get("site_id").(string))
	virtualHostIds := expandStringList(d.Get("virtual_host_ids").(*schema.Set).List())
	var vhIds []*int
	for _, i := range virtualHostIds {
		text, _ := strconv.Atoi(*i)
		vhIds = append(vhIds, &text)
	}

	application := &models.ApplicationView{
		AgentId:           Int(d.Get("agent_id").(int)),
		Name:              String(d.Get("name").(string)),
		ApplicationType:   String(d.Get("application_type").(string)),
		ContextRoot:       String(d.Get("context_root").(string)),
		SiteId:            Int(siteID),
		SpaSupportEnabled: Bool(d.Get("spa_support_enabled").(bool)),
		VirtualHostIds:    &vhIds,
		CaseSensitivePath: Bool(d.Get("case_sensitive_path").(bool)),
		AccessValidatorId: Int(d.Get("access_validator_id").(int)),
	}

	if *application.ApplicationType != "Dynamic" {
		application.DefaultAuthType = application.ApplicationType
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
