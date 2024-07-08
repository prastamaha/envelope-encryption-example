local_resource(
    'generate_cert',
    'openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout ./dev/vault-key.pem -out ./dev/vault-cert.pem -subj "/CN=localhost"',
    deps=[],
    resource_deps=[],
)

docker_compose("./dev/docker-compose.yaml")

local_resource(
    'unseal_vault',
    './dev/unseal_vault.sh',
    deps=[],
    resource_deps=['vault'],
)
