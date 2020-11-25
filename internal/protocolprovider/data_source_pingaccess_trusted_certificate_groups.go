package protocol

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/tftypes"
	"github.com/iwarapter/pingaccess-sdk-go/services/trustedCertificateGroups"
)

type dataPingAccessTrustedCertificateGroups struct {
	client trustedCertificateGroups.TrustedCertificateGroupsAPI
}

func (d dataPingAccessTrustedCertificateGroups) trustedCertGroupType() tftypes.Type {
	return tftypes.Object{
		AttributeTypes: d.trustedCertGroupTypes(),
	}
}

func (d dataPingAccessTrustedCertificateGroups) trustedCertGroupTypes() map[string]tftypes.Type {
	return map[string]tftypes.Type{
		"id": tftypes.String,
		"cert_ids": tftypes.List{
			ElementType: tftypes.String,
		},
		"ignore_all_certificate_errors": tftypes.Bool,
		"name":                          tftypes.String,
		"skip_certificate_date_check":   tftypes.Bool,
		"system_group":                  tftypes.Bool,
		"use_java_trust_store":          tftypes.Bool,
	}
}

func (d dataPingAccessTrustedCertificateGroups) schema() *tfprotov5.Schema {
	return &tfprotov5.Schema{
		Version: 1,
		Block: &tfprotov5.SchemaBlock{
			Attributes: []*tfprotov5.SchemaAttribute{
				{
					Name:     "id",
					Type:     tftypes.String,
					Computed: true,
				},
				{
					Name: "cert_ids",
					Type: tftypes.List{
						ElementType: tftypes.String,
					},
					Computed: true,
				},
				{
					Name:     "ignore_all_certificate_errors",
					Type:     tftypes.Bool,
					Computed: true,
				},
				{
					Name:     "name",
					Type:     tftypes.String,
					Required: true,
				},
				{
					Name:     "skip_certificate_date_check",
					Type:     tftypes.Bool,
					Computed: true,
				},
				{
					Name:     "system_group",
					Type:     tftypes.Bool,
					Computed: true,
				},
				{
					Name:     "use_java_trust_store",
					Type:     tftypes.Bool,
					Computed: true,
				},
			},
		},
	}
}

func (d dataPingAccessTrustedCertificateGroups) ValidateDataSourceConfig(ctx context.Context, req *tfprotov5.ValidateDataSourceConfigRequest) (*tfprotov5.ValidateDataSourceConfigResponse, error) {
	val, err := req.Config.Unmarshal(d.trustedCertGroupType())
	if err != nil {
		return &tfprotov5.ValidateDataSourceConfigResponse{
			Diagnostics: []*tfprotov5.Diagnostic{
				{
					Severity: tfprotov5.DiagnosticSeverityError,
					Summary:  "Unexpected configuration format",
					Detail:   "The data source got a configuration that did not match its schema, This may indication an error in the provider.\n\nError: " + err.Error(),
				},
			},
		}, nil
	}
	if !val.Is(d.trustedCertGroupType()) {
		return &tfprotov5.ValidateDataSourceConfigResponse{
			Diagnostics: []*tfprotov5.Diagnostic{
				{
					Severity: tfprotov5.DiagnosticSeverityError,
					Summary:  "Unexpected configuration format",
					Detail:   "The data source got a configuration that did not match its schema, This may indication an error in the provider.",
				},
			},
		}, nil
	}
	values := map[string]tftypes.Value{}
	err = val.As(&values)
	if err != nil {
		return &tfprotov5.ValidateDataSourceConfigResponse{
			Diagnostics: []*tfprotov5.Diagnostic{
				{
					Severity: tfprotov5.DiagnosticSeverityError,
					Summary:  "Unexpected configuration format",
					Detail:   "The data source got a configuration that did not match its schema, This may indication an error in the provider.\n\nError: " + err.Error(),
				},
			},
		}, nil
	}

	return &tfprotov5.ValidateDataSourceConfigResponse{}, nil
}

func (d dataPingAccessTrustedCertificateGroups) ReadDataSource(ctx context.Context, req *tfprotov5.ReadDataSourceRequest) (*tfprotov5.ReadDataSourceResponse, error) {
	val, err := req.Config.Unmarshal(d.trustedCertGroupType())
	if err != nil {
		return &tfprotov5.ReadDataSourceResponse{
			Diagnostics: []*tfprotov5.Diagnostic{
				{
					Severity: tfprotov5.DiagnosticSeverityError,
					Summary:  "Unexpected configuration format",
					Detail:   "The data source got a configuration that did not match its schema, This may indication an error in the provider.\n\nError: " + err.Error(),
				},
			},
		}, nil
	}
	if !val.Is(d.trustedCertGroupType()) {
		return &tfprotov5.ReadDataSourceResponse{
			Diagnostics: []*tfprotov5.Diagnostic{
				{
					Severity: tfprotov5.DiagnosticSeverityError,
					Summary:  "Unexpected configuration format",
					Detail:   "The data source got a configuration that did not match its schema, This may indication an error in the provider.",
				},
			},
		}, nil
	}
	values := map[string]tftypes.Value{}
	err = val.As(&values)
	if err != nil {
		return &tfprotov5.ReadDataSourceResponse{
			Diagnostics: []*tfprotov5.Diagnostic{
				{
					Severity: tfprotov5.DiagnosticSeverityError,
					Summary:  "Unexpected configuration format",
					Detail:   "The data source got a configuration that did not match its schema, This may indication an error in the provider.\n\nError: " + err.Error(),
				},
			},
		}, nil
	}
	if values["name"].IsKnown() {
		var name string
		err = values["name"].As(&name)
		if err != nil {
			return &tfprotov5.ReadDataSourceResponse{
				Diagnostics: []*tfprotov5.Diagnostic{
					{
						Severity: tfprotov5.DiagnosticSeverityError,
						Summary:  "Unexpected configuration format",
						Detail:   "The data source got a configuration that did not match its schema, This may indication an error in the provider.\n\nError: " + err.Error(),
					},
				},
			}, nil
		}
		input := &trustedCertificateGroups.GetTrustedCertificateGroupsCommandInput{
			Name: name,
		}
		result, _, err := d.client.GetTrustedCertificateGroupsCommand(input)
		if err != nil {
			return &tfprotov5.ReadDataSourceResponse{
				Diagnostics: []*tfprotov5.Diagnostic{
					{
						Severity: tfprotov5.DiagnosticSeverityError,
						Summary:  "Unexpected configuration format",
						Detail:   fmt.Sprintf("unable to find TrustedCertificateGroup with the name '%s', result was nil", name),
					},
				},
			}, nil
		}
		if result == nil {
			return &tfprotov5.ReadDataSourceResponse{
				Diagnostics: []*tfprotov5.Diagnostic{
					{
						Severity: tfprotov5.DiagnosticSeverityError,
						Summary:  "Unexpected configuration format",
						Detail:   fmt.Sprintf("unable to find TrustedCertificateGroup with the name '%s' found '%d' results", name, len(result.Items)),
					},
				},
			}, nil
		}
		if len(result.Items) != 1 {
			return &tfprotov5.ReadDataSourceResponse{
				Diagnostics: []*tfprotov5.Diagnostic{
					{
						Severity: tfprotov5.DiagnosticSeverityError,
						Summary:  "Unexpected configuration format",
						Detail:   fmt.Sprintf("unable to find TrustedCertificateGroup with the name '%s' found '%d' results", name, len(result.Items)),
					},
				},
			}, nil
		}
		state, err := tfprotov5.NewDynamicValue(d.trustedCertGroupType(), tftypes.NewValue(tftypes.Object{
			AttributeTypes: d.trustedCertGroupTypes(),
		}, map[string]tftypes.Value{
			"id":                            tftypes.NewValue(tftypes.String, result.Items[0].Id.String()),
			"name":                          tftypes.NewValue(tftypes.String, result.Items[0].Name),
			"ignore_all_certificate_errors": tftypes.NewValue(tftypes.Bool, result.Items[0].IgnoreAllCertificateErrors),
			"skip_certificate_date_check":   tftypes.NewValue(tftypes.Bool, result.Items[0].SkipCertificateDateCheck),
			"system_group":                  tftypes.NewValue(tftypes.Bool, result.Items[0].SystemGroup),
			"use_java_trust_store":          tftypes.NewValue(tftypes.Bool, result.Items[0].UseJavaTrustStore),
			"cert_ids":                      tftypes.NewValue(tftypes.List{ElementType: tftypes.String}, nil),
		}))
		if err != nil {
			return &tfprotov5.ReadDataSourceResponse{
				Diagnostics: []*tfprotov5.Diagnostic{
					{
						Severity: tfprotov5.DiagnosticSeverityError,
						Summary:  "Error encoding state",
						Detail:   fmt.Sprintf("Error encoding state: %s", err.Error()),
					},
				},
			}, nil
		}
		//setResourceDataBoolWithDiagnostic(d, "ignore_all_certificate_errors", input.IgnoreAllCertificateErrors, &diags)
		//setResourceDataStringWithDiagnostic(d, "name", input.Name, &diags)
		//setResourceDataBoolWithDiagnostic(d, "", input.SkipCertificateDateCheck, &diags)
		//setResourceDataBoolWithDiagnostic(d, "", input.SystemGroup, &diags)
		//setResourceDataBoolWithDiagnostic(d, "use_java_trust_store", input.UseJavaTrustStore, &diags)

		return &tfprotov5.ReadDataSourceResponse{
			State: &state,
		}, nil
	}
	return &tfprotov5.ReadDataSourceResponse{
		Diagnostics: []*tfprotov5.Diagnostic{
			{
				Severity: tfprotov5.DiagnosticSeverityError,
				Summary:  "Unexpected configuration format",
				Detail:   fmt.Sprintf("name field not known??"),
			},
		},
	}, nil
}

//
//func dataSourcePingAccessTrustedCertificateGroupsSchema() map[string]*schema.Schema {
//	return map[string]*schema.Schema{
//		"cert_ids": {
//			Type:     schema.TypeList,
//			Computed: true,
//			Elem: &schema.Schema{
//				Type: schema.TypeString,
//			},
//		},
//		"ignore_all_certificate_errors": {
//			Type:     schema.TypeBool,
//			Computed: true,
//		},
//		"name": {
//			Type:     schema.TypeString,
//			Required: true,
//		},
//		"skip_certificate_date_check": {
//			Type:     schema.TypeBool,
//			Computed: true,
//		},
//		"system_group": {
//			Type:     schema.TypeBool,
//			Computed: true,
//		},
//		"use_java_trust_store": {
//			Type:     schema.TypeBool,
//			Computed: true,
//		},
//	}
//}
//
//func dataSourcePingAccessTrustedCertificateGroupsRead(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
//	svc := m.(paClient).TrustedCertificateGroups
//	input := &trustedCertificateGroups.GetTrustedCertificateGroupsCommandInput{
//		Name: d.Get("name").(string),
//	}
//	result, _, err := svc.GetTrustedCertificateGroupsCommand(input)
//	if err != nil {
//		return diag.Errorf("unable to read TrustedCertificateGroup: %s", err)
//	}
//	if result == nil {
//		return diag.Errorf("unable to find TrustedCertificateGroup with the name '%s', result was nil", d.Get("name").(string))
//	}
//	if len(result.Items) != 1 {
//		return diag.Errorf("unable to find TrustedCertificateGroup with the name '%s' found '%d' results", d.Get("name").(string), len(result.Items))
//	}
//	d.SetId(result.Items[0].Id.String())
//	return resourcePingAccessTrustedCertificateGroupsReadResult(d, result.Items[0])
//}
