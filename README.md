# Terraform Jose Provider

Expose the [go-jose package](https://github.com/go-jose/go-jose) generator as a Terraform resource.

Example usage:

```
resource "jose_keyset" "default" {
    use = "sig"
    alg = "RS256"
    size = 2048
}
```

