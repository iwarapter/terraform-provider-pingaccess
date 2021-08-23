package protocol

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/hashicorp/terraform-plugin-go-contrib/asgotypes"

	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess/models"
	"github.com/iwarapter/pingaccess-sdk-go/v62/services/accessTokenValidators"
)

type resourcePingAccessAccessTokenValidator struct {
	client accessTokenValidators.AccessTokenValidatorsAPI
	genericPluginResource
}

func (r resourcePingAccessAccessTokenValidator) schema() *tfprotov5.Schema {
	return &tfprotov5.Schema{
		Version: 1,
		Block: &tfprotov5.SchemaBlock{
			Description: `Provides configuration for Access Token Validators within PingAccess.

-> The PingAccess API does not provider repeatable means of querying a sensitive value, we are unable to detect configuration drift of any sensitive fields in the ` + "configuration" + ` block.
`,
			Attributes: []*tfprotov5.SchemaAttribute{
				{
					Name:        "id",
					Type:        tftypes.String,
					Computed:    true,
					Description: "When creating a new AccessTokenValidator, this is the ID for the AccessTokenValidator. If not specified, an ID will be automatically assigned.",
				},
				{
					Name:        "name",
					Type:        tftypes.String,
					Required:    true,
					Description: "The access token validator's name.",
				},
				{
					Name:        "class_name",
					Type:        tftypes.String,
					Required:    true,
					Description: "The access token validator's class name.",
				},
				{
					Name:        "configuration",
					Type:        tftypes.DynamicPseudoType,
					Required:    true,
					Description: "The access token validator's configuration data.",
				},
			},
		},
	}
}

func (r resourcePingAccessAccessTokenValidator) UpgradeResourceState(_ context.Context, req *tfprotov5.UpgradeResourceStateRequest) (*tfprotov5.UpgradeResourceStateResponse, error) {
	switch req.Version {
	case 1:
		val, err := req.RawState.Unmarshal(r.resourceType())
		if err != nil {
			return &tfprotov5.UpgradeResourceStateResponse{
				Diagnostics: []*tfprotov5.Diagnostic{
					{
						Severity: tfprotov5.DiagnosticSeverityError,
						Summary:  unexpectedConfigurationFormat,
						Detail:   "The resource got a configuration that did not match its schema, This may indication an error in the provider.\n\nError: " + err.Error(),
					},
				},
			}, nil
		}
		dv, err := tfprotov5.NewDynamicValue(r.resourceType(), val)
		if err != nil {
			return &tfprotov5.UpgradeResourceStateResponse{
				Diagnostics: []*tfprotov5.Diagnostic{
					{
						Severity: tfprotov5.DiagnosticSeverityError,
						Summary:  unexpectedConfigurationFormat,
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
					Summary:  unexpectedConfigurationFormat,
					Detail:   "The provider doesn't know how to upgrade from the current state version. Try an earlier release of the provider.",
				},
			},
		}, nil
	}
}

func (r resourcePingAccessAccessTokenValidator) ReadResource(_ context.Context, req *tfprotov5.ReadResourceRequest) (*tfprotov5.ReadResourceResponse, error) {
	values, diags := resourceDynamicValueToTftypesValues(req.CurrentState, r.resourceType())
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
		return readResourceChangeError(fmt.Errorf("unable to find AccessTokenValidator with the id '%s', result was nil", id)), nil
	}
	if result == nil {
		return readResourceChangeError(fmt.Errorf("unable to find AccessTokenValidator with the id '%s', result was nil", id)), nil
	}
	var configuration asgotypes.GoPrimitive
	_ = values["configuration"].As(&configuration)
	var v tftypes.Value
	if _, ok := configuration.Value.(string); ok {
		b, _ := json.Marshal(result.Configuration)
		if suppressEquivalentJSONDiffs(configuration.Value.(string), string(b)) {
			v = tftypes.NewValue(tftypes.String, configuration.Value.(string))
		} else {
			v = tftypes.NewValue(tftypes.String, string(b))
		}
	} else {
		_, v, _ = marshal(result.Configuration)
	}
	state, err := createGenericClassResourceState(r.resourceType(), r.resourceTypes(), result.Id.String(), *result.Name, *result.ClassName, v)
	if err != nil {
		return &tfprotov5.ReadResourceResponse{Diagnostics: []*tfprotov5.Diagnostic{stateEncodingDiagnostic(err)}}, nil
	}
	return &tfprotov5.ReadResourceResponse{
		NewState: &state,
	}, nil
}

func (r resourcePingAccessAccessTokenValidator) ApplyResourceChange(_ context.Context, req *tfprotov5.ApplyResourceChangeRequest) (*tfprotov5.ApplyResourceChangeResponse, error) {
	planned, err := req.PlannedState.Unmarshal(r.resourceType())
	if err != nil {
		return applyResourceChangeError(err), nil
	}
	prior, err := req.PriorState.Unmarshal(r.resourceType())
	if err != nil {
		return applyResourceChangeError(err), nil
	}

	switch {
	case prior.IsNull():
		{ //create
			return r.genericPluginResourceCreate(planned, func(name, class string, dat map[string]interface{}) (string, string, string, map[string]interface{}, error) {
				input := &accessTokenValidators.AddAccessTokenValidatorCommandInput{
					Body: models.AccessTokenValidatorView{
						ClassName:     String(class),
						Configuration: dat,
						Name:          String(name),
					},
				}
				if result, _, err := r.client.AddAccessTokenValidatorCommand(input); err != nil {
					return "", "", "", nil, fmt.Errorf("unable to create AccessTokenValidator: %s", err)
				} else {
					return result.Id.String(), *result.Name, *result.ClassName, result.Configuration, nil
				}
			})
		}
	case planned.IsNull():
		{ //delete
			return genericPluginResourceDelete(req, prior, func(id string) error {
				input := &accessTokenValidators.DeleteAccessTokenValidatorCommandInput{
					Id: id,
				}
				if _, err = r.client.DeleteAccessTokenValidatorCommand(input); err != nil {
					return fmt.Errorf("unable to delete AccessTokenValidator: %s", err)
				}
				return nil
			})
		}
	case !planned.IsNull() && !prior.IsNull():
		{ //update
			return r.genericPluginResourceUpdate(planned, func(id, name, class string, dat map[string]interface{}) (string, string, string, map[string]interface{}, error) {
				input := &accessTokenValidators.UpdateAccessTokenValidatorCommandInput{
					Id: id,
					Body: models.AccessTokenValidatorView{
						ClassName:     String(class),
						Configuration: dat,
						Name:          String(name),
					},
				}
				if result, _, err := r.client.UpdateAccessTokenValidatorCommand(input); err != nil {
					return "", "", "", nil, fmt.Errorf("unable to update AccessTokenValidator: %s", err)
				} else {
					return result.Id.String(), *result.Name, *result.ClassName, result.Configuration, nil
				}
			})
		}
	}
	return nil, nil
}

func createGenericClassResourceState(typ tftypes.Type, attrTypes map[string]tftypes.Type, id interface{}, name, class string, config tftypes.Value) (tfprotov5.DynamicValue, error) {
	return tfprotov5.NewDynamicValue(typ, tftypes.NewValue(tftypes.Object{
		AttributeTypes: attrTypes,
	}, map[string]tftypes.Value{
		"id":            tftypes.NewValue(tftypes.String, id),
		"name":          tftypes.NewValue(tftypes.String, name),
		"class_name":    tftypes.NewValue(tftypes.String, class),
		"configuration": config,
	}))
}

func (r resourcePingAccessAccessTokenValidator) ImportResourceState(_ context.Context, req *tfprotov5.ImportResourceStateRequest) (*tfprotov5.ImportResourceStateResponse, error) {
	result, _, err := r.client.GetAccessTokenValidatorCommand(&accessTokenValidators.GetAccessTokenValidatorCommandInput{Id: req.ID})
	if err != nil {
		return importResourceError(fmt.Sprintf("The provider was unable to retrieve the access token validator with ID: '%s'.\n\nError:\n%s", req.ID, err.Error())), nil
	}
	var v tftypes.Value
	_, v, _ = marshal(result.Configuration)
	state, err := createGenericClassResourceState(r.resourceType(), r.resourceTypes(), result.Id.String(), *result.Name, *result.ClassName, v)
	if err != nil {
		return &tfprotov5.ImportResourceStateResponse{Diagnostics: []*tfprotov5.Diagnostic{stateEncodingDiagnostic(err)}}, nil
	}
	return &tfprotov5.ImportResourceStateResponse{
		ImportedResources: []*tfprotov5.ImportedResource{
			{
				TypeName: req.TypeName,
				State:    &state,
			},
		},
	}, nil
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
