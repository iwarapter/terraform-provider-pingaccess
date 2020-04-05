package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/iwarapter/terraform-provider-pingaccess/pingaccess"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: pingaccess.Provider})
}
