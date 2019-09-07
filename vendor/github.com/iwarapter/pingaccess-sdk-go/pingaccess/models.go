package pingaccess

import (
	"encoding/json"
	"time"
)

//AccessTokenValidatorView
type AccessTokenValidatorView struct {
	ClassName     *string                `json:"className"`
	Configuration map[string]interface{} `json:"configuration"`
	Id            json.Number            `json:"id,omitempty"`
	Name          *string                `json:"name"`
}

//AccessTokenValidatorsView
type AccessTokenValidatorsView struct {
	Items []*AccessTokenValidatorView `json:"items"`
}

//AdminBasicWebSessionView
type AdminBasicWebSessionView struct {
	Audience                     *string `json:"audience"`
	CookieDomain                 *string `json:"cookieDomain,omitempty"`
	CookieType                   *string `json:"cookieType"`
	ExpirationWarningInMinutes   *int    `json:"expirationWarningInMinutes"`
	IdleTimeoutInMinutes         *int    `json:"idleTimeoutInMinutes"`
	SessionPollIntervalInSeconds *int    `json:"sessionPollIntervalInSeconds"`
	SessionTimeoutInMinutes      *int    `json:"sessionTimeoutInMinutes"`
}

//AdminConfigurationView
type AdminConfigurationView struct {
	HostPort     *string `json:"hostPort"`
	HttpProxyId  *int    `json:"httpProxyId,omitempty"`
	HttpsProxyId *int    `json:"httpsProxyId,omitempty"`
}

//AdminWebSessionOidcConfigurationView
type AdminWebSessionOidcConfigurationView struct {
	CacheUserAttributes           *bool                       `json:"cacheUserAttributes,omitempty"`
	ClientCredentials             *OAuthClientCredentialsView `json:"clientCredentials"`
	EnableRefreshUser             *bool                       `json:"enableRefreshUser,omitempty"`
	OidcLoginType                 *string                     `json:"oidcLoginType,omitempty"`
	PfsessionStateCacheInSeconds  *int                        `json:"pfsessionStateCacheInSeconds,omitempty"`
	RefreshUserInfoClaimsInterval *int                        `json:"refreshUserInfoClaimsInterval,omitempty"`
	RequestPreservationType       *string                     `json:"requestPreservationType,omitempty"`
	RequestProfile                *bool                       `json:"requestProfile,omitempty"`
	Scopes                        *[]*string                  `json:"scopes,omitempty"`
	SendRequestedUrlToProvider    *bool                       `json:"sendRequestedUrlToProvider,omitempty"`
	ValidateSessionIsAlive        *bool                       `json:"validateSessionIsAlive,omitempty"`
	WebStorageType                *string                     `json:"webStorageType,omitempty"`
}

//AgentCertificateView
type AgentCertificateView struct {
	Alias                   *string        `json:"alias"`
	ChainCertificate        *bool          `json:"chainCertificate"`
	Expires                 *string        `json:"expires"`
	Id                      json.Number    `json:"id,omitempty"`
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

//AgentCertificatesView
type AgentCertificatesView struct {
	Items []*AgentCertificateView `json:"items"`
}

//AgentView
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

//AgentsView
type AgentsView struct {
	Items []*AgentView `json:"items"`
}

//AlgorithmView
type AlgorithmView struct {
	Description *string `json:"description"`
	Name        *string `json:"name"`
}

//AlgorithmsView
type AlgorithmsView struct {
	Items []*AlgorithmView `json:"items"`
}

//ApplicationView - An application.
type ApplicationView struct {
	AccessValidatorId                     *int                      `json:"accessValidatorId,omitempty"`
	AgentCacheInvalidatedExpiration       *int                      `json:"agentCacheInvalidatedExpiration,omitempty"`
	AgentCacheInvalidatedResponseDuration *int                      `json:"agentCacheInvalidatedResponseDuration,omitempty"`
	AgentId                               *int                      `json:"agentId"`
	ApplicationType                       *string                   `json:"applicationType,omitempty"`
	CaseSensitivePath                     *bool                     `json:"caseSensitivePath,omitempty"`
	ContextRoot                           *string                   `json:"contextRoot"`
	DefaultAuthType                       *string                   `json:"defaultAuthType"`
	Description                           *string                   `json:"description,omitempty"`
	Destination                           *string                   `json:"destination,omitempty"`
	Enabled                               *bool                     `json:"enabled,omitempty"`
	Id                                    json.Number               `json:"id,omitempty"`
	IdentityMappingIds                    map[string]*int           `json:"identityMappingIds,omitempty"`
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

//AuthTokenManagementView
type AuthTokenManagementView struct {
	Issuer               *string `json:"issuer,omitempty"`
	KeyRollEnabled       *bool   `json:"keyRollEnabled,omitempty"`
	KeyRollPeriodInHours *int    `json:"keyRollPeriodInHours,omitempty"`
	SigningAlgorithm     *string `json:"signingAlgorithm,omitempty"`
}

//AuthnReqListView
type AuthnReqListView struct {
	AuthnReqs *[]*string  `json:"authnReqs"`
	Id        json.Number `json:"id,omitempty"`
	Name      *string     `json:"name"`
}

//AuthnReqListsView
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

//BasicAuthConfigView
type BasicAuthConfigView struct {
	Enabled *bool `json:"enabled,omitempty"`
}

//BasicConfig
type BasicConfig struct {
	Enabled *bool `json:"enabled"`
}

//CSRResponseImportDocView
type CSRResponseImportDocView struct {
	FileData           *string `json:"fileData"`
	TrustedCertGroupId *int    `json:"trustedCertGroupId"`
}

//ChainCertificateView
type ChainCertificateView struct {
	Alias                   *string        `json:"alias"`
	Expires                 json.Number    `json:"expires"`
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
	ValidFrom               json.Number    `json:"validFrom"`
}

//ChainCertificates
type ChainCertificates struct {
	AddChainCertificates *[]*string  `json:"addChainCertificates"`
	Id                   json.Number `json:"id,omitempty"`
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
	Options      *[]*ConfigurationOption   `json:"options"`
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

//CookieTypesView
type CookieTypesView struct {
	Items []*ItemView `json:"items"`
}

//DescriptorView
type DescriptorView struct {
	ClassName           *string               `json:"className"`
	ConfigurationFields []*ConfigurationField `json:"configurationFields"`
	Label               *string               `json:"label"`
	Type                *string               `json:"type"`
}

//DescriptorsView
type DescriptorsView struct {
	Items []*DescriptorView `json:"items"`
}

//EngineCertificateView
type EngineCertificateView struct {
	Alias                   *string        `json:"alias"`
	ChainCertificate        *bool          `json:"chainCertificate"`
	Expires                 *string        `json:"expires"`
	Id                      json.Number    `json:"id,omitempty"`
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

//EngineHealthStatus - Engine clustering health details.
type EngineHealthStatus struct {
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
	Id     json.Number `json:"id,omitempty"`
	Name   *string     `json:"name"`
	Port   *int        `json:"port"`
	Secure *bool       `json:"secure,omitempty"`
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

//ExportData
type ExportData struct {
	Data          map[string]interface{} `json:"data"`
	EncryptionKey *JsonWebKey            `json:"encryptionKey"`
	MasterKeys    *MasterKeysView        `json:"masterKeys"`
	Version       *string                `json:"version"`
}

//ExportParameters
type ExportParameters struct {
	Id       *int    `json:"id"`
	Password *string `json:"password"`
}

//GeneralName
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

//Hash
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

//HiddenFieldView
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

//HttpsListenerView - An HTTPS listener.
type HttpsListenerView struct {
	Id                        json.Number `json:"id,omitempty"`
	KeyPairId                 *int        `json:"keyPairId"`
	Name                      *string     `json:"name"`
	UseServerCipherSuiteOrder *bool       `json:"useServerCipherSuiteOrder"`
}

//HttpsListenersView - A collection of HTTPS listeners.
type HttpsListenersView struct {
	Items []*HttpsListenerView `json:"items"`
}

//IdentityMappingView
type IdentityMappingView struct {
	ClassName     *string                `json:"className"`
	Configuration map[string]interface{} `json:"configuration"`
	Id            json.Number            `json:"id,omitempty"`
	Name          *string                `json:"name"`
}

//IdentityMappingsView
type IdentityMappingsView struct {
	Items []*IdentityMappingView `json:"items"`
}

//IpMultiValueSourceView - Configuration for the IP source.
type IpMultiValueSourceView struct {
	FallbackToLastHopIp *bool      `json:"fallbackToLastHopIp,omitempty"`
	HeaderNameList      *[]*string `json:"headerNameList"`
	ListValueLocation   *string    `json:"listValueLocation"`
}

//ItemView
type ItemView struct {
	Description *string `json:"description"`
	Name        *string `json:"name"`
}

//JsonWebKey
type JsonWebKey struct {
	Algorithm *string        `json:"algorithm"`
	Key       *Key           `json:"key"`
	KeyId     *string        `json:"keyId"`
	KeyOps    *[]*string     `json:"keyOps"`
	KeyType   *string        `json:"keyType"`
	PublicKey *PublicKeyView `json:"publicKey"`
	Use       *string        `json:"use"`
}

//Key
type Key struct {
	Algorithm *string  `json:"algorithm"`
	Encoded   *[]*byte `json:"encoded"`
	Format    *string  `json:"format"`
}

//KeyAlgorithm
type KeyAlgorithm struct {
	DefaultKeySize            *int       `json:"defaultKeySize"`
	DefaultSignatureAlgorithm *string    `json:"defaultSignatureAlgorithm"`
	KeySizes                  *[]*int    `json:"keySizes"`
	Name                      *string    `json:"name"`
	SignatureAlgorithms       *[]*string `json:"signatureAlgorithms"`
}

//KeyAlgorithmsView
type KeyAlgorithmsView struct {
	Items []*KeyAlgorithm `json:"items"`
}

//KeyPairView
type KeyPairView struct {
	Alias                   *string                 `json:"alias"`
	ChainCertificates       []*ChainCertificateView `json:"chainCertificates,omitempty"`
	CsrPending              *bool                   `json:"csrPending"`
	Expires                 json.Number             `json:"expires"`
	Id                      json.Number             `json:"id,omitempty"`
	IssuerDn                *string                 `json:"issuerDn"`
	Md5sum                  *string                 `json:"md5sum"`
	SerialNumber            *string                 `json:"serialNumber"`
	Sha1sum                 *string                 `json:"sha1sum"`
	SignatureAlgorithm      *string                 `json:"signatureAlgorithm"`
	Status                  *string                 `json:"status"`
	SubjectAlternativeNames []*GeneralName          `json:"subjectAlternativeNames,omitempty"`
	SubjectCn               *string                 `json:"subjectCn,omitempty"`
	SubjectDn               *string                 `json:"subjectDn"`
	ValidFrom               json.Number             `json:"validFrom"`
}

//KeyPairsView
type KeyPairsView struct {
	Items []*KeyPairView `json:"items"`
}

//KeySetView
type KeySetView struct {
	KeySet *string `json:"keySet"`
	Nonce  *string `json:"nonce"`
}

//LicenseImportDocView
type LicenseImportDocView struct {
	FileData *string `json:"fileData"`
}

//LicenseView
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

//LinkView - A reference to the associated resource
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

//MasterKeysView
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

//NewKeyPairConfig
type NewKeyPairConfig struct {
	Alias                   *string        `json:"alias"`
	City                    *string        `json:"city"`
	CommonName              *string        `json:"commonName"`
	Country                 *string        `json:"country"`
	KeyAlgorithm            *string        `json:"keyAlgorithm"`
	KeySize                 *int           `json:"keySize"`
	Organization            *string        `json:"organization"`
	OrganizationUnit        *string        `json:"organizationUnit"`
	SignatureAlgorithm      *string        `json:"signatureAlgorithm,omitempty"`
	State                   *string        `json:"state"`
	SubjectAlternativeNames []*GeneralName `json:"subjectAlternativeNames,omitempty"`
	ValidDays               *int           `json:"validDays"`
}

//OAuthClientCredentialsView
type OAuthClientCredentialsView struct {
	ClientId     *string          `json:"clientId"`
	ClientSecret *HiddenFieldView `json:"clientSecret,omitempty"`
}

//OAuthConfigView
type OAuthConfigView struct {
	ClientId             *string                       `json:"clientId"`
	ClientSecret         *HiddenFieldView              `json:"clientSecret,omitempty"`
	Enabled              *bool                         `json:"enabled,omitempty"`
	RoleMapping          *RoleMappingConfigurationView `json:"roleMapping,omitempty"`
	Scope                *string                       `json:"scope"`
	SubjectAttributeName *string                       `json:"subjectAttributeName,omitempty"`
}

//OAuthKeyManagementView
type OAuthKeyManagementView struct {
	KeyRollEnabled       *bool `json:"keyRollEnabled,omitempty"`
	KeyRollPeriodInHours *int  `json:"keyRollPeriodInHours,omitempty"`
}

//OIDCProviderMetadata
type OIDCProviderMetadata struct {
	Authorization_endpoint                      *string    `json:"authorization_endpoint"`
	Claim_types_supported                       *[]*string `json:"claim_types_supported"`
	Claims_parameter_supported                  *bool      `json:"claims_parameter_supported"`
	Claims_supported                            *[]*string `json:"claims_supported"`
	Code_challenge_methods_supported            *[]*string `json:"code_challenge_methods_supported"`
	End_session_endpoint                        *string    `json:"end_session_endpoint"`
	Grant_types_supported                       *[]*string `json:"grant_types_supported"`
	Id_token_signing_alg_values_supported       *[]*string `json:"id_token_signing_alg_values_supported"`
	Introspection_endpoint                      *string    `json:"introspection_endpoint"`
	Issuer                                      *string    `json:"issuer"`
	Jwks_uri                                    *string    `json:"jwks_uri"`
	Ping_end_session_endpoint                   *string    `json:"ping_end_session_endpoint"`
	Ping_revoked_sris_endpoint                  *string    `json:"ping_revoked_sris_endpoint"`
	Request_object_signing_alg_values_supported *[]*string `json:"request_object_signing_alg_values_supported"`
	Request_parameter_supported                 *bool      `json:"request_parameter_supported"`
	Request_uri_parameter_supported             *bool      `json:"request_uri_parameter_supported"`
	Response_modes_supported                    *[]*string `json:"response_modes_supported"`
	Response_types_supported                    *[]*string `json:"response_types_supported"`
	Revocation_endpoint                         *string    `json:"revocation_endpoint"`
	Scopes_supported                            *[]*string `json:"scopes_supported"`
	Subject_types_supported                     *[]*string `json:"subject_types_supported"`
	Token_endpoint                              *string    `json:"token_endpoint"`
	Token_endpoint_auth_methods_supported       *[]*string `json:"token_endpoint_auth_methods_supported"`
	Userinfo_endpoint                           *string    `json:"userinfo_endpoint"`
	Userinfo_signing_alg_values_supported       *[]*string `json:"userinfo_signing_alg_values_supported"`
}

//OIDCProviderPluginView
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

//OidcConfigView
type OidcConfigView struct {
	AuthnReqListId    *int                                  `json:"authnReqListId,omitempty"`
	Enabled           *bool                                 `json:"enabled,omitempty"`
	OidcConfiguration *AdminWebSessionOidcConfigurationView `json:"oidcConfiguration"`
	RoleMapping       *RoleMappingConfigurationView         `json:"roleMapping,omitempty"`
	UseSlo            *bool                                 `json:"useSlo,omitempty"`
}

//OidcLoginTypesView
type OidcLoginTypesView struct {
	Items []*ItemView `json:"items"`
}

//OptionalAttributeMappingView - A set of user attributes that define an optional role mapping.
type OptionalAttributeMappingView struct {
	Attributes []*AttributeView `json:"attributes"`
	Enabled    *bool            `json:"enabled,omitempty"`
}

//PKCS12FileImportDocView
type PKCS12FileImportDocView struct {
	Alias             *string    `json:"alias"`
	ChainCertificates *[]*string `json:"chainCertificates"`
	FileData          *string    `json:"fileData"`
	Password          *string    `json:"password"`
}

//PathPatternView - A pattern for matching request URI paths.
type PathPatternView struct {
	Pattern *string `json:"pattern"`
	Type    *string `json:"type"`
}

//PingFederateAccessTokenView
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

//PingFederateAdminView
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

//PingFederateRuntimeView
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

//PingOne4CView - The Ping One for Customers OIDC provider configuration.
type PingOne4CView struct {
	Description               *string `json:"description,omitempty"`
	Issuer                    *string `json:"issuer"`
	TrustedCertificateGroupId *int    `json:"trustedCertificateGroupId,omitempty"`
	UseProxy                  *bool   `json:"useProxy,omitempty"`
}

//PolicyItem - The policy items associated with the application
type PolicyItem struct {
	Id   json.Number `json:"id,omitempty"`
	Type *string     `json:"type,omitempty"`
}

//ProtocolSourceView - Configuration for the protocol source.
type ProtocolSourceView struct {
	HeaderName *string `json:"headerName"`
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
	Id           json.Number         `json:"id,omitempty"`
	ResponseCode *int                `json:"responseCode,omitempty"`
	Source       *HostPortView       `json:"source,omitempty"`
	Target       *TargetHostPortView `json:"target,omitempty"`
}

//RedirectsView - A collection of Redirects.
type RedirectsView struct {
	Items []*RedirectView `json:"items"`
}

//RejectionHandlerView
type RejectionHandlerView struct {
	ClassName     *string                `json:"className"`
	Configuration map[string]interface{} `json:"configuration"`
	Id            json.Number            `json:"id,omitempty"`
	Name          *string                `json:"name"`
}

//RejectionHandlersView
type RejectionHandlersView struct {
	Items []*RejectionHandlerView `json:"items"`
}

//ReplicaAdminView
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

//ReplicaAdminsView
type ReplicaAdminsView struct {
	Items []*ReplicaAdminView `json:"items"`
}

//RequestPreservationTypesView
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
	Id      json.Number                  `json:"id,omitempty"`
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

//ResourcesView
type ResourcesView struct {
	Id    json.Number     `json:"id,omitempty"`
	Items []*ResourceView `json:"items"`
}

//RoleMappingConfigurationView - Configuration for mapping user attributes to roles.
type RoleMappingConfigurationView struct {
	Administrator *RequiredAttributeMappingView `json:"administrator,omitempty"`
	Auditor       *OptionalAttributeMappingView `json:"auditor,omitempty"`
	Enabled       *bool                         `json:"enabled,omitempty"`
}

//RuleDescriptorView
type RuleDescriptorView struct {
	AgentCachingDisabled *bool                 `json:"agentCachingDisabled"`
	Category             *string               `json:"category"`
	ClassName            *string               `json:"className"`
	ConfigurationFields  []*ConfigurationField `json:"configurationFields"`
	Label                *string               `json:"label"`
	Modes                *[]*string            `json:"modes"`
	Type                 *string               `json:"type"`
}

//RuleDescriptorsView
type RuleDescriptorsView struct {
	Items []*RuleDescriptorView `json:"items"`
}

//RuleSetElementTypesView
type RuleSetElementTypesView struct {
	Items []*ItemView `json:"items"`
}

//RuleSetSuccessCriteriaView
type RuleSetSuccessCriteriaView struct {
	Items []*ItemView `json:"items"`
}

//RuleSetView
type RuleSetView struct {
	ElementType     *string     `json:"elementType,omitempty"`
	Id              json.Number `json:"id,omitempty"`
	Name            *string     `json:"name"`
	Policy          *[]*int     `json:"policy"`
	SuccessCriteria *string     `json:"successCriteria,omitempty"`
}

//RuleSetsView
type RuleSetsView struct {
	Items []*RuleSetView `json:"items"`
}

//RuleView
type RuleView struct {
	ClassName             *string                `json:"className"`
	Configuration         map[string]interface{} `json:"configuration"`
	Id                    json.Number            `json:"id,omitempty"`
	Name                  *string                `json:"name"`
	SupportedDestinations *[]*string             `json:"supportedDestinations,omitempty"`
}

//RulesView
type RulesView struct {
	Items []*RuleView `json:"items"`
}

//SanType - The valid General Names for creating Subject Alternative Names
type SanType struct {
	Description *string `json:"description,omitempty"`
	Name        *string `json:"name,omitempty"`
}

//SanTypes
type SanTypes struct {
	Items *[]*SanType `json:"items"`
}

//SessionInfo
type SessionInfo struct {
	AccessControlDirectives *[]*string `json:"accessControlDirectives"`
	Exp                     *int       `json:"exp"`
	ExpWarn                 *int       `json:"expWarn"`
	Flash                   *string    `json:"flash"`
	Iat                     *int       `json:"iat"`
	MaxFileUploadSize       *int       `json:"maxFileUploadSize"`
	PollIntervalSeconds     *int       `json:"pollIntervalSeconds"`
	Roles                   *[]*string `json:"roles"`
	SesTimeout              *int       `json:"sesTimeout"`
	ShowWarning             *bool      `json:"showWarning"`
	SniEnabled              *bool      `json:"sniEnabled"`
	Sub                     *string    `json:"sub"`
	UseSlo                  *bool      `json:"useSlo"`
}

//SharedSecretView
type SharedSecretView struct {
	Created *string          `json:"created,omitempty"`
	Id      json.Number      `json:"id,omitempty"`
	Secret  *HiddenFieldView `json:"secret"`
}

//SharedSecretsView
type SharedSecretsView struct {
	Items []*SharedSecretView `json:"items"`
}

//SigningAlgorithmsView
type SigningAlgorithmsView struct {
	Items []*AlgorithmView `json:"items"`
}

//SiteAuthenticatorView
type SiteAuthenticatorView struct {
	ClassName     *string                `json:"className"`
	Configuration map[string]interface{} `json:"configuration"`
	Id            json.Number            `json:"id,omitempty"`
	Name          *string                `json:"name"`
}

//SiteAuthenticatorsView
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
	AvailabilityProfileId     *int        `json:"availabilityProfileId,omitempty"`
	ExpectedHostname          *string     `json:"expectedHostname,omitempty"`
	HostValue                 *string     `json:"hostValue,omitempty"`
	Id                        json.Number `json:"id,omitempty"`
	LoadBalancingStrategyId   *int        `json:"loadBalancingStrategyId,omitempty"`
	MaxConnections            *int        `json:"maxConnections,omitempty"`
	Name                      *string     `json:"name"`
	Secure                    *bool       `json:"secure,omitempty"`
	SkipHostnameVerification  *bool       `json:"skipHostnameVerification,omitempty"`
	Targets                   *[]*string  `json:"targets"`
	TrustedCertificateGroupId *int        `json:"trustedCertificateGroupId,omitempty"`
	UseProxy                  *bool       `json:"useProxy,omitempty"`
}

//ThirdPartyServicesView - A collection of third-party service items.
type ThirdPartyServicesView struct {
	Items []*ThirdPartyServiceView `json:"items"`
}

//TokenProviderSettingView
type TokenProviderSettingView struct {
	Type          *string `json:"type,omitempty"`
	UseThirdParty *bool   `json:"useThirdParty,omitempty"`
}

//TrustedCertView
type TrustedCertView struct {
	Alias                   *string        `json:"alias"`
	Expires                 json.Number    `json:"expires"`
	Id                      json.Number    `json:"id,omitempty"`
	IssuerDn                *string        `json:"issuerDn"`
	Md5sum                  *string        `json:"md5sum"`
	SerialNumber            *string        `json:"serialNumber"`
	Sha1sum                 *string        `json:"sha1sum"`
	SignatureAlgorithm      *string        `json:"signatureAlgorithm"`
	Status                  *string        `json:"status"`
	SubjectAlternativeNames []*GeneralName `json:"subjectAlternativeNames,omitempty"`
	SubjectCn               *string        `json:"subjectCn,omitempty"`
	SubjectDn               *string        `json:"subjectDn"`
	ValidFrom               json.Number    `json:"validFrom"`
}

//TrustedCertificateGroupView
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

//TrustedCertsView
type TrustedCertsView struct {
	Items []*TrustedCertView `json:"items"`
}

//UnknownResourceSettingsView
type UnknownResourceSettingsView struct {
	AgentDefaultCacheTTL *int    `json:"agentDefaultCacheTTL"`
	AgentDefaultMode     *string `json:"agentDefaultMode"`
	ErrorContentType     *string `json:"errorContentType"`
	ErrorStatusCode      *int    `json:"errorStatusCode"`
	ErrorTemplateFile    *string `json:"errorTemplateFile"`
}

//UserPasswordView
type UserPasswordView struct {
	CurrentPassword *string     `json:"currentPassword"`
	Id              json.Number `json:"id,omitempty"`
	NewPassword     *string     `json:"newPassword"`
}

//UserView
type UserView struct {
	Email        *string     `json:"email,omitempty"`
	FirstLogin   *bool       `json:"firstLogin,omitempty"`
	Id           json.Number `json:"id,omitempty"`
	ShowTutorial *bool       `json:"showTutorial,omitempty"`
	SlaAccepted  *bool       `json:"slaAccepted,omitempty"`
	Username     *string     `json:"username"`
}

//UsersView
type UsersView struct {
	Items []*UserView `json:"items"`
}

//VersionDocClass
type VersionDocClass struct {
	Version *string `json:"version"`
}

//VirtualHostView
type VirtualHostView struct {
	AgentResourceCacheTTL     *int        `json:"agentResourceCacheTTL,omitempty"`
	Host                      *string     `json:"host"`
	Id                        json.Number `json:"id,omitempty"`
	KeyPairId                 *int        `json:"keyPairId,omitempty"`
	Port                      *int        `json:"port"`
	TrustedCertificateGroupId *int        `json:"trustedCertificateGroupId,omitempty"`
}

//VirtualHostsView
type VirtualHostsView struct {
	Items []*VirtualHostView `json:"items"`
}

//WebSessionManagementView
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

//WebSessionView
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
	RefreshUserInfoClaimsInterval *int                        `json:"refreshUserInfoClaimsInterval,omitempty"`
	RequestPreservationType       *string                     `json:"requestPreservationType,omitempty"`
	RequestProfile                *bool                       `json:"requestProfile,omitempty"`
	Scopes                        *[]*string                  `json:"scopes,omitempty"`
	SecureCookie                  *bool                       `json:"secureCookie,omitempty"`
	SendRequestedUrlToProvider    *bool                       `json:"sendRequestedUrlToProvider,omitempty"`
	SessionTimeoutInMinutes       *int                        `json:"sessionTimeoutInMinutes,omitempty"`
	ValidateSessionIsAlive        *bool                       `json:"validateSessionIsAlive,omitempty"`
	WebStorageType                *string                     `json:"webStorageType,omitempty"`
}

//WebSessionsView
type WebSessionsView struct {
	Items []*WebSessionView `json:"items"`
}

//WebStorageTypesView
type WebStorageTypesView struct {
	Items []*ItemView `json:"items"`
}

//X509FileImportDocView
type X509FileImportDocView struct {
	Alias    *string `json:"alias"`
	FileData *string `json:"fileData"`
}
