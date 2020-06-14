package pingaccess

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type AdminConfigService service

type AdminConfigAPI interface {
	DeleteAdminConfigurationCommand() (resp *http.Response, err error)
	GetAdminConfigurationCommand() (result *AdminConfigurationView, resp *http.Response, err error)
	UpdateAdminConfigurationCommand(input *UpdateAdminConfigurationCommandInput) (result *AdminConfigurationView, resp *http.Response, err error)
	GetReplicaAdminsCommand() (result *ReplicaAdminsView, resp *http.Response, err error)
	AddReplicaAdminCommand(input *AddReplicaAdminCommandInput) (result *ReplicaAdminView, resp *http.Response, err error)
	DeleteReplicaAdminCommand(input *DeleteReplicaAdminCommandInput) (resp *http.Response, err error)
	GetReplicaAdminCommand(input *GetReplicaAdminCommandInput) (result *ReplicaAdminView, resp *http.Response, err error)
	UpdateAdminReplicaCommand(input *UpdateAdminReplicaCommandInput) (result *ReplicaAdminView, resp *http.Response, err error)
	GetAdminReplicaFileCommand(input *GetAdminReplicaFileCommandInput) (resp *http.Response, err error)
}

//DeleteAdminConfigurationCommand - Resets the Admin Config to default values
//RequestType: DELETE
//Input:
func (s *AdminConfigService) DeleteAdminConfigurationCommand() (resp *http.Response, err error) {
	path := "/adminConfig"
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

//GetAdminConfigurationCommand - Get the Admin Config
//RequestType: GET
//Input:
func (s *AdminConfigService) GetAdminConfigurationCommand() (result *AdminConfigurationView, resp *http.Response, err error) {
	path := "/adminConfig"
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

//UpdateAdminConfigurationCommand - Update the Admin Config
//RequestType: PUT
//Input: input *UpdateAdminConfigurationCommandInput
func (s *AdminConfigService) UpdateAdminConfigurationCommand(input *UpdateAdminConfigurationCommandInput) (result *AdminConfigurationView, resp *http.Response, err error) {
	path := "/adminConfig"
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

type UpdateAdminConfigurationCommandInput struct {
	Body AdminConfigurationView
}

//GetReplicaAdminsCommand - Get list of ReplicaAdmins
//RequestType: GET
//Input:
func (s *AdminConfigService) GetReplicaAdminsCommand() (result *ReplicaAdminsView, resp *http.Response, err error) {
	path := "/adminConfig/replicaAdmins"
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

//AddReplicaAdminCommand - Add a ReplicaAdmin
//RequestType: POST
//Input: input *AddReplicaAdminCommandInput
func (s *AdminConfigService) AddReplicaAdminCommand(input *AddReplicaAdminCommandInput) (result *ReplicaAdminView, resp *http.Response, err error) {
	path := "/adminConfig/replicaAdmins"
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

type AddReplicaAdminCommandInput struct {
	Body ReplicaAdminView
}

//DeleteReplicaAdminCommand - Delete a ReplicaAdmin
//RequestType: DELETE
//Input: input *DeleteReplicaAdminCommandInput
func (s *AdminConfigService) DeleteReplicaAdminCommand(input *DeleteReplicaAdminCommandInput) (resp *http.Response, err error) {
	path := "/adminConfig/replicaAdmins/{id}"
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

type DeleteReplicaAdminCommandInput struct {
	Id string
}

//GetReplicaAdminCommand - Get a ReplicaAdmin
//RequestType: GET
//Input: input *GetReplicaAdminCommandInput
func (s *AdminConfigService) GetReplicaAdminCommand(input *GetReplicaAdminCommandInput) (result *ReplicaAdminView, resp *http.Response, err error) {
	path := "/adminConfig/replicaAdmins/{id}"
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

type GetReplicaAdminCommandInput struct {
	Id string
}

//UpdateAdminReplicaCommand - Update a ReplicaAdmin
//RequestType: PUT
//Input: input *UpdateAdminReplicaCommandInput
func (s *AdminConfigService) UpdateAdminReplicaCommand(input *UpdateAdminReplicaCommandInput) (result *ReplicaAdminView, resp *http.Response, err error) {
	path := "/adminConfig/replicaAdmins/{id}"
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

type UpdateAdminReplicaCommandInput struct {
	Body ReplicaAdminView
	Id   string
}

//GetAdminReplicaFileCommand - Get configuration file for a given ReplicaAdmin
//RequestType: POST
//Input: input *GetAdminReplicaFileCommandInput
func (s *AdminConfigService) GetAdminReplicaFileCommand(input *GetAdminReplicaFileCommandInput) (resp *http.Response, err error) {
	path := "/adminConfig/replicaAdmins/{id}/config"
	path = strings.Replace(path, "{id}", input.Id, -1)

	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("POST", rel, nil)
	if err != nil {
		return nil, err
	}

	resp, err = s.client.do(req, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil

}

type GetAdminReplicaFileCommandInput struct {
	Id string
}
