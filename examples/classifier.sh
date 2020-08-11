#!/bin/bash

DIR=$(cd `dirname $0` && pwd)
source ${DIR}/shared.sh

PE_CONSOLE=$1

if [[ -z "${PE_CONSOLE}"  ]];then
  echo "Usage: $0 [pe-console-fqdn]"
  exit 2
fi

FACT='{"fact": {"certname": "outward-gown.delivery.puppetlabs.net"}}'     

bold "Aquiring token"

token=$(go run cmd/test/main.go ${PE_CONSOLE} admin compliance | grep Token: | awk '{ print $4 }' | cut -d\" -f2)

green "Token=${token}"

bold "Selecting a node from PDB"

node=$(go run cmd/test/main.go pdb nodes ${PE_CONSOLE} $token | grep Certname | awk '{ print $4 }' | head -1 | cut -d\" -f2)

green "Selecting a classified node from the result=${node}"

bold "Performing the classifier API call on node=${node}"

groups=$(go run cmd/test/main.go classifier node ${PE_CONSOLE} ${node} $token)
echo "Responses=${groups}"



