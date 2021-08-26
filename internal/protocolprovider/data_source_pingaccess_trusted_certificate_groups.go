package protocol

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/iwarapter/pingaccess-sdk-go/v62/services/trustedCertificateGroups"
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
			Description: "Use this data source to get the ID of a trusted certificate group in Ping Access, you can reference it by name without having to hard code the IDs as input.",
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
					Computed:    true,
					Description: "The IDs of the certificates that are in the trusted certificate group.",
				},
				{
					Name:        "ignore_all_certificate_errors",
					Type:        tftypes.Bool,
					Computed:    true,
					Description: "This field is read-only and is only set to true for the Trust Any certificate group.",
				},
				{
					Name:        "name",
					Type:        tftypes.String,
					Required:    true,
					Description: "The name of trusted certificate group.",
				},
				{
					Name:        "skip_certificate_date_check",
					Type:        tftypes.Bool,
					Computed:    true,
					Description: "This field is true if certificates that have expired or are not yet valid but have passed the other certificate checks should be trusted.",
				},
				{
					Name:        "system_group",
					Type:        tftypes.Bool,
					Computed:    true,
					Description: "This field is read-only and indicates the trusted certificate group cannot be modified.",
				},
				{
					Name:        "use_java_trust_store",
					Type:        tftypes.Bool,
					Computed:    true,
					Description: "This field is true if the certificates in the group should also include all certificates in the Java Trust Store.",
				},
			},
		},
	}
}

func (d dataPingAccessTrustedCertificateGroups) ValidateDataSourceConfig(ctx context.Context, req *tfprotov5.ValidateDataSourceConfigRequest) (*tfprotov5.ValidateDataSourceConfigResponse, error) {
	_, diags := dynamicValueToTftypesValues(req.Config, d.trustedCertGroupType())
	if len(diags) > 0 {
		return &tfprotov5.ValidateDataSourceConfigResponse{
			Diagnostics: diags,
		}, nil
	}

	return &tfprotov5.ValidateDataSourceConfigResponse{}, nil
}

func (d dataPingAccessTrustedCertificateGroups) ReadDataSource(ctx context.Context, req *tfprotov5.ReadDataSourceRequest) (*tfprotov5.ReadDataSourceResponse, error) {
	values, diags := dynamicValueToTftypesValues(req.Config, d.trustedCertGroupType())
	if len(diags) > 0 {
		return &tfprotov5.ReadDataSourceResponse{
			Diagnostics: diags,
		}, nil
	}
	if values["name"].IsKnown() && !values["name"].IsNull() {
		var name string
		err := values["name"].As(&name)
		if err != nil {
			return &tfprotov5.ReadDataSourceResponse{
				Diagnostics: []*tfprotov5.Diagnostic{
					{
						Severity: tfprotov5.DiagnosticSeverityError,
						Summary:  unexpectedConfigurationFormat,
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
						Summary:  unexpectedConfigurationFormat,
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
						Summary:  unexpectedConfigurationFormat,
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
						Summary:  unexpectedConfigurationFormat,
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
		return &tfprotov5.ReadDataSourceResponse{
			State: &state,
		}, nil
	}
	return &tfprotov5.ReadDataSourceResponse{
		Diagnostics: []*tfprotov5.Diagnostic{
			{
				Severity: tfprotov5.DiagnosticSeverityError,
				Summary:  unexpectedConfigurationFormat,
				Detail:   "name field not known??",
			},
		},
	}, nil
}
