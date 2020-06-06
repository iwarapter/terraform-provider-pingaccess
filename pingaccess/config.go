package pingaccess

import (
	"crypto/tls"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"net/http"
	"net/url"
	"os"

	"github.com/iwarapter/pingaccess-sdk-go/pingaccess"
)

const (
	anonymous                 = "anonymous"
	applicationID             = "application_id"
	auditLevel                = "audit_level"
	availabilityProfileID     = "availability_profile_id"
	configuration             = "configuration"
	defaultAuthTypeOverride   = "default_auth_type_override"
	enabled                   = "enabled"
	expectedHostname          = "expected_hostname"
	host                      = "host"
	id                        = "id"
	keepAliveTimeout          = "keep_alive_timeout"
	loadBalancingStrategyID   = "load_balancing_strategy_id"
	maxConnections            = "max_connections"
	maxWebSocketConnections   = "max_web_socket_connections"
	methods                   = "methods"
	name                      = "name"
	pathPrefixes              = "path_prefixes"
	pathPatterns              = "path_patterns"
	rootResource              = "root_resource"
	secure                    = "secure"
	sendPaCookie              = "send_pa_cookie"
	siteAuthenticatorIDs      = "site_authenticator_ids"
	skipHostnameVerification  = "skip_hostname_verification"
	targets                   = "targets"
	trustedCertificateGroupID = "trusted_certificate_group_id"
)

type Config struct {
	Username string
	Password string
	Context  string
	BaseURL  string
}

// Client configures and returns a fully initialized PAClient
func (c *Config) Client() (interface{}, diag.Diagnostics) {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	url, _ := url.Parse(c.BaseURL)
	client := pingaccess.NewClient(c.Username, c.Password, url, c.Context, nil)

	if os.Getenv("TF_LOG") == "DEBUG" || os.Getenv("TF_LOG") == "TRACE" {
		client.LogDebug(true)
	}
	return client, nil
}
