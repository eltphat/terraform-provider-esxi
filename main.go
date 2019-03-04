package main

import (
<<<<<<< HEAD
	"github.com/terraform-providers/terraform-provider-esxi/esxi"
=======
	"github.com/josenk/terraform-provider-esxi/esxi"
>>>>>>> b4c5ff5490318e1efa277b1a2fd1b7e4213cb655
	"github.com/hashicorp/terraform/plugin"
	"github.com/hashicorp/terraform/terraform"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return esxi.Provider()
		},
	})
}
