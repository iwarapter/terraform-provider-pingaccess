package pingaccess

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type HttpsListenersService service

//GetHttpsListenersCommand - Get all HTTPS Listeners
//RequestType: GET
//Input: input *GetHttpsListenersCommandInput
func (s *HttpsListenersService) GetHttpsListenersCommand(input *GetHttpsListenersCommandInput) (result *HttpsListenersView, resp *http.Response, err error) {
	path := "/httpsListeners"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	q := rel.Query()
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

type GetHttpsListenersCommandInput struct {
	SortKey string
	Order   string
}

//GetHttpsListenerCommand - Get an HTTPS Listener
//RequestType: GET
//Input: input *GetHttpsListenerCommandInput
func (s *HttpsListenersService) GetHttpsListenerCommand(input *GetHttpsListenerCommandInput) (result *HttpsListenerView, resp *http.Response, err error) {
	path := "/httpsListeners/{id}"
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

type GetHttpsListenerCommandInput struct {
	Id string
}

//UpdateHttpsListener - Update an HTTPS Listener
//RequestType: PUT
//Input: input *UpdateHttpsListenerInput
func (s *HttpsListenersService) UpdateHttpsListener(input *UpdateHttpsListenerInput) (result *HttpsListenerView, resp *http.Response, err error) {
	path := "/httpsListeners/{id}"
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

type UpdateHttpsListenerInput struct {
	Body HttpsListenerView
	Id   string
}
