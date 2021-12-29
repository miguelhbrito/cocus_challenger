#!/usr/bin/env bash
cd $(dirname $0)

curl -d @newsecretpayload.json -H "X-Vault-Token: myroot" -X POST http://localhost:8200/v1/secret/data/cocuschallenge/development