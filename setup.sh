export VAULT_ADDR=http://vault.k8s.local
export OPEN_TOKEN=$(vault kv get -format=json kv/switchbot|jq -r '.data.data.TOKEN')
export SECRET_TOKEN=$(vault kv get -format=json kv/switchbot|jq -r '.data.data.CLIENT_SECRET')