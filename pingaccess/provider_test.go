package pingaccess

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/iwarapter/pingaccess-sdk-go/pingaccess"
	"github.com/ory/dockertest/v3"
)

func TestMain(m *testing.M) {
	_, acceptanceTesting := os.LookupEnv("TF_ACC")
	if acceptanceTesting {
		pool, err := dockertest.NewPool("")
		if err != nil {
			log.Fatalf("Could not connect to docker: %s", err)
		}
		server := httptest.NewUnstartedServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			// Send response to be tested
			rw.Header().Set("Content-Type", "application/json;charset=utf-8")
			rw.Write([]byte(`
{
  "issuer": "https://localhost:9031",
  "authorization_endpoint": "https://localhost:9031/as/authorization.oauth2",
  "token_endpoint": "https://localhost:9031/as/token.oauth2",
  "revocation_endpoint": "https://localhost:9031/as/revoke_token.oauth2",
  "userinfo_endpoint": "https://localhost:9031/idp/userinfo.openid",
  "introspection_endpoint": "https://localhost:9031/as/introspect.oauth2",
  "jwks_uri": "https://localhost:9031/pf/JWKS",
  "registration_endpoint": "https://localhost:9031/as/clients.oauth2",
  "ping_revoked_sris_endpoint": "https://localhost:9031/pf-ws/rest/sessionMgmt/revokedSris",
  "ping_end_session_endpoint": "https://localhost:9031/idp/startSLO.ping",
  "device_authorization_endpoint": "https://localhost:9031/as/device_authz.oauth2",
  "scopes_supported": [ "address", "mail", "phone", "openid", "profile", "group1" ],
  "response_types_supported": [ "code", "token", "id_token", "code token", "code id_token", "token id_token", "code token id_token" ],
  "response_modes_supported": [ "fragment", "query", "form_post" ],
  "grant_types_supported": [ "implicit", "authorization_code", "refresh_token", "password", "client_credentials", "urn:pingidentity.com:oauth2:grant_type:validate_bearer", "urn:ietf:params:oauth:grant-type:jwt-bearer", "urn:ietf:params:oauth:grant-type:saml2-bearer", "urn:ietf:params:oauth:grant-type:device_code", "urn:openid:params:grant-type:ciba" ],
  "subject_types_supported": [ "public", "pairwise" ],
  "id_token_signing_alg_values_supported": [ "none", "HS256", "HS384", "HS512", "RS256", "RS384", "RS512", "ES256", "ES384", "ES512", "PS256", "PS384", "PS512" ],
  "token_endpoint_auth_methods_supported": [ "client_secret_basic", "client_secret_post", "private_key_jwt" ],
  "token_endpoint_auth_signing_alg_values_supported":  [ "RS256", "RS384", "RS512", "ES256", "ES384", "ES512", "PS256", "PS384", "PS512" ],
  "claim_types_supported": [ "normal" ],
  "claims_parameter_supported": false,
  "request_parameter_supported": true,
  "request_uri_parameter_supported": false,
  "request_object_signing_alg_values_supported": [ "RS256", "RS384", "RS512", "ES256", "ES384", "ES512", "PS256", "PS384", "PS512" ],
  "id_token_encryption_alg_values_supported": [ "dir", "A128KW", "A192KW", "A256KW", "A128GCMKW", "A192GCMKW", "A256GCMKW", "ECDH-ES", "ECDH-ES+A128KW", "ECDH-ES+A192KW", "ECDH-ES+A256KW", "RSA-OAEP" ],
  "id_token_encryption_enc_values_supported": [ "A128CBC-HS256", "A192CBC-HS384", "A256CBC-HS512", "A128GCM", "A192GCM", "A256GCM" ],
  "backchannel_authentication_endpoint": "https://localhost:9031/as/bc-auth.ciba",
  "backchannel_token_delivery_modes_supported": [ "poll", "ping" ],
  "backchannel_authentication_request_signing_alg_values_supported": [ "RS256", "RS384", "RS512", "ES256", "ES384", "ES512", "PS256", "PS384", "PS512" ],
  "backchannel_user_code_parameter_supported": false
}
`))
		}))
		l, err := net.Listen("tcp", ":0")
		server.Listener = l //for CI tests as host.docker.internal is window/macosx
		server.StartTLS()
		// Close the server when test finishes
		defer server.Close()

		devOpsUser, devOpsUserExists := os.LookupEnv("PING_IDENTITY_DEVOPS_USER")
		devOpsKey, devOpsKeyExists := os.LookupEnv("PING_IDENTITY_DEVOPS_KEY")

		if devOpsKeyExists != true || devOpsUserExists != true {
			log.Fatalf("Both PING_IDENTITY_DEVOPS_USER and PING_IDENTITY_DEVOPS_KEY environment variables must be set for acceptance tests.")
		}

		randomID := randomString(10)
		paOpts := &dockertest.RunOptions{
			Name:       fmt.Sprintf("pa-%s", randomID),
			Repository: "pingidentity/pingaccess",
			Tag:        "6.0.2-edge",
			//ExtraHosts: []string{"host.docker.internal:host-gateway"},
			Env: []string{"PING_IDENTITY_ACCEPT_EULA=YES", fmt.Sprintf("PING_IDENTITY_DEVOPS_USER=%s", devOpsUser), fmt.Sprintf("PING_IDENTITY_DEVOPS_KEY=%s", devOpsKey)},
		}
		paCont, err := pool.RunWithOptions(paOpts)
		if err != nil {
			log.Fatalf("Could not create pingaccess container: %s", err)
		}
		defer paCont.Close()

		pool.MaxWait = time.Minute * 2

		// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
		u, _ := url.Parse(fmt.Sprintf("https://localhost:%s", paCont.GetPort("9000/tcp")))
		log.Printf("Setting PingAccess admin API: %s", u.String())
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
		client := pingaccess.NewClient("administrator", "2FederateM0re", u, "/pa-admin-api/v3", nil)
		if err = pool.Retry(func() error {
			log.Println("Attempting to connect to PingAccess admin API....")
			_, _, err = client.Version.VersionCommand()
			return err
		}); err != nil {
			log.Fatalf("Could not connect to pingaccess: %s", err)
		}
		os.Setenv("PINGACCESS_BASEURL", fmt.Sprintf("https://localhost:%s", paCont.GetPort("9000/tcp")))
		os.Setenv("PINGACCESS_PASSWORD", "2FederateM0re")
		host , _ := os.Hostname() //for CI tests as host.docker.internal is window/macosx
		os.Setenv("PINGFEDERATE_TEST_IP", strings.Replace(server.URL, "[::]", host, -1))
		log.Println("Connected to PingAccess admin API....")

		version, _, err := client.Version.VersionCommand()
		if err != nil {
			log.Fatalf("Failed to retrieve version from server: %v", err)
		}
		log.Printf("Connected to PingAccess version: %s", *version.Version)

		paCont.Expire(360)
		code := m.Run()
		paCont.Close()
		log.Printf("Tests complete shutting down container")

		os.Exit(code)
	} else {
		m.Run()
	}
}

var testAccProviders map[string]*schema.Provider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider()
	testAccProviders = map[string]*schema.Provider{
		"pingaccess": testAccProvider,
	}
}

func testAccPreCheck(t *testing.T) {
	err := testAccProvider.Configure(context.TODO(), terraform.NewResourceConfigRaw(nil))
	if err != nil {
		t.Fatal(err)
	}
}

func randomString(length int) string {
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	charset := "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

// equals fails the test if exp is not equal to act.
func equals(tb testing.TB, exp, act interface{}) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\033[39m\n\n", filepath.Base(file), line, exp, act)
		tb.FailNow()
	}
}
