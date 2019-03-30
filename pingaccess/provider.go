package pingaccess

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

//Provider does stuff
//
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"username": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions["username"],
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{"PINGACCESS_USERNAME"}, "Administrator"),
			},
			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions["password"],
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{"PINGACCESS_PASSWORD"}, "2Access"),
			},
			"context": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions["context"],
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{"PINGACCESS_CONTEXT"}, "/pa-admin-api/v3"),
			},
			"base_url": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions["base_url"],
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{"PINGACCESS_BASEURL"}, "https://localhost:9000"),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"pingaccess_certificate":               resourcePingAccessCertificate(),
			"pingaccess_identity_mapping":          resourcePingAccessIdentityMapping(),
			"pingaccess_rule":                      resourcePingAccessRule(),
			"pingaccess_ruleset":                   resourcePingAccessRuleSet(),
			"pingaccess_virtualhost":               resourcePingAccessVirtualHost(),
			"pingaccess_site":                      resourcePingAccessSite(),
			"pingaccess_application":               resourcePingAccessApplication(),
			"pingaccess_application_resource":      resourcePingAccessApplicationResource(),
			"pingaccess_websession":                resourcePingAccessWebSession(),
			"pingaccess_site_authenticator":        resourcePingAccessSiteAuthenticator(),
			"pingaccess_third_party_service":       resourcePingAccessThirdPartyService(),
			"pingaccess_trusted_certificate_group": resourcePingAccessTrustedCertificateGroups(),
			"pingaccess_pingfederate_runtime":      resourcePingAccessPingFederateRuntime(),
			"pingaccess_pingfederate_oauth":        resourcePingAccessPingFederateOAuth(),
			"pingaccess_oauth_server":              resourcePingAccessOAuthServer(),
		},
		ConfigureFunc: providerConfigure,
	}
}

var descriptions map[string]string

func init() {
	descriptions = map[string]string{
		"username": "The username for pingaccess API.",
		"password": "The password for pingaccess API.",
		"base_url": "The base url of the pingaccess API.",
		"context":  "The context path of the pingaccess API.",
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := &Config{
		Username: d.Get("username").(string),
		Password: d.Get("password").(string),
		BaseURL:  d.Get("base_url").(string),
		Context:  d.Get("context").(string),
	}

	return config.Client()
}

// Bool is a helper routine that allocates a new bool value
// to store v and returns a pointer to it.
func Bool(v bool) *bool { return &v }

// Int is a helper routine that allocates a new int value
// to store v and returns a pointer to it.
func Int(v int) *int { return &v }

// Int64 is a helper routine that allocates a new int64 value
// to store v and returns a pointer to it.
func Int64(v int64) *int64 { return &v }

// String is a helper routine that allocates a new string value
// to store v and returns a pointer to it.
func String(v string) *string { return &v }

func setResourceDataString(d *schema.ResourceData, name string, data *string) error {
	if data != nil {
		if err := d.Set(name, *data); err != nil {
			return err
		}
	}
	return nil
}

func setResourceDataInt(d *schema.ResourceData, name string, data *int) error {
	if data != nil {
		if err := d.Set(name, *data); err != nil {
			return err
		}
	}
	return nil
}

func setResourceDataBool(d *schema.ResourceData, name string, data *bool) error {
	if data != nil {
		if err := d.Set(name, *data); err != nil {
			return err
		}
	}
	return nil
}
