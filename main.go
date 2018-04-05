package main

import (
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	opts := plugin.ServeOpts{
		ProviderFunc: Provider,
	}
	plugin.Serve(&opts)
}
