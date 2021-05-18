package joseprovider

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"gopkg.in/square/go-jose.v2"
)

func dataSourceExpand() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceExpandRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Description: "The data sources ID",
				Type:        schema.TypeInt,
				Computed:    true,
			},
			"format": &schema.Schema{
				Description: "The format, either JWS or JWE",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "JWS",
			},
			"payload": &schema.Schema{
				Description: "The payload to expand",
				Type:        schema.TypeString,
				Required:    true,
			},
			"result": &schema.Schema{
				Description: "The expanded payload",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func dataSourceExpandRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	var serialized string
	var err error

	payload := d.Get("payload").(string)
	format := d.Get("format").(string)

	switch format {
	case "JWE":
		var jwe *jose.JSONWebEncryption
		jwe, err = jose.ParseEncrypted(payload)
		if err == nil {
			serialized = jwe.FullSerialize()
		}
	case "JWS":
		var jws *jose.JSONWebSignature
		jws, err = jose.ParseSigned(payload)
		if err == nil {
			serialized = jws.FullSerialize()
		}
	default:
		return diag.FromErr(fmt.Errorf("format should be either JWS or JWE"))
	}

	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("result", serialized)
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
