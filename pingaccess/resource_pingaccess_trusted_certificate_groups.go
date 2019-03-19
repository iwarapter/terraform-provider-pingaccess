package pingaccess

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"
	pa "github.com/iwarapter/pingaccess-sdk-go/pingaccess"
)

func resourcePingAccessTrustedCertificateGroups() *schema.Resource {
	return &schema.Resource{
		Create: resourcePingAccessTrustedCertificateGroupsCreate,
		Read:   resourcePingAccessTrustedCertificateGroupsRead,
		Update: resourcePingAccessTrustedCertificateGroupsUpdate,
		Delete: resourcePingAccessTrustedCertificateGroupsDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: resourcePingAccessTrustedCertificateGroupsSchema(),
	}
}

func resourcePingAccessTrustedCertificateGroupsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"cert_ids": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"ignore_all_certificate_errors": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
		},
		"name": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"skip_certificate_date_check": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
		},
		"system_group": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
		},
		"use_java_trust_store": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
	}
}

func resourcePingAccessTrustedCertificateGroupsCreate(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pa.Client).TrustedCertificateGroups
	input := pa.AddTrustedCertificateGroupCommandInput{
		Body: *resourcePingAccessTrustedCertificateGroupsReadData(d),
	}

	result, _, err := svc.AddTrustedCertificateGroupCommand(&input)
	if err != nil {
		return fmt.Errorf("Error creating third party service: %s", err)
	}

	d.SetId(result.Id.String())
	return resourcePingAccessTrustedCertificateGroupsReadResult(d, result)
}

func resourcePingAccessTrustedCertificateGroupsRead(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pa.Client).TrustedCertificateGroups
	input := &pa.GetTrustedCertificateGroupCommandInput{
		Id: d.Id(),
	}
	result, _, err := svc.GetTrustedCertificateGroupCommand(input)
	if err != nil {
		return fmt.Errorf("Error reading third party service: %s", err)
	}
	return resourcePingAccessTrustedCertificateGroupsReadResult(d, result)
}

func resourcePingAccessTrustedCertificateGroupsUpdate(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pa.Client).TrustedCertificateGroups
	input := pa.UpdateTrustedCertificateGroupCommandInput{
		Body: *resourcePingAccessTrustedCertificateGroupsReadData(d),
		Id:   d.Id(),
	}
	result, _, err := svc.UpdateTrustedCertificateGroupCommand(&input)
	if err != nil {
		return fmt.Errorf("Error updating third party service: %s", err)
	}
	return resourcePingAccessTrustedCertificateGroupsReadResult(d, result)
}

func resourcePingAccessTrustedCertificateGroupsDelete(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pa.Client).TrustedCertificateGroups
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	input := &pa.DeleteTrustedCertificateGroupCommandInput{
		Id: d.Id(),
	}

	_, err := svc.DeleteTrustedCertificateGroupCommand(input)
	if err != nil {
		return fmt.Errorf("Error deleting third party service: %s", err)
	}
	return nil
}

func resourcePingAccessTrustedCertificateGroupsReadResult(d *schema.ResourceData, input *pa.TrustedCertificateGroupView) (err error) {
	setResourceDataBool(d, "ignore_all_certificate_errors", input.IgnoreAllCertificateErrors)
	setResourceDataString(d, "name", input.Name)
	setResourceDataBool(d, "skip_certificate_date_check", input.SkipCertificateDateCheck)
	setResourceDataBool(d, "system_group", input.SystemGroup)
	setResourceDataBool(d, "use_java_trust_store", input.UseJavaTrustStore)

	if input.CertIds != nil {
		certs := []string{}
		for _, cert := range *input.CertIds {
			certs = append(certs, strconv.Itoa(*cert))
		}
		if err := d.Set("cert_ids", certs); err != nil {
			return err
		}
	}

	return nil
}

func resourcePingAccessTrustedCertificateGroupsReadData(d *schema.ResourceData) *pa.TrustedCertificateGroupView {

	tcg := &pa.TrustedCertificateGroupView{
		Name: String(d.Get("name").(string)),
	}

	if v, ok := d.GetOkExists("cert_ids"); ok {
		certs := expandStringList(v.([]interface{}))
		certIDs := []*int{}
		for _, i := range certs {
			text, _ := strconv.Atoi(*i)
			certIDs = append(certIDs, &text)
		}
		tcg.CertIds = &certIDs
	}

	if v, ok := d.GetOkExists("ignore_all_certificate_errors"); ok {
		tcg.IgnoreAllCertificateErrors = Bool(v.(bool))
	}

	if v, ok := d.GetOkExists("skip_certificate_date_check"); ok {
		tcg.SkipCertificateDateCheck = Bool(v.(bool))
	}

	if v, ok := d.GetOkExists("system_group"); ok {
		tcg.SystemGroup = Bool(v.(bool))
	}

	if v, ok := d.GetOkExists("use_java_trust_store"); ok {
		tcg.UseJavaTrustStore = Bool(v.(bool))
	}

	return tcg
}
