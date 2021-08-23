package protocol

import (
	"context"
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-go-contrib/asgotypes"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess/models"
)

type genericPluginResource struct {
	descriptors *models.DescriptorsView
}

func (r genericPluginResource) resourceType() tftypes.Type {
	return tftypes.Object{
		AttributeTypes: r.resourceTypes(),
	}
}

func (r genericPluginResource) resourceTypes() map[string]tftypes.Type {
	return map[string]tftypes.Type{
		"id":            tftypes.String,
		"name":          tftypes.String,
		"class_name":    tftypes.String,
		"configuration": tftypes.DynamicPseudoType,
	}
}

func (r genericPluginResource) schema() *tfprotov5.Schema {
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

func (r genericPluginResource) ValidateResourceTypeConfig(_ context.Context, req *tfprotov5.ValidateResourceTypeConfigRequest) (*tfprotov5.ValidateResourceTypeConfigResponse, error) {
	resp, values := valuesFromTypeConfigRequest(req, r.resourceType())
	if resp != nil {
		return resp, nil
	}
	if r.descriptors == nil { //no client config - we're done.
		return &tfprotov5.ValidateResourceTypeConfigResponse{}, nil
	}
	if !values["configuration"].IsFullyKnown() {
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

func (r genericPluginResource) PlanResourceChange(_ context.Context, req *tfprotov5.PlanResourceChangeRequest) (*tfprotov5.PlanResourceChangeResponse, error) {
	proposed, err := req.ProposedNewState.Unmarshal(r.resourceType())
	if err != nil {
		return planResourceChangeError(err), nil
	}
	proposedValues := map[string]tftypes.Value{}
	err = proposed.As(&proposedValues)
	if err != nil {
		return planResourceChangeError(err), nil
	}
	prior, err := req.PriorState.Unmarshal(r.resourceType())
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
	if !proposedValues["configuration"].IsFullyKnown() {
		v = proposedValues["configuration"]
	}

	state, err := createGenericClassResourceState(r.resourceType(), r.resourceTypes(), id, name, className, v)
	if err != nil {
		return planResourceChangeError(err), nil
	}

	return &tfprotov5.PlanResourceChangeResponse{
		PlannedState: &state,
	}, nil
}

func genericPluginResourceDelete(req *tfprotov5.ApplyResourceChangeRequest, prior tftypes.Value, cb func(id string) error) (*tfprotov5.ApplyResourceChangeResponse, error) {
	values := map[string]tftypes.Value{}
	err := prior.As(&values)
	if err != nil {
		return applyResourceChangeError(err), nil
	}
	var id string
	err = values["id"].As(&id)
	if err != nil {
		return applyResourceChangeError(err), nil
	}
	err = cb(id)
	if err != nil {
		return &tfprotov5.ApplyResourceChangeResponse{
			Diagnostics: []*tfprotov5.Diagnostic{
				{
					Severity: tfprotov5.DiagnosticSeverityError,
					Summary:  unexpectedConfigurationFormat,
					Detail:   err.Error(),
				},
			},
		}, nil
	}
	return &tfprotov5.ApplyResourceChangeResponse{
		NewState: req.PlannedState,
	}, nil
}

func (r genericPluginResource) genericPluginResourceCreate(planned tftypes.Value, cb func(name, class string, dat map[string]interface{}) (string, string, string, map[string]interface{}, error)) (*tfprotov5.ApplyResourceChangeResponse, error) {
	values := map[string]tftypes.Value{}
	err := planned.As(&values)
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
	id, name, class, result, err := cb(name, className, dat)
	if err != nil {
		return &tfprotov5.ApplyResourceChangeResponse{
			Diagnostics: []*tfprotov5.Diagnostic{
				{
					Severity: tfprotov5.DiagnosticSeverityError,
					Summary:  unexpectedConfigurationFormat,
					Detail:   err.Error(),
				},
			},
		}, nil
	}

	var v tftypes.Value
	if _, ok := configuration.Value.(string); ok {
		b, _ := json.Marshal(result)
		str := maskConfigFromDescriptors(r.descriptors, className, string(b), configuration.Value.(string))
		if suppressEquivalentJSONDiffs(configuration.Value.(string), str) {
			v = tftypes.NewValue(tftypes.String, configuration.Value.(string))
		} else {
			v = tftypes.NewValue(tftypes.String, str)
		}
	} else {
		var dat map[string]interface{}
		s := maskConfigFromDescriptorsAsMap(r.descriptors, className, result, configuration.Value.(map[string]interface{}))
		_ = json.Unmarshal([]byte(s), &dat)
		_, v, _ = marshal(dat)
	}
	state, err := createGenericClassResourceState(r.resourceType(), r.resourceTypes(), id, name, class, v)
	if err != nil {
		return &tfprotov5.ApplyResourceChangeResponse{Diagnostics: []*tfprotov5.Diagnostic{stateEncodingDiagnostic(err)}}, nil
	}
	return &tfprotov5.ApplyResourceChangeResponse{
		NewState: &state,
	}, nil
}

func (r genericPluginResource) genericPluginResourceUpdate(planned tftypes.Value, cb func(id, name, class string, dat map[string]interface{}) (string, string, string, map[string]interface{}, error)) (*tfprotov5.ApplyResourceChangeResponse, error) {
	values := map[string]tftypes.Value{}
	err := planned.As(&values)
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
	id, name, class, result, err := cb(id, name, className, dat)

	if err != nil {
		return &tfprotov5.ApplyResourceChangeResponse{
			Diagnostics: []*tfprotov5.Diagnostic{
				{
					Severity: tfprotov5.DiagnosticSeverityError,
					Summary:  unexpectedConfigurationFormat,
					Detail:   err.Error(),
				},
			},
		}, nil
	}
	var v tftypes.Value
	if _, ok := configuration.Value.(string); ok {
		b, _ := json.Marshal(result)
		str := maskConfigFromDescriptors(r.descriptors, className, string(b), configuration.Value.(string))
		if suppressEquivalentJSONDiffs(configuration.Value.(string), str) {
			v = tftypes.NewValue(tftypes.String, configuration.Value.(string))
		} else {
			v = tftypes.NewValue(tftypes.String, str)
		}
	} else {
		var dat map[string]interface{}
		s := maskConfigFromDescriptorsAsMap(r.descriptors, className, result, configuration.Value.(map[string]interface{}))
		_ = json.Unmarshal([]byte(s), &dat)
		_, v, _ = marshal(dat)
	}
	state, err := createGenericClassResourceState(r.resourceType(), r.resourceTypes(), id, name, class, v)
	if err != nil {
		return &tfprotov5.ApplyResourceChangeResponse{Diagnostics: []*tfprotov5.Diagnostic{stateEncodingDiagnostic(err)}}, nil
	}
	return &tfprotov5.ApplyResourceChangeResponse{
		NewState: &state,
	}, nil
}
