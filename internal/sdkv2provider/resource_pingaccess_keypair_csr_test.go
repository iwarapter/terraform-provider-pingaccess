package sdkv2provider

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"log"
	"math/big"
	"strings"
	"testing"
	"time"

	"github.com/iwarapter/pingaccess-sdk-go/v62/services/keyPairs"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccPingAccessKeyPairCsr(t *testing.T) {
	key, _ := rsa.GenerateKey(rand.Reader, 2048)
	ca := &x509.Certificate{
		SerialNumber: big.NewInt(2019),
		Subject: pkix.Name{
			Organization: []string{"Ping Identity"},
			Country:      []string{"US"},
			CommonName:   "localhost",
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(10, 0, 0),
		IsCA:                  true,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		BasicConstraintsValid: true,
	}

	caBytes, err := x509.CreateCertificate(rand.Reader, ca, ca, &key.PublicKey, key)
	if err != nil {
		log.Fatalf("Failed to create certificateCa: %s", err)
	}
	caBuf := new(bytes.Buffer)
	_ = pem.Encode(caBuf, &pem.Block{Type: "CERTIFICATE", Bytes: caBytes})
	svc := keyPairs.New(conf)
	csrPem, _, err := svc.GenerateCsrCommand(&keyPairs.GenerateCsrCommandInput{Id: "1"})
	if err != nil {
		t.Fatalf("unable to get CSR")
	}
	*csrPem = strings.ReplaceAll(*csrPem, " NEW ", " ")
	b, _ := pem.Decode([]byte(*csrPem))
	csr, err := x509.ParseCertificateRequest(b.Bytes)
	if err != nil {
		t.Fatalf("unable to parse csr: %s", err)
	}
	template := x509.Certificate{
		Signature:          csr.Signature,
		SignatureAlgorithm: csr.SignatureAlgorithm,

		PublicKeyAlgorithm: csr.PublicKeyAlgorithm,
		PublicKey:          csr.PublicKey,

		SerialNumber: big.NewInt(2),
		Issuer:       ca.Subject,
		Subject:      csr.Subject,
		NotBefore:    time.Now(),
		NotAfter:     time.Now().Add(24 * time.Hour),
		KeyUsage:     x509.KeyUsageDigitalSignature,
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
	}
	certBytes, err := x509.CreateCertificate(rand.Reader, &template, ca, csr.PublicKey, key)
	if err != nil {
		t.Fatalf("unable to sign certificate request: %s", err)
	}
	buf := new(bytes.Buffer)
	err = pem.Encode(buf, &pem.Block{Type: "CERTIFICATE", Bytes: certBytes})
	if err != nil {
		t.Fatalf("unable to encode certificate: %s", err)
	}
	signedCert := buf.String()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV5ProviderFactories: testAccProviders,
		CheckDestroy:             testAccCheckPingAccessKeyPairCsrDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingAccessKeyPairCsrConfig(signedCert, caBuf.String()),
				//Check: resource.ComposeTestCheckFunc(),
			},
		},
	})
}

func testAccCheckPingAccessKeyPairCsrDestroy(s *terraform.State) error {
	return nil
}

func testAccPingAccessKeyPairCsrConfig(signedCert, chain string) string {
	return fmt.Sprintf(`

data "pingaccess_trusted_certificate_group" "test" {
	name = "Trust Any"
}

resource "pingaccess_keypair_csr" "test" {
  keypair_id = "1"
  file_data = "%s"
  chain_certificates = ["%s"]
  trusted_certificate_group_id = data.pingaccess_trusted_certificate_group.test.id
}
`, base64.StdEncoding.EncodeToString([]byte(signedCert)), base64.StdEncoding.EncodeToString([]byte(chain)))
}
