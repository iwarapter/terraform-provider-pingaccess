package pingaccess

import (
	"fmt"
	"net/http"
	"net/url"
)

type LicenseService service

type LicenseAPI interface {
	GetLicenseCommand() (result *LicenseView, resp *http.Response, err error)
	ImportLicenseCommand(input *ImportLicenseCommandInput) (result *LicenseView, resp *http.Response, err error)
}

//GetLicenseCommand - Get the License File
//RequestType: GET
//Input:
func (s *LicenseService) GetLicenseCommand() (result *LicenseView, resp *http.Response, err error) {
	path := "/license"
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

//ImportLicenseCommand - Import a License
//RequestType: POST
//Input: input *ImportLicenseCommandInput
func (s *LicenseService) ImportLicenseCommand(input *ImportLicenseCommandInput) (result *LicenseView, resp *http.Response, err error) {
	path := "/license"
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

type ImportLicenseCommandInput struct {
	Body LicenseImportDocView
}
