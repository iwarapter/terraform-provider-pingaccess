package protocol

import (
	"context"

	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
)

type errUnsupportedResource string

func (e errUnsupportedResource) Error() string {
	return "unsupported resource: " + string(e)
}

func (p *provider) ValidateResourceTypeConfig(ctx context.Context, req *tfprotov5.ValidateResourceTypeConfigRequest) (*tfprotov5.ValidateResourceTypeConfigResponse, error) {
	return nil, errUnsupportedResource(req.TypeName)
}

func (p *provider) UpgradeResourceState(ctx context.Context, req *tfprotov5.UpgradeResourceStateRequest) (*tfprotov5.UpgradeResourceStateResponse, error) {
	return nil, errUnsupportedResource(req.TypeName)
}

func (p *provider) ReadResource(ctx context.Context, req *tfprotov5.ReadResourceRequest) (*tfprotov5.ReadResourceResponse, error) {
	return nil, errUnsupportedResource(req.TypeName)
}

func (p *provider) PlanResourceChange(ctx context.Context, req *tfprotov5.PlanResourceChangeRequest) (*tfprotov5.PlanResourceChangeResponse, error) {
	return nil, errUnsupportedResource(req.TypeName)
}

func (p *provider) ApplyResourceChange(ctx context.Context, req *tfprotov5.ApplyResourceChangeRequest) (*tfprotov5.ApplyResourceChangeResponse, error) {
	return nil, errUnsupportedResource(req.TypeName)
}

func (p *provider) ImportResourceState(ctx context.Context, req *tfprotov5.ImportResourceStateRequest) (*tfprotov5.ImportResourceStateResponse, error) {
	return nil, errUnsupportedResource(req.TypeName)
}
