package joseprovider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var (
	_ provider.Provider = &joseProvider{}
)

// New is a helper function to simplify provider server and testing implementation.
func New() provider.Provider {
	return &joseProvider{}
}

// hashicupsProvider is the provider implementation.
type joseProvider struct{}

// Metadata returns the provider type name.
func (p *joseProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "jose"
}

// Configure prepares a HashiCups API client for data sources and resources.
func (p *joseProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
}

// DataSources defines the data sources implemented in the provider.
func (p *joseProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return nil
}

// Resources defines the resources implemented in the provider.
func (p *joseProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewKeysetResource,
	}
}

func (p *joseProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{},
	}
}
