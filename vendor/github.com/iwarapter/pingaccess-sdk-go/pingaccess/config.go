package pingaccess

import (
	"fmt"
	"net/http"
	"net/url"
)

type ConfigService service

//ConfigExportCommand - Export a JSON backup of the entire system
//RequestType: GET
//Input:
func (s *ConfigService) ConfigExportCommand() (result *ExportData, resp *http.Response, err error) {
	path := "/config/export"
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

//ConfigImportCommand - Import JSON backup
//RequestType: POST
//Input: input *ConfigImportCommandInput
func (s *ConfigService) ConfigImportCommand(input *ConfigImportCommandInput) (resp *http.Response, err error) {
	path := "/config/import"
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

type ConfigImportCommandInput struct {
	Body string
}
