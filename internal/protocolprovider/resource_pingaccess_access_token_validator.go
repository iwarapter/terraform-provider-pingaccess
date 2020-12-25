package protocol

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/tftypes"
	"github.com/iwarapter/pingaccess-sdk-go/pingaccess/models"
	"github.com/iwarapter/pingaccess-sdk-go/services/accessTokenValidators"
)

type resourcePingAccessAccessTokenValidator struct {
	client accessTokenValidators.AccessTokenValidatorsAPI
}

func (r resourcePingAccessAccessTokenValidator) accessTokenValidatorType() tftypes.Type {
	return tftypes.Object{
		AttributeTypes: r.accessTokenValidatorTypes(),
	}
}

func (r resourcePingAccessAccessTokenValidator) accessTokenValidatorTypes() map[string]tftypes.Type {
	return map[string]tftypes.Type{
		"id":            tftypes.String,
		"name":          tftypes.String,
		"class_name":    tftypes.String,
		"configuration": tftypes.String,
	}
}

func (r resourcePingAccessAccessTokenValidator) schema() *tfprotov5.Schema {
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
					Name:     "name",
					Type:     tftypes.String,
					Required: true,
				},
				{
					Name:     "class_name",
					Type:     tftypes.String,
					Required: true,
				},
				{
					Name:     "configuration",
					Type:     tftypes.String,
					Required: true,
				},
			},
		},
	}
}

func (r resourcePingAccessAccessTokenValidator) ValidateResourceTypeConfig(_ context.Context, req *tfprotov5.ValidateResourceTypeConfigRequest) (*tfprotov5.ValidateResourceTypeConfigResponse, error) {
	val, err := req.Config.Unmarshal(r.accessTokenValidatorType())
	if err != nil {
		return &tfprotov5.ValidateResourceTypeConfigResponse{
			Diagnostics: []*tfprotov5.Diagnostic{
				{
					Severity: tfprotov5.DiagnosticSeverityError,
					Summary:  "Unexpected configuration format",
					Detail:   "The resource got a configuration that did not match its schema, This may indication an error in the provider.\n\nError: " + err.Error(),
				},
			},
		}, nil
	}
	if !val.Is(r.accessTokenValidatorType()) {
		return &tfprotov5.ValidateResourceTypeConfigResponse{
			Diagnostics: []*tfprotov5.Diagnostic{
				{
					Severity: tfprotov5.DiagnosticSeverityError,
					Summary:  "Unexpected configuration format",
					Detail:   "The resource got a configuration that did not match its schema, This may indication an error in the provider.",
				},
			},
		}, nil
	}
	values := map[string]tftypes.Value{}
	err = val.As(&values)
	if err != nil {
		return &tfprotov5.ValidateResourceTypeConfigResponse{
			Diagnostics: []*tfprotov5.Diagnostic{
				{
					Severity: tfprotov5.DiagnosticSeverityError,
					Summary:  "Unexpected configuration format",
					Detail:   "The resource got a configuration that did not match its schema, This may indication an error in the provider.\n\nError: " + err.Error(),
				},
			},
		}, nil
	}

	return &tfprotov5.ValidateResourceTypeConfigResponse{}, nil
}

func (r resourcePingAccessAccessTokenValidator) UpgradeResourceState(_ context.Context, req *tfprotov5.UpgradeResourceStateRequest) (*tfprotov5.UpgradeResourceStateResponse, error) {
	switch req.Version {
	case 1:
		val, err := req.RawState.Unmarshal(r.accessTokenValidatorType())
		if err != nil {
			return &tfprotov5.UpgradeResourceStateResponse{
				Diagnostics: []*tfprotov5.Diagnostic{
					{
						Severity: tfprotov5.DiagnosticSeverityError,
						Summary:  "Unexpected configuration format",
						Detail:   "The resource got a configuration that did not match its schema, This may indication an error in the provider.\n\nError: " + err.Error(),
					},
				},
			}, nil
		}
		dv, err := tfprotov5.NewDynamicValue(r.accessTokenValidatorType(), val)
		if err != nil {
			return &tfprotov5.UpgradeResourceStateResponse{
				Diagnostics: []*tfprotov5.Diagnostic{
					{
						Severity: tfprotov5.DiagnosticSeverityError,
						Summary:  "Unexpected configuration format",
						Detail:   "The resource got a configuration that did not match its schema, This may indication an error in the provider.\n\nError: " + err.Error(),
					},
				},
			}, nil
		}
		return &tfprotov5.UpgradeResourceStateResponse{
			UpgradedState: &dv,
		}, nil
	default:
		return &tfprotov5.UpgradeResourceStateResponse{
			Diagnostics: []*tfprotov5.Diagnostic{
				{
					Severity: tfprotov5.DiagnosticSeverityError,
					Summary:  "Unexpected configuration format",
					Detail:   "The provider doesn't know how to upgrade from the current state version. Try an earlier release of the provider.",
				},
			},
		}, nil
	}
}

func (r resourcePingAccessAccessTokenValidator) ReadResource(_ context.Context, req *tfprotov5.ReadResourceRequest) (*tfprotov5.ReadResourceResponse, error) {
	values, diags := resourceDynamicValueToTftypesValues(req.CurrentState, r.accessTokenValidatorType())
	if len(diags) > 0 {
		return &tfprotov5.ReadResourceResponse{
			Diagnostics: diags,
		}, nil
	}
	var id string
	_ = values["id"].As(&id)

	input := &accessTokenValidators.GetAccessTokenValidatorCommandInput{
		Id: id,
	}
	result, _, err := r.client.GetAccessTokenValidatorCommand(input)
	if err != nil {
		return &tfprotov5.ReadResourceResponse{
			Diagnostics: []*tfprotov5.Diagnostic{
				{
					Severity: tfprotov5.DiagnosticSeverityError,
					Summary:  "Unexpected configuration format",
					Detail:   fmt.Sprintf("unable to find AccessTokenValidator with the id '%s', result was nil", id),
				},
			},
		}, nil
	}
	if result == nil {
		return &tfprotov5.ReadResourceResponse{
			Diagnostics: []*tfprotov5.Diagnostic{
				{
					Severity: tfprotov5.DiagnosticSeverityError,
					Summary:  "Unexpected configuration format",
					Detail:   fmt.Sprintf("unable to find AccessTokenValidator with the id '%s', result was nil", id),
				},
			},
		}, nil
	}
	b, _ := json.Marshal(result.Configuration)
	state, err := tfprotov5.NewDynamicValue(r.accessTokenValidatorType(), tftypes.NewValue(tftypes.Object{
		AttributeTypes: r.accessTokenValidatorTypes(),
	}, map[string]tftypes.Value{
		"id":            tftypes.NewValue(tftypes.String, result.Id.String()),
		"name":          tftypes.NewValue(tftypes.String, result.Name),
		"class_name":    tftypes.NewValue(tftypes.String, result.ClassName),
		"configuration": tftypes.NewValue(tftypes.String, string(b)),
	}))
	if err != nil {
		return &tfprotov5.ReadResourceResponse{
			Diagnostics: []*tfprotov5.Diagnostic{
				{
					Severity: tfprotov5.DiagnosticSeverityError,
					Summary:  "Error encoding state",
					Detail:   fmt.Sprintf("Error encoding state: %s", err.Error()),
				},
			},
		}, nil
	}
	return &tfprotov5.ReadResourceResponse{
		NewState: &state,
	}, nil
}

func (r resourcePingAccessAccessTokenValidator) PlanResourceChange(_ context.Context, req *tfprotov5.PlanResourceChangeRequest) (*tfprotov5.PlanResourceChangeResponse, error) {
	proposed, err := req.ProposedNewState.Unmarshal(r.accessTokenValidatorType())
	if err != nil {
		return planResourceChangeError(err), nil
	}
	proposedValues := map[string]tftypes.Value{}
	err = proposed.As(&proposedValues)
	if err != nil {
		return planResourceChangeError(err), nil
	}
	prior, err := req.PriorState.Unmarshal(r.accessTokenValidatorType())
	if err != nil {
		return planResourceChangeError(err), nil
	}
	priorValues := map[string]tftypes.Value{}
	err = prior.As(&priorValues)
	if err != nil {
		return planResourceChangeError(err), nil
	}

	if proposed.IsNull() {
		//we plan to delete the resource
		return &tfprotov5.PlanResourceChangeResponse{
			PlannedState: req.ProposedNewState,
		}, nil
	}
	var name, className, proposedConf, priorConf, ob string
	proposedValues["name"].As(&name)
	proposedValues["class_name"].As(&className)
	_ = proposedValues["configuration"].As(&proposedConf)
	_ = priorValues["configuration"].As(&priorConf)

	var id interface{}
	if proposedValues["id"].IsKnown() && !proposedValues["id"].IsNull() {
		var str string
		proposedValues["id"].As(&str)
		id = str
	} else {
		id = tftypes.UnknownValue
	}
	ob = priorConf
	if !suppressEquivalentJSONDiffs(priorConf, proposedConf) {
		ob = proposedConf
	}

	state, err := tfprotov5.NewDynamicValue(r.accessTokenValidatorType(), tftypes.NewValue(tftypes.Object{
		AttributeTypes: r.accessTokenValidatorTypes(),
	}, map[string]tftypes.Value{
		"id":            tftypes.NewValue(tftypes.String, id),
		"name":          tftypes.NewValue(tftypes.String, name),
		"class_name":    tftypes.NewValue(tftypes.String, className),
		"configuration": tftypes.NewValue(tftypes.String, ob),
	}))
	if err != nil {
		return planResourceChangeError(err), nil
	}

	return &tfprotov5.PlanResourceChangeResponse{
		PlannedState: &state,
	}, nil
}

func (r resourcePingAccessAccessTokenValidator) ApplyResourceChange(_ context.Context, req *tfprotov5.ApplyResourceChangeRequest) (*tfprotov5.ApplyResourceChangeResponse, error) {
	planned, err := req.PlannedState.Unmarshal(r.accessTokenValidatorType())
	if err != nil {
		return applyResourceChangeError(err), nil
	}
	prior, err := req.PriorState.Unmarshal(r.accessTokenValidatorType())
	if err != nil {
		return applyResourceChangeError(err), nil
	}

	switch {
	case prior.IsNull():
		{ //create
			values := map[string]tftypes.Value{}
			err = planned.As(&values)
			if err != nil {
				return applyResourceChangeError(err), nil
			}
			var name, className, configuration string
			err = values["name"].As(&name)
			if err != nil {
				return applyResourceChangeError(err), nil
			}
			err = values["class_name"].As(&className)
			if err != nil {
				return applyResourceChangeError(err), nil
			}
			err = values["configuration"].As(&configuration)
			if err != nil {
				return applyResourceChangeError(err), nil
			}
			var dat map[string]interface{}
			_ = json.Unmarshal([]byte(configuration), &dat)
			input := &accessTokenValidators.AddAccessTokenValidatorCommandInput{
				Body: models.AccessTokenValidatorView{
					ClassName:     String(className),
					Configuration: dat,
					Name:          String(name),
				},
			}
			result, _, err := r.client.AddAccessTokenValidatorCommand(input)
			if err != nil {
				return &tfprotov5.ApplyResourceChangeResponse{
					Diagnostics: []*tfprotov5.Diagnostic{
						{
							Severity: tfprotov5.DiagnosticSeverityError,
							Summary:  "Unexpected configuration format",
							Detail:   fmt.Sprintf("unable to create AccessTokenValidator: %s", err),
						},
					},
				}, nil
			}
			b, _ := json.Marshal(result.Configuration)

			ob := string(b)
			if suppressEquivalentJSONDiffs(configuration, ob) {
				ob = configuration
			}
			state, err := tfprotov5.NewDynamicValue(r.accessTokenValidatorType(), tftypes.NewValue(tftypes.Object{
				AttributeTypes: r.accessTokenValidatorTypes(),
			}, map[string]tftypes.Value{
				"id":            tftypes.NewValue(tftypes.String, result.Id.String()),
				"name":          tftypes.NewValue(tftypes.String, result.Name),
				"class_name":    tftypes.NewValue(tftypes.String, result.ClassName),
				"configuration": tftypes.NewValue(tftypes.String, ob),
			}))
			if err != nil {
				return &tfprotov5.ApplyResourceChangeResponse{
					Diagnostics: []*tfprotov5.Diagnostic{
						{
							Severity: tfprotov5.DiagnosticSeverityError,
							Summary:  "Error encoding state",
							Detail:   fmt.Sprintf("Error encoding state: %s", err.Error()),
						},
					},
				}, nil
			}
			return &tfprotov5.ApplyResourceChangeResponse{
				NewState: &state,
			}, nil
		}
	case planned.IsNull():
		{ //delete
			values := map[string]tftypes.Value{}
			err = prior.As(&values)
			if err != nil {
				return applyResourceChangeError(err), nil
			}
			var id string
			err = values["id"].As(&id)
			if err != nil {
				return applyResourceChangeError(err), nil
			}
			input := &accessTokenValidators.DeleteAccessTokenValidatorCommandInput{
				Id: id,
			}
			_, err := r.client.DeleteAccessTokenValidatorCommand(input)
			if err != nil {
				return &tfprotov5.ApplyResourceChangeResponse{
					Diagnostics: []*tfprotov5.Diagnostic{
						{
							Severity: tfprotov5.DiagnosticSeverityError,
							Summary:  "Unexpected configuration format",
							Detail:   fmt.Sprintf("unable to delete AccessTokenValidator: %s", err),
						},
					},
				}, nil
			}
			return &tfprotov5.ApplyResourceChangeResponse{
				NewState: req.PlannedState,
			}, nil
		}
	case !planned.IsNull() && !prior.IsNull():
		{ //update
			values := map[string]tftypes.Value{}
			err = planned.As(&values)
			if err != nil {
				return applyResourceChangeError(err), nil
			}
			var id, name, className, configuration string
			err = values["id"].As(&id)
			if err != nil {
				return applyResourceChangeError(err), nil
			}
			err = values["name"].As(&name)
			if err != nil {
				return applyResourceChangeError(err), nil
			}
			err = values["class_name"].As(&className)
			if err != nil {
				return applyResourceChangeError(err), nil
			}
			err = values["configuration"].As(&configuration)
			if err != nil {
				return applyResourceChangeError(err), nil
			}
			var dat map[string]interface{}
			_ = json.Unmarshal([]byte(configuration), &dat)
			input := &accessTokenValidators.UpdateAccessTokenValidatorCommandInput{
				Id: id,
				Body: models.AccessTokenValidatorView{
					ClassName:     String(className),
					Configuration: dat,
					Name:          String(name),
				},
			}
			result, _, err := r.client.UpdateAccessTokenValidatorCommand(input)
			if err != nil {
				return &tfprotov5.ApplyResourceChangeResponse{
					Diagnostics: []*tfprotov5.Diagnostic{
						{
							Severity: tfprotov5.DiagnosticSeverityError,
							Summary:  "Unexpected configuration format",
							Detail:   fmt.Sprintf("unable to create AccessTokenValidator: %s", err),
						},
					},
				}, nil
			}
			b, _ := json.Marshal(result.Configuration)

			ob := string(b)
			if suppressEquivalentJSONDiffs(configuration, string(b)) {
				ob = configuration
			}
			state, err := tfprotov5.NewDynamicValue(r.accessTokenValidatorType(), tftypes.NewValue(tftypes.Object{
				AttributeTypes: r.accessTokenValidatorTypes(),
			}, map[string]tftypes.Value{
				"id":            tftypes.NewValue(tftypes.String, result.Id.String()),
				"name":          tftypes.NewValue(tftypes.String, result.Name),
				"class_name":    tftypes.NewValue(tftypes.String, result.ClassName),
				"configuration": tftypes.NewValue(tftypes.String, ob),
			}))
			if err != nil {
				return &tfprotov5.ApplyResourceChangeResponse{
					Diagnostics: []*tfprotov5.Diagnostic{
						{
							Severity: tfprotov5.DiagnosticSeverityError,
							Summary:  "Error encoding state",
							Detail:   fmt.Sprintf("Error encoding state: %s", err.Error()),
						},
					},
				}, nil
			}
			return &tfprotov5.ApplyResourceChangeResponse{
				NewState: &state,
			}, nil
		}
	}
	return nil, nil
}

func (r resourcePingAccessAccessTokenValidator) ImportResourceState(_ context.Context, req *tfprotov5.ImportResourceStateRequest) (*tfprotov5.ImportResourceStateResponse, error) {
	return &tfprotov5.ImportResourceStateResponse{}, nil
}

func suppressEquivalentJSONDiffs(old, new string) bool {
	ob := bytes.NewBufferString("")
	if err := json.Compact(ob, []byte(old)); err != nil {
		return false
	}

	nb := bytes.NewBufferString("")
	if err := json.Compact(nb, []byte(new)); err != nil {
		return false
	}

	return jsonBytesEqual(ob.Bytes(), nb.Bytes())
}

func jsonBytesEqual(b1, b2 []byte) bool {
	var o1 interface{}
	if err := json.Unmarshal(b1, &o1); err != nil {
		return false
	}

	var o2 interface{}
	if err := json.Unmarshal(b2, &o2); err != nil {
		return false
	}

	return reflect.DeepEqual(o1, o2)
}
