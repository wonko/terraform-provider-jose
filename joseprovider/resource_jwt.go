package joseprovider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"
)

func resourceJWT() *schema.Resource {
	return &schema.Resource{
		Description:   "The resource `jose_jwt` generates a JWT token",
		CreateContext: CreateJWT,
		Read:          schema.Noop,
		Delete:        schema.RemoveFromState,

		Schema: map[string]*schema.Schema{
			"private_key": {
				Description: "Private key used for signing",
				Required:    true,
				Type:        schema.TypeString,
				ForceNew:    true,
			},
			"alg": {
				Description: "Algorithm used for signing",
				Optional:    true,
				Type:        schema.TypeString,
				ForceNew:    true,
			},
			"claims": &schema.Schema{
				Type:     schema.TypeList,
				ForceNew: true,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"claim_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"issuer": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"subject": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"audience": &schema.Schema{
							Type: schema.TypeSet,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Optional: true,
							ForceNew: true,
						},
						"expiry": &schema.Schema{
							Type: schema.TypeString,
							//							ValidateFunc: validation.IsRFC3339Time,
							Optional: true,
							ForceNew: true,
						},
						"not_before": &schema.Schema{
							Type: schema.TypeString,
							//							ValidateFunc: validation.IsRFC3339Time,
							Optional: true,
							ForceNew: true,
						},
						"issued_at": &schema.Schema{
							Type: schema.TypeString,
							//							ValidateFunc: validation.IsRFC3339Time,
							Optional: true,
							ForceNew: true,
						},
						"private_claims": {
							Description: "Map of claims to include in the token",
							Type:        schema.TypeMap,
							Optional:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							ForceNew: true,
						},
					},
				},
			},
			"jwt": {
				Description: "The generated JWT",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"id": {
				Description: "Token ID.",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func CreateJWT(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	key := d.Get("private_key").(string)
	alg := d.Get("alg").(string)
	var jwk jose.JSONWebKey
	err := jwk.UnmarshalJSON([]byte(key))
	if err != nil {
		return diag.FromErr(err)
	}

	if jwk.IsPublic() {
		return diag.Errorf("given key is a public key")
	}

	salg := jose.SignatureAlgorithm(jwk.Algorithm)
	if alg != "" {
		salg = jose.SignatureAlgorithm(alg)
	}

	sig, err := jose.NewSigner(jose.SigningKey{Algorithm: salg, Key: jwk.Key}, (&jose.SignerOptions{}).WithType("JWT"))
	if err != nil {
		return diag.FromErr(err)
	}

	given_claims := d.Get("claims").([]interface{})[0].(map[string]interface{})

	cl := jwt.Claims{
		Subject: given_claims["subject"].(string),
	}
	raw, err := jwt.Signed(sig).Claims(cl).CompactSerialize()
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(raw)

	return diags
}
