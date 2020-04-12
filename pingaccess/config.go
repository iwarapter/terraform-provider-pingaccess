package pingaccess

import (
	"crypto/tls"
	"net/http"
	"net/url"
	"os"

	"github.com/iwarapter/pingaccess-sdk-go/pingaccess"
)

const (
	accessValidatorID         = "access_validator_id"
	agentID                   = "agent_id"
	agentResourceCacheTtl     = "agent_resource_cache_ttl"
	anonymous                 = "anonymous"
	applicationID             = "application_id"
	applicationType           = "application_type"
	auditLevel                = "audit_level"
	availabilityProfileID     = "availability_profile_id"
	caseSensitivePath         = "case_sensitive_path"
	className                 = "class_name"
	configuration             = "configuration"
	contextRoot               = "context_root"
	defaultAuthType           = "default_auth_type"
	defaultAuthTypeOverride   = "default_auth_type_override"
	description               = "description"
	destination               = "destination"
	enabled                   = "enabled"
	expectedHostname          = "expected_hostname"
	host                      = "host"
	id                        = "id"
	identityMappingIds        = "identity_mapping_ids"
	keepAliveTimeout          = "keep_alive_timeout"
	keyPairID                 = "key_pair_id"
	loadBalancingStrategyID   = "load_balancing_strategy_id"
	maxConnections            = "max_connections"
	maxWebSocketConnections   = "max_web_socket_connections"
	methods                   = "methods"
	name                      = "name"
	pathPrefixes              = "path_prefixes"
	pathPatterns              = "path_patterns"
	policy                    = "policy"
	port                      = "port"
	realm                     = "realm"
	requireHTTPS              = "require_https"
	rootResource              = "root_resource"
	secure                    = "secure"
	sendPaCookie              = "send_pa_cookie"
	siteAuthenticatorIDs      = "site_authenticator_ids"
	siteID                    = "site_id"
	skipHostnameVerification  = "skip_hostname_verification"
	supportedDestinations     = "supported_destinations"
	targets                   = "targets"
	trustedCertificateGroupID = "trusted_certificate_group_id"
	useProxy                  = "use_proxy"
	useTargetHostHeader       = "use_target_host_header"
	virtualHostIDs            = "virtual_host_ids"
	webSessionId              = "web_session_id"
)

type Config struct {
	Username string
	Password string
	Context  string
	BaseURL  string
}

// Client configures and returns a fully initialized PAClient
func (c *Config) Client() (interface{}, error) {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	url, _ := url.Parse(c.BaseURL)
	client := pingaccess.NewClient(c.Username, c.Password, url, c.Context, nil)

	if os.Getenv("TF_LOG") == "DEBUG" || os.Getenv("TF_LOG") == "TRACE" {
		client.LogDebug(true)
	}
	return client, nil
}
