package joseprovider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	jose "gopkg.in/go-jose/go-jose.v2"
)

func resourceKeyset() *schema.Resource {
	return &schema.Resource{
		Description:   "The resource `jose` generates a JWKS keypair",
		CreateContext: CreateKeyset,
		Read:          schema.Noop,
		Delete:        schema.RemoveFromState,

		Schema: map[string]*schema.Schema{
			"use": {
				Description: "Desired public key usage (use header), one of [enc sig]. Defaults to 'sig'",
				Default:     "sig",
				Optional:    true,
				Type:        schema.TypeString,
				ForceNew:    true,
			},
			"alg": {
				Description: "The algorithm to use. Should be one of " +
					"'" + string(jose.EdDSA) + "', " +
					"'" + string(jose.ES256) + "', " +
					"'" + string(jose.ES384) + "', " +
					"'" + string(jose.ES512) + "', " +
					"'" + string(jose.RS256) + "' (default), " +
					"'" + string(jose.RS384) + "', " +
					"'" + string(jose.RS512) + "', " +
					"'" + string(jose.PS256) + "', " +
					"'" + string(jose.PS384) + "', " +
					"'" + string(jose.PS512) + "' " +
					" for verification or one of " +
					"'" + string(jose.RSA1_5) + "', " +
					"'" + string(jose.RSA_OAEP) + "', " +
					"'" + string(jose.RSA_OAEP_256) + "', " +
					"'" + string(jose.ECDH_ES) + "', " +
					"'" + string(jose.ECDH_ES_A128KW) + "', " +
					"'" + string(jose.ECDH_ES_A192KW) + "', " +
					"'" + string(jose.ECDH_ES_A256KW) + "'" +
					" for encryption.",
				Default:  string(jose.RS256),
				Optional: true,
				Type:     schema.TypeString,
				ForceNew: true,
			},
			"size": {
				Description: "Key size in bits (e.g. 2048 if generating an RSA key). Default is 4096",
				Type:        schema.TypeInt,
				Default:     4096,
				Optional:    true,
				ForceNew:    true,
			},
			"public_key": {
				Description: "Generated public key",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"private_key": {
				Description: "Generated private key",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"public_key_pem": {
				Description: "Generated public key (PEM formatted)",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"private_key_pem": {
				Description: "Generated private key (PEM formatted)",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"id": {
				Description: "Generated key id (kid).",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func CreateKeyset(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	use := d.Get("use").(string)
	alg := d.Get("alg").(string)
	size := d.Get("size").(int)

	pubkey, privkey, kid, err := generateKey(use, alg, size)

	if err != nil {
		return diag.FromErr(err)
	}

	err = d.Set("public_key", pubkey.Json)
	if err != nil {
		return diag.FromErr(err)
	}

	err = d.Set("private_key", privkey.Json)
	if err != nil {
		return diag.FromErr(err)
	}

	err = d.Set("public_key_pem", pubkey.Pem)
	if err != nil {
		return diag.FromErr(err)
	}

	err = d.Set("private_key_pem", privkey.Pem)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(kid)

	return diags
}
