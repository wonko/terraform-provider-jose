package main

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/wonko/terraform-provider-jose/joseprovider"
)

func main() {
	providerserver.Serve(context.Background(), joseprovider.New, providerserver.ServeOpts{
		Address: "registry.terraform.io/wonko/jose",
	})
}
