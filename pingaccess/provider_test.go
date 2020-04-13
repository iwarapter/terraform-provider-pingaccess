package pingaccess

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/iwarapter/pingaccess-sdk-go/pingaccess"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

func TestMain(m *testing.M) {
	_, acceptanceTesting := os.LookupEnv("TF_ACC")
	if acceptanceTesting {

		devOpsUser, devOpsUserExists := os.LookupEnv("PING_IDENTITY_DEVOPS_USER")
		devOpsKey, devOpsKeyExists := os.LookupEnv("PING_IDENTITY_DEVOPS_KEY")

		if devOpsKeyExists != true || devOpsUserExists != true {
			log.Fatalf("Both PING_IDENTITY_DEVOPS_USER and PING_IDENTITY_DEVOPS_KEY environment variables must be set for acceptance tests.")
		}

		ctx := context.Background()
		//networkName := "tf-pa-test-network"
		gcr := testcontainers.GenericContainerRequest{
			ContainerRequest: testcontainers.ContainerRequest{
				FromDockerfile: testcontainers.FromDockerfile{
					Context:    "../",
					Dockerfile: "Dockerfile",
				},
				//Networks:     []string{networkName},
				ExposedPorts: []string{"9000/tcp"},
				WaitingFor:   wait.ForLog("INFO: No file named /instance/data/data.json found, skipping import."),
				Name:         "terraform-provider-pingaccess-test",
				Env:          map[string]string{"PING_IDENTITY_ACCEPT_EULA": "YES", "PING_IDENTITY_DEVOPS_USER": devOpsUser, "PING_IDENTITY_DEVOPS_KEY": devOpsKey},
			},
			Started: true,
		}
		//provider, _ := gcr.ProviderType.GetProvider()
		//net, _ := provider.CreateNetwork(ctx, testcontainers.NetworkRequest{
		//	Name:           networkName,
		//	CheckDuplicate: true,
		//})

		paContainer, err := testcontainers.GenericContainer(ctx, gcr)
		if err != nil {
			log.Fatal(err)
		}
		//defer net.Remove(ctx)
		defer paContainer.Terminate(ctx)

		port, _ := paContainer.MappedPort(ctx, "9000")
		url, _ := url.Parse(fmt.Sprintf("https://localhost:%s", port.Port()))
		os.Setenv("PINGACCESS_BASEURL", url.String())
		os.Setenv("PINGACCESS_PASSWORD", "2FederateM0re")
		log.Printf("Setting PingAccess admin API: %s", url.String())
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
		client := pingaccess.NewClient("Administrator", "2FederateM0re", url, "/pa-admin-api/v3", nil)
		version, _, err := client.Version.VersionCommand()
		if err != nil {
			log.Fatalf("Failed to retrieve version from server: %v", err)
		}
		log.Printf("Connected to PingAccess version: %s", *version.Version)

		code := m.Run()
		log.Println("Tests complete shutting down container")

		os.Exit(code)
	} else {
		m.Run()
	}
}

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider
var testAccProviderFactories func(providers *[]*schema.Provider) map[string]terraform.ResourceProviderFactory
var testAccTemplateProvider *schema.Provider

func init() {
	testAccProvider = Provider().(*schema.Provider)
	//testAccTemplateProvider = template.Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"pingaccess": testAccProvider,
		"template":   testAccTemplateProvider,
	}
	testAccProviderFactories = func(providers *[]*schema.Provider) map[string]terraform.ResourceProviderFactory {
		return map[string]terraform.ResourceProviderFactory{
			"pingaccess": func() (terraform.ResourceProvider, error) {
				p := Provider()
				*providers = append(*providers, p.(*schema.Provider))
				return p, nil
			},
		}
	}
}

func testAccPreCheck(t *testing.T) {
	err := testAccProvider.Configure(terraform.NewResourceConfigRaw(nil))
	if err != nil {
		t.Fatal(err)
	}
}

// assert fails the test if the condition is false.
func assert(tb testing.TB, condition bool, msg string, v ...interface{}) {
	if !condition {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: "+msg+"\033[39m\n\n", append([]interface{}{filepath.Base(file), line}, v...)...)
		tb.FailNow()
	}
}

// ok fails the test if an err is not nil.
func ok(tb testing.TB, err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: unexpected error: %s\033[39m\n\n", filepath.Base(file), line, err.Error())
		tb.FailNow()
	}
}

// equals fails the test if exp is not equal to act.
func equals(tb testing.TB, exp, act interface{}) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\033[39m\n\n", filepath.Base(file), line, exp, act)
		tb.FailNow()
	}
}
