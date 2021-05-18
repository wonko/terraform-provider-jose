module github.com/wonko/terraform-provider-jose

go 1.15

require (
	github.com/hashicorp/terraform-plugin-sdk v1.17.2
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.6.1
	gopkg.in/square/go-jose.v2 v2.5.1
)

//replace github.com/square/go-jose => github.com/go-jose/go-jose v0.0.0-20200630053402-0a67ce9b0693
