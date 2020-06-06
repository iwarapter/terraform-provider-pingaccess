package pingaccess

import (
	"bytes"
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/iwarapter/pingaccess-sdk-go/pingaccess"
	"github.com/ory/dockertest/v3"
	"github.com/tidwall/sjson"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	acctest.UseBinaryDriver("pingaccess", Provider)
	_, acceptanceTesting := os.LookupEnv("TF_ACC")
	if acceptanceTesting {
		pool, err := dockertest.NewPool("")
		if err != nil {
			log.Fatalf("Could not connect to docker: %s", err)
		}
		networkName := "tf-pa-test-network"
		network, err := pool.CreateNetwork(networkName)
		if err != nil {
			log.Fatalf("Could not create docker network: %s", err)
		}
		defer network.Close()

		devOpsUser, devOpsUserExists := os.LookupEnv("PING_IDENTITY_DEVOPS_USER")
		devOpsKey, devOpsKeyExists := os.LookupEnv("PING_IDENTITY_DEVOPS_KEY")

		if devOpsKeyExists != true || devOpsUserExists != true {
			log.Fatalf("Both PING_IDENTITY_DEVOPS_USER and PING_IDENTITY_DEVOPS_KEY environment variables must be set for acceptance tests.")
		}

		randomID := randomString(10)
		paOpts := &dockertest.RunOptions{
			Name:       fmt.Sprintf("pa-%s", randomID),
			Repository: "pingidentity/pingaccess",
			Tag:        "6.0.1-edge",
			NetworkID:  network.Network.ID,
			Env:        []string{"PING_IDENTITY_ACCEPT_EULA=YES", fmt.Sprintf("PING_IDENTITY_DEVOPS_USER=%s", devOpsUser), fmt.Sprintf("PING_IDENTITY_DEVOPS_KEY=%s", devOpsKey)},
		}
		pfOpts := &dockertest.RunOptions{
			Name:       fmt.Sprintf("pf-%s", randomID),
			Repository: "pingidentity/pingfederate",
			Tag:        "10.0.2-edge",
			NetworkID:  network.Network.ID,
			Env: []string{
				"PING_IDENTITY_ACCEPT_EULA=YES",
				fmt.Sprintf("PING_IDENTITY_DEVOPS_USER=%s", devOpsUser),
				fmt.Sprintf("PING_IDENTITY_DEVOPS_KEY=%s", devOpsKey),
				"SERVER_PROFILE_URL=https://github.com/pingidentity/pingidentity-server-profiles.git",
				"SERVER_PROFILE_PATH=getting-started/pingfederate",
			},
		}
		// pulls an image, creates a container based on it and runs it
		paCont, err := pool.RunWithOptions(paOpts)
		if err != nil {
			log.Fatalf("Could not create pingaccess container: %s", err)
		}
		defer paCont.Close()
		pfCont, err := pool.RunWithOptions(pfOpts)
		if err != nil {
			log.Fatalf("Could not create pingfederate container: %s", err)
		}
		defer pfCont.Close()

		pool.MaxWait = time.Minute * 2

		// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
		url, _ := url.Parse(fmt.Sprintf("https://localhost:%s", paCont.GetPort("9000/tcp")))
		log.Printf("Setting PingAccess admin API: %s", url.String())
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
		client := pingaccess.NewClient("administrator", "2FederateM0re", url, "/pa-admin-api/v3", nil)
		if err = pool.Retry(func() error {
			log.Println("Attempting to connect to PingAccess admin API....")
			_, _, err = client.Version.VersionCommand()
			return err
		}); err != nil {
			log.Fatalf("Could not connect to pingaccess: %s", err)
		}
		os.Setenv("PINGACCESS_BASEURL", fmt.Sprintf("https://localhost:%s", paCont.GetPort("9000/tcp")))
		os.Setenv("PINGACCESS_PASSWORD", "2FederateM0re")
		os.Setenv("PINGFEDERATE_TEST_IP", pfCont.GetIPInNetwork(network))
		log.Println("Connected to PingAccess admin API....")

		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
		version, _, err := client.Version.VersionCommand()
		if err != nil {
			log.Fatalf("Failed to retrieve version from server: %v", err)
		}
		log.Printf("Connected to PingAccess version: %s", *version.Version)
		pfPort := pfCont.GetPort("9999/tcp")
		pfUrl, _ := url.Parse(fmt.Sprintf("https://localhost:%s", pfPort))
		pfClient := &http.Client{}
		if err = pool.Retry(func() error {
			log.Println("Attempting to connect to PingFederate admin API....")
			_, err := connectToPF(pfClient, pfUrl.String())
			return err
		}); err != nil {
			log.Fatalf("Could not connect to pingaccess: %s", err)
		}
		err = setupPF(pfUrl.String())
		if err != nil {
			log.Fatalf("Failed to setup PF server: %v", err)
		}
		log.Printf("Connected to PingFederate setup complete")

		paCont.Expire(360)
		pfCont.Expire(360)
		//resource.TestMain(m)
		code := m.Run()
		paCont.Close()
		pfCont.Close()
		network.Close()
		log.Printf("Tests complete shutting down container")

		os.Exit(code)
	} else {
		m.Run()
		//resource.TestMain(m)
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

func connectToPF(client *http.Client, adminUrl string) (*http.Response, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/pf-admin-api/v1/serverSettings", adminUrl), nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth("Administrator", "2FederateM0re")
	req.Header.Add("X-Xsrf-Header", "pingfederate")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New("Incorrect response code from admin: " + resp.Status)
	}
	return resp, nil
}

func setupPF(adminUrl string) error {
	client := &http.Client{}
	resp, err := connectToPF(client, adminUrl)
	if err != nil {
		return err
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	s := string(bodyText)
	s, err = sjson.Set(s, "rolesAndProtocols.oauthRole.enableOpenIdConnect", true)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/pf-admin-api/v1/serverSettings", adminUrl), bytes.NewBuffer([]byte(s)))
	req.SetBasicAuth("Administrator", "2FederateM0re")
	req.Header.Add("X-Xsrf-Header", "pingfederate")
	req.Header.Set("Content-Type", "application/json")
	resp, err = client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return errors.New("Incorrect response code from admin: " + resp.Status)
	}
	return nil
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
