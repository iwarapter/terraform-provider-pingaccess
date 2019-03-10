package pingaccess

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	pa "github.com/iwarapter/pingaccess-sdk-go/pingaccess"
	"github.com/ory/dockertest"
	"github.com/terraform-providers/terraform-provider-template/template"
)

func TestMain(m *testing.M) {
	_, acceptanceTesting := os.LookupEnv("TF_ACC")
	if acceptanceTesting {
		pool, err := dockertest.NewPool("")
		if err != nil {
			log.Fatalf("Could not connect to docker: %s", err)
		}

		dir, _ := os.Getwd()
		options := &dockertest.RunOptions{
			Repository: "pingidentity/pingaccess",
			Tag:        "latest",
			Mounts:     []string{dir + "/pingaccess.lic:/opt/in/instance/conf/pingaccess.lic"},
		}

		// pulls an image, creates a container based on it and runs it
		resource, err := pool.RunWithOptions(options)
		if err != nil {
			log.Fatalf("Could not start resource: %s", err)
		}

		// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
		if err := pool.Retry(func() error {
			var err error
			http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
			url, _ := url.Parse(fmt.Sprintf("https://localhost:%s", resource.GetPort("9000/tcp")))
			client := pa.NewClient("Administrator", "2Access", url, "/pa-admin-api/v3", nil)

			_, _, err = client.Applications.GetApplicationsCommand(&pa.GetApplicationsCommandInput{})
			return err
		}); err != nil {
			log.Fatalf("Could not connect to docker: %s", err)
		}

		os.Setenv("PINGACCESS_BASEURL", fmt.Sprintf("https://localhost:%s", resource.GetPort("9000/tcp")))
		code := m.Run()

		// You can't defer this because os.Exit doesn't care for defer
		if err := pool.Purge(resource); err != nil {
			log.Fatalf("Could not purge resource: %s", err)
		}

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
	testAccTemplateProvider = template.Provider().(*schema.Provider)
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
	err := testAccProvider.Configure(terraform.NewResourceConfig(nil))
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
