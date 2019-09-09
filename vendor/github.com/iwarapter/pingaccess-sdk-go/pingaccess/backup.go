package pingaccess

import (
	"fmt"
	"net/http"
	"net/url"
)

type BackupService service

//BackupCommand - Create a local database backup
//RequestType: GET
//Input:
func (s *BackupService) BackupCommand() (resp *http.Response, err error) {
	path := "/backup"
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
