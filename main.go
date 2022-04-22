package main

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/karelorigin/terraform-provider-local/local"
)

// Generate docs using tfplugindocs
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs

func main() {
	tfsdk.Serve(context.Background(), local.New, tfsdk.ServeOpts{
		Name: "local",
	})
}
