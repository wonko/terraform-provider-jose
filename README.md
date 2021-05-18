# Terraform Jose Provider

Expose the [go-jose packages](https://github.com/go-jose/go-jose) generator option as a Terraform resource.

All the heavy lifting is done through the re-use of parts of the go-jose package, the Terraform provider
code is a simple wrapper around it. 

I had to copy over the content of parts of the generator submod, as the rename of the mod origin 
[isn't fully done yet](https://github.com/go-jose/go-jose/pull/1). The content of `generate.go` is pretty
much a copy-paste with slight adaptions to fit the provider framework.

Example usage:

```
resource "jose_keypair" "default" {
    use = "sig"
    alg = "RS256"
    size = 2048
}
```
