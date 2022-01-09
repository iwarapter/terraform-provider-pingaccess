package protocol

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

	"github.com/hashicorp/terraform-plugin-go/tfprotov5"

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

	paCfg "github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess/config"
)

type cfg struct {
	Username string
	Password string
	Context  string
	BaseURL  string
}

type paClient struct {
	AccessTokenValidators            accessTokenValidators.AccessTokenValidatorsAPI
	AccessTokenValidatorsDescriptors *models.DescriptorsView
	Acme                             acme.AcmeAPI
	AdminConfig                      adminConfig.AdminConfigAPI
	AdminSessionInfo                 adminSessionInfo.AdminSessionInfoAPI
	Agents                           agents.AgentsAPI
	Applications                     applications.ApplicationsAPI
	Auth                             auth.AuthAPI
	AuthTokenManagement              authTokenManagement.AuthTokenManagementAPI
	AuthnReqLists                    authnReqLists.AuthnReqListsAPI
	Backup                           backup.BackupAPI
	Certificates                     certificates.CertificatesAPI
	Config                           config.ConfigAPI
	EngineListeners                  engineListeners.EngineListenersAPI
	Engines                          engines.EnginesAPI
	GlobalUnprotectedResources       globalUnprotectedResources.GlobalUnprotectedResourcesAPI
	HighAvailability                 highAvailability.HighAvailabilityAPI
	HsmProviders                     hsmProviders.HsmProvidersAPI
	HttpConfig                       httpConfig.HttpConfigAPI
	HttpsListeners                   httpsListeners.HttpsListenersAPI
	IdentityMappings                 identityMappings.IdentityMappingsAPI
	KeyPairs                         keyPairs.KeyPairsAPI
	License                          license.LicenseAPI
	Oauth                            oauth.OauthAPI
	OauthKeyManagement               oauthKeyManagement.OauthKeyManagementAPI
	Oidc                             oidc.OidcAPI
	Pingfederate                     pingfederate.PingfederateAPI
	Pingone                          pingone.PingoneAPI
	Proxies                          proxies.ProxiesAPI
	Redirects                        redirects.RedirectsAPI
	RejectionHandlers                rejectionHandlers.RejectionHandlersAPI
	Rules                            rules.RulesAPI
	Rulesets                         rulesets.RulesetsAPI
	SharedSecrets                    sharedSecrets.SharedSecretsAPI
	SiteAuthenticators               siteAuthenticators.SiteAuthenticatorsAPI
	SiteAuthenticatorsDescriptors    *models.DescriptorsView
	Sites                            sites.SitesAPI
	ThirdPartyServices               thirdPartyServices.ThirdPartyServicesAPI
	TokenProvider                    tokenProvider.TokenProviderAPI
	TrustedCertificateGroups         trustedCertificateGroups.TrustedCertificateGroupsAPI
	UnknownResources                 unknownResources.UnknownResourcesAPI
	Users                            users.UsersAPI
	Version                          version.VersionAPI
	Virtualhosts                     virtualhosts.VirtualhostsAPI
	WebSessionManagement             webSessionManagement.WebSessionManagementAPI
	WebSessions                      webSessions.WebSessionsAPI

	apiVersion string
}

// Client configures and returns a fully initialized PAClient
func (c *cfg) Client() (*paClient, *tfprotov5.Diagnostic) {
	/* #nosec */
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	u, err := url.ParseRequestURI(c.BaseURL)
	if err != nil {
		return nil, &tfprotov5.Diagnostic{
			Severity: tfprotov5.DiagnosticSeverityError,
			Summary:  "Invalid URL",
			Detail:   fmt.Sprintf("Unable to parse base_url for client: %s", err),
		}
	}

	cfg := paCfg.NewConfig().WithEndpoint(u.String() + c.Context).WithUsername(c.Username).WithPassword(c.Password)

	if os.Getenv("TF_LOG") == "DEBUG" || os.Getenv("TF_LOG") == "TRACE" || os.Getenv("TF_LOG_PROVIDER") == "DEBUG" || os.Getenv("TF_LOG_PROVIDER") == "TRACE" {
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
		return nil, &tfprotov5.Diagnostic{
			Severity: tfprotov5.DiagnosticSeverityError,
			Summary:  "Connection Error",
			Detail:   fmt.Sprintf("Unable to connect to PingAccess: %s", checkErr(err)),
		}
	}
	client.apiVersion = *v.Version
	client.AccessTokenValidatorsDescriptors, _, err = client.AccessTokenValidators.GetAccessTokenValidatorDescriptorsCommand()
	if err != nil {
		return nil, &tfprotov5.Diagnostic{
			Severity: tfprotov5.DiagnosticSeverityError,
			Summary:  "Setup Error",
			Detail:   fmt.Sprintf("Unable to retrieve AccessTokenValidatorsDescriptors: %s", err.Error()),
		}
	}
	client.SiteAuthenticatorsDescriptors, _, err = client.SiteAuthenticators.GetSiteAuthenticatorDescriptorsCommand()
	if err != nil {
		return nil, &tfprotov5.Diagnostic{
			Severity: tfprotov5.DiagnosticSeverityError,
			Summary:  "Setup Error",
			Detail:   fmt.Sprintf("Unable to retrieve SiteAuthenticatorsDescriptors: %s", err.Error()),
		}
	}

	return &client, nil
}

// Checks whether we are running against PingAccess 6.1 or above and can track password changes
func (c paClient) CanMaskPasswords() bool {
	re := regexp.MustCompile(`^(6\.[1-9])`)
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
