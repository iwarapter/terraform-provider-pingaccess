package pingaccess

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/iwarapter/pingaccess-sdk-go/pingaccess/models"
	"github.com/iwarapter/pingaccess-sdk-go/services/applications"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePingAccessApplicationResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePingAccessApplicationResourceCreate,
		ReadContext:   resourcePingAccessApplicationResourceRead,
		UpdateContext: resourcePingAccessApplicationResourceUpdate,
		DeleteContext: resourcePingAccessApplicationResourceDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourcePingAccessApplicationResourceImport,
		},

		Schema: resourcePingAccessApplicationResourceSchema(),
	}
}

func resourcePingAccessApplicationResourceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"anonymous": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"application_id": {
			Type:     schema.TypeString,
			Required: true,
		},
		"audit_level": {
			Type:             schema.TypeString,
			Optional:         true,
			Default:          "ON",
			ValidateDiagFunc: validateAuditLevel,
		},
		"default_auth_type_override": {
			Type:             schema.TypeString,
			Optional:         true,
			ValidateDiagFunc: validateWebOrAPI,
		},
		"enabled": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"methods": {
			Type:     schema.TypeSet,
			Required: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"path_patterns": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"pattern": {
						Type:     schema.TypeString,
						Required: true,
					},
					"type": {
						Type:     schema.TypeString,
						Required: true,
					},
				},
			},
		},
		"path_prefixes": {
			Type:       schema.TypeSet,
			Optional:   true,
			Deprecated: "To be removed in a future release; please use 'path_patterns' instead",
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"policy": applicationPolicySchema(),
		"root_resource": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"unprotected": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
	}
}

func resourcePingAccessApplicationResourceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	svc := m.(paClient).Applications
	applicationID := d.Get("application_id").(string)

	if d.Get("root_resource").(bool) {
		//Root Resources are created automatically, if one exists we will import it instead of failing to create.
		input := applications.GetApplicationResourcesCommandInput{
			Id:   applicationID,
			Name: "Root Resource",
		}
		result, _, err := svc.GetApplicationResourcesCommand(&input)
		if err != nil {
			return diag.Errorf("unable to create ApplicationResource: %s", err)
		}
		rv := result.Items[0]
		d.SetId(rv.Id.String())
		setResourceDataStringWithDiagnostic(d, "application_id", String(strconv.Itoa(*rv.ApplicationId)), &diags)
		return resourcePingAccessApplicationResourceUpdate(ctx, d, m)
	}

	input := applications.AddApplicationResourceCommandInput{
		Id:   applicationID,
		Body: *resourcePingAccessApplicationResourceReadData(d),
	}

	result, _, err := svc.AddApplicationResourceCommand(&input)
	if err != nil {
		return diag.Errorf("unable to create ApplicationResource: %s", err)
	}

	d.SetId(result.Id.String())
	return resourcePingAccessApplicationResourceReadResult(d, result)
}

func resourcePingAccessApplicationResourceRead(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(paClient).Applications
	input := &applications.GetApplicationResourceCommandInput{
		ApplicationId: d.Get("application_id").(string),
		ResourceId:    d.Id(),
	}

	result, _, err := svc.GetApplicationResourceCommand(input)
	if err != nil {
		return diag.Errorf("unable to read ApplicationResource: %s", err)
	}

	return resourcePingAccessApplicationResourceReadResult(d, result)
}

func resourcePingAccessApplicationResourceUpdate(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(paClient).Applications
	input := applications.UpdateApplicationResourceCommandInput{
		ApplicationId: d.Get("application_id").(string),
		ResourceId:    d.Id(),
		Body:          *resourcePingAccessApplicationResourceReadData(d),
	}

	result, _, err := svc.UpdateApplicationResourceCommand(&input)
	if err != nil {
		return diag.Errorf("unable to update ApplicationResource: %s", err)
	}
	return resourcePingAccessApplicationResourceReadResult(d, result)
}

func resourcePingAccessApplicationResourceDelete(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(paClient).Applications

	if d.Get("root_resource").(bool) {
		return nil
	}

	input := &applications.DeleteApplicationResourceCommandInput{
		ResourceId:    d.Id(),
		ApplicationId: d.Get("application_id").(string),
	}

	_, err := svc.DeleteApplicationResourceCommand(input)
	if err != nil {
		return diag.Errorf("unable to delete ApplicationResource: %s", err)
	}
	return nil
}

func resourcePingAccessApplicationResourceImport(_ context.Context, d *schema.ResourceData, _ interface{}) ([]*schema.ResourceData, error) {
	idParts := strings.SplitN(d.Id(), "/", 2)
	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		return nil, fmt.Errorf("unexpected format of ID (%q), expected <application_id>/<resource_id>", d.Id())
	}
	_ = d.Set("application_id", idParts[0])
	d.SetId(idParts[1])
	return []*schema.ResourceData{d}, nil
}

func resourcePingAccessApplicationResourceReadResult(d *schema.ResourceData, rv *models.ResourceView) diag.Diagnostics {
	var diags diag.Diagnostics
	setResourceDataBoolWithDiagnostic(d, "anonymous", rv.Anonymous, &diags)
	setResourceDataStringWithDiagnostic(d, "application_id", String(strconv.Itoa(*rv.ApplicationId)), &diags)
	setResourceDataStringWithDiagnostic(d, "audit_level", rv.AuditLevel, &diags)
	setResourceDataStringWithDiagnostic(d, "default_auth_type_override", rv.DefaultAuthTypeOverride, &diags)
	setResourceDataBoolWithDiagnostic(d, "enabled", rv.Enabled, &diags)
	if err := d.Set("methods", *rv.Methods); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	if err := d.Set("name", *rv.Name); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	if rv.PathPrefixes != nil {
		if err := d.Set("path_prefixes", *rv.PathPrefixes); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}
	setResourceDataBoolWithDiagnostic(d, "root_resource", rv.RootResource, &diags)
	setResourceDataBoolWithDiagnostic(d, "unprotected", rv.Unprotected, &diags)

	if rv.Policy != nil && (len(*rv.Policy["Web"]) != 0 || len(*rv.Policy["API"]) != 0) {
		if err := d.Set("policy", flattenPolicy(rv.Policy)); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}
	return diags
}

func resourcePingAccessApplicationResourceReadData(d *schema.ResourceData) *models.ResourceView {
	methods := expandStringList(d.Get("methods").(*schema.Set).List())

	resource := &models.ResourceView{
		Name:    String(d.Get("name").(string)),
		Methods: &methods,
	}

	resource.Anonymous = Bool(d.Get("anonymous").(bool))

	if v, ok := d.GetOk("application_id"); ok {
		applicationID, _ := strconv.Atoi(v.(string))
		resource.ApplicationId = Int(applicationID)
	}

	if v, ok := d.GetOk("audit_level"); ok {
		resource.AuditLevel = String(v.(string))
	}

	if v, ok := d.GetOk("default_auth_type_override"); ok {
		resource.DefaultAuthTypeOverride = String(v.(string))
	}

	resource.Enabled = Bool(d.Get("enabled").(bool))
	resource.RootResource = Bool(d.Get("root_resource").(bool))
	resource.Unprotected = Bool(d.Get("unprotected").(bool))

	if v, ok := d.GetOk("path_prefixes"); ok {
		pathPrefixes := expandStringList(v.(*schema.Set).List())
		resource.PathPrefixes = &pathPrefixes
	}

	if v, ok := d.GetOk("path_patterns"); ok {
		pathPatterns := v.(*schema.Set).List()
		for _, raw := range pathPatterns {
			l := raw.(map[string]interface{})
			p := &models.PathPatternView{
				Pattern: String(l["pattern"].(string)),
				Type:    String(l["type"].(string)),
			}
			resource.PathPatterns = append(resource.PathPatterns, p)
		}
	}

	if v, ok := d.GetOk("policy"); ok {
		resource.Policy = expandPolicy(v.([]interface{}))

	}

	return resource
}
