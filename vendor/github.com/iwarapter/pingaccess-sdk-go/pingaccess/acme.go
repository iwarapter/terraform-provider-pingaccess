package pingaccess

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type AcmeService service

//GetAcmeServersCommand - Get all ACME Servers
//RequestType: GET
//Input: input *GetAcmeServersCommandInput
func (s *AcmeService) GetAcmeServersCommand(input *GetAcmeServersCommandInput) (result *AcmeServersView, resp *http.Response, err error) {
	path := "/acme/servers"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	q := rel.Query()
	if input.Page != "" {
		q.Set("page", input.Page)
	}
	if input.NumberPerPage != "" {
		q.Set("numberPerPage", input.NumberPerPage)
	}
	if input.Filter != "" {
		q.Set("filter", input.Filter)
	}
	if input.Name != "" {
		q.Set("name", input.Name)
	}
	if input.SortKey != "" {
		q.Set("sortKey", input.SortKey)
	}
	if input.Order != "" {
		q.Set("order", input.Order)
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

type GetAcmeServersCommandInput struct {
	Page          string
	NumberPerPage string
	Filter        string
	Name          string
	SortKey       string
	Order         string
}

//AddAcmeServerCommand - Add an ACME Server
//RequestType: POST
//Input: input *AddAcmeServerCommandInput
func (s *AcmeService) AddAcmeServerCommand(input *AddAcmeServerCommandInput) (result *AcmeServerView, resp *http.Response, err error) {
	path := "/acme/servers"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("POST", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

type AddAcmeServerCommandInput struct {
	Body AcmeServerView
}

//GetDefaultAcmeServerCommand - Get the default ACME Server
//RequestType: GET
//Input:
func (s *AcmeService) GetDefaultAcmeServerCommand() (result *LinkView, resp *http.Response, err error) {
	path := "/acme/servers/default"
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

//UpdateDefaultAcmeServerCommand - Update the default ACME Server
//RequestType: PUT
//Input: input *UpdateDefaultAcmeServerCommandInput
func (s *AcmeService) UpdateDefaultAcmeServerCommand(input *UpdateDefaultAcmeServerCommandInput) (result *LinkView, resp *http.Response, err error) {
	path := "/acme/servers/default"
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

type UpdateDefaultAcmeServerCommandInput struct {
	Body LinkView
}

//DeleteAcmeServerCommand - Delete an ACME Server
//RequestType: DELETE
//Input: input *DeleteAcmeServerCommandInput
func (s *AcmeService) DeleteAcmeServerCommand(input *DeleteAcmeServerCommandInput) (result *AcmeServerView, resp *http.Response, err error) {
	path := "/acme/servers/{acmeServerId}"
	path = strings.Replace(path, "{acmeServerId}", input.AcmeServerId, -1)

	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("DELETE", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

type DeleteAcmeServerCommandInput struct {
	AcmeServerId string
}

//GetAcmeServerCommand - Get an ACME Server
//RequestType: GET
//Input: input *GetAcmeServerCommandInput
func (s *AcmeService) GetAcmeServerCommand(input *GetAcmeServerCommandInput) (result *AcmeServerView, resp *http.Response, err error) {
	path := "/acme/servers/{acmeServerId}"
	path = strings.Replace(path, "{acmeServerId}", input.AcmeServerId, -1)

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

type GetAcmeServerCommandInput struct {
	AcmeServerId string
}

//GetAcmeAccountsCommand - Get all ACME Accounts
//RequestType: GET
//Input: input *GetAcmeAccountsCommandInput
func (s *AcmeService) GetAcmeAccountsCommand(input *GetAcmeAccountsCommandInput) (result *AcmeAccountView, resp *http.Response, err error) {
	path := "/acme/servers/{acmeServerId}/accounts"
	path = strings.Replace(path, "{acmeServerId}", input.AcmeServerId, -1)

	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	q := rel.Query()
	if input.Page != "" {
		q.Set("page", input.Page)
	}
	if input.NumberPerPage != "" {
		q.Set("numberPerPage", input.NumberPerPage)
	}
	if input.SortKey != "" {
		q.Set("sortKey", input.SortKey)
	}
	if input.Order != "" {
		q.Set("order", input.Order)
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

type GetAcmeAccountsCommandInput struct {
	Page          string
	NumberPerPage string
	SortKey       string
	Order         string

	AcmeServerId string
}

//AddAcmeAccountCommand - Add an ACME Account
//RequestType: POST
//Input: input *AddAcmeAccountCommandInput
func (s *AcmeService) AddAcmeAccountCommand(input *AddAcmeAccountCommandInput) (result *AcmeAccountView, resp *http.Response, err error) {
	path := "/acme/servers/{acmeServerId}/accounts"
	path = strings.Replace(path, "{acmeServerId}", input.AcmeServerId, -1)

	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("POST", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

type AddAcmeAccountCommandInput struct {
	Body         AcmeAccountView
	AcmeServerId string
}

//DeleteAcmeAccountCommand - Delete an ACME Account
//RequestType: DELETE
//Input: input *DeleteAcmeAccountCommandInput
func (s *AcmeService) DeleteAcmeAccountCommand(input *DeleteAcmeAccountCommandInput) (result *AcmeAccountView, resp *http.Response, err error) {
	path := "/acme/servers/{acmeServerId}/accounts/{acmeAccountId}"
	path = strings.Replace(path, "{acmeServerId}", input.AcmeServerId, -1)

	path = strings.Replace(path, "{acmeAccountId}", input.AcmeAccountId, -1)

	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("DELETE", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

type DeleteAcmeAccountCommandInput struct {
	AcmeServerId  string
	AcmeAccountId string
}

//GetAcmeAccountCommand - Get an ACME Account
//RequestType: GET
//Input: input *GetAcmeAccountCommandInput
func (s *AcmeService) GetAcmeAccountCommand(input *GetAcmeAccountCommandInput) (result *AcmeAccountView, resp *http.Response, err error) {
	path := "/acme/servers/{acmeServerId}/accounts/{acmeAccountId}"
	path = strings.Replace(path, "{acmeServerId}", input.AcmeServerId, -1)

	path = strings.Replace(path, "{acmeAccountId}", input.AcmeAccountId, -1)

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

type GetAcmeAccountCommandInput struct {
	AcmeServerId  string
	AcmeAccountId string
}

//GetAcmeCertificateRequestsCommand - Get all ACME Certificate Requests
//RequestType: GET
//Input: input *GetAcmeCertificateRequestsCommandInput
func (s *AcmeService) GetAcmeCertificateRequestsCommand(input *GetAcmeCertificateRequestsCommandInput) (result *AcmeCertificateRequestView, resp *http.Response, err error) {
	path := "/acme/servers/{acmeServerId}/accounts/{acmeAccountId}/certificateRequests"
	path = strings.Replace(path, "{acmeServerId}", input.AcmeServerId, -1)

	path = strings.Replace(path, "{acmeAccountId}", input.AcmeAccountId, -1)

	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	q := rel.Query()
	if input.KeyPairId != "" {
		q.Set("keyPairId", input.KeyPairId)
	}
	if input.Page != "" {
		q.Set("page", input.Page)
	}
	if input.NumberPerPage != "" {
		q.Set("numberPerPage", input.NumberPerPage)
	}
	if input.SortKey != "" {
		q.Set("sortKey", input.SortKey)
	}
	if input.Order != "" {
		q.Set("order", input.Order)
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

type GetAcmeCertificateRequestsCommandInput struct {
	KeyPairId     string
	Page          string
	NumberPerPage string
	SortKey       string
	Order         string

	AcmeServerId  string
	AcmeAccountId string
}

//AddAcmeCertificateRequestCommand - Initiate the ACME protocol
//RequestType: POST
//Input: input *AddAcmeCertificateRequestCommandInput
func (s *AcmeService) AddAcmeCertificateRequestCommand(input *AddAcmeCertificateRequestCommandInput) (result *AcmeCertificateRequestView, resp *http.Response, err error) {
	path := "/acme/servers/{acmeServerId}/accounts/{acmeAccountId}/certificateRequests"
	path = strings.Replace(path, "{acmeServerId}", input.AcmeServerId, -1)

	path = strings.Replace(path, "{acmeAccountId}", input.AcmeAccountId, -1)

	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("POST", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

type AddAcmeCertificateRequestCommandInput struct {
	Body          AcmeCertificateRequestView
	AcmeServerId  string
	AcmeAccountId string
}

//DeleteAcmeCertificateRequestCommand - Delete an ACME Certificate Request
//RequestType: DELETE
//Input: input *DeleteAcmeCertificateRequestCommandInput
func (s *AcmeService) DeleteAcmeCertificateRequestCommand(input *DeleteAcmeCertificateRequestCommandInput) (result *AcmeCertificateRequestView, resp *http.Response, err error) {
	path := "/acme/servers/{acmeServerId}/accounts/{acmeAccountId}/certificateRequests/{acmeCertificateRequestId}"
	path = strings.Replace(path, "{acmeServerId}", input.AcmeServerId, -1)

	path = strings.Replace(path, "{acmeAccountId}", input.AcmeAccountId, -1)

	path = strings.Replace(path, "{acmeCertificateRequestId}", input.AcmeCertificateRequestId, -1)

	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("DELETE", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

type DeleteAcmeCertificateRequestCommandInput struct {
	AcmeServerId             string
	AcmeAccountId            string
	AcmeCertificateRequestId string
}

//GetAcmeCertificateRequestCommand - Get an ACME Certificate Request
//RequestType: GET
//Input: input *GetAcmeCertificateRequestCommandInput
func (s *AcmeService) GetAcmeCertificateRequestCommand(input *GetAcmeCertificateRequestCommandInput) (result *AcmeCertificateRequestView, resp *http.Response, err error) {
	path := "/acme/servers/{acmeServerId}/accounts/{acmeAccountId}/certificateRequests/{acmeCertificateRequestId}"
	path = strings.Replace(path, "{acmeServerId}", input.AcmeServerId, -1)

	path = strings.Replace(path, "{acmeAccountId}", input.AcmeAccountId, -1)

	path = strings.Replace(path, "{acmeCertificateRequestId}", input.AcmeCertificateRequestId, -1)

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

type GetAcmeCertificateRequestCommandInput struct {
	AcmeServerId             string
	AcmeAccountId            string
	AcmeCertificateRequestId string
}
