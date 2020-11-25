package protocol

import (
	"context"

	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
)

type errUnsupportedDataSource string

func (e errUnsupportedDataSource) Error() string {
	return "unsupported data source: " + string(e)
}

func (p *provider) ValidateDataSourceConfig(ctx context.Context, req *tfprotov5.ValidateDataSourceConfigRequest) (*tfprotov5.ValidateDataSourceConfigResponse, error) {
	switch req.TypeName {
	case "pingaccess_trusted_certificate_group":
		res := &dataPingAccessTrustedCertificateGroups{
			client: nil,
		}
		return res.ValidateDataSourceConfig(ctx, req)

	}
	return nil, errUnsupportedDataSource(req.TypeName)
}

func (p *provider) ReadDataSource(ctx context.Context, req *tfprotov5.ReadDataSourceRequest) (*tfprotov5.ReadDataSourceResponse, error) {
	switch req.TypeName {
	case "pingaccess_trusted_certificate_group":
		res := &dataPingAccessTrustedCertificateGroups{
			client: p.client.TrustedCertificateGroups,
		}
		return res.ReadDataSource(ctx, req)

	}
	return nil, errUnsupportedDataSource(req.TypeName)
}
