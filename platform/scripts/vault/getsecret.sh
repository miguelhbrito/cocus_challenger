#!/usr/bin/env bash

curl -s -H "X-Vault-Token: myroot" -X GET http://localhost:8200/v1/secret/data/cocuschallenge/development | jq -r '.data.data|to_entries|map("\(.key)=\(.value)")|.[]' > platform/app.env
