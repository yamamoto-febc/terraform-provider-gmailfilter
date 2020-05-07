package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/yamamoto-febc/terraform-provider-gmailfilter/gmailfilter"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: gmailfilter.Provider,
	})
}
