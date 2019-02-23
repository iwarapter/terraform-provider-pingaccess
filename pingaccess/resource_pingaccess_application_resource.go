package pingaccess

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/iwarapter/pingaccess-sdk-go/pingaccess"
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

		Schema: map[string]*schema.Schema{
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
			rootResource: &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
		},
	}
}

func resourcePingAccessApplicationResourceCreate(d *schema.ResourceData, m interface{}) error {
	log.Printf("[INFO] resourcePingAccessApplicationResourceCreate")
	svc := m.(*pingaccess.Client).Applications
	anonymous := d.Get(anonymous).(bool)
	application_id := d.Get(applicationID).(string)
	audit_level := d.Get(auditLevel).(string)
	enabled := d.Get(enabled).(bool)
	// methods := d.Get("methods").([]*string)
	methods := expandStringList(d.Get(methods).(*schema.Set).List())
	name := d.Get(name).(string)
	// path_prefixes := d.Get("path_prefixes").([]*string)
	path_prefixes := expandStringList(d.Get(pathPrefixes).(*schema.Set).List())
	// policy := d.Get("policy").(map[string]interface{})
	root_resource := d.Get(rootResource).(bool)

	appId, _ := strconv.Atoi(application_id)

	// log.Printf("[INFO] RootResourceCheck: %s %b", name, root_resource)
	if root_resource {
		//Root Resources are created automatically, if one exists we will import it instead of failing to create.
		input := pingaccess.GetApplicationResourcesCommandInput{
			Id:   application_id,
			Name: "Root Resource",
		}
		result, _, err := svc.GetApplicationResourcesCommand(&input)
		b, _ := json.Marshal(result)
		log.Printf("[INFO] RootResourceCheck: %s", b)
		if err != nil {
			return fmt.Errorf("Error creating application resource: %s", err)
		}
		// d.SetId(result.Id.String())
		rv := result.Items[0]
		d.SetId(rv.Id.String())
		d.Set("application_id", strconv.Itoa(*result.Items[0].ApplicationId))
		// log.Printf("[INFO] RootResourceCheck: set application ID: %d and resource ID: %d", rv.Id.String(), *result.Items[0].ApplicationId)
		return resourcePingAccessApplicationResourceReadResult(d, result.Items[0])
	}

	log.Printf("[INFO] Not Root Resource, creating normally")
	input := pingaccess.AddApplicationResourceCommandInput{
		Id: application_id,
		Body: pingaccess.ResourceView{
			Anonymous:     Bool(anonymous),
			ApplicationId: Int(appId),
			AuditLevel:    String(audit_level),
			Enabled:       Bool(enabled),
			Methods:       &methods,
			Name:          String(name),
			PathPrefixes:  &path_prefixes,
			RootResource:  Bool(root_resource),
		},
	}

	if default_auth_type_override, ok := d.GetOk(defaultAuthTypeOverride); ok {
		input.Body.DefaultAuthTypeOverride = String(default_auth_type_override.(string))
	}
	log.Printf("[DEBUG-README] looking for patterns")
	if _, ok := d.GetOk(pathPatterns); ok {
		path_patterns := d.Get(pathPatterns).(*schema.Set).List()
		log.Printf("[DEBUG-README] found patterns")
		for _, raw := range path_patterns {
			l := raw.(map[string]interface{})
			p := &pingaccess.PathPatternView{
				Pattern: String(l["pattern"].(string)),
				Type:    String(l["type"].(string)),
			}
			log.Printf("[DEBUG-README] Adding: %s %s", l["pattern"].(string), l["type"].(string))
			input.Body.PathPatterns = append(input.Body.PathPatterns, p)
		}
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
	// default_auth_type_override := d.GetOk(defaultAuthTypeOverride)

	// default_auth_type_override := d.Get(defaultAuthTypeOverride).(string)

	result, resp, err := svc.AddApplicationResourceCommand(&input)
	log.Printf("[DEBUG-README] AddApplicationResourceCommand-Response: %d", resp.StatusCode)
	if err != nil {
		return fmt.Errorf("Error creating application: %s", err)
	}

	d.SetId(result.Id.String())
	return resourcePingAccessApplicationResourceReadResult(d, result)
}

func resourcePingAccessApplicationResourceRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("[INFO] resourcePingAccessApplicationResourceRead")
	svc := m.(*pingaccess.Client).Applications

	input := &pingaccess.GetApplicationResourceCommandInput{
		ApplicationId: d.Get("application_id").(string),
		ResourceId:    d.Id(),
	}

	log.Printf("[INFO] ResourceID: %s", d.Id())
	log.Printf("[INFO] GetApplicationResourceCommandInput: %s", input.ApplicationId)
	result, _, _ := svc.GetApplicationResourceCommand(input)
	b, _ := json.Marshal(result)
	log.Printf("[INFO] read rv: %s", b)
	// b := new(bytes.Buffer)
	// json.NewEncoder(b).Encode(result)
	// rs := pingaccess.ResourceView{}
	// json.NewDecoder(b).Decode(&rs)

	return resourcePingAccessApplicationResourceReadResult(d, result)
}

func resourcePingAccessApplicationResourceUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("[INFO] resourcePingAccessApplicationResourceUpdate")
	anonymous := d.Get(anonymous).(bool)
	application_id := d.Get(applicationID).(string)
	audit_level := d.Get(auditLevel).(string)
	// default_auth_type_override := d.Get(defaultAuthTypeOverride).(string)
	enabled := d.Get(enabled).(bool)
	// methods := d.Get("methods").([]*string)
	methods := expandStringList(d.Get(methods).(*schema.Set).List())
	name := d.Get(name).(string)
	// path_prefixes := d.Get("path_prefixes").([]*string)
	path_prefixes := expandStringList(d.Get(pathPrefixes).(*schema.Set).List())
	// policy := d.Get("policy").(map[string]interface{})
	root_resource := d.Get(rootResource).(bool)

	appId, _ := strconv.Atoi(application_id)

	input := pingaccess.UpdateApplicationResourceCommandInput{
		ApplicationId: application_id,
		ResourceId:    d.Id(),
		Body: pingaccess.ResourceView{
			Anonymous:     Bool(anonymous),
			ApplicationId: Int(appId),
			AuditLevel:    String(audit_level),
			Enabled:       Bool(enabled),
			Methods:       &methods,
			Name:          String(name),
			PathPrefixes:  &path_prefixes,
			RootResource:  Bool(root_resource),
		},
	}
	input.ApplicationId = d.Get("application_id").(string)

	svc := m.(*pingaccess.Client).Applications

	if default_auth_type_override, ok := d.GetOk(defaultAuthTypeOverride); ok {
		input.Body.DefaultAuthTypeOverride = String(default_auth_type_override.(string))
	}
	log.Printf("[DEBUG-README] looking for patterns")
	if _, ok := d.GetOk(pathPatterns); ok {
		path_patterns := d.Get(pathPatterns).(*schema.Set).List()
		log.Printf("[DEBUG-README] found patterns")
		for _, raw := range path_patterns {
			l := raw.(map[string]interface{})
			p := &pingaccess.PathPatternView{
				Pattern: String(l["pattern"].(string)),
				Type:    String(l["type"].(string)),
			}
			log.Printf("[DEBUG-README] Adding: %s %s", l["pattern"].(string), l["type"].(string))
			input.Body.PathPatterns = append(input.Body.PathPatterns, p)
		}
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

	_, resp, err := svc.UpdateApplicationResourceCommand(&input)
	log.Printf("[DEBUG-README] UpdateApplicationResourceCommand-Response: %d", resp.StatusCode)
	if err != nil {
		return fmt.Errorf("Error updating application: %s", err)
	}
	log.Println("[INFO] End - resourcePingAccessApplicationResourceUpdate")
	return resourcePingAccessApplicationResourceRead(d, m)
}

func resourcePingAccessApplicationResourceDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("[INFO] resourcePingAccessApplicationResourceDelete")
	svc := m.(*pingaccess.Client).Applications

	input := &pingaccess.DeleteApplicationResourceCommandInput{
		ResourceId:    d.Id(),
		ApplicationId: d.Get(applicationID).(string),
	}

	log.Printf("[INFO] ResourceID: %s", d.Id())
	log.Printf("[INFO] DeleteApplicationCommandInput: %s", input.ResourceId)
	_, err := svc.DeleteApplicationResourceCommand(input)
	if err != nil {
		return fmt.Errorf("Error deleting application resource: %s", err)
	}
	log.Println("[INFO] End - resourcePingAccessApplicationResourceDelete")
	return nil
}

func resourcePingAccessApplicationResourceImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	idParts := strings.SplitN(d.Id(), "/", 2)
	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		return nil, fmt.Errorf("unexpected format of ID (%q), expected <application_id>/<resource_id>", d.Id())
	}
	application_id := idParts[0]
	resource_id := idParts[1]
	d.Set("application_id", application_id)
	d.SetId(resource_id)
	return []*schema.ResourceData{d}, nil
}

func resourcePingAccessApplicationResourceReadResult(d *schema.ResourceData, rv *pingaccess.ResourceView) error {
	log.Printf("[INFO] BEGIN - resourcePingAccessApplicationResourceReadResult")
	b, _ := json.Marshal(*rv)
	log.Printf("[INFO] rv: %s", b)
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
	// path_prefixes := expandStringList(d.Get("path_prefixes").(*schema.Set).List())
	// if err := d.Set("policy", rv.Policy); err != nil {
	// 	return err
	// }
	// policy := d.Get("policy").(map[string]interface{})
	if err := d.Set(rootResource, *rv.RootResource); err != nil {
		return err
	}
	log.Printf("[INFO] END - resourcePingAccessApplicationResourceReadResult")
	return nil
}
