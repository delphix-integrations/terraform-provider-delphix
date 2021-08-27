package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	opts := plugin.ServeOpts{
		ProviderFunc: Provider,
	}
	plugin.Serve(&opts)
}
