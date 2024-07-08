storage "file" {
  path = "/vault/data"
}

listener "tcp" {
  address     = "0.0.0.0:8200"
  tls_cert_file = "/vault/config/vault-cert.pem"
  tls_key_file = "/vault/config/vault-key.pem"
}

default_lease_ttl = "168h"
max_lease_ttl     = "720h"
api_addr          = "https://0.0.0.0:8200"
ui                = 1