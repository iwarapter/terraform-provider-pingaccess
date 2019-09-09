package pingaccess

import (
	"fmt"
	"net/http"
	"net/url"
)

type WebSessionManagementService service

//DeleteWebSessionManagementCommand - Resets the Web Session Management configuration to default values
//RequestType: DELETE
//Input:
func (s *WebSessionManagementService) DeleteWebSessionManagementCommand() (resp *http.Response, err error) {
	path := "/webSessionManagement"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("DELETE", rel, nil)
	if err != nil {
		return nil, err
	}

	resp, err = s.client.do(req, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil

}

//GetWebSessionManagementCommand - Get the Web Session Management configuration
//RequestType: GET
//Input:
func (s *WebSessionManagementService) GetWebSessionManagementCommand() (result *WebSessionManagementView, resp *http.Response, err error) {
	path := "/webSessionManagement"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("GET", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//UpdateWebSessionManagementCommand - Update the Web Session Management configuration
//RequestType: PUT
//Input: input *UpdateWebSessionManagementCommandInput
func (s *WebSessionManagementService) UpdateWebSessionManagementCommand(input *UpdateWebSessionManagementCommandInput) (result *WebSessionManagementView, resp *http.Response, err error) {
	path := "/webSessionManagement"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("PUT", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

type UpdateWebSessionManagementCommandInput struct {
	Body WebSessionManagementView
}

//GetCookieTypes - Get the valid OIDC Cookie Types
//RequestType: GET
//Input:
func (s *WebSessionManagementService) GetCookieTypes() (result *CookieTypesView, resp *http.Response, err error) {
	path := "/webSessionManagement/cookieTypes"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("GET", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//GetWebSessionSupportedEncryptionAlgorithmsCommand - Get the valid algorithms for Web Session Cookie Encryption
//RequestType: GET
//Input:
func (s *WebSessionManagementService) GetWebSessionSupportedEncryptionAlgorithmsCommand() (result *AlgorithmsView, resp *http.Response, err error) {
	path := "/webSessionManagement/encryptionAlgorithms"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("GET", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//GetWebSessionKeySetCommand - Get the Web Session key set configuration
//RequestType: GET
//Input:
func (s *WebSessionManagementService) GetWebSessionKeySetCommand() (result *KeySetView, resp *http.Response, err error) {
	path := "/webSessionManagement/keySet"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("GET", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//UpdateWebSessionKeySetCommand - Update the WebSession key set configuration
//RequestType: PUT
//Input: input *UpdateWebSessionKeySetCommandInput
func (s *WebSessionManagementService) UpdateWebSessionKeySetCommand(input *UpdateWebSessionKeySetCommandInput) (result *KeySetView, resp *http.Response, err error) {
	path := "/webSessionManagement/keySet"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("PUT", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

type UpdateWebSessionKeySetCommandInput struct {
	Body KeySetView
}

//GetOidcLoginTypes - Get the valid OIDC Login Types
//RequestType: GET
//Input:
func (s *WebSessionManagementService) GetOidcLoginTypes() (result *OidcLoginTypesView, resp *http.Response, err error) {
	path := "/webSessionManagement/oidcLoginTypes"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("GET", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//GetOidcScopesCommand - Get the scopes supported by the current OIDC Provider
//RequestType: GET
//Input: input *GetOidcScopesCommandInput
func (s *WebSessionManagementService) GetOidcScopesCommand(input *GetOidcScopesCommandInput) (result *SupportedScopesView, resp *http.Response, err error) {
	path := "/webSessionManagement/oidcScopes"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	q := rel.Query()
	if input.ClientId != "" {
		q.Set("clientId", input.ClientId)
	}
	rel.RawQuery = q.Encode()
	req, err := s.client.newRequest("GET", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

type GetOidcScopesCommandInput struct {
	ClientId string
}

//GetRequestPreservationTypes - Get the valid Request Preservation Types
//RequestType: GET
//Input:
func (s *WebSessionManagementService) GetRequestPreservationTypes() (result *RequestPreservationTypesView, resp *http.Response, err error) {
	path := "/webSessionManagement/requestPreservationTypes"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("GET", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//GetWebSessionSupportedSigningAlgorithms - Get the valid algorithms for Web Session Cookie Signing
//RequestType: GET
//Input:
func (s *WebSessionManagementService) GetWebSessionSupportedSigningAlgorithms() (result *SigningAlgorithmsView, resp *http.Response, err error) {
	path := "/webSessionManagement/signingAlgorithms"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("GET", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//GetWebStorageTypes - Get the valid Web Storage Types
//RequestType: GET
//Input:
func (s *WebSessionManagementService) GetWebStorageTypes() (result *WebStorageTypesView, resp *http.Response, err error) {
	path := "/webSessionManagement/webStorageTypes"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("GET", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}
