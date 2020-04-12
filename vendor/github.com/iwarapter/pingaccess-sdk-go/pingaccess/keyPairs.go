package pingaccess

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type KeyPairsService service

//GetKeyPairsCommand - Get all Key Pairs
//RequestType: GET
//Input: input *GetKeyPairsCommandInput
func (s *KeyPairsService) GetKeyPairsCommand(input *GetKeyPairsCommandInput) (result *KeyPairsView, resp *http.Response, err error) {
	path := "/keyPairs"
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

type GetKeyPairsCommandInput struct {
	Page          string
	NumberPerPage string
	Filter        string
	Alias         string
	SortKey       string
	Order         string
}

//GenerateKeyPairCommand - Generate a Key Pair
//RequestType: POST
//Input: input *GenerateKeyPairCommandInput
func (s *KeyPairsService) GenerateKeyPairCommand(input *GenerateKeyPairCommandInput) (result *KeyPairView, resp *http.Response, err error) {
	path := "/keyPairs/generate"
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

type GenerateKeyPairCommandInput struct {
	Body NewKeyPairConfigView
}

//ImportKeyPairCommand - Import a Key Pair from a PKCS12 file
//RequestType: POST
//Input: input *ImportKeyPairCommandInput
func (s *KeyPairsService) ImportKeyPairCommand(input *ImportKeyPairCommandInput) (result *KeyPairView, resp *http.Response, err error) {
	path := "/keyPairs/import"
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

type ImportKeyPairCommandInput struct {
	Body PKCS12FileImportDocView
}

//KeyAlgorithms - Get the key algorithms supported by Key Pair generation
//RequestType: GET
//Input:
func (s *KeyPairsService) KeyAlgorithms() (result *KeyAlgorithmsView, resp *http.Response, err error) {
	path := "/keyPairs/keyAlgorithms"
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

//GetKeypairsCreatableGeneralNamesCommand - Get the valid General Names for creating Subject Alternative Names
//RequestType: GET
//Input:
func (s *KeyPairsService) GetKeypairsCreatableGeneralNamesCommand() (result *SanTypes, resp *http.Response, err error) {
	path := "/keyPairs/subjectAlternativeTypes"
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

//DeleteKeyPairCommand - Delete a Key Pair
//RequestType: DELETE
//Input: input *DeleteKeyPairCommandInput
func (s *KeyPairsService) DeleteKeyPairCommand(input *DeleteKeyPairCommandInput) (resp *http.Response, err error) {
	path := "/keyPairs/{id}"
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

type DeleteKeyPairCommandInput struct {
	Id string
}

//GetKeyPairCommand - Get a Key Pair
//RequestType: GET
//Input: input *GetKeyPairCommandInput
func (s *KeyPairsService) GetKeyPairCommand(input *GetKeyPairCommandInput) (result *KeyPairView, resp *http.Response, err error) {
	path := "/keyPairs/{id}"
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

type GetKeyPairCommandInput struct {
	Id string
}

//PatchKeyPairCommand - Update the chainCertificates of a Key Pair
//RequestType: PATCH
//Input: input *PatchKeyPairCommandInput
func (s *KeyPairsService) PatchKeyPairCommand(input *PatchKeyPairCommandInput) (result *KeyPairView, resp *http.Response, err error) {
	path := "/keyPairs/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("PATCH", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

type PatchKeyPairCommandInput struct {
	Body ChainCertificatesDocView
	Id   string
}

//UpdateKeyPairCommand - Update a Key Pair
//RequestType: PUT
//Input: input *UpdateKeyPairCommandInput
func (s *KeyPairsService) UpdateKeyPairCommand(input *UpdateKeyPairCommandInput) (result *KeyPairView, resp *http.Response, err error) {
	path := "/keyPairs/{id}"
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

type UpdateKeyPairCommandInput struct {
	Body PKCS12FileImportDocView
	Id   string
}

//ExportKeyPairCert - Export only the Certificate from a Key Pair
//RequestType: GET
//Input: input *ExportKeyPairCertInput
func (s *KeyPairsService) ExportKeyPairCert(input *ExportKeyPairCertInput) (resp *http.Response, err error) {
	path := "/keyPairs/{id}/certificate"
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

type ExportKeyPairCertInput struct {
	Id string
}

//GenerateCsrCommand - Generate a Certificate Signing Request for a Key Pair
//RequestType: GET
//Input: input *GenerateCsrCommandInput
func (s *KeyPairsService) GenerateCsrCommand(input *GenerateCsrCommandInput) (resp *http.Response, err error) {
	path := "/keyPairs/{id}/csr"
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

type GenerateCsrCommandInput struct {
	Id string
}

//ImportCSRResponseCommand - Import a Certificate Signing Request response
//RequestType: POST
//Input: input *ImportCSRResponseCommandInput
func (s *KeyPairsService) ImportCSRResponseCommand(input *ImportCSRResponseCommandInput) (result *KeyPairView, resp *http.Response, err error) {
	path := "/keyPairs/{id}/csr"
	path = strings.Replace(path, "{id}", input.Id, -1)

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

type ImportCSRResponseCommandInput struct {
	Body CSRResponseImportDocView
	Id   string
}

//ExportKeyPair - Export a Key Pair in the PKCS12 file format
//RequestType: POST
//Input: input *ExportKeyPairInput
func (s *KeyPairsService) ExportKeyPair(input *ExportKeyPairInput) (resp *http.Response, err error) {
	path := "/keyPairs/{id}/pkcs12"
	path = strings.Replace(path, "{id}", input.Id, -1)

	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("POST", rel, input.Body)
	if err != nil {
		return nil, err
	}

	resp, err = s.client.do(req, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil

}

type ExportKeyPairInput struct {
	Body ExportParameters
	Id   string
}

//DeleteChainCertificateCommand - Delete a Chain Certificate
//RequestType: DELETE
//Input: input *DeleteChainCertificateCommandInput
func (s *KeyPairsService) DeleteChainCertificateCommand(input *DeleteChainCertificateCommandInput) (resp *http.Response, err error) {
	path := "/keyPairs/{keyPairId}/chainCertificates/{chainCertificateId}"
	path = strings.Replace(path, "{keyPairId}", input.KeyPairId, -1)

	path = strings.Replace(path, "{chainCertificateId}", input.ChainCertificateId, -1)

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

type DeleteChainCertificateCommandInput struct {
	KeyPairId          string
	ChainCertificateId string
}
