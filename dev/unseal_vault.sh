#!/bin/sh

sleep 10

IS_INITIALIZED=$(docker exec vault vault status -format=json | jq -r '.initialized')

if [ "$IS_INITIALIZED" = "true" ]; then
  echo "Vault is already initialized. check your existing vault_init_keys.txt"
  exit 0
fi

INIT_OUTPUT=$(docker exec vault vault operator init -format=json)
UNSEAL_KEYS=$(echo $INIT_OUTPUT | jq -r '.unseal_keys_b64[]')
ROOT_TOKEN=$(echo $INIT_OUTPUT | jq -r '.root_token')

for KEY in $UNSEAL_KEYS; do
  docker exec vault vault operator unseal $KEY
done

echo "Vault has been initialized and unsealed."
echo "Root Token: $ROOT_TOKEN"

# Optionally, store the unseal keys and root token in a secure location
echo "Unseal Keys: $UNSEAL_KEYS" > ./dev/vault_init_keys.txt
echo "Root Token: $ROOT_TOKEN" >> ./dev/vault_init_keys.txt

echo "Unseal keys and root token have been saved to vault_init_keys.txt"
