package pingaccess

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/iwarapter/pingaccess-sdk-go/pingaccess"
	pa "github.com/iwarapter/pingaccess-sdk-go/pingaccess"
)

func resourcePingAccessApplicationResource() *schema.Resource {
	return &schema.Resource{
		Create: resourcePingAccessApplicationResourceCreate,
		Read:   resourcePingAccessApplicationResourceRead,
		Update: resourcePingAccessApplicationResourceUpdate,
		Delete: resourcePingAccessApplicationResourceDelete,
		Importer: &schema.ResourceImporter{
			State: resourcePingAccessApplicationResourceImport,
		},

		Schema: resourcePingAccessApplicationResourceSchema(),
	}
}

func resourcePingAccessApplicationResourceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		anonymous: &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
		},
		applicationID: &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		auditLevel: &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		defaultAuthTypeOverride: &schema.Schema{
			Type:         schema.TypeString,
			Optional:     true,
			ValidateFunc: validateWebOrAPI,
		},
		enabled: &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
		},
		methods: &schema.Schema{
			Type:     schema.TypeSet,
			Required: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		name: &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		pathPatterns: &schema.Schema{
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
		pathPrefixes: &schema.Schema{
			Type:     schema.TypeSet,
			Required: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"policy": applicationPolicySchema(),
		rootResource: &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
		},
		"unprotected": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
		},
	}
}

func resourcePingAccessApplicationResourceCreate(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pingaccess.Client).Applications
	applicationID := d.Get("application_id").(string)

	if d.Get("root_resource").(bool) {
		//Root Resources are created automatically, if one exists we will import it instead of failing to create.
		input := pingaccess.GetApplicationResourcesCommandInput{
			Id:   applicationID,
			Name: "Root Resource",
		}
		result, _, err := svc.GetApplicationResourcesCommand(&input)
		if err != nil {
			return fmt.Errorf("Error creating application resource: %s", err)
		}
		rv := result.Items[0]
		d.SetId(rv.Id.String())
		d.Set("application_id", strconv.Itoa(*rv.ApplicationId))
		return resourcePingAccessApplicationResourceReadResult(d, rv)
	}

	input := pingaccess.AddApplicationResourceCommandInput{
		Id:   applicationID,
		Body: *resourcePingAccessApplicationResourceReadData(d),
	}

	result, _, err := svc.AddApplicationResourceCommand(&input)
	if err != nil {
		return fmt.Errorf("Error creating application resource: %s", err)
	}

	d.SetId(result.Id.String())
	return resourcePingAccessApplicationResourceReadResult(d, result)
}

func resourcePingAccessApplicationResourceRead(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pingaccess.Client).Applications
	input := &pingaccess.GetApplicationResourceCommandInput{
		ApplicationId: d.Get("application_id").(string),
		ResourceId:    d.Id(),
	}

	result, _, err := svc.GetApplicationResourceCommand(input)
	if err != nil {
		return fmt.Errorf("Error reading application resource: %s", err)
	}

	return resourcePingAccessApplicationResourceReadResult(d, result)
}

func resourcePingAccessApplicationResourceUpdate(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pingaccess.Client).Applications
	input := pingaccess.UpdateApplicationResourceCommandInput{
		ApplicationId: d.Get("application_id").(string),
		ResourceId:    d.Id(),
		Body:          *resourcePingAccessApplicationResourceReadData(d),
	}

	result, _, err := svc.UpdateApplicationResourceCommand(&input)
	if err != nil {
		return fmt.Errorf("Error updating application: %s", err)
	}
	return resourcePingAccessApplicationResourceReadResult(d, result)
}

func resourcePingAccessApplicationResourceDelete(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pingaccess.Client).Applications

	if d.Get(rootResource).(bool) {
		return nil
	}

	input := &pingaccess.DeleteApplicationResourceCommandInput{
		ResourceId:    d.Id(),
		ApplicationId: d.Get(applicationID).(string),
	}

	_, err := svc.DeleteApplicationResourceCommand(input)
	if err != nil {
		return fmt.Errorf("Error deleting application resource: %s", err)
	}
	return nil
}

func resourcePingAccessApplicationResourceImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	idParts := strings.SplitN(d.Id(), "/", 2)
	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		return nil, fmt.Errorf("unexpected format of ID (%q), expected <application_id>/<resource_id>", d.Id())
	}
	d.Set("application_id", idParts[0])
	d.SetId(idParts[1])
	return []*schema.ResourceData{d}, nil
}

func resourcePingAccessApplicationResourceReadResult(d *schema.ResourceData, rv *pingaccess.ResourceView) error {
	setResourceDataBool(d, anonymous, rv.Anonymous)
	// if err := d.Set(anonymous, *rv.Anonymous); err != nil {
	// 	return err
	// }
	// anonymous := d.Get("anonymous").(bool)
	// setResourceDataString(d, applicationId, *rv.ApplicationId)
	if rv.ApplicationId != nil {
		if err := d.Set(applicationID, strconv.Itoa(*rv.ApplicationId)); err != nil {
			return err
		}
	}
	// return nil
	// if err := d.Set(applicationID, *rv.ApplicationId); err != nil {
	// 	return err
	// }
	// application_id := d.Get("application_id").(string)
	setResourceDataString(d, auditLevel, rv.AuditLevel)
	// if err := d.Set(auditLevel, *rv.AuditLevel); err != nil {
	// 	return err
	// }
	// audit_level := d.Get("audit_level").(string)
	setResourceDataString(d, defaultAuthTypeOverride, rv.DefaultAuthTypeOverride)
	// if err := d.Set(defaultAuthTypeOverride, *rv.DefaultAuthTypeOverride); err != nil {
	// 	return err
	// }
	// default_auth_type_override := d.Get("default_auth_type_override").(string)
	setResourceDataBool(d, enabled, rv.Enabled)
	// if err := d.Set(enabled, *rv.Enabled); err != nil {
	// 	return err
	// }
	// enabled := d.Get("enabled").(bool)
	// methods := d.Get("methods").([]*string)
	if err := d.Set(methods, *rv.Methods); err != nil {
		return err
	}
	// methods := expandStringList(d.Get("methods").(*schema.Set).List())
	if err := d.Set(name, *rv.Name); err != nil {
		return err
	}
	// name := d.Get("name").(string)
	// path_prefixes := d.Get("path_prefixes").([]*string)
	if err := d.Set(pathPrefixes, *rv.PathPrefixes); err != nil {
		return err
	}

	// if err := d.Set("path_patterns", *rv.PathPrefixes); err != nil {
	// 	return err
	// }
	// path_prefixes := expandStringList(d.Get("path_prefixes").(*schema.Set).List())
	// if err := d.Set("policy", rv.Policy); err != nil {
	// 	return err
	// }
	// policy := d.Get("policy").(map[string]interface{})
	setResourceDataBool(d, "root_resource", rv.RootResource)
	setResourceDataBool(d, "unprotected", rv.Unprotected)

	if rv.Policy != nil && (len(*rv.Policy["API"]) > 0 || len(*rv.Policy["Web"]) > 0) {
		if err := d.Set("policy", flattenPolicy(rv.Policy)); err != nil {
			return err
		}
	}
	return nil
}

func resourcePingAccessApplicationResourceReadData(d *schema.ResourceData) *pa.ResourceView {
	methods := expandStringList(d.Get(methods).(*schema.Set).List())

	resource := &pa.ResourceView{
		Name:    String(d.Get("name").(string)),
		Methods: &methods,
	}

	if _, ok := d.GetOkExists("anonymous"); ok {
		resource.Anonymous = Bool(d.Get("anonymous").(bool))
	}

	if _, ok := d.GetOkExists("application_id"); ok {
		applicationID, _ := strconv.Atoi(d.Get("application_id").(string))
		resource.ApplicationId = Int(applicationID)
	}

	if _, ok := d.GetOkExists("audit_level"); ok {
		resource.AuditLevel = String(d.Get("audit_level").(string))
	}

	if _, ok := d.GetOkExists("default_auth_type_override"); ok {
		resource.DefaultAuthTypeOverride = String(d.Get("default_auth_type_override").(string))
	}

	if _, ok := d.GetOkExists("enabled"); ok {
		resource.Enabled = Bool(d.Get("enabled").(bool))
	}

	if _, ok := d.GetOkExists("root_resource"); ok {
		resource.RootResource = Bool(d.Get("root_resource").(bool))
	}

	if _, ok := d.GetOkExists("unprotected"); ok {
		resource.Unprotected = Bool(d.Get("unprotected").(bool))
	}
	if _, ok := d.GetOkExists("path_prefixes"); ok {
		pathPrefixes := expandStringList(d.Get("path_prefixes").(*schema.Set).List())
		resource.PathPrefixes = &pathPrefixes
	}

	if _, ok := d.GetOk("path_patterns"); ok {
		pathPatterns := d.Get("path_patterns").(*schema.Set).List()
		for _, raw := range pathPatterns {
			l := raw.(map[string]interface{})
			p := &pingaccess.PathPatternView{
				Pattern: String(l["pattern"].(string)),
				Type:    String(l["type"].(string)),
			}
			resource.PathPatterns = append(resource.PathPatterns, p)
		}
	}

	if val, ok := d.GetOkExists("policy"); ok {
		resource.Policy = expandPolicy(val.([]interface{}))
	}

	return resource
}
