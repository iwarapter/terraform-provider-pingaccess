package main

import (
	"context"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	tf5server "github.com/hashicorp/terraform-plugin-go/tfprotov5/server"
	tfmux "github.com/hashicorp/terraform-plugin-mux"
	"github.com/iwarapter/terraform-provider-pingaccess/pingaccess"
	"log"
	"os"
)

func main() {
	ctx := context.Background()
	sdkv2 := pingaccess.Provider().GRPCProvider
	factory, err := tfmux.NewSchemaServerFactory(ctx, sdkv2)
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
	tf5server.Serve("registry.terraform.io/iwarapter/pingaccess", func() tfprotov5.ProviderServer {
		return factory.Server()
	})
}
