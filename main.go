package main

import (
	"context"

	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/tf5server"
	tfmux "github.com/hashicorp/terraform-plugin-mux/tf5muxserver"
	protocol "github.com/iwarapter/terraform-provider-pingaccess/internal/protocolprovider"
	"github.com/iwarapter/terraform-provider-pingaccess/internal/sdkv2provider"
)

func main() {
	ctx := context.Background()
	sdkv2 := sdkv2provider.Provider().GRPCProvider
	factory, err := tfmux.NewMuxServer(ctx, sdkv2, protocol.Server)
	if err != nil {
		panic(err)
	}
	err = tf5server.Serve("registry.terraform.io/iwarapter/pingaccess", func() tfprotov5.ProviderServer {
		return factory.ProviderServer()
	})
	if err != nil {
		panic(err)
	}
}
