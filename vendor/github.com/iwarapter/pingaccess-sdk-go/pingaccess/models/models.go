package models

import (
	"encoding/json"
	"time"
)

//AccessTokenValidatorView - An access token validator.
type AccessTokenValidatorView struct {
	ClassName     *string                `json:"className"`
	Configuration map[string]interface{} `json:"configuration"`
	Id            json.Number            `json:"id,omitempty"`
	Name          *string                `json:"name"`
}

//AccessTokenValidatorsView - A collection of access token validators.
type AccessTokenValidatorsView struct {
	Items []*AccessTokenValidatorView `json:"items"`
}

//AcmeAccountView - An ACME Account associated with a CA.
type AcmeAccountView struct {
	AcmeServerId *string          `json:"acmeServerId,omitempty"`
	Id           *string          `json:"id,omitempty"`
	KeyAlgorithm *string          `json:"keyAlgorithm,omitempty"`
	PrivateKey   *HiddenFieldView `json:"privateKey,omitempty"`
	PublicKey    *PublicKeyView   `json:"publicKey,omitempty"`
	Url          *string          `json:"url,omitempty"`
}

//AcmeCertStatusView - The status of a certificate.
type AcmeCertStatusView struct {
	Problems map[string]*ProblemDocumentView `json:"problems"`
	State    *string                         `json:"state"`
}

//AcmeCertificateRequestView - A reference to a Key Pair to be signed by the ACME protocol.
type AcmeCertificateRequestView struct {
	AcmeAccountId  *string             `json:"acmeAccountId"`
	AcmeCertStatus *AcmeCertStatusView `json:"acmeCertStatus"`
	AcmeServerId   *string             `json:"acmeServerId"`
	Id             *string             `json:"id,omitempty"`
	KeyPairId      *int                `json:"keyPairId"`
	Url            *string             `json:"url"`
}

//AcmeServerView - An ACME server.
type AcmeServerView struct {
	AcmeAccounts []*LinkView `json:"acmeAccounts,omitempty"`
	Id           *string     `json:"id,omitempty"`
	Name         *string     `json:"name"`
	Url          *string     `json:"url"`
}

//AcmeServersView - A collection of ACME servers.
type AcmeServersView struct {
	Items []*AcmeServerView `json:"items"`
}

//AdminBasicWebSessionView - An admin basic web session.
type AdminBasicWebSessionView struct {
	Audience                     *string `json:"audience"`
	CookieDomain                 *string `json:"cookieDomain,omitempty"`
	CookieType                   *string `json:"cookieType"`
	ExpirationWarningInMinutes   *int    `json:"expirationWarningInMinutes"`
	IdleTimeoutInMinutes         *int    `json:"idleTimeoutInMinutes"`
	SessionPollIntervalInSeconds *int    `json:"sessionPollIntervalInSeconds"`
	SessionTimeoutInMinutes      *int    `json:"sessionTimeoutInMinutes"`
}

//AdminConfigurationView - An admin configuration.
type AdminConfigurationView struct {
	HostPort     *string `json:"hostPort"`
	HttpProxyId  *int    `json:"httpProxyId,omitempty"`
	HttpsProxyId *int    `json:"httpsProxyId,omitempty"`
}

//AdminWebSessionOidcConfigurationView - An admin web session OIDC configuration.
type AdminWebSessionOidcConfigurationView struct {
	CacheUserAttributes           *bool                       `json:"cacheUserAttributes,omitempty"`
	ClientCredentials             *OAuthClientCredentialsView `json:"clientCredentials"`
	EnableRefreshUser             *bool                       `json:"enableRefreshUser,omitempty"`
	OidcLoginType                 *string                     `json:"oidcLoginType,omitempty"`
	PfsessionStateCacheInSeconds  *int                        `json:"pfsessionStateCacheInSeconds,omitempty"`
	PkceChallengeType             *string                     `json:"pkceChallengeType,omitempty"`
	RefreshUserInfoClaimsInterval *int                        `json:"refreshUserInfoClaimsInterval,omitempty"`
	RequestPreservationType       *string                     `json:"requestPreservationType,omitempty"`
	RequestProfile                *bool                       `json:"requestProfile,omitempty"`
	Scopes                        *[]*string                  `json:"scopes,omitempty"`
	SendRequestedUrlToProvider    *bool                       `json:"sendRequestedUrlToProvider,omitempty"`
	ValidateSessionIsAlive        *bool                       `json:"validateSessionIsAlive,omitempty"`
	WebStorageType                *string                     `json:"webStorageType,omitempty"`
}

//AgentCertificateView - An agent certificate.
type AgentCertificateView struct {
	Alias                   *string        `json:"alias"`
	ChainCertificate        *bool          `json:"chainCertificate"`
	Expires                 *string        `json:"expires"`
	Id                      *int           `json:"id,omitempty"`
	IssuerDn                *string        `json:"issuerDn"`
	KeyPair                 *bool          `json:"keyPair"`
	Md5sum                  *string        `json:"md5sum"`
	SerialNumber            *string        `json:"serialNumber"`
	Sha1sum                 *string        `json:"sha1sum"`
	SignatureAlgorithm      *string        `json:"signatureAlgorithm"`
	Status                  *string        `json:"status"`
	SubjectAlternativeNames []*GeneralName `json:"subjectAlternativeNames,omitempty"`
	SubjectCn               *string        `json:"subjectCn,omitempty"`
	SubjectDn               *string        `json:"subjectDn"`
	TrustedCertificate      *bool          `json:"trustedCertificate"`
	ValidFrom               *string        `json:"validFrom"`
}

//AgentCertificatesView - A collection of agent certificates.
type AgentCertificatesView struct {
	Items []*AgentCertificateView `json:"items"`
}

//AgentView - An agent.
type AgentView struct {
	CertificateHash       *Hash                   `json:"certificateHash,omitempty"`
	Description           *string                 `json:"description,omitempty"`
	FailedRetryTimeout    *int                    `json:"failedRetryTimeout,omitempty"`
	FailoverHosts         *[]*string              `json:"failoverHosts,omitempty"`
	Hostname              *string                 `json:"hostname"`
	Id                    json.Number             `json:"id,omitempty"`
	IpSource              *IpMultiValueSourceView `json:"ipSource,omitempty"`
	MaxRetries            *int                    `json:"maxRetries,omitempty"`
	Name                  *string                 `json:"name"`
	OverrideIpSource      *bool                   `json:"overrideIpSource,omitempty"`
	Port                  *int                    `json:"port"`
	SelectedCertificateId *int                    `json:"selectedCertificateId,omitempty"`
	SharedSecretIds       *[]*int                 `json:"sharedSecretIds"`
	UnknownResourceMode   *string                 `json:"unknownResourceMode,omitempty"`
}

//AgentsView - A collection of agents.
type AgentsView struct {
	Items []*AgentView `json:"items"`
}

//AlgorithmView - An algorithm.
type AlgorithmView struct {
	Description *string `json:"description"`
	Name        *string `json:"name"`
}

//AlgorithmsView - A collection of valid web session encryption algorithms.
type AlgorithmsView struct {
	Items []*AlgorithmView `json:"items"`
}

//ApiErrorView - An API error used by the PingAccess Administrative UI.
type ApiErrorView struct {
	Flash *[]*string            `json:"flash"`
	Form  map[string]*[]*string `json:"form"`
}

//ApplicationView - An application.
type ApplicationView struct {
	AccessValidatorId                     *int                      `json:"accessValidatorId,omitempty"`
	AgentCacheInvalidatedExpiration       *int                      `json:"agentCacheInvalidatedExpiration,omitempty"`
	AgentCacheInvalidatedResponseDuration *int                      `json:"agentCacheInvalidatedResponseDuration,omitempty"`
	AgentId                               *int                      `json:"agentId"`
	AllowEmptyPathSegments                *bool                     `json:"allowEmptyPathSegments,omitempty"`
	ApplicationType                       *string                   `json:"applicationType,omitempty"`
	CaseSensitivePath                     *bool                     `json:"caseSensitivePath,omitempty"`
	ContextRoot                           *string                   `json:"contextRoot"`
	DefaultAuthType                       *string                   `json:"defaultAuthType"`
	Description                           *string                   `json:"description,omitempty"`
	Destination                           *string                   `json:"destination,omitempty"`
	Enabled                               *bool                     `json:"enabled,omitempty"`
	Id                                    json.Number               `json:"id,omitempty"`
	IdentityMappingIds                    map[string]*int           `json:"identityMappingIds,omitempty"`
	Issuer                                *string                   `json:"issuer,omitempty"`
	LastModified                          *int                      `json:"lastModified,omitempty"`
	ManualOrderingEnabled                 *bool                     `json:"manualOrderingEnabled,omitempty"`
	Name                                  *string                   `json:"name"`
	Policy                                map[string]*[]*PolicyItem `json:"policy,omitempty"`
	Realm                                 *string                   `json:"realm,omitempty"`
	RequireHTTPS                          *bool                     `json:"requireHTTPS,omitempty"`
	ResourceOrder                         *[]*int                   `json:"resourceOrder,omitempty"`
	SiteId                                *int                      `json:"siteId"`
	SpaSupportEnabled                     *bool                     `json:"spaSupportEnabled"`
	VirtualHostIds                        *[]*int                   `json:"virtualHostIds"`
	WebSessionId                          *int                      `json:"webSessionId,omitempty"`
}

//ApplicationsView - A collection of applications.
type ApplicationsView struct {
	Items []*ApplicationView `json:"items"`
}

//AttributeView - A name-value pair of user attributes.
type AttributeView struct {
	AttributeName  *string `json:"attributeName"`
	AttributeValue *string `json:"attributeValue"`
}

//AuthTokenManagementView - An auth token management configuration.
type AuthTokenManagementView struct {
	Issuer               *string `json:"issuer,omitempty"`
	KeyRollEnabled       *bool   `json:"keyRollEnabled,omitempty"`
	KeyRollPeriodInHours *int    `json:"keyRollPeriodInHours,omitempty"`
	SigningAlgorithm     *string `json:"signingAlgorithm,omitempty"`
}

//AuthnReqListView - An authentication requirements list.
type AuthnReqListView struct {
	AuthnReqs *[]*string  `json:"authnReqs"`
	Id        json.Number `json:"id,omitempty"`
	Name      *string     `json:"name"`
}

//AuthnReqListsView - A collection of authentication requirements lists.
type AuthnReqListsView struct {
	Items []*AuthnReqListView `json:"items"`
}

//AuthorizationServerView - The third-party OAuth 2.0 Authorization Server configuration.
type AuthorizationServerView struct {
	AuditLevel                *string                     `json:"auditLevel,omitempty"`
	CacheTokens               *bool                       `json:"cacheTokens,omitempty"`
	ClientCredentials         *OAuthClientCredentialsView `json:"clientCredentials"`
	Description               *string                     `json:"description,omitempty"`
	IntrospectionEndpoint     *string                     `json:"introspectionEndpoint"`
	Secure                    *bool                       `json:"secure,omitempty"`
	SendAudience              *bool                       `json:"sendAudience,omitempty"`
	SubjectAttributeName      *string                     `json:"subjectAttributeName"`
	Targets                   *[]*string                  `json:"targets"`
	TokenTimeToLiveSeconds    *int                        `json:"tokenTimeToLiveSeconds,omitempty"`
	TrustedCertificateGroupId *int                        `json:"trustedCertificateGroupId"`
	UseProxy                  *bool                       `json:"useProxy,omitempty"`
}

//AvailabilityProfileView - An availability profile.
type AvailabilityProfileView struct {
	ClassName     *string                `json:"className"`
	Configuration map[string]interface{} `json:"configuration"`
	Id            json.Number            `json:"id,omitempty"`
	Name          *string                `json:"name"`
}

//AvailabilityProfilesView - A collection of availability profiles.
type AvailabilityProfilesView struct {
	Items []*AvailabilityProfileView `json:"items"`
}

//BasicAuthConfigView - A basic authentication configuration.
type BasicAuthConfigView struct {
	Enabled *bool `json:"enabled,omitempty"`
}

//BasicConfig - A basic authentication configuration.
type BasicConfig struct {
	Enabled *bool `json:"enabled"`
}

//CSRResponseImportDocView - A CSR response.
type CSRResponseImportDocView struct {
	ChainCertificates  *[]*string `json:"chainCertificates"`
	FileData           *string    `json:"fileData"`
	HsmProviderId      *int       `json:"hsmProviderId"`
	TrustedCertGroupId *int       `json:"trustedCertGroupId"`
}

//ChainCertificateView - A chain certificate.
type ChainCertificateView struct {
	Alias                   *string        `json:"alias"`
	Expires                 *int           `json:"expires"`
	Id                      *int           `json:"id,omitempty"`
	IssuerDn                *string        `json:"issuerDn"`
	Md5sum                  *string        `json:"md5sum"`
	SerialNumber            *string        `json:"serialNumber"`
	Sha1sum                 *string        `json:"sha1sum"`
	SignatureAlgorithm      *string        `json:"signatureAlgorithm"`
	Status                  *string        `json:"status"`
	SubjectAlternativeNames []*GeneralName `json:"subjectAlternativeNames,omitempty"`
	SubjectCn               *string        `json:"subjectCn,omitempty"`
	SubjectDn               *string        `json:"subjectDn"`
	ValidFrom               *int           `json:"validFrom"`
}

//ChainCertificatesDocView
type ChainCertificatesDocView struct {
	AddChainCertificates *[]*string `json:"addChainCertificates"`
}

//ConfigStatusView - An import or export configuration.
type ConfigStatusView struct {
	ApiErrorView  *ApiErrorView          `json:"apiErrorView,omitempty"`
	CurrentEntity map[string]interface{} `json:"currentEntity,omitempty"`
	Id            *int                   `json:"id,omitempty"`
	Status        *string                `json:"status,omitempty"`
	TotalEntities *int                   `json:"totalEntities,omitempty"`
	Warnings      *[]*string             `json:"warnings"`
}

//ConfigStatusesView - A collection of import or export configuration workflows.
type ConfigStatusesView struct {
	Items []*ConfigStatusView `json:"items"`
}

//ConfigurationDependentFieldOption - Configuration of the dependent field option.
type ConfigurationDependentFieldOption struct {
	Options []*ConfigurationOption `json:"options"`
	Value   *string                `json:"value"`
}

//ConfigurationField - Details for configuration in the administrator console.
type ConfigurationField struct {
	Advanced     *bool                     `json:"advanced"`
	ButtonGroup  *string                   `json:"buttonGroup"`
	Default      *string                   `json:"default"`
	Deselectable *bool                     `json:"deselectable"`
	Fields       []*ConfigurationField     `json:"fields"`
	Help         *Help                     `json:"help"`
	Label        *string                   `json:"label"`
	Name         *string                   `json:"name"`
	Options      []*ConfigurationOption    `json:"options"`
	ParentField  *ConfigurationParentField `json:"parentField"`
	Required     *bool                     `json:"required"`
	Type         *string                   `json:"type"`
}

//ConfigurationOption - The configuration option attributes.
type ConfigurationOption struct {
	Category *string `json:"category"`
	Label    *string `json:"label"`
	Value    *string `json:"value"`
}

//ConfigurationParentField - Configuration of the parent field.
type ConfigurationParentField struct {
	DependentFieldOptions []*ConfigurationDependentFieldOption `json:"dependentFieldOptions"`
	FieldName             *string                              `json:"fieldName"`
}

//CookieTypesView - A collection of valid values for the web session cookie type.
type CookieTypesView struct {
	Items []*ItemView `json:"items"`
}

//DescriptorView - A descriptor.
type DescriptorView struct {
	ClassName           *string               `json:"className"`
	ConfigurationFields []*ConfigurationField `json:"configurationFields"`
	Label               *string               `json:"label"`
	Type                *string               `json:"type"`
}

//DescriptorsView - A list of descriptors.
type DescriptorsView struct {
	Items []*DescriptorView `json:"items"`
}

//EngineCertificateView - An engine certificate.
type EngineCertificateView struct {
	Alias                   *string        `json:"alias"`
	ChainCertificate        *bool          `json:"chainCertificate"`
	Expires                 *string        `json:"expires"`
	Id                      *int           `json:"id,omitempty"`
	IssuerDn                *string        `json:"issuerDn"`
	KeyPair                 *bool          `json:"keyPair"`
	Md5sum                  *string        `json:"md5sum"`
	SerialNumber            *string        `json:"serialNumber"`
	Sha1sum                 *string        `json:"sha1sum"`
	SignatureAlgorithm      *string        `json:"signatureAlgorithm"`
	Status                  *string        `json:"status"`
	SubjectAlternativeNames []*GeneralName `json:"subjectAlternativeNames,omitempty"`
	SubjectCn               *string        `json:"subjectCn,omitempty"`
	SubjectDn               *string        `json:"subjectDn"`
	TrustedCertificate      *bool          `json:"trustedCertificate"`
	ValidFrom               *string        `json:"validFrom"`
}

//EngineHealthStatusView
type EngineHealthStatusView struct {
	CurrentServerTime *int                   `json:"currentServerTime"`
	EnginesStatus     map[string]*EngineInfo `json:"enginesStatus"`
}

//EngineInfo
type EngineInfo struct {
	Description  *string `json:"description"`
	LastUpdated  *int    `json:"lastUpdated"`
	Name         *string `json:"name"`
	PollingDelay *int    `json:"pollingDelay"`
}

//EngineListenerView - An engine listener.
type EngineListenerView struct {
	Id                        json.Number `json:"id,omitempty"`
	Name                      *string     `json:"name"`
	Port                      *int        `json:"port"`
	Secure                    *bool       `json:"secure,omitempty"`
	TrustedCertificateGroupId *int        `json:"trustedCertificateGroupId,omitempty"`
}

//EngineListenersView - A collection of engine listeners.
type EngineListenersView struct {
	Items []*EngineListenerView `json:"items"`
}

//EngineView - An engine.
type EngineView struct {
	CertificateHash          *Hash            `json:"certificateHash,omitempty"`
	ConfigReplicationEnabled *bool            `json:"configReplicationEnabled,omitempty"`
	Description              *string          `json:"description,omitempty"`
	HttpProxyId              *int             `json:"httpProxyId,omitempty"`
	HttpsProxyId             *int             `json:"httpsProxyId,omitempty"`
	Id                       json.Number      `json:"id,omitempty"`
	Keys                     []*PublicKeyView `json:"keys,omitempty"`
	Name                     *string          `json:"name"`
	SelectedCertificateId    *int             `json:"selectedCertificateId,omitempty"`
}

//EnginesView - A collection of engines.
type EnginesView struct {
	Items []*EngineView `json:"items"`
}

//ExportData - A JSON backup file.
type ExportData struct {
	Data          map[string]interface{} `json:"data"`
	EncryptionKey *JsonWebKey            `json:"encryptionKey"`
	Id            json.Number            `json:"id,omitempty"`
	MasterKeys    *MasterKeysView        `json:"masterKeys"`
	Version       *string                `json:"version"`
}

//ExportParameters - The export parameters for a key pair.
type ExportParameters struct {
	HsmProviderId *int    `json:"hsmProviderId"`
	Id            *int    `json:"id"`
	Password      *string `json:"password"`
}

//GeneralName - A subject alternative name.
type GeneralName struct {
	Name  *string `json:"name"`
	Value *string `json:"value"`
}

//GlobalUnprotectedResourceView - A global unprotected resource.
type GlobalUnprotectedResourceView struct {
	AuditLevel   *string     `json:"auditLevel,omitempty"`
	Description  *string     `json:"description,omitempty"`
	Enabled      *bool       `json:"enabled,omitempty"`
	Id           json.Number `json:"id,omitempty"`
	Name         *string     `json:"name"`
	WildcardPath *string     `json:"wildcardPath"`
}

//GlobalUnprotectedResourcesView - A collection of global unprotected resource items.
type GlobalUnprotectedResourcesView struct {
	Items []*GlobalUnprotectedResourceView `json:"items"`
}

//Hash - A hash.
type Hash struct {
	Algorithm *string `json:"algorithm"`
	HexValue  *string `json:"hexValue"`
}

//Help - Configuration of the help context of a configuration field.
type Help struct {
	Content *string `json:"content"`
	Title   *string `json:"title"`
	Url     *string `json:"url"`
}

//HiddenFieldView - A hidden field.
type HiddenFieldView struct {
	EncryptedValue *string `json:"encryptedValue,omitempty"`
	Value          *string `json:"value,omitempty"`
}

//HostMultiValueSourceView - Configuration for the host source.
type HostMultiValueSourceView struct {
	HeaderNameList    *[]*string `json:"headerNameList"`
	ListValueLocation *string    `json:"listValueLocation"`
}

//HostPortView - A redirect source.
type HostPortView struct {
	Host *string `json:"host"`
	Port *int    `json:"port"`
}

//HsmProviderView - An HSM provider.
type HsmProviderView struct {
	ClassName     *string                `json:"className"`
	Configuration map[string]interface{} `json:"configuration"`
	Id            json.Number            `json:"id,omitempty"`
	Name          *string                `json:"name"`
}

//HttpClientProxyView - A proxy.
type HttpClientProxyView struct {
	Description            *string          `json:"description,omitempty"`
	Host                   *string          `json:"host"`
	Id                     json.Number      `json:"id,omitempty"`
	Name                   *string          `json:"name"`
	Password               *HiddenFieldView `json:"password,omitempty"`
	Port                   *int             `json:"port"`
	RequiresAuthentication *bool            `json:"requiresAuthentication,omitempty"`
	Username               *string          `json:"username,omitempty"`
}

//HttpMonitoringView
type HttpMonitoringView struct {
	AuditLevel *string `json:"auditLevel"`
}

//HttpsListenerView - An HTTPS listener.
type HttpsListenerView struct {
	Id                        json.Number `json:"id,omitempty"`
	KeyPairId                 *int        `json:"keyPairId"`
	Name                      *string     `json:"name"`
	RestartRequired           *bool       `json:"restartRequired"`
	UseServerCipherSuiteOrder *bool       `json:"useServerCipherSuiteOrder"`
}

//HttpsListenersView - A collection of HTTPS listeners.
type HttpsListenersView struct {
	Items []*HttpsListenerView `json:"items"`
}

//IdentityMappingView - An identity mapping.
type IdentityMappingView struct {
	ClassName     *string                `json:"className"`
	Configuration map[string]interface{} `json:"configuration"`
	Id            json.Number            `json:"id,omitempty"`
	Name          *string                `json:"name"`
}

//IdentityMappingsView - A collection of identity mappings.
type IdentityMappingsView struct {
	Items []*IdentityMappingView `json:"items"`
}

//IpMultiValueSourceView - Configuration for the IP source.
type IpMultiValueSourceView struct {
	FallbackToLastHopIp *bool      `json:"fallbackToLastHopIp,omitempty"`
	HeaderNameList      *[]*string `json:"headerNameList"`
	ListValueLocation   *string    `json:"listValueLocation"`
}

//ItemView - An item.
type ItemView struct {
	Description *string `json:"description"`
	Name        *string `json:"name"`
}

//JsonWebKey - A JSON Web Key.
type JsonWebKey struct {
	Algorithm *string    `json:"algorithm"`
	Key       *Key       `json:"key"`
	KeyId     *string    `json:"keyId"`
	KeyOps    *[]*string `json:"keyOps"`
	KeyType   *string    `json:"keyType"`
	PublicKey *PublicKey `json:"publicKey"`
	Use       *string    `json:"use"`
}

//Key - A key.
type Key struct {
	Algorithm *string  `json:"algorithm"`
	Encoded   *[]*byte `json:"encoded"`
	Format    *string  `json:"format"`
}

//KeyAlgorithm - A key algorithm.
type KeyAlgorithm struct {
	DefaultKeySize            *int       `json:"defaultKeySize"`
	DefaultSignatureAlgorithm *string    `json:"defaultSignatureAlgorithm"`
	KeySizes                  *[]*int    `json:"keySizes"`
	Name                      *string    `json:"name"`
	SignatureAlgorithms       *[]*string `json:"signatureAlgorithms"`
}

//KeyAlgorithmsView - A collection of key algorithms.
type KeyAlgorithmsView struct {
	Items []*KeyAlgorithm `json:"items"`
}

//KeyPairView - A key pair.
type KeyPairView struct {
	Alias                   *string                 `json:"alias"`
	ChainCertificates       []*ChainCertificateView `json:"chainCertificates,omitempty"`
	CsrPending              *bool                   `json:"csrPending"`
	Expires                 *int                    `json:"expires"`
	HsmProviderId           *int                    `json:"hsmProviderId,omitempty"`
	Id                      *int                    `json:"id,omitempty"`
	IssuerDn                *string                 `json:"issuerDn"`
	Md5sum                  *string                 `json:"md5sum"`
	SerialNumber            *string                 `json:"serialNumber"`
	Sha1sum                 *string                 `json:"sha1sum"`
	SignatureAlgorithm      *string                 `json:"signatureAlgorithm"`
	Status                  *string                 `json:"status"`
	SubjectAlternativeNames []*GeneralName          `json:"subjectAlternativeNames,omitempty"`
	SubjectCn               *string                 `json:"subjectCn,omitempty"`
	SubjectDn               *string                 `json:"subjectDn"`
	ValidFrom               *int                    `json:"validFrom"`
}

//KeyPairsView - A collection of key pairs.
type KeyPairsView struct {
	Items []*KeyPairView `json:"items"`
}

//KeySetView - An auth token key set configuration.
type KeySetView struct {
	KeySet *string `json:"keySet"`
	Nonce  *string `json:"nonce"`
}

//LicenseImportDocView - A license file.
type LicenseImportDocView struct {
	FileData *string `json:"fileData"`
}

//LicenseView - A Ping Identity license.
type LicenseView struct {
	EnforcementType *int    `json:"enforcementType"`
	ExpirationDate  *string `json:"expirationDate"`
	Id              *int    `json:"id"`
	IssueDate       *string `json:"issueDate"`
	MaxApplications *int    `json:"maxApplications"`
	Name            *string `json:"name"`
	Organization    *string `json:"organization"`
	Product         *string `json:"product"`
	Tier            *string `json:"tier"`
	Version         *string `json:"version"`
}

//LinkView - A reference to the associated resource.
type LinkView struct {
	Id       *string `json:"id"`
	Location *string `json:"location"`
}

//List
type List struct {
	Empty *bool `json:"empty"`
}

//LoadBalancingStrategiesView - A collection of load balancing strategies.
type LoadBalancingStrategiesView struct {
	Items []*LoadBalancingStrategyView `json:"items"`
}

//LoadBalancingStrategyView - A load balancing strategy.
type LoadBalancingStrategyView struct {
	ClassName     *string                `json:"className"`
	Configuration map[string]interface{} `json:"configuration"`
	Id            json.Number            `json:"id,omitempty"`
	Name          *string                `json:"name"`
}

//MasterKeysView - An encrypted master key.
type MasterKeysView struct {
	EncryptedValue *[]*byte `json:"encryptedValue"`
	KeyId          *string  `json:"keyId"`
}

//MethodView - HTTP Method
type MethodView struct {
	Name *string `json:"name"`
}

//MethodsView
type MethodsView struct {
	Items []*MethodView `json:"items"`
}

//NewKeyPairConfigView - A new key pair.
type NewKeyPairConfigView struct {
	Alias                   *string        `json:"alias"`
	City                    *string        `json:"city"`
	CommonName              *string        `json:"commonName"`
	Country                 *string        `json:"country"`
	HsmProviderId           *int           `json:"hsmProviderId"`
	Id                      *int           `json:"id,omitempty"`
	KeyAlgorithm            *string        `json:"keyAlgorithm"`
	KeySize                 *int           `json:"keySize"`
	Organization            *string        `json:"organization"`
	OrganizationUnit        *string        `json:"organizationUnit"`
	SignatureAlgorithm      *string        `json:"signatureAlgorithm,omitempty"`
	State                   *string        `json:"state"`
	SubjectAlternativeNames []*GeneralName `json:"subjectAlternativeNames,omitempty"`
	ValidDays               *int           `json:"validDays"`
}

//OAuthClientCredentialsView - OAuth client credentials.
type OAuthClientCredentialsView struct {
	ClientId     *string          `json:"clientId"`
	ClientSecret *HiddenFieldView `json:"clientSecret,omitempty"`
}

//OAuthConfigView - An OAuth authentication configuration.
type OAuthConfigView struct {
	ClientId             *string                       `json:"clientId"`
	ClientSecret         *HiddenFieldView              `json:"clientSecret,omitempty"`
	Enabled              *bool                         `json:"enabled,omitempty"`
	RoleMapping          *RoleMappingConfigurationView `json:"roleMapping,omitempty"`
	Scope                *string                       `json:"scope"`
	SubjectAttributeName *string                       `json:"subjectAttributeName,omitempty"`
}

//OAuthKeyManagementView - An OAuth key management configuration.
type OAuthKeyManagementView struct {
	KeyRollEnabled       *bool `json:"keyRollEnabled,omitempty"`
	KeyRollPeriodInHours *int  `json:"keyRollPeriodInHours,omitempty"`
}

//OIDCProviderMetadata - The OpenID Connect provider's metadata.
type OIDCProviderMetadata struct {
	AuthorizationEndpoint                  *string    `json:"authorization_endpoint"`
	BackchannelAuthenticationEndpoint      *string    `json:"backchannel_authentication_endpoint"`
	ClaimTypesSupported                    *[]*string `json:"claim_types_supported"`
	ClaimsParameterSupported               *bool      `json:"claims_parameter_supported"`
	ClaimsSupported                        *[]*string `json:"claims_supported"`
	CodeChallengeMethodsSupported          *[]*string `json:"code_challenge_methods_supported"`
	EndSessionEndpoint                     *string    `json:"end_session_endpoint"`
	GrantTypesSupported                    *[]*string `json:"grant_types_supported"`
	IdTokenSigningAlgValuesSupported       *[]*string `json:"id_token_signing_alg_values_supported"`
	IntrospectionEndpoint                  *string    `json:"introspection_endpoint"`
	Issuer                                 *string    `json:"issuer"`
	JwksUri                                *string    `json:"jwks_uri"`
	PingEndSessionEndpoint                 *string    `json:"ping_end_session_endpoint"`
	PingRevokedSrisEndpoint                *string    `json:"ping_revoked_sris_endpoint"`
	RequestObjectSigningAlgValuesSupported *[]*string `json:"request_object_signing_alg_values_supported"`
	RequestParameterSupported              *bool      `json:"request_parameter_supported"`
	RequestUriParameterSupported           *bool      `json:"request_uri_parameter_supported"`
	ResponseModesSupported                 *[]*string `json:"response_modes_supported"`
	ResponseTypesSupported                 *[]*string `json:"response_types_supported"`
	RevocationEndpoint                     *string    `json:"revocation_endpoint"`
	ScopesSupported                        *[]*string `json:"scopes_supported"`
	SubjectTypesSupported                  *[]*string `json:"subject_types_supported"`
	TokenEndpoint                          *string    `json:"token_endpoint"`
	TokenEndpointAuthMethodsSupported      *[]*string `json:"token_endpoint_auth_methods_supported"`
	UserinfoEndpoint                       *string    `json:"userinfo_endpoint"`
	UserinfoSigningAlgValuesSupported      *[]*string `json:"userinfo_signing_alg_values_supported"`
}

//OIDCProviderPluginView - An OpenID Connect provider plugin.
type OIDCProviderPluginView struct {
	ClassName     *string                `json:"className"`
	Configuration map[string]interface{} `json:"configuration"`
}

//OIDCProviderView - The third-party OpenID Connect provider configuration.
type OIDCProviderView struct {
	AuditLevel                 *string                 `json:"auditLevel,omitempty"`
	Description                *string                 `json:"description,omitempty"`
	Issuer                     *string                 `json:"issuer"`
	Plugin                     *OIDCProviderPluginView `json:"plugin,omitempty"`
	QueryParameters            []*QueryParameterView   `json:"queryParameters,omitempty"`
	RequestSupportedScopesOnly *bool                   `json:"requestSupportedScopesOnly,omitempty"`
	TrustedCertificateGroupId  *int                    `json:"trustedCertificateGroupId,omitempty"`
	UseProxy                   *bool                   `json:"useProxy,omitempty"`
	UseSlo                     *bool                   `json:"useSlo,omitempty"`
}

//OidcConfigView - An OIDC authentication configuration.
type OidcConfigView struct {
	AuthnReqListId    *int                                  `json:"authnReqListId,omitempty"`
	Enabled           *bool                                 `json:"enabled,omitempty"`
	OidcConfiguration *AdminWebSessionOidcConfigurationView `json:"oidcConfiguration"`
	RoleMapping       *RoleMappingConfigurationView         `json:"roleMapping,omitempty"`
	UseSlo            *bool                                 `json:"useSlo,omitempty"`
}

//OidcLoginTypesView - A collection of valid web session OIDC login types.
type OidcLoginTypesView struct {
	Items []*ItemView `json:"items"`
}

//OptionalAttributeMappingView - A set of user attributes that define an optional role mapping.
type OptionalAttributeMappingView struct {
	Attributes []*AttributeView `json:"attributes"`
	Enabled    *bool            `json:"enabled,omitempty"`
}

//PKCS12FileImportDocView - A PKCS#12 file.
type PKCS12FileImportDocView struct {
	Alias             *string    `json:"alias"`
	ChainCertificates *[]*string `json:"chainCertificates"`
	FileData          *string    `json:"fileData"`
	HsmProviderId     *int       `json:"hsmProviderId"`
	Password          *string    `json:"password"`
}

//PathPatternView - A pattern for matching request URI paths.
type PathPatternView struct {
	Pattern *string `json:"pattern"`
	Type    *string `json:"type"`
}

//PingFederateAccessTokenView - A PingAccess OAuth client configuration.
type PingFederateAccessTokenView struct {
	AccessValidatorId      *int             `json:"accessValidatorId,omitempty"`
	CacheTokens            *bool            `json:"cacheTokens,omitempty"`
	ClientId               *string          `json:"clientId"`
	ClientSecret           *HiddenFieldView `json:"clientSecret,omitempty"`
	Name                   *string          `json:"name,omitempty"`
	SendAudience           *bool            `json:"sendAudience,omitempty"`
	SubjectAttributeName   *string          `json:"subjectAttributeName"`
	TokenTimeToLiveSeconds *int             `json:"tokenTimeToLiveSeconds,omitempty"`
	UseTokenIntrospection  *bool            `json:"useTokenIntrospection,omitempty"`
}

//PingFederateAdminView - A PingFederate Admin configuration.
type PingFederateAdminView struct {
	AdminPassword             *HiddenFieldView `json:"adminPassword"`
	AdminUsername             *string          `json:"adminUsername"`
	AuditLevel                *string          `json:"auditLevel,omitempty"`
	BasePath                  *string          `json:"basePath,omitempty"`
	Host                      *string          `json:"host"`
	Port                      *int             `json:"port"`
	Secure                    *bool            `json:"secure,omitempty"`
	TrustedCertificateGroupId *int             `json:"trustedCertificateGroupId,omitempty"`
	UseProxy                  *bool            `json:"useProxy,omitempty"`
}

//PingFederateMetadataRuntimeView - A PingFederate configuration.
type PingFederateMetadataRuntimeView struct {
	Description               *string `json:"description,omitempty"`
	Issuer                    *string `json:"issuer"`
	SkipHostnameVerification  *bool   `json:"skipHostnameVerification,omitempty"`
	StsTokenExchangeEndpoint  *string `json:"stsTokenExchangeEndpoint,omitempty"`
	TrustedCertificateGroupId *int    `json:"trustedCertificateGroupId,omitempty"`
	UseProxy                  *bool   `json:"useProxy,omitempty"`
	UseSlo                    *bool   `json:"useSlo,omitempty"`
}

//PingFederateRuntimeView - A PingFederate configuration.
type PingFederateRuntimeView struct {
	AuditLevel                *string    `json:"auditLevel,omitempty"`
	BackChannelBasePath       *string    `json:"backChannelBasePath,omitempty"`
	BackChannelSecure         *bool      `json:"backChannelSecure,omitempty"`
	BasePath                  *string    `json:"basePath,omitempty"`
	ExpectedHostname          *string    `json:"expectedHostname,omitempty"`
	Host                      *string    `json:"host"`
	Port                      *int       `json:"port"`
	Secure                    *bool      `json:"secure,omitempty"`
	SkipHostnameVerification  *bool      `json:"skipHostnameVerification,omitempty"`
	Targets                   *[]*string `json:"targets,omitempty"`
	TrustedCertificateGroupId *int       `json:"trustedCertificateGroupId,omitempty"`
	UseProxy                  *bool      `json:"useProxy,omitempty"`
	UseSlo                    *bool      `json:"useSlo,omitempty"`
}

//PingOne4CView - The PingOne for Customers OIDC provider configuration.
type PingOne4CView struct {
	Description               *string `json:"description,omitempty"`
	Issuer                    *string `json:"issuer"`
	TrustedCertificateGroupId *int    `json:"trustedCertificateGroupId,omitempty"`
	UseProxy                  *bool   `json:"useProxy,omitempty"`
}

//PolicyItem - The policy items associated with the application - Now an undocumented model type
type PolicyItem struct {
	Id   json.Number `json:"id,omitempty"`
	Type *string     `json:"type,omitempty"`
}

//ProblemDocumentView - An RFC 7807 problem details object.
type ProblemDocumentView struct {
	Detail *string `json:"detail"`
	Type   *string `json:"type"`
}

//ProtocolSourceView - Configuration for the protocol source.
type ProtocolSourceView struct {
	HeaderName *string `json:"headerName"`
}

//PublicKey - An undocumented model type
type PublicKey struct {
}

//PublicKeyView - A public key.
type PublicKeyView struct {
	Created *time.Time             `json:"created,omitempty"`
	Jwk     map[string]interface{} `json:"jwk"`
}

//QueryParameterView - A name-value pair of custom query parameters.
type QueryParameterView struct {
	Name  *string `json:"name"`
	Value *string `json:"value"`
}

//RedirectView - A Redirect.
type RedirectView struct {
	AuditLevel   *string             `json:"auditLevel,omitempty"`
	ResponseCode *int                `json:"responseCode,omitempty"`
	Source       *HostPortView       `json:"source,omitempty"`
	Target       *TargetHostPortView `json:"target,omitempty"`
}

//RedirectsView - A collection of Redirects.
type RedirectsView struct {
	Items []*RedirectView `json:"items"`
}

//RejectionHandlerView - A rejection handler.
type RejectionHandlerView struct {
	ClassName     *string                `json:"className"`
	Configuration map[string]interface{} `json:"configuration"`
	Id            json.Number            `json:"id,omitempty"`
	Name          *string                `json:"name"`
}

//RejectionHandlersView - A collection of rejection handlers.
type RejectionHandlersView struct {
	Items []*RejectionHandlerView `json:"items"`
}

//ReplicaAdminView - A replica admin.
type ReplicaAdminView struct {
	CertificateHash          *Hash            `json:"certificateHash,omitempty"`
	ConfigReplicationEnabled *bool            `json:"configReplicationEnabled,omitempty"`
	Description              *string          `json:"description,omitempty"`
	HostPort                 *string          `json:"hostPort"`
	HttpProxyId              *int             `json:"httpProxyId,omitempty"`
	HttpsProxyId             *int             `json:"httpsProxyId,omitempty"`
	Id                       json.Number      `json:"id,omitempty"`
	Keys                     []*PublicKeyView `json:"keys,omitempty"`
	Name                     *string          `json:"name"`
	SelectedCertificateId    *int             `json:"selectedCertificateId,omitempty"`
}

//ReplicaAdminsView - A list of replica admins.
type ReplicaAdminsView struct {
	Items []*ReplicaAdminView `json:"items"`
}

//RequestPreservationTypesView - A collection of valid web session request preservation types.
type RequestPreservationTypesView struct {
	Items []*ItemView `json:"items"`
}

//RequiredAttributeMappingView - A set of user attributes that define a mandatory role mapping.
type RequiredAttributeMappingView struct {
	Attributes []*AttributeView `json:"attributes"`
}

//ReservedApplicationView - The reserved application.
type ReservedApplicationView struct {
	ContextRoot *string `json:"contextRoot"`
}

//ResourceMatchingEntryView - A resource matching entry.
type ResourceMatchingEntryView struct {
	Link        *LinkView  `json:"link"`
	Methods     *[]*string `json:"methods"`
	Name        *string    `json:"name"`
	Pattern     *string    `json:"pattern"`
	PatternType *string    `json:"patternType"`
	Type        *string    `json:"type"`
}

//ResourceMatchingEvaluationOrderView - Specifies an ordering of Resource Matching Entries.
type ResourceMatchingEvaluationOrderView struct {
	Entries []*ResourceMatchingEntryView `json:"entries"`
}

//ResourceOrderView - Specifies an ordering of Application Resources.
type ResourceOrderView struct {
	Id          json.Number `json:"id,omitempty"`
	ResourceIds *[]*int     `json:"resourceIds"`
}

//ResourceView - A resource.
type ResourceView struct {
	Anonymous               *bool                     `json:"anonymous,omitempty"`
	ApplicationId           *int                      `json:"applicationId,omitempty"`
	AuditLevel              *string                   `json:"auditLevel,omitempty"`
	DefaultAuthTypeOverride *string                   `json:"defaultAuthTypeOverride"`
	Enabled                 *bool                     `json:"enabled,omitempty"`
	Id                      json.Number               `json:"id,omitempty"`
	Methods                 *[]*string                `json:"methods"`
	Name                    *string                   `json:"name"`
	PathPatterns            []*PathPatternView        `json:"pathPatterns,omitempty"`
	PathPrefixes            *[]*string                `json:"pathPrefixes,omitempty"`
	Policy                  map[string]*[]*PolicyItem `json:"policy,omitempty"`
	RootResource            *bool                     `json:"rootResource,omitempty"`
	Unprotected             *bool                     `json:"unprotected,omitempty"`
}

//ResourcesView - A collection of resources.
type ResourcesView struct {
	Items []*ResourceView `json:"items"`
}

//RoleMappingConfigurationView - Configuration for mapping user attributes to roles.
type RoleMappingConfigurationView struct {
	Administrator *RequiredAttributeMappingView `json:"administrator,omitempty"`
	Auditor       *OptionalAttributeMappingView `json:"auditor,omitempty"`
	Enabled       *bool                         `json:"enabled,omitempty"`
}

//RuleDescriptorView - A rule descriptor.
type RuleDescriptorView struct {
	AgentCachingDisabled *bool                 `json:"agentCachingDisabled"`
	Category             *string               `json:"category"`
	ClassName            *string               `json:"className"`
	ConfigurationFields  []*ConfigurationField `json:"configurationFields"`
	Label                *string               `json:"label"`
	Modes                []*string             `json:"modes"`
	Type                 *string               `json:"type"`
}

//RuleDescriptorsView - A collection of rule descriptors.
type RuleDescriptorsView struct {
	Items []*RuleDescriptorView `json:"items"`
}

//RuleSetElementTypesView - A collection of rule set element types.
type RuleSetElementTypesView struct {
	Items []*ItemView `json:"items"`
}

//RuleSetSuccessCriteriaView - A collection of success criteria.
type RuleSetSuccessCriteriaView struct {
	Items []*ItemView `json:"items"`
}

//RuleSetView - A rule set.
type RuleSetView struct {
	ElementType     *string     `json:"elementType,omitempty"`
	Id              json.Number `json:"id,omitempty"`
	Name            *string     `json:"name"`
	Policy          *[]*int     `json:"policy"`
	SuccessCriteria *string     `json:"successCriteria,omitempty"`
}

//RuleSetsView - A collection of rule sets.
type RuleSetsView struct {
	Items []*RuleSetView `json:"items"`
}

//RuleView - A rule.
type RuleView struct {
	ClassName             *string                `json:"className"`
	Configuration         map[string]interface{} `json:"configuration"`
	Id                    json.Number            `json:"id,omitempty"`
	Name                  *string                `json:"name"`
	SupportedDestinations *[]*string             `json:"supportedDestinations,omitempty"`
}

//RulesView - A collection of rules.
type RulesView struct {
	Items []*RuleView `json:"items"`
}

//SanType - An undocumented model type
type SanType struct {
	Description *string `json:"description,omitempty"`
	Name        *string `json:"name,omitempty"`
}

//SanTypes - A collection of available general names.
type SanTypes struct {
	Items *[]*SanType `json:"items"`
}

//SessionInfo - A session.
type SessionInfo struct {
	AccessControlDirectives *[]*string          `json:"accessControlDirectives"`
	ConfigurationExports    *ConfigStatusesView `json:"configurationExports"`
	ConfigurationImports    *ConfigStatusesView `json:"configurationImports"`
	Exp                     *int                `json:"exp"`
	ExpWarn                 *int                `json:"expWarn"`
	Flash                   *string             `json:"flash"`
	Iat                     *int                `json:"iat"`
	MaxFileUploadSize       *int                `json:"maxFileUploadSize"`
	PollIntervalSeconds     *int                `json:"pollIntervalSeconds"`
	Roles                   *[]*string          `json:"roles"`
	SesTimeout              *int                `json:"sesTimeout"`
	ShowWarning             *bool               `json:"showWarning"`
	SniEnabled              *bool               `json:"sniEnabled"`
	Sub                     *string             `json:"sub"`
	UseSlo                  *bool               `json:"useSlo"`
}

//SharedSecretView - A shared secret.
type SharedSecretView struct {
	Created *string          `json:"created,omitempty"`
	Id      json.Number      `json:"id,omitempty"`
	Secret  *HiddenFieldView `json:"secret"`
}

//SharedSecretsView - A collection of shared secrets.
type SharedSecretsView struct {
	Items []*SharedSecretView `json:"items"`
}

//SigningAlgorithmsView - A collection of valid web session signing algorithms.
type SigningAlgorithmsView struct {
	Items []*AlgorithmView `json:"items"`
}

//SiteAuthenticatorView - A site authenticator.
type SiteAuthenticatorView struct {
	ClassName     *string                `json:"className"`
	Configuration map[string]interface{} `json:"configuration"`
	Id            json.Number            `json:"id,omitempty"`
	Name          *string                `json:"name"`
}

//SiteAuthenticatorsView - A collection of site authenticators.
type SiteAuthenticatorsView struct {
	Items []*SiteAuthenticatorView `json:"items"`
}

//SiteView - A site.
type SiteView struct {
	AvailabilityProfileId     *int        `json:"availabilityProfileId,omitempty"`
	ExpectedHostname          *string     `json:"expectedHostname,omitempty"`
	Id                        json.Number `json:"id,omitempty"`
	KeepAliveTimeout          *int        `json:"keepAliveTimeout,omitempty"`
	LoadBalancingStrategyId   *int        `json:"loadBalancingStrategyId,omitempty"`
	MaxConnections            *int        `json:"maxConnections,omitempty"`
	MaxWebSocketConnections   *int        `json:"maxWebSocketConnections,omitempty"`
	Name                      *string     `json:"name"`
	Secure                    *bool       `json:"secure,omitempty"`
	SendPaCookie              *bool       `json:"sendPaCookie,omitempty"`
	SiteAuthenticatorIds      *[]*int     `json:"siteAuthenticatorIds,omitempty"`
	SkipHostnameVerification  *bool       `json:"skipHostnameVerification,omitempty"`
	Targets                   *[]*string  `json:"targets"`
	TrustedCertificateGroupId *int        `json:"trustedCertificateGroupId,omitempty"`
	UseProxy                  *bool       `json:"useProxy,omitempty"`
	UseTargetHostHeader       *bool       `json:"useTargetHostHeader,omitempty"`
}

//SitesView - A collection of sites items.
type SitesView struct {
	Items []*SiteView `json:"items"`
}

//SupportedScopesView - A set of scopes supported by the OIDC Provider.
type SupportedScopesView struct {
	ClientId *string    `json:"clientId"`
	Scopes   *[]*string `json:"scopes"`
}

//TargetHostPortView - A redirect target.
type TargetHostPortView struct {
	Host   *string `json:"host"`
	Port   *int    `json:"port"`
	Secure *bool   `json:"secure"`
}

//ThirdPartyServiceView - A third-party service.
type ThirdPartyServiceView struct {
	AvailabilityProfileId     *int       `json:"availabilityProfileId,omitempty"`
	ExpectedHostname          *string    `json:"expectedHostname,omitempty"`
	HostValue                 *string    `json:"hostValue,omitempty"`
	Id                        *string    `json:"id,omitempty"`
	LoadBalancingStrategyId   *int       `json:"loadBalancingStrategyId,omitempty"`
	MaxConnections            *int       `json:"maxConnections,omitempty"`
	Name                      *string    `json:"name"`
	Secure                    *bool      `json:"secure,omitempty"`
	SkipHostnameVerification  *bool      `json:"skipHostnameVerification,omitempty"`
	Targets                   *[]*string `json:"targets"`
	TrustedCertificateGroupId *int       `json:"trustedCertificateGroupId,omitempty"`
	UseProxy                  *bool      `json:"useProxy,omitempty"`
}

//ThirdPartyServicesView - A collection of third-party service items.
type ThirdPartyServicesView struct {
	Items []*ThirdPartyServiceView `json:"items"`
}

//TokenProviderSettingView - Settings for a token provider.
type TokenProviderSettingView struct {
	Type          *string `json:"type,omitempty"`
	UseThirdParty *bool   `json:"useThirdParty,omitempty"`
}

//TrustedCertView - A trusted certificate.
type TrustedCertView struct {
	Alias                   *string        `json:"alias"`
	Expires                 *int           `json:"expires"`
	Id                      *int           `json:"id,omitempty"`
	IssuerDn                *string        `json:"issuerDn"`
	Md5sum                  *string        `json:"md5sum"`
	SerialNumber            *string        `json:"serialNumber"`
	Sha1sum                 *string        `json:"sha1sum"`
	SignatureAlgorithm      *string        `json:"signatureAlgorithm"`
	Status                  *string        `json:"status"`
	SubjectAlternativeNames []*GeneralName `json:"subjectAlternativeNames,omitempty"`
	SubjectCn               *string        `json:"subjectCn,omitempty"`
	SubjectDn               *string        `json:"subjectDn"`
	ValidFrom               *int           `json:"validFrom"`
}

//TrustedCertificateGroupView - A trusted certificate group.
type TrustedCertificateGroupView struct {
	CertIds                    *[]*int     `json:"certIds,omitempty"`
	Id                         json.Number `json:"id,omitempty"`
	IgnoreAllCertificateErrors *bool       `json:"ignoreAllCertificateErrors,omitempty"`
	Name                       *string     `json:"name"`
	SkipCertificateDateCheck   *bool       `json:"skipCertificateDateCheck,omitempty"`
	SystemGroup                *bool       `json:"systemGroup,omitempty"`
	UseJavaTrustStore          *bool       `json:"useJavaTrustStore,omitempty"`
}

//TrustedCertificateGroupsView - A collection of trusted certificate group items.
type TrustedCertificateGroupsView struct {
	Items []*TrustedCertificateGroupView `json:"items"`
}

//TrustedCertsView - A collection of trusted certificates.
type TrustedCertsView struct {
	Items []*TrustedCertView `json:"items"`
}

//UnknownResourceSettingsView - Global settings for unknown resources.
type UnknownResourceSettingsView struct {
	AgentDefaultCacheTTL *int    `json:"agentDefaultCacheTTL"`
	AgentDefaultMode     *string `json:"agentDefaultMode"`
	AuditLevel           *string `json:"auditLevel,omitempty"`
	ErrorContentType     *string `json:"errorContentType"`
	ErrorStatusCode      *int    `json:"errorStatusCode"`
	ErrorTemplateFile    *string `json:"errorTemplateFile"`
}

//UserPasswordView - Settings to update a password.
type UserPasswordView struct {
	CurrentPassword *string     `json:"currentPassword"`
	Id              json.Number `json:"id,omitempty"`
	NewPassword     *string     `json:"newPassword"`
}

//UserView - A user.
type UserView struct {
	Email        *string     `json:"email,omitempty"`
	FirstLogin   *bool       `json:"firstLogin,omitempty"`
	Id           json.Number `json:"id,omitempty"`
	ShowTutorial *bool       `json:"showTutorial,omitempty"`
	SlaAccepted  *bool       `json:"slaAccepted,omitempty"`
	Username     *string     `json:"username"`
}

//UsersView - A collection of users.
type UsersView struct {
	Items []*UserView `json:"items"`
}

//VersionDocClass - A version.
type VersionDocClass struct {
	Version *string `json:"version"`
}

//VirtualHostView - A virtual host.
type VirtualHostView struct {
	AgentResourceCacheTTL     *int        `json:"agentResourceCacheTTL,omitempty"`
	Host                      *string     `json:"host"`
	Id                        json.Number `json:"id,omitempty"`
	KeyPairId                 *int        `json:"keyPairId,omitempty"`
	Port                      *int        `json:"port"`
	TrustedCertificateGroupId *int        `json:"trustedCertificateGroupId,omitempty"`
}

//VirtualHostsView - A collection of virtual hosts.
type VirtualHostsView struct {
	Items []*VirtualHostView `json:"items"`
}

//WebSessionManagementView - A web session management configuration.
type WebSessionManagementView struct {
	CookieName                     *string `json:"cookieName,omitempty"`
	EncryptionAlgorithm            *string `json:"encryptionAlgorithm,omitempty"`
	Issuer                         *string `json:"issuer,omitempty"`
	KeyRollEnabled                 *bool   `json:"keyRollEnabled,omitempty"`
	KeyRollPeriodInHours           *int    `json:"keyRollPeriodInHours,omitempty"`
	NonceCookieTimeToLiveInMinutes *int    `json:"nonceCookieTimeToLiveInMinutes,omitempty"`
	SessionStateCookieName         *string `json:"sessionStateCookieName,omitempty"`
	SigningAlgorithm               *string `json:"signingAlgorithm,omitempty"`
	UpdateTokenWindowInSeconds     *int    `json:"updateTokenWindowInSeconds,omitempty"`
}

//WebSessionView - A web session.
type WebSessionView struct {
	Audience                      *string                     `json:"audience"`
	CacheUserAttributes           *bool                       `json:"cacheUserAttributes,omitempty"`
	ClientCredentials             *OAuthClientCredentialsView `json:"clientCredentials"`
	CookieDomain                  *string                     `json:"cookieDomain,omitempty"`
	CookieType                    *string                     `json:"cookieType,omitempty"`
	EnableRefreshUser             *bool                       `json:"enableRefreshUser,omitempty"`
	HttpOnlyCookie                *bool                       `json:"httpOnlyCookie,omitempty"`
	Id                            json.Number                 `json:"id,omitempty"`
	IdleTimeoutInMinutes          *int                        `json:"idleTimeoutInMinutes,omitempty"`
	Name                          *string                     `json:"name"`
	OidcLoginType                 *string                     `json:"oidcLoginType,omitempty"`
	PfsessionStateCacheInSeconds  *int                        `json:"pfsessionStateCacheInSeconds,omitempty"`
	PkceChallengeType             *string                     `json:"pkceChallengeType,omitempty"`
	RefreshUserInfoClaimsInterval *int                        `json:"refreshUserInfoClaimsInterval,omitempty"`
	RequestPreservationType       *string                     `json:"requestPreservationType,omitempty"`
	RequestProfile                *bool                       `json:"requestProfile,omitempty"`
	SameSite                      *string                     `json:"sameSite,omitempty"`
	Scopes                        *[]*string                  `json:"scopes,omitempty"`
	SecureCookie                  *bool                       `json:"secureCookie,omitempty"`
	SendRequestedUrlToProvider    *bool                       `json:"sendRequestedUrlToProvider,omitempty"`
	SessionTimeoutInMinutes       *int                        `json:"sessionTimeoutInMinutes,omitempty"`
	ValidateSessionIsAlive        *bool                       `json:"validateSessionIsAlive,omitempty"`
	WebStorageType                *string                     `json:"webStorageType,omitempty"`
}

//WebSessionsView - A collection of web sessions.
type WebSessionsView struct {
	Items []*WebSessionView `json:"items"`
}

//WebStorageTypesView - A collection of valid web storage types.
type WebStorageTypesView struct {
	Items []*ItemView `json:"items"`
}

//X509FileImportDocView - An X.509 certificate.
type X509FileImportDocView struct {
	Alias    *string `json:"alias"`
	FileData *string `json:"fileData"`
}
