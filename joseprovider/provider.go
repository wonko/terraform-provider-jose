package joseprovider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"jose_keypair": resourceKeypair(),
			"jose_jwt":     resourceJWT(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"jose_expand": dataSourceExpand(),
		},
	}
}
