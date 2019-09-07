package pingaccess

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type CertificatesService service

//GetTrustedCerts - Get all Certificates
//RequestType: GET
//Input: input *GetTrustedCertsInput
func (s *CertificatesService) GetTrustedCerts(input *GetTrustedCertsInput) (result *TrustedCertsView, resp *http.Response, err error) {
	path := "/certificates"
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
	if input.Alias != "" {
		q.Set("alias", input.Alias)
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

type GetTrustedCertsInput struct {
	Page          string
	NumberPerPage string
	Filter        string
	Alias         string
	SortKey       string
	Order         string
}

//ImportTrustedCert - Import a Certificate
//RequestType: POST
//Input: input *ImportTrustedCertInput
func (s *CertificatesService) ImportTrustedCert(input *ImportTrustedCertInput) (result *TrustedCertView, resp *http.Response, err error) {
	path := "/certificates"
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

type ImportTrustedCertInput struct {
	Body X509FileImportDocView
}

//DeleteTrustedCertCommand - Delete a Certificate
//RequestType: DELETE
//Input: input *DeleteTrustedCertCommandInput
func (s *CertificatesService) DeleteTrustedCertCommand(input *DeleteTrustedCertCommandInput) (resp *http.Response, err error) {
	path := "/certificates/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

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

type DeleteTrustedCertCommandInput struct {
	Id string
}

//GetTrustedCert - Get a Certificate
//RequestType: GET
//Input: input *GetTrustedCertInput
func (s *CertificatesService) GetTrustedCert(input *GetTrustedCertInput) (result *TrustedCertView, resp *http.Response, err error) {
	path := "/certificates/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

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

type GetTrustedCertInput struct {
	Id string
}

//UpdateTrustedCert - Update a Certificate
//RequestType: PUT
//Input: input *UpdateTrustedCertInput
func (s *CertificatesService) UpdateTrustedCert(input *UpdateTrustedCertInput) (result *TrustedCertView, resp *http.Response, err error) {
	path := "/certificates/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

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

type UpdateTrustedCertInput struct {
	Body X509FileImportDocView
	Id   string
}

//ExportTrustedCert - Export a Certificate
//RequestType: GET
//Input: input *ExportTrustedCertInput
func (s *CertificatesService) ExportTrustedCert(input *ExportTrustedCertInput) (resp *http.Response, err error) {
	path := "/certificates/{id}/file"
	path = strings.Replace(path, "{id}", input.Id, -1)

	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("GET", rel, nil)
	if err != nil {
		return nil, err
	}

	resp, err = s.client.do(req, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil

}

type ExportTrustedCertInput struct {
	Id string
}
