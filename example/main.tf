resource "jose_keyset" "default" {
    # use = "sig"
    # alg = "RS256"
    # size = 2048
}

output "default_keyset" {
  value = jose_keyset.default
}