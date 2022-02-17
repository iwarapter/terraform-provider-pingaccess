package sdkv2provider

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess/models"
	"github.com/iwarapter/pingaccess-sdk-go/v62/services/applications"

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
		Schema:      resourcePingAccessApplicationResourceSchema(),
		Description: "Provides configuration for Application Resources within PingAccess.",
	}
}

func resourcePingAccessApplicationResourceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"anonymous": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
			Description: "True if the resource is anonymous.",
		},
		"application_id": {
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
			Description: "The id of the associated application. This field is read-only.",
		},
		"audit_level": {
			Type:             schema.TypeString,
			Optional:         true,
			Default:          "ON",
			ValidateDiagFunc: validateAuditLevel,
			Description:      "Indicates if audit logging is enabled for the resource.",
		},
		"default_auth_type_override": {
			Type:             schema.TypeString,
			Optional:         true,
			ValidateDiagFunc: validateWebOrAPI,
			Description:      "For Web + API applications (dynamic) default_auth_type selects the processing mode when a request: does not have a token (web session, OAuth bearer) or has both tokens. default_auth_type_override overrides the default_auth_type at the application level for this resource. A value of null indicates the resource should not override the default_auth_type.",
		},
		"enabled": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     true,
			Description: "True if the resource is enabled.",
		},
		"methods": {
			Type:        schema.TypeSet,
			Required:    true,
			Description: "An array of HTTP methods configured for the resource.",
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The name of the resource.",
		},
		"path_patterns": {
			Type:        schema.TypeSet,
			Optional:    true,
			Description: "A list of one or more request path-matching patterns.",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"pattern": {
						Type:        schema.TypeString,
						Required:    true,
						Description: "The path-matching pattern, relative to the Application context root (interpreted according to the pattern 'type').",
					},
					"type": {
						Type:        schema.TypeString,
						Required:    true,
						Description: "The pattern syntax type.",
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
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
			Description: "True if the resource is the root resource for the application.",
		},
		"unprotected": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
			Description: "True if the resource is unprotected.",
		},
		"resource_type": {
			Type:             schema.TypeString,
			Optional:         true,
			Default:          "Standard",
			ValidateDiagFunc: validateResourceType,
			Description:      "The type of this resource. 'Standard' resources are those served by the protected applications. 'Virtual' resources do not have a corresponding resource in the protected application. Instead, when accessing the resource, PingAccess returns a response created by the response generator defined in the resource type configuration. The default type is 'Standard'.",
		},
		"resource_type_configuration": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "A container for configuration specific to different types of resources.",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"response_generator": {
						Type:        schema.TypeList,
						Required:    true,
						Description: "The path-matching pattern, relative to the Application context root (interpreted according to the pattern 'type').",
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"class_name": {
									Type:        schema.TypeString,
									Required:    true,
									Description: "The response generator's class name.",
								},
								"configuration": {
									Type:             schema.TypeString,
									Optional:         true,
									DiffSuppressFunc: suppressEquivalentJSONDiffs,
									Description:      "The response generator's configuration data.",
								},
							},
						},
					},
				},
			},
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
	if rv.PathPatterns != nil {
		if err := d.Set("path_patterns", flattenPathPatternView(rv.PathPatterns)); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}
	setResourceDataBoolWithDiagnostic(d, "root_resource", rv.RootResource, &diags)
	setResourceDataBoolWithDiagnostic(d, "unprotected", rv.Unprotected, &diags)
	diags = append(diags, flattenPolicies(d, rv.Policy)...)

	setResourceDataStringWithDiagnostic(d, "resource_type", rv.ResourceType, &diags)

	if rv.ResourceTypeConfiguration != nil {
		if err := d.Set("resource_type_configuration", flattenResourceTypeConfiguration(rv.ResourceTypeConfiguration)); err != nil {
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

	if v, ok := d.GetOk("resource_type"); ok {
		resource.ResourceType = String(v.(string))
	}

	if v, ok := d.GetOk("resource_type_configuration"); ok {
		resource.ResourceTypeConfiguration = expandResourceTypeConfiguration(v.([]interface{}))
	}

	return resource
}
