resource "jose_keypair" "default" {
    use = "sig"
    alg = "RS256"
    size = 2048
}

resource "jose_jwt" "default" {
    private_key = jose_keypair.default.private_key
    alg = "RS512"
    claims {
        subject = "unfreezing fry"
        issued_at = "3000-01-01T00:00:00.000Z"
        audience = ["leela", "bender"]
        private_claims = {
            has_professor = "true"
            executive_director_1 = "matt groening"
            executive_director_2 = "david x cohen"
        }
    }
}

output "default_keypair" {
  value = jose_keypair.default
}

output "default_jwt" {
    value = jose_jwt.default
}