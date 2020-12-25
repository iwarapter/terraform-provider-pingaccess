package protocol

import (
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/tftypes"
)

func dynamicValueToTftypesValues(conf *tfprotov5.DynamicValue, types tftypes.Type) (map[string]tftypes.Value, []*tfprotov5.Diagnostic) {
	val, err := conf.Unmarshal(types)
	if err != nil {
		return nil, []*tfprotov5.Diagnostic{
			{
				Severity: tfprotov5.DiagnosticSeverityError,
				Summary:  "Unexpected configuration format",
				Detail:   "The data source got a configuration that did not match its schema, This may indication an error in the provider.\n\nError: " + err.Error(),
			},
		}
	}
	if !val.Is(types) {
		return nil, []*tfprotov5.Diagnostic{
			{
				Severity: tfprotov5.DiagnosticSeverityError,
				Summary:  "Unexpected configuration format",
				Detail:   "The data source got a configuration that did not match its schema, This may indication an error in the provider.",
			},
		}
	}
	values := map[string]tftypes.Value{}
	err = val.As(&values)
	if err != nil {
		return nil, []*tfprotov5.Diagnostic{
			{
				Severity: tfprotov5.DiagnosticSeverityError,
				Summary:  "Unexpected configuration format",
				Detail:   "The data source got a configuration that did not match its schema, This may indication an error in the provider.\n\nError: " + err.Error(),
			},
		}
	}
	return values, nil
}

func resourceDynamicValueToTftypesValues(conf *tfprotov5.DynamicValue, types tftypes.Type) (map[string]tftypes.Value, []*tfprotov5.Diagnostic) {
	val, err := conf.Unmarshal(types)
	if err != nil {
		return nil, []*tfprotov5.Diagnostic{
			{
				Severity: tfprotov5.DiagnosticSeverityError,
				Summary:  "Unexpected configuration format",
				Detail:   "The resource got a configuration that did not match its schema, This may indication an error in the provider.\n\nError: " + err.Error(),
			},
		}
	}
	if !val.Is(types) {
		return nil, []*tfprotov5.Diagnostic{
			{
				Severity: tfprotov5.DiagnosticSeverityError,
				Summary:  "Unexpected configuration format",
				Detail:   "The resource got a configuration that did not match its schema, This may indication an error in the provider.",
			},
		}
	}
	values := map[string]tftypes.Value{}
	err = val.As(&values)
	if err != nil {
		return nil, []*tfprotov5.Diagnostic{
			{
				Severity: tfprotov5.DiagnosticSeverityError,
				Summary:  "Unexpected configuration format",
				Detail:   "The resource got a configuration that did not match its schema, This may indication an error in the provider.\n\nError: " + err.Error(),
			},
		}
	}
	return values, nil
}

func planResourceChangeError(err error) *tfprotov5.PlanResourceChangeResponse {
	return &tfprotov5.PlanResourceChangeResponse{
		Diagnostics: []*tfprotov5.Diagnostic{
			{
				Severity: tfprotov5.DiagnosticSeverityError,
				Summary:  "Unexpected configuration format",
				Detail:   "The resource got a configuration that did not match its schema, This may indication an error in the provider.\n\nError: " + err.Error(),
			},
		},
	}
}

func applyResourceChangeError(err error) *tfprotov5.ApplyResourceChangeResponse {
	return &tfprotov5.ApplyResourceChangeResponse{
		Diagnostics: []*tfprotov5.Diagnostic{
			{
				Severity: tfprotov5.DiagnosticSeverityError,
				Summary:  "Unexpected configuration format",
				Detail:   "The resource got a configuration that did not match its schema, This may indication an error in the provider.\n\nError: " + err.Error(),
			},
		},
	}
}
