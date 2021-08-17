package protocol

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-go-contrib/asgotypes"

	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/iwarapter/pingaccess-sdk-go/pingaccess/models"
	"github.com/iwarapter/pingaccess-sdk-go/services/siteAuthenticators"
)

type resourcePingAccessSiteAuthenticator struct {
	client siteAuthenticators.SiteAuthenticatorsAPI
	genericPluginResource
}

func (r resourcePingAccessSiteAuthenticator) schema() *tfprotov5.Schema {
	return r.genericPluginResource.schema()
}

func (r resourcePingAccessSiteAuthenticator) UpgradeResourceState(_ context.Context, req *tfprotov5.UpgradeResourceStateRequest) (*tfprotov5.UpgradeResourceStateResponse, error) {
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

func (r resourcePingAccessSiteAuthenticator) ReadResource(_ context.Context, req *tfprotov5.ReadResourceRequest) (*tfprotov5.ReadResourceResponse, error) {
	values, diags := resourceDynamicValueToTftypesValues(req.CurrentState, r.resourceType())
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
	state, err := createGenericClassResourceState(r.resourceType(), r.resourceTypes(), result.Id.String(), *result.Name, *result.ClassName, v)
	if err != nil {
		return &tfprotov5.ReadResourceResponse{Diagnostics: []*tfprotov5.Diagnostic{stateEncodingDiagnostic(err)}}, nil
	}
	return &tfprotov5.ReadResourceResponse{
		NewState: &state,
	}, nil
}

func (r resourcePingAccessSiteAuthenticator) ApplyResourceChange(_ context.Context, req *tfprotov5.ApplyResourceChangeRequest) (*tfprotov5.ApplyResourceChangeResponse, error) {
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
				input := &siteAuthenticators.AddSiteAuthenticatorCommandInput{
					Body: models.SiteAuthenticatorView{
						ClassName:     String(class),
						Configuration: dat,
						Name:          String(name),
					},
				}
				if result, _, err := r.client.AddSiteAuthenticatorCommand(input); err != nil {
					return "", "", "", nil, fmt.Errorf("unable to create SiteAuthenticator: %s", err)
				} else {
					return result.Id.String(), *result.Name, *result.ClassName, result.Configuration, nil
				}
			})
		}
	case planned.IsNull():
		{ //delete
			return genericPluginResourceDelete(req, prior, func(id string) error {
				input := &siteAuthenticators.DeleteSiteAuthenticatorCommandInput{
					Id: id,
				}
				if _, err = r.client.DeleteSiteAuthenticatorCommand(input); err != nil {
					return fmt.Errorf("unable to delete SiteAuthenticator: %s", err)
				}
				return nil
			})
		}
	case !planned.IsNull() && !prior.IsNull():
		{ //update
			return r.genericPluginResourceUpdate(planned, func(id, name, class string, dat map[string]interface{}) (string, string, string, map[string]interface{}, error) {
				input := &siteAuthenticators.UpdateSiteAuthenticatorCommandInput{
					Id: id,
					Body: models.SiteAuthenticatorView{
						ClassName:     String(class),
						Configuration: dat,
						Name:          String(name),
					},
				}
				if result, _, err := r.client.UpdateSiteAuthenticatorCommand(input); err != nil {
					return "", "", "", nil, fmt.Errorf("unable to update AccessTokenValidator: %s", err)
				} else {
					return result.Id.String(), *result.Name, *result.ClassName, result.Configuration, nil
				}
			})
		}
	}
	return nil, nil
}

func (r resourcePingAccessSiteAuthenticator) ImportResourceState(_ context.Context, req *tfprotov5.ImportResourceStateRequest) (*tfprotov5.ImportResourceStateResponse, error) {
	return &tfprotov5.ImportResourceStateResponse{}, nil
}
