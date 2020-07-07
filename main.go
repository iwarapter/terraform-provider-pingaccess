package main

import (
	"context"
	"flag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/iwarapter/terraform-provider-pingaccess/pingaccess"
	"log"
)

func main() {
	var debugMode bool

	flag.BoolVar(&debugMode, "debuggable", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	if debugMode {
		err := plugin.Debug(context.Background(), "registry.terraform.io/iwarapter/pingaccess",
			&plugin.ServeOpts{
				ProviderFunc: pingaccess.Provider,
			})
		if err != nil {
			log.Println(err.Error())
		}
	} else {
		plugin.Serve(&plugin.ServeOpts{
			ProviderFunc: pingaccess.Provider})
	}

}
