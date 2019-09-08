package pingaccess

import (
	"fmt"
	"net/http"
	"net/url"
)

type VersionService service

//VersionCommand - Get the PingAccess version number
//RequestType: GET
//Input:
func (s *VersionService) VersionCommand() (result *VersionDocClass, resp *http.Response, err error) {
	path := "/version"
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
