package protocol

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-go-contrib/asgotypes"

	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/tftypes"
	"github.com/iwarapter/pingaccess-sdk-go/pingaccess/models"
	"github.com/iwarapter/pingaccess-sdk-go/services/siteAuthenticators"
)

type resourcePingAccessSiteAuthenticator struct {
	client      siteAuthenticators.SiteAuthenticatorsAPI
	descriptors *models.DescriptorsView
}

func (r resourcePingAccessSiteAuthenticator) siteAuthenticatorType() tftypes.Type {
	return tftypes.Object{
		AttributeTypes: r.siteAuthenticatorTypes(),
	}
}

func (r resourcePingAccessSiteAuthenticator) siteAuthenticatorTypes() map[string]tftypes.Type {
	return map[string]tftypes.Type{
		"id":            tftypes.String,
		"name":          tftypes.String,
		"class_name":    tftypes.String,
		"configuration": tftypes.DynamicPseudoType,
	}
}

func (r resourcePingAccessSiteAuthenticator) schema() *tfprotov5.Schema {
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
					Type:     tftypes.DynamicPseudoType,
					Required: true,
				},
			},
		},
	}
}

func (r resourcePingAccessSiteAuthenticator) ValidateResourceTypeConfig(_ context.Context, req *tfprotov5.ValidateResourceTypeConfigRequest) (*tfprotov5.ValidateResourceTypeConfigResponse, error) {
	resp, values := valuesFromTypeConfigRequest(req, r.siteAuthenticatorType())
	if resp != nil {
		return resp, nil
	}
	if r.descriptors == nil { //no client config - we're done.
		return &tfprotov5.ValidateResourceTypeConfigResponse{}, nil
	}
	var name, className string
	var configuration asgotypes.GoPrimitive
	_ = values["name"].As(&name)
	_ = values["class_name"].As(&className)
	_ = values["configuration"].As(&configuration)
	diags := []*tfprotov5.Diagnostic{}
	diags = append(diags, descriptorsHasClassName(className, r.descriptors))
	if diags = append(diags, validateConfiguration(className, configuration, r.descriptors)...); len(diags) > 0 {
		return &tfprotov5.ValidateResourceTypeConfigResponse{
			Diagnostics: diags,
		}, nil
	}

	return &tfprotov5.ValidateResourceTypeConfigResponse{}, nil
}

func (r resourcePingAccessSiteAuthenticator) UpgradeResourceState(_ context.Context, req *tfprotov5.UpgradeResourceStateRequest) (*tfprotov5.UpgradeResourceStateResponse, error) {
	switch req.Version {
	case 1:
		val, err := req.RawState.Unmarshal(r.siteAuthenticatorType())
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
		dv, err := tfprotov5.NewDynamicValue(r.siteAuthenticatorType(), val)
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

func (r resourcePingAccessSiteAuthenticator) ReadResource(_ context.Context, req *tfprotov5.ReadResourceRequest) (*tfprotov5.ReadResourceResponse, error) {
	values, diags := resourceDynamicValueToTftypesValues(req.CurrentState, r.siteAuthenticatorType())
	if len(diags) > 0 {
		return &tfprotov5.ReadResourceResponse{
			Diagnostics: diags,
		}, nil
	}
	var id string
	_ = values["id"].As(&id)

	input := &siteAuthenticators.GetSiteAuthenticatorCommandInput{
		Id: id,
	}
	result, _, err := r.client.GetSiteAuthenticatorCommand(input)
	if err != nil {
		return readResourceChangeError(fmt.Errorf("unable to find SiteAuthenticator with the id '%s', result was nil", id)), nil
	}
	if result == nil {
		return readResourceChangeError(fmt.Errorf("unable to find SiteAuthenticator with the id '%s', result was nil", id)), nil
	}
	var className string
	_ = values["class_name"].As(&className)
	var configuration asgotypes.GoPrimitive
	_ = values["configuration"].As(&configuration)
	var v tftypes.Value
	if _, ok := configuration.Value.(string); ok {
		b, _ := json.Marshal(result.Configuration)
		str := maskConfigFromDescriptors(r.descriptors, className, string(b), configuration.Value.(string))
		if suppressEquivalentJSONDiffs(configuration.Value.(string), str) {
			v = tftypes.NewValue(tftypes.String, configuration.Value.(string))
		} else {
			v = tftypes.NewValue(tftypes.String, str)
		}
	} else {
		var dat map[string]interface{}
		s := maskConfigFromDescriptorsAsMap(r.descriptors, className, result.Configuration, configuration.Value.(map[string]interface{}))
		_ = json.Unmarshal([]byte(s), &dat)
		_, v, _ = marshal(dat)
	}
	state, err := tfprotov5.NewDynamicValue(r.siteAuthenticatorType(), tftypes.NewValue(tftypes.Object{
		AttributeTypes: r.siteAuthenticatorTypes(),
	}, map[string]tftypes.Value{
		"id":            tftypes.NewValue(tftypes.String, result.Id.String()),
		"name":          tftypes.NewValue(tftypes.String, result.Name),
		"class_name":    tftypes.NewValue(tftypes.String, result.ClassName),
		"configuration": v,
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

func (r resourcePingAccessSiteAuthenticator) PlanResourceChange(_ context.Context, req *tfprotov5.PlanResourceChangeRequest) (*tfprotov5.PlanResourceChangeResponse, error) {
	proposed, err := req.ProposedNewState.Unmarshal(r.siteAuthenticatorType())
	if err != nil {
		return planResourceChangeError(err), nil
	}
	proposedValues := map[string]tftypes.Value{}
	err = proposed.As(&proposedValues)
	if err != nil {
		return planResourceChangeError(err), nil
	}
	prior, err := req.PriorState.Unmarshal(r.siteAuthenticatorType())
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
	var name, className string
	var proposedConf, priorConf, configuration asgotypes.GoPrimitive
	_ = proposedValues["name"].As(&name)
	_ = proposedValues["class_name"].As(&className)
	_ = proposedValues["configuration"].As(&proposedConf)
	_ = priorValues["configuration"].As(&priorConf)

	var id interface{}
	if proposedValues["id"].IsKnown() && !proposedValues["id"].IsNull() {
		var str string
		_ = proposedValues["id"].As(&str)
		id = str
	} else {
		id = tftypes.UnknownValue
	}
	configuration = proposedConf
	var v tftypes.Value
	if _, ok := configuration.Value.(string); ok {
		v = tftypes.NewValue(tftypes.String, configuration.Value.(string))
	} else {
		_, v, _ = marshal(configuration.Value)
	}
	state, err := tfprotov5.NewDynamicValue(r.siteAuthenticatorType(), tftypes.NewValue(tftypes.Object{
		AttributeTypes: r.siteAuthenticatorTypes(),
	}, map[string]tftypes.Value{
		"id":            tftypes.NewValue(tftypes.String, id),
		"name":          tftypes.NewValue(tftypes.String, name),
		"class_name":    tftypes.NewValue(tftypes.String, className),
		"configuration": v,
	}))
	if err != nil {
		return planResourceChangeError(err), nil
	}

	return &tfprotov5.PlanResourceChangeResponse{
		PlannedState: &state,
	}, nil
}

func (r resourcePingAccessSiteAuthenticator) ApplyResourceChange(_ context.Context, req *tfprotov5.ApplyResourceChangeRequest) (*tfprotov5.ApplyResourceChangeResponse, error) {
	planned, err := req.PlannedState.Unmarshal(r.siteAuthenticatorType())
	if err != nil {
		return applyResourceChangeError(err), nil
	}
	prior, err := req.PriorState.Unmarshal(r.siteAuthenticatorType())
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
			var name, className string
			err = values["name"].As(&name)
			if err != nil {
				return applyResourceChangeError(err), nil
			}
			err = values["class_name"].As(&className)
			if err != nil {
				return applyResourceChangeError(err), nil
			}
			var configuration asgotypes.GoPrimitive
			err = values["configuration"].As(&configuration)
			if err != nil {
				return applyResourceChangeError(err), nil
			}
			var dat map[string]interface{}
			if _, ok := configuration.Value.(string); ok {
				_ = json.Unmarshal([]byte(configuration.Value.(string)), &dat)
			} else {
				dat = configuration.Value.(map[string]interface{})
			}
			input := &siteAuthenticators.AddSiteAuthenticatorCommandInput{
				Body: models.SiteAuthenticatorView{
					ClassName:     String(className),
					Configuration: dat,
					Name:          String(name),
				},
			}
			result, _, err := r.client.AddSiteAuthenticatorCommand(input)
			if err != nil {
				return &tfprotov5.ApplyResourceChangeResponse{
					Diagnostics: []*tfprotov5.Diagnostic{
						{
							Severity: tfprotov5.DiagnosticSeverityError,
							Summary:  "Unexpected configuration format",
							Detail:   fmt.Sprintf("unable to create SiteAuthenticator: %s", err),
						},
					},
				}, nil
			}
			var v tftypes.Value
			if _, ok := configuration.Value.(string); ok {
				b, _ := json.Marshal(result.Configuration)
				str := maskConfigFromDescriptors(r.descriptors, className, string(b), configuration.Value.(string))
				if suppressEquivalentJSONDiffs(configuration.Value.(string), str) {
					v = tftypes.NewValue(tftypes.String, configuration.Value.(string))
				} else {
					v = tftypes.NewValue(tftypes.String, str)
				}
			} else {
				var dat map[string]interface{}
				s := maskConfigFromDescriptorsAsMap(r.descriptors, className, result.Configuration, configuration.Value.(map[string]interface{}))
				_ = json.Unmarshal([]byte(s), &dat)
				_, v, _ = marshal(dat)
			}

			state, err := tfprotov5.NewDynamicValue(r.siteAuthenticatorType(), tftypes.NewValue(tftypes.Object{
				AttributeTypes: r.siteAuthenticatorTypes(),
			}, map[string]tftypes.Value{
				"id":            tftypes.NewValue(tftypes.String, result.Id.String()),
				"name":          tftypes.NewValue(tftypes.String, result.Name),
				"class_name":    tftypes.NewValue(tftypes.String, result.ClassName),
				"configuration": v,
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
			input := &siteAuthenticators.DeleteSiteAuthenticatorCommandInput{
				Id: id,
			}
			_, err := r.client.DeleteSiteAuthenticatorCommand(input)
			if err != nil {
				return &tfprotov5.ApplyResourceChangeResponse{
					Diagnostics: []*tfprotov5.Diagnostic{
						{
							Severity: tfprotov5.DiagnosticSeverityError,
							Summary:  "Unexpected configuration format",
							Detail:   fmt.Sprintf("unable to delete SiteAuthenticator: %s", err),
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
			var id, name, className string
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
			var configuration asgotypes.GoPrimitive
			err = values["configuration"].As(&configuration)
			if err != nil {
				return applyResourceChangeError(err), nil
			}
			var dat map[string]interface{}
			if _, ok := configuration.Value.(string); ok {
				_ = json.Unmarshal([]byte(configuration.Value.(string)), &dat)
			} else {
				dat = configuration.Value.(map[string]interface{})
			}
			input := &siteAuthenticators.UpdateSiteAuthenticatorCommandInput{
				Id: id,
				Body: models.SiteAuthenticatorView{
					ClassName:     String(className),
					Configuration: dat,
					Name:          String(name),
				},
			}
			result, _, err := r.client.UpdateSiteAuthenticatorCommand(input)
			if err != nil {
				return &tfprotov5.ApplyResourceChangeResponse{
					Diagnostics: []*tfprotov5.Diagnostic{
						{
							Severity: tfprotov5.DiagnosticSeverityError,
							Summary:  "Unexpected configuration format",
							Detail:   fmt.Sprintf("unable to create SiteAuthenticator: %s", err),
						},
					},
				}, nil
			}
			var v tftypes.Value
			if _, ok := configuration.Value.(string); ok {
				b, _ := json.Marshal(result.Configuration)
				str := maskConfigFromDescriptors(r.descriptors, className, string(b), configuration.Value.(string))
				if suppressEquivalentJSONDiffs(configuration.Value.(string), str) {
					v = tftypes.NewValue(tftypes.String, configuration.Value.(string))
				} else {
					v = tftypes.NewValue(tftypes.String, str)
				}
			} else {
				var dat map[string]interface{}
				s := maskConfigFromDescriptorsAsMap(r.descriptors, className, result.Configuration, configuration.Value.(map[string]interface{}))
				_ = json.Unmarshal([]byte(s), &dat)
				_, v, _ = marshal(dat)
			}

			state, err := tfprotov5.NewDynamicValue(r.siteAuthenticatorType(), tftypes.NewValue(tftypes.Object{
				AttributeTypes: r.siteAuthenticatorTypes(),
			}, map[string]tftypes.Value{
				"id":            tftypes.NewValue(tftypes.String, result.Id.String()),
				"name":          tftypes.NewValue(tftypes.String, result.Name),
				"class_name":    tftypes.NewValue(tftypes.String, result.ClassName),
				"configuration": v,
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

func (r resourcePingAccessSiteAuthenticator) ImportResourceState(_ context.Context, req *tfprotov5.ImportResourceStateRequest) (*tfprotov5.ImportResourceStateResponse, error) {
	return &tfprotov5.ImportResourceStateResponse{}, nil
}