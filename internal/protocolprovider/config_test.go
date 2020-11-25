package protocol

import (
	"net"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
)

func TestConfig_Client(t *testing.T) {
	server := httptest.NewUnstartedServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Send response to be tested
		rw.Header().Set("Content-Type", "application/json;charset=utf-8")
		rw.WriteHeader(http.StatusUnauthorized)
		//rw.Write([]byte(`{"resultId":"invalid_credentials","message":"The credentials you provided were not recognized."}`))
	}))
	l, _ := net.Listen("tcp", ":0")
	server.Listener = l //for CI tests as host.docker.internal is window/macosx
	server.StartTLS()
	// Close the server when test finishes
	defer server.Close()

	tests := []struct {
		name     string
		username string
		password string
		baseUrl  string
		want     *tfprotov5.Diagnostic
	}{
		{
			name:     "handle malformed urls",
			username: "foo",
			password: "bar",
			baseUrl:  "not a url",
			want: &tfprotov5.Diagnostic{
				Severity: tfprotov5.DiagnosticSeverityError,
				Summary:  "Invalid URL",
				Detail:   "Unable to parse base_url for client: parse \"not a url\": invalid URI for request",
			},
		},
		{
			name:     "handle unresponsive server",
			username: "foo",
			password: "bar",
			baseUrl:  "https://localhost:19999",
			want: &tfprotov5.Diagnostic{
				Severity: tfprotov5.DiagnosticSeverityError,
				Summary:  "Connection Error",
				Detail:   "Unable to connect to PingAccess: Unknown host/port",
			},
		},
		{
			name:     "unauthenticated",
			username: "foo",
			password: "bar",
			baseUrl:  server.URL,
			want: &tfprotov5.Diagnostic{
				Severity: tfprotov5.DiagnosticSeverityError,
				Summary:  "Connection Error",
				Detail:   "Unable to connect to PingAccess: unauthorized",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &cfg{
				Username: tt.username,
				Password: tt.password,
				BaseURL:  tt.baseUrl,
			}
			_, diags := c.Client()
			if !reflect.DeepEqual(diags, tt.want) {
				t.Errorf("Client() diags = %v, want %v", diags, tt.want)
			}
		})
	}
}
