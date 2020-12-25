package protocol

import (
	"context"
	"os"

	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/tftypes"
)

type provider struct {
	providerSchema     *tfprotov5.Schema
	providerMetaSchema *tfprotov5.Schema
	resourceSchemas    map[string]*tfprotov5.Schema
	dataSourceSchemas  map[string]*tfprotov5.Schema

	resourceRouter   map[string]tfprotov5.ResourceServer
	dataSourceRouter map[string]tfprotov5.DataSourceServer

	client *paClient
}

func (p *provider) GetProviderSchema(_ context.Context, _ *tfprotov5.GetProviderSchemaRequest) (*tfprotov5.GetProviderSchemaResponse, error) {
	return &tfprotov5.GetProviderSchemaResponse{
		Provider:          p.providerSchema,
		ProviderMeta:      p.providerMetaSchema,
		ResourceSchemas:   p.resourceSchemas,
		DataSourceSchemas: p.dataSourceSchemas,
	}, nil
}

func (p *provider) PrepareProviderConfig(_ context.Context, req *tfprotov5.PrepareProviderConfigRequest) (*tfprotov5.PrepareProviderConfigResponse, error) {
	return &tfprotov5.PrepareProviderConfigResponse{
		PreparedConfig: nil,
	}, nil
}

func (p *provider) ConfigureProvider(_ context.Context, req *tfprotov5.ConfigureProviderRequest) (*tfprotov5.ConfigureProviderResponse, error) {
	c := cfg{
		Username: "Administrator",
		Password: "2Access",
		Context:  "/pa-admin-api/v3",
		BaseURL:  "https://localhost:9000",
	}
	configType := tftypes.Object{
		AttributeTypes: map[string]tftypes.Type{
			"username": tftypes.String,
			"password": tftypes.String,
			"context":  tftypes.String,
			"base_url": tftypes.String,
		},
	}
	val, err := req.Config.Unmarshal(configType)
	if err != nil {
		return &tfprotov5.ConfigureProviderResponse{
			Diagnostics: []*tfprotov5.Diagnostic{
				{
					Severity: tfprotov5.DiagnosticSeverityError,
					Summary:  "Unexpected provider configuration",
					Detail:   "The provider got a configuration that did not match it's schema. This may indicate an error in the provider.\n\nError: " + err.Error(),
				},
			},
		}, nil
	}
	values := map[string]tftypes.Value{}
	err = val.As(&values)
	if err != nil {
		return &tfprotov5.ConfigureProviderResponse{
			Diagnostics: []*tfprotov5.Diagnostic{
				{
					Severity: tfprotov5.DiagnosticSeverityError,
					Summary:  "Unexpected provider configuration",
					Detail:   "The provider got a configuration that did not match it's schema. This may indicate an error in the provider.\n\nError: " + err.Error(),
				},
			},
		}, nil
	}
	if values["username"].IsKnown() && !values["username"].IsNull() {
		err = values["username"].As(&c.Username)
		if err != nil {
			return &tfprotov5.ConfigureProviderResponse{
				Diagnostics: []*tfprotov5.Diagnostic{
					{
						Severity: tfprotov5.DiagnosticSeverityError,
						Summary:  "Unexpected provider configuration",
						Detail:   "The provider got a configuration that did not match it's schema. This may indicate an error in the provider.\n\nError: " + err.Error(),
						Attribute: &tftypes.AttributePath{
							Steps: []tftypes.AttributePathStep{
								tftypes.AttributeName("username"),
							},
						},
					},
				},
			}, nil
		}
	}
	if values["password"].IsKnown() && !values["password"].IsNull() {
		err = values["password"].As(&c.Password)
		if err != nil {
			return &tfprotov5.ConfigureProviderResponse{
				Diagnostics: []*tfprotov5.Diagnostic{
					{
						Severity: tfprotov5.DiagnosticSeverityError,
						Summary:  "Unexpected provider configuration",
						Detail:   "The provider got a configuration that did not match it's schema. This may indicate an error in the provider.\n\nError: " + err.Error(),
						Attribute: &tftypes.AttributePath{
							Steps: []tftypes.AttributePathStep{
								tftypes.AttributeName("password"),
							},
						},
					},
				},
			}, nil
		}
	}
	if values["context"].IsKnown() && !values["context"].IsNull() {
		err = values["context"].As(&c.Context)
		if err != nil {
			return &tfprotov5.ConfigureProviderResponse{
				Diagnostics: []*tfprotov5.Diagnostic{
					{
						Severity: tfprotov5.DiagnosticSeverityError,
						Summary:  "Unexpected provider configuration",
						Detail:   "The provider got a configuration that did not match it's schema. This may indicate an error in the provider.\n\nError: " + err.Error(),
						Attribute: &tftypes.AttributePath{
							Steps: []tftypes.AttributePathStep{
								tftypes.AttributeName("context"),
							},
						},
					},
				},
			}, nil
		}
	}
	if values["base_url"].IsKnown() && !values["base_url"].IsNull() {
		err = values["base_url"].As(&c.BaseURL)
		if err != nil {
			return &tfprotov5.ConfigureProviderResponse{
				Diagnostics: []*tfprotov5.Diagnostic{
					{
						Severity: tfprotov5.DiagnosticSeverityError,
						Summary:  "Unexpected provider configuration",
						Detail:   "The provider got a configuration that did not match it's schema. This may indicate an error in the provider.\n\nError: " + err.Error(),
						Attribute: &tftypes.AttributePath{
							Steps: []tftypes.AttributePathStep{
								tftypes.AttributeName("base_url"),
							},
						},
					},
				},
			}, nil
		}
	}
	if v := os.Getenv("PINGACCESS_USERNAME"); v != "" {
		c.Username = v
	}
	if v := os.Getenv("PINGACCESS_PASSWORD"); v != "" {
		c.Password = v
	}
	if v := os.Getenv("PINGACCESS_CONTEXT"); v != "" {
		c.Context = v
	}
	if v := os.Getenv("PINGACCESS_BASEURL"); v != "" {
		c.BaseURL = v
	}

	var diags []*tfprotov5.Diagnostic

	client, diag := c.Client()
	if diag != nil {
		diags = append(diags, diag)
	}
	p.client = client
	return &tfprotov5.ConfigureProviderResponse{
		Diagnostics: diags,
	}, nil
}

func (p *provider) StopProvider(ctx context.Context, req *tfprotov5.StopProviderRequest) (*tfprotov5.StopProviderResponse, error) {
	return &tfprotov5.StopProviderResponse{}, nil
}

func Server() tfprotov5.ProviderServer {
	return &provider{
		providerSchema: &tfprotov5.Schema{
			Version: 0,
			Block: &tfprotov5.SchemaBlock{
				Version: 0,
				Attributes: []*tfprotov5.SchemaAttribute{
					{
						Name:            "base_url",
						Optional:        true,
						Type:            tftypes.String,
						Description:     "The base url of the pingaccess API.",
						DescriptionKind: tfprotov5.StringKindPlain,
					},
					{
						Name:            "context",
						Optional:        true,
						Type:            tftypes.String,
						Description:     "The context path of the pingaccess API.",
						DescriptionKind: tfprotov5.StringKindPlain,
					},
					{
						Name:            "password",
						Optional:        true,
						Type:            tftypes.String,
						Description:     "The password for pingaccess API.",
						DescriptionKind: tfprotov5.StringKindPlain,
						Sensitive:       true,
					},
					{
						Name:            "username",
						Optional:        true,
						Type:            tftypes.String,
						Description:     "The username for pingaccess API.",
						DescriptionKind: tfprotov5.StringKindPlain,
					},
				},
			},
		},
		dataSourceSchemas: map[string]*tfprotov5.Schema{
			"pingaccess_trusted_certificate_group": dataPingAccessTrustedCertificateGroups{}.schema(),
		},
		dataSourceRouter: map[string]tfprotov5.DataSourceServer{
			"pingaccess_trusted_certificate_group": dataPingAccessTrustedCertificateGroups{},
		},
		resourceSchemas: map[string]*tfprotov5.Schema{
			"pingaccess_access_token_validator": resourcePingAccessAccessTokenValidator{}.schema(),
		},
		resourceRouter: map[string]tfprotov5.ResourceServer{
			"pingaccess_access_token_validator": resourcePingAccessAccessTokenValidator{},
		},
	}
}

// Bool is a helper routine that allocates a new bool value
// to store v and returns a pointer to it.
func Bool(v bool) *bool { return &v }

// Int is a helper routine that allocates a new int value
// to store v and returns a pointer to it.
func Int(v int) *int { return &v }

// String is a helper routine that allocates a new string value
// to store v and returns a pointer to it.
func String(v string) *string { return &v }
