package sdkv2provider

import (
	"fmt"
	"net"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
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
		want     diag.Diagnostics
	}{
		{
			name:     "handle malformed urls",
			username: "foo",
			password: "bar",
			baseUrl:  "not a url",
			want: diag.Diagnostics{
				diag.Diagnostic{
					Severity: diag.Error,
					Summary:  "Invalid URL",
					Detail:   "Unable to parse base_url for client: parse \"not a url\": invalid URI for request",
				},
			},
		},
		{
			name:     "handle unresponsive server",
			username: "foo",
			password: "bar",
			baseUrl:  "https://localhost:19999",
			want: diag.Diagnostics{
				diag.Diagnostic{
					Severity: diag.Error,
					Summary:  "Connection Error",
					Detail:   "Unable to connect to PingAccess: Unknown host/port",
				},
			},
		},
		{
			name:     "unauthenticated",
			username: "foo",
			password: "bar",
			baseUrl:  server.URL,
			want: diag.Diagnostics{
				diag.Diagnostic{
					Severity: diag.Error,
					Summary:  "Connection Error",
					Detail:   "Unable to connect to PingAccess: unauthorized",
				},
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

func TestIs60OrAbove(t *testing.T) {
	cli := paClient{}
	tests := []struct {
		version string
		expect  bool
	}{
		{"6.0", true},
		{"6.1", true},
		{"6.2", true},
		{"6.3", true},
		{"7.0", true},
		{"7.1", true},
		{"7.2", true},
		{"7.3", true},
		{"8.0", true},
		{"8.1", true},
		{"8.2", true},
		{"8.3", true},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("we handle %s", tt.version), func(t *testing.T) {
			cli.apiVersion = tt.version
			assert.Equal(t, tt.expect, cli.Is60OrAbove())
		})
	}
}

func TestIs61OrAbove(t *testing.T) {
	cli := paClient{}
	tests := []struct {
		version string
		expect  bool
	}{
		{"6.0", false},
		{"6.1", true},
		{"6.2", true},
		{"6.3", true},
		{"7.0", true},
		{"7.1", true},
		{"7.2", true},
		{"7.3", true},
		{"8.0", true},
		{"8.1", true},
		{"8.2", true},
		{"8.3", true},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("we handle %s", tt.version), func(t *testing.T) {
			cli.apiVersion = tt.version
			assert.Equal(t, tt.expect, cli.Is61OrAbove())
		})
	}
}

func TestIs62OrAbove(t *testing.T) {
	cli := paClient{}
	tests := []struct {
		version string
		expect  bool
	}{
		{"6.0", false},
		{"6.1", false},
		{"6.2", true},
		{"6.3", true},
		{"7.0", true},
		{"7.1", true},
		{"7.2", true},
		{"7.3", true},
		{"8.0", true},
		{"8.1", true},
		{"8.2", true},
		{"8.3", true},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("we handle %s", tt.version), func(t *testing.T) {
			cli.apiVersion = tt.version
			assert.Equal(t, tt.expect, cli.Is62OrAbove())
		})
	}
}

func Test_CanMaskPasswords(t *testing.T) {
	cli := paClient{}
	tests := []struct {
		version string
		expect  bool
	}{
		{"6.0", false},
		{"6.1", true},
		{"6.2", true},
		{"6.3", true},
		{"7.0", true},
		{"7.1", true},
		{"7.2", true},
		{"7.3", true},
		{"8.0", true},
		{"8.1", true},
		{"8.2", true},
		{"8.3", true},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("we handle %s", tt.version), func(t *testing.T) {
			cli.apiVersion = tt.version
			assert.Equal(t, tt.expect, cli.CanMaskPasswords())
		})
	}
}
