package pingaccess

import (
	"crypto/tls"
	"net/http"
	"net/url"
	"os"
	"regexp"

	"github.com/iwarapter/pingaccess-sdk-go/services/accessTokenValidators"
	"github.com/iwarapter/pingaccess-sdk-go/services/acme"
	"github.com/iwarapter/pingaccess-sdk-go/services/adminConfig"
	"github.com/iwarapter/pingaccess-sdk-go/services/adminSessionInfo"
	"github.com/iwarapter/pingaccess-sdk-go/services/agents"
	"github.com/iwarapter/pingaccess-sdk-go/services/applications"
	"github.com/iwarapter/pingaccess-sdk-go/services/auth"
	"github.com/iwarapter/pingaccess-sdk-go/services/authTokenManagement"
	"github.com/iwarapter/pingaccess-sdk-go/services/authnReqLists"
	"github.com/iwarapter/pingaccess-sdk-go/services/backup"
	"github.com/iwarapter/pingaccess-sdk-go/services/certificates"
	"github.com/iwarapter/pingaccess-sdk-go/services/config"
	"github.com/iwarapter/pingaccess-sdk-go/services/engineListeners"
	"github.com/iwarapter/pingaccess-sdk-go/services/engines"
	"github.com/iwarapter/pingaccess-sdk-go/services/globalUnprotectedResources"
	"github.com/iwarapter/pingaccess-sdk-go/services/highAvailability"
	"github.com/iwarapter/pingaccess-sdk-go/services/hsmProviders"
	"github.com/iwarapter/pingaccess-sdk-go/services/httpConfig"
	"github.com/iwarapter/pingaccess-sdk-go/services/httpsListeners"
	"github.com/iwarapter/pingaccess-sdk-go/services/identityMappings"
	"github.com/iwarapter/pingaccess-sdk-go/services/keyPairs"
	"github.com/iwarapter/pingaccess-sdk-go/services/license"
	"github.com/iwarapter/pingaccess-sdk-go/services/oauth"
	"github.com/iwarapter/pingaccess-sdk-go/services/oauthKeyManagement"
	"github.com/iwarapter/pingaccess-sdk-go/services/oidc"
	"github.com/iwarapter/pingaccess-sdk-go/services/pingfederate"
	"github.com/iwarapter/pingaccess-sdk-go/services/pingone"
	"github.com/iwarapter/pingaccess-sdk-go/services/proxies"
	"github.com/iwarapter/pingaccess-sdk-go/services/redirects"
	"github.com/iwarapter/pingaccess-sdk-go/services/rejectionHandlers"
	"github.com/iwarapter/pingaccess-sdk-go/services/rules"
	"github.com/iwarapter/pingaccess-sdk-go/services/rulesets"
	"github.com/iwarapter/pingaccess-sdk-go/services/sharedSecrets"
	"github.com/iwarapter/pingaccess-sdk-go/services/siteAuthenticators"
	"github.com/iwarapter/pingaccess-sdk-go/services/sites"
	"github.com/iwarapter/pingaccess-sdk-go/services/thirdPartyServices"
	"github.com/iwarapter/pingaccess-sdk-go/services/tokenProvider"
	"github.com/iwarapter/pingaccess-sdk-go/services/trustedCertificateGroups"
	"github.com/iwarapter/pingaccess-sdk-go/services/unknownResources"
	"github.com/iwarapter/pingaccess-sdk-go/services/users"
	"github.com/iwarapter/pingaccess-sdk-go/services/version"
	"github.com/iwarapter/pingaccess-sdk-go/services/virtualhosts"
	"github.com/iwarapter/pingaccess-sdk-go/services/webSessionManagement"
	"github.com/iwarapter/pingaccess-sdk-go/services/webSessions"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"

	paCfg "github.com/iwarapter/pingaccess-sdk-go/pingaccess/config"
)

type cfg struct {
	Username string
	Password string
	Context  string
	BaseURL  string
}

type paClient struct {
	AccessTokenValidators      accessTokenValidators.AccessTokenValidatorsAPI
	Acme                       acme.AcmeAPI
	AdminConfig                adminConfig.AdminConfigAPI
	AdminSessionInfo           adminSessionInfo.AdminSessionInfoAPI
	Agents                     agents.AgentsAPI
	Applications               applications.ApplicationsAPI
	Auth                       auth.AuthAPI
	AuthTokenManagement        authTokenManagement.AuthTokenManagementAPI
	AuthnReqLists              authnReqLists.AuthnReqListsAPI
	Backup                     backup.BackupAPI
	Certificates               certificates.CertificatesAPI
	Config                     config.ConfigAPI
	EngineListeners            engineListeners.EngineListenersAPI
	Engines                    engines.EnginesAPI
	GlobalUnprotectedResources globalUnprotectedResources.GlobalUnprotectedResourcesAPI
	HighAvailability           highAvailability.HighAvailabilityAPI
	HsmProviders               hsmProviders.HsmProvidersAPI
	HttpConfig                 httpConfig.HttpConfigAPI
	HttpsListeners             httpsListeners.HttpsListenersAPI
	IdentityMappings           identityMappings.IdentityMappingsAPI
	KeyPairs                   keyPairs.KeyPairsAPI
	License                    license.LicenseAPI
	Oauth                      oauth.OauthAPI
	OauthKeyManagement         oauthKeyManagement.OauthKeyManagementAPI
	Oidc                       oidc.OidcAPI
	Pingfederate               pingfederate.PingfederateAPI
	Pingone                    pingone.PingoneAPI
	Proxies                    proxies.ProxiesAPI
	Redirects                  redirects.RedirectsAPI
	RejectionHandlers          rejectionHandlers.RejectionHandlersAPI
	Rules                      rules.RulesAPI
	Rulesets                   rulesets.RulesetsAPI
	SharedSecrets              sharedSecrets.SharedSecretsAPI
	SiteAuthenticators         siteAuthenticators.SiteAuthenticatorsAPI
	Sites                      sites.SitesAPI
	ThirdPartyServices         thirdPartyServices.ThirdPartyServicesAPI
	TokenProvider              tokenProvider.TokenProviderAPI
	TrustedCertificateGroups   trustedCertificateGroups.TrustedCertificateGroupsAPI
	UnknownResources           unknownResources.UnknownResourcesAPI
	Users                      users.UsersAPI
	Version                    version.VersionAPI
	Virtualhosts               virtualhosts.VirtualhostsAPI
	WebSessionManagement       webSessionManagement.WebSessionManagementAPI
	WebSessions                webSessions.WebSessionsAPI

	apiVersion string
}

// Client configures and returns a fully initialized PAClient
func (c *cfg) Client() (interface{}, diag.Diagnostics) {
	/* #nosec */
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	u, _ := url.Parse(c.BaseURL)

	cfg := paCfg.NewConfig().WithEndpoint(u.String() + c.Context).WithUsername(c.Username).WithPassword(c.Password)

	if os.Getenv("TF_LOG") == "DEBUG" || os.Getenv("TF_LOG") == "TRACE" {
		cfg.WithDebug(true)
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
		return nil, diag.Errorf("unable to retrieve version %s", err)
	}
	client.apiVersion = *v.Version

	return client, nil
}

// Checks whether we are running against PingAccess 6.1 or above and can track password changes
func (c paClient) CanMaskPasswords() bool {
	re := regexp.MustCompile(`^(6\.[1-9])`)
	return re.MatchString(c.apiVersion)
}
