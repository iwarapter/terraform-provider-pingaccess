package pingaccess

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type ConfigService service

type ConfigAPI interface {
	ConfigExportCommand() (result *ExportData, resp *http.Response, err error)
	GetConfigExportWorkflowsCommand() (result *ConfigStatusesView, resp *http.Response, err error)
	AddConfigExportWorkflowCommand() (result *ConfigStatusView, resp *http.Response, err error)
	GetConfigExportWorkflowCommand(input *GetConfigExportWorkflowCommandInput) (result *ConfigStatusView, resp *http.Response, err error)
	GetConfigExportWorkflowDataCommand(input *GetConfigExportWorkflowDataCommandInput) (result *ExportData, resp *http.Response, err error)
	ConfigImportCommand(input *ConfigImportCommandInput) (resp *http.Response, err error)
	GetConfigImportWorkflowsCommand() (result *ConfigStatusesView, resp *http.Response, err error)
	AddConfigImportWorkflowCommand(input *AddConfigImportWorkflowCommandInput) (resp *http.Response, err error)
	GetConfigImportWorkflowCommand(input *GetConfigImportWorkflowCommandInput) (result *ConfigStatusView, resp *http.Response, err error)
}

//ConfigExportCommand - [Attention: The endpoint "/config/export" is deprecated. The "config/export/workflows" endpoint should be used instead.]    Export a JSON backup of the entire system
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

//GetConfigExportWorkflowsCommand - Get the status of pending Exports
//RequestType: GET
//Input:
func (s *ConfigService) GetConfigExportWorkflowsCommand() (result *ConfigStatusesView, resp *http.Response, err error) {
	path := "/config/export/workflows"
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

//AddConfigExportWorkflowCommand - Start a JSON backup of the entire system for export
//RequestType: POST
//Input:
func (s *ConfigService) AddConfigExportWorkflowCommand() (result *ConfigStatusView, resp *http.Response, err error) {
	path := "/config/export/workflows"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("POST", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//GetConfigExportWorkflowCommand - Check the status of an Export
//RequestType: GET
//Input: input *GetConfigExportWorkflowCommandInput
func (s *ConfigService) GetConfigExportWorkflowCommand(input *GetConfigExportWorkflowCommandInput) (result *ConfigStatusView, resp *http.Response, err error) {
	path := "/config/export/workflows/{id}"
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

type GetConfigExportWorkflowCommandInput struct {
	Id string
}

//GetConfigExportWorkflowDataCommand - Export a JSON backup of the entire system
//RequestType: GET
//Input: input *GetConfigExportWorkflowDataCommandInput
func (s *ConfigService) GetConfigExportWorkflowDataCommand(input *GetConfigExportWorkflowDataCommandInput) (result *ExportData, resp *http.Response, err error) {
	path := "/config/export/workflows/{id}/data"
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

type GetConfigExportWorkflowDataCommandInput struct {
	Id string
}

//ConfigImportCommand - [Attention: The endpoint "/config/import" is deprecated. The "config/import/workflows" endpoint should be used instead.]   Import JSON backup.
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

//GetConfigImportWorkflowsCommand - Get the status of pending imports
//RequestType: GET
//Input:
func (s *ConfigService) GetConfigImportWorkflowsCommand() (result *ConfigStatusesView, resp *http.Response, err error) {
	path := "/config/import/workflows"
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

//AddConfigImportWorkflowCommand - Start an Import of a JSON backup
//RequestType: POST
//Input: input *AddConfigImportWorkflowCommandInput
func (s *ConfigService) AddConfigImportWorkflowCommand(input *AddConfigImportWorkflowCommandInput) (resp *http.Response, err error) {
	path := "/config/import/workflows"
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

type AddConfigImportWorkflowCommandInput struct {
	Body ExportData
}

//GetConfigImportWorkflowCommand - Check the status of an Export
//RequestType: GET
//Input: input *GetConfigImportWorkflowCommandInput
func (s *ConfigService) GetConfigImportWorkflowCommand(input *GetConfigImportWorkflowCommandInput) (result *ConfigStatusView, resp *http.Response, err error) {
	path := "/config/import/workflows/{id}"
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

type GetConfigImportWorkflowCommandInput struct {
	Id string
}
