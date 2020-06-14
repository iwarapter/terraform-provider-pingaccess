package pingaccess

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

//Provider does stuff
//
func Provider() *schema.Provider {
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
		DataSourcesMap: map[string]*schema.Resource{
			"pingaccess_acme_default":                  dataSourcePingAccessAcmeDefault(),
			"pingaccess_certificate":                   dataSourcePingAccessCertificate(),
			"pingaccess_keypair":                       dataSourcePingAccessKeyPair(),
			"pingaccess_pingfederate_runtime_metadata": dataSourcePingAccessPingFederateRuntimeMetadata(),
			"pingaccess_trusted_certificate_group":     dataSourcePingAccessTrustedCertificateGroups(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"pingaccess_access_token_validator":          resourcePingAccessAccessTokenValidator(),
			"pingaccess_acme_server":                     resourcePingAccessAcmeServer(),
			"pingaccess_auth_token_management":           resourcePingAccessAuthTokenManagement(),
			"pingaccess_authn_req_list":                  resourcePingAccessAuthnReqList(),
			"pingaccess_certificate":                     resourcePingAccessCertificate(),
			"pingaccess_engine_listener":                 resourcePingAccessEngineListener(),
			"pingaccess_hsm_provider":                    resourcePingAccessHsmProvider(),
			"pingaccess_https_listener":                  resourcePingAccessHTTPSListener(),
			"pingaccess_identity_mapping":                resourcePingAccessIdentityMapping(),
			"pingaccess_keypair":                         resourcePingAccessKeyPair(),
			"pingaccess_rule":                            resourcePingAccessRule(),
			"pingaccess_ruleset":                         resourcePingAccessRuleSet(),
			"pingaccess_virtualhost":                     resourcePingAccessVirtualHost(),
			"pingaccess_site":                            resourcePingAccessSite(),
			"pingaccess_application":                     resourcePingAccessApplication(),
			"pingaccess_application_resource":            resourcePingAccessApplicationResource(),
			"pingaccess_websession":                      resourcePingAccessWebSession(),
			"pingaccess_site_authenticator":              resourcePingAccessSiteAuthenticator(),
			"pingaccess_third_party_service":             resourcePingAccessThirdPartyService(),
			"pingaccess_trusted_certificate_group":       resourcePingAccessTrustedCertificateGroups(),
			"pingaccess_pingfederate_admin":              resourcePingAccessPingFederateAdmin(),
			"pingaccess_pingfederate_runtime":            resourcePingAccessPingFederateRuntime(),
			"pingaccess_pingfederate_oauth":              resourcePingAccessPingFederateOAuth(),
			"pingaccess_oauth_server":                    resourcePingAccessOAuthServer(),
			"pingaccess_http_config_request_host_source": resourcePingAccessHTTPConfigRequestHostSource(),
		},
		ConfigureContextFunc: providerConfigure,
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

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {

	config := &config{
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

// String is a helper routine that allocates a new string value
// to store v and returns a pointer to it.
func String(v string) *string { return &v }

func setResourceDataStringWithDiagnostic(d *schema.ResourceData, name string, data *string, diags *diag.Diagnostics) {
	if data != nil {
		if err := d.Set(name, *data); err != nil {
			*diags = append(*diags, diag.FromErr(err)...)
		}
	}
}

func setResourceDataIntWithDiagnostic(d *schema.ResourceData, name string, data *int, diags *diag.Diagnostics) {
	if data != nil {
		if err := d.Set(name, *data); err != nil {
			*diags = append(*diags, diag.FromErr(err)...)
		}
	}
}

func setResourceDataBoolWithDiagnostic(d *schema.ResourceData, name string, data *bool, diags *diag.Diagnostics) {
	if data != nil {
		if err := d.Set(name, *data); err != nil {
			*diags = append(*diags, diag.FromErr(err)...)
		}
	}
}
