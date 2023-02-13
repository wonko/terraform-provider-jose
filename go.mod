module github.com/wonko/terraform-provider-jose

go 1.15

require (
	github.com/go-jose/go-jose/v3 v3.0.0
	github.com/hashicorp/hcl/v2 v2.8.2 // indirect
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.6.1
	github.com/mattn/go-colorable v0.1.8 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	google.golang.org/genproto v0.0.0-20200904004341-0bd0a958aa1d // indirect
	gopkg.in/square/go-jose.v2 v2.5.1
)

//replace github.com/square/go-jose => github.com/go-jose/go-jose v0.0.0-20200630053402-0a67ce9b0693
