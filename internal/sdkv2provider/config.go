package sdkv2provider

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"syscall"

	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess/models"

	"github.com/iwarapter/pingaccess-sdk-go/v62/services/accessTokenValidators"
	"github.com/iwarapter/pingaccess-sdk-go/v62/services/acme"
	"github.com/iwarapter/pingaccess-sdk-go/v62/services/adminConfig"
	"github.com/iwarapter/pingaccess-sdk-go/v62/services/adminSessionInfo"
	"github.com/iwarapter/pingaccess-sdk-go/v62/services/agents"
	"github.com/iwarapter/pingaccess-sdk-go/v62/services/applications"
	"github.com/iwarapter/pingaccess-sdk-go/v62/services/auth"
	"github.com/iwarapter/pingaccess-sdk-go/v62/services/authTokenManagement"
	"github.com/iwarapter/pingaccess-sdk-go/v62/services/authnReqLists"
	"github.com/iwarapter/pingaccess-sdk-go/v62/services/backup"
	"github.com/iwarapter/pingaccess-sdk-go/v62/services/certificates"
	"github.com/iwarapter/pingaccess-sdk-go/v62/services/config"
	"github.com/iwarapter/pingaccess-sdk-go/v62/services/engineListeners"
	"github.com/iwarapter/pingaccess-sdk-go/v62/services/engines"
	"github.com/iwarapter/pingaccess-sdk-go/v62/services/globalUnprotectedResources"
	"github.com/iwarapter/pingaccess-sdk-go/v62/services/highAvailability"
	"github.com/iwarapter/pingaccess-sdk-go/v62/services/hsmProviders"
	"github.com/iwarapter/pingaccess-sdk-go/v62/services/httpConfig"
	"github.com/iwarapter/pingaccess-sdk-go/v62/services/httpsListeners"
	"github.com/iwarapter/pingaccess-sdk-go/v62/services/identityMappings"
	"github.com/iwarapter/pingaccess-sdk-go/v62/services/keyPairs"
	"github.com/iwarapter/pingaccess-sdk-go/v62/services/license"
	"github.com/iwarapter/pingaccess-sdk-go/v62/services/oauth"
	"github.com/iwarapter/pingaccess-sdk-go/v62/services/oauthKeyManagement"
	"github.com/iwarapter/pingaccess-sdk-go/v62/services/oidc"
	"github.com/iwarapter/pingaccess-sdk-go/v62/services/pingfederate"
	"github.com/iwarapter/pingaccess-sdk-go/v62/services/pingone"
	"github.com/iwarapter/pingaccess-sdk-go/v62/services/proxies"
	"github.com/iwarapter/pingaccess-sdk-go/v62/services/redirects"
	"github.com/iwarapter/pingaccess-sdk-go/v62/services/rejectionHandlers"
	"github.com/iwarapter/pingaccess-sdk-go/v62/services/rules"
	"github.com/iwarapter/pingaccess-sdk-go/v62/services/rulesets"
	"github.com/iwarapter/pingaccess-sdk-go/v62/services/sharedSecrets"
	"github.com/iwarapter/pingaccess-sdk-go/v62/services/siteAuthenticators"
	"github.com/iwarapter/pingaccess-sdk-go/v62/services/sites"
	"github.com/iwarapter/pingaccess-sdk-go/v62/services/thirdPartyServices"
	"github.com/iwarapter/pingaccess-sdk-go/v62/services/tokenProvider"
	"github.com/iwarapter/pingaccess-sdk-go/v62/services/trustedCertificateGroups"
	"github.com/iwarapter/pingaccess-sdk-go/v62/services/unknownResources"
	"github.com/iwarapter/pingaccess-sdk-go/v62/services/users"
	"github.com/iwarapter/pingaccess-sdk-go/v62/services/version"
	"github.com/iwarapter/pingaccess-sdk-go/v62/services/virtualhosts"
	"github.com/iwarapter/pingaccess-sdk-go/v62/services/webSessionManagement"
	"github.com/iwarapter/pingaccess-sdk-go/v62/services/webSessions"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"

	paCfg "github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess/config"

	paCfg60 "github.com/iwarapter/pingaccess-sdk-go/v60/pingaccess/config"
	keyPairs60 "github.com/iwarapter/pingaccess-sdk-go/v60/services/keyPairs"
)

type cfg struct {
	Username string
	Password string
	Context  string
	BaseURL  string
}

type paClient struct {
	AccessTokenValidators       accessTokenValidators.AccessTokenValidatorsAPI
	Acme                        acme.AcmeAPI
	AdminConfig                 adminConfig.AdminConfigAPI
	AdminSessionInfo            adminSessionInfo.AdminSessionInfoAPI
	Agents                      agents.AgentsAPI
	Applications                applications.ApplicationsAPI
	Auth                        auth.AuthAPI
	AuthTokenManagement         authTokenManagement.AuthTokenManagementAPI
	AuthnReqLists               authnReqLists.AuthnReqListsAPI
	Backup                      backup.BackupAPI
	Certificates                certificates.CertificatesAPI
	Config                      config.ConfigAPI
	EngineListeners             engineListeners.EngineListenersAPI
	Engines                     engines.EnginesAPI
	GlobalUnprotectedResources  globalUnprotectedResources.GlobalUnprotectedResourcesAPI
	HighAvailability            highAvailability.HighAvailabilityAPI
	HighAvailabilityDescriptors *models.DescriptorsView
	HsmProviders                hsmProviders.HsmProvidersAPI
	HttpConfig                  httpConfig.HttpConfigAPI
	HttpsListeners              httpsListeners.HttpsListenersAPI
	IdentityMappings            identityMappings.IdentityMappingsAPI
	IdentityMappingDescriptors  *models.DescriptorsView
	KeyPairs                    keyPairs.KeyPairsAPI
	KeyPairsV60                 keyPairs60.KeyPairsAPI
	License                     license.LicenseAPI
	Oauth                       oauth.OauthAPI
	OauthKeyManagement          oauthKeyManagement.OauthKeyManagementAPI
	Oidc                        oidc.OidcAPI
	Pingfederate                pingfederate.PingfederateAPI
	Pingone                     pingone.PingoneAPI
	Proxies                     proxies.ProxiesAPI
	Redirects                   redirects.RedirectsAPI
	RejectionHandlers           rejectionHandlers.RejectionHandlersAPI
	Rules                       rules.RulesAPI
	RuleDescriptions            *models.RuleDescriptorsView
	Rulesets                    rulesets.RulesetsAPI
	SharedSecrets               sharedSecrets.SharedSecretsAPI
	SiteAuthenticators          siteAuthenticators.SiteAuthenticatorsAPI
	Sites                       sites.SitesAPI
	ThirdPartyServices          thirdPartyServices.ThirdPartyServicesAPI
	TokenProvider               tokenProvider.TokenProviderAPI
	TrustedCertificateGroups    trustedCertificateGroups.TrustedCertificateGroupsAPI
	UnknownResources            unknownResources.UnknownResourcesAPI
	Users                       users.UsersAPI
	Version                     version.VersionAPI
	Virtualhosts                virtualhosts.VirtualhostsAPI
	WebSessionManagement        webSessionManagement.WebSessionManagementAPI
	WebSessions                 webSessions.WebSessionsAPI

	apiVersion string
}

// Client configures and returns a fully initialized PAClient
func (c *cfg) Client() (interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics
	/* #nosec */
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	u, err := url.ParseRequestURI(c.BaseURL)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Invalid URL",
			Detail:   fmt.Sprintf("Unable to parse base_url for client: %s", err),
		})
		return nil, diags
	}

	cfg := paCfg.NewConfig().WithEndpoint(u.String() + c.Context).WithUsername(c.Username).WithPassword(c.Password)
	cfg60 := paCfg60.NewConfig().WithEndpoint(u.String() + c.Context).WithUsername(c.Username).WithPassword(c.Password)
	if os.Getenv("TF_LOG") == "DEBUG" || os.Getenv("TF_LOG") == "TRACE" || os.Getenv("TF_LOG_PROVIDER") == "DEBUG" || os.Getenv("TF_LOG_PROVIDER") == "TRACE" {
		cfg.WithDebug(true)
		cfg60.WithDebug(true)
	}

	client := paClient{
		AccessTokenValidators:      accessTokenValidators.New(cfg),
		Acme:                       acme.New(cfg),
		AdminConfig:                adminConfig.New(cfg),
		AdminSessionInfo:           adminSessionInfo.New(cfg),
		Agents:                     agents.New(cfg),
		Applications:               applications.New(cfg),
		Auth:                       auth.New(cfg),
		AuthTokenManagement:        authTokenManagement.New(cfg),
		AuthnReqLists:              authnReqLists.New(cfg),
		Backup:                     backup.New(cfg),
		Certificates:               certificates.New(cfg),
		Config:                     config.New(cfg),
		EngineListeners:            engineListeners.New(cfg),
		Engines:                    engines.New(cfg),
		GlobalUnprotectedResources: globalUnprotectedResources.New(cfg),
		HighAvailability:           highAvailability.New(cfg),
		HsmProviders:               hsmProviders.New(cfg),
		HttpConfig:                 httpConfig.New(cfg),
		HttpsListeners:             httpsListeners.New(cfg),
		IdentityMappings:           identityMappings.New(cfg),
		KeyPairs:                   keyPairs.New(cfg),
		KeyPairsV60:                keyPairs60.New(cfg60),
		License:                    license.New(cfg),
		Oauth:                      oauth.New(cfg),
		OauthKeyManagement:         oauthKeyManagement.New(cfg),
		Oidc:                       oidc.New(cfg),
		Pingfederate:               pingfederate.New(cfg),
		Pingone:                    pingone.New(cfg),
		Proxies:                    proxies.New(cfg),
		Redirects:                  redirects.New(cfg),
		RejectionHandlers:          rejectionHandlers.New(cfg),
		Rules:                      rules.New(cfg),
		Rulesets:                   rulesets.New(cfg),
		SharedSecrets:              sharedSecrets.New(cfg),
		SiteAuthenticators:         siteAuthenticators.New(cfg),
		Sites:                      sites.New(cfg),
		ThirdPartyServices:         thirdPartyServices.New(cfg),
		TokenProvider:              tokenProvider.New(cfg),
		TrustedCertificateGroups:   trustedCertificateGroups.New(cfg),
		UnknownResources:           unknownResources.New(cfg),
		Users:                      users.New(cfg),
		Version:                    version.New(cfg),
		Virtualhosts:               virtualhosts.New(cfg),
		WebSessionManagement:       webSessionManagement.New(cfg),
		WebSessions:                webSessions.New(cfg),
	}

	v, _, err := client.Version.VersionCommand()
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Connection Error",
			Detail:   fmt.Sprintf("Unable to connect to PingAccess: %s", checkErr(err)),
		})
		return nil, diags
	}

	client.apiVersion = *v.Version

	client.RuleDescriptions, _, err = client.Rules.GetRuleDescriptorsCommand()
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Connection Error",
			Detail:   fmt.Sprintf("Unable to connect to PingAccess: %s", checkErr(err)),
		})
		return nil, diags
	}
	client.IdentityMappingDescriptors, _, err = client.IdentityMappings.GetIdentityMappingDescriptorsCommand()
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Connection Error",
			Detail:   fmt.Sprintf("Unable to connect to PingAccess: %s", checkErr(err)),
		})
		return nil, diags
	}
	client.HighAvailabilityDescriptors, _, err = client.HighAvailability.GetAvailabilityProfileDescriptorsCommand()
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Connection Error",
			Detail:   fmt.Sprintf("Unable to connect to PingAccess: %s", checkErr(err)),
		})
		return nil, diags
	}

	return client, nil
}

// Checks whether we are running against PingAccess 6.1 or above and can track password changes
func (c paClient) CanMaskPasswords() bool {
	re := regexp.MustCompile(`^(6\.0)`)
	return c.Is60OrAbove() && !re.MatchString(c.apiVersion)
}

// Checks whether we are running against PingAccess 6.2 or above
func (c paClient) Is62OrAbove() bool {
	re := regexp.MustCompile(`^(6\.[0-1])`)
	return c.Is60OrAbove() && !re.MatchString(c.apiVersion)
}

// Checks whether we are running against PingAccess 6.2 or above
func (c paClient) Is61OrAbove() bool {
	re := regexp.MustCompile(`^(6\.[0])`)
	return c.Is60OrAbove() && !re.MatchString(c.apiVersion)
}

// Checks whether we are running against PingAccess 6.0 or above
func (c paClient) Is60OrAbove() bool {
	re := regexp.MustCompile(`^([6-9]\.[0-9])`)
	return re.MatchString(c.apiVersion)
}

func checkErr(err error) string {
	if netError, ok := err.(net.Error); ok && netError.Timeout() {
		return "Timeout"
	}

	switch t := err.(type) {
	case *net.OpError:
		if t.Op == "dial" {
			return "Unknown host/port"
		} else if t.Op == "read" {
			return "Connection refused"
		}
	case *url.Error:
		return checkErr(t.Err)
	case syscall.Errno:
		if t == syscall.ECONNREFUSED {
			return "Connection refused"
		}
	}
	return err.Error()
}
